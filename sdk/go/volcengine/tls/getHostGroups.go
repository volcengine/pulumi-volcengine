// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of tls host groups
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/tls"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tls.GetHostGroups(ctx, &tls.GetHostGroupsArgs{
//				HostGroupId:   pulumi.StringRef("fbea6619-7b0c-40f3-ac7e-45c63e3f676e"),
//				HostGroupName: pulumi.StringRef("cn"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetHostGroups(ctx *pulumi.Context, args *GetHostGroupsArgs, opts ...pulumi.InvokeOption) (*GetHostGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetHostGroupsResult
	err := ctx.Invoke("volcengine:tls/getHostGroups:getHostGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHostGroups.
type GetHostGroupsArgs struct {
	// Whether enable auto update.
	AutoUpdate *bool `pulumi:"autoUpdate"`
	// The id of host group.
	HostGroupId *string `pulumi:"hostGroupId"`
	// The name of host group.
	HostGroupName *string `pulumi:"hostGroupName"`
	// The identifier of host.
	HostIdentifier *string `pulumi:"hostIdentifier"`
	// The project name of iam.
	IamProjectName *string `pulumi:"iamProjectName"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// Whether enable service logging.
	ServiceLogging *bool `pulumi:"serviceLogging"`
}

// A collection of values returned by getHostGroups.
type GetHostGroupsResult struct {
	// Whether enable auto update.
	AutoUpdate *bool `pulumi:"autoUpdate"`
	// The id of host group.
	HostGroupId *string `pulumi:"hostGroupId"`
	// The name of host group.
	HostGroupName *string `pulumi:"hostGroupName"`
	// The identifier of host.
	HostIdentifier *string `pulumi:"hostIdentifier"`
	// The project name of iam.
	IamProjectName *string `pulumi:"iamProjectName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The collection of query.
	Infos      []GetHostGroupsInfo `pulumi:"infos"`
	OutputFile *string             `pulumi:"outputFile"`
	// Whether enable service logging.
	ServiceLogging *bool `pulumi:"serviceLogging"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func GetHostGroupsOutput(ctx *pulumi.Context, args GetHostGroupsOutputArgs, opts ...pulumi.InvokeOption) GetHostGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetHostGroupsResult, error) {
			args := v.(GetHostGroupsArgs)
			r, err := GetHostGroups(ctx, &args, opts...)
			var s GetHostGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetHostGroupsResultOutput)
}

// A collection of arguments for invoking getHostGroups.
type GetHostGroupsOutputArgs struct {
	// Whether enable auto update.
	AutoUpdate pulumi.BoolPtrInput `pulumi:"autoUpdate"`
	// The id of host group.
	HostGroupId pulumi.StringPtrInput `pulumi:"hostGroupId"`
	// The name of host group.
	HostGroupName pulumi.StringPtrInput `pulumi:"hostGroupName"`
	// The identifier of host.
	HostIdentifier pulumi.StringPtrInput `pulumi:"hostIdentifier"`
	// The project name of iam.
	IamProjectName pulumi.StringPtrInput `pulumi:"iamProjectName"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Whether enable service logging.
	ServiceLogging pulumi.BoolPtrInput `pulumi:"serviceLogging"`
}

func (GetHostGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHostGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getHostGroups.
type GetHostGroupsResultOutput struct{ *pulumi.OutputState }

func (GetHostGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHostGroupsResult)(nil)).Elem()
}

func (o GetHostGroupsResultOutput) ToGetHostGroupsResultOutput() GetHostGroupsResultOutput {
	return o
}

func (o GetHostGroupsResultOutput) ToGetHostGroupsResultOutputWithContext(ctx context.Context) GetHostGroupsResultOutput {
	return o
}

// Whether enable auto update.
func (o GetHostGroupsResultOutput) AutoUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetHostGroupsResult) *bool { return v.AutoUpdate }).(pulumi.BoolPtrOutput)
}

// The id of host group.
func (o GetHostGroupsResultOutput) HostGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHostGroupsResult) *string { return v.HostGroupId }).(pulumi.StringPtrOutput)
}

// The name of host group.
func (o GetHostGroupsResultOutput) HostGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHostGroupsResult) *string { return v.HostGroupName }).(pulumi.StringPtrOutput)
}

// The identifier of host.
func (o GetHostGroupsResultOutput) HostIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHostGroupsResult) *string { return v.HostIdentifier }).(pulumi.StringPtrOutput)
}

// The project name of iam.
func (o GetHostGroupsResultOutput) IamProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHostGroupsResult) *string { return v.IamProjectName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetHostGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetHostGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The collection of query.
func (o GetHostGroupsResultOutput) Infos() GetHostGroupsInfoArrayOutput {
	return o.ApplyT(func(v GetHostGroupsResult) []GetHostGroupsInfo { return v.Infos }).(GetHostGroupsInfoArrayOutput)
}

func (o GetHostGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHostGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// Whether enable service logging.
func (o GetHostGroupsResultOutput) ServiceLogging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetHostGroupsResult) *bool { return v.ServiceLogging }).(pulumi.BoolPtrOutput)
}

// The total count of query.
func (o GetHostGroupsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetHostGroupsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetHostGroupsResultOutput{})
}
