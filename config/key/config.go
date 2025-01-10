package key

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sentry_key", func(r *config.Resource) {
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}

			if a, ok := attr["dsn_csp"].(string); ok {
				conn["dsn_csp"] = []byte(a)
			}

			if a, ok := attr["dsn_public"].(string); ok {
				conn["dsn_public"] = []byte(a)
			}

			if a, ok := attr["dsn_secret"].(string); ok {
				conn["dsn_secret"] = []byte(a)
			}

			return conn, nil
		}
		r.References["project"] = config.Reference{
			TerraformName: "sentry_project",
		}
	})
}
