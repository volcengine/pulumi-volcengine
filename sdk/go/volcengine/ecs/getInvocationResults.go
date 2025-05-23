// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of ecs invocation results
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
//			_, err := ecs.GetInvocationResults(ctx, &ecs.GetInvocationResultsArgs{
//				InvocationId: "ivk-ych9y4vujvl8j01c****",
//				InvocationResultStatuses: []string{
//					"Success",
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
func GetInvocationResults(ctx *pulumi.Context, args *GetInvocationResultsArgs, opts ...pulumi.InvokeOption) (*GetInvocationResultsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInvocationResultsResult
	err := ctx.Invoke("volcengine:ecs/getInvocationResults:getInvocationResults", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInvocationResults.
type GetInvocationResultsArgs struct {
	// The id of ecs command.
	CommandId *string `pulumi:"commandId"`
	// The id of ecs instance.
	InstanceId *string `pulumi:"instanceId"`
	// The id of ecs invocation.
	InvocationId string `pulumi:"invocationId"`
	// The list of status of ecs invocation in a single instance. Valid values: `Pending`, `Running`, `Success`, `Failed`, `Timeout`.
	InvocationResultStatuses []string `pulumi:"invocationResultStatuses"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getInvocationResults.
type GetInvocationResultsResult struct {
	// The id of the ecs command.
	CommandId *string `pulumi:"commandId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The id of the ecs instance.
	InstanceId *string `pulumi:"instanceId"`
	// The id of the ecs invocation.
	InvocationId string `pulumi:"invocationId"`
	// The status of ecs invocation in a single instance.
	InvocationResultStatuses []string `pulumi:"invocationResultStatuses"`
	// The collection of query.
	InvocationResults []GetInvocationResultsInvocationResult `pulumi:"invocationResults"`
	OutputFile        *string                                `pulumi:"outputFile"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func GetInvocationResultsOutput(ctx *pulumi.Context, args GetInvocationResultsOutputArgs, opts ...pulumi.InvokeOption) GetInvocationResultsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInvocationResultsResult, error) {
			args := v.(GetInvocationResultsArgs)
			r, err := GetInvocationResults(ctx, &args, opts...)
			var s GetInvocationResultsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInvocationResultsResultOutput)
}

// A collection of arguments for invoking getInvocationResults.
type GetInvocationResultsOutputArgs struct {
	// The id of ecs command.
	CommandId pulumi.StringPtrInput `pulumi:"commandId"`
	// The id of ecs instance.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// The id of ecs invocation.
	InvocationId pulumi.StringInput `pulumi:"invocationId"`
	// The list of status of ecs invocation in a single instance. Valid values: `Pending`, `Running`, `Success`, `Failed`, `Timeout`.
	InvocationResultStatuses pulumi.StringArrayInput `pulumi:"invocationResultStatuses"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetInvocationResultsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInvocationResultsArgs)(nil)).Elem()
}

// A collection of values returned by getInvocationResults.
type GetInvocationResultsResultOutput struct{ *pulumi.OutputState }

func (GetInvocationResultsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInvocationResultsResult)(nil)).Elem()
}

func (o GetInvocationResultsResultOutput) ToGetInvocationResultsResultOutput() GetInvocationResultsResultOutput {
	return o
}

func (o GetInvocationResultsResultOutput) ToGetInvocationResultsResultOutputWithContext(ctx context.Context) GetInvocationResultsResultOutput {
	return o
}

// The id of the ecs command.
func (o GetInvocationResultsResultOutput) CommandId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInvocationResultsResult) *string { return v.CommandId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInvocationResultsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInvocationResultsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The id of the ecs instance.
func (o GetInvocationResultsResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInvocationResultsResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

// The id of the ecs invocation.
func (o GetInvocationResultsResultOutput) InvocationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInvocationResultsResult) string { return v.InvocationId }).(pulumi.StringOutput)
}

// The status of ecs invocation in a single instance.
func (o GetInvocationResultsResultOutput) InvocationResultStatuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInvocationResultsResult) []string { return v.InvocationResultStatuses }).(pulumi.StringArrayOutput)
}

// The collection of query.
func (o GetInvocationResultsResultOutput) InvocationResults() GetInvocationResultsInvocationResultArrayOutput {
	return o.ApplyT(func(v GetInvocationResultsResult) []GetInvocationResultsInvocationResult { return v.InvocationResults }).(GetInvocationResultsInvocationResultArrayOutput)
}

func (o GetInvocationResultsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInvocationResultsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o GetInvocationResultsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetInvocationResultsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInvocationResultsResultOutput{})
}
