// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package private_zone

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of private zone records
func Records(ctx *pulumi.Context, args *RecordsArgs, opts ...pulumi.InvokeOption) (*RecordsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv RecordsResult
	err := ctx.Invoke("volcengine:private_zone/records:Records", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Records.
type RecordsArgs struct {
	// The host of Private Zone Record.
	Host *string `pulumi:"host"`
	// The last operator account id of Private Zone Record.
	LastOperator *string `pulumi:"lastOperator"`
	// The subnet id of Private Zone Record. This field is only effected when the `intelligentMode` of the private zone is true.
	Line *string `pulumi:"line"`
	// The domain name of Private Zone Record.
	Name *string `pulumi:"name"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The id of Private Zone Record.
	RecordId *string `pulumi:"recordId"`
	// The search mode of query `host`. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
	SearchMode *string `pulumi:"searchMode"`
	// The type of Private Zone Record.
	Type *string `pulumi:"type"`
	// The value of Private Zone Record.
	Value *string `pulumi:"value"`
	// The zid of Private Zone.
	Zid *int `pulumi:"zid"`
}

// A collection of values returned by Records.
type RecordsResult struct {
	// The host of the private zone record.
	Host *string `pulumi:"host"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The last operator account id of the private zone record.
	LastOperator *string `pulumi:"lastOperator"`
	// The subnet id of the private zone record. This field is only effected when the `intelligentMode` of the private zone is true.
	Line       *string `pulumi:"line"`
	Name       *string `pulumi:"name"`
	OutputFile *string `pulumi:"outputFile"`
	// The id of the private zone record.
	RecordId *string `pulumi:"recordId"`
	// The collection of query.
	Records    []RecordsRecord `pulumi:"records"`
	SearchMode *string         `pulumi:"searchMode"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
	// The type of the private zone record.
	Type *string `pulumi:"type"`
	// The value of the private zone record.
	Value *string `pulumi:"value"`
	// The zid of the private zone record.
	Zid *int `pulumi:"zid"`
}

func RecordsOutput(ctx *pulumi.Context, args RecordsOutputArgs, opts ...pulumi.InvokeOption) RecordsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (RecordsResult, error) {
			args := v.(RecordsArgs)
			r, err := Records(ctx, &args, opts...)
			var s RecordsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(RecordsResultOutput)
}

// A collection of arguments for invoking Records.
type RecordsOutputArgs struct {
	// The host of Private Zone Record.
	Host pulumi.StringPtrInput `pulumi:"host"`
	// The last operator account id of Private Zone Record.
	LastOperator pulumi.StringPtrInput `pulumi:"lastOperator"`
	// The subnet id of Private Zone Record. This field is only effected when the `intelligentMode` of the private zone is true.
	Line pulumi.StringPtrInput `pulumi:"line"`
	// The domain name of Private Zone Record.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The id of Private Zone Record.
	RecordId pulumi.StringPtrInput `pulumi:"recordId"`
	// The search mode of query `host`. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
	SearchMode pulumi.StringPtrInput `pulumi:"searchMode"`
	// The type of Private Zone Record.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// The value of Private Zone Record.
	Value pulumi.StringPtrInput `pulumi:"value"`
	// The zid of Private Zone.
	Zid pulumi.IntPtrInput `pulumi:"zid"`
}

func (RecordsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordsArgs)(nil)).Elem()
}

// A collection of values returned by Records.
type RecordsResultOutput struct{ *pulumi.OutputState }

func (RecordsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordsResult)(nil)).Elem()
}

func (o RecordsResultOutput) ToRecordsResultOutput() RecordsResultOutput {
	return o
}

func (o RecordsResultOutput) ToRecordsResultOutputWithContext(ctx context.Context) RecordsResultOutput {
	return o
}

// The host of the private zone record.
func (o RecordsResultOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecordsResult) *string { return v.Host }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o RecordsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RecordsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The last operator account id of the private zone record.
func (o RecordsResultOutput) LastOperator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecordsResult) *string { return v.LastOperator }).(pulumi.StringPtrOutput)
}

// The subnet id of the private zone record. This field is only effected when the `intelligentMode` of the private zone is true.
func (o RecordsResultOutput) Line() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecordsResult) *string { return v.Line }).(pulumi.StringPtrOutput)
}

func (o RecordsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecordsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RecordsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecordsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The id of the private zone record.
func (o RecordsResultOutput) RecordId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecordsResult) *string { return v.RecordId }).(pulumi.StringPtrOutput)
}

// The collection of query.
func (o RecordsResultOutput) Records() RecordsRecordArrayOutput {
	return o.ApplyT(func(v RecordsResult) []RecordsRecord { return v.Records }).(RecordsRecordArrayOutput)
}

func (o RecordsResultOutput) SearchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecordsResult) *string { return v.SearchMode }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o RecordsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v RecordsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The type of the private zone record.
func (o RecordsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecordsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// The value of the private zone record.
func (o RecordsResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecordsResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

// The zid of the private zone record.
func (o RecordsResultOutput) Zid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RecordsResult) *int { return v.Zid }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RecordsResultOutput{})
}