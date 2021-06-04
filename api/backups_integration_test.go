package api_test

import (
	"context"
	"os"
	"testing"

	"github.com/joyent/solidfire-sdk/api"
	"github.com/stretchr/testify/assert"
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
		t.Fatal("Environment variables SOLIDFIRE_HOST, SOLIDFIRE_USER, and SOLIDFIRE_PASS must be set")
	}

	c, err := api.BuildClient(host, username, password, "12.3", 443, 3)
	if err != nil {
		t.Fatalf("Error connecting: %s\n", err)
	}
	return c
}

// Note: These tests simply verify that the request is accepted and a handle to the async operation is returned
// It's expected that the actual operation will fail (and in fact desired since we don't want the restore tests
// to overwrite anything)

func Test_RemoteS3Backup(t *testing.T) {
	skip.If(t, IntegrationTestsDisabled, IntegrationTestHelp)
	subject := testClient(t)
	vol, snap := identifyBackupSnapshot(t, subject)
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
	subject := testClient(t)
	vol, snap := identifyBackupSnapshot(t, subject)
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
	subject := testClient(t)
	// subject.HTTPClient.Debug = true
	vol := identifyRestoreVolume(t, subject)
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
	subject := testClient(t)
	vol := identifyRestoreVolume(t, subject)
	id, key, err := subject.StartRemoteSolidFireRestore(context.Background(), vol, api.FormatNative)

	assert.NoError(t, err)
	assert.NotZero(t, id)
	assert.NotEmpty(t, key)
}

func Test_ListAsyncResults(t *testing.T) {
	skip.If(t, IntegrationTestsDisabled, IntegrationTestHelp)
	subject := testClient(t)
	_, err := subject.ListAllAsyncTasks(context.Background(), api.ListAsyncResultsRequest{})
	assert.NoError(t, err)
}

func Test_GetAsyncResult(t *testing.T) {
	skip.If(t, IntegrationTestsDisabled, IntegrationTestHelp)
	subject := testClient(t)
	id := fetchAsyncTask(t, subject)
	_, err := subject.GetAsyncTask(context.Background(), api.GetAsyncResultRequest{
		AsyncHandle: id,
	})
	assert.NoError(t, err)
	assert.NotZero(t, id)
}

func Test_ListEvents(t *testing.T) {
	skip.If(t, IntegrationTestsDisabled, IntegrationTestHelp)
	subject := testClient(t)
	_, err := subject.GetEventList(context.Background(), api.ListEventsRequest{})
	assert.NoError(t, err)
}

func identifyBackupSnapshot(t *testing.T, c *api.Client) (volumeId, snapshotId int64) {
	snaps, err := c.ListSnapshots(context.Background(), api.ListSnapshotsRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if len(snaps) < 1 {
		t.Fatal(("found no snapshot for testing"))
	}
	return snaps[0].VolumeID, snaps[0].SnapshotID
}

func identifyRestoreVolume(t *testing.T, c *api.Client) (volumeId int64) {
	vols, err := c.ListVolumes(context.Background(), api.ListVolumesRequest{
		Limit: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(vols) < 1 {
		t.Fatal("found no volume for testing")
	}
	return vols[0].VolumeID
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
