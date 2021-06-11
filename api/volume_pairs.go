package api

import (
	"context"
	"fmt"
)

const (
	VolumePairingModeAsync         = "Async"
	VolumePairingModeSync          = "Sync"
	VolumePairingModeSnapshotsOnly = "SnapshotsOnly"
)

func (c *Client) StartVolumePairing(ctx context.Context, volId int64, mode string) (volumePairingKey string, err error) {
	req := StartVolumePairingRequest{
		VolumeID: volId,
		Mode:     mode,
	}
	result := StartVolumePairingResult{}
	err = c.request(ctx, "StartVolumePairing", req, &result)
	volumePairingKey = result.VolumePairingKey
	return volumePairingKey, err
}

func (c *Client) CompleteVolumePairing(ctx context.Context, volId int64, volumePairingKey string) (err error) {
	req := CompleteVolumePairingRequest{
		VolumeID:         volId,
		VolumePairingKey: volumePairingKey,
	}
	return c.request(ctx, "CompleteVolumePairing", req, nil)
}

func (c *Client) ModifyVolumePair(ctx context.Context, req ModifyVolumePairRequest) (err error) {
	return c.request(ctx, "ModifyVolumePair", req, nil)
}

func (c *Client) RemoveVolumePair(ctx context.Context, volId int64) (err error) {
	req := RemoveVolumePairRequest{
		VolumeID: volId,
	}
	return c.request(ctx, "RemoveVolumePair", req, nil)
}

func (c *Client) ListActivePairedVolumes(ctx context.Context, req ListActivePairedVolumesRequest) (resp []Volume, err error) {
	result := ListActivePairedVolumesResult{}
	err = c.request(ctx, "ListActivePairedVolumes", req, &result)
	resp = result.Volumes
	return resp, err
}

func (c *Client) GetActivePairedVolume(ctx context.Context, volId int64) (resp *Volume, err error) {
	req := ListActivePairedVolumesRequest{
		StartVolumeID: volId,
		Limit:         1,
	}
	res, err := c.ListActivePairedVolumes(ctx, req)
	if err != nil || len(res) == 0 {
		err = BuildRequestError(ErrVolumeIDDoesNotExist, fmt.Sprintf("Active paired volume with the given id %d does not exist", volId))
	} else if err == nil {
		resp = &res[0]
	}
	return resp, err
}
