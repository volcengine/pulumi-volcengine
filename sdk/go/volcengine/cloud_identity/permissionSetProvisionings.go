// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud_identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of cloud identity permission set provisionings
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/cloud_identity"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloud_identity.PermissionSetProvisionings(ctx, &cloud_identity.PermissionSetProvisioningsArgs{
//				TargetId: pulumi.StringRef("210026****"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func PermissionSetProvisionings(ctx *pulumi.Context, args *PermissionSetProvisioningsArgs, opts ...pulumi.InvokeOption) (*PermissionSetProvisioningsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv PermissionSetProvisioningsResult
	err := ctx.Invoke("volcengine:cloud_identity/permissionSetProvisionings:PermissionSetProvisionings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking PermissionSetProvisionings.
type PermissionSetProvisioningsArgs struct {
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The id of cloud identity permission set.
	PermissionSetId *string `pulumi:"permissionSetId"`
	// The target account id of cloud identity permission set.
	TargetId *string `pulumi:"targetId"`
}

// A collection of values returned by PermissionSetProvisionings.
type PermissionSetProvisioningsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The collection of query.
	PermissionProvisionings []PermissionSetProvisioningsPermissionProvisioning `pulumi:"permissionProvisionings"`
	// The id of the cloud identity permission set.
	PermissionSetId *string `pulumi:"permissionSetId"`
	// The target account id of the cloud identity permission set provisioning.
	TargetId *string `pulumi:"targetId"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func PermissionSetProvisioningsOutput(ctx *pulumi.Context, args PermissionSetProvisioningsOutputArgs, opts ...pulumi.InvokeOption) PermissionSetProvisioningsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (PermissionSetProvisioningsResult, error) {
			args := v.(PermissionSetProvisioningsArgs)
			r, err := PermissionSetProvisionings(ctx, &args, opts...)
			var s PermissionSetProvisioningsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(PermissionSetProvisioningsResultOutput)
}

// A collection of arguments for invoking PermissionSetProvisionings.
type PermissionSetProvisioningsOutputArgs struct {
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The id of cloud identity permission set.
	PermissionSetId pulumi.StringPtrInput `pulumi:"permissionSetId"`
	// The target account id of cloud identity permission set.
	TargetId pulumi.StringPtrInput `pulumi:"targetId"`
}

func (PermissionSetProvisioningsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionSetProvisioningsArgs)(nil)).Elem()
}

// A collection of values returned by PermissionSetProvisionings.
type PermissionSetProvisioningsResultOutput struct{ *pulumi.OutputState }

func (PermissionSetProvisioningsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionSetProvisioningsResult)(nil)).Elem()
}

func (o PermissionSetProvisioningsResultOutput) ToPermissionSetProvisioningsResultOutput() PermissionSetProvisioningsResultOutput {
	return o
}

func (o PermissionSetProvisioningsResultOutput) ToPermissionSetProvisioningsResultOutputWithContext(ctx context.Context) PermissionSetProvisioningsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o PermissionSetProvisioningsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PermissionSetProvisioningsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o PermissionSetProvisioningsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PermissionSetProvisioningsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o PermissionSetProvisioningsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PermissionSetProvisioningsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The collection of query.
func (o PermissionSetProvisioningsResultOutput) PermissionProvisionings() PermissionSetProvisioningsPermissionProvisioningArrayOutput {
	return o.ApplyT(func(v PermissionSetProvisioningsResult) []PermissionSetProvisioningsPermissionProvisioning {
		return v.PermissionProvisionings
	}).(PermissionSetProvisioningsPermissionProvisioningArrayOutput)
}

// The id of the cloud identity permission set.
func (o PermissionSetProvisioningsResultOutput) PermissionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PermissionSetProvisioningsResult) *string { return v.PermissionSetId }).(pulumi.StringPtrOutput)
}

// The target account id of the cloud identity permission set provisioning.
func (o PermissionSetProvisioningsResultOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PermissionSetProvisioningsResult) *string { return v.TargetId }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o PermissionSetProvisioningsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v PermissionSetProvisioningsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(PermissionSetProvisioningsResultOutput{})
}