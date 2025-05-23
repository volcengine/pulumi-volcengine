// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of scaling lifecycle hooks
func GetScalingLifecycleHooks(ctx *pulumi.Context, args *GetScalingLifecycleHooksArgs, opts ...pulumi.InvokeOption) (*GetScalingLifecycleHooksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetScalingLifecycleHooksResult
	err := ctx.Invoke("volcengine:autoscaling/getScalingLifecycleHooks:getScalingLifecycleHooks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getScalingLifecycleHooks.
type GetScalingLifecycleHooksArgs struct {
	// A list of lifecycle hook ids.
	Ids []string `pulumi:"ids"`
	// A list of lifecycle hook names.
	LifecycleHookNames []string `pulumi:"lifecycleHookNames"`
	// A Name Regex of lifecycle hook.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// An id of scaling group id.
	ScalingGroupId string `pulumi:"scalingGroupId"`
}

// A collection of values returned by getScalingLifecycleHooks.
type GetScalingLifecycleHooksResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                 string   `pulumi:"id"`
	Ids                []string `pulumi:"ids"`
	LifecycleHookNames []string `pulumi:"lifecycleHookNames"`
	// The collection of lifecycle hook query.
	LifecycleHooks []GetScalingLifecycleHooksLifecycleHook `pulumi:"lifecycleHooks"`
	NameRegex      *string                                 `pulumi:"nameRegex"`
	OutputFile     *string                                 `pulumi:"outputFile"`
	// The id of the scaling group.
	ScalingGroupId string `pulumi:"scalingGroupId"`
	// The total count of lifecycle hook query.
	TotalCount int `pulumi:"totalCount"`
}

func GetScalingLifecycleHooksOutput(ctx *pulumi.Context, args GetScalingLifecycleHooksOutputArgs, opts ...pulumi.InvokeOption) GetScalingLifecycleHooksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetScalingLifecycleHooksResult, error) {
			args := v.(GetScalingLifecycleHooksArgs)
			r, err := GetScalingLifecycleHooks(ctx, &args, opts...)
			var s GetScalingLifecycleHooksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetScalingLifecycleHooksResultOutput)
}

// A collection of arguments for invoking getScalingLifecycleHooks.
type GetScalingLifecycleHooksOutputArgs struct {
	// A list of lifecycle hook ids.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A list of lifecycle hook names.
	LifecycleHookNames pulumi.StringArrayInput `pulumi:"lifecycleHookNames"`
	// A Name Regex of lifecycle hook.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// An id of scaling group id.
	ScalingGroupId pulumi.StringInput `pulumi:"scalingGroupId"`
}

func (GetScalingLifecycleHooksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetScalingLifecycleHooksArgs)(nil)).Elem()
}

// A collection of values returned by getScalingLifecycleHooks.
type GetScalingLifecycleHooksResultOutput struct{ *pulumi.OutputState }

func (GetScalingLifecycleHooksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetScalingLifecycleHooksResult)(nil)).Elem()
}

func (o GetScalingLifecycleHooksResultOutput) ToGetScalingLifecycleHooksResultOutput() GetScalingLifecycleHooksResultOutput {
	return o
}

func (o GetScalingLifecycleHooksResultOutput) ToGetScalingLifecycleHooksResultOutputWithContext(ctx context.Context) GetScalingLifecycleHooksResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetScalingLifecycleHooksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetScalingLifecycleHooksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetScalingLifecycleHooksResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetScalingLifecycleHooksResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetScalingLifecycleHooksResultOutput) LifecycleHookNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetScalingLifecycleHooksResult) []string { return v.LifecycleHookNames }).(pulumi.StringArrayOutput)
}

// The collection of lifecycle hook query.
func (o GetScalingLifecycleHooksResultOutput) LifecycleHooks() GetScalingLifecycleHooksLifecycleHookArrayOutput {
	return o.ApplyT(func(v GetScalingLifecycleHooksResult) []GetScalingLifecycleHooksLifecycleHook {
		return v.LifecycleHooks
	}).(GetScalingLifecycleHooksLifecycleHookArrayOutput)
}

func (o GetScalingLifecycleHooksResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScalingLifecycleHooksResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetScalingLifecycleHooksResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScalingLifecycleHooksResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The id of the scaling group.
func (o GetScalingLifecycleHooksResultOutput) ScalingGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetScalingLifecycleHooksResult) string { return v.ScalingGroupId }).(pulumi.StringOutput)
}

// The total count of lifecycle hook query.
func (o GetScalingLifecycleHooksResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetScalingLifecycleHooksResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetScalingLifecycleHooksResultOutput{})
}
