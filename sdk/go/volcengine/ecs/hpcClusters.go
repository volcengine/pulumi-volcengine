// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of ecs hpc clusters
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ecs.HpcClusters(ctx, &ecs.HpcClustersArgs{
//				ZoneId: pulumi.StringRef("cn-beijing-a"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func HpcClusters(ctx *pulumi.Context, args *HpcClustersArgs, opts ...pulumi.InvokeOption) (*HpcClustersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv HpcClustersResult
	err := ctx.Invoke("volcengine:ecs/hpcClusters:HpcClusters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking HpcClusters.
type HpcClustersArgs struct {
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The zone id of the hpc cluster.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by HpcClusters.
type HpcClustersResult struct {
	// The collection of query.
	HpcClusters []HpcClustersHpcCluster `pulumi:"hpcClusters"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
	// The zone id of the hpc cluster.
	ZoneId *string `pulumi:"zoneId"`
}

func HpcClustersOutput(ctx *pulumi.Context, args HpcClustersOutputArgs, opts ...pulumi.InvokeOption) HpcClustersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (HpcClustersResult, error) {
			args := v.(HpcClustersArgs)
			r, err := HpcClusters(ctx, &args, opts...)
			var s HpcClustersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(HpcClustersResultOutput)
}

// A collection of arguments for invoking HpcClusters.
type HpcClustersOutputArgs struct {
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The zone id of the hpc cluster.
	ZoneId pulumi.StringPtrInput `pulumi:"zoneId"`
}

func (HpcClustersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HpcClustersArgs)(nil)).Elem()
}

// A collection of values returned by HpcClusters.
type HpcClustersResultOutput struct{ *pulumi.OutputState }

func (HpcClustersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HpcClustersResult)(nil)).Elem()
}

func (o HpcClustersResultOutput) ToHpcClustersResultOutput() HpcClustersResultOutput {
	return o
}

func (o HpcClustersResultOutput) ToHpcClustersResultOutputWithContext(ctx context.Context) HpcClustersResultOutput {
	return o
}

// The collection of query.
func (o HpcClustersResultOutput) HpcClusters() HpcClustersHpcClusterArrayOutput {
	return o.ApplyT(func(v HpcClustersResult) []HpcClustersHpcCluster { return v.HpcClusters }).(HpcClustersHpcClusterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o HpcClustersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v HpcClustersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o HpcClustersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HpcClustersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o HpcClustersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HpcClustersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o HpcClustersResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v HpcClustersResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The zone id of the hpc cluster.
func (o HpcClustersResultOutput) ZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HpcClustersResult) *string { return v.ZoneId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(HpcClustersResultOutput{})
}