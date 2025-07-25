// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vmp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of vmp rules
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vmp"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vmp.GetRules(ctx, &vmp.GetRulesArgs{
//				Kind:        "Recording",
//				WorkspaceId: "baa02ffb-6f22-43c4-841b-ecf90ded****",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetRules(ctx *pulumi.Context, args *GetRulesArgs, opts ...pulumi.InvokeOption) (*GetRulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRulesResult
	err := ctx.Invoke("volcengine:vmp/getRules:getRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRules.
type GetRulesArgs struct {
	// The kind of rule.
	Kind string `pulumi:"kind"`
	// The name of rule.
	Name *string `pulumi:"name"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The name of rule file.
	RuleFileNames []string `pulumi:"ruleFileNames"`
	// The name of rule group.
	RuleGroupNames []string `pulumi:"ruleGroupNames"`
	// The status of rule.
	Status *string `pulumi:"status"`
	// The id of workspace.
	WorkspaceId string `pulumi:"workspaceId"`
}

// A collection of values returned by getRules.
type GetRulesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The kind of rule.
	Kind string `pulumi:"kind"`
	// The name of rule.
	Name           *string  `pulumi:"name"`
	OutputFile     *string  `pulumi:"outputFile"`
	RuleFileNames  []string `pulumi:"ruleFileNames"`
	RuleGroupNames []string `pulumi:"ruleGroupNames"`
	// The collection of query.
	Rules []GetRulesRule `pulumi:"rules"`
	// The status of rule.
	Status *string `pulumi:"status"`
	// The total count of query.
	TotalCount  int    `pulumi:"totalCount"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func GetRulesOutput(ctx *pulumi.Context, args GetRulesOutputArgs, opts ...pulumi.InvokeOption) GetRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRulesResult, error) {
			args := v.(GetRulesArgs)
			r, err := GetRules(ctx, &args, opts...)
			var s GetRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRulesResultOutput)
}

// A collection of arguments for invoking getRules.
type GetRulesOutputArgs struct {
	// The kind of rule.
	Kind pulumi.StringInput `pulumi:"kind"`
	// The name of rule.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The name of rule file.
	RuleFileNames pulumi.StringArrayInput `pulumi:"ruleFileNames"`
	// The name of rule group.
	RuleGroupNames pulumi.StringArrayInput `pulumi:"ruleGroupNames"`
	// The status of rule.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The id of workspace.
	WorkspaceId pulumi.StringInput `pulumi:"workspaceId"`
}

func (GetRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRulesArgs)(nil)).Elem()
}

// A collection of values returned by getRules.
type GetRulesResultOutput struct{ *pulumi.OutputState }

func (GetRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRulesResult)(nil)).Elem()
}

func (o GetRulesResultOutput) ToGetRulesResultOutput() GetRulesResultOutput {
	return o
}

func (o GetRulesResultOutput) ToGetRulesResultOutputWithContext(ctx context.Context) GetRulesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of rule.
func (o GetRulesResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of rule.
func (o GetRulesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRulesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetRulesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRulesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetRulesResultOutput) RuleFileNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRulesResult) []string { return v.RuleFileNames }).(pulumi.StringArrayOutput)
}

func (o GetRulesResultOutput) RuleGroupNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRulesResult) []string { return v.RuleGroupNames }).(pulumi.StringArrayOutput)
}

// The collection of query.
func (o GetRulesResultOutput) Rules() GetRulesRuleArrayOutput {
	return o.ApplyT(func(v GetRulesResult) []GetRulesRule { return v.Rules }).(GetRulesRuleArrayOutput)
}

// The status of rule.
func (o GetRulesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRulesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o GetRulesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetRulesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func (o GetRulesResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRulesResultOutput{})
}
