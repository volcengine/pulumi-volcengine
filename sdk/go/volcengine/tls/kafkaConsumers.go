// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of tls kafka consumers
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/tls"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tls.KafkaConsumers(ctx, &tls.KafkaConsumersArgs{
//				Ids: []string{
//					"65d67d34-c5b4-4ec8-b3a9-175d33668b45",
//					"cfb5c08b-0c7a-44fa-8971-8afc12f1b123",
//					"edf051ed-3c46-49ba-9339-bea628fedc15",
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
func KafkaConsumers(ctx *pulumi.Context, args *KafkaConsumersArgs, opts ...pulumi.InvokeOption) (*KafkaConsumersResult, error) {
	var rv KafkaConsumersResult
	err := ctx.Invoke("volcengine:tls/kafkaConsumers:KafkaConsumers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking KafkaConsumers.
type KafkaConsumersArgs struct {
	// A list of topic IDs.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by KafkaConsumers.
type KafkaConsumersResult struct {
	// The collection of query.
	Datas []KafkaConsumersData `pulumi:"datas"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func KafkaConsumersOutput(ctx *pulumi.Context, args KafkaConsumersOutputArgs, opts ...pulumi.InvokeOption) KafkaConsumersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (KafkaConsumersResult, error) {
			args := v.(KafkaConsumersArgs)
			r, err := KafkaConsumers(ctx, &args, opts...)
			var s KafkaConsumersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(KafkaConsumersResultOutput)
}

// A collection of arguments for invoking KafkaConsumers.
type KafkaConsumersOutputArgs struct {
	// A list of topic IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (KafkaConsumersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KafkaConsumersArgs)(nil)).Elem()
}

// A collection of values returned by KafkaConsumers.
type KafkaConsumersResultOutput struct{ *pulumi.OutputState }

func (KafkaConsumersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KafkaConsumersResult)(nil)).Elem()
}

func (o KafkaConsumersResultOutput) ToKafkaConsumersResultOutput() KafkaConsumersResultOutput {
	return o
}

func (o KafkaConsumersResultOutput) ToKafkaConsumersResultOutputWithContext(ctx context.Context) KafkaConsumersResultOutput {
	return o
}

// The collection of query.
func (o KafkaConsumersResultOutput) Datas() KafkaConsumersDataArrayOutput {
	return o.ApplyT(func(v KafkaConsumersResult) []KafkaConsumersData { return v.Datas }).(KafkaConsumersDataArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o KafkaConsumersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v KafkaConsumersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o KafkaConsumersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v KafkaConsumersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o KafkaConsumersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KafkaConsumersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o KafkaConsumersResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v KafkaConsumersResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(KafkaConsumersResultOutput{})
}