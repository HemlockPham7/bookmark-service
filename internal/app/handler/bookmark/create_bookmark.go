package bookmark

import (
	"errors"
	"net/http"

	"github.com/HemlockPham7/bookmark-service/internal/app/model"
	"github.com/HemlockPham7/common-libs/pkg/dbutils"
	"github.com/HemlockPham7/common-libs/pkg/requestutils"
	"github.com/HemlockPham7/common-libs/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog/log"
)

type createBookmarkRequest struct {
	Description string `json:"description" example:"Google" validate:"lte=255"`
	URL         string `json:"url" example:"https://www.google.com" validate:"required,url,lte=2048"`
}

type createBookmarkResponse struct {
	Data    *model.Bookmark `json:"data"`
	Message string          `json:"message"`
}

// CreateBookmark generates a Gin framework handler that creates a new bookmark for the authenticated user.
// @Summary      Create a new bookmark
// @Description  Create a new bookmark for the authenticated user
// @Tags         Bookmarks
// @Accept       application/json
// @Produce      application/json
// @Param        request  body      createBookmarkRequest  true  "Create Bookmark Request"
// @Success      201      {object}  createBookmarkResponse
// @Failure      400      {object}  object{message=string}
// @Failure      401      {object}  object{message=string}
// @Failure      500      {object}  object{message=string}
// @Security     BearerAuth
// @Router       /v1/bookmarks [post]
func (h *bookmarkHandler) CreateBookmark(c *gin.Context) {
	span := newrelic.FromContext(c).StartSegment("CreateBookmark_BookmarkHandler")
	defer span.End()

	// get input
	input, uid, err := requestutils.BindInputFromRequestWithAuth[createBookmarkRequest](c)
	if err != nil {
		return
	}

	// call service
	bookmark, err := h.bookmarkService.CreateBookmark(c, input.Description, input.URL, uid)
	switch {
	case errors.Is(err, dbutils.ErrDuplicationType):
		c.JSON(http.StatusBadRequest, response.InputErrResponse)
		return
	case errors.Is(err, dbutils.ErrForeignKeyType):
		c.JSON(http.StatusUnauthorized, response.UnauthorizedResponse)
		return
	case errors.Is(err, nil):
	default:
		log.Error().
			Str("operation", "CreateBookmark").
			Err(err).
			Msg("service return error when create bookmark")
		c.JSON(http.StatusInternalServerError, response.InternalErrResponse)
		return
	}

	c.JSON(http.StatusCreated, &createBookmarkResponse{
		Data:    bookmark,
		Message: "Create a bookmark successfully!",
	})
}
