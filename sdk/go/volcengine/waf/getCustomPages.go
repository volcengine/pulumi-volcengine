// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of waf custom pages
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
//			_, err := waf.GetCustomPages(ctx, &waf.GetCustomPagesArgs{
//				Host: "www.tf-test.com",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCustomPages(ctx *pulumi.Context, args *GetCustomPagesArgs, opts ...pulumi.InvokeOption) (*GetCustomPagesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCustomPagesResult
	err := ctx.Invoke("volcengine:waf/getCustomPages:getCustomPages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCustomPages.
type GetCustomPagesArgs struct {
	// The domain names that need to be viewed.
	Host string `pulumi:"host"`
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The name of the project to which your domain names belong.
	ProjectName *string `pulumi:"projectName"`
	// Unique identification of the rules.
	RuleTag *string `pulumi:"ruleTag"`
}

// A collection of values returned by getCustomPages.
type GetCustomPagesResult struct {
	// Details of the rules.
	Datas []GetCustomPagesData `pulumi:"datas"`
	// Domain name to be protected.
	Host string `pulumi:"host"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	NameRegex   *string `pulumi:"nameRegex"`
	OutputFile  *string `pulumi:"outputFile"`
	ProjectName *string `pulumi:"projectName"`
	// Unique identification of the rules.
	RuleTag *string `pulumi:"ruleTag"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func GetCustomPagesOutput(ctx *pulumi.Context, args GetCustomPagesOutputArgs, opts ...pulumi.InvokeOption) GetCustomPagesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCustomPagesResult, error) {
			args := v.(GetCustomPagesArgs)
			r, err := GetCustomPages(ctx, &args, opts...)
			var s GetCustomPagesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCustomPagesResultOutput)
}

// A collection of arguments for invoking getCustomPages.
type GetCustomPagesOutputArgs struct {
	// The domain names that need to be viewed.
	Host pulumi.StringInput `pulumi:"host"`
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The name of the project to which your domain names belong.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
	// Unique identification of the rules.
	RuleTag pulumi.StringPtrInput `pulumi:"ruleTag"`
}

func (GetCustomPagesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomPagesArgs)(nil)).Elem()
}

// A collection of values returned by getCustomPages.
type GetCustomPagesResultOutput struct{ *pulumi.OutputState }

func (GetCustomPagesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomPagesResult)(nil)).Elem()
}

func (o GetCustomPagesResultOutput) ToGetCustomPagesResultOutput() GetCustomPagesResultOutput {
	return o
}

func (o GetCustomPagesResultOutput) ToGetCustomPagesResultOutputWithContext(ctx context.Context) GetCustomPagesResultOutput {
	return o
}

// Details of the rules.
func (o GetCustomPagesResultOutput) Datas() GetCustomPagesDataArrayOutput {
	return o.ApplyT(func(v GetCustomPagesResult) []GetCustomPagesData { return v.Datas }).(GetCustomPagesDataArrayOutput)
}

// Domain name to be protected.
func (o GetCustomPagesResultOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomPagesResult) string { return v.Host }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCustomPagesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomPagesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCustomPagesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCustomPagesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetCustomPagesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCustomPagesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetCustomPagesResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCustomPagesResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// Unique identification of the rules.
func (o GetCustomPagesResultOutput) RuleTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCustomPagesResult) *string { return v.RuleTag }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o GetCustomPagesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetCustomPagesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCustomPagesResultOutput{})
}
