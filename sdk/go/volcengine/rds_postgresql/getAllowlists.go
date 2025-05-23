// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds_postgresql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of rds postgresql allowlists
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/rds_postgresql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rds_postgresql.GetAllowlists(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetAllowlists(ctx *pulumi.Context, args *GetAllowlistsArgs, opts ...pulumi.InvokeOption) (*GetAllowlistsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAllowlistsResult
	err := ctx.Invoke("volcengine:rds_postgresql/getAllowlists:getAllowlists", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAllowlists.
type GetAllowlistsArgs struct {
	// The id of the postgresql Instance.
	InstanceId *string `pulumi:"instanceId"`
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getAllowlists.
type GetAllowlistsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The id of the postgresql instance.
	InstanceId *string `pulumi:"instanceId"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The list of postgresql allowed list.
	PostgresqlAllowLists []GetAllowlistsPostgresqlAllowList `pulumi:"postgresqlAllowLists"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func GetAllowlistsOutput(ctx *pulumi.Context, args GetAllowlistsOutputArgs, opts ...pulumi.InvokeOption) GetAllowlistsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAllowlistsResult, error) {
			args := v.(GetAllowlistsArgs)
			r, err := GetAllowlists(ctx, &args, opts...)
			var s GetAllowlistsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAllowlistsResultOutput)
}

// A collection of arguments for invoking getAllowlists.
type GetAllowlistsOutputArgs struct {
	// The id of the postgresql Instance.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetAllowlistsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAllowlistsArgs)(nil)).Elem()
}

// A collection of values returned by getAllowlists.
type GetAllowlistsResultOutput struct{ *pulumi.OutputState }

func (GetAllowlistsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAllowlistsResult)(nil)).Elem()
}

func (o GetAllowlistsResultOutput) ToGetAllowlistsResultOutput() GetAllowlistsResultOutput {
	return o
}

func (o GetAllowlistsResultOutput) ToGetAllowlistsResultOutputWithContext(ctx context.Context) GetAllowlistsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetAllowlistsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAllowlistsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The id of the postgresql instance.
func (o GetAllowlistsResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAllowlistsResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o GetAllowlistsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAllowlistsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetAllowlistsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAllowlistsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The list of postgresql allowed list.
func (o GetAllowlistsResultOutput) PostgresqlAllowLists() GetAllowlistsPostgresqlAllowListArrayOutput {
	return o.ApplyT(func(v GetAllowlistsResult) []GetAllowlistsPostgresqlAllowList { return v.PostgresqlAllowLists }).(GetAllowlistsPostgresqlAllowListArrayOutput)
}

// The total count of query.
func (o GetAllowlistsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetAllowlistsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAllowlistsResultOutput{})
}
