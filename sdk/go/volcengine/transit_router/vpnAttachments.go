// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package transit_router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of transit router vpn attachments
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/transit_router"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := transit_router.VpnAttachments(ctx, &transit_router.VpnAttachmentsArgs{
//				Ids: []string{
//					"tr-attach-3rf2xi7ae6y9s5zsk2hm6pibt",
//				},
//				TransitRouterId: "tr-2d6fr7f39unsw58ozfe1ow21x",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func VpnAttachments(ctx *pulumi.Context, args *VpnAttachmentsArgs, opts ...pulumi.InvokeOption) (*VpnAttachmentsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv VpnAttachmentsResult
	err := ctx.Invoke("volcengine:transit_router/vpnAttachments:VpnAttachments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking VpnAttachments.
type VpnAttachmentsArgs struct {
	// The ID list of the VPN attachment.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The id of the transit router.
	TransitRouterId string `pulumi:"transitRouterId"`
	// The ID of the IPSec connection.
	VpnConnectionId *string `pulumi:"vpnConnectionId"`
}

// A collection of values returned by VpnAttachments.
type VpnAttachmentsResult struct {
	// The collection of query.
	Attachments []VpnAttachmentsAttachment `pulumi:"attachments"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
	// The id of the transit router.
	TransitRouterId string `pulumi:"transitRouterId"`
	// The ID of the IPSec connection.
	VpnConnectionId *string `pulumi:"vpnConnectionId"`
}

func VpnAttachmentsOutput(ctx *pulumi.Context, args VpnAttachmentsOutputArgs, opts ...pulumi.InvokeOption) VpnAttachmentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (VpnAttachmentsResult, error) {
			args := v.(VpnAttachmentsArgs)
			r, err := VpnAttachments(ctx, &args, opts...)
			var s VpnAttachmentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(VpnAttachmentsResultOutput)
}

// A collection of arguments for invoking VpnAttachments.
type VpnAttachmentsOutputArgs struct {
	// The ID list of the VPN attachment.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The id of the transit router.
	TransitRouterId pulumi.StringInput `pulumi:"transitRouterId"`
	// The ID of the IPSec connection.
	VpnConnectionId pulumi.StringPtrInput `pulumi:"vpnConnectionId"`
}

func (VpnAttachmentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnAttachmentsArgs)(nil)).Elem()
}

// A collection of values returned by VpnAttachments.
type VpnAttachmentsResultOutput struct{ *pulumi.OutputState }

func (VpnAttachmentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnAttachmentsResult)(nil)).Elem()
}

func (o VpnAttachmentsResultOutput) ToVpnAttachmentsResultOutput() VpnAttachmentsResultOutput {
	return o
}

func (o VpnAttachmentsResultOutput) ToVpnAttachmentsResultOutputWithContext(ctx context.Context) VpnAttachmentsResultOutput {
	return o
}

// The collection of query.
func (o VpnAttachmentsResultOutput) Attachments() VpnAttachmentsAttachmentArrayOutput {
	return o.ApplyT(func(v VpnAttachmentsResult) []VpnAttachmentsAttachment { return v.Attachments }).(VpnAttachmentsAttachmentArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o VpnAttachmentsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VpnAttachmentsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o VpnAttachmentsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VpnAttachmentsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o VpnAttachmentsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnAttachmentsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o VpnAttachmentsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v VpnAttachmentsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The id of the transit router.
func (o VpnAttachmentsResultOutput) TransitRouterId() pulumi.StringOutput {
	return o.ApplyT(func(v VpnAttachmentsResult) string { return v.TransitRouterId }).(pulumi.StringOutput)
}

// The ID of the IPSec connection.
func (o VpnAttachmentsResultOutput) VpnConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnAttachmentsResult) *string { return v.VpnConnectionId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VpnAttachmentsResultOutput{})
}