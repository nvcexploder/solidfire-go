package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

const testVolumePairVolumeId int64 = 3577

var testVolumePairVolume = map[string]interface{}{
	"volumeID":        testVolumePairVolumeId,
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
	"volumeAccessGroups": []VolumeAccessGroup{},
	"volumePairs": []map[string]interface{}{
		{
			"clusterPairID":    1,
			"remoteVolumeID":   3776,
			"remoteSliceID":    3776,
			"remoteVolumeName": "1",
			"volumePairUUID":   "951d5294-bb2c-48d9-9a9e-99cd34fafd2b",
			"remoteReplication": map[string]interface{}{
				"mode":            "Sync",
				"pauseLimit":      3145728000,
				"remoteServiceID": 0,
				"resumeDetails":   "",
				"snapshotReplication": map[string]interface{}{
					"state":        "PausedMisconfigured",
					"stateDetails": "",
				},
				"state":        "PausedMisconfigured",
				"stateDetails": ""},
		}},
	"sliceCount": 1,
	"totalSize":  int64(1.5 * Gigabytes),
	"blockSize":  4096,
	"attributes": map[string]interface{}{},
}

func TestStartVolumePairing(t *testing.T) {
	c := getTestClient(t)
	testPairingKey := "ABCDEF-1234-12344321-0987654321-ABCDEF"
	mockResp := buildSFResponseWrapper(map[string]interface{}{"VolumePairingKey": testPairingKey})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	key, err := c.StartVolumePairing(ctx, testVolumePairVolumeId, VolumePairingModeAsync)
	require.Nil(t, err)
	require.Equal(t, testPairingKey, key)
}

func TestCompleteVolumePairing(t *testing.T) {
	c := getTestClient(t)
	testPairingKey := "ABCDEF-1234-12344321-0987654321-ABCDEF"
	mockResp := buildSFResponseWrapper(nil)
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	err := c.CompleteVolumePairing(ctx, testVolumePairVolumeId, testPairingKey)
	require.Nil(t, err)
}

func TestModifyVolumePair(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(nil)
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	req := ModifyVolumePairRequest{
		VolumeID:     testVolumePairVolumeId,
		PausedManual: false,
		Mode:         VolumePairingModeSync,
	}
	err := c.ModifyVolumePair(ctx, req)
	require.Nil(t, err)
}

func TestRemoveVolumePair(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(nil)
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	err := c.RemoveVolumePair(ctx, testVolumePairVolumeId)
	require.Nil(t, err)
}

func TestListActivePairedVolumes(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(map[string]interface{}{"volumes": []map[string]interface{}{testVolumePairVolume}})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	req := ListActivePairedVolumesRequest{StartVolumeID: testVolumePairVolumeId, Limit: 1}
	resp, err := c.ListActivePairedVolumes(ctx, req)
	require.Nil(t, err)
	require.True(t, len(resp) > 0)
	require.Equal(t, testVolumePairVolumeId, resp[0].VolumeID)
}
