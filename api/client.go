package api

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

type Client struct {
	Target       string
	Port         int
	RequestCount int64
	Timeout      int
	Version      string
	ApiUrl       string
	Name         string
	HTTPClient   *resty.Client
}

type Credentials struct {
	username string
	password string
}

type SFResponse struct {
	Id     int32                  `json:"id"`
	Result map[string]interface{} `json:"result"`
	Error  SFAPIError             `json:"error"`
}

type SFAPIError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Name    string `json:"name"`
}

type SFError interface {
	GetName() string
	GetMessage() string
}

type RequestError struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("%s : %s", e.GetName(), e.GetMessage())
}
func (e *RequestError) GetName() string    { return e.Name }
func (e *RequestError) GetMessage() string { return e.Message }

type ServiceError struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("%s : %s", e.Name, e.Message)
}
func (e *ServiceError) GetName() string    { return e.Name }
func (e *ServiceError) GetMessage() string { return e.Message }

type ResourceNotFoundError struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

func (e *ResourceNotFoundError) Error() string {
	return fmt.Sprintf("%s : %s", e.Name, e.Message)
}
func (e *ResourceNotFoundError) GetName() string    { return e.Name }
func (e *ResourceNotFoundError) GetMessage() string { return e.Message }

const (
	ErrNoTarget                        = "Client requires a valid target"
	ErrNoCredentials                   = "Client requires a valid username and password"
	ErrInvalidCredentials              = "Provided credentials are invalid"
	ErrUnexpectedServerError           = "Unexpected server error"
	ErrVolumeIDDoesNotExist            = "xVolumeIDDoesNotExist"
	ErrSnapshotIDDoesNotExist          = "xSnapshotIDDoesNotExist"
	ErrAccountIDDoesNotExist           = "xAccountIDDoesNotExist"
	ErrQoSPolicyDoesNotExist           = "xQoSPolicyDoesNotExist"
	ErrVolumeAccessGroupIDDoesNotExist = "xVolumeAccessGroupIdDoesNotExist"
	ErrInitiatorDoesNotExist           = "xInitiatorDoesNotExist"
	ErrInitiatorExists                 = "xInitiatorExists"
	ErrExceededLimit                   = "xExceededLimit"
	ErrUnrecognizedEnumString          = "xUnrecognizedEnumString"
	ErrInvalidAPIParameter             = "xInvalidAPIParameter"
	ErrInvalidParameter                = "xInvalidParameter"
	ErrInvalidParameterType            = "xInvalidParameterType"
	ErrMVIPNotPaired                   = "xMVIPNotPaired"
)

func BuildClient(target string, username string, password string, version string, port int, timeoutSecs int) (c *Client, err error) {
	// sanity check inputs
	if target == "" {
		err = errors.New(ErrNoTarget)
		return nil, err
	}
	if username == "" || password == "" {
		err = errors.New(ErrNoCredentials)
		return nil, err
	}
	if port == 0 {
		port = 443
	}
	if timeoutSecs == 0 {
		timeoutSecs = 40
	}
	if version == "" {
		version = "12.3"
	}
	creds := Credentials{username: username, password: password}
	apiUrl := fmt.Sprintf("https://%s:%d/json-rpc/%s", target, port, version)
	r := resty.New().
		SetHeader("Accept", "application/json").
		SetBasicAuth(creds.username, creds.password).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	SFClient := &Client{
		Target:     target,
		ApiUrl:     apiUrl,
		Version:    version,
		Port:       port,
		HTTPClient: r,
		Timeout:    timeoutSecs}
	return SFClient, nil
}

func BuildRequestError(name string, message string) *RequestError {
	return &RequestError{
		Name:    name,
		Message: message,
	}
}

func (c *Client) request(ctx context.Context, method string, params interface{}, result interface{}) (err error) {
	sfr := SFResponse{}
	response, err := c.HTTPClient.R().
		SetBody(map[string]interface{}{
			"id":     c.RequestCount,
			"method": method,
			"params": params,
		}).
		SetResult(&sfr).
		Post(c.ApiUrl)
	c.RequestCount++
	if err != nil {
		return err
	}

	// resty documentation isn't explicit about this but it seems like `err` is for connection/transport
	// errors and http protocol errors are provided in the *resty.Response.  Most of the time protocol
	// errors won't be expected as those usually become json-rpc errors.  However, there are some edge
	// cases where HTTP errors are expected and no json object is available
	if response.IsError() {
		// auth errors return a 401 and html instead of json
		if response.StatusCode() == 401 {
			return BuildRequestError(ErrInvalidCredentials, response.Status())
		}
		// making a bit of an assumption here that any other kind of http error will also not include
		// json so we'll just return something generic.  Unfortunately, resty doesn't tell us if
		// if the response wasn't able to be unmarshaled so our SFResponse would appear to have
		// a 0 error code if we attempted to check the result object in the case where we got something
		// other than the expected json schema.
		return &ServiceError{
			Name:    ErrUnexpectedServerError,
			Message: response.Status(),
		}
	}

	if sfr.Error.Code != 0 {
		switch sfr.Error.Name {
		case ErrVolumeIDDoesNotExist, ErrSnapshotIDDoesNotExist, ErrAccountIDDoesNotExist,
			ErrQoSPolicyDoesNotExist, ErrVolumeAccessGroupIDDoesNotExist, ErrInitiatorDoesNotExist:
			return &ResourceNotFoundError{
				Name:    sfr.Error.Name,
				Message: sfr.Error.Message,
			}
		case ErrExceededLimit, ErrUnrecognizedEnumString, ErrInvalidAPIParameter,
			ErrInvalidParameter, ErrInvalidParameterType, ErrInitiatorExists, ErrMVIPNotPaired:
			return &RequestError{
				Name:    sfr.Error.Name,
				Message: sfr.Error.Message,
			}
		default:
			return &ServiceError{
				Name:    sfr.Error.Name,
				Message: sfr.Error.Message,
			}
		}
	}
	if err = mapstructure.Decode(sfr.Result, &result); err != nil {
		return err
	}
	return nil
}
