// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of cloudfs namespaces
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/cloudfs"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfs.GetNamespaces(ctx, &cloudfs.GetNamespacesArgs{
//				FsName: "tf-test-fs",
//				NsId:   pulumi.StringRef("1801439850948****"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetNamespaces(ctx *pulumi.Context, args *GetNamespacesArgs, opts ...pulumi.InvokeOption) (*GetNamespacesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNamespacesResult
	err := ctx.Invoke("volcengine:cloudfs/getNamespaces:getNamespaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNamespaces.
type GetNamespacesArgs struct {
	// The name of file system.
	FsName string `pulumi:"fsName"`
	// The id of namespace.
	NsId *string `pulumi:"nsId"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The name of tos bucket.
	TosBucket *string `pulumi:"tosBucket"`
}

// A collection of values returned by getNamespaces.
type GetNamespacesResult struct {
	FsName string `pulumi:"fsName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The collection of query.
	Namespaces []GetNamespacesNamespace `pulumi:"namespaces"`
	NsId       *string                  `pulumi:"nsId"`
	OutputFile *string                  `pulumi:"outputFile"`
	// The name of the tos bucket.
	TosBucket *string `pulumi:"tosBucket"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func GetNamespacesOutput(ctx *pulumi.Context, args GetNamespacesOutputArgs, opts ...pulumi.InvokeOption) GetNamespacesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNamespacesResult, error) {
			args := v.(GetNamespacesArgs)
			r, err := GetNamespaces(ctx, &args, opts...)
			var s GetNamespacesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNamespacesResultOutput)
}

// A collection of arguments for invoking getNamespaces.
type GetNamespacesOutputArgs struct {
	// The name of file system.
	FsName pulumi.StringInput `pulumi:"fsName"`
	// The id of namespace.
	NsId pulumi.StringPtrInput `pulumi:"nsId"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The name of tos bucket.
	TosBucket pulumi.StringPtrInput `pulumi:"tosBucket"`
}

func (GetNamespacesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNamespacesArgs)(nil)).Elem()
}

// A collection of values returned by getNamespaces.
type GetNamespacesResultOutput struct{ *pulumi.OutputState }

func (GetNamespacesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNamespacesResult)(nil)).Elem()
}

func (o GetNamespacesResultOutput) ToGetNamespacesResultOutput() GetNamespacesResultOutput {
	return o
}

func (o GetNamespacesResultOutput) ToGetNamespacesResultOutputWithContext(ctx context.Context) GetNamespacesResultOutput {
	return o
}

func (o GetNamespacesResultOutput) FsName() pulumi.StringOutput {
	return o.ApplyT(func(v GetNamespacesResult) string { return v.FsName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNamespacesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNamespacesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The collection of query.
func (o GetNamespacesResultOutput) Namespaces() GetNamespacesNamespaceArrayOutput {
	return o.ApplyT(func(v GetNamespacesResult) []GetNamespacesNamespace { return v.Namespaces }).(GetNamespacesNamespaceArrayOutput)
}

func (o GetNamespacesResultOutput) NsId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNamespacesResult) *string { return v.NsId }).(pulumi.StringPtrOutput)
}

func (o GetNamespacesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNamespacesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The name of the tos bucket.
func (o GetNamespacesResultOutput) TosBucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNamespacesResult) *string { return v.TosBucket }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o GetNamespacesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetNamespacesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNamespacesResultOutput{})
}
