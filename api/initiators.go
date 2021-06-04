package api

import (
	"context"

	"github.com/pkg/errors"
)

const (
	ErrNoInitiatorFound = "No Initiator found for the given id"
)

func (c *Client) CreateInitiators(ctx context.Context, initiators []CreateInitiator) (results []Initiator, err error) {
	req := CreateInitiatorsRequest{
		Initiators: initiators,
	}
	ciResult := CreateInitiatorsResult{}
	err = c.request(ctx, "CreateInitiators", req, &ciResult)
	results = ciResult.Initiators
	return results, err
}

func (c *Client) ModifyInitiators(ctx context.Context, req []ModifyInitiator) (results []Initiator, err error) {
	modReq := ModifyInitiatorsRequest{
		Initiators: req,
	}
	miResult := ModifyInitiatorsResult{}
	err = c.request(ctx, "ModifyInitiators", modReq, &miResult)
	results = miResult.Initiators
	return results, err
}

func (c *Client) DeleteInitiators(ctx context.Context, ids []int64) (err error) {
	req := DeleteInitiatorsRequest{
		Initiators: ids,
	}
	err = c.request(ctx, "DeleteInitiators", req, nil)
	return err
}

func (c *Client) ListInitiators(ctx context.Context, req ListInitiatorsRequest) (results []Initiator, err error) {
	liResult := ListInitiatorsResult{}
	err = c.request(ctx, "ListInitiators", req, &liResult)
	results = liResult.Initiators
	return results, err
}

func (c *Client) GetInitiator(ctx context.Context, id int64) (result *Initiator, err error) {
	req := ListInitiatorsRequest{
		Initiators: []int64{id},
	}
	initiators, err := c.ListInitiators(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(initiators) > 0 {
		result = &initiators[0]
		return result, err
	} else {
		return nil, errors.New(ErrNoInitiatorFound)
	}
}
