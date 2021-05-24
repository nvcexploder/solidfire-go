package api

import (
	"context"
	"testing"

	"github.com/joyent/solidfire-sdk/types"
	"github.com/stretchr/testify/require"
)

const testVolumeId int64 = 3576
const testAccountId int64 = 1

var testVolume = map[string]interface{}{
	"volumeID":        testVolumeId,
	"name":            "solidfire-sdk-test",
	"accountID":       testAccountId,
	"createTime":      "2021-05-21T17:44:29Z",
	"status":          "active",
	"access":          "readWrite",
	"enable512e":      true,
	"iqn":             "iqn.2010-01.com.solidfire:youn.solidfire-sdk-test.3576",
	"scsiEUIDeviceID": "796f756e00000df8f47acc0100000000",
	"scsiNAADeviceID": "6f47acc100000000796f756e00000df8",
	"qos": map[string]interface{}{
		"minIOPS":   50,
		"maxIOPS":   15000,
		"burstIOPS": 15000,
		"burstTime": 60,
		"curve": map[string]float64{
			"1048576": 15000,
			"131072":  1950,
			"16384":   270,
			"262144":  3900,
			"32768":   500,
			"4096":    100,
			"524288":  7600,
			"65536":   1000,
			"8192":    160,
		},
	},
	"volumeAccessGroups": []types.VolumeAccessGroup{},
	"volumePairs":        []types.VolumePair{},
	"sliceCount":         1,
	"totalSize":          1500000000,
	"blockSize":          4096,
	"attributes":         map[string]interface{}{},
}

func TestCreateVolume(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(map[string]interface{}{"Volume": testVolume})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	var totalSize int64 = 1500000000
	req := types.CreateVolumeRequest{
		Name:       "solidfire-sdk-test",
		AccountID:  testAccountId,
		TotalSize:  totalSize,
		Enable512e: true,
	}
	resp, err := c.CreateVolume(ctx, req)
	require.Nil(t, err)
	require.Equal(t, testVolumeId, resp.VolumeID)
	require.Equal(t, testAccountId, resp.AccountID)
	require.Equal(t, "solidfire-sdk-test", resp.Name)
	require.Equal(t, "active", resp.Status)
	require.Equal(t, "readWrite", resp.Access)
	require.Equal(t, totalSize, resp.TotalSize)
	require.Equal(t, "iqn.2010-01.com.solidfire:youn.solidfire-sdk-test.3576", resp.Iqn)
}

func TestDeleteVolume(t *testing.T) {
	c := getTestClient(t)
	testVolume2 := make(map[string]interface{})
	for k, v := range testVolume {
		testVolume2[k] = v
	}

	testVolume2["Status"] = "deleted"
	mockResp := buildSFResponseWrapper(map[string]interface{}{"Volume": testVolume2})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	resp, err := c.DeleteVolume(ctx, testVolumeId)
	require.Nil(t, err)
	require.Equal(t, testVolumeId, resp.VolumeID)
	require.Equal(t, testAccountId, resp.AccountID)
	require.Equal(t, "deleted", resp.Status)
}

func TestModifyVolume(t *testing.T) {
	c := getTestClient(t)
	testVolume2 := make(map[string]interface{})
	for k, v := range testVolume {
		testVolume2[k] = v
	}
	var newTotalSize int64 = 2000000000
	testVolume2["TotalSize"] = newTotalSize
	mockResp := buildSFResponseWrapper(map[string]interface{}{"Volume": testVolume2})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	req := types.ModifyVolumeRequest{
		TotalSize: newTotalSize,
	}
	resp, err := c.ModifyVolume(ctx, req)
	require.Nil(t, err)
	require.Equal(t, testVolumeId, resp.VolumeID)
	require.Equal(t, testAccountId, resp.AccountID)
	require.Equal(t, resp.TotalSize, newTotalSize)
}

func TestListVolumes(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(map[string]interface{}{"Volumes": []map[string]interface{}{testVolume}})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	req := types.ListVolumesRequest{}
	resp, err := c.ListVolumes(ctx, req)
	require.Nil(t, err)
	require.Equal(t, testVolumeId, resp[0].VolumeID)
	require.Equal(t, testAccountId, resp[0].AccountID)
	require.Equal(t, "solidfire-sdk-test", resp[0].Name)
	require.Equal(t, "active", resp[0].Status)
	require.Equal(t, "readWrite", resp[0].Access)
	require.Equal(t, "iqn.2010-01.com.solidfire:youn.solidfire-sdk-test.3576", resp[0].Iqn)
}

func TestGetVolumeById(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(map[string]interface{}{"Volumes": []map[string]interface{}{testVolume}})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	resp, err := c.GetVolumeById(ctx, testVolumeId)
	require.Nil(t, err)
	require.Equal(t, testVolumeId, resp.VolumeID)
	require.Equal(t, testAccountId, resp.AccountID)
	require.Equal(t, "solidfire-sdk-test", resp.Name)
	require.Equal(t, "active", resp.Status)
	require.Equal(t, "readWrite", resp.Access)
	require.Equal(t, "iqn.2010-01.com.solidfire:youn.solidfire-sdk-test.3576", resp.Iqn)
}
