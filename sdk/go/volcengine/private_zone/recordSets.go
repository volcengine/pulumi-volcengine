// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package private_zone

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of private zone record sets
func RecordSets(ctx *pulumi.Context, args *RecordSetsArgs, opts ...pulumi.InvokeOption) (*RecordSetsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv RecordSetsResult
	err := ctx.Invoke("volcengine:private_zone/recordSets:RecordSets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking RecordSets.
type RecordSetsArgs struct {
	// The host of Private Zone Record Set.
	Host *string `pulumi:"host"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The id of Private Zone Record Set.
	RecordSetId *string `pulumi:"recordSetId"`
	// The search mode of query `host`. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
	SearchMode *string `pulumi:"searchMode"`
	// The zid of Private Zone.
	Zid int `pulumi:"zid"`
}

// A collection of values returned by RecordSets.
type RecordSetsResult struct {
	// The host of the private zone record.
	Host *string `pulumi:"host"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	OutputFile *string `pulumi:"outputFile"`
	// The id of the private zone record set.
	RecordSetId *string `pulumi:"recordSetId"`
	// The collection of query.
	RecordSets []RecordSetsRecordSet `pulumi:"recordSets"`
	SearchMode *string               `pulumi:"searchMode"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
	Zid        int `pulumi:"zid"`
}

func RecordSetsOutput(ctx *pulumi.Context, args RecordSetsOutputArgs, opts ...pulumi.InvokeOption) RecordSetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (RecordSetsResult, error) {
			args := v.(RecordSetsArgs)
			r, err := RecordSets(ctx, &args, opts...)
			var s RecordSetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(RecordSetsResultOutput)
}

// A collection of arguments for invoking RecordSets.
type RecordSetsOutputArgs struct {
	// The host of Private Zone Record Set.
	Host pulumi.StringPtrInput `pulumi:"host"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The id of Private Zone Record Set.
	RecordSetId pulumi.StringPtrInput `pulumi:"recordSetId"`
	// The search mode of query `host`. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
	SearchMode pulumi.StringPtrInput `pulumi:"searchMode"`
	// The zid of Private Zone.
	Zid pulumi.IntInput `pulumi:"zid"`
}

func (RecordSetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordSetsArgs)(nil)).Elem()
}

// A collection of values returned by RecordSets.
type RecordSetsResultOutput struct{ *pulumi.OutputState }

func (RecordSetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordSetsResult)(nil)).Elem()
}

func (o RecordSetsResultOutput) ToRecordSetsResultOutput() RecordSetsResultOutput {
	return o
}

func (o RecordSetsResultOutput) ToRecordSetsResultOutputWithContext(ctx context.Context) RecordSetsResultOutput {
	return o
}

// The host of the private zone record.
func (o RecordSetsResultOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecordSetsResult) *string { return v.Host }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o RecordSetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RecordSetsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o RecordSetsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecordSetsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The id of the private zone record set.
func (o RecordSetsResultOutput) RecordSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecordSetsResult) *string { return v.RecordSetId }).(pulumi.StringPtrOutput)
}

// The collection of query.
func (o RecordSetsResultOutput) RecordSets() RecordSetsRecordSetArrayOutput {
	return o.ApplyT(func(v RecordSetsResult) []RecordSetsRecordSet { return v.RecordSets }).(RecordSetsRecordSetArrayOutput)
}

func (o RecordSetsResultOutput) SearchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecordSetsResult) *string { return v.SearchMode }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o RecordSetsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v RecordSetsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func (o RecordSetsResultOutput) Zid() pulumi.IntOutput {
	return o.ApplyT(func(v RecordSetsResult) int { return v.Zid }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(RecordSetsResultOutput{})
}