// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of tls schedule sql tasks
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
//			_, err := tls.GetScheduleSqlTasks(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: volcengine.tls.ScheduleSqlTasks has been deprecated in favor of volcengine.tls.getScheduleSqlTasks
func ScheduleSqlTasks(ctx *pulumi.Context, args *ScheduleSqlTasksArgs, opts ...pulumi.InvokeOption) (*ScheduleSqlTasksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ScheduleSqlTasksResult
	err := ctx.Invoke("volcengine:tls/scheduleSqlTasks:ScheduleSqlTasks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking ScheduleSqlTasks.
type ScheduleSqlTasksArgs struct {
	// IAM log project name.
	IamProjectName *string `pulumi:"iamProjectName"`
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The log project ID to which the source log topic belongs.
	ProjectId *string `pulumi:"projectId"`
	// The name of the log item to which the source log topic belongs.
	ProjectName *string `pulumi:"projectName"`
	// Source log topic name.
	SourceTopicName *string `pulumi:"sourceTopicName"`
	// Timed SQL analysis task status.
	Status *string `pulumi:"status"`
	// Timed SQL analysis task ID.
	TaskId *string `pulumi:"taskId"`
	// Timed SQL analysis task name.
	TaskName *string `pulumi:"taskName"`
	// Source log topic ID.
	TopicId *string `pulumi:"topicId"`
}

// A collection of values returned by ScheduleSqlTasks.
type ScheduleSqlTasksResult struct {
	IamProjectName *string `pulumi:"iamProjectName"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	NameRegex   *string `pulumi:"nameRegex"`
	OutputFile  *string `pulumi:"outputFile"`
	ProjectId   *string `pulumi:"projectId"`
	ProjectName *string `pulumi:"projectName"`
	// The name of the source log topic where the original log for timed SQL analysis is located.
	SourceTopicName *string `pulumi:"sourceTopicName"`
	// Whether to start the scheduled SQL analysis task immediately after completing the task configuration.
	Status *string `pulumi:"status"`
	// Timed SQL analysis task ID.
	TaskId *string `pulumi:"taskId"`
	// Timed SQL analysis task name.
	TaskName *string `pulumi:"taskName"`
	// The List of timed SQL analysis tasks.
	Tasks   []ScheduleSqlTasksTask `pulumi:"tasks"`
	TopicId *string                `pulumi:"topicId"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func ScheduleSqlTasksOutput(ctx *pulumi.Context, args ScheduleSqlTasksOutputArgs, opts ...pulumi.InvokeOption) ScheduleSqlTasksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ScheduleSqlTasksResult, error) {
			args := v.(ScheduleSqlTasksArgs)
			r, err := ScheduleSqlTasks(ctx, &args, opts...)
			var s ScheduleSqlTasksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ScheduleSqlTasksResultOutput)
}

// A collection of arguments for invoking ScheduleSqlTasks.
type ScheduleSqlTasksOutputArgs struct {
	// IAM log project name.
	IamProjectName pulumi.StringPtrInput `pulumi:"iamProjectName"`
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The log project ID to which the source log topic belongs.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The name of the log item to which the source log topic belongs.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
	// Source log topic name.
	SourceTopicName pulumi.StringPtrInput `pulumi:"sourceTopicName"`
	// Timed SQL analysis task status.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// Timed SQL analysis task ID.
	TaskId pulumi.StringPtrInput `pulumi:"taskId"`
	// Timed SQL analysis task name.
	TaskName pulumi.StringPtrInput `pulumi:"taskName"`
	// Source log topic ID.
	TopicId pulumi.StringPtrInput `pulumi:"topicId"`
}

func (ScheduleSqlTasksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleSqlTasksArgs)(nil)).Elem()
}

// A collection of values returned by ScheduleSqlTasks.
type ScheduleSqlTasksResultOutput struct{ *pulumi.OutputState }

func (ScheduleSqlTasksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleSqlTasksResult)(nil)).Elem()
}

func (o ScheduleSqlTasksResultOutput) ToScheduleSqlTasksResultOutput() ScheduleSqlTasksResultOutput {
	return o
}

func (o ScheduleSqlTasksResultOutput) ToScheduleSqlTasksResultOutputWithContext(ctx context.Context) ScheduleSqlTasksResultOutput {
	return o
}

func (o ScheduleSqlTasksResultOutput) IamProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleSqlTasksResult) *string { return v.IamProjectName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o ScheduleSqlTasksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleSqlTasksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ScheduleSqlTasksResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleSqlTasksResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o ScheduleSqlTasksResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleSqlTasksResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o ScheduleSqlTasksResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleSqlTasksResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o ScheduleSqlTasksResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleSqlTasksResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// The name of the source log topic where the original log for timed SQL analysis is located.
func (o ScheduleSqlTasksResultOutput) SourceTopicName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleSqlTasksResult) *string { return v.SourceTopicName }).(pulumi.StringPtrOutput)
}

// Whether to start the scheduled SQL analysis task immediately after completing the task configuration.
func (o ScheduleSqlTasksResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleSqlTasksResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// Timed SQL analysis task ID.
func (o ScheduleSqlTasksResultOutput) TaskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleSqlTasksResult) *string { return v.TaskId }).(pulumi.StringPtrOutput)
}

// Timed SQL analysis task name.
func (o ScheduleSqlTasksResultOutput) TaskName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleSqlTasksResult) *string { return v.TaskName }).(pulumi.StringPtrOutput)
}

// The List of timed SQL analysis tasks.
func (o ScheduleSqlTasksResultOutput) Tasks() ScheduleSqlTasksTaskArrayOutput {
	return o.ApplyT(func(v ScheduleSqlTasksResult) []ScheduleSqlTasksTask { return v.Tasks }).(ScheduleSqlTasksTaskArrayOutput)
}

func (o ScheduleSqlTasksResultOutput) TopicId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleSqlTasksResult) *string { return v.TopicId }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o ScheduleSqlTasksResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v ScheduleSqlTasksResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduleSqlTasksResultOutput{})
}
