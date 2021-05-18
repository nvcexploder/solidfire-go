package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
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
}

type Credentials struct {
	username string
	password string
}

type APIError struct {
	Id    int `json:"id"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Name    string `json:"name"`
	} `json:"error"`
}

/*

   func Create(target string, username string, password string, version string, port int, timeoutSecs int) (c *Client, err error) {
	client, err := NewFromOpts(target, username, password, ver, port, timeoutSecs)
	if err != nil {
		log.Errorf("Err: %v", err)
		return nil, err
	}

	// set to current version
	getApiResult, err := client.GetAPI()
	if err != nil {
		log.Errorf("Error retrieving Cluster version: %v", err)
		return nil, err
	}
	client.Version = strconv.FormatFloat(getApiResult.CurrentVersion, 'f', 1, 64)

	if client.Port == 443 {
		getClusterInfoResult, err := client.GetClusterInfo()
		if err != nil {
			log.Errorf("Error retrieving Cluster name: %v", err)
			return nil, err
		}
		client.Name = getClusterInfoResult.ClusterInfo.Name
	} else if client.Port == 442 {
		getConfigResult, err := client.GetConfig()
		if err != nil {
			log.Errorf("Error retrieving Node name: %v", err)
			return nil, err
		}
		client.Name = getConfigResult.Config.Cluster.Name
	}
	return client, err
    }
*/

func BuildClient(target string, username string, password string, version string, port int, timeoutSecs int) (c *Client, err error) {
	// sanity check inputs
	if target == "" {
		log.Error("Target is not set, unable to issue requests")
		err = errors.New("Unable to issue json-rpc requests without specifying Target")
		return nil, err
	}
	if username == "" || password == "" {
		log.Error("Credentials are not set, unable to issue requests")
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
	SFClient := &Client{
		target:      target,
		ApiUrl:      apiUrl,
		credentials: creds,
		Version:     version,
		Port:        port,
		Timeout:     timeoutSecs}
	return SFClient, nil
}

func (c *Client) SendRequest(method string, params interface{}) (response map[string]interface{}, err error) {
	// increment the request counter
	c.RequestCount++

	// create the request json
	data, err := json.Marshal(map[string]interface{}{
		"method": method,
		"id":     c.RequestCount,
		"params": params,
	})

	// create the http client with proper settings
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true},
		},
	}

	// make the request
	req, err := http.NewRequest("POST", c.ApiUrl, strings.NewReader(string(data)))
	req.SetBasicAuth(c.credentials.username, c.credentials.password)
	log.Debugf("Request: %+v", string(data))
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error encountered posting request: %v", err)
		return nil, err
	}

	// read the response into a byte array
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Debugf("Response: %+v", string(body))
	if err != nil {
		log.Errorf("Error reading response body: %v", err)
		return nil, err
	}

	// decode the response into a map
	r := bytes.NewReader(body)
	var result map[string]interface{}
	err = json.NewDecoder(r).Decode(&result)
	if err != nil {
		log.Errorf("Error decoding response into map: %v", err)
		return nil, err
	}

	// check for errors
	errresp := APIError{}
	err = json.Unmarshal([]byte(body), &errresp)
	if err != nil {
		log.Errorf("Error unmarshalling response: %v", err)
		return nil, err
	}
	if errresp.Error.Code != 0 {
		err = errors.New("Received error response from API request: " + errresp.Error.Message)
		return nil, err
	}

	// return successful response
	return result, nil
}
