// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of waf cc rules
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
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// _, err := waf.GetCcRules(ctx, &waf.GetCcRulesArgs{
// CcTypes: interface{}{
// 1,
// },
// Host: "www.tf-test.com",
// PathOrderBy: "ASC",
// RuleName: pulumi.StringRef("tf"),
// }, nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func GetCcRules(ctx *pulumi.Context, args *GetCcRulesArgs, opts ...pulumi.InvokeOption) (*GetCcRulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCcRulesResult
	err := ctx.Invoke("volcengine:waf/getCcRules:getCcRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCcRules.
type GetCcRulesArgs struct {
	// The actions performed on subsequent requests after meeting the statistical conditions.
	CcTypes []int `pulumi:"ccTypes"`
	// Website domain names that require the setting of protection rules.
	Host string `pulumi:"host"`
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The list shows the order.
	PathOrderBy string `pulumi:"pathOrderBy"`
	// Search by rule name in a fuzzy manner.
	RuleName *string `pulumi:"ruleName"`
	// Search precisely according to the rule ID.
	RuleTag *string `pulumi:"ruleTag"`
	// Fuzzy search by the requested path.
	Url *string `pulumi:"url"`
}

// A collection of values returned by getCcRules.
type GetCcRulesResult struct {
	// The actions performed on subsequent requests after meeting the statistical conditions.
	CcTypes []int `pulumi:"ccTypes"`
	// The collection of query.
	Datas []GetCcRulesData `pulumi:"datas"`
	// Protected website domain names.
	Host string `pulumi:"host"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	NameRegex   *string `pulumi:"nameRegex"`
	OutputFile  *string `pulumi:"outputFile"`
	PathOrderBy string  `pulumi:"pathOrderBy"`
	RuleName    *string `pulumi:"ruleName"`
	// Rule label, that is, the complete rule ID.
	RuleTag *string `pulumi:"ruleTag"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
	// The requested path.
	Url *string `pulumi:"url"`
}

func GetCcRulesOutput(ctx *pulumi.Context, args GetCcRulesOutputArgs, opts ...pulumi.InvokeOption) GetCcRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCcRulesResult, error) {
			args := v.(GetCcRulesArgs)
			r, err := GetCcRules(ctx, &args, opts...)
			var s GetCcRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCcRulesResultOutput)
}

// A collection of arguments for invoking getCcRules.
type GetCcRulesOutputArgs struct {
	// The actions performed on subsequent requests after meeting the statistical conditions.
	CcTypes pulumi.IntArrayInput `pulumi:"ccTypes"`
	// Website domain names that require the setting of protection rules.
	Host pulumi.StringInput `pulumi:"host"`
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The list shows the order.
	PathOrderBy pulumi.StringInput `pulumi:"pathOrderBy"`
	// Search by rule name in a fuzzy manner.
	RuleName pulumi.StringPtrInput `pulumi:"ruleName"`
	// Search precisely according to the rule ID.
	RuleTag pulumi.StringPtrInput `pulumi:"ruleTag"`
	// Fuzzy search by the requested path.
	Url pulumi.StringPtrInput `pulumi:"url"`
}

func (GetCcRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCcRulesArgs)(nil)).Elem()
}

// A collection of values returned by getCcRules.
type GetCcRulesResultOutput struct{ *pulumi.OutputState }

func (GetCcRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCcRulesResult)(nil)).Elem()
}

func (o GetCcRulesResultOutput) ToGetCcRulesResultOutput() GetCcRulesResultOutput {
	return o
}

func (o GetCcRulesResultOutput) ToGetCcRulesResultOutputWithContext(ctx context.Context) GetCcRulesResultOutput {
	return o
}

// The actions performed on subsequent requests after meeting the statistical conditions.
func (o GetCcRulesResultOutput) CcTypes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetCcRulesResult) []int { return v.CcTypes }).(pulumi.IntArrayOutput)
}

// The collection of query.
func (o GetCcRulesResultOutput) Datas() GetCcRulesDataArrayOutput {
	return o.ApplyT(func(v GetCcRulesResult) []GetCcRulesData { return v.Datas }).(GetCcRulesDataArrayOutput)
}

// Protected website domain names.
func (o GetCcRulesResultOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v GetCcRulesResult) string { return v.Host }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCcRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCcRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCcRulesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCcRulesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetCcRulesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCcRulesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetCcRulesResultOutput) PathOrderBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetCcRulesResult) string { return v.PathOrderBy }).(pulumi.StringOutput)
}

func (o GetCcRulesResultOutput) RuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCcRulesResult) *string { return v.RuleName }).(pulumi.StringPtrOutput)
}

// Rule label, that is, the complete rule ID.
func (o GetCcRulesResultOutput) RuleTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCcRulesResult) *string { return v.RuleTag }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o GetCcRulesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetCcRulesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The requested path.
func (o GetCcRulesResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCcRulesResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCcRulesResultOutput{})
}
