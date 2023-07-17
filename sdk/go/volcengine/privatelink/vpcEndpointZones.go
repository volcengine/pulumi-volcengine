// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of privatelink vpc endpoint zones
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/privatelink"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := privatelink.VpcEndpointZones(ctx, &privatelink.VpcEndpointZonesArgs{
//				EndpointId: pulumi.StringRef("ep-2byz5npiuu1hc2dx0efkv****"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func VpcEndpointZones(ctx *pulumi.Context, args *VpcEndpointZonesArgs, opts ...pulumi.InvokeOption) (*VpcEndpointZonesResult, error) {
	var rv VpcEndpointZonesResult
	err := ctx.Invoke("volcengine:privatelink/vpcEndpointZones:VpcEndpointZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking VpcEndpointZones.
type VpcEndpointZonesArgs struct {
	// The endpoint id of query.
	EndpointId *string `pulumi:"endpointId"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by VpcEndpointZones.
type VpcEndpointZonesResult struct {
	EndpointId *string `pulumi:"endpointId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	OutputFile *string `pulumi:"outputFile"`
	// Returns the total amount of the data list.
	TotalCount int `pulumi:"totalCount"`
	// The collection of query.
	VpcEndpointZones []VpcEndpointZonesVpcEndpointZone `pulumi:"vpcEndpointZones"`
}

func VpcEndpointZonesOutput(ctx *pulumi.Context, args VpcEndpointZonesOutputArgs, opts ...pulumi.InvokeOption) VpcEndpointZonesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (VpcEndpointZonesResult, error) {
			args := v.(VpcEndpointZonesArgs)
			r, err := VpcEndpointZones(ctx, &args, opts...)
			var s VpcEndpointZonesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(VpcEndpointZonesResultOutput)
}

// A collection of arguments for invoking VpcEndpointZones.
type VpcEndpointZonesOutputArgs struct {
	// The endpoint id of query.
	EndpointId pulumi.StringPtrInput `pulumi:"endpointId"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (VpcEndpointZonesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpointZonesArgs)(nil)).Elem()
}

// A collection of values returned by VpcEndpointZones.
type VpcEndpointZonesResultOutput struct{ *pulumi.OutputState }

func (VpcEndpointZonesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpointZonesResult)(nil)).Elem()
}

func (o VpcEndpointZonesResultOutput) ToVpcEndpointZonesResultOutput() VpcEndpointZonesResultOutput {
	return o
}

func (o VpcEndpointZonesResultOutput) ToVpcEndpointZonesResultOutputWithContext(ctx context.Context) VpcEndpointZonesResultOutput {
	return o
}

func (o VpcEndpointZonesResultOutput) EndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcEndpointZonesResult) *string { return v.EndpointId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o VpcEndpointZonesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VpcEndpointZonesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o VpcEndpointZonesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcEndpointZonesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// Returns the total amount of the data list.
func (o VpcEndpointZonesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v VpcEndpointZonesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The collection of query.
func (o VpcEndpointZonesResultOutput) VpcEndpointZones() VpcEndpointZonesVpcEndpointZoneArrayOutput {
	return o.ApplyT(func(v VpcEndpointZonesResult) []VpcEndpointZonesVpcEndpointZone { return v.VpcEndpointZones }).(VpcEndpointZonesVpcEndpointZoneArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VpcEndpointZonesResultOutput{})
}