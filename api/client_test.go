package api

import (
	"context"
	"time"

	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

var (
	defaultTarget   = "localhost"
	defaultUsername = "test-username"
	defaultPassword = "supersecret"
	defaultVersion  = "12.3"
	defaultPort     = 443
)

func TestBuildClientErrors(t *testing.T) {
	var err error
	opts1 := ClientOptions{
		Target:   "",
		Username: defaultUsername,
		Password: defaultPassword,
	}
	_, err = BuildClient(opts1)
	require.NotNil(t, err)
	require.Equal(t, err.Error(), ErrNoTarget)
	opts2 := ClientOptions{
		Target:   defaultTarget,
		Username: "",
		Password: defaultPassword,
	}
	_, err = BuildClient(opts2)
	require.NotNil(t, err)
	require.Equal(t, err.Error(), ErrNoCredentials)
	opts3 := ClientOptions{
		Target:   defaultTarget,
		Username: defaultUsername,
		Password: "",
	}
	_, err = BuildClient(opts3)
	require.NotNil(t, err)
	require.Equal(t, err.Error(), ErrNoCredentials)
}

func TestBuildClient(t *testing.T) {
	opts := ClientOptions{
		Target:   defaultTarget,
		Username: defaultUsername,
		Password: defaultPassword,
	}
	c, err := BuildClient(opts)
	require.Nil(t, err)
	require.Equal(t, c.ApiUrl, fmt.Sprintf("https://%s:%d/json-rpc/%s", defaultTarget, defaultPort, defaultVersion))
}

func TestBuildClientNoRetriesRequestError(t *testing.T) {
	// Build a client that will (quickly) retry on service error
	retryCount := 1
	opts := ClientOptions{
		Target:           defaultTarget,
		Username:         defaultUsername,
		Password:         defaultPassword,
		TimeoutSecs:      time.Second * 10,
		UseRetry:         true,
		RetryCount:       retryCount,
		RetryWaitTime:    time.Millisecond * 1,
		RetryMaxWaitTime: time.Millisecond * 1,
	}
	c, err := BuildClient(opts)
	require.Nil(t, err)

	mockReset := activateMock(t, c, SFResponse{
		Error: SFAPIError{
			Code:    1,
			Name:    ErrUnrecognizedEnumString,
			Message: "Given Access value is invalid",
		},
		Result: nil,
		Id:     1,
	})
	defer mockReset()
	ctx := context.Background()
	req := CreateVolumeRequest{
		Name:       "testvolume1",
		AccountID:  1,
		TotalSize:  1 * Gigabytes,
		Enable512e: true,
		Access:     "notAValidAccessValue",
	}
	_, err = c.CreateVolume(ctx, req)
	callCount := httpmock.DefaultTransport.GetTotalCallCount()

	require.NotNil(t, err)
	var r *RequestError
	require.True(t, errors.As(err, &r))
	require.Equal(t, 1, callCount)
}

func TestClientRequestErrors(t *testing.T) {
	testCases := []struct {
		desc       string
		httpStatus int
		errMessage string
		errName    string
		verify     func(*testing.T, error)
	}{
		{
			desc:       "request error json",
			errMessage: "The snapMirrorLabel cannot be greater than 31 characters in length. snapMirrorLabel: sdk-test-asdkflkasdjflkasdjflkdjfdkjdkfjdfkdjkdjkdjdkjfdfdfdfdfdfdfdfdf",
			errName:    ErrExceededLimit,
			verify: func(t *testing.T, e error) {
				var r *RequestError
				require.True(t, errors.As(e, &r))
			},
		},
		{
			desc:       "service error json",
			errName:    "xUnexpected",
			errMessage: "Unexpected error. Failed to reticulate splines in time. Try again in a moment",
			verify: func(t *testing.T, e error) {
				var r *ServiceError
				require.True(t, errors.As(e, &r))
			},
		},
		{
			desc:       "auth error",
			httpStatus: 401,
			errName:    ErrInvalidCredentials,
			errMessage: "401",
			verify: func(t *testing.T, e error) {
				var r *RequestError
				require.True(t, errors.As(e, &r))
			},
		},
		{
			desc:       "unexpected http error",
			httpStatus: 503,
			errName:    ErrUnexpectedServerError,
			errMessage: "503",
			verify: func(t *testing.T, e error) {
				var r *ServiceError
				require.True(t, errors.As(e, &r))
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			c := getTestClient(t)

			var mockReset func()
			if tC.httpStatus > 0 {
				mockReset = activateMockHttpErr(c, tC.httpStatus)
			} else {
				mockReset = activateMock(t, c, SFResponse{
					Error: SFAPIError{
						Code:    500,
						Message: tC.errMessage,
						Name:    tC.errName,
					},
					Result: nil,
					Id:     1,
				})
			}
			defer mockReset()
			ctx := context.Background()
			req := CreateSnapshotRequest{
				VolumeID: testSnapshotVolumeId,
				Name:     testSnapshotName,
			}
			_, err := c.CreateSnapshot(ctx, req)

			tC.verify(t, err)
			sfe := err.(SFError)
			require.Equal(t, tC.errName, sfe.GetName())
			require.Equal(t, tC.errMessage, sfe.GetMessage())
		})
	}
}
