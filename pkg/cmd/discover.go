package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/brutella/dnssd"
	"github.com/urfave/cli/v2"
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

func List(c *cli.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.Int("duration")))
	defer cancel()

	addFn := func(srv dnssd.Service) {
		data := &Data{}
		if err := json.Unmarshal([]byte(strings.ReplaceAll(srv.Text["data1"], "\\", "")), data); err != nil {
			log.Printf("could not parse data1: %v", err)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, '\t', 0)
		_, _ = fmt.Fprintf(w, "IP:PORT\tHost\tType\tSignal\n")
		_, _ = fmt.Fprintf(w, "-------\t----\t----\t------\n")
		_, _ = fmt.Fprintf(w, "%s:%d\t%s.%s\t%s\t%d dB\n",
			srv.IPs[0].String(), srv.Port, srv.Host, srv.Domain, srv.Text["type"], data.RSSI)
		_ = w.Flush()
	}

	err := dnssd.LookupType(ctx, service, addFn, nil)
	if err == context.DeadlineExceeded {
		err = nil
	}
	return err
}
