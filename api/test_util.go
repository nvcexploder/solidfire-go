package api

import (
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func getTestClient(t *testing.T) (client *Client) {
	var (
		defaultTarget   = "localhost"
		defaultUsername = "test-username"
		defaultPassword = "supersecret"
		defaultVersion  = "12.3"
		defaultPort     = 443
		defaultTimeout  = 10 * time.Second
	)
	client, err := BuildClient(defaultTarget, defaultUsername, defaultPassword, defaultVersion, defaultPort, defaultTimeout)
	if err != nil {
		require.Fail(t, "Failed to build test client", err)
	}
	return client
}

func activateMock(t *testing.T, c *Client, respBody interface{}) (mockReset func()) {
	httpmock.ActivateNonDefault(c.HTTPClient.GetClient())

	responder, err := httpmock.NewJsonResponder(http.StatusOK, respBody)
	if err != nil {
		require.Fail(t, "Failed to Mock response with ", respBody, err)
	}
	httpmock.RegisterResponder("POST", c.ApiUrl, responder)
	return httpmock.DeactivateAndReset
}

func activateMockHttpErr(c *Client, status int) (mockReset func()) {
	httpmock.ActivateNonDefault(c.HTTPClient.GetClient())

	responder := httpmock.NewStringResponder(status, "not json")
	httpmock.RegisterResponder("POST", c.ApiUrl, responder)
	return httpmock.DeactivateAndReset
}

func buildSFResponseWrapper(resultValue map[string]interface{}) (response SFResponse) {
	response = SFResponse{
		Id:     1,
		Result: resultValue,
	}
	return response
}
