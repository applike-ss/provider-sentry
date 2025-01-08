package issuealert

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sentry_issue_alert", func(r *config.Resource) {
		r.ShortGroup = "sentry"
		r.Kind = "IssueAlert"
		r.References["project"] = config.Reference{
			TerraformName: "sentry_project",
		}
	})
}
