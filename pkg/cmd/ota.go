package cmd

import (
	"net"

	"github.com/urfave/cli/v2"

	"github.com/sam-kamerer/sonoff-diy/pkg/client"
)

func UnlockOTA(c *cli.Context) error {
	cl := client.New(net.ParseIP(c.String("ip")), c.Int("port"), c.String("device-id"))
	return cl.UnlockOTA()
}

func FlashFirmware(c *cli.Context) error {
	cl := client.New(net.ParseIP(c.String("ip")), c.Int("port"), c.String("device-id"))
	return cl.FlashFirmware(c.String("file"))
}
