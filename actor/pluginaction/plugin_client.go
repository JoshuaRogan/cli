package pluginaction

import "code.cloudfoundry.org/cli/api/plugin"

//go:generate counterfeiter . PluginClient

type PluginClient interface {
	GetPluginRepository(repositoryURL string) (plugin.PluginRepository, error)
}
