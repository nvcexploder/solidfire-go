package api

import (
	"context"

	"github.com/joyent/solidfire-sdk/types"
)

func (c *Client) ListActiveVolumes(ctx context.Context, req types.ListActiveVolumesRequest) (result *types.ListActiveVolumesResult, err error) {
	err = c.request(ctx, "ListActiveVolumes", req, &result)
	return result, err
}
