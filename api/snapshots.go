package api

import (
	"context"
	"fmt"

	"github.com/joyent/solidfire-sdk/types"
)

func (c *Client) CreateSnapshot(ctx context.Context, req types.CreateSnapshotRequest) (result *types.Snapshot, err error) {
	csr := types.CreateSnapshotResult{}
	err = c.request(ctx, "CreateSnapshot", req, &csr)
	result = &csr.Snapshot
	return result, err
}

func (c *Client) ModifySnapshot(ctx context.Context, req types.ModifySnapshotRequest) (result *types.Snapshot, err error) {
	msr := types.ModifySnapshotResult{}
	err = c.request(ctx, "ModifySnapshot", req, &msr)
	result = &msr.Snapshot
	return result, err
}

func (c *Client) DeleteSnapshot(ctx context.Context, id int64) error {
	req := types.DeleteSnapshotRequest{
		SnapshotID: id,
	}
	return c.request(ctx, "DeleteSnapshot", req, nil)
}

func (c *Client) ListSnapshots(ctx context.Context, req types.ListSnapshotsRequest) (result []types.Snapshot, err error) {
	lsr := types.ListSnapshotsResult{}
	err = c.request(ctx, "ListSnapshots", req, &lsr)
	result = lsr.Snapshots
	return result, err
}

func (c *Client) GetSnapshotById(ctx context.Context, id int64) (result *types.Snapshot, err error) {
	req := types.ListSnapshotsRequest{
		SnapshotID: id,
	}
	resp, err := c.ListSnapshots(ctx, req)
	if len(resp) > 0 {
		result = &resp[0]
	} else if err == nil {
		err = fmt.Errorf("No snapshot found with the given id: %d", id)
	}
	return result, err
}

func (c *Client) GetSnapshotsByVolumeId(ctx context.Context, id int64) (result []types.Snapshot, err error) {
	req := types.ListSnapshotsRequest{
		VolumeID: id,
	}
	return c.ListSnapshots(ctx, req)
}
