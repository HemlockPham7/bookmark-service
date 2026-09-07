package link

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// GetLinkFromCode retrieves the original URL associated with a shortened link or bookmark code.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//   - requestCode: the shortened link or bookmark code used to retrieve the URL.
//
// Returns:
//   - The original URL associated with the code.
//   - ErrCodeNotFound if the code has an unsupported length or does not exist.
//   - An error if the URL or bookmark cannot be retrieved.
func (s *linkService) GetLinkFromCode(ctx context.Context, requestCode string) (string, error) {
	span := newrelic.FromContext(ctx).StartSegment("GetLinkFromCode_LinkService")
	defer span.End()

	switch {
	case len(requestCode) == codeLength:
		return s.linkRepository.GetURL(ctx, requestCode)
	case len(requestCode) == codeLengthBookmark:
		bookmark, err := s.bookmarkRepository.GetBookmarkByCode(ctx, requestCode)
		if err != nil {
			return "", err
		}
		return bookmark.URL, nil
	default:
		return "", ErrCodeNotFound
	}
}
