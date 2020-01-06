package cmd

import (
	"encoding/json"
	"github.com/sam-kamerer/sonoff-diy/pkg/client"
	"github.com/urfave/cli/v2"
	"net"
	"os"
)

func DeviceInfo(c *cli.Context) error {
	cl := client.New(net.ParseIP(c.String("ip")), c.Int("port"), c.String("device-id"))
	di, err := cl.DeviceInfo()
	if err != nil {
		return err
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(di)
}
