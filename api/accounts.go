package api

import (
	"context"

	"github.com/joyent/solidfire-sdk/types"
)

func (c *Client) ListAccounts(ctx context.Context, req types.ListAccountsRequest) (result []types.Account, err error) {
	lar := types.ListAccountsResult{}
	err = c.request(ctx, "ListAccounts", req, &lar)
	result = lar.Accounts
	return result, err
}
