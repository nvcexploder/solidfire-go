package api_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/joyent/solidfire-sdk/api"
	"github.com/pkg/errors"
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

	opts := api.ClientOptions{
		Target:   host,
		Username: username,
		Password: password,
	}
	c, err := api.BuildClient(opts)
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

	opts := api.ClientOptions{
		Target:   host,
		Username: username,
		Password: password,
	}
	c, err := api.BuildClient(opts)
	if err != nil {
		t.Fatalf("Error connecting: %s\n", err)
	}
	return c
}

type EphemeralEntity struct {
	Entity  interface{}
	Destroy func()
}

func buildEphmeralEntity(entity interface{}, entityDestroy func() (err error)) (e EphemeralEntity) {
	destroyer := func() {
		err := entityDestroy()
		if err != nil {
			fmt.Printf("Failed to delete entity %#v during test cleanup. Error was: %s\n", entity, err)
		}
	}
	return EphemeralEntity{
		Entity:  entity,
		Destroy: destroyer,
	}
}

func createEphemeralInitiator(t *testing.T, c *api.Client, initiatorArgs api.CreateInitiator) (result EphemeralEntity) {
	ctx := context.Background()
	initiators, err := c.CreateInitiators(ctx, []api.CreateInitiator{initiatorArgs})
	if err != nil {
		t.Fatal(err)
	}
	if len(initiators) < 1 {
		// Should be impossible
		t.Fatal(errors.New("Failed to create initiator"))
	}
	initiator := initiators[0]
	return buildEphmeralEntity(initiator, func() (err error) { return c.DeleteInitiators(ctx, []int64{initiator.InitiatorID}) })
}

func createEphemeralVolumeAccessGroup(t *testing.T, c *api.Client, req api.CreateVolumeAccessGroupRequest) (result EphemeralEntity) {
	ctx := context.Background()
	volumeAccessGroup, err := c.CreateVolumeAccessGroup(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	delReq := api.DeleteVolumeAccessGroupRequest{
		VolumeAccessGroupID:    volumeAccessGroup.VolumeAccessGroupID,
		DeleteOrphanInitiators: false,
		Force:                  true,
	}
	return buildEphmeralEntity(*volumeAccessGroup, func() (err error) { return c.DeleteVolumeAccessGroup(ctx, delReq) })
}

func createEphemeralVolume(t *testing.T, c *api.Client, req api.CreateVolumeRequest) (result EphemeralEntity) {
	ctx := context.Background()
	volume, err := c.CreateVolume(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	return buildEphmeralEntity(*volume, func() (err error) {
		_, err = c.DeleteVolume(ctx, volume.VolumeID)
		return err
	})
}
