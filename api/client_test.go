package api

import (
	"context"

	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

var (
	defaultTarget   = "localhost"
	defaultUsername = "test-username"
	defaultPassword = "supersecret"
	defaultVersion  = "12.3"
	defaultPort     = 443
	defaultTimeout  = 10
)

func TestBuldClientErrors(t *testing.T) {
	var err error
	_, err = BuildClient("", defaultUsername, defaultPassword, defaultVersion, defaultPort, defaultTimeout)
	require.Equal(t, err.Error(), ErrNoTarget)
	_, err = BuildClient(defaultTarget, "", defaultPassword, defaultVersion, defaultPort, defaultTimeout)
	require.Equal(t, err.Error(), ErrNoCredentials)
	_, err = BuildClient(defaultTarget, defaultUsername, "", defaultVersion, defaultPort, defaultTimeout)
	require.Equal(t, err.Error(), ErrNoCredentials)
}

func TestBuldClient(t *testing.T) {
	c, err := BuildClient(defaultTarget, defaultUsername, defaultPassword, defaultVersion, 443, 0)
	require.Nil(t, err)
	require.Equal(t, c.ApiUrl, fmt.Sprintf("https://%s:%d/json-rpc/%s", defaultTarget, defaultPort, defaultVersion))
}

func TestClientRequestError(t *testing.T) {
	c := getTestClient(t)
	errName := "xExceededLimit"
	errMessage := "The snapMirrorLabel cannot be greater than 31 characters in length. snapMirrorLabel: sdk-test-asdkflkasdjflkasdjflkdjfdkjdkfjdfkdjkdjkdjdkjfdfdfdfdfdfdfdfdf"
	mockResp := SFResponse{
		Error: SFAPIError{
			Code:    500,
			Message: errMessage,
			Name:    errName,
		},
		Result: nil,
		Id:     1,
	}
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	req := CreateSnapshotRequest{
		VolumeID: testSnapshotVolumeId,
		Name:     testSnapshotName,
	}
	_, err := c.CreateSnapshot(ctx, req)

	var reqErr *RequestError
	require.True(t, errors.As(err, &reqErr))
	require.Equal(t, errName, reqErr.GetName())
	require.Equal(t, errMessage, reqErr.GetMessage())

}

func TestClientServiceError(t *testing.T) {
	c := getTestClient(t)
	errName := "xUnexpected"
	errMessage := "Unexpected error. Failed to reticulate splines in time. Try again in a moment"
	mockResp := SFResponse{
		Error: SFAPIError{
			Code:    500,
			Message: errMessage,
			Name:    errName,
		},
		Result: nil,
		Id:     1,
	}
	mockReset := activateMock(t, c, mockResp)
	defer mockReset()

	ctx := context.Background()
	req := CreateSnapshotRequest{
		VolumeID: testSnapshotVolumeId,
		Name:     testSnapshotName,
	}
	_, err := c.CreateSnapshot(ctx, req)

	var re *RequestError
	require.False(t, errors.As(err, &re))
	var serviceErr *ServiceError
	require.True(t, errors.As(err, &serviceErr))
	require.Equal(t, errName, serviceErr.GetName())
	require.Equal(t, errMessage, serviceErr.GetMessage())

}
