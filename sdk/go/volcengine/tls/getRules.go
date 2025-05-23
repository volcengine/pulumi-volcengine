// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of tls rules
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
//			_, err := tls.GetRules(ctx, &tls.GetRulesArgs{
//				ProjectId: "cc44f8b6-0328-4622-b043-023fca735cd4",
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
	err := ctx.Invoke("volcengine:tls/getRules:getRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRules.
type GetRulesArgs struct {
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The project id.
	ProjectId string `pulumi:"projectId"`
	// The rule id.
	RuleId *string `pulumi:"ruleId"`
	// The rule name.
	RuleName *string `pulumi:"ruleName"`
	// The topic id.
	TopicId *string `pulumi:"topicId"`
	// The topic name.
	TopicName *string `pulumi:"topicName"`
}

// A collection of values returned by getRules.
type GetRulesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	OutputFile *string `pulumi:"outputFile"`
	ProjectId  string  `pulumi:"projectId"`
	// The rule id.
	RuleId *string `pulumi:"ruleId"`
	// The rule name.
	RuleName *string `pulumi:"ruleName"`
	// The rules list.
	Rules []GetRulesRule `pulumi:"rules"`
	// The topic id.
	TopicId *string `pulumi:"topicId"`
	// The topic name.
	TopicName *string `pulumi:"topicName"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
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
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The project id.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// The rule id.
	RuleId pulumi.StringPtrInput `pulumi:"ruleId"`
	// The rule name.
	RuleName pulumi.StringPtrInput `pulumi:"ruleName"`
	// The topic id.
	TopicId pulumi.StringPtrInput `pulumi:"topicId"`
	// The topic name.
	TopicName pulumi.StringPtrInput `pulumi:"topicName"`
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

func (o GetRulesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRulesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetRulesResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRulesResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The rule id.
func (o GetRulesResultOutput) RuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRulesResult) *string { return v.RuleId }).(pulumi.StringPtrOutput)
}

// The rule name.
func (o GetRulesResultOutput) RuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRulesResult) *string { return v.RuleName }).(pulumi.StringPtrOutput)
}

// The rules list.
func (o GetRulesResultOutput) Rules() GetRulesRuleArrayOutput {
	return o.ApplyT(func(v GetRulesResult) []GetRulesRule { return v.Rules }).(GetRulesRuleArrayOutput)
}

// The topic id.
func (o GetRulesResultOutput) TopicId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRulesResult) *string { return v.TopicId }).(pulumi.StringPtrOutput)
}

// The topic name.
func (o GetRulesResultOutput) TopicName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRulesResult) *string { return v.TopicName }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o GetRulesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetRulesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRulesResultOutput{})
}
