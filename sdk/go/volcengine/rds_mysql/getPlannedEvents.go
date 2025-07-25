// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds_mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of rds mysql planned events
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/rds_mysql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rds_mysql.GetPlannedEvents(ctx, &rds_mysql.GetPlannedEventsArgs{
//				InstanceId: pulumi.StringRef("mysql-b51d37110dd1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetPlannedEvents(ctx *pulumi.Context, args *GetPlannedEventsArgs, opts ...pulumi.InvokeOption) (*GetPlannedEventsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPlannedEventsResult
	err := ctx.Invoke("volcengine:rds_mysql/getPlannedEvents:getPlannedEvents", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPlannedEvents.
type GetPlannedEventsArgs struct {
	// The start time of the planned event.
	BeginTime *string `pulumi:"beginTime"`
	// The end time of the planned event.
	EndTime *string `pulumi:"endTime"`
	// The id of the planned event.
	EventId *string `pulumi:"eventId"`
	// The type of the planned event.
	EventTypes []string `pulumi:"eventTypes"`
	// The id of the instance.
	InstanceId *string `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The status of the planned event.
	Statuses []string `pulumi:"statuses"`
}

// A collection of values returned by getPlannedEvents.
type GetPlannedEventsResult struct {
	BeginTime *string `pulumi:"beginTime"`
	EndTime   *string `pulumi:"endTime"`
	// The id of the planned event.
	EventId *string `pulumi:"eventId"`
	// The type of the planned event.
	EventTypes []string `pulumi:"eventTypes"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The id of the instance.
	InstanceId *string `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
	// The collection of query.
	PlannedEvents []GetPlannedEventsPlannedEvent `pulumi:"plannedEvents"`
	// Event status.
	Statuses []string `pulumi:"statuses"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func GetPlannedEventsOutput(ctx *pulumi.Context, args GetPlannedEventsOutputArgs, opts ...pulumi.InvokeOption) GetPlannedEventsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPlannedEventsResult, error) {
			args := v.(GetPlannedEventsArgs)
			r, err := GetPlannedEvents(ctx, &args, opts...)
			var s GetPlannedEventsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPlannedEventsResultOutput)
}

// A collection of arguments for invoking getPlannedEvents.
type GetPlannedEventsOutputArgs struct {
	// The start time of the planned event.
	BeginTime pulumi.StringPtrInput `pulumi:"beginTime"`
	// The end time of the planned event.
	EndTime pulumi.StringPtrInput `pulumi:"endTime"`
	// The id of the planned event.
	EventId pulumi.StringPtrInput `pulumi:"eventId"`
	// The type of the planned event.
	EventTypes pulumi.StringArrayInput `pulumi:"eventTypes"`
	// The id of the instance.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The status of the planned event.
	Statuses pulumi.StringArrayInput `pulumi:"statuses"`
}

func (GetPlannedEventsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPlannedEventsArgs)(nil)).Elem()
}

// A collection of values returned by getPlannedEvents.
type GetPlannedEventsResultOutput struct{ *pulumi.OutputState }

func (GetPlannedEventsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPlannedEventsResult)(nil)).Elem()
}

func (o GetPlannedEventsResultOutput) ToGetPlannedEventsResultOutput() GetPlannedEventsResultOutput {
	return o
}

func (o GetPlannedEventsResultOutput) ToGetPlannedEventsResultOutputWithContext(ctx context.Context) GetPlannedEventsResultOutput {
	return o
}

func (o GetPlannedEventsResultOutput) BeginTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPlannedEventsResult) *string { return v.BeginTime }).(pulumi.StringPtrOutput)
}

func (o GetPlannedEventsResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPlannedEventsResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

// The id of the planned event.
func (o GetPlannedEventsResultOutput) EventId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPlannedEventsResult) *string { return v.EventId }).(pulumi.StringPtrOutput)
}

// The type of the planned event.
func (o GetPlannedEventsResultOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPlannedEventsResult) []string { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPlannedEventsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPlannedEventsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The id of the instance.
func (o GetPlannedEventsResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPlannedEventsResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o GetPlannedEventsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPlannedEventsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The collection of query.
func (o GetPlannedEventsResultOutput) PlannedEvents() GetPlannedEventsPlannedEventArrayOutput {
	return o.ApplyT(func(v GetPlannedEventsResult) []GetPlannedEventsPlannedEvent { return v.PlannedEvents }).(GetPlannedEventsPlannedEventArrayOutput)
}

// Event status.
func (o GetPlannedEventsResultOutput) Statuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPlannedEventsResult) []string { return v.Statuses }).(pulumi.StringArrayOutput)
}

// The total count of query.
func (o GetPlannedEventsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetPlannedEventsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPlannedEventsResultOutput{})
}
