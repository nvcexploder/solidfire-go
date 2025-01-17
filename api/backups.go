package api

import (
	"context"
)

func (c *Client) StartBulkVolumeRead(ctx context.Context, r StartBulkVolumeReadRequest) (result StartBulkVolumeReadResult, err error) {
	result = StartBulkVolumeReadResult{}
	err = c.request(ctx, "StartBulkVolumeRead", r, &result)
	return result, err
}

func (c *Client) StartBulkVolumeWrite(ctx context.Context, r StartBulkVolumeWriteRequest) (result StartBulkVolumeWriteResult, err error) {
	result = StartBulkVolumeWriteResult{}
	err = c.request(ctx, "StartBulkVolumeWrite", r, &result)
	return result, err
}

func (c *Client) StartRemoteS3Backup(ctx context.Context, r S3BackupRequest) (result AsyncResultID, err error) {
	bvreq := StartBulkVolumeReadRequest{
		VolumeID:   r.VolumeID,
		SnapshotID: r.SnapshotID,
		Format:     r.Params.Format,
		Script:     BulkVolumeScript,
		ScriptParameters: S3WriteParameters{
			Range: r.Range,
			Write: s3Params{
				S3Params: r.Params,
				Endpoint: EndpointS3,
			},
		},
	}
	bvresult, err := c.StartBulkVolumeRead(ctx, bvreq)
	if err != nil {
		return result, err
	}
	return AsyncResultID(bvresult.AsyncHandle), nil
}

func (c *Client) StartRemoteSolidFireBackup(ctx context.Context, r SolidFireBackupRequest) (result AsyncResultID, err error) {
	bvreq := StartBulkVolumeReadRequest{
		VolumeID:   r.VolumeID,
		SnapshotID: r.SnapshotID,
		Format:     r.Params.Format,
		Script:     BulkVolumeScript,
		ScriptParameters: solidFireWriteParameters{
			Range: r.Range,
			Write: solidFireParams{
				SolidFireParams: r.Params,
				Endpoint:        EndpointSolidFire,
			},
		},
	}
	bvresult, err := c.StartBulkVolumeRead(ctx, bvreq)
	if err != nil {
		return result, err
	}
	return AsyncResultID(bvresult.AsyncHandle), nil
}

func (c *Client) StartRemoteSolidFireRestore(ctx context.Context, volumeID int64, format string) (result AsyncResultID, key string, err error) {
	bvresult, err := c.StartBulkVolumeWrite(ctx, StartBulkVolumeWriteRequest{
		VolumeID: volumeID,
		Format:   format,
	})
	if err != nil {
		return result, key, err
	}
	return AsyncResultID(bvresult.AsyncHandle), bvresult.Key, nil
}

func (c *Client) StartRemoteS3Restore(ctx context.Context, r S3RestoreRequest) (result AsyncResultID, err error) {
	bvreq := StartBulkVolumeWriteRequest{
		VolumeID: r.VolumeID,
		Format:   r.Params.Format,
		Script:   BulkVolumeScript,
		ScriptParameters: S3ReadParameters{
			Range: r.Range,
			Read: s3Params{
				S3Params: r.Params,
				Endpoint: EndpointS3,
			},
		},
	}
	bvresult, err := c.StartBulkVolumeWrite(ctx, bvreq)
	if err != nil {
		return result, err
	}
	return AsyncResultID(bvresult.AsyncHandle), nil
}

func (c *Client) ListAllAsyncTasks(ctx context.Context, r ListAsyncResultsRequest) (results ListAsyncResultsResult, err error) {
	err = c.request(ctx, "ListAsyncResults", r, &results)
	if err != nil {
		return ListAsyncResultsResult{}, err
	}
	return results, nil
}

func (c *Client) GetAsyncTask(ctx context.Context, r GetAsyncResultRequest) (result GetAsyncResult, err error) {
	err = c.request(ctx, "GetAsyncResult", r, &result)
	if err != nil {
		return GetAsyncResult{}, err
	}
	return result, nil
}

func (c *Client) GetEventList(ctx context.Context, r ListEventsRequest) (result ListEventsResult, err error) {
	err = c.request(ctx, "ListEvents", r, &result)
	if err != nil {
		return ListEventsResult{}, err
	}
	return result, nil
}
