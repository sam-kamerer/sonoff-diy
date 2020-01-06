package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sam-kamerer/sonoff-diy/pkg/vars"
	"io/ioutil"
	"log"
	"net/http"
)

const requestContentType = "application/json"

type (
	Response struct {
		// The order of device status update (also the order of TXT Record update)
		Seq int `json:"seq"`

		// Whether the device successfully sets the specified device information.
		// - 0:   Successfully
		// - 400: The operation failed and the request was formatted incorrectly.
		//        The request body is not a valid JSON format.
		// - 401: The operation failed and the request was unauthorized.
		//        Device information encryption is enabled on the device, but the request is not encrypted.
		// - 404: The operation failed and the device does not exist.
		//        The device does not support the requested deviceid.
		// - 422: The operation failed and the request parameters are invalid.
		//        For example, the device does not support setting specific device information.
		Error int `json:"error"`

		// Object type, it returns specific device info when check the device information
		Data json.RawMessage `json:"data"`
	}
	ResponseError struct {
		code int
		msg  string
	}
)

func (e ResponseError) Error() string {
	return e.msg
}

func (e ResponseError) Code() int {
	return e.code
}

func request(url string, data []byte) (*Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", requestContentType)

	if vars.Debug {
		log.Printf("request url: %s", url)
		log.Printf("request body: %v", string(data))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if vars.Debug {
		log.Printf("response status: %v", resp.Status)
		log.Println("response headers:")
		jsonString, err := json.MarshalIndent(resp.Header, "", "  ")
		if err != nil {
			log.Printf("error: %v", err)
		}
		log.Printf("%s", jsonString)
		log.Printf("response body:\n%v", string(body))
	}

	if resp.StatusCode >= 400 {
		respErr := ResponseError{code: resp.StatusCode}
		switch resp.StatusCode {
		case 400:
			respErr.msg = fmt.Sprintf("the request body is not a valid JSON format: %s", string(data))
		case 401:
			respErr.msg = "device information encryption is enabled on the device, but the request is not encrypted"
		case 404:
			respErr.msg = "the device does not support the requested deviceid"
		case 422:
			respErr.msg = "the request parameters are invalid"
		default:
			respErr.msg = fmt.Sprintf("undefined error: %d", resp.StatusCode)
		}
		return nil, respErr
	}

	res := &Response{}
	return res, json.Unmarshal(body, res)
}

func (r Response) UnmarshalData(v interface{}) error {
	if len(r.Data) < 2 {
		return nil
	}
	d := bytes.ReplaceAll(r.Data, []byte{0x5c}, []byte{})
	d = d[1 : len(d)-1]
	return json.Unmarshal(d, v)
}
