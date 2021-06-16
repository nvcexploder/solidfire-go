package api_test

import (
	"os"
	"testing"

	"github.com/joyent/solidfire-sdk/api"
)

const IntegrationTestHelp = "Set $SOLIDFIRE_HOST, $SOLIDFIRE_USER, and $SOLIDFIRE_PASS to enable integration tests"

func IntegrationTestsDisabled() bool {
	host := os.Getenv("SOLIDFIRE_HOST")
	host2 := os.Getenv("SOLIDFIRE_HOST2")
	username := os.Getenv("SOLIDFIRE_USER")
	password := os.Getenv("SOLIDFIRE_PASS")
	return host == "" || host2 == "" || username == "" || password == ""
}

func BuildTestClient(t *testing.T) *api.Client {
	host := os.Getenv("SOLIDFIRE_HOST")
	username := os.Getenv("SOLIDFIRE_USER")
	password := os.Getenv("SOLIDFIRE_PASS")
	if host == "" || username == "" || password == "" {
		t.Fatal("Environment variables SOLIDFIRE_HOST, SOLIDFIRE_HOST2, SOLIDFIRE_USER, and SOLIDFIRE_PASS must be set")
	}

	opts := api.ClientOptions{}
	c, err := api.BuildClient(host, username, password, "12.3", 443, opts)
	if err != nil {
		t.Fatalf("Error connecting: %s\n", err)
	}
	return c
}

func BuildTestClientHost2(t *testing.T) *api.Client {
	host := os.Getenv("SOLIDFIRE_HOST2")
	username := os.Getenv("SOLIDFIRE_USER")
	password := os.Getenv("SOLIDFIRE_PASS")
	if host == "" || username == "" || password == "" {
		t.Fatal("Environment variables SOLIDFIRE_HOST, SOLIDFIRE_HOST2, SOLIDFIRE_USER, and SOLIDFIRE_PASS must be set")
	}

	opts := api.ClientOptions{}
	c, err := api.BuildClient(host, username, password, "12.3", 443, opts)
	if err != nil {
		t.Fatalf("Error connecting: %s\n", err)
	}
	return c
}
