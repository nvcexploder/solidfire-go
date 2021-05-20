package api

import (
	"context"
	"fmt"

	"github.com/joyent/solidfire-sdk/types"
)

func (c *Client) CreateVolume(ctx context.Context, req types.CreateVolumeRequest) (result *types.CreateVolumeResult, err error) {
	err = c.request(ctx, "CreateVolume", req, &result)
	return result, err
}

func (c *Client) ModifyVolume(ctx context.Context, req types.ModifyVolumeRequest) (result *types.Volume, err error) {
	if req.CreateTime != "" {
		req.SetCreateTime = true
	}
	mvr := types.ModifyVolumeResult{}
	err = c.request(ctx, "ModifyVolume", req, &mvr)
	result = &mvr.Volume
	return result, err
}

func (c *Client) DeleteVolume(ctx context.Context, id int64) (result *types.Volume, err error) {
	req := types.DeleteVolumeRequest{
		VolumeID: id,
	}
	dvr := types.DeleteVolumeResult{}
	err = c.request(ctx, "DeleteVolume", req, &dvr)
	result = &dvr.Volume
	return result, err
}

func (c *Client) ListVolumes(ctx context.Context, req types.ListVolumesRequest) (result []types.Volume, err error) {
	lvr := types.ListVolumesResult{}
	err = c.request(ctx, "ListVolumes", req, &lvr)
	result = lvr.Volumes
	return result, err
}

func (c *Client) GetVolumeById(ctx context.Context, id int64) (result *types.Volume, err error) {
	req := types.ListVolumesRequest{
		VolumeIDs: []int64{id},
	}
	lvr := types.ListVolumesResult{}
	err = c.request(ctx, "ListVolumes", req, &lvr)
	if len(lvr.Volumes) > 0 {
		result = &lvr.Volumes[0]
	} else if err == nil {
		err = fmt.Errorf("No volume found with the given id: %d", id)
	}
	return result, err
}

func (c *Client) ListVolumeStats(ctx context.Context, ids []int64) (result []types.VolumeStats, err error) {
	req := types.ListVolumeStatsRequest{
		VolumeIDs: ids,
	}
	lvsr := types.ListVolumeStatsResult{}
	err = c.request(ctx, "ListVolumeStats", req, &lvsr)
	result = lvsr.VolumeStats
	return result, err
}
