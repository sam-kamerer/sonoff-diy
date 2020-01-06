package cmd

import (
	"fmt"
	"github.com/sam-kamerer/sonoff-diy/pkg/client"
	"github.com/urfave/cli/v2"
	"net"
)

func WiFiSignal(c *cli.Context) error {
	cl := client.New(net.ParseIP(c.String("ip")), c.Int("port"), c.String("device-id"))
	s, err := cl.WiFiSignal()
	if err != nil {
		return err
	}
	fmt.Printf("Signal strength: %d dB\n", s)
	return nil
}

func WiFiConfig(c *cli.Context) error {
	cl := client.New(net.ParseIP(c.String("ip")), c.Int("port"), c.String("device-id"))
	return cl.WiFiConfig(c.String("ssid"), c.String("password"))
}