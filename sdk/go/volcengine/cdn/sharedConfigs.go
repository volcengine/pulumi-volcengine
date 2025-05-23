// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of cdn shared configs
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/cdn"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cdn.GetSharedConfigs(ctx, &cdn.GetSharedConfigsArgs{
//				ConfigName:  pulumi.StringRef("tf-test"),
//				ConfigType:  pulumi.StringRef("allow_ip_access_rule"),
//				ProjectName: pulumi.StringRef("default"),
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
// Deprecated: volcengine.cdn.SharedConfigs has been deprecated in favor of volcengine.cdn.getSharedConfigs
func SharedConfigs(ctx *pulumi.Context, args *SharedConfigsArgs, opts ...pulumi.InvokeOption) (*SharedConfigsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv SharedConfigsResult
	err := ctx.Invoke("volcengine:cdn/sharedConfigs:SharedConfigs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking SharedConfigs.
type SharedConfigsArgs struct {
	// The name of the shared config.
	ConfigName *string `pulumi:"configName"`
	// The type of the shared config.
	ConfigType *string `pulumi:"configType"`
	// The config type list. The parameter value can be a combination of available values for ConfigType. ConfigType and ConfigTypeList cannot be specified at the same time.
	ConfigTypeLists []string `pulumi:"configTypeLists"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The name of the project.
	ProjectName *string `pulumi:"projectName"`
}

// A collection of values returned by SharedConfigs.
type SharedConfigsResult struct {
	// The collection of query.
	ConfigDatas []SharedConfigsConfigData `pulumi:"configDatas"`
	// The name of the config.
	ConfigName *string `pulumi:"configName"`
	// The type of the config.
	ConfigType      *string  `pulumi:"configType"`
	ConfigTypeLists []string `pulumi:"configTypeLists"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	OutputFile *string `pulumi:"outputFile"`
	// The name of the project.
	ProjectName *string `pulumi:"projectName"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func SharedConfigsOutput(ctx *pulumi.Context, args SharedConfigsOutputArgs, opts ...pulumi.InvokeOption) SharedConfigsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (SharedConfigsResult, error) {
			args := v.(SharedConfigsArgs)
			r, err := SharedConfigs(ctx, &args, opts...)
			var s SharedConfigsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(SharedConfigsResultOutput)
}

// A collection of arguments for invoking SharedConfigs.
type SharedConfigsOutputArgs struct {
	// The name of the shared config.
	ConfigName pulumi.StringPtrInput `pulumi:"configName"`
	// The type of the shared config.
	ConfigType pulumi.StringPtrInput `pulumi:"configType"`
	// The config type list. The parameter value can be a combination of available values for ConfigType. ConfigType and ConfigTypeList cannot be specified at the same time.
	ConfigTypeLists pulumi.StringArrayInput `pulumi:"configTypeLists"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The name of the project.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
}

func (SharedConfigsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedConfigsArgs)(nil)).Elem()
}

// A collection of values returned by SharedConfigs.
type SharedConfigsResultOutput struct{ *pulumi.OutputState }

func (SharedConfigsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedConfigsResult)(nil)).Elem()
}

func (o SharedConfigsResultOutput) ToSharedConfigsResultOutput() SharedConfigsResultOutput {
	return o
}

func (o SharedConfigsResultOutput) ToSharedConfigsResultOutputWithContext(ctx context.Context) SharedConfigsResultOutput {
	return o
}

// The collection of query.
func (o SharedConfigsResultOutput) ConfigDatas() SharedConfigsConfigDataArrayOutput {
	return o.ApplyT(func(v SharedConfigsResult) []SharedConfigsConfigData { return v.ConfigDatas }).(SharedConfigsConfigDataArrayOutput)
}

// The name of the config.
func (o SharedConfigsResultOutput) ConfigName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedConfigsResult) *string { return v.ConfigName }).(pulumi.StringPtrOutput)
}

// The type of the config.
func (o SharedConfigsResultOutput) ConfigType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedConfigsResult) *string { return v.ConfigType }).(pulumi.StringPtrOutput)
}

func (o SharedConfigsResultOutput) ConfigTypeLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SharedConfigsResult) []string { return v.ConfigTypeLists }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o SharedConfigsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SharedConfigsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o SharedConfigsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedConfigsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The name of the project.
func (o SharedConfigsResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedConfigsResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o SharedConfigsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v SharedConfigsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(SharedConfigsResultOutput{})
}
