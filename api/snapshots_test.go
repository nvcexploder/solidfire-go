package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

const testSnapshotName = "solidfire-sdk-snapshot-test"
const testSnapshotId int64 = 9500
const testSnapshotUUID = "1ffbfa6e-e9b6-4069-a9a2-ad33f1191b01"
const testSnapshotVolumeId int64 = 1234

var testSnapshot = map[string]interface{}{
	"snapshotID":              9500,
	"volumeID":                testSnapshotVolumeId,
	"name":                    testSnapshotName,
	"checksum":                "0x0",
	"enableRemoteReplication": false,
	"expirationReason":        "None",
	"status":                  "done",
	"snapshotUUID":            testSnapshotUUID,
	"totalSize":               1101004800,
	"groupSnapshotUUID":       "00000000-0000-0000-0000-000000000000",
	"createTime":              "2021-05-24T15:18:14Z",
	"attributes":              map[string]interface{}{},
}

func TestCreateSnapshot(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(map[string]interface{}{"Snapshot": testSnapshot})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	req := CreateSnapshotRequest{
		VolumeID: testSnapshotVolumeId,
		Name:     testSnapshotName,
	}
	resp, err := c.CreateSnapshot(ctx, req)
	require.Nil(t, err)
	require.Equal(t, testSnapshotVolumeId, resp.VolumeID)
	require.Equal(t, testSnapshotName, resp.Name)
	require.Equal(t, testSnapshotUUID, resp.SnapshotUUID)
}

func TestModifySnapshot(t *testing.T) {

	c := getTestClient(t)
	testSnapshot2 := make(map[string]interface{})
	for k, v := range testSnapshot {
		testSnapshot2[k] = v
	}
	snapshotName := "solidfire-sdk-snapshot-test-2"
	testSnapshot2["name"] = snapshotName
	mockResp := buildSFResponseWrapper(map[string]interface{}{"Snapshot": testSnapshot2})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	req := ModifySnapshotRequest{
		SnapshotID: testSnapshotId,
		Name:       snapshotName,
	}
	resp, err := c.ModifySnapshot(ctx, req)
	require.Nil(t, err)
	require.Equal(t, testSnapshotId, resp.SnapshotID)
	require.Equal(t, snapshotName, resp.Name)
}

func TestDeleteSnapshot(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(nil)
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	err := c.DeleteSnapshot(ctx, testSnapshotId)
	require.Nil(t, err)
}

func TestListSnapshots(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(map[string]interface{}{"Snapshots": []map[string]interface{}{testSnapshot}})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()
	ctx := context.Background()
	req := ListSnapshotsRequest{}
	resp, err := c.ListSnapshots(ctx, req)
	require.Nil(t, err)
	require.True(t, (len(resp) > 0))
	require.Equal(t, testSnapshotUUID, resp[0].SnapshotUUID)
}

func TestGetSnapshotById(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(map[string]interface{}{"Snapshots": []map[string]interface{}{testSnapshot}})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()
	ctx := context.Background()
	resp, err := c.GetSnapshotById(ctx, testSnapshotId)
	require.Nil(t, err)
	require.Equal(t, testSnapshotUUID, resp.SnapshotUUID)
}

func TestGetSnapshotByVolumeId(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(map[string]interface{}{"Snapshots": []map[string]interface{}{testSnapshot}})
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()
	ctx := context.Background()
	resp, err := c.GetSnapshotsByVolumeId(ctx, testSnapshotVolumeId)
	require.Nil(t, err)
	require.Equal(t, testSnapshotVolumeId, resp[0].VolumeID)
}
