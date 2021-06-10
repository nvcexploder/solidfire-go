package api

import (
	"context"
	"fmt"
)

const (
	// Valid volume access values
	VolumeAccessPolicyReadOnly          = "readOnly"
	VolumeAccessPolicyReadWrite         = "readWrite"
	VolumeAccessPolicyLocked            = "locked"
	VolumeAccessPolicyReplicationTarget = "replicationTarget"
)

func (c *Client) CreateVolume(ctx context.Context, req CreateVolumeRequest) (result *Volume, err error) {
	cvr := CreateVolumeResult{}
	err = c.request(ctx, "CreateVolume", req, &cvr)
	result = &cvr.Volume
	return result, err
}

func (c *Client) ModifyVolume(ctx context.Context, req ModifyVolumeRequest) (result *Volume, err error) {
	mvr := ModifyVolumeResult{}
	err = c.request(ctx, "ModifyVolume", req, &mvr)
	result = &mvr.Volume
	return result, err
}

func (c *Client) DeleteVolume(ctx context.Context, id int64) (result *Volume, err error) {
	req := DeleteVolumeRequest{
		VolumeID: id,
	}
	dvr := DeleteVolumeResult{}
	err = c.request(ctx, "DeleteVolume", req, &dvr)
	result = &dvr.Volume
	return result, err
}

func (c *Client) ListVolumes(ctx context.Context, req ListVolumesRequest) (result []Volume, err error) {
	lvr := ListVolumesResult{}
	err = c.request(ctx, "ListVolumes", req, &lvr)
	result = lvr.Volumes
	return result, err
}

func (c *Client) GetVolumeById(ctx context.Context, id int64) (result *Volume, err error) {
	req := ListVolumesRequest{
		VolumeIDs: []int64{id},
	}
	lvr := ListVolumesResult{}
	err = c.request(ctx, "ListVolumes", req, &lvr)
	if len(lvr.Volumes) > 0 {
		result = &lvr.Volumes[0]
	} else if err == nil {
		err = BuildRequestError(ErrVolumeIDDoesNotExist, fmt.Sprintf("Volume with the given id %d does not exist", id))
	}
	return result, err
}

func (c *Client) ListVolumeStats(ctx context.Context, ids []int64) (result []VolumeStats, err error) {
	req := ListVolumeStatsRequest{
		VolumeIDs: ids,
	}
	lvsr := ListVolumeStatsResult{}
	err = c.request(ctx, "ListVolumeStats", req, &lvsr)
	result = lvsr.VolumeStats
	return result, err
}
