// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of cr vpc endpoints
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/cr"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cr.VpcEndpoints(ctx, &cr.VpcEndpointsArgs{
//				Registry: "enterprise-1",
//				Statuses: []string{
//					"Enabled",
//					"Enabling",
//					"Disabling",
//					"Failed",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func VpcEndpoints(ctx *pulumi.Context, args *VpcEndpointsArgs, opts ...pulumi.InvokeOption) (*VpcEndpointsResult, error) {
	var rv VpcEndpointsResult
	err := ctx.Invoke("volcengine:cr/vpcEndpoints:VpcEndpoints", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking VpcEndpoints.
type VpcEndpointsArgs struct {
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The CR registry name.
	Registry string `pulumi:"registry"`
	// VPC access entry state array, used to filter out VPC access entries in the specified state. Available values are Enabling, Enabled, Disabling, Failed.
	Statuses []string `pulumi:"statuses"`
}

// A collection of values returned by VpcEndpoints.
type VpcEndpointsResult struct {
	// List of CR vpc endpoints.
	Endpoints []VpcEndpointsEndpoint `pulumi:"endpoints"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	OutputFile *string `pulumi:"outputFile"`
	// The name of CR registry.
	Registry string   `pulumi:"registry"`
	Statuses []string `pulumi:"statuses"`
	// The total count of CR vpc endpoints query.
	TotalCount int `pulumi:"totalCount"`
}

func VpcEndpointsOutput(ctx *pulumi.Context, args VpcEndpointsOutputArgs, opts ...pulumi.InvokeOption) VpcEndpointsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (VpcEndpointsResult, error) {
			args := v.(VpcEndpointsArgs)
			r, err := VpcEndpoints(ctx, &args, opts...)
			var s VpcEndpointsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(VpcEndpointsResultOutput)
}

// A collection of arguments for invoking VpcEndpoints.
type VpcEndpointsOutputArgs struct {
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The CR registry name.
	Registry pulumi.StringInput `pulumi:"registry"`
	// VPC access entry state array, used to filter out VPC access entries in the specified state. Available values are Enabling, Enabled, Disabling, Failed.
	Statuses pulumi.StringArrayInput `pulumi:"statuses"`
}

func (VpcEndpointsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpointsArgs)(nil)).Elem()
}

// A collection of values returned by VpcEndpoints.
type VpcEndpointsResultOutput struct{ *pulumi.OutputState }

func (VpcEndpointsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpointsResult)(nil)).Elem()
}

func (o VpcEndpointsResultOutput) ToVpcEndpointsResultOutput() VpcEndpointsResultOutput {
	return o
}

func (o VpcEndpointsResultOutput) ToVpcEndpointsResultOutputWithContext(ctx context.Context) VpcEndpointsResultOutput {
	return o
}

// List of CR vpc endpoints.
func (o VpcEndpointsResultOutput) Endpoints() VpcEndpointsEndpointArrayOutput {
	return o.ApplyT(func(v VpcEndpointsResult) []VpcEndpointsEndpoint { return v.Endpoints }).(VpcEndpointsEndpointArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o VpcEndpointsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VpcEndpointsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o VpcEndpointsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcEndpointsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The name of CR registry.
func (o VpcEndpointsResultOutput) Registry() pulumi.StringOutput {
	return o.ApplyT(func(v VpcEndpointsResult) string { return v.Registry }).(pulumi.StringOutput)
}

func (o VpcEndpointsResultOutput) Statuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VpcEndpointsResult) []string { return v.Statuses }).(pulumi.StringArrayOutput)
}

// The total count of CR vpc endpoints query.
func (o VpcEndpointsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v VpcEndpointsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(VpcEndpointsResultOutput{})
}