package api

import (
	"net/http"
	"testing"

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
		defaultTimeout  = 10
	)
	client, err := BuildClient(defaultTarget, defaultUsername, defaultPassword, defaultVersion, defaultPort, defaultTimeout)
	if err != nil {
		require.Fail(t, "Failed to build test client", err)
	}
	return client
}

func activateMock(t *testing.T, c *Client, respBody interface{}) (mockReset func()) {
	httpmock.ActivateNonDefault(c.HTTPClient.GetClient())
	mockReset = func() {
		httpmock.DeactivateAndReset()
	}

	responder, err := httpmock.NewJsonResponder(http.StatusOK, respBody)
	if err != nil {
		require.Fail(t, "Failed to Mock response with ", respBody, err)
	}
	httpmock.RegisterResponder("POST", c.ApiUrl, responder)
	return mockReset
}

func buildSFResponseWrapper(resultValue map[string]interface{}) (response SFResponse) {
	response = SFResponse{
		Id:     1,
		Result: resultValue,
	}
	return response
}