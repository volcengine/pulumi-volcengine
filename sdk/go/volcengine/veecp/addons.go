// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package veecp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of veecp addons
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/veecp"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooAddon, err := veecp.NewAddon(ctx, "fooAddon", &veecp.AddonArgs{
//				ClusterId:      pulumi.String("ccvd7mte6t101fno98u60"),
//				Version:        pulumi.String("1.8.6-edge.4"),
//				DeployNodeType: pulumi.String("Node"),
//				DeployMode:     pulumi.String("Unmanaged"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = veecp.GetAddonsOutput(ctx, veecp.GetAddonsOutputArgs{
//				Name: fooAddon.Name,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: volcengine.veecp.Addons has been deprecated in favor of volcengine.veecp.getAddons
func Addons(ctx *pulumi.Context, args *AddonsArgs, opts ...pulumi.InvokeOption) (*AddonsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv AddonsResult
	err := ctx.Invoke("volcengine:veecp/addons:Addons", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Addons.
type AddonsArgs struct {
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

// A collection of values returned by Addons.
type AddonsResult struct {
	// The collection of addons query.
	Addons      []AddonsAddon `pulumi:"addons"`
	Categories  []string      `pulumi:"categories"`
	DeployModes []string      `pulumi:"deployModes"`
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
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func AddonsOutput(ctx *pulumi.Context, args AddonsOutputArgs, opts ...pulumi.InvokeOption) AddonsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (AddonsResult, error) {
			args := v.(AddonsArgs)
			r, err := Addons(ctx, &args, opts...)
			var s AddonsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(AddonsResultOutput)
}

// A collection of arguments for invoking Addons.
type AddonsOutputArgs struct {
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

func (AddonsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonsArgs)(nil)).Elem()
}

// A collection of values returned by Addons.
type AddonsResultOutput struct{ *pulumi.OutputState }

func (AddonsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonsResult)(nil)).Elem()
}

func (o AddonsResultOutput) ToAddonsResultOutput() AddonsResultOutput {
	return o
}

func (o AddonsResultOutput) ToAddonsResultOutputWithContext(ctx context.Context) AddonsResultOutput {
	return o
}

// The collection of addons query.
func (o AddonsResultOutput) Addons() AddonsAddonArrayOutput {
	return o.ApplyT(func(v AddonsResult) []AddonsAddon { return v.Addons }).(AddonsAddonArrayOutput)
}

func (o AddonsResultOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AddonsResult) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

func (o AddonsResultOutput) DeployModes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AddonsResult) []string { return v.DeployModes }).(pulumi.StringArrayOutput)
}

// The deploy node types.
func (o AddonsResultOutput) DeployNodeTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AddonsResult) []string { return v.DeployNodeTypes }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o AddonsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AddonsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o AddonsResultOutput) KubernetesVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AddonsResult) []string { return v.KubernetesVersions }).(pulumi.StringArrayOutput)
}

// The name of addon.
func (o AddonsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddonsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o AddonsResultOutput) Necessaries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AddonsResult) []string { return v.Necessaries }).(pulumi.StringArrayOutput)
}

func (o AddonsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddonsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The network modes of pod.
func (o AddonsResultOutput) PodNetworkModes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AddonsResult) []string { return v.PodNetworkModes }).(pulumi.StringArrayOutput)
}

// The total count of query.
func (o AddonsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v AddonsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(AddonsResultOutput{})
}
