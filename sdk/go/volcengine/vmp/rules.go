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
//
// Deprecated: volcengine.vmp.Rules has been deprecated in favor of volcengine.vmp.getRules
func Rules(ctx *pulumi.Context, args *RulesArgs, opts ...pulumi.InvokeOption) (*RulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv RulesResult
	err := ctx.Invoke("volcengine:vmp/rules:Rules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Rules.
type RulesArgs struct {
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

// A collection of values returned by Rules.
type RulesResult struct {
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
	Rules []RulesRule `pulumi:"rules"`
	// The status of rule.
	Status *string `pulumi:"status"`
	// The total count of query.
	TotalCount  int    `pulumi:"totalCount"`
	WorkspaceId string `pulumi:"workspaceId"`
}

func RulesOutput(ctx *pulumi.Context, args RulesOutputArgs, opts ...pulumi.InvokeOption) RulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (RulesResult, error) {
			args := v.(RulesArgs)
			r, err := Rules(ctx, &args, opts...)
			var s RulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(RulesResultOutput)
}

// A collection of arguments for invoking Rules.
type RulesOutputArgs struct {
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

func (RulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesArgs)(nil)).Elem()
}

// A collection of values returned by Rules.
type RulesResultOutput struct{ *pulumi.OutputState }

func (RulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesResult)(nil)).Elem()
}

func (o RulesResultOutput) ToRulesResultOutput() RulesResultOutput {
	return o
}

func (o RulesResultOutput) ToRulesResultOutputWithContext(ctx context.Context) RulesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o RulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RulesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of rule.
func (o RulesResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v RulesResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of rule.
func (o RulesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RulesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o RulesResultOutput) RuleFileNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RulesResult) []string { return v.RuleFileNames }).(pulumi.StringArrayOutput)
}

func (o RulesResultOutput) RuleGroupNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RulesResult) []string { return v.RuleGroupNames }).(pulumi.StringArrayOutput)
}

// The collection of query.
func (o RulesResultOutput) Rules() RulesRuleArrayOutput {
	return o.ApplyT(func(v RulesResult) []RulesRule { return v.Rules }).(RulesRuleArrayOutput)
}

// The status of rule.
func (o RulesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o RulesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v RulesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func (o RulesResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v RulesResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RulesResultOutput{})
}
