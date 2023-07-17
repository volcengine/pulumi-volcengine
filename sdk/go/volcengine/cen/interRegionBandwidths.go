// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of cen inter region bandwidths
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/cen"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cen.InterRegionBandwidths(ctx, &cen.InterRegionBandwidthsArgs{
//				Ids: []string{
//					"cirb-274q484wxao007fap8tlvl6si",
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
func InterRegionBandwidths(ctx *pulumi.Context, args *InterRegionBandwidthsArgs, opts ...pulumi.InvokeOption) (*InterRegionBandwidthsResult, error) {
	var rv InterRegionBandwidthsResult
	err := ctx.Invoke("volcengine:cen/interRegionBandwidths:InterRegionBandwidths", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking InterRegionBandwidths.
type InterRegionBandwidthsArgs struct {
	// A list of cen inter region bandwidth IDs.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by InterRegionBandwidths.
type InterRegionBandwidthsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The collection of cen inter region bandwidth query.
	InterRegionBandwidths []InterRegionBandwidthsInterRegionBandwidth `pulumi:"interRegionBandwidths"`
	OutputFile            *string                                     `pulumi:"outputFile"`
	// The total count of cen inter region bandwidth query.
	TotalCount int `pulumi:"totalCount"`
}

func InterRegionBandwidthsOutput(ctx *pulumi.Context, args InterRegionBandwidthsOutputArgs, opts ...pulumi.InvokeOption) InterRegionBandwidthsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (InterRegionBandwidthsResult, error) {
			args := v.(InterRegionBandwidthsArgs)
			r, err := InterRegionBandwidths(ctx, &args, opts...)
			var s InterRegionBandwidthsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(InterRegionBandwidthsResultOutput)
}

// A collection of arguments for invoking InterRegionBandwidths.
type InterRegionBandwidthsOutputArgs struct {
	// A list of cen inter region bandwidth IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (InterRegionBandwidthsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InterRegionBandwidthsArgs)(nil)).Elem()
}

// A collection of values returned by InterRegionBandwidths.
type InterRegionBandwidthsResultOutput struct{ *pulumi.OutputState }

func (InterRegionBandwidthsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InterRegionBandwidthsResult)(nil)).Elem()
}

func (o InterRegionBandwidthsResultOutput) ToInterRegionBandwidthsResultOutput() InterRegionBandwidthsResultOutput {
	return o
}

func (o InterRegionBandwidthsResultOutput) ToInterRegionBandwidthsResultOutputWithContext(ctx context.Context) InterRegionBandwidthsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o InterRegionBandwidthsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v InterRegionBandwidthsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o InterRegionBandwidthsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InterRegionBandwidthsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The collection of cen inter region bandwidth query.
func (o InterRegionBandwidthsResultOutput) InterRegionBandwidths() InterRegionBandwidthsInterRegionBandwidthArrayOutput {
	return o.ApplyT(func(v InterRegionBandwidthsResult) []InterRegionBandwidthsInterRegionBandwidth {
		return v.InterRegionBandwidths
	}).(InterRegionBandwidthsInterRegionBandwidthArrayOutput)
}

func (o InterRegionBandwidthsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InterRegionBandwidthsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of cen inter region bandwidth query.
func (o InterRegionBandwidthsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v InterRegionBandwidthsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(InterRegionBandwidthsResultOutput{})
}