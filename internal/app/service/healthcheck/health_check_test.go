package healthcheck

import (
	"context"
	"testing"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	mocksHealthCheck "github.com/HemlockPham7/bookmark-service/internal/app/repository/healthcheck/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMockPing func(ctx context.Context) *mocksHealthCheck.Repository

		expectedResp *model.HealthCheckResponse
		expectedErr  error
	}{
		{
			name: "health check successfully",

			setupMockPing: func(ctx context.Context) *mocksHealthCheck.Repository {
				mockPing := mocksHealthCheck.NewRepository(t)
				mockPing.On("RedisPing", ctx).Return(nil)
				return mockPing
			},

			expectedResp: &model.HealthCheckResponse{
				Message:     "OK",
				ServiceName: "bookmark-service",
				InstanceID:  "instance-1",
			},
			expectedErr: nil,
		},
		{
			name: "health check failed",

			setupMockPing: func(ctx context.Context) *mocksHealthCheck.Repository {
				mockPing := mocksHealthCheck.NewRepository(t)
				mockPing.
					On("RedisPing", ctx).
					Return(assert.AnError)

				return mockPing
			},

			expectedResp: nil,
			expectedErr:  assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			mockPing := tc.setupMockPing(ctx)
			testService := NewHealthCheckService("bookmark-service", "instance-1", mockPing)

			resp, err := testService.HealthCheck(ctx)
			assert.Equal(t, tc.expectedResp, resp)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
