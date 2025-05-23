// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds parameter templates
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
//			_, err := rds.GetParameterTemplates(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetParameterTemplates(ctx *pulumi.Context, args *GetParameterTemplatesArgs, opts ...pulumi.InvokeOption) (*GetParameterTemplatesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetParameterTemplatesResult
	err := ctx.Invoke("volcengine:rds/getParameterTemplates:getParameterTemplates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getParameterTemplates.
type GetParameterTemplatesArgs struct {
	// A Name Regex of RDS parameter template.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// Parameter template type, range of values:
	// DBEngine - Engine parameters.
	TemplateCategory *string `pulumi:"templateCategory"`
	// Template source, value range:
	// System - System
	// User - the user.
	TemplateSource *string `pulumi:"templateSource"`
	// Parameter template database type, range of values:
	// MySQL - MySQL database.
	TemplateType *string `pulumi:"templateType"`
	// Parameter template database version, value range:
	// MySQL_Community_5_7 - MySQL 5.7
	// MySQL_8_0 - MySQL 8.0.
	TemplateTypeVersion *string `pulumi:"templateTypeVersion"`
}

// A collection of values returned by getParameterTemplates.
type GetParameterTemplatesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The collection of RDS parameter templates query.
	RdsParameterTemplates []GetParameterTemplatesRdsParameterTemplate `pulumi:"rdsParameterTemplates"`
	TemplateCategory      *string                                     `pulumi:"templateCategory"`
	TemplateSource        *string                                     `pulumi:"templateSource"`
	// Parameter template database type, range of values:
	// MySQL - MySQL database.
	TemplateType *string `pulumi:"templateType"`
	// Parameter template database version, value range:
	// MySQL_Community_5_7 - MySQL 5.7
	// MySQL_8_0 - MySQL 8.0.
	TemplateTypeVersion *string `pulumi:"templateTypeVersion"`
	// The total count of RDS parameter templates query.
	TotalCount int `pulumi:"totalCount"`
}

func GetParameterTemplatesOutput(ctx *pulumi.Context, args GetParameterTemplatesOutputArgs, opts ...pulumi.InvokeOption) GetParameterTemplatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetParameterTemplatesResult, error) {
			args := v.(GetParameterTemplatesArgs)
			r, err := GetParameterTemplates(ctx, &args, opts...)
			var s GetParameterTemplatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetParameterTemplatesResultOutput)
}

// A collection of arguments for invoking getParameterTemplates.
type GetParameterTemplatesOutputArgs struct {
	// A Name Regex of RDS parameter template.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Parameter template type, range of values:
	// DBEngine - Engine parameters.
	TemplateCategory pulumi.StringPtrInput `pulumi:"templateCategory"`
	// Template source, value range:
	// System - System
	// User - the user.
	TemplateSource pulumi.StringPtrInput `pulumi:"templateSource"`
	// Parameter template database type, range of values:
	// MySQL - MySQL database.
	TemplateType pulumi.StringPtrInput `pulumi:"templateType"`
	// Parameter template database version, value range:
	// MySQL_Community_5_7 - MySQL 5.7
	// MySQL_8_0 - MySQL 8.0.
	TemplateTypeVersion pulumi.StringPtrInput `pulumi:"templateTypeVersion"`
}

func (GetParameterTemplatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetParameterTemplatesArgs)(nil)).Elem()
}

// A collection of values returned by getParameterTemplates.
type GetParameterTemplatesResultOutput struct{ *pulumi.OutputState }

func (GetParameterTemplatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetParameterTemplatesResult)(nil)).Elem()
}

func (o GetParameterTemplatesResultOutput) ToGetParameterTemplatesResultOutput() GetParameterTemplatesResultOutput {
	return o
}

func (o GetParameterTemplatesResultOutput) ToGetParameterTemplatesResultOutputWithContext(ctx context.Context) GetParameterTemplatesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetParameterTemplatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetParameterTemplatesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetParameterTemplatesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetParameterTemplatesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetParameterTemplatesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetParameterTemplatesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The collection of RDS parameter templates query.
func (o GetParameterTemplatesResultOutput) RdsParameterTemplates() GetParameterTemplatesRdsParameterTemplateArrayOutput {
	return o.ApplyT(func(v GetParameterTemplatesResult) []GetParameterTemplatesRdsParameterTemplate {
		return v.RdsParameterTemplates
	}).(GetParameterTemplatesRdsParameterTemplateArrayOutput)
}

func (o GetParameterTemplatesResultOutput) TemplateCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetParameterTemplatesResult) *string { return v.TemplateCategory }).(pulumi.StringPtrOutput)
}

func (o GetParameterTemplatesResultOutput) TemplateSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetParameterTemplatesResult) *string { return v.TemplateSource }).(pulumi.StringPtrOutput)
}

// Parameter template database type, range of values:
// MySQL - MySQL database.
func (o GetParameterTemplatesResultOutput) TemplateType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetParameterTemplatesResult) *string { return v.TemplateType }).(pulumi.StringPtrOutput)
}

// Parameter template database version, value range:
// MySQL_Community_5_7 - MySQL 5.7
// MySQL_8_0 - MySQL 8.0.
func (o GetParameterTemplatesResultOutput) TemplateTypeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetParameterTemplatesResult) *string { return v.TemplateTypeVersion }).(pulumi.StringPtrOutput)
}

// The total count of RDS parameter templates query.
func (o GetParameterTemplatesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetParameterTemplatesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetParameterTemplatesResultOutput{})
}
