package api_test

import (
	"context"
	"os"
	"testing"

	"github.com/joyent/solidfire-sdk/api"
	"gotest.tools/skip"
)

const IntegrationTestHelp = "Set $SOLIDFIRE_HOST, $SOLIDFIRE_USER, and $SOLIDFIRE_PASS to enable integration tests"

func IntegrationTestsDisabled() bool {
	host := os.Getenv("SOLIDFIRE_HOST")
	username := os.Getenv("SOLIDFIRE_USER")
	password := os.Getenv("SOLIDFIRE_PASS")
	return host == "" || username == "" || password == ""
}

func testClient(t *testing.T) *api.Client {
	host := os.Getenv("SOLIDFIRE_HOST")
	username := os.Getenv("SOLIDFIRE_USER")
	password := os.Getenv("SOLIDFIRE_PASS")
	if host == "" || username == "" || password == "" {
		t.Fatal("Environment variables SOLIDFIRE_HOST, SOLIDFIRE_USER, and SOLIDER_PASS must be set")
	}

	c, err := api.BuildClient(host, username, password, "12.3", 443, 3)
	if err != nil {
		t.Fatalf("Error connecting: %s\n", err)
	}
	return c
}

func Test(t *testing.T) {
	skip.If(t, IntegrationTestsDisabled, IntegrationTestHelp)
	subject := testClient(t)
	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			subject.StartBulkVolumeRead(context.Background())
		})
	}
}
