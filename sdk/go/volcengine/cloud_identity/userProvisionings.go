// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud_identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of cloud identity user provisionings
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
//			_, err := cloud_identity.UserProvisionings(ctx, &cloud_identity.UserProvisioningsArgs{
//				AccountId: pulumi.StringRef("210026****"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func UserProvisionings(ctx *pulumi.Context, args *UserProvisioningsArgs, opts ...pulumi.InvokeOption) (*UserProvisioningsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv UserProvisioningsResult
	err := ctx.Invoke("volcengine:cloud_identity/userProvisionings:UserProvisionings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking UserProvisionings.
type UserProvisioningsArgs struct {
	// The account id.
	AccountId *string `pulumi:"accountId"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by UserProvisionings.
type UserProvisioningsResult struct {
	AccountId *string `pulumi:"accountId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	OutputFile *string `pulumi:"outputFile"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
	// The collection of query.
	UserProvisionings []UserProvisioningsUserProvisioning `pulumi:"userProvisionings"`
}

func UserProvisioningsOutput(ctx *pulumi.Context, args UserProvisioningsOutputArgs, opts ...pulumi.InvokeOption) UserProvisioningsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (UserProvisioningsResult, error) {
			args := v.(UserProvisioningsArgs)
			r, err := UserProvisionings(ctx, &args, opts...)
			var s UserProvisioningsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(UserProvisioningsResultOutput)
}

// A collection of arguments for invoking UserProvisionings.
type UserProvisioningsOutputArgs struct {
	// The account id.
	AccountId pulumi.StringPtrInput `pulumi:"accountId"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (UserProvisioningsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserProvisioningsArgs)(nil)).Elem()
}

// A collection of values returned by UserProvisionings.
type UserProvisioningsResultOutput struct{ *pulumi.OutputState }

func (UserProvisioningsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserProvisioningsResult)(nil)).Elem()
}

func (o UserProvisioningsResultOutput) ToUserProvisioningsResultOutput() UserProvisioningsResultOutput {
	return o
}

func (o UserProvisioningsResultOutput) ToUserProvisioningsResultOutputWithContext(ctx context.Context) UserProvisioningsResultOutput {
	return o
}

func (o UserProvisioningsResultOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserProvisioningsResult) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o UserProvisioningsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v UserProvisioningsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o UserProvisioningsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserProvisioningsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o UserProvisioningsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v UserProvisioningsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The collection of query.
func (o UserProvisioningsResultOutput) UserProvisionings() UserProvisioningsUserProvisioningArrayOutput {
	return o.ApplyT(func(v UserProvisioningsResult) []UserProvisioningsUserProvisioning { return v.UserProvisionings }).(UserProvisioningsUserProvisioningArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(UserProvisioningsResultOutput{})
}