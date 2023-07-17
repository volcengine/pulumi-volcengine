// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tos

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of tos objects
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/tos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tos.BucketObjects(ctx, &tos.BucketObjectsArgs{
//				BucketName: "test",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func BucketObjects(ctx *pulumi.Context, args *BucketObjectsArgs, opts ...pulumi.InvokeOption) (*BucketObjectsResult, error) {
	var rv BucketObjectsResult
	err := ctx.Invoke("volcengine:tos/bucketObjects:BucketObjects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking BucketObjects.
type BucketObjectsArgs struct {
	// The name the TOS bucket.
	BucketName string `pulumi:"bucketName"`
	// A Name Regex of TOS Object.
	NameRegex *string `pulumi:"nameRegex"`
	// The name the TOS Object.
	ObjectName *string `pulumi:"objectName"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by BucketObjects.
type BucketObjectsResult struct {
	BucketName string `pulumi:"bucketName"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	NameRegex  *string `pulumi:"nameRegex"`
	ObjectName *string `pulumi:"objectName"`
	// The collection of TOS Object query.
	Objects    []BucketObjectsObject `pulumi:"objects"`
	OutputFile *string               `pulumi:"outputFile"`
	// The total count of TOS Object query.
	TotalCount int `pulumi:"totalCount"`
}

func BucketObjectsOutput(ctx *pulumi.Context, args BucketObjectsOutputArgs, opts ...pulumi.InvokeOption) BucketObjectsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (BucketObjectsResult, error) {
			args := v.(BucketObjectsArgs)
			r, err := BucketObjects(ctx, &args, opts...)
			var s BucketObjectsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(BucketObjectsResultOutput)
}

// A collection of arguments for invoking BucketObjects.
type BucketObjectsOutputArgs struct {
	// The name the TOS bucket.
	BucketName pulumi.StringInput `pulumi:"bucketName"`
	// A Name Regex of TOS Object.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// The name the TOS Object.
	ObjectName pulumi.StringPtrInput `pulumi:"objectName"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (BucketObjectsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketObjectsArgs)(nil)).Elem()
}

// A collection of values returned by BucketObjects.
type BucketObjectsResultOutput struct{ *pulumi.OutputState }

func (BucketObjectsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketObjectsResult)(nil)).Elem()
}

func (o BucketObjectsResultOutput) ToBucketObjectsResultOutput() BucketObjectsResultOutput {
	return o
}

func (o BucketObjectsResultOutput) ToBucketObjectsResultOutputWithContext(ctx context.Context) BucketObjectsResultOutput {
	return o
}

func (o BucketObjectsResultOutput) BucketName() pulumi.StringOutput {
	return o.ApplyT(func(v BucketObjectsResult) string { return v.BucketName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o BucketObjectsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v BucketObjectsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o BucketObjectsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BucketObjectsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o BucketObjectsResultOutput) ObjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BucketObjectsResult) *string { return v.ObjectName }).(pulumi.StringPtrOutput)
}

// The collection of TOS Object query.
func (o BucketObjectsResultOutput) Objects() BucketObjectsObjectArrayOutput {
	return o.ApplyT(func(v BucketObjectsResult) []BucketObjectsObject { return v.Objects }).(BucketObjectsObjectArrayOutput)
}

func (o BucketObjectsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BucketObjectsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of TOS Object query.
func (o BucketObjectsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v BucketObjectsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(BucketObjectsResultOutput{})
}