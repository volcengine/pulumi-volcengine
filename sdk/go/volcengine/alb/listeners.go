// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of alb listeners
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
//			_, err := alb.GetListeners(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: volcengine.alb.Listeners has been deprecated in favor of volcengine.alb.getListeners
func Listeners(ctx *pulumi.Context, args *ListenersArgs, opts ...pulumi.InvokeOption) (*ListenersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ListenersResult
	err := ctx.Invoke("volcengine:alb/listeners:Listeners", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Listeners.
type ListenersArgs struct {
	// A list of Listener IDs.
	Ids []string `pulumi:"ids"`
	// The name of the Listener.
	ListenerName *string `pulumi:"listenerName"`
	// The id of the Alb.
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	// A Name Regex of Listener.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The project name of the listener.
	ProjectName *string `pulumi:"projectName"`
}

// A collection of values returned by Listeners.
type ListenersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The name of the Listener.
	ListenerName *string `pulumi:"listenerName"`
	// The collection of Listener query.
	Listeners []ListenersListener `pulumi:"listeners"`
	// The load balancer ID that the listener belongs to.
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	NameRegex      *string `pulumi:"nameRegex"`
	OutputFile     *string `pulumi:"outputFile"`
	// The project name of the listener.
	ProjectName *string `pulumi:"projectName"`
	// The total count of Listener query.
	TotalCount int `pulumi:"totalCount"`
}

func ListenersOutput(ctx *pulumi.Context, args ListenersOutputArgs, opts ...pulumi.InvokeOption) ListenersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListenersResult, error) {
			args := v.(ListenersArgs)
			r, err := Listeners(ctx, &args, opts...)
			var s ListenersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListenersResultOutput)
}

// A collection of arguments for invoking Listeners.
type ListenersOutputArgs struct {
	// A list of Listener IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The name of the Listener.
	ListenerName pulumi.StringPtrInput `pulumi:"listenerName"`
	// The id of the Alb.
	LoadBalancerId pulumi.StringPtrInput `pulumi:"loadBalancerId"`
	// A Name Regex of Listener.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The project name of the listener.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
}

func (ListenersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenersArgs)(nil)).Elem()
}

// A collection of values returned by Listeners.
type ListenersResultOutput struct{ *pulumi.OutputState }

func (ListenersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenersResult)(nil)).Elem()
}

func (o ListenersResultOutput) ToListenersResultOutput() ListenersResultOutput {
	return o
}

func (o ListenersResultOutput) ToListenersResultOutputWithContext(ctx context.Context) ListenersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o ListenersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListenersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListenersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListenersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The name of the Listener.
func (o ListenersResultOutput) ListenerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListenersResult) *string { return v.ListenerName }).(pulumi.StringPtrOutput)
}

// The collection of Listener query.
func (o ListenersResultOutput) Listeners() ListenersListenerArrayOutput {
	return o.ApplyT(func(v ListenersResult) []ListenersListener { return v.Listeners }).(ListenersListenerArrayOutput)
}

// The load balancer ID that the listener belongs to.
func (o ListenersResultOutput) LoadBalancerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListenersResult) *string { return v.LoadBalancerId }).(pulumi.StringPtrOutput)
}

func (o ListenersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListenersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o ListenersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListenersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The project name of the listener.
func (o ListenersResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListenersResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// The total count of Listener query.
func (o ListenersResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v ListenersResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(ListenersResultOutput{})
}
