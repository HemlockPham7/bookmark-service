package gencode

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HemlockPham7/bookmark-service/internal/api"
	redisPkg "github.com/HemlockPham7/common-libs/pkg/redis"
	"github.com/HemlockPham7/common-libs/pkg/sqldb"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGenCodeEndpoint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupTestHTTP func(api api.Engine) *httptest.ResponseRecorder

		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "normal case",

			setupTestHTTP: func(api api.Engine) *httptest.ResponseRecorder {
				req := httptest.NewRequest("GET", "/gencode", nil)
				respRecorder := httptest.NewRecorder()

				api.ServeHTTP(respRecorder, req)
				return respRecorder
			},

			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"code":`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			setupRedisClient := redisPkg.InitMockRedis(t)
			setupDB := sqldb.InitMockDB(t)
			testAPI := api.NewEngine(&api.EngineOpts{
				App:         gin.Default(),
				Cfg:         &api.Config{},
				RedisClient: setupRedisClient,
				DbClient:    setupDB,
			})
			recorder := tc.setupTestHTTP(testAPI)

			assert.Equal(t, tc.expectedStatusCode, recorder.Code)
			assert.Contains(t, recorder.Body.String(), tc.expectedResponseBody)
		})
	}
}
