// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of traffic mirror filter rules
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpc.GetTrafficMirrorFilterRules(ctx, &vpc.GetTrafficMirrorFilterRulesArgs{
//				TrafficMirrorFilterIds: []string{
//					"tmf-mivro9v5x24g5smt1bsq****",
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
//
// Deprecated: volcengine.vpc.TrafficMirrorFilterRules has been deprecated in favor of volcengine.vpc.getTrafficMirrorFilterRules
func TrafficMirrorFilterRules(ctx *pulumi.Context, args *TrafficMirrorFilterRulesArgs, opts ...pulumi.InvokeOption) (*TrafficMirrorFilterRulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv TrafficMirrorFilterRulesResult
	err := ctx.Invoke("volcengine:vpc/trafficMirrorFilterRules:TrafficMirrorFilterRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking TrafficMirrorFilterRules.
type TrafficMirrorFilterRulesArgs struct {
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The project name of traffic mirror filter.
	ProjectName *string `pulumi:"projectName"`
	// Tags.
	Tags []TrafficMirrorFilterRulesTag `pulumi:"tags"`
	// A list of traffic mirror filter IDs.
	TrafficMirrorFilterIds []string `pulumi:"trafficMirrorFilterIds"`
	// A list of traffic mirror filter names.
	TrafficMirrorFilterNames []string `pulumi:"trafficMirrorFilterNames"`
}

// A collection of values returned by TrafficMirrorFilterRules.
type TrafficMirrorFilterRulesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id          string                        `pulumi:"id"`
	OutputFile  *string                       `pulumi:"outputFile"`
	ProjectName *string                       `pulumi:"projectName"`
	Tags        []TrafficMirrorFilterRulesTag `pulumi:"tags"`
	// The total count of query.
	TotalCount               int      `pulumi:"totalCount"`
	TrafficMirrorFilterIds   []string `pulumi:"trafficMirrorFilterIds"`
	TrafficMirrorFilterNames []string `pulumi:"trafficMirrorFilterNames"`
	// The collection of query.
	TrafficMirrorFilterRules []TrafficMirrorFilterRulesTrafficMirrorFilterRule `pulumi:"trafficMirrorFilterRules"`
}

func TrafficMirrorFilterRulesOutput(ctx *pulumi.Context, args TrafficMirrorFilterRulesOutputArgs, opts ...pulumi.InvokeOption) TrafficMirrorFilterRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (TrafficMirrorFilterRulesResult, error) {
			args := v.(TrafficMirrorFilterRulesArgs)
			r, err := TrafficMirrorFilterRules(ctx, &args, opts...)
			var s TrafficMirrorFilterRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(TrafficMirrorFilterRulesResultOutput)
}

// A collection of arguments for invoking TrafficMirrorFilterRules.
type TrafficMirrorFilterRulesOutputArgs struct {
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The project name of traffic mirror filter.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
	// Tags.
	Tags TrafficMirrorFilterRulesTagArrayInput `pulumi:"tags"`
	// A list of traffic mirror filter IDs.
	TrafficMirrorFilterIds pulumi.StringArrayInput `pulumi:"trafficMirrorFilterIds"`
	// A list of traffic mirror filter names.
	TrafficMirrorFilterNames pulumi.StringArrayInput `pulumi:"trafficMirrorFilterNames"`
}

func (TrafficMirrorFilterRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficMirrorFilterRulesArgs)(nil)).Elem()
}

// A collection of values returned by TrafficMirrorFilterRules.
type TrafficMirrorFilterRulesResultOutput struct{ *pulumi.OutputState }

func (TrafficMirrorFilterRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficMirrorFilterRulesResult)(nil)).Elem()
}

func (o TrafficMirrorFilterRulesResultOutput) ToTrafficMirrorFilterRulesResultOutput() TrafficMirrorFilterRulesResultOutput {
	return o
}

func (o TrafficMirrorFilterRulesResultOutput) ToTrafficMirrorFilterRulesResultOutputWithContext(ctx context.Context) TrafficMirrorFilterRulesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o TrafficMirrorFilterRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v TrafficMirrorFilterRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o TrafficMirrorFilterRulesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrafficMirrorFilterRulesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o TrafficMirrorFilterRulesResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrafficMirrorFilterRulesResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

func (o TrafficMirrorFilterRulesResultOutput) Tags() TrafficMirrorFilterRulesTagArrayOutput {
	return o.ApplyT(func(v TrafficMirrorFilterRulesResult) []TrafficMirrorFilterRulesTag { return v.Tags }).(TrafficMirrorFilterRulesTagArrayOutput)
}

// The total count of query.
func (o TrafficMirrorFilterRulesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v TrafficMirrorFilterRulesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func (o TrafficMirrorFilterRulesResultOutput) TrafficMirrorFilterIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TrafficMirrorFilterRulesResult) []string { return v.TrafficMirrorFilterIds }).(pulumi.StringArrayOutput)
}

func (o TrafficMirrorFilterRulesResultOutput) TrafficMirrorFilterNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TrafficMirrorFilterRulesResult) []string { return v.TrafficMirrorFilterNames }).(pulumi.StringArrayOutput)
}

// The collection of query.
func (o TrafficMirrorFilterRulesResultOutput) TrafficMirrorFilterRules() TrafficMirrorFilterRulesTrafficMirrorFilterRuleArrayOutput {
	return o.ApplyT(func(v TrafficMirrorFilterRulesResult) []TrafficMirrorFilterRulesTrafficMirrorFilterRule {
		return v.TrafficMirrorFilterRules
	}).(TrafficMirrorFilterRulesTrafficMirrorFilterRuleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(TrafficMirrorFilterRulesResultOutput{})
}
