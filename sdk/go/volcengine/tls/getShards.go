// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of tls shards
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/tls"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tls.GetShards(ctx, &tls.GetShardsArgs{
//				TopicId: "edf051ed-3c46-49ba-9339-bea628fedc15",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetShards(ctx *pulumi.Context, args *GetShardsArgs, opts ...pulumi.InvokeOption) (*GetShardsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetShardsResult
	err := ctx.Invoke("volcengine:tls/getShards:getShards", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getShards.
type GetShardsArgs struct {
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The id of topic.
	TopicId string `pulumi:"topicId"`
}

// A collection of values returned by getShards.
type GetShardsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	OutputFile *string `pulumi:"outputFile"`
	// The collection of query.
	Shards []GetShardsShard `pulumi:"shards"`
	// The ID of topic.
	TopicId string `pulumi:"topicId"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func GetShardsOutput(ctx *pulumi.Context, args GetShardsOutputArgs, opts ...pulumi.InvokeOption) GetShardsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetShardsResult, error) {
			args := v.(GetShardsArgs)
			r, err := GetShards(ctx, &args, opts...)
			var s GetShardsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetShardsResultOutput)
}

// A collection of arguments for invoking getShards.
type GetShardsOutputArgs struct {
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The id of topic.
	TopicId pulumi.StringInput `pulumi:"topicId"`
}

func (GetShardsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetShardsArgs)(nil)).Elem()
}

// A collection of values returned by getShards.
type GetShardsResultOutput struct{ *pulumi.OutputState }

func (GetShardsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetShardsResult)(nil)).Elem()
}

func (o GetShardsResultOutput) ToGetShardsResultOutput() GetShardsResultOutput {
	return o
}

func (o GetShardsResultOutput) ToGetShardsResultOutputWithContext(ctx context.Context) GetShardsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetShardsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetShardsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetShardsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetShardsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The collection of query.
func (o GetShardsResultOutput) Shards() GetShardsShardArrayOutput {
	return o.ApplyT(func(v GetShardsResult) []GetShardsShard { return v.Shards }).(GetShardsShardArrayOutput)
}

// The ID of topic.
func (o GetShardsResultOutput) TopicId() pulumi.StringOutput {
	return o.ApplyT(func(v GetShardsResult) string { return v.TopicId }).(pulumi.StringOutput)
}

// The total count of query.
func (o GetShardsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetShardsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetShardsResultOutput{})
}
