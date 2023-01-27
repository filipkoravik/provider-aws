/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/cloudwatchlogs/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Domain.
func (mg *Domain) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.LogPublishingOptions); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogPublishingOptions[i3].CloudwatchLogGroupArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.LogPublishingOptions[i3].CloudwatchLogGroupArnRef,
			Selector:     mg.Spec.ForProvider.LogPublishingOptions[i3].CloudwatchLogGroupArnSelector,
			To: reference.To{
				List:    &v1beta1.GroupList{},
				Managed: &v1beta1.Group{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.LogPublishingOptions[i3].CloudwatchLogGroupArn")
		}
		mg.Spec.ForProvider.LogPublishingOptions[i3].CloudwatchLogGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.LogPublishingOptions[i3].CloudwatchLogGroupArnRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this DomainPolicy.
func (mg *DomainPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DomainName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DomainNameRef,
		Selector:     mg.Spec.ForProvider.DomainNameSelector,
		To: reference.To{
			List:    &DomainList{},
			Managed: &Domain{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DomainName")
	}
	mg.Spec.ForProvider.DomainName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DomainNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DomainSAMLOptions.
func (mg *DomainSAMLOptions) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DomainName),
		Extract:      resource.ExtractParamPath("domain_name", false),
		Reference:    mg.Spec.ForProvider.DomainNameRef,
		Selector:     mg.Spec.ForProvider.DomainNameSelector,
		To: reference.To{
			List:    &DomainList{},
			Managed: &Domain{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DomainName")
	}
	mg.Spec.ForProvider.DomainName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DomainNameRef = rsp.ResolvedReference

	return nil
}
