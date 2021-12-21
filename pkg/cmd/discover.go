package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/grandcat/zeroconf"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
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

const (
	serviceType   = "_ewelink._tcp"
	serviceDomain = "local."
)

func List(c *cli.Context) error {
	service := c.String("serviceType")
	if service == "" {
		service = serviceType
	}
	domain := c.String("serviceDomain")
	if domain == "" {
		domain = serviceDomain
	}

	// Discover all services on the network (e.g. _workstation._tcp)
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		return errors.Wrap(err, "failed to initialize resolver")
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, '\t', 0)
	_, _ = fmt.Fprintf(w, "IP\tPort\tHost\tType\tSignal\n")
	_, _ = fmt.Fprintf(w, "--\t----\t----\t----\t------\n")

	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			printEntry(w, entry)
		}
	}(entries)

	fmt.Printf("Started services discovering...\n\n")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.Int("duration")))
	defer cancel()
	err = resolver.Browse(ctx, service, domain, entries)
	if err != nil {
		return errors.Wrap(err, "failed to browse services")
	}
	<-ctx.Done()
	_ = w.Flush()

	return nil
}

func printEntry(w io.Writer, entry *zeroconf.ServiceEntry) {
	dataValues, err := url.ParseQuery(strings.Join(entry.Text, "&"))
	if err != nil {
		log.Printf("could not parse entry text: %v", err)
	}

	data := Data{}
	if err := json.Unmarshal([]byte(strings.ReplaceAll(dataValues.Get("data1"), "\\", "")), &data); err != nil {
		log.Printf("could not parse data1: %v", err)
	}

	ips := make([]string, 0)
	for _, ipV4 := range entry.AddrIPv4 {
		ips = append(ips, ipV4.String())
	}
	if len(ips) == 0 {
		for _, ipV6 := range entry.AddrIPv6 {
			ips = append(ips, ipV6.String())
		}
	}
	_, _ = fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%d dB\n",
		strings.Join(ips, ", "), entry.Port, entry.ServiceInstanceName(), dataValues.Get("type"), data.RSSI)
}
