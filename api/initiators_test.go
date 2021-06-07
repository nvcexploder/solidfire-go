package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

const testInitiatorName = "iqn.1993-08.org.debian:01:c84ffd71216"
const testInitiatorId = int64(1)
const testInitiatorAlias = "solidfire-sdk-initiator-test-1"

var testInitiators = map[string]interface{}{
	"initiators": []map[string]interface{}{
		{
			"alias":              testInitiatorAlias,
			"attributes":         map[string]interface{}{},
			"initiatorID":        testInitiatorId,
			"initiatorName":      testInitiatorName,
			"volumeAccessGroups": []int64{},
		},
	},
}

func TestCreateInitiator(t *testing.T) {

	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(testInitiators)
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	cInits := []CreateInitiator{
		{Name: testInitiatorName, Alias: testInitiatorAlias},
	}
	resp, err := c.CreateInitiators(ctx, cInits)

	require.Nil(t, err)
	require.True(t, len(resp) > 0)
	respInitiator := resp[0]
	require.Equal(t, testInitiatorAlias, respInitiator.Alias)
	require.Equal(t, testInitiatorId, respInitiator.InitiatorID)
	require.Equal(t, testInitiatorName, respInitiator.InitiatorName)
}

func TestDeleteInitiator(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(nil)
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	err := c.DeleteInitiators(ctx, []int64{1})
	require.Nil(t, err)
}

func TestListInitiators(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(testInitiators)
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	req := ListInitiatorsRequest{Initiators: []int64{testInitiatorId}}
	resp, err := c.ListInitiators(ctx, req)
	require.Nil(t, err)
	require.True(t, len(resp) > 0)
	respInitiator := resp[0]
	require.Equal(t, testInitiatorAlias, respInitiator.Alias)
	require.Equal(t, testInitiatorId, respInitiator.InitiatorID)
	require.Equal(t, testInitiatorName, respInitiator.InitiatorName)
}

func TestGetInitiators(t *testing.T) {
	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(testInitiators)
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	resp, err := c.GetInitiator(ctx, testInitiatorId)
	require.Nil(t, err)
	require.Equal(t, testInitiatorAlias, resp.Alias)
	require.Equal(t, testInitiatorId, resp.InitiatorID)
	require.Equal(t, testInitiatorName, resp.InitiatorName)
}

func TestModifyInitiators(t *testing.T) {
	const testInitiatorAlias2 = "solidfire-sdk-initiator-test-2"

	var testInitiators2 = map[string]interface{}{
		"initiators": []map[string]interface{}{
			{
				"alias":              testInitiatorAlias2,
				"attributes":         map[string]interface{}{},
				"initiatorID":        testInitiatorId,
				"initiatorName":      testInitiatorName,
				"volumeAccessGroups": []int64{},
			},
		},
	}

	c := getTestClient(t)
	mockResp := buildSFResponseWrapper(testInitiators2)
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	mInits := []ModifyInitiator{
		{InitiatorID: testInitiatorId, Alias: testInitiatorAlias2},
	}
	resp, err := c.ModifyInitiators(ctx, mInits)

	require.Nil(t, err)
	require.True(t, len(resp) > 0)
	respInitiator := resp[0]
	require.Equal(t, testInitiatorAlias2, respInitiator.Alias)
	require.Equal(t, testInitiatorId, respInitiator.InitiatorID)
	require.Equal(t, testInitiatorName, respInitiator.InitiatorName)
}
