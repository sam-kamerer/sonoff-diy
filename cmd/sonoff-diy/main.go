package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/sam-kamerer/sonoff-diy/pkg/cmd"
	"github.com/sam-kamerer/sonoff-diy/pkg/vars"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	compiled, _ := time.Parse(time.RFC3339, date)
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "Version:\t%s\nBuild at:\t%s\nCommit:\t\t%s\n", c.App.Version, date, commit)
	}
	app := &cli.App{
		Name:     "sonoff-diy",
		Version:  version,
		Compiled: compiled,
		Usage:    "The tool for work with Sonoff devices (Basic R3/RFR3/Mini) in the DIY MODE",
		Authors: []*cli.Author{{
			Name:  "Andrey Semerun",
			Email: "andrey.semerun@gmail.com",
		}},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "ip",
				Usage:   "sets the device ip address",
				EnvVars: nil,
				Value:   "127.0.0.1",
			},
			&cli.IntFlag{
				Name:    "port",
				Usage:   "sets the device port",
				EnvVars: nil,
				Value:   8081,
			},
			&cli.StringFlag{
				Name:    "device-id",
				Usage:   "sets the device id (optional)",
				EnvVars: nil,
			},
			&cli.BoolFlag{
				Name:    "debug",
				Usage:   "sets the debug mode on",
				EnvVars: nil,
			},
		},
		Before: func(c *cli.Context) error {
			if c.Bool("debug") {
				vars.Debug = true
			}
			return nil
		},
		Commands: []*cli.Command{{
			Name:    "discover",
			Aliases: []string{"d"},
			Usage:   "Discover devices",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:    "duration",
					Aliases: []string{"d"},
					Usage:   "sets the duration of device discovery in seconds",
					Value:   10,
				},
			},
			Action: cmd.List,
		}, {
			Name:    "device-info",
			Aliases: []string{"di"},
			Usage:   "Print specified device info",
			Action:  cmd.DeviceInfo,
		}, {
			Name:    "switch",
			Aliases: []string{"sw"},
			Usage:   "Switch device state",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "state",
					Aliases: []string{"s"},
					Usage:   "sets the device state (ex: on or off)",
				},
			},
			Action: cmd.Switch,
		}, {
			Name:    "power-on-state",
			Aliases: []string{"pos"},
			Usage:   "Sets power on device state",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "state",
					Aliases:  []string{"s"},
					Usage:    "sets the power on device state (ex: on, off or stay)",
					Required: true,
				},
			},
			Action: cmd.PowerOnState,
		}, {
			Name:    "sleep-timer",
			Aliases: []string{"st"},
			Usage:   "Sets sleep timer",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "state",
					Aliases:  []string{"s"},
					Usage:    "sets the device sleep timer state (ex: on or off)",
					Required: true,
				},
				&cli.IntFlag{
					Name:    "duration",
					Aliases: []string{"d"},
					Usage:   "sets timer duration in seconds",
					Value:   10,
				},
			},
			Action: cmd.SleepTimer,
		}, {
			Name:    "wifi-signal",
			Aliases: []string{"wfs"},
			Usage:   "Print the WiFi signal strength of the specified device",
			Action:  cmd.WiFiSignal,
		}, {
			Name:    "wifi-config",
			Aliases: []string{"wfc"},
			Usage:   "Sets the WiFi SSID and password for specified device",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "ssid",
					Aliases:  []string{"s"},
					Usage:    "sets the WiFi network SSID",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "password",
					Aliases:  []string{"p"},
					Usage:    "sets the WiFi network password",
					Required: true,
				},
			},
			Action: cmd.WiFiConfig,
		}, {
			Name:    "unlock-ota",
			Aliases: []string{"uo"},
			Usage:   "Unlocks ability for flash firmware over the air for the specified device",
			Action:  cmd.UnlockOTA,
		}, {
			Name:    "flash-firmware",
			Aliases: []string{"ff"},
			Usage:   "Flashing the firmware over the air to the specified device",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "file",
					Aliases:  []string{"f"},
					Usage:    "sets the file path to the firmware",
					Required: true,
				},
			},
			Action: cmd.FlashFirmware,
		}},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
