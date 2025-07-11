// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of waf custom bots
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/waf"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := waf.GetCustomBots(ctx, &waf.GetCustomBotsArgs{
//				Host: pulumi.StringRef("www.tf-test.com"),
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
// Deprecated: volcengine.waf.CustomBots has been deprecated in favor of volcengine.waf.getCustomBots
func CustomBots(ctx *pulumi.Context, args *CustomBotsArgs, opts ...pulumi.InvokeOption) (*CustomBotsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv CustomBotsResult
	err := ctx.Invoke("volcengine:waf/customBots:CustomBots", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking CustomBots.
type CustomBotsArgs struct {
	// The domain names that need to be viewed.
	Host *string `pulumi:"host"`
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by CustomBots.
type CustomBotsResult struct {
	// The Details of Custom bot.
	Datas []CustomBotsData `pulumi:"datas"`
	Host  *string          `pulumi:"host"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func CustomBotsOutput(ctx *pulumi.Context, args CustomBotsOutputArgs, opts ...pulumi.InvokeOption) CustomBotsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (CustomBotsResult, error) {
			args := v.(CustomBotsArgs)
			r, err := CustomBots(ctx, &args, opts...)
			var s CustomBotsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(CustomBotsResultOutput)
}

// A collection of arguments for invoking CustomBots.
type CustomBotsOutputArgs struct {
	// The domain names that need to be viewed.
	Host pulumi.StringPtrInput `pulumi:"host"`
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (CustomBotsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomBotsArgs)(nil)).Elem()
}

// A collection of values returned by CustomBots.
type CustomBotsResultOutput struct{ *pulumi.OutputState }

func (CustomBotsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomBotsResult)(nil)).Elem()
}

func (o CustomBotsResultOutput) ToCustomBotsResultOutput() CustomBotsResultOutput {
	return o
}

func (o CustomBotsResultOutput) ToCustomBotsResultOutputWithContext(ctx context.Context) CustomBotsResultOutput {
	return o
}

// The Details of Custom bot.
func (o CustomBotsResultOutput) Datas() CustomBotsDataArrayOutput {
	return o.ApplyT(func(v CustomBotsResult) []CustomBotsData { return v.Datas }).(CustomBotsDataArrayOutput)
}

func (o CustomBotsResultOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomBotsResult) *string { return v.Host }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o CustomBotsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CustomBotsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o CustomBotsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomBotsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o CustomBotsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomBotsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o CustomBotsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v CustomBotsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomBotsResultOutput{})
}
