// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of tls projects
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
//			_, err := tls.GetProjects(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetProjects(ctx *pulumi.Context, args *GetProjectsArgs, opts ...pulumi.InvokeOption) (*GetProjectsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectsResult
	err := ctx.Invoke("volcengine:tls/getProjects:getProjects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjects.
type GetProjectsArgs struct {
	// The IAM project name of the tls project.
	IamProjectName *string `pulumi:"iamProjectName"`
	// Whether to match accurately when filtering based on ProjectName.
	IsFullName *bool `pulumi:"isFullName"`
	// A Name Regex of tls project.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The id of tls project. This field supports fuzzy queries. It is not supported to specify both ProjectName and ProjectId at the same time.
	ProjectId *string `pulumi:"projectId"`
	// The name of tls project. This field supports fuzzy queries. It is not supported to specify both ProjectName and ProjectId at the same time.
	ProjectName *string `pulumi:"projectName"`
	// Tags.
	Tags []GetProjectsTag `pulumi:"tags"`
}

// A collection of values returned by getProjects.
type GetProjectsResult struct {
	// The IAM project name of the tls project.
	IamProjectName *string `pulumi:"iamProjectName"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	IsFullName *bool   `pulumi:"isFullName"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the tls project.
	ProjectId *string `pulumi:"projectId"`
	// The name of the tls project.
	ProjectName *string `pulumi:"projectName"`
	// Tags.
	Tags []GetProjectsTag `pulumi:"tags"`
	// The collection of tls project query.
	TlsProjects []GetProjectsTlsProject `pulumi:"tlsProjects"`
	// The total count of tls project query.
	TotalCount int `pulumi:"totalCount"`
}

func GetProjectsOutput(ctx *pulumi.Context, args GetProjectsOutputArgs, opts ...pulumi.InvokeOption) GetProjectsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectsResult, error) {
			args := v.(GetProjectsArgs)
			r, err := GetProjects(ctx, &args, opts...)
			var s GetProjectsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProjectsResultOutput)
}

// A collection of arguments for invoking getProjects.
type GetProjectsOutputArgs struct {
	// The IAM project name of the tls project.
	IamProjectName pulumi.StringPtrInput `pulumi:"iamProjectName"`
	// Whether to match accurately when filtering based on ProjectName.
	IsFullName pulumi.BoolPtrInput `pulumi:"isFullName"`
	// A Name Regex of tls project.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The id of tls project. This field supports fuzzy queries. It is not supported to specify both ProjectName and ProjectId at the same time.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The name of tls project. This field supports fuzzy queries. It is not supported to specify both ProjectName and ProjectId at the same time.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
	// Tags.
	Tags GetProjectsTagArrayInput `pulumi:"tags"`
}

func (GetProjectsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsArgs)(nil)).Elem()
}

// A collection of values returned by getProjects.
type GetProjectsResultOutput struct{ *pulumi.OutputState }

func (GetProjectsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsResult)(nil)).Elem()
}

func (o GetProjectsResultOutput) ToGetProjectsResultOutput() GetProjectsResultOutput {
	return o
}

func (o GetProjectsResultOutput) ToGetProjectsResultOutputWithContext(ctx context.Context) GetProjectsResultOutput {
	return o
}

// The IAM project name of the tls project.
func (o GetProjectsResultOutput) IamProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.IamProjectName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetProjectsResultOutput) IsFullName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *bool { return v.IsFullName }).(pulumi.BoolPtrOutput)
}

func (o GetProjectsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetProjectsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The ID of the tls project.
func (o GetProjectsResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// The name of the tls project.
func (o GetProjectsResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// Tags.
func (o GetProjectsResultOutput) Tags() GetProjectsTagArrayOutput {
	return o.ApplyT(func(v GetProjectsResult) []GetProjectsTag { return v.Tags }).(GetProjectsTagArrayOutput)
}

// The collection of tls project query.
func (o GetProjectsResultOutput) TlsProjects() GetProjectsTlsProjectArrayOutput {
	return o.ApplyT(func(v GetProjectsResult) []GetProjectsTlsProject { return v.TlsProjects }).(GetProjectsTlsProjectArrayOutput)
}

// The total count of tls project query.
func (o GetProjectsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetProjectsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectsResultOutput{})
}
