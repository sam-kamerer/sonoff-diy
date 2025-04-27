package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const requestContentType = "application/json"

const (
	// HTTPStatusBadRequest indicates that the request was invalid or cannot be otherwise processed.
	HTTPStatusBadRequest = 400
	// HTTPStatusUnauthorized indicates that the request has not been applied because it lacks valid authentication credentials for the targeting resource.
	HTTPStatusUnauthorized = 401
	// HTTPStatusNotFound indicates that the target resource does not exist.
	HTTPStatusNotFound = 404
	// HTTPStatusUnprocessableEntity indicates that the server understands the content type of the request entity, and the syntax of the request entity is correct but was unable to process the contained instructions.
	HTTPStatusUnprocessableEntity = 422
)

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

func sendRequest(url string, data []byte) (*Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", requestContentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return handleResponse(resp)
}

func handleResponse(resp *http.Response) (*Response, error) {
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return createErrorResponse(resp.StatusCode, body)
	}

	res := &Response{}
	return res, json.Unmarshal(body, res)
}

func createErrorResponse(statusCode int, body []byte) (*Response, error) {
	respErr := ResponseError{code: statusCode}
	switch statusCode {
	case HTTPStatusBadRequest:
		respErr.msg = fmt.Sprintf("the request body is not a valid JSON format: %s", string(body))
	case HTTPStatusUnauthorized:
		respErr.msg = "device information encryption is enabled on the device, but the request is not encrypted"
	case HTTPStatusNotFound:
		respErr.msg = "the device does not support the requested deviceid"
	case HTTPStatusUnprocessableEntity:
		respErr.msg = "the request parameters are invalid"
	default:
		respErr.msg = fmt.Sprintf("undefined error: %d", statusCode)
	}
	return nil, respErr
}

func (r Response) UnmarshalData(v interface{}) error {
	if len(r.Data) < 2 {
		return nil
	}
	d := bytes.ReplaceAll(r.Data, []byte{0x5c}, []byte{})
	d = d[1 : len(d)-1]
	return json.Unmarshal(d, v)
}
