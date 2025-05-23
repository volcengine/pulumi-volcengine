// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rocketmq

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of rocketmq access keys
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/rocketmq"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rocketmq.GetAccessKeys(ctx, &rocketmq.GetAccessKeysArgs{
//				InstanceId: "rocketmq-cnoeea6b32118fc2",
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
// Deprecated: volcengine.rocketmq.AccessKeys has been deprecated in favor of volcengine.rocketmq.getAccessKeys
func AccessKeys(ctx *pulumi.Context, args *AccessKeysArgs, opts ...pulumi.InvokeOption) (*AccessKeysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv AccessKeysResult
	err := ctx.Invoke("volcengine:rocketmq/accessKeys:AccessKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking AccessKeys.
type AccessKeysArgs struct {
	// The access key id of the rocketmq key.
	AccessKey *string `pulumi:"accessKey"`
	// The id of rocketmq instance.
	InstanceId string `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by AccessKeys.
type AccessKeysResult struct {
	// The access key id of the rocketmq key.
	AccessKey *string `pulumi:"accessKey"`
	// The collection of query.
	AccessKeys []AccessKeysAccessKey `pulumi:"accessKeys"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The id of rocketmq instance.
	InstanceId string  `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func AccessKeysOutput(ctx *pulumi.Context, args AccessKeysOutputArgs, opts ...pulumi.InvokeOption) AccessKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (AccessKeysResult, error) {
			args := v.(AccessKeysArgs)
			r, err := AccessKeys(ctx, &args, opts...)
			var s AccessKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(AccessKeysResultOutput)
}

// A collection of arguments for invoking AccessKeys.
type AccessKeysOutputArgs struct {
	// The access key id of the rocketmq key.
	AccessKey pulumi.StringPtrInput `pulumi:"accessKey"`
	// The id of rocketmq instance.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (AccessKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessKeysArgs)(nil)).Elem()
}

// A collection of values returned by AccessKeys.
type AccessKeysResultOutput struct{ *pulumi.OutputState }

func (AccessKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessKeysResult)(nil)).Elem()
}

func (o AccessKeysResultOutput) ToAccessKeysResultOutput() AccessKeysResultOutput {
	return o
}

func (o AccessKeysResultOutput) ToAccessKeysResultOutputWithContext(ctx context.Context) AccessKeysResultOutput {
	return o
}

// The access key id of the rocketmq key.
func (o AccessKeysResultOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessKeysResult) *string { return v.AccessKey }).(pulumi.StringPtrOutput)
}

// The collection of query.
func (o AccessKeysResultOutput) AccessKeys() AccessKeysAccessKeyArrayOutput {
	return o.ApplyT(func(v AccessKeysResult) []AccessKeysAccessKey { return v.AccessKeys }).(AccessKeysAccessKeyArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o AccessKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AccessKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

// The id of rocketmq instance.
func (o AccessKeysResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v AccessKeysResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o AccessKeysResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessKeysResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o AccessKeysResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v AccessKeysResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessKeysResultOutput{})
}
