package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

const testVolumeAccessGroupId = int64(37)
const testVolumeAccessGroupName = "1"

var testVolumeAccessGroup = map[string]interface{}{
	"deletedVolume":       []int64{},
	"volumeAccessGroupID": testVolumeAccessGroupId,
	"name":                testVolumeAccessGroupName,
	"initiatorIDs":        []int64{},
	"initiators":          []string{},
	"volumes":             []int64{},
	"attributes":          map[string]interface{}{},
}

func TestCreateVolumeAccessGroup(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(map[string]interface{}{"volumeAccessGroup": testVolumeAccessGroup})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	req := CreateVolumeAccessGroupRequest{
		Name: testVolumeAccessGroupName,
	}
	resp, err := c.CreateVolumeAccessGroup(ctx, req)

	require.Nil(t, err)
	require.Equal(t, testVolumeAccessGroupId, resp.VolumeAccessGroupID)
	require.Equal(t, testVolumeAccessGroupName, resp.Name)
}

func TestDeleteVolumeAccessGroup(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(nil)
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	req := DeleteVolumeAccessGroupRequest{
		VolumeAccessGroupID:    testVolumeAccessGroupId,
		DeleteOrphanInitiators: true,
		Force:                  true,
	}
	err := c.DeleteVolumeAccessGroup(ctx, req)
	require.Nil(t, err)
}

func TestModifyVolumeAccessGroup(t *testing.T) {
	testVolumeAccessGroup2 := make(map[string]interface{})
	for k, v := range testVolumeAccessGroup {
		testVolumeAccessGroup2[k] = v
	}
	const modifiedTestVolumeAccessGroupName = "2"
	testVolumeAccessGroup2["name"] = modifiedTestVolumeAccessGroupName
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(map[string]interface{}{"volumeAccessGroup": testVolumeAccessGroup2})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	req := ModifyVolumeAccessGroupRequest{
		VolumeAccessGroupID: testVolumeAccessGroupId,
		Name:                modifiedTestVolumeAccessGroupName,
	}
	resp, err := c.ModifyVolumeAccessGroup(ctx, req)

	require.Nil(t, err)
	require.Equal(t, testVolumeAccessGroupId, resp.VolumeAccessGroupID)
	require.Equal(t, modifiedTestVolumeAccessGroupName, resp.Name)
}

func TestListVolumeAccessGroups(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(map[string]interface{}{"volumeAccessGroups": []map[string]interface{}{testVolumeAccessGroup}})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	req := ListVolumeAccessGroupsRequest{
		VolumeAccessGroups: []int64{testVolumeAccessGroupId},
	}
	resp, err := c.ListVolumeAccessGroups(ctx, req)

	require.Nil(t, err)
	require.True(t, len(resp) > 0)
	firstVolumeAccessGroup := resp[0]
	require.Equal(t, testVolumeAccessGroupId, firstVolumeAccessGroup.VolumeAccessGroupID)
	require.Equal(t, testVolumeAccessGroupName, firstVolumeAccessGroup.Name)
}

func TestGetVolumeAccessGroup(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(map[string]interface{}{"volumeAccessGroups": []map[string]interface{}{testVolumeAccessGroup}})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	resp, err := c.GetVolumeAccessGroup(ctx, testVolumeAccessGroupId)

	require.Nil(t, err)
	require.Equal(t, testVolumeAccessGroupId, resp.VolumeAccessGroupID)
	require.Equal(t, testVolumeAccessGroupName, resp.Name)
}

func TestAddInitiatorsToVolumeAccessGroup(t *testing.T) {
	vagWithInits := make(map[string]interface{})
	for k, v := range testVolumeAccessGroup {
		vagWithInits[k] = v
	}
	initiatorIds := []int64{1}
	initiators := []string{"iqn.1993-08.org.debian:01:181324777"}
	vagWithInits["initiatorIDs"] = initiatorIds
	vagWithInits["initiators"] = initiators
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(map[string]interface{}{"volumeAccessGroup": vagWithInits})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	resp, err := c.AddInitiatorsToVolumeAccessGroup(ctx, testVolumeAccessGroupId, initiatorIds)

	require.Nil(t, err)
	require.Equal(t, testVolumeAccessGroupId, resp.VolumeAccessGroupID)
	require.Equal(t, initiatorIds, resp.InitiatorIDs)
	require.Equal(t, initiators, resp.Initiators)
}

func TestRemoveInitiatorsFromVolumeAccessGroup(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(map[string]interface{}{"volumeAccessGroup": testVolumeAccessGroup})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	resp, err := c.RemoveInitiatorsFromVolumeAccessGroup(ctx, testVolumeAccessGroupId, []int64{1}, false)

	require.Nil(t, err)
	require.Equal(t, testVolumeAccessGroupId, resp.VolumeAccessGroupID)
	require.Equal(t, []int64{}, resp.InitiatorIDs)
	require.Equal(t, []string{}, resp.Initiators)
}

func TestAddVolumesToVolumeAccessGroup(t *testing.T) {
	vagWithVolumes := make(map[string]interface{})
	for k, v := range testVolumeAccessGroup {
		vagWithVolumes[k] = v
	}
	volumeIds := []int64{1}
	vagWithVolumes["volumes"] = volumeIds
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(map[string]interface{}{"volumeAccessGroup": vagWithVolumes})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	resp, err := c.AddVolumesToVolumeAccessGroup(ctx, testVolumeAccessGroupId, volumeIds)

	require.Nil(t, err)
	require.Equal(t, testVolumeAccessGroupId, resp.VolumeAccessGroupID)
	require.Equal(t, volumeIds, resp.Volumes)
}

func TestRemoveVolumesFromVolumeAccessGroup(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(map[string]interface{}{"volumeAccessGroup": testVolumeAccessGroup})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	resp, err := c.RemoveVolumesFromVolumeAccessGroup(ctx, testVolumeAccessGroupId, []int64{1})

	require.Nil(t, err)
	require.Equal(t, testVolumeAccessGroupId, resp.VolumeAccessGroupID)
	require.Equal(t, []int64{}, resp.Volumes)
}
