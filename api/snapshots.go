package api

import (
	"context"
	"errors"
)

const (
	ErrNoSnapshotFound = "No snapshot found with the given id"
)

func (c *Client) CreateSnapshot(ctx context.Context, req CreateSnapshotRequest) (result *Snapshot, err error) {
	csr := CreateSnapshotResult{}
	err = c.request(ctx, "CreateSnapshot", req, &csr)
	result = &csr.Snapshot
	return result, err
}

func (c *Client) ModifySnapshot(ctx context.Context, req ModifySnapshotRequest) (result *Snapshot, err error) {
	msr := ModifySnapshotResult{}
	err = c.request(ctx, "ModifySnapshot", req, &msr)
	result = &msr.Snapshot
	return result, err
}

func (c *Client) DeleteSnapshot(ctx context.Context, id int64) error {
	req := DeleteSnapshotRequest{
		SnapshotID: id,
	}
	return c.request(ctx, "DeleteSnapshot", req, nil)
}

func (c *Client) ListSnapshots(ctx context.Context, req ListSnapshotsRequest) (result []Snapshot, err error) {
	lsr := ListSnapshotsResult{}
	err = c.request(ctx, "ListSnapshots", req, &lsr)
	result = lsr.Snapshots
	return result, err
}

func (c *Client) GetSnapshotById(ctx context.Context, id int64) (result *Snapshot, err error) {
	req := ListSnapshotsRequest{
		SnapshotID: id,
	}
	resp, err := c.ListSnapshots(ctx, req)
	if len(resp) > 0 {
		result = &resp[0]
	} else if err == nil {
		err = errors.New(ErrNoSnapshotFound)
	}
	return result, err
}

func (c *Client) GetSnapshotsByVolumeId(ctx context.Context, id int64) (result []Snapshot, err error) {
	req := ListSnapshotsRequest{
		VolumeID: id,
	}
	return c.ListSnapshots(ctx, req)
}
