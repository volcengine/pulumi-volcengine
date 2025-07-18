// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of waf acl rules
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
// _, err := waf.GetAclRules(ctx, &waf.GetAclRulesArgs{
// AclType: "Block",
// Actions: []string{
// "observe",
// },
// DefenceHosts: []string{
// "www.tf-test.com",
// },
// Enables: interface{}{
// 1,
// },
// ProjectName: pulumi.StringRef("default"),
// RuleName: pulumi.StringRef("tf-test"),
// TimeOrderBy: pulumi.StringRef("ASC"),
// }, nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func GetAclRules(ctx *pulumi.Context, args *GetAclRulesArgs, opts ...pulumi.InvokeOption) (*GetAclRulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAclRulesResult
	err := ctx.Invoke("volcengine:waf/getAclRules:getAclRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAclRules.
type GetAclRulesArgs struct {
	// The types of access control rules.
	AclType string `pulumi:"aclType"`
	// Action to be taken on requests that match the rule.
	Actions []string `pulumi:"actions"`
	// The list of queried domain names.
	DefenceHosts []string `pulumi:"defenceHosts"`
	// The enabled status of the rule.
	Enables []int `pulumi:"enables"`
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The name of the project to which your domain names belong.
	ProjectName *string `pulumi:"projectName"`
	// Rule name, fuzzy search.
	RuleName *string `pulumi:"ruleName"`
	// Rule unique identifier, precise search.
	RuleTag *string `pulumi:"ruleTag"`
	// The list shows the timing sequence.
	TimeOrderBy *string `pulumi:"timeOrderBy"`
}

// A collection of values returned by getAclRules.
type GetAclRulesResult struct {
	AclType string `pulumi:"aclType"`
	// Action to be taken on requests that match the rule.
	Actions      []string `pulumi:"actions"`
	DefenceHosts []string `pulumi:"defenceHosts"`
	// Whether to enable the rule.
	Enables []int `pulumi:"enables"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	NameRegex   *string `pulumi:"nameRegex"`
	OutputFile  *string `pulumi:"outputFile"`
	ProjectName *string `pulumi:"projectName"`
	RuleName    *string `pulumi:"ruleName"`
	// Rule unique identifier.
	RuleTag *string `pulumi:"ruleTag"`
	// Details of the rules.
	Rules       []GetAclRulesRule `pulumi:"rules"`
	TimeOrderBy *string           `pulumi:"timeOrderBy"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func GetAclRulesOutput(ctx *pulumi.Context, args GetAclRulesOutputArgs, opts ...pulumi.InvokeOption) GetAclRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAclRulesResult, error) {
			args := v.(GetAclRulesArgs)
			r, err := GetAclRules(ctx, &args, opts...)
			var s GetAclRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAclRulesResultOutput)
}

// A collection of arguments for invoking getAclRules.
type GetAclRulesOutputArgs struct {
	// The types of access control rules.
	AclType pulumi.StringInput `pulumi:"aclType"`
	// Action to be taken on requests that match the rule.
	Actions pulumi.StringArrayInput `pulumi:"actions"`
	// The list of queried domain names.
	DefenceHosts pulumi.StringArrayInput `pulumi:"defenceHosts"`
	// The enabled status of the rule.
	Enables pulumi.IntArrayInput `pulumi:"enables"`
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The name of the project to which your domain names belong.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
	// Rule name, fuzzy search.
	RuleName pulumi.StringPtrInput `pulumi:"ruleName"`
	// Rule unique identifier, precise search.
	RuleTag pulumi.StringPtrInput `pulumi:"ruleTag"`
	// The list shows the timing sequence.
	TimeOrderBy pulumi.StringPtrInput `pulumi:"timeOrderBy"`
}

func (GetAclRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAclRulesArgs)(nil)).Elem()
}

// A collection of values returned by getAclRules.
type GetAclRulesResultOutput struct{ *pulumi.OutputState }

func (GetAclRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAclRulesResult)(nil)).Elem()
}

func (o GetAclRulesResultOutput) ToGetAclRulesResultOutput() GetAclRulesResultOutput {
	return o
}

func (o GetAclRulesResultOutput) ToGetAclRulesResultOutputWithContext(ctx context.Context) GetAclRulesResultOutput {
	return o
}

func (o GetAclRulesResultOutput) AclType() pulumi.StringOutput {
	return o.ApplyT(func(v GetAclRulesResult) string { return v.AclType }).(pulumi.StringOutput)
}

// Action to be taken on requests that match the rule.
func (o GetAclRulesResultOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAclRulesResult) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

func (o GetAclRulesResultOutput) DefenceHosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAclRulesResult) []string { return v.DefenceHosts }).(pulumi.StringArrayOutput)
}

// Whether to enable the rule.
func (o GetAclRulesResultOutput) Enables() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetAclRulesResult) []int { return v.Enables }).(pulumi.IntArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAclRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAclRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAclRulesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAclRulesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetAclRulesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAclRulesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetAclRulesResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAclRulesResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

func (o GetAclRulesResultOutput) RuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAclRulesResult) *string { return v.RuleName }).(pulumi.StringPtrOutput)
}

// Rule unique identifier.
func (o GetAclRulesResultOutput) RuleTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAclRulesResult) *string { return v.RuleTag }).(pulumi.StringPtrOutput)
}

// Details of the rules.
func (o GetAclRulesResultOutput) Rules() GetAclRulesRuleArrayOutput {
	return o.ApplyT(func(v GetAclRulesResult) []GetAclRulesRule { return v.Rules }).(GetAclRulesRuleArrayOutput)
}

func (o GetAclRulesResultOutput) TimeOrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAclRulesResult) *string { return v.TimeOrderBy }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o GetAclRulesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetAclRulesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAclRulesResultOutput{})
}
