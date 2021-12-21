package client

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/schollz/progressbar/v2"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type ProgressBar struct {
	Bar *progressbar.ProgressBar
}

func (b *ProgressBar) Write(p []byte) (int, error) {
	n := len(p)
	_ = b.Bar.Add(n)
	return n, nil
}

func serveFile(filePath string, endCh chan struct{}) (string, string, error) {
	fw, err := os.Open(filePath)
	if err != nil {
		return "", "", err
	}
	defer fw.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, fw)
	if err != nil {
		return "", "", err
	}
	hashSum := fmt.Sprintf("%x", hash.Sum(nil))
	tmpFile, err := os.CreateTemp("sonoff-diy", "firmware-*")
	if err != nil {
		return "", "", err
	}
	defer tmpFile.Close()

	_, err = fw.Seek(0, io.SeekStart)
	if err != nil {
		return hashSum, "", err
	}
	_, err = io.Copy(tmpFile, fw)
	if err != nil {
		return hashSum, "", err
	}
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return hashSum, "", err
	}
	fwFileName := filepath.Base(tmpFile.Name())
	go func(l net.Listener) {
		defer func() {
			endCh <- struct{}{}
		}()
		respEndCh := make(chan struct{}, 1)
		server := &http.Server{}
		http.Handle("/"+fwFileName, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			f, err := os.Open(tmpFile.Name())
			if err != nil {
				log.Println(err)
				w.WriteHeader(500)
				return
			}
			defer f.Close()

			fStat, _ := f.Stat()
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Header().Set("Content-Disposition", `attachment; filename="`+fwFileName+`"`)
			c, err := io.Copy(w, io.TeeReader(f, &ProgressBar{Bar: progressbar.New64(fStat.Size())}))
			fmt.Println("")
			if err != nil {
				log.Println(err)
				w.WriteHeader(500)
				return
			}
			w.Header().Set("Content-Length", strconv.Itoa(int(c)))
			respEndCh <- struct{}{}
		}))
		go func() {
			<-respEndCh
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			_ = server.Shutdown(ctx)
		}()
		if err = server.Serve(l); err != nil && err != http.ErrServerClosed {
			log.Println(err)
		}
		_ = l.Close()
	}(l)
	return hashSum, fmt.Sprintf("http://%s/%s", l.Addr().String(), fwFileName), err
}
