// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of alb rules
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/alb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := alb.Rules(ctx, &alb.RulesArgs{
//				ListenerId: "lsn-1iidd19u4oni874adhezjkyj3",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func Rules(ctx *pulumi.Context, args *RulesArgs, opts ...pulumi.InvokeOption) (*RulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv RulesResult
	err := ctx.Invoke("volcengine:alb/rules:Rules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Rules.
type RulesArgs struct {
	// The Id of listener.
	ListenerId string `pulumi:"listenerId"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by Rules.
type RulesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	ListenerId string  `pulumi:"listenerId"`
	OutputFile *string `pulumi:"outputFile"`
	// The collection of Rule query.
	Rules []RulesRule `pulumi:"rules"`
	// The total count of Rule query.
	TotalCount int `pulumi:"totalCount"`
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
	// The Id of listener.
	ListenerId pulumi.StringInput `pulumi:"listenerId"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
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

func (o RulesResultOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v RulesResult) string { return v.ListenerId }).(pulumi.StringOutput)
}

func (o RulesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The collection of Rule query.
func (o RulesResultOutput) Rules() RulesRuleArrayOutput {
	return o.ApplyT(func(v RulesResult) []RulesRule { return v.Rules }).(RulesRuleArrayOutput)
}

// The total count of Rule query.
func (o RulesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v RulesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(RulesResultOutput{})
}