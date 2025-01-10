// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	providerconfig "github.com/justtrack/provider-sentry/internal/controller/providerconfig"
	issuealert "github.com/justtrack/provider-sentry/internal/controller/sentry/issuealert"
	key "github.com/justtrack/provider-sentry/internal/controller/sentry/key"
	plugin "github.com/justtrack/provider-sentry/internal/controller/sentry/plugin"
	project "github.com/justtrack/provider-sentry/internal/controller/sentry/project"
	team "github.com/justtrack/provider-sentry/internal/controller/sentry/team"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		issuealert.Setup,
		key.Setup,
		plugin.Setup,
		project.Setup,
		team.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
