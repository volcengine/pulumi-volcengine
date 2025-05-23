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
func GetAccessKeys(ctx *pulumi.Context, args *GetAccessKeysArgs, opts ...pulumi.InvokeOption) (*GetAccessKeysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAccessKeysResult
	err := ctx.Invoke("volcengine:rocketmq/getAccessKeys:getAccessKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccessKeys.
type GetAccessKeysArgs struct {
	// The access key id of the rocketmq key.
	AccessKey *string `pulumi:"accessKey"`
	// The id of rocketmq instance.
	InstanceId string `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getAccessKeys.
type GetAccessKeysResult struct {
	// The access key id of the rocketmq key.
	AccessKey *string `pulumi:"accessKey"`
	// The collection of query.
	AccessKeys []GetAccessKeysAccessKey `pulumi:"accessKeys"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The id of rocketmq instance.
	InstanceId string  `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func GetAccessKeysOutput(ctx *pulumi.Context, args GetAccessKeysOutputArgs, opts ...pulumi.InvokeOption) GetAccessKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAccessKeysResult, error) {
			args := v.(GetAccessKeysArgs)
			r, err := GetAccessKeys(ctx, &args, opts...)
			var s GetAccessKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAccessKeysResultOutput)
}

// A collection of arguments for invoking getAccessKeys.
type GetAccessKeysOutputArgs struct {
	// The access key id of the rocketmq key.
	AccessKey pulumi.StringPtrInput `pulumi:"accessKey"`
	// The id of rocketmq instance.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetAccessKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccessKeysArgs)(nil)).Elem()
}

// A collection of values returned by getAccessKeys.
type GetAccessKeysResultOutput struct{ *pulumi.OutputState }

func (GetAccessKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccessKeysResult)(nil)).Elem()
}

func (o GetAccessKeysResultOutput) ToGetAccessKeysResultOutput() GetAccessKeysResultOutput {
	return o
}

func (o GetAccessKeysResultOutput) ToGetAccessKeysResultOutputWithContext(ctx context.Context) GetAccessKeysResultOutput {
	return o
}

// The access key id of the rocketmq key.
func (o GetAccessKeysResultOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessKeysResult) *string { return v.AccessKey }).(pulumi.StringPtrOutput)
}

// The collection of query.
func (o GetAccessKeysResultOutput) AccessKeys() GetAccessKeysAccessKeyArrayOutput {
	return o.ApplyT(func(v GetAccessKeysResult) []GetAccessKeysAccessKey { return v.AccessKeys }).(GetAccessKeysAccessKeyArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccessKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

// The id of rocketmq instance.
func (o GetAccessKeysResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessKeysResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetAccessKeysResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessKeysResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o GetAccessKeysResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetAccessKeysResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccessKeysResultOutput{})
}
