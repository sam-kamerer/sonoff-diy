package client

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net"
)

const (
	PowerOnStateOn   PowerOnState = "on"
	PowerOnStateOff  PowerOnState = "off"
	PowerOnStateStay PowerOnState = "stay"
)

type (
	PowerOnState string
	Client       struct {
		ip       net.IP
		port     int
		deviceId string
	}
	DeviceInfoData struct {
		Switch     string `json:"switch"`
		Startup    string `json:"startup"`
		Pulse      string `json:"pulse"`
		PulseWidth int    `json:"pulseWidth"`
		SSID       string `json:"ssid"`
		OTAUnlock  bool   `json:"otaUnlock"`
	}
	WiFiSignalData struct {
		SignalStrength int `json:"signalStrength"`
	}
)

func New(ip net.IP, port int, deviceId string) *Client {
	return &Client{
		ip:       ip,
		port:     port,
		deviceId: deviceId,
	}
}

func (c Client) url(path string) string {
	return fmt.Sprintf("http://%s:%d%s", c.ip, c.port, path)
}

func (c Client) DeviceInfo() (*DeviceInfoData, error) {
	data, err := json.Marshal(map[string]interface{}{
		"deviceid": c.deviceId,
		"data":     struct{}{},
	})
	if err != nil {
		return nil, err
	}
	resp, err := request(c.url("/zeroconf/info"), data)
	if err != nil {
		return nil, err
	}
	di := &DeviceInfoData{}
	return di, resp.UnmarshalData(di)
}

func (c Client) Switch(on bool) error {
	v := "off"
	if on {
		v = "on"
	}
	data, err := json.Marshal(map[string]interface{}{
		"deviceid": c.deviceId,
		"data":     map[string]interface{}{"switch": v},
	})
	if err != nil {
		return err
	}
	_, err = request(c.url("/zeroconf/switch"), data)
	return err
}

func (c Client) PowerOnState(state PowerOnState) error {
	data, err := json.Marshal(map[string]interface{}{
		"deviceid": c.deviceId,
		"data":     map[string]interface{}{"startup": state},
	})
	if err != nil {
		return err
	}
	_, err = request(c.url("/zeroconf/startup"), data)
	return err
}

func (c Client) Pulsate(on bool, pulseDurationMs int) error {
	v := "off"
	if on {
		v = "on"
	}
	d := map[string]interface{}{"pulse": v}

	if on {
		if pulseDurationMs < 500 {
			pulseDurationMs = 500
		} else if pulseDurationMs > 36000000 {
			pulseDurationMs = 36000000
		}
		d["pulseWidth"] = int(float64(pulseDurationMs)/500.0) * 500
	}
	data, err := json.Marshal(map[string]interface{}{
		"deviceid": c.deviceId,
		"data":     d,
	})
	if err != nil {
		return err
	}
	_, err = request(c.url("/zeroconf/pulse"), data)
	return err
}

func (c Client) WiFiSignal() (int, error) {
	data, err := json.Marshal(map[string]interface{}{
		"deviceid": c.deviceId,
		"data":     struct{}{},
	})
	if err != nil {
		return 0, err
	}
	resp, err := request(c.url("/zeroconf/signal_strength"), data)
	if err != nil {
		return 0, err
	}
	d := &WiFiSignalData{}
	return d.SignalStrength, resp.UnmarshalData(d)
}

func (c Client) WiFiConfig(ssid, password string) error {
	data, err := json.Marshal(map[string]interface{}{
		"deviceid": c.deviceId,
		"data": map[string]interface{}{
			"ssid":     ssid,
			"password": password,
		},
	})
	if err != nil {
		return err
	}
	_, err = request(c.url("/zeroconf/wifi"), data)
	return err
}

func (c Client) UnlockOTA() error {
	data, err := json.Marshal(map[string]interface{}{
		"deviceid": c.deviceId,
		"data":     struct{}{},
	})
	if err != nil {
		return err
	}
	_, err = request(c.url("/zeroconf/ota_unlock"), data)
	return err
}

func (c Client) FlashFirmware(filePath string) error {
	endCh := make(chan struct{}, 1)
	hash, url, err := serveFile(filePath, endCh)
	if err != nil {
		return err
	}
	data, err := json.Marshal(map[string]interface{}{
		"deviceid": c.deviceId,
		"data": map[string]interface{}{
			"downloadUrl": url,
			"sha256sum":   hash,
		},
	})
	if err != nil {
		return err
	}
	_, err = request(c.url("/zeroconf/ota_flash"), data)
	if err == nil {
		return nil
	}

	if e, ok := err.(ResponseError); ok {
		switch e.Code() {
		case 403:
			err = errors.New("the OTA function was not unlocked, the interface '3.2.6OTA function unlocking' must be successfully called first")
		case 408:
			err = errors.New("the pre-download firmware timed out, you can try to call this interface again after optimizing the network environment or increasing the network speed")
		case 413:
			err = errors.New("the request body size is too large, the size of the new OTA firmware exceeds the firmware size limit allowed by the device")
		case 424:
			err = errors.New("the firmware could not be downloaded, the URL address is unreachable (IP address is unreachable, HTTP protocol is unreachable, firmware does not exist, server does not support Range request header, etc.)")
		case 471:
			err = errors.New("the firmware integrity check failed, the SHA256 checksum of the downloaded new firmware does not match the value of the request body's sha256sum field. Restarting the device will cause bricking issue")
		}
		return err
	}
	<-endCh
	return err
}
