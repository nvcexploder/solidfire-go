package api

import (
	"context"

	"github.com/pkg/errors"
)

const (
	ErrNoVolumeAccessGroupFound = "No VolumeAccessGroup found for the given id"
)

func (c *Client) CreateVolumeAccessGroup(ctx context.Context, req CreateVolumeAccessGroupRequest) (result *VolumeAccessGroup, err error) {
	cvagResult := CreateVolumeAccessGroupResult{}
	err = c.request(ctx, "CreateVolumeAccessGroup", req, &cvagResult)
	result = &cvagResult.VolumeAccessGroup
	return result, err
}

func (c *Client) DeleteVolumeAccessGroup(ctx context.Context, req DeleteVolumeAccessGroupRequest) (err error) {
	return c.request(ctx, "DeleteVolumeAccessGroup", req, nil)
}

func (c *Client) ModifyVolumeAccessGroup(ctx context.Context, req ModifyVolumeAccessGroupRequest) (result *VolumeAccessGroup, err error) {
	mvagResult := ModifyVolumeAccessGroupResult{}
	err = c.request(ctx, "ModifyVolumeAccessGroup", req, &mvagResult)
	result = &mvagResult.VolumeAccessGroup
	return result, err
}

func (c *Client) ListVolumeAccessGroups(ctx context.Context, req ListVolumeAccessGroupsRequest) (result []VolumeAccessGroup, err error) {
	lvagResult := ListVolumeAccessGroupsResult{}
	err = c.request(ctx, "ListVolumeAccessGroups", req, &lvagResult)
	result = lvagResult.VolumeAccessGroups
	return result, err
}

func (c *Client) GetVolumeAccessGroup(ctx context.Context, id int64) (result *VolumeAccessGroup, err error) {
	req := ListVolumeAccessGroupsRequest{
		VolumeAccessGroups: []int64{id},
	}
	accessGroups, err := c.ListVolumeAccessGroups(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(accessGroups) > 0 {
		result = &accessGroups[0]
		return result, err
	} else {
		return nil, errors.New(ErrNoVolumeAccessGroupFound)
	}
}

func (c *Client) AddInitiatorsToVolumeAccessGroup(ctx context.Context, vagId int64, initiators []int64) (result *VolumeAccessGroup, err error) {
	req := AddInitiatorsToVolumeAccessGroupRequest{
		VolumeAccessGroupID: vagId,
		Initiators:          initiators,
	}
	aivr := AddInitiatorsToVolumeAccessGroupResult{}
	err = c.request(ctx, "AddInitiatorsToVolumeAccessGroup", req, &aivr)
	result = &aivr.VolumeAccessGroup
	return result, err
}

func (c *Client) AddVolumesToVolumeAccessGroup(ctx context.Context, vagId int64, volumes []int64) (result *VolumeAccessGroup, err error) {
	req := AddVolumesToVolumeAccessGroupRequest{
		VolumeAccessGroupID: vagId,
		Volumes:             volumes,
	}
	avvr := AddVolumesToVolumeAccessGroup{}
	err = c.request(ctx, "AddVolumesToVolumeAccessGroup", req, &avvr)
	result = &avvr.VolumeAccessGroup
	return result, err
}

func (c *Client) RemoveInitiatorsFromVolumeAccessGroup(ctx context.Context, vagId int64, initiators []int64, deleteOrphanInitiators bool) (result *VolumeAccessGroup, err error) {
	req := RemoveInitiatorsFromVolumeAccessGroupRequest{
		VolumeAccessGroupID:    vagId,
		Initiators:             initiators,
		DeleteOrphanInitiators: deleteOrphanInitiators,
	}
	rivr := RemoveInitiatorsFromVolumeAccessGroupResult{}
	err = c.request(ctx, "RemoveInitiatorsFromVolumeAccessGroup", req, &rivr)
	result = &rivr.VolumeAccessGroup
	return result, err
}

func (c *Client) RemoveVolumesFromVolumeAccessGroup(ctx context.Context, vagId int64, volumes []int64) (result *VolumeAccessGroup, err error) {
	req := RemoveVolumesFromVolumeAccessGroupRequest{
		VolumeAccessGroupID: vagId,
		Volumes:             volumes,
	}
	rvvr := RemoveVolumesFromVolumeAccessGroupResult{}
	err = c.request(ctx, "RemoveVolumesFromVolumeAccessGroup", req, &rvvr)
	result = &rvvr.VolumeAccessGroup
	return result, err
}
