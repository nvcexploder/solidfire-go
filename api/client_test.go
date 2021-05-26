package api

import (
	"fmt"
	"testing"

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
	_, err = BuildClient("", defaultUsername, defaultPassword, defaultVersion, defaultPort, defaultTimeout)
	require.Equal(t, err.Error(), ErrNoTarget)
	_, err = BuildClient(defaultTarget, "", defaultPassword, defaultVersion, defaultPort, defaultTimeout)
	require.Equal(t, err.Error(), ErrNoCredentials)
	_, err = BuildClient(defaultTarget, defaultUsername, "", defaultVersion, defaultPort, defaultTimeout)
	require.Equal(t, err.Error(), ErrNoCredentials)
}

func TestBuldClient(t *testing.T) {
	c, err := BuildClient(defaultTarget, defaultUsername, defaultPassword, defaultVersion, 443, 0)
	require.Nil(t, err)
	require.Equal(t, c.ApiUrl, fmt.Sprintf("https://%s:%d/json-rpc/%s", defaultTarget, defaultPort, defaultVersion))
}
