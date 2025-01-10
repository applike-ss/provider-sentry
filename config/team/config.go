package team

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("sentry_team", func(r *config.Resource) {
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}

			if a, ok := attr["internal_id"].(string); ok {
				conn["internal_id"] = []byte(a)
			}

			if a, ok := attr["slug"].(string); ok {
				conn["slug"] = []byte(a)
			}

			if a, ok := attr["team_id"].(string); ok {
				conn["team_id"] = []byte(a)
			}

			return conn, nil
		}
	})
}
