// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vefaas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of vefaas kafka triggers
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vefaas"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vefaas.GetKafkaTriggers(ctx, &vefaas.GetKafkaTriggersArgs{
//				FunctionId: "f0zvcxxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: volcengine.vefaas.KafkaTriggers has been deprecated in favor of volcengine.vefaas.getKafkaTriggers
func KafkaTriggers(ctx *pulumi.Context, args *KafkaTriggersArgs, opts ...pulumi.InvokeOption) (*KafkaTriggersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv KafkaTriggersResult
	err := ctx.Invoke("volcengine:vefaas/kafkaTriggers:KafkaTriggers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking KafkaTriggers.
type KafkaTriggersArgs struct {
	// The ID of Function.
	FunctionId string `pulumi:"functionId"`
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by KafkaTriggers.
type KafkaTriggersResult struct {
	// The ID of Function.
	FunctionId string `pulumi:"functionId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of timer trigger.
	Items      []KafkaTriggersItem `pulumi:"items"`
	NameRegex  *string             `pulumi:"nameRegex"`
	OutputFile *string             `pulumi:"outputFile"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func KafkaTriggersOutput(ctx *pulumi.Context, args KafkaTriggersOutputArgs, opts ...pulumi.InvokeOption) KafkaTriggersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (KafkaTriggersResult, error) {
			args := v.(KafkaTriggersArgs)
			r, err := KafkaTriggers(ctx, &args, opts...)
			var s KafkaTriggersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(KafkaTriggersResultOutput)
}

// A collection of arguments for invoking KafkaTriggers.
type KafkaTriggersOutputArgs struct {
	// The ID of Function.
	FunctionId pulumi.StringInput `pulumi:"functionId"`
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (KafkaTriggersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KafkaTriggersArgs)(nil)).Elem()
}

// A collection of values returned by KafkaTriggers.
type KafkaTriggersResultOutput struct{ *pulumi.OutputState }

func (KafkaTriggersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KafkaTriggersResult)(nil)).Elem()
}

func (o KafkaTriggersResultOutput) ToKafkaTriggersResultOutput() KafkaTriggersResultOutput {
	return o
}

func (o KafkaTriggersResultOutput) ToKafkaTriggersResultOutputWithContext(ctx context.Context) KafkaTriggersResultOutput {
	return o
}

// The ID of Function.
func (o KafkaTriggersResultOutput) FunctionId() pulumi.StringOutput {
	return o.ApplyT(func(v KafkaTriggersResult) string { return v.FunctionId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o KafkaTriggersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v KafkaTriggersResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of timer trigger.
func (o KafkaTriggersResultOutput) Items() KafkaTriggersItemArrayOutput {
	return o.ApplyT(func(v KafkaTriggersResult) []KafkaTriggersItem { return v.Items }).(KafkaTriggersItemArrayOutput)
}

func (o KafkaTriggersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KafkaTriggersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o KafkaTriggersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KafkaTriggersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o KafkaTriggersResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v KafkaTriggersResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(KafkaTriggersResultOutput{})
}
