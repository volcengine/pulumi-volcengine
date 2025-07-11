// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of vpcs
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpc.GetVpcs(ctx, &vpc.GetVpcsArgs{
//				Ids: []string{
//					"vpc-mizl7m1kqccg5smt1bdpijuj",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetVpcs(ctx *pulumi.Context, args *GetVpcsArgs, opts ...pulumi.InvokeOption) (*GetVpcsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVpcsResult
	err := ctx.Invoke("volcengine:vpc/getVpcs:getVpcs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcs.
type GetVpcsArgs struct {
	// A list of VPC IDs.
	Ids []string `pulumi:"ids"`
	// A Name Regex of Vpc.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The ProjectName of the VPC.
	ProjectName *string `pulumi:"projectName"`
	// Tags.
	Tags []GetVpcsTag `pulumi:"tags"`
	// The vpc name to query.
	VpcName *string `pulumi:"vpcName"`
	// The owner ID of the vpc.
	VpcOwnerId *int `pulumi:"vpcOwnerId"`
}

// A collection of values returned by getVpcs.
type GetVpcsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	OutputFile *string  `pulumi:"outputFile"`
	// The ProjectName of the VPC.
	ProjectName *string `pulumi:"projectName"`
	// Tags.
	Tags []GetVpcsTag `pulumi:"tags"`
	// The total count of Vpc query.
	TotalCount int `pulumi:"totalCount"`
	// The name of VPC.
	VpcName    *string `pulumi:"vpcName"`
	VpcOwnerId *int    `pulumi:"vpcOwnerId"`
	// The collection of Vpc query.
	Vpcs []GetVpcsVpc `pulumi:"vpcs"`
}

func GetVpcsOutput(ctx *pulumi.Context, args GetVpcsOutputArgs, opts ...pulumi.InvokeOption) GetVpcsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVpcsResult, error) {
			args := v.(GetVpcsArgs)
			r, err := GetVpcs(ctx, &args, opts...)
			var s GetVpcsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVpcsResultOutput)
}

// A collection of arguments for invoking getVpcs.
type GetVpcsOutputArgs struct {
	// A list of VPC IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A Name Regex of Vpc.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ProjectName of the VPC.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
	// Tags.
	Tags GetVpcsTagArrayInput `pulumi:"tags"`
	// The vpc name to query.
	VpcName pulumi.StringPtrInput `pulumi:"vpcName"`
	// The owner ID of the vpc.
	VpcOwnerId pulumi.IntPtrInput `pulumi:"vpcOwnerId"`
}

func (GetVpcsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcsArgs)(nil)).Elem()
}

// A collection of values returned by getVpcs.
type GetVpcsResultOutput struct{ *pulumi.OutputState }

func (GetVpcsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcsResult)(nil)).Elem()
}

func (o GetVpcsResultOutput) ToGetVpcsResultOutput() GetVpcsResultOutput {
	return o
}

func (o GetVpcsResultOutput) ToGetVpcsResultOutputWithContext(ctx context.Context) GetVpcsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetVpcsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVpcsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVpcsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetVpcsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetVpcsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The ProjectName of the VPC.
func (o GetVpcsResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcsResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// Tags.
func (o GetVpcsResultOutput) Tags() GetVpcsTagArrayOutput {
	return o.ApplyT(func(v GetVpcsResult) []GetVpcsTag { return v.Tags }).(GetVpcsTagArrayOutput)
}

// The total count of Vpc query.
func (o GetVpcsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetVpcsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The name of VPC.
func (o GetVpcsResultOutput) VpcName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcsResult) *string { return v.VpcName }).(pulumi.StringPtrOutput)
}

func (o GetVpcsResultOutput) VpcOwnerId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetVpcsResult) *int { return v.VpcOwnerId }).(pulumi.IntPtrOutput)
}

// The collection of Vpc query.
func (o GetVpcsResultOutput) Vpcs() GetVpcsVpcArrayOutput {
	return o.ApplyT(func(v GetVpcsResult) []GetVpcsVpc { return v.Vpcs }).(GetVpcsVpcArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVpcsResultOutput{})
}
