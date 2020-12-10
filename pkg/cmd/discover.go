package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/brutella/dnssd"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

type (
	Data struct {
		Switch     string `json:"switch"`
		Startup    string `json:"startup"`
		Pulse      string `json:"pulse"`
		PulseWidth int    `json:"pulseWidth"`
		RSSI       int    `json:"rssi"`
	}
)

const service = "_ewelink._tcp.local."

func addServiceFunc(w io.Writer) dnssd.AddServiceFunc {
	return func(srv dnssd.Service) {
		data := Data{}
		if err := json.Unmarshal([]byte(strings.ReplaceAll(srv.Text["data1"], "\\", "")), &data); err != nil {
			log.Printf("could not parse data1: %v", err)
		}

		_, _ = fmt.Fprintf(w, "%s:%d\t%s.%s\t%s\t%d dB\n",
			srv.IPs[0].String(), srv.Port, srv.Host, srv.Domain, srv.Text["type"], data.RSSI)
	}
}

func List(c *cli.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.Int("duration")))
	defer cancel()

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, '\t', 0)
	_, _ = fmt.Fprintf(w, "IP:Port\tHost\tType\tSignal\n")
	_, _ = fmt.Fprintf(w, "-------\t----\t----\t------\n")

	err := dnssd.LookupType(ctx, service, addServiceFunc(w), nil)
	if err == context.DeadlineExceeded {
		err = nil
	}
	_ = w.Flush()
	return err
}
