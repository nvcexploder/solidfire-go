package api

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestClient_StartRemoteS3Backup(t *testing.T) {
	defer gock.Off()

	c := getTestClient(t)
	e := map[string]interface{}{
		"id":     0,
		"method": "StartBulkVolumeRead",
		"params": StartBulkVolumeReadRequest{
			VolumeID: 57,
			Format:   FormatNative,
			Script:   BulkVolumeScript,
			ScriptParameters: S3WriteParameters{
				Range: BulkVolumeRange{
					LBA:   0,
					Range: 10,
				},
				Write: S3Params{
					AWSAccessKeyID:     "fake",
					AWSSecretAccessKey: "fake",
					Bucket:             "fake",
					Prefix:             "fake",
					Endpoint:           EndpointS3,
					Format:             FormatNative,
					Hostname:           "fake",
				},
			},
		},
	}

	gock.New(c.ApiUrl).Post("").MatchType("application/json").JSON(e).
		Reply(200).JSON(buildSFResponseWrapper(
		map[string]interface{}{
			"asyncHandle": 123,
			"key":         "key",
			"url":         "https://example.com",
		},
	))
	gock.InterceptClient(c.HTTPClient.GetClient())
	id, err := c.StartRemoteS3Backup(context.Background(), S3BackupRequest{
		VolumeID: 57,
		Range: BulkVolumeRange{
			LBA:   0,
			Range: 10,
		},
		Params: S3Params{
			AWSAccessKeyID:     "fake",
			AWSSecretAccessKey: "fake",
			Bucket:             "fake",
			Prefix:             "fake",
			Endpoint:           EndpointS3,
			Format:             FormatNative,
			Hostname:           "fake",
		},
	})
	spew.Dump(gock.GetUnmatchedRequests())
	assert.NoError(t, err)
	assert.Equal(t, AsyncResultID(123), id)
}
