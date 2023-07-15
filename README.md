## sonoff-diy

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/sam-kamerer/sonoff-diy?sort=semver)](https://github.com/sam-kamerer/sonoff-diy/releases/latest)
[![GitHub All Releases](https://img.shields.io/github/downloads/sam-kamerer/sonoff-diy/total)](https://github.com/sam-kamerer/sonoff-diy/releases)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/sam-kamerer/sonoff-diy/Release)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sam-kamerer/sonoff-diy)
![GitHub top language](https://img.shields.io/github/languages/top/sam-kamerer/sonoff-diy)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/sam-kamerer/sonoff-diy)

Sonoff-diy is a command-line tool for controlling Sonoff devices in DIY mode. It allows you to discover devices, print device information, switch device state, set power on state, set sleep timer, print WiFi signal strength, set WiFi SSID and password, unlock OTA flashing, and flash firmware over the air.

The tool is easy to use and can be used from the command line. To get started, simply install the tool and then run the `discover` command to discover all of your Sonoff devices. Once you have discovered your devices, you can use the other commands to control them.

For example, to switch the state of a device, you would use the `switch` command. To print the device information, you would use the `device-info` command. And to set the WiFi SSID and password, you would use the `wifi-config` command.

The tool is also very flexible and allows you to customize the behavior of each command. For example, you can set the device IP address, port, and device ID. You can also enable debug mode to get more information about the tool's operation.

Overall, sonoff-diy is a powerful and versatile tool for controlling Sonoff devices in DIY mode. It is easy to use and can be customized to meet your specific needs.

Here are some additional benefits of using sonoff-diy:

* It is free and open source.
* It is cross-platform and can be used on Windows, macOS, and Linux.
* It is constantly being updated with new features and bug fixes.

If you are looking for a powerful and versatile tool for controlling Sonoff devices in DIY mode, then sonoff-diy is the perfect choice for you.

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
