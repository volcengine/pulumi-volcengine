// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of security groups
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpc.SecurityGroups(ctx, &vpc.SecurityGroupsArgs{
//				Ids: []string{
//					"sg-273ycgql3ig3k7fap8t3dyvqx",
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
func SecurityGroups(ctx *pulumi.Context, args *SecurityGroupsArgs, opts ...pulumi.InvokeOption) (*SecurityGroupsResult, error) {
	var rv SecurityGroupsResult
	err := ctx.Invoke("volcengine:vpc/securityGroups:SecurityGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking SecurityGroups.
type SecurityGroupsArgs struct {
	// A list of SecurityGroup IDs.
	Ids []string `pulumi:"ids"`
	// A Name Regex of SecurityGroup.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The ProjectName of SecurityGroup.
	ProjectName *string `pulumi:"projectName"`
	// The list of security group name to query.
	SecurityGroupNames []string `pulumi:"securityGroupNames"`
	// Tags.
	Tags []SecurityGroupsTag `pulumi:"tags"`
	// The ID of vpc where security group is located.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by SecurityGroups.
type SecurityGroupsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	OutputFile *string  `pulumi:"outputFile"`
	// The ProjectName of SecurityGroup.
	ProjectName        *string  `pulumi:"projectName"`
	SecurityGroupNames []string `pulumi:"securityGroupNames"`
	// The collection of SecurityGroup query.
	SecurityGroups []SecurityGroupsSecurityGroup `pulumi:"securityGroups"`
	// Tags.
	Tags []SecurityGroupsTag `pulumi:"tags"`
	// The total count of SecurityGroup query.
	TotalCount int `pulumi:"totalCount"`
	// The ID of Vpc.
	VpcId *string `pulumi:"vpcId"`
}

func SecurityGroupsOutput(ctx *pulumi.Context, args SecurityGroupsOutputArgs, opts ...pulumi.InvokeOption) SecurityGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (SecurityGroupsResult, error) {
			args := v.(SecurityGroupsArgs)
			r, err := SecurityGroups(ctx, &args, opts...)
			var s SecurityGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(SecurityGroupsResultOutput)
}

// A collection of arguments for invoking SecurityGroups.
type SecurityGroupsOutputArgs struct {
	// A list of SecurityGroup IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A Name Regex of SecurityGroup.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ProjectName of SecurityGroup.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
	// The list of security group name to query.
	SecurityGroupNames pulumi.StringArrayInput `pulumi:"securityGroupNames"`
	// Tags.
	Tags SecurityGroupsTagArrayInput `pulumi:"tags"`
	// The ID of vpc where security group is located.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (SecurityGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroupsArgs)(nil)).Elem()
}

// A collection of values returned by SecurityGroups.
type SecurityGroupsResultOutput struct{ *pulumi.OutputState }

func (SecurityGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroupsResult)(nil)).Elem()
}

func (o SecurityGroupsResultOutput) ToSecurityGroupsResultOutput() SecurityGroupsResultOutput {
	return o
}

func (o SecurityGroupsResultOutput) ToSecurityGroupsResultOutputWithContext(ctx context.Context) SecurityGroupsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o SecurityGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o SecurityGroupsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityGroupsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o SecurityGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o SecurityGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The ProjectName of SecurityGroup.
func (o SecurityGroupsResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityGroupsResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

func (o SecurityGroupsResultOutput) SecurityGroupNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityGroupsResult) []string { return v.SecurityGroupNames }).(pulumi.StringArrayOutput)
}

// The collection of SecurityGroup query.
func (o SecurityGroupsResultOutput) SecurityGroups() SecurityGroupsSecurityGroupArrayOutput {
	return o.ApplyT(func(v SecurityGroupsResult) []SecurityGroupsSecurityGroup { return v.SecurityGroups }).(SecurityGroupsSecurityGroupArrayOutput)
}

// Tags.
func (o SecurityGroupsResultOutput) Tags() SecurityGroupsTagArrayOutput {
	return o.ApplyT(func(v SecurityGroupsResult) []SecurityGroupsTag { return v.Tags }).(SecurityGroupsTagArrayOutput)
}

// The total count of SecurityGroup query.
func (o SecurityGroupsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v SecurityGroupsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The ID of Vpc.
func (o SecurityGroupsResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityGroupsResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityGroupsResultOutput{})
}