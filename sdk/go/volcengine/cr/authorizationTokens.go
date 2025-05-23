// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of cr authorization tokens
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/cr"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cr.GetAuthorizationTokens(ctx, &cr.GetAuthorizationTokensArgs{
//				Registry: "tf-1",
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
// Deprecated: volcengine.cr.AuthorizationTokens has been deprecated in favor of volcengine.cr.getAuthorizationTokens
func AuthorizationTokens(ctx *pulumi.Context, args *AuthorizationTokensArgs, opts ...pulumi.InvokeOption) (*AuthorizationTokensResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv AuthorizationTokensResult
	err := ctx.Invoke("volcengine:cr/authorizationTokens:AuthorizationTokens", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking AuthorizationTokens.
type AuthorizationTokensArgs struct {
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The cr instance name want to query.
	Registry string `pulumi:"registry"`
}

// A collection of values returned by AuthorizationTokens.
type AuthorizationTokensResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	OutputFile *string `pulumi:"outputFile"`
	Registry   string  `pulumi:"registry"`
	// The collection of users.
	Tokens []AuthorizationTokensToken `pulumi:"tokens"`
	// The total count of instance query.
	TotalCount int `pulumi:"totalCount"`
}

func AuthorizationTokensOutput(ctx *pulumi.Context, args AuthorizationTokensOutputArgs, opts ...pulumi.InvokeOption) AuthorizationTokensResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (AuthorizationTokensResult, error) {
			args := v.(AuthorizationTokensArgs)
			r, err := AuthorizationTokens(ctx, &args, opts...)
			var s AuthorizationTokensResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(AuthorizationTokensResultOutput)
}

// A collection of arguments for invoking AuthorizationTokens.
type AuthorizationTokensOutputArgs struct {
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The cr instance name want to query.
	Registry pulumi.StringInput `pulumi:"registry"`
}

func (AuthorizationTokensOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationTokensArgs)(nil)).Elem()
}

// A collection of values returned by AuthorizationTokens.
type AuthorizationTokensResultOutput struct{ *pulumi.OutputState }

func (AuthorizationTokensResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationTokensResult)(nil)).Elem()
}

func (o AuthorizationTokensResultOutput) ToAuthorizationTokensResultOutput() AuthorizationTokensResultOutput {
	return o
}

func (o AuthorizationTokensResultOutput) ToAuthorizationTokensResultOutputWithContext(ctx context.Context) AuthorizationTokensResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o AuthorizationTokensResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AuthorizationTokensResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o AuthorizationTokensResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthorizationTokensResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o AuthorizationTokensResultOutput) Registry() pulumi.StringOutput {
	return o.ApplyT(func(v AuthorizationTokensResult) string { return v.Registry }).(pulumi.StringOutput)
}

// The collection of users.
func (o AuthorizationTokensResultOutput) Tokens() AuthorizationTokensTokenArrayOutput {
	return o.ApplyT(func(v AuthorizationTokensResult) []AuthorizationTokensToken { return v.Tokens }).(AuthorizationTokensTokenArrayOutput)
}

// The total count of instance query.
func (o AuthorizationTokensResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v AuthorizationTokensResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorizationTokensResultOutput{})
}
