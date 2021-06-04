package tests

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/joyent/solidfire-sdk/api"
	"github.com/stretchr/testify/assert"
	"gotest.tools/skip"
)

const (
	// N.B. - Tests assume a test account already exists with id 1 (no account creation in SDK yet)
	testAccountId             = 1
	testInitiatorName         = "solidfire-sdk_volume-integration-test_initiator_1"
	testInitiatorName2        = "solidfire-sdk_volume-integration-test_initiator_2"
	testVolumeAccessGroupName = "1"
	testVolumeName            = "1"
)

var (
	defaultInitiatorRequestData = api.CreateInitiator{
		Name: testInitiatorName,
	}
	defaultVolumeAccessGroupRequestData = api.CreateVolumeAccessGroupRequest{
		Name: testVolumeAccessGroupName,
	}
	defaultVolumeRequestData = api.CreateVolumeRequest{
		Name:       testVolumeName,
		AccountID:  testAccountId,
		TotalSize:  1 * api.Gigabytes,
		Enable512e: true,
	}
)

type EphemeralEntity struct {
	Entity  interface{}
	Destroy func()
}

func buildEphmeralEntity(entity interface{}, entityDestroy func() (err error)) (e EphemeralEntity) {
	destroyer := func() {
		err := entityDestroy()
		if err != nil {
			fmt.Printf("Failed to entity %#v during test cleanup. Error was: %s\n", entity, err)
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

func Test_CreateVolumeAndInitiator(t *testing.T) {
	skip.If(t, IntegrationTestsDisabled, IntegrationTestHelp)
	subject := BuildTestClient(t)
	eInit := createEphemeralInitiator(t, subject, defaultInitiatorRequestData)
	defer eInit.Destroy()
	initiator := eInit.Entity.(api.Initiator)
	eVag := createEphemeralVolumeAccessGroup(t, subject, defaultVolumeAccessGroupRequestData)
	volumeAccessGroup := eVag.Entity.(api.VolumeAccessGroup)
	defer eVag.Destroy()
	eVol := createEphemeralVolume(t, subject, defaultVolumeRequestData)
	defer eVol.Destroy()
	volume := eVol.Entity.(api.Volume)
	assert.Equal(t, testInitiatorName, initiator.InitiatorName)
	assert.Equal(t, testVolumeAccessGroupName, volumeAccessGroup.Name)
	assert.Equal(t, testVolumeName, volume.Name)
}

func Test_ModifyVolumeAndInitiator(t *testing.T) {
	skip.If(t, IntegrationTestsDisabled, IntegrationTestHelp)
	subject := BuildTestClient(t)
	// build entities for modification
	eInit := createEphemeralInitiator(t, subject, defaultInitiatorRequestData)
	defer eInit.Destroy()
	initiator := eInit.Entity.(api.Initiator)
	eVag := createEphemeralVolumeAccessGroup(t, subject, defaultVolumeAccessGroupRequestData)
	volumeAccessGroup := eVag.Entity.(api.VolumeAccessGroup)
	defer eVag.Destroy()
	eVol := createEphemeralVolume(t, subject, defaultVolumeRequestData)
	defer eVol.Destroy()
	volume := eVol.Entity.(api.Volume)
	// Modify each entity
	ctx := context.Background()
	testInitiatorAlias := "solidfire-sdk-test-initiator-1"
	modInitReq := api.ModifyInitiator{
		InitiatorID: initiator.InitiatorID,
		Alias:       testInitiatorAlias,
	}
	modifiedInitiators, miErr := subject.ModifyInitiators(ctx, []api.ModifyInitiator{modInitReq})
	assert.Nil(t, miErr)
	modifiedInitiator := modifiedInitiators[0]
	modifiedVagName := "2"
	modVagReq := api.ModifyVolumeAccessGroupRequest{
		VolumeAccessGroupID: volumeAccessGroup.VolumeAccessGroupID,
		Initiators:          []int64{initiator.InitiatorID},
		Name:                modifiedVagName,
	}
	modifiedVag, modVagErr := subject.ModifyVolumeAccessGroup(ctx, modVagReq)
	assert.Nil(t, modVagErr)
	modifiedVolumeAccess := api.VolumeAccessPolicyReadOnly
	modVolReq := api.ModifyVolumeRequest{
		VolumeID: volume.VolumeID,
		Access:   modifiedVolumeAccess,
	}
	modifiedVolume, modVolErr := subject.ModifyVolume(ctx, modVolReq)
	assert.Nil(t, modVolErr)
	assert.Equal(t, testInitiatorAlias, modifiedInitiator.Alias)
	assert.Equal(t, modifiedVagName, modifiedVag.Name)
	assert.Equal(t, []int64{initiator.InitiatorID}, modifiedVag.InitiatorIDs)
	assert.Equal(t, modifiedVolumeAccess, modifiedVolume.Access)
}

func Test_AddRemoveFromVolumeAccessGroup(t *testing.T) {
	skip.If(t, IntegrationTestsDisabled, IntegrationTestHelp)
	subject := BuildTestClient(t)
	// establish default entities
	eInit := createEphemeralInitiator(t, subject, defaultInitiatorRequestData)
	defer eInit.Destroy()
	initiator := eInit.Entity.(api.Initiator)
	eVag := createEphemeralVolumeAccessGroup(t, subject, defaultVolumeAccessGroupRequestData)
	volumeAccessGroup := eVag.Entity.(api.VolumeAccessGroup)
	defer eVag.Destroy()
	eVol := createEphemeralVolume(t, subject, defaultVolumeRequestData)
	defer eVol.Destroy()
	volume := eVol.Entity.(api.Volume)
	// Add/remove initiator and volume from volume access group
	ctx := context.Background()
	initiatorData2 := api.CreateInitiator{
		Name: testInitiatorName2,
	}
	eInit2 := createEphemeralInitiator(t, subject, initiatorData2)
	defer eInit2.Destroy()
	initiator2 := eInit2.Entity.(api.Initiator)
	modifiedVag, modVagErr := subject.AddVolumesToVolumeAccessGroup(ctx, volumeAccessGroup.VolumeAccessGroupID, []int64{volume.VolumeID})
	assert.Nil(t, modVagErr)
	assert.Equal(t, volume.VolumeID, modifiedVag.Volumes[0])
	modifiedVag2, modVagErr2 := subject.RemoveVolumesFromVolumeAccessGroup(ctx, volumeAccessGroup.VolumeAccessGroupID, []int64{volume.VolumeID})
	assert.Nil(t, modVagErr2)
	assert.Equal(t, 0, len(modifiedVag2.Volumes))
	modifiedVag3, modVagErr3 := subject.AddInitiatorsToVolumeAccessGroup(ctx, volumeAccessGroup.VolumeAccessGroupID, []int64{initiator.InitiatorID, initiator2.InitiatorID})
	assert.Nil(t, modVagErr3)
	assert.Equal(t, []int64{initiator.InitiatorID, initiator2.InitiatorID}, modifiedVag3.InitiatorIDs)
	modifiedVag5, modVagErr4 := subject.RemoveInitiatorsFromVolumeAccessGroup(ctx, volumeAccessGroup.VolumeAccessGroupID, []int64{initiator2.InitiatorID}, false)
	assert.Nil(t, modVagErr4)
	assert.Equal(t, []int64{initiator.InitiatorID}, modifiedVag5.InitiatorIDs)
}
