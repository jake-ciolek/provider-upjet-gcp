// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *ConnectivityTest) ResolveReferences( // ResolveReferences of this ConnectivityTest.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Destination); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Address", "AddressList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Destination[i3].IPAddress),
				Extract:      resource.ExtractParamPath("address", false),
				Reference:    mg.Spec.ForProvider.Destination[i3].IPAddressRef,
				Selector:     mg.Spec.ForProvider.Destination[i3].IPAddressSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Destination[i3].IPAddress")
		}
		mg.Spec.ForProvider.Destination[i3].IPAddress = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Destination[i3].IPAddressRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Destination); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Instance", "InstanceList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Destination[i3].Instance),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Destination[i3].InstanceRef,
				Selector:     mg.Spec.ForProvider.Destination[i3].InstanceSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Destination[i3].Instance")
		}
		mg.Spec.ForProvider.Destination[i3].Instance = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Destination[i3].InstanceRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Destination); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Destination[i3].Network),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Destination[i3].NetworkRef,
				Selector:     mg.Spec.ForProvider.Destination[i3].NetworkSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Destination[i3].Network")
		}
		mg.Spec.ForProvider.Destination[i3].Network = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Destination[i3].NetworkRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Destination); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Address", "AddressList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Destination[i3].ProjectID),
				Extract:      resource.ExtractParamPath("project", false),
				Reference:    mg.Spec.ForProvider.Destination[i3].ProjectIDRef,
				Selector:     mg.Spec.ForProvider.Destination[i3].ProjectIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Destination[i3].ProjectID")
		}
		mg.Spec.ForProvider.Destination[i3].ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Destination[i3].ProjectIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Source); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Address", "AddressList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Source[i3].IPAddress),
				Extract:      resource.ExtractParamPath("address", false),
				Reference:    mg.Spec.ForProvider.Source[i3].IPAddressRef,
				Selector:     mg.Spec.ForProvider.Source[i3].IPAddressSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Source[i3].IPAddress")
		}
		mg.Spec.ForProvider.Source[i3].IPAddress = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Source[i3].IPAddressRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Source); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Instance", "InstanceList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Source[i3].Instance),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Source[i3].InstanceRef,
				Selector:     mg.Spec.ForProvider.Source[i3].InstanceSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Source[i3].Instance")
		}
		mg.Spec.ForProvider.Source[i3].Instance = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Source[i3].InstanceRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Source); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Source[i3].Network),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Source[i3].NetworkRef,
				Selector:     mg.Spec.ForProvider.Source[i3].NetworkSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Source[i3].Network")
		}
		mg.Spec.ForProvider.Source[i3].Network = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Source[i3].NetworkRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Source); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Address", "AddressList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Source[i3].ProjectID),
				Extract:      resource.ExtractParamPath("project", false),
				Reference:    mg.Spec.ForProvider.Source[i3].ProjectIDRef,
				Selector:     mg.Spec.ForProvider.Source[i3].ProjectIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Source[i3].ProjectID")
		}
		mg.Spec.ForProvider.Source[i3].ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Source[i3].ProjectIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Destination); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Address", "AddressList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Destination[i3].IPAddress),
				Extract:      resource.ExtractParamPath("address", false),
				Reference:    mg.Spec.InitProvider.Destination[i3].IPAddressRef,
				Selector:     mg.Spec.InitProvider.Destination[i3].IPAddressSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Destination[i3].IPAddress")
		}
		mg.Spec.InitProvider.Destination[i3].IPAddress = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Destination[i3].IPAddressRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Destination); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Instance", "InstanceList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Destination[i3].Instance),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Destination[i3].InstanceRef,
				Selector:     mg.Spec.InitProvider.Destination[i3].InstanceSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Destination[i3].Instance")
		}
		mg.Spec.InitProvider.Destination[i3].Instance = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Destination[i3].InstanceRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Destination); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Destination[i3].Network),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Destination[i3].NetworkRef,
				Selector:     mg.Spec.InitProvider.Destination[i3].NetworkSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Destination[i3].Network")
		}
		mg.Spec.InitProvider.Destination[i3].Network = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Destination[i3].NetworkRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Destination); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Address", "AddressList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Destination[i3].ProjectID),
				Extract:      resource.ExtractParamPath("project", false),
				Reference:    mg.Spec.InitProvider.Destination[i3].ProjectIDRef,
				Selector:     mg.Spec.InitProvider.Destination[i3].ProjectIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Destination[i3].ProjectID")
		}
		mg.Spec.InitProvider.Destination[i3].ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Destination[i3].ProjectIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Source); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Address", "AddressList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Source[i3].IPAddress),
				Extract:      resource.ExtractParamPath("address", false),
				Reference:    mg.Spec.InitProvider.Source[i3].IPAddressRef,
				Selector:     mg.Spec.InitProvider.Source[i3].IPAddressSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Source[i3].IPAddress")
		}
		mg.Spec.InitProvider.Source[i3].IPAddress = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Source[i3].IPAddressRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Source); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Instance", "InstanceList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Source[i3].Instance),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Source[i3].InstanceRef,
				Selector:     mg.Spec.InitProvider.Source[i3].InstanceSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Source[i3].Instance")
		}
		mg.Spec.InitProvider.Source[i3].Instance = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Source[i3].InstanceRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Source); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Source[i3].Network),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Source[i3].NetworkRef,
				Selector:     mg.Spec.InitProvider.Source[i3].NetworkSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Source[i3].Network")
		}
		mg.Spec.InitProvider.Source[i3].Network = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Source[i3].NetworkRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Source); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Address", "AddressList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Source[i3].ProjectID),
				Extract:      resource.ExtractParamPath("project", false),
				Reference:    mg.Spec.InitProvider.Source[i3].ProjectIDRef,
				Selector:     mg.Spec.InitProvider.Source[i3].ProjectIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Source[i3].ProjectID")
		}
		mg.Spec.InitProvider.Source[i3].ProjectID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Source[i3].ProjectIDRef = rsp.ResolvedReference

	}

	return nil
}
