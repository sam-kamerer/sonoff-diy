## sonoff-diy

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/sam-kamerer/sonoff-diy?sort=semver)](https://github.com/sam-kamerer/sonoff-diy/releases/latest)
[![GitHub All Releases](https://img.shields.io/github/downloads/sam-kamerer/sonoff-diy/total)](https://github.com/sam-kamerer/sonoff-diy/releases)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/sam-kamerer/sonoff-diy/Release)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sam-kamerer/sonoff-diy)
![GitHub top language](https://img.shields.io/github/languages/top/sam-kamerer/sonoff-diy)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/sam-kamerer/sonoff-diy)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/sam-kamerer/sonoff-diy?sort=semver)](https://github.com/sam-kamerer/sonoff-diy/releases/latest)

The tool for work with Sonoff devices (Basic R3/RFR3/Mini) in the DIY MODE.

```
$ sonoff-diy --help

NAME:
   sonoff-diy - The tool for work with Sonoff devices (Basic R3/RFR3/Mini) in the DIY MODE

USAGE:
   sonoff-diy [global options] command [command options] [arguments...]

VERSION:
   v1.1.1

COMMANDS:
   discover, d          Discover devices
   device-info, di      Print specified device info
   switch, sw           Switch device state
   power-on-state, pos  Sets power on device state
   sleep-timer, st      Sets sleep timer
   wifi-signal, wfs     Print the WiFi signal strength of the specified device
   wifi-config, wfc     Sets the WiFi SSID and password for specified device
   unlock-ota, uo       Unlocks ability for flash firmware over the air for the specified device
   flash-firmware, ff   Flashing the firmware over the air to the specified device
   help, h              Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --ip value         sets the device ip address (default: "127.0.0.1")
   --port value       sets the device port (default: 8081)
   --device-id value  sets the device id (optional)
   --debug            sets the debug mode on (default: false)
   --help, -h         show help (default: false)
   --version, -v      print the version (default: false)
```
