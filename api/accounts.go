package api

import (
	"context"
)

func (c *Client) ListAccounts(ctx context.Context, req ListAccountsRequest) (result []Account, err error) {
	lar := ListAccountsResult{}
	err = c.request(ctx, "ListAccounts", req, &lar)
	result = lar.Accounts
	return result, err
}
