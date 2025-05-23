// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vke

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of vke support addons
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vke"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vke.GetSupportAddons(ctx, &vke.GetSupportAddonsArgs{
//				Categories: []string{
//					"Monitor",
//				},
//				Name: pulumi.StringRef("metrics-server"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSupportAddons(ctx *pulumi.Context, args *GetSupportAddonsArgs, opts ...pulumi.InvokeOption) (*GetSupportAddonsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSupportAddonsResult
	err := ctx.Invoke("volcengine:vke/getSupportAddons:getSupportAddons", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSupportAddons.
type GetSupportAddonsArgs struct {
	// The categories of addons, the value is `Storage` or `Network` or `Monitor` or `Scheduler` or `Dns` or `Security` or `Gpu` or `Image`.
	Categories []string `pulumi:"categories"`
	// The deploy model, the value is `Managed` or `Unmanaged`.
	DeployModes []string `pulumi:"deployModes"`
	// The deploy node types, the value is `Node` or `VirtualNode`. Only effected when deployMode is `Unmanaged`.
	DeployNodeTypes []string `pulumi:"deployNodeTypes"`
	// A list of Kubernetes Versions.
	KubernetesVersions []string `pulumi:"kubernetesVersions"`
	// The name of the addon.
	Name *string `pulumi:"name"`
	// The necessaries of addons, the value is `Required` or `Recommended` or `OnDemand`.
	Necessaries []string `pulumi:"necessaries"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The container network model, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.
	PodNetworkModes []string `pulumi:"podNetworkModes"`
}

// A collection of values returned by getSupportAddons.
type GetSupportAddonsResult struct {
	// The collection of addons query.
	Addons      []GetSupportAddonsAddon `pulumi:"addons"`
	Categories  []string                `pulumi:"categories"`
	DeployModes []string                `pulumi:"deployModes"`
	// The deploy node types.
	DeployNodeTypes []string `pulumi:"deployNodeTypes"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string   `pulumi:"id"`
	KubernetesVersions []string `pulumi:"kubernetesVersions"`
	// The name of addon.
	Name        *string  `pulumi:"name"`
	Necessaries []string `pulumi:"necessaries"`
	OutputFile  *string  `pulumi:"outputFile"`
	// The network modes of pod.
	PodNetworkModes []string `pulumi:"podNetworkModes"`
	// The total count of addons query.
	TotalCount int `pulumi:"totalCount"`
}

func GetSupportAddonsOutput(ctx *pulumi.Context, args GetSupportAddonsOutputArgs, opts ...pulumi.InvokeOption) GetSupportAddonsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSupportAddonsResult, error) {
			args := v.(GetSupportAddonsArgs)
			r, err := GetSupportAddons(ctx, &args, opts...)
			var s GetSupportAddonsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSupportAddonsResultOutput)
}

// A collection of arguments for invoking getSupportAddons.
type GetSupportAddonsOutputArgs struct {
	// The categories of addons, the value is `Storage` or `Network` or `Monitor` or `Scheduler` or `Dns` or `Security` or `Gpu` or `Image`.
	Categories pulumi.StringArrayInput `pulumi:"categories"`
	// The deploy model, the value is `Managed` or `Unmanaged`.
	DeployModes pulumi.StringArrayInput `pulumi:"deployModes"`
	// The deploy node types, the value is `Node` or `VirtualNode`. Only effected when deployMode is `Unmanaged`.
	DeployNodeTypes pulumi.StringArrayInput `pulumi:"deployNodeTypes"`
	// A list of Kubernetes Versions.
	KubernetesVersions pulumi.StringArrayInput `pulumi:"kubernetesVersions"`
	// The name of the addon.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The necessaries of addons, the value is `Required` or `Recommended` or `OnDemand`.
	Necessaries pulumi.StringArrayInput `pulumi:"necessaries"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The container network model, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.
	PodNetworkModes pulumi.StringArrayInput `pulumi:"podNetworkModes"`
}

func (GetSupportAddonsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSupportAddonsArgs)(nil)).Elem()
}

// A collection of values returned by getSupportAddons.
type GetSupportAddonsResultOutput struct{ *pulumi.OutputState }

func (GetSupportAddonsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSupportAddonsResult)(nil)).Elem()
}

func (o GetSupportAddonsResultOutput) ToGetSupportAddonsResultOutput() GetSupportAddonsResultOutput {
	return o
}

func (o GetSupportAddonsResultOutput) ToGetSupportAddonsResultOutputWithContext(ctx context.Context) GetSupportAddonsResultOutput {
	return o
}

// The collection of addons query.
func (o GetSupportAddonsResultOutput) Addons() GetSupportAddonsAddonArrayOutput {
	return o.ApplyT(func(v GetSupportAddonsResult) []GetSupportAddonsAddon { return v.Addons }).(GetSupportAddonsAddonArrayOutput)
}

func (o GetSupportAddonsResultOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSupportAddonsResult) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

func (o GetSupportAddonsResultOutput) DeployModes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSupportAddonsResult) []string { return v.DeployModes }).(pulumi.StringArrayOutput)
}

// The deploy node types.
func (o GetSupportAddonsResultOutput) DeployNodeTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSupportAddonsResult) []string { return v.DeployNodeTypes }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSupportAddonsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSupportAddonsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSupportAddonsResultOutput) KubernetesVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSupportAddonsResult) []string { return v.KubernetesVersions }).(pulumi.StringArrayOutput)
}

// The name of addon.
func (o GetSupportAddonsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSupportAddonsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetSupportAddonsResultOutput) Necessaries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSupportAddonsResult) []string { return v.Necessaries }).(pulumi.StringArrayOutput)
}

func (o GetSupportAddonsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSupportAddonsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The network modes of pod.
func (o GetSupportAddonsResultOutput) PodNetworkModes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSupportAddonsResult) []string { return v.PodNetworkModes }).(pulumi.StringArrayOutput)
}

// The total count of addons query.
func (o GetSupportAddonsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetSupportAddonsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSupportAddonsResultOutput{})
}
