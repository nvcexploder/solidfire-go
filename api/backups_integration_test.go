package api_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/joyent/solidfire-sdk/api"
	"github.com/stretchr/testify/assert"
	"gotest.tools/skip"
)

// Note: These tests simply verify that the request is accepted and a handle to the async operation is returned
// It's expected that the actual operation will fail (and in fact desired since we don't want the restore tests
// to overwrite anything)

func Test_RemoteS3Backup(t *testing.T) {
	skip.If(t, IntegrationTestsDisabled, IntegrationTestHelp)
	subject := BuildTestClient(t)
	vol, snap, cleanup := createTestSnapshot(t, subject)
	defer cleanup()
	id, err := subject.StartRemoteS3Backup(context.Background(), api.S3BackupRequest{
		VolumeID:   vol,
		SnapshotID: snap,
		Range: api.BulkVolumeRange{
			LBA:   0,
			Range: 10,
		},
		Params: api.S3Params{
			AWSAccessKeyID:     "fake",
			AWSSecretAccessKey: "fake",
			Bucket:             "fake",
			Prefix:             "fake",
			Format:             api.FormatNative,
			Hostname:           "fake",
		},
	})
	assert.NoError(t, err)
	assert.NotZero(t, id)
}

func Test_StartRemoteSolidFireBackup(t *testing.T) {
	skip.If(t, IntegrationTestsDisabled, IntegrationTestHelp)
	subject := BuildTestClient(t)
	vol, snap, cleanup := createTestSnapshot(t, subject)
	defer cleanup()
	id, err := subject.StartRemoteSolidFireBackup(context.Background(), api.SolidFireBackupRequest{
		VolumeID:   vol,
		SnapshotID: snap,
		Range: api.BulkVolumeRange{
			LBA:   0,
			Range: 10,
		},
		Params: api.SolidFireParams{
			ManagementVIP:  "fake",
			Username:       "fake",
			Password:       "fake",
			Format:         api.FormatNative,
			VolumeWriteKey: "fake",
		},
	})
	assert.NoError(t, err)
	assert.NotZero(t, id)
}

func Test_StartRemoteS3Restore(t *testing.T) {
	skip.If(t, IntegrationTestsDisabled, IntegrationTestHelp)
	subject := BuildTestClient(t)
	// subject.HTTPClient.Debug = true
	vol, cleanup := createTestVolume(t, subject)
	defer cleanup()
	id, err := subject.StartRemoteS3Restore(context.Background(), api.S3RestoreRequest{
		VolumeID: vol,
		Range: api.BulkVolumeRange{
			LBA:   0,
			Range: 10,
		},
		Params: api.S3Params{
			AWSAccessKeyID:     "fake",
			AWSSecretAccessKey: "fake",
			Bucket:             "fake",
			Prefix:             "fake",
			Format:             api.FormatNative,
			Hostname:           "fake",
		},
	})
	assert.NoError(t, err)
	assert.NotZero(t, id)
}

func Test_StartRemoteSolidFireRestore(t *testing.T) {
	skip.If(t, IntegrationTestsDisabled, IntegrationTestHelp)
	subject := BuildTestClient(t)
	vol, cleanup := createTestVolume(t, subject)
	defer cleanup()
	id, key, err := subject.StartRemoteSolidFireRestore(context.Background(), vol, api.FormatNative)

	assert.NoError(t, err)
	assert.NotZero(t, id)
	assert.NotEmpty(t, key)
}

func Test_ListAsyncResults(t *testing.T) {
	skip.If(t, IntegrationTestsDisabled, IntegrationTestHelp)
	subject := BuildTestClient(t)
	_, err := subject.ListAllAsyncTasks(context.Background(), api.ListAsyncResultsRequest{})
	assert.NoError(t, err)
}

func Test_GetAsyncResult(t *testing.T) {
	skip.If(t, IntegrationTestsDisabled, IntegrationTestHelp)
	subject := BuildTestClient(t)
	id := fetchAsyncTask(t, subject)
	_, err := subject.GetAsyncTask(context.Background(), api.GetAsyncResultRequest{
		AsyncHandle: id,
	})
	assert.NoError(t, err)
	assert.NotZero(t, id)
}

func Test_ListEvents(t *testing.T) {
	skip.If(t, IntegrationTestsDisabled, IntegrationTestHelp)
	subject := BuildTestClient(t)
	_, err := subject.GetEventList(context.Background(), api.ListEventsRequest{})
	assert.NoError(t, err)
}

func createTestSnapshot(t *testing.T, c *api.Client) (volumeId, snapshotId int64, cleanup func()) {
	createVolReq := api.CreateVolumeRequest{
		Name:       testVolumeName,
		AccountID:  testAccountId,
		TotalSize:  1 * api.Gigabytes,
		Enable512e: true,
	}
	eVol := createEphemeralVolume(t, c, createVolReq)
	volume := eVol.Entity.(api.Volume)
	snapReq := api.CreateSnapshotRequest{
		VolumeID: volume.VolumeID,
		Name:     "test-snapshot",
	}
	snap, err := c.CreateSnapshot(context.Background(), snapReq)
	if err != nil {
		eVol.Destroy()
		t.Fatal(err)
	}
	cleanup = func() {
		eVol.Destroy()
		err = c.DeleteSnapshot(context.Background(), snap.SnapshotID)
		if err != nil {
			fmt.Printf("Failed to delete entity %#v during test cleanup. Error was: %s\n", snap, err)
		}
	}
	return volume.VolumeID, snap.SnapshotID, cleanup
}

func createTestVolume(t *testing.T, c *api.Client) (volumeId int64, cleanup func()) {
	createVolReq := api.CreateVolumeRequest{
		Name:       testVolumeName,
		AccountID:  testAccountId,
		TotalSize:  1 * api.Gigabytes,
		Enable512e: true,
	}
	eVol := createEphemeralVolume(t, c, createVolReq)
	volume := eVol.Entity.(api.Volume)
	return volume.VolumeID, eVol.Destroy
}

func fetchAsyncTask(t *testing.T, c *api.Client) (taskID api.AsyncResultID) {
	res, err := c.ListAllAsyncTasks(context.Background(), api.ListAsyncResultsRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if len(res.AsyncHandles) < 1 {
		t.Fatal("found no async task for testing")
	}
	return res.AsyncHandles[0].AsyncResultID
}
