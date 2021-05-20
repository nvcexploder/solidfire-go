package api

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
)

type Client struct {
	target       string
	Port         int
	RequestCount int64
	credentials  Credentials
	Timeout      int
	Version      string
	ApiUrl       string
	Name         string
	RestyClient  *resty.Client
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

func (e *SFAPIError) Error() string {
	return fmt.Sprintf("%d : %s - %s", e.Code, e.Name, e.Message)
}

func BuildClient(target string, username string, password string, version string, port int, timeoutSecs int) (c *Client, err error) {
	// sanity check inputs
	if target == "" {
		err = errors.New("Unable to issue json-rpc requests without specifying Target")
		return nil, err
	}
	if username == "" || password == "" {
		err = errors.New("Unable to issue json-rpc requests without specifying Credentials")
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
		target:      target,
		ApiUrl:      apiUrl,
		credentials: creds,
		Version:     version,
		Port:        port,
		RestyClient: r,
		Timeout:     timeoutSecs}
	return SFClient, nil
}

func (c *Client) request(ctx context.Context, method string, params interface{}, result interface{}) (err error) {
	sfr := SFResponse{}
	_, err = c.RestyClient.R().
		SetBody(map[string]interface{}{
			// TODO: Investigate replacing id w/ something concurrency safe
			"id":     c.RequestCount,
			"method": method,
			"params": params,
		}).
		SetResult(&sfr).
		Post(c.ApiUrl)
	if err != nil {
		return err
	}
	if sfr.Error.Code != 0 {
		return errors.New(sfr.Error.Error())
	}
	if err = mapstructure.Decode(sfr.Result, &result); err != nil {
		return err
	}
	return nil
}
