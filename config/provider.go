/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/justtrack/provider-sentry/config/issuealert"
	"github.com/justtrack/provider-sentry/config/key"
	"github.com/justtrack/provider-sentry/config/plugin"
	"github.com/justtrack/provider-sentry/config/project"
	"github.com/justtrack/provider-sentry/config/team"
)

const (
	resourcePrefix = "sentry"
	modulePath     = "github.com/justtrack/provider-sentry"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("justtrack.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		issuealert.Configure,
		key.Configure,
		plugin.Configure,
		project.Configure,
		team.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
