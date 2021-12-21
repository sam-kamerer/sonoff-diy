package cmd

import (
	"net"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/sam-kamerer/sonoff-diy/pkg/client"
)

func PowerOnState(c *cli.Context) error {
	cl := client.New(net.ParseIP(c.String("ip")), c.Int("port"), c.String("device-id"))
	var state client.PowerOnState
	s := c.String("state")
	switch s {
	case "on":
		state = client.PowerOnStateOn
	case "off":
		state = client.PowerOnStateOff
	case "stay":
		state = client.PowerOnStateStay
	default:
		return errors.Errorf("invalid power on state: %s", s)
	}
	return cl.PowerOnState(state)
}

func Switch(c *cli.Context) error {
	cl := client.New(net.ParseIP(c.String("ip")), c.Int("port"), c.String("device-id"))
	return cl.Switch(client.State(c.String("state")))
}

func SleepTimer(c *cli.Context) error {
	cl := client.New(net.ParseIP(c.String("ip")), c.Int("port"), c.String("device-id"))
	return cl.SleepTimer(client.State(c.String("state")), c.Int("duration"))
}
