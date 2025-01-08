// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this IssueAlert.
func (mg *IssueAlert) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Project),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ProjectRef,
		Selector:     mg.Spec.ForProvider.ProjectSelector,
		To: reference.To{
			List:    &ProjectList{},
			Managed: &Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Project")
	}
	mg.Spec.ForProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Project),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ProjectRef,
		Selector:     mg.Spec.InitProvider.ProjectSelector,
		To: reference.To{
			List:    &ProjectList{},
			Managed: &Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Project")
	}
	mg.Spec.InitProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProjectRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Key.
func (mg *Key) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Project),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ProjectRef,
		Selector:     mg.Spec.ForProvider.ProjectSelector,
		To: reference.To{
			List:    &ProjectList{},
			Managed: &Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Project")
	}
	mg.Spec.ForProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Project),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ProjectRef,
		Selector:     mg.Spec.InitProvider.ProjectSelector,
		To: reference.To{
			List:    &ProjectList{},
			Managed: &Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Project")
	}
	mg.Spec.InitProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProjectRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Plugin.
func (mg *Plugin) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Project),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ProjectRef,
		Selector:     mg.Spec.ForProvider.ProjectSelector,
		To: reference.To{
			List:    &ProjectList{},
			Managed: &Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Project")
	}
	mg.Spec.ForProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Project),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ProjectRef,
		Selector:     mg.Spec.InitProvider.ProjectSelector,
		To: reference.To{
			List:    &ProjectList{},
			Managed: &Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Project")
	}
	mg.Spec.InitProvider.Project = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ProjectRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Project.
func (mg *Project) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Teams),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.TeamsRefs,
		Selector:      mg.Spec.ForProvider.TeamsSelector,
		To: reference.To{
			List:    &TeamList{},
			Managed: &Team{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Teams")
	}
	mg.Spec.ForProvider.Teams = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.TeamsRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Teams),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.TeamsRefs,
		Selector:      mg.Spec.InitProvider.TeamsSelector,
		To: reference.To{
			List:    &TeamList{},
			Managed: &Team{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Teams")
	}
	mg.Spec.InitProvider.Teams = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.TeamsRefs = mrsp.ResolvedReferences

	return nil
}