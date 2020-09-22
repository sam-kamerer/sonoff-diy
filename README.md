## sonoff-diy
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fsam-kamerer%2Fsonoff-diy.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fsam-kamerer%2Fsonoff-diy?ref=badge_shield)

The tool for work with Sonoff devices (Basic R3/RFR3/Mini) in the DIY MODE.

```
$ sonoff-diy --help

NAME:
   sonoff-diy - The tool for work with Sonoff devices (Basic R3/RFR3/Mini) in the DIY MODE

USAGE:
   sonoff-diy [global options] command [command options] [arguments...]

VERSION:
   v1.0.0

AUTHOR:
   Andrey Semerun <andrey.semerun@gmail.com>

COMMANDS:
   discover, d          Discover devices
   device-info, di      Print specified device info
   switch, sw           Switch device state
   power-on-state, pos  Sets power on device state
   pulsate, pu          Sets pulsing state
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

## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fsam-kamerer%2Fsonoff-diy.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fsam-kamerer%2Fsonoff-diy?ref=badge_large)