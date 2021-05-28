package api

import (
	"context"
	"fmt"

	"github.com/joyent/solidfire-sdk/types"
)

func (c *Client) StartBulkVolumeRead(ctx context.Context, r types.StartBulkVolumeReadRequest) (result types.StartBulkVolumeReadResult, err error) {
	result = types.StartBulkVolumeReadResult{}
	err = c.request(ctx, "StartBulkVolumeRead", r, &result)
	return result, err
}

func (c *Client) StartBulkVolumeWrite(ctx context.Context, r types.StartBulkVolumeWriteRequest) (result types.StartBulkVolumeWriteResult, err error) {
	result = types.StartBulkVolumeWriteResult{}
	err = c.request(ctx, "StartBulkVolumeWrite", r, &result)
	return result, err
}

func (c *Client) StartRemoteS3Backup(ctx context.Context, r types.S3BackupRequest) (result types.StartBulkVolumeReadResult, err error) {
	bvreq := types.StartBulkVolumeReadRequest{
		VolumeID:   r.VolumeID,
		SnapshotID: r.SnapshotID,
		Format:     r.Params.Format,
		Script:     types.BulkVolumeScript,
		ScriptParameters: types.S3ScriptParameters{
			Range: r.Range,
			Write: r.Params,
		},
	}
	return c.StartBulkVolumeRead(ctx, bvreq)
}

func (c *Client) StartRemoteSolidFireBackup(ctx context.Context) {

}

func (c *Client) StartRemoteSolidFireRestore(ctx context.Context, volumeID int64, format string) (result types.StartBulkVolumeWriteResult, err error) {
	return c.StartBulkVolumeWrite(ctx, types.StartBulkVolumeWriteRequest{
		VolumeID: volumeID,
		Format:   format,
	})
}

func validateFormat(format string) error {
	if format != types.FormatNative && format != types.FormatUncompressed {
		// TODO: use error types
		return fmt.Errorf("supported formats are '%s' and '%s'", types.FormatNative, types.FormatUncompressed)
	}

	return nil
}
