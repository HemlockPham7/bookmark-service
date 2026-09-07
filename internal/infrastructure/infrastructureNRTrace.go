package infrastructure

import (
	"github.com/HemlockPham7/common-libs/pkg/common"
	"github.com/HemlockPham7/common-libs/pkg/nrtrace"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// CreateNRClient creates a New Relic application using configuration loaded
// from environment variables.
//
// Parameters:
//   - envPrefix: the environment variable prefix used to load New Relic configuration.
//
// Returns:
//   - A configured New Relic application.
func CreateNRClient() *newrelic.Application {
	nrClient, err := nrtrace.NewClient("bookmark")
	common.HandleError(err)
	return nrClient
}
