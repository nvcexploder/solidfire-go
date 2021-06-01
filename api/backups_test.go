package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

const (
	AWSAccessKeyID     = "fake-accesskey"
	AWSSecretAccessKey = "fake-secretkey"
	Bucket             = "fake-bucket"
	Prefix             = "fake-prefix"
	Hostname           = "fake-hostname"
	ManagementVIP      = "fake-mvip"
	Username           = "fake-username"
	Password           = "fake-password"
	VolumeWriteKey     = "fake-vwritekey"
)

func TestClient_StartRemoteS3Backup(t *testing.T) {
	defer gock.Off()

	c := getTestClient(t)
	expectedBody := map[string]interface{}{
		"id":     0,
		"method": "StartBulkVolumeRead",
		"params": StartBulkVolumeReadRequest{
			VolumeID: 100,
			Format:   FormatNative,
			Script:   BulkVolumeScript,
			ScriptParameters: S3WriteParameters{
				Range: BulkVolumeRange{
					LBA:   0,
					Range: 10,
				},
				Write: s3Params{
					S3Params: S3Params{
						AWSAccessKeyID:     AWSAccessKeyID,
						AWSSecretAccessKey: AWSSecretAccessKey,
						Bucket:             Bucket,
						Prefix:             Prefix,
						Format:             FormatNative,
						Hostname:           Hostname,
					},
					Endpoint: EndpointS3,
				},
			},
		},
	}

	gock.New(c.ApiUrl).Post("").MatchType("application/json").JSON(expectedBody).
		Reply(200).JSON(buildSFResponseWrapper(
		map[string]interface{}{
			"asyncHandle": 123,
			"key":         "key",
			"url":         "https://example.com",
		},
	))
	gock.InterceptClient(c.HTTPClient.GetClient())
	id, err := c.StartRemoteS3Backup(context.Background(), S3BackupRequest{
		VolumeID: 100,
		Range: BulkVolumeRange{
			LBA:   0,
			Range: 10,
		},
		Params: S3Params{
			AWSAccessKeyID:     AWSAccessKeyID,
			AWSSecretAccessKey: AWSSecretAccessKey,
			Bucket:             Bucket,
			Prefix:             Prefix,
			Format:             FormatNative,
			Hostname:           Hostname,
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, AsyncResultID(123), id)
}

func TestClient_StartRemoteSolidFireBackup(t *testing.T) {
	defer gock.Off()

	c := getTestClient(t)
	expectedBody := map[string]interface{}{
		"id":     0,
		"method": "StartBulkVolumeRead",
		"params": StartBulkVolumeReadRequest{
			VolumeID: 100,
			Format:   FormatNative,
			Script:   BulkVolumeScript,
			ScriptParameters: solidFireWriteParameters{
				Range: BulkVolumeRange{
					LBA:   0,
					Range: 10,
				},
				Write: solidFireParams{
					SolidFireParams: SolidFireParams{
						ManagementVIP:  ManagementVIP,
						Username:       Username,
						Password:       Password,
						Format:         FormatNative,
						VolumeWriteKey: VolumeWriteKey,
					},
					Endpoint: EndpointSolidFire,
				},
			},
		},
	}

	gock.New(c.ApiUrl).Post("").MatchType("application/json").JSON(expectedBody).
		Reply(200).JSON(buildSFResponseWrapper(
		map[string]interface{}{
			"asyncHandle": 123,
			"key":         "key",
			"url":         "https://example.com",
		},
	))
	gock.InterceptClient(c.HTTPClient.GetClient())
	id, err := c.StartRemoteSolidFireBackup(context.Background(), SolidFireBackupRequest{
		VolumeID: 100,
		Range: BulkVolumeRange{
			LBA:   0,
			Range: 10,
		},
		Params: SolidFireParams{
			ManagementVIP:  ManagementVIP,
			Username:       Username,
			Password:       Password,
			Format:         FormatNative,
			VolumeWriteKey: VolumeWriteKey,
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, AsyncResultID(123), id)
}

func TestClient_StartRemoteS3Restore(t *testing.T) {
	defer gock.Off()

	c := getTestClient(t)
	expectedBody := map[string]interface{}{
		"id":     0,
		"method": "StartBulkVolumeWrite",
		"params": StartBulkVolumeWriteRequest{
			VolumeID: 100,
			Format:   FormatNative,
			Script:   BulkVolumeScript,
			ScriptParameters: S3ReadParameters{
				Range: BulkVolumeRange{
					LBA:   0,
					Range: 10,
				},
				Read: s3Params{
					S3Params: S3Params{
						AWSAccessKeyID:     AWSAccessKeyID,
						AWSSecretAccessKey: AWSSecretAccessKey,
						Bucket:             Bucket,
						Prefix:             Prefix,
						Format:             FormatNative,
						Hostname:           Hostname,
					},
					Endpoint: EndpointS3,
				},
			},
		},
	}

	gock.New(c.ApiUrl).Post("").MatchType("application/json").JSON(expectedBody).
		Reply(200).JSON(buildSFResponseWrapper(
		map[string]interface{}{
			"asyncHandle": 123,
			"key":         "key",
			"url":         "https://example.com",
		},
	))
	gock.InterceptClient(c.HTTPClient.GetClient())
	id, err := c.StartRemoteS3Restore(context.Background(), S3RestoreRequest{
		VolumeID: 100,
		Range: BulkVolumeRange{
			LBA:   0,
			Range: 10,
		},
		Params: S3Params{
			AWSAccessKeyID:     AWSAccessKeyID,
			AWSSecretAccessKey: AWSSecretAccessKey,
			Bucket:             Bucket,
			Prefix:             Prefix,
			Format:             FormatNative,
			Hostname:           Hostname,
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, AsyncResultID(123), id)
}

func TestClient_StartRemoteSolidFireRestore(t *testing.T) {
	defer gock.Off()

	c := getTestClient(t)
	expectedBody := map[string]interface{}{
		"id":     0,
		"method": "StartBulkVolumeWrite",
		"params": StartBulkVolumeWriteRequest{
			VolumeID: 100,
			Format:   FormatNative,
		},
	}

	gock.New(c.ApiUrl).Post("").MatchType("application/json").JSON(expectedBody).
		Reply(200).JSON(buildSFResponseWrapper(
		map[string]interface{}{
			"asyncHandle": 123,
			"key":         "key",
			"url":         "https://example.com",
		},
	))
	gock.InterceptClient(c.HTTPClient.GetClient())
	id, key, err := c.StartRemoteSolidFireRestore(context.Background(), 100, FormatNative)
	assert.NoError(t, err)
	assert.Equal(t, AsyncResultID(123), id)
	assert.Equal(t, "key", key)
}
