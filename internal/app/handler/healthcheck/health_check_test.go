package healthcheck

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	mockHealthCheck "github.com/HemlockPham7/bookmark-service/internal/app/service/healthcheck/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck_HealthCheck(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupRequest     func(ctx *gin.Context)
		setupMockService func(ctx context.Context) *mockHealthCheck.Service

		expectedStatus int
		expectedBody   string
	}{
		{
			name: "health check successfully",

			setupRequest: func(ctx *gin.Context) {
				ctx.Request = httptest.NewRequest(http.MethodGet, "/health-check", nil)
			},
			setupMockService: func(ctx context.Context) *mockHealthCheck.Service {
				serviceMock := mockHealthCheck.NewService(t)
				serviceMock.On("HealthCheck", ctx).Return(&model.HealthCheckResponse{
					Message:     "OK",
					ServiceName: "bookmark-service",
					InstanceID:  "instsance-1",
				}, nil)
				return serviceMock
			},

			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"OK","service_name":"bookmark-service","instance_id":"instsance-1"}`,
		},
		{
			name: "health check failed",

			setupRequest: func(ctx *gin.Context) {
				ctx.Request = httptest.NewRequest(http.MethodGet, "/health-check", nil)
			},
			setupMockService: func(ctx context.Context) *mockHealthCheck.Service {
				serviceMock := mockHealthCheck.NewService(t)
				serviceMock.On("HealthCheck", ctx).Return(nil, assert.AnError)
				return serviceMock
			},

			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"Instance is not ready!"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)

			tc.setupRequest(ctx)
			mockService := tc.setupMockService(ctx)
			testHandler := NewHealthcheckHandler(mockService)
			testHandler.HealthCheck(ctx)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Equal(t, tc.expectedBody, rec.Body.String())

		})
	}
}
