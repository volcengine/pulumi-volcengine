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
func GetCustomBots(ctx *pulumi.Context, args *GetCustomBotsArgs, opts ...pulumi.InvokeOption) (*GetCustomBotsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCustomBotsResult
	err := ctx.Invoke("volcengine:waf/getCustomBots:getCustomBots", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCustomBots.
type GetCustomBotsArgs struct {
	// The domain names that need to be viewed.
	Host *string `pulumi:"host"`
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getCustomBots.
type GetCustomBotsResult struct {
	// The Details of Custom bot.
	Datas []GetCustomBotsData `pulumi:"datas"`
	Host  *string             `pulumi:"host"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func GetCustomBotsOutput(ctx *pulumi.Context, args GetCustomBotsOutputArgs, opts ...pulumi.InvokeOption) GetCustomBotsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCustomBotsResult, error) {
			args := v.(GetCustomBotsArgs)
			r, err := GetCustomBots(ctx, &args, opts...)
			var s GetCustomBotsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCustomBotsResultOutput)
}

// A collection of arguments for invoking getCustomBots.
type GetCustomBotsOutputArgs struct {
	// The domain names that need to be viewed.
	Host pulumi.StringPtrInput `pulumi:"host"`
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetCustomBotsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomBotsArgs)(nil)).Elem()
}

// A collection of values returned by getCustomBots.
type GetCustomBotsResultOutput struct{ *pulumi.OutputState }

func (GetCustomBotsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomBotsResult)(nil)).Elem()
}

func (o GetCustomBotsResultOutput) ToGetCustomBotsResultOutput() GetCustomBotsResultOutput {
	return o
}

func (o GetCustomBotsResultOutput) ToGetCustomBotsResultOutputWithContext(ctx context.Context) GetCustomBotsResultOutput {
	return o
}

// The Details of Custom bot.
func (o GetCustomBotsResultOutput) Datas() GetCustomBotsDataArrayOutput {
	return o.ApplyT(func(v GetCustomBotsResult) []GetCustomBotsData { return v.Datas }).(GetCustomBotsDataArrayOutput)
}

func (o GetCustomBotsResultOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCustomBotsResult) *string { return v.Host }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCustomBotsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomBotsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCustomBotsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCustomBotsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetCustomBotsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCustomBotsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o GetCustomBotsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetCustomBotsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCustomBotsResultOutput{})
}
