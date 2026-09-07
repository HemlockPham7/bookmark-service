package link

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// CreateShortenLink creates a shortened link for the specified URL and stores it with an expiration time.
//
// Parameters:
//   - ctx: the context used to control the lifetime of the operation.
//   - url: the original URL to shorten.
//   - expSecond: the number of seconds before the shortened link expires.
//
// Returns:
//   - The generated shortened link code.
//   - An error if the code cannot be generated or the URL cannot be stored.
func (s *linkService) CreateShortenLink(ctx context.Context, url string, expSecond int64) (string, error) {
	span := newrelic.FromContext(ctx).StartSegment("CreateShortenLink_LinkService")
	defer span.End()

	// tao code
	code, err := s.codeGenerator.GenerateCode(codeLength)
	if err != nil {
		return "", err
	}
	// goi repo de store url
	err = s.linkRepository.StoreURL(ctx, code, url, expSecond)
	if err != nil {
		return "", err
	}
	// return code
	return code, nil
}
