package bookmark

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mock_bookmark "github.com/HemlockPham7/bookmark-service/internal/app/service/bookmark/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestBookmarkHandler_UpdateBookmarkByID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupRequest func(ctx *gin.Context)
		setupMockSvc func(ctx *gin.Context) *mock_bookmark.Service

		expectedCode     int
		expectedResponse string
	}{
		{
			name: "successful update bookmark by ID",

			setupRequest: func(ctx *gin.Context) {
				setupUpdateBookmarkRequest(ctx, "d7c13097-67a7-4eae-a60e-0b9b533b7bd6", "updated description", "updated url", true)
			},

			setupMockSvc: func(ctx *gin.Context) *mock_bookmark.Service {
				mockService := mock_bookmark.NewService(t)
				mockService.On("UpdateBookmarkByID", ctx, "updated description", "updated url", "d7c13097-67a7-4eae-a60e-0b9b533b7bd4", "d7c13097-67a7-4eae-a60e-0b9b533b7bd6").
					Return(nil, nil)
				return mockService
			},

			expectedCode:     http.StatusOK,
			expectedResponse: `{"message":"Success"}`,
		},
		{
			name: "claim not exist",

			setupRequest: func(ctx *gin.Context) {
				setupUpdateBookmarkRequest(ctx, "d7c13097-67a7-4eae-a60e-0b9b533b7bd6", "updated description", "updated url", false)
			},

			setupMockSvc: func(ctx *gin.Context) *mock_bookmark.Service {
				return mock_bookmark.NewService(t)
			},

			expectedCode:     http.StatusUnauthorized,
			expectedResponse: `{"message":"claim not exist"}`,
		},
		{
			name: "Bookmark ID is required",

			setupRequest: func(ctx *gin.Context) {
				setupUpdateBookmarkRequest(ctx, "", "updated description", "updated url", true)
			},

			setupMockSvc: func(ctx *gin.Context) *mock_bookmark.Service {
				return mock_bookmark.NewService(t)
			},

			expectedCode:     http.StatusBadRequest,
			expectedResponse: `{"message":"Bookmark ID is required"}`,
		},
		{
			name: "internal error",

			setupRequest: func(ctx *gin.Context) {
				setupUpdateBookmarkRequest(ctx, "d7c13097-67a7-4eae-a60e-0b9b533b7bd6", "updated description", "updated url", true)
			},

			setupMockSvc: func(ctx *gin.Context) *mock_bookmark.Service {
				mockService := mock_bookmark.NewService(t)
				mockService.On("UpdateBookmarkByID", ctx, "updated description", "updated url", "d7c13097-67a7-4eae-a60e-0b9b533b7bd4", "d7c13097-67a7-4eae-a60e-0b9b533b7bd6").
					Return(nil, assert.AnError)
				return mockService
			},

			expectedCode:     http.StatusInternalServerError,
			expectedResponse: `{"message":"Processing Error"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)

			tc.setupRequest(ctx)
			serviceMock := tc.setupMockSvc(ctx)

			handler := NewHandler(serviceMock, nil)

			handler.UpdateBookmarkByID(ctx)

			assert.Equal(t, tc.expectedCode, rec.Code)
			assert.Equal(t, tc.expectedResponse, strings.TrimSpace(rec.Body.String()))
		})
	}
}

func setupUpdateBookmarkRequest(ctx *gin.Context, bookmarkID, description, url string, haveClaim bool) {
	input := &updateBookmarkRequest{
		Description: description,
		URL:         url,
	}
	reqBody, _ := json.Marshal(input)

	target := fmt.Sprintf("/v1/bookmarks/%s", bookmarkID)

	ctx.Request = httptest.NewRequest(http.MethodPut, target, strings.NewReader(string(reqBody)))
	ctx.Request.Header.Set("Content-Type", "application/json")
	if bookmarkID != "" {
		ctx.Params = gin.Params{{Key: "id", Value: bookmarkID}}
	}

	if haveClaim {
		ctx.Set("claims", jwt.MapClaims{
			"sub": "d7c13097-67a7-4eae-a60e-0b9b533b7bd4",
		})
	}
}
