// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds ip lists
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/rds"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rds.GetIpLists(ctx, &rds.GetIpListsArgs{
//				InstanceId: "mysql-0fdd3bab2e7c",
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
// Deprecated: volcengine.rds.IpLists has been deprecated in favor of volcengine.rds.getIpLists
func IpLists(ctx *pulumi.Context, args *IpListsArgs, opts ...pulumi.InvokeOption) (*IpListsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv IpListsResult
	err := ctx.Invoke("volcengine:rds/ipLists:IpLists", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking IpLists.
type IpListsArgs struct {
	// The id of the RDS instance.
	InstanceId string `pulumi:"instanceId"`
	// A Name Regex of RDS ip list.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by IpLists.
type IpListsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstanceId string  `pulumi:"instanceId"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The collection of RDS ip list account query.
	RdsIpLists []IpListsRdsIpList `pulumi:"rdsIpLists"`
	// The total count of RDS ip list query.
	TotalCount int `pulumi:"totalCount"`
}

func IpListsOutput(ctx *pulumi.Context, args IpListsOutputArgs, opts ...pulumi.InvokeOption) IpListsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (IpListsResult, error) {
			args := v.(IpListsArgs)
			r, err := IpLists(ctx, &args, opts...)
			var s IpListsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(IpListsResultOutput)
}

// A collection of arguments for invoking IpLists.
type IpListsOutputArgs struct {
	// The id of the RDS instance.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// A Name Regex of RDS ip list.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (IpListsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpListsArgs)(nil)).Elem()
}

// A collection of values returned by IpLists.
type IpListsResultOutput struct{ *pulumi.OutputState }

func (IpListsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpListsResult)(nil)).Elem()
}

func (o IpListsResultOutput) ToIpListsResultOutput() IpListsResultOutput {
	return o
}

func (o IpListsResultOutput) ToIpListsResultOutputWithContext(ctx context.Context) IpListsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o IpListsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IpListsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o IpListsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v IpListsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o IpListsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpListsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o IpListsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpListsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The collection of RDS ip list account query.
func (o IpListsResultOutput) RdsIpLists() IpListsRdsIpListArrayOutput {
	return o.ApplyT(func(v IpListsResult) []IpListsRdsIpList { return v.RdsIpLists }).(IpListsRdsIpListArrayOutput)
}

// The total count of RDS ip list query.
func (o IpListsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v IpListsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(IpListsResultOutput{})
}
