package api_test

import (
	"fmt"
	"testing"

	"github.com/joyent/solidfire-sdk/api"
	"github.com/stretchr/testify/require"
)

var (
	defaultTarget   = "localhost"
	defaultUsername = "test-username"
	defaultPassword = "supersecret"
	defaultVersion  = "12.3"
	defaultPort     = 443
	defaultTimeout  = 10
)

func TestBuldClientErrors(t *testing.T) {
	var err error
	_, err = api.BuildClient("", defaultUsername, defaultPassword, defaultVersion, defaultPort, defaultTimeout)
	require.Equal(t, err.Error(), "Client requires a valid target")
	_, err = api.BuildClient(defaultTarget, "", defaultPassword, defaultVersion, defaultPort, defaultTimeout)
	require.Equal(t, err.Error(), "Client requires a valid username and password")
	_, err = api.BuildClient(defaultTarget, defaultUsername, "", defaultVersion, defaultPort, defaultTimeout)
	require.Equal(t, err.Error(), "Client requires a valid username and password")
}

func TestBuldClient(t *testing.T) {
	c, err := api.BuildClient(defaultTarget, defaultUsername, defaultPassword, defaultVersion, 443, 0)
	require.Nil(t, err)
	require.Equal(t, c.ApiUrl, fmt.Sprintf("https://%s:%d/json-rpc/%s", defaultTarget, defaultPort, defaultVersion))
}