package link

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

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
