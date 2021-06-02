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

func TestBuildClientErrors(t *testing.T) {
	var err error
	_, err = BuildClient("", defaultUsername, defaultPassword, defaultVersion, defaultPort, defaultTimeout)
	require.Equal(t, err.Error(), ErrNoTarget)
	_, err = BuildClient(defaultTarget, "", defaultPassword, defaultVersion, defaultPort, defaultTimeout)
	require.Equal(t, err.Error(), ErrNoCredentials)
	_, err = BuildClient(defaultTarget, defaultUsername, "", defaultVersion, defaultPort, defaultTimeout)
	require.Equal(t, err.Error(), ErrNoCredentials)
}

func TestBuildClient(t *testing.T) {
	c, err := BuildClient(defaultTarget, defaultUsername, defaultPassword, defaultVersion, 443, 0)
	require.Nil(t, err)
	require.Equal(t, c.ApiUrl, fmt.Sprintf("https://%s:%d/json-rpc/%s", defaultTarget, defaultPort, defaultVersion))
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
