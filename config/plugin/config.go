package plugin

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sentry_plugin", func(r *config.Resource) {
		r.References["project"] = config.Reference{
			TerraformName: "sentry_project",
		}
	})
}
