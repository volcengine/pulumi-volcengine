// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package transit_router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of transit router route table associations
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
//			_, err := transit_router.RouteTableAssociations(ctx, &transit_router.RouteTableAssociationsArgs{
//				TransitRouterAttachmentId: pulumi.StringRef("tr-attach-im73ng3n5kao8gbssz2ddpuq"),
//				TransitRouterRouteTableId: "tr-rtb-12b7qd3fmzf2817q7y2jkbd55",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func RouteTableAssociations(ctx *pulumi.Context, args *RouteTableAssociationsArgs, opts ...pulumi.InvokeOption) (*RouteTableAssociationsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv RouteTableAssociationsResult
	err := ctx.Invoke("volcengine:transit_router/routeTableAssociations:RouteTableAssociations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking RouteTableAssociations.
type RouteTableAssociationsArgs struct {
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the network instance connection.
	TransitRouterAttachmentId *string `pulumi:"transitRouterAttachmentId"`
	// The ID of the routing table associated with the transit router instance.
	TransitRouterRouteTableId string `pulumi:"transitRouterRouteTableId"`
}

// A collection of values returned by RouteTableAssociations.
type RouteTableAssociationsResult struct {
	// The list of route table associations.
	Associations []RouteTableAssociationsAssociation `pulumi:"associations"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	OutputFile *string `pulumi:"outputFile"`
	// The total count of data query.
	TotalCount int `pulumi:"totalCount"`
	// The ID of the network instance connection.
	TransitRouterAttachmentId *string `pulumi:"transitRouterAttachmentId"`
	// The ID of the routing table associated with the transit router instance.
	TransitRouterRouteTableId string `pulumi:"transitRouterRouteTableId"`
}

func RouteTableAssociationsOutput(ctx *pulumi.Context, args RouteTableAssociationsOutputArgs, opts ...pulumi.InvokeOption) RouteTableAssociationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (RouteTableAssociationsResult, error) {
			args := v.(RouteTableAssociationsArgs)
			r, err := RouteTableAssociations(ctx, &args, opts...)
			var s RouteTableAssociationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(RouteTableAssociationsResultOutput)
}

// A collection of arguments for invoking RouteTableAssociations.
type RouteTableAssociationsOutputArgs struct {
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of the network instance connection.
	TransitRouterAttachmentId pulumi.StringPtrInput `pulumi:"transitRouterAttachmentId"`
	// The ID of the routing table associated with the transit router instance.
	TransitRouterRouteTableId pulumi.StringInput `pulumi:"transitRouterRouteTableId"`
}

func (RouteTableAssociationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTableAssociationsArgs)(nil)).Elem()
}

// A collection of values returned by RouteTableAssociations.
type RouteTableAssociationsResultOutput struct{ *pulumi.OutputState }

func (RouteTableAssociationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTableAssociationsResult)(nil)).Elem()
}

func (o RouteTableAssociationsResultOutput) ToRouteTableAssociationsResultOutput() RouteTableAssociationsResultOutput {
	return o
}

func (o RouteTableAssociationsResultOutput) ToRouteTableAssociationsResultOutputWithContext(ctx context.Context) RouteTableAssociationsResultOutput {
	return o
}

// The list of route table associations.
func (o RouteTableAssociationsResultOutput) Associations() RouteTableAssociationsAssociationArrayOutput {
	return o.ApplyT(func(v RouteTableAssociationsResult) []RouteTableAssociationsAssociation { return v.Associations }).(RouteTableAssociationsAssociationArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o RouteTableAssociationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RouteTableAssociationsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o RouteTableAssociationsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteTableAssociationsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of data query.
func (o RouteTableAssociationsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v RouteTableAssociationsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The ID of the network instance connection.
func (o RouteTableAssociationsResultOutput) TransitRouterAttachmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteTableAssociationsResult) *string { return v.TransitRouterAttachmentId }).(pulumi.StringPtrOutput)
}

// The ID of the routing table associated with the transit router instance.
func (o RouteTableAssociationsResultOutput) TransitRouterRouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v RouteTableAssociationsResult) string { return v.TransitRouterRouteTableId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RouteTableAssociationsResultOutput{})
}