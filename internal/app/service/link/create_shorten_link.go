package link

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

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
