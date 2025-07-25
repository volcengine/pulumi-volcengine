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
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/private_zone"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := private_zone.GetRecords(ctx, &private_zone.GetRecordsArgs{
//				RecordId: pulumi.StringRef("907925684878276****"),
//				Zid:      pulumi.IntRef(2450000),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetRecords(ctx *pulumi.Context, args *GetRecordsArgs, opts ...pulumi.InvokeOption) (*GetRecordsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRecordsResult
	err := ctx.Invoke("volcengine:private_zone/getRecords:getRecords", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRecords.
type GetRecordsArgs struct {
	// The host of Private Zone Record.
	Host *string `pulumi:"host"`
	// The last operator account id of Private Zone Record.
	LastOperator *string `pulumi:"lastOperator"`
	// The subnet id of Private Zone Record.
	Line *string `pulumi:"line"`
	// The domain name of Private Zone Record.
	Name *string `pulumi:"name"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// This field is deprecated, please use `recordIds` instead. The id of Private Zone Record.
	//
	// Deprecated: This field is deprecated, please use `recordIds` instead.
	RecordId *string `pulumi:"recordId"`
	// The ids of Private Zone Record.
	RecordIds []string `pulumi:"recordIds"`
	// The search mode of query `host`. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
	SearchMode *string `pulumi:"searchMode"`
	// The type of Private Zone Record.
	Type *string `pulumi:"type"`
	// The value of Private Zone Record.
	Value *string `pulumi:"value"`
	// The zid of Private Zone.
	Zid *int `pulumi:"zid"`
}

// A collection of values returned by getRecords.
type GetRecordsResult struct {
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
	//
	// Deprecated: This field is deprecated, please use `recordIds` instead.
	RecordId  *string  `pulumi:"recordId"`
	RecordIds []string `pulumi:"recordIds"`
	// The collection of query.
	Records    []GetRecordsRecord `pulumi:"records"`
	SearchMode *string            `pulumi:"searchMode"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
	// The type of the private zone record.
	Type *string `pulumi:"type"`
	// The value of the private zone record.
	Value *string `pulumi:"value"`
	// The zid of the private zone record.
	Zid *int `pulumi:"zid"`
}

func GetRecordsOutput(ctx *pulumi.Context, args GetRecordsOutputArgs, opts ...pulumi.InvokeOption) GetRecordsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRecordsResult, error) {
			args := v.(GetRecordsArgs)
			r, err := GetRecords(ctx, &args, opts...)
			var s GetRecordsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRecordsResultOutput)
}

// A collection of arguments for invoking getRecords.
type GetRecordsOutputArgs struct {
	// The host of Private Zone Record.
	Host pulumi.StringPtrInput `pulumi:"host"`
	// The last operator account id of Private Zone Record.
	LastOperator pulumi.StringPtrInput `pulumi:"lastOperator"`
	// The subnet id of Private Zone Record.
	Line pulumi.StringPtrInput `pulumi:"line"`
	// The domain name of Private Zone Record.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// This field is deprecated, please use `recordIds` instead. The id of Private Zone Record.
	//
	// Deprecated: This field is deprecated, please use `recordIds` instead.
	RecordId pulumi.StringPtrInput `pulumi:"recordId"`
	// The ids of Private Zone Record.
	RecordIds pulumi.StringArrayInput `pulumi:"recordIds"`
	// The search mode of query `host`. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
	SearchMode pulumi.StringPtrInput `pulumi:"searchMode"`
	// The type of Private Zone Record.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// The value of Private Zone Record.
	Value pulumi.StringPtrInput `pulumi:"value"`
	// The zid of Private Zone.
	Zid pulumi.IntPtrInput `pulumi:"zid"`
}

func (GetRecordsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRecordsArgs)(nil)).Elem()
}

// A collection of values returned by getRecords.
type GetRecordsResultOutput struct{ *pulumi.OutputState }

func (GetRecordsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRecordsResult)(nil)).Elem()
}

func (o GetRecordsResultOutput) ToGetRecordsResultOutput() GetRecordsResultOutput {
	return o
}

func (o GetRecordsResultOutput) ToGetRecordsResultOutputWithContext(ctx context.Context) GetRecordsResultOutput {
	return o
}

// The host of the private zone record.
func (o GetRecordsResultOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecordsResult) *string { return v.Host }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRecordsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRecordsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The last operator account id of the private zone record.
func (o GetRecordsResultOutput) LastOperator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecordsResult) *string { return v.LastOperator }).(pulumi.StringPtrOutput)
}

// The subnet id of the private zone record. This field is only effected when the `intelligentMode` of the private zone is true.
func (o GetRecordsResultOutput) Line() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecordsResult) *string { return v.Line }).(pulumi.StringPtrOutput)
}

func (o GetRecordsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecordsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetRecordsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecordsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The id of the private zone record.
//
// Deprecated: This field is deprecated, please use `recordIds` instead.
func (o GetRecordsResultOutput) RecordId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecordsResult) *string { return v.RecordId }).(pulumi.StringPtrOutput)
}

func (o GetRecordsResultOutput) RecordIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRecordsResult) []string { return v.RecordIds }).(pulumi.StringArrayOutput)
}

// The collection of query.
func (o GetRecordsResultOutput) Records() GetRecordsRecordArrayOutput {
	return o.ApplyT(func(v GetRecordsResult) []GetRecordsRecord { return v.Records }).(GetRecordsRecordArrayOutput)
}

func (o GetRecordsResultOutput) SearchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecordsResult) *string { return v.SearchMode }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o GetRecordsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetRecordsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The type of the private zone record.
func (o GetRecordsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecordsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// The value of the private zone record.
func (o GetRecordsResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecordsResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

// The zid of the private zone record.
func (o GetRecordsResultOutput) Zid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetRecordsResult) *int { return v.Zid }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRecordsResultOutput{})
}
