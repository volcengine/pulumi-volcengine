// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vke

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of vke addons
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
//			_, err := vke.GetAddons(ctx, &vke.GetAddonsArgs{
//				ClusterIds: []string{
//					"cccctv1vqtofp49d96ujg",
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
func GetAddons(ctx *pulumi.Context, args *GetAddonsArgs, opts ...pulumi.InvokeOption) (*GetAddonsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAddonsResult
	err := ctx.Invoke("volcengine:vke/getAddons:getAddons", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAddons.
type GetAddonsArgs struct {
	// The IDs of Cluster.
	ClusterIds []string `pulumi:"clusterIds"`
	// ClientToken when the addon is created successfully. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
	CreateClientToken *string `pulumi:"createClientToken"`
	// The deploy model, the value is `Managed` or `Unmanaged`.
	DeployModes []string `pulumi:"deployModes"`
	// The deploy node types, the value is `Node` or `VirtualNode`. Only effected when deployMode is `Unmanaged`.
	DeployNodeTypes []string `pulumi:"deployNodeTypes"`
	// A Name Regex of addon.
	NameRegex *string `pulumi:"nameRegex"`
	// The Names of addons.
	Names []string `pulumi:"names"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// Array of addon states to filter.
	Statuses []GetAddonsStatus `pulumi:"statuses"`
	// The ClientToken when the last addon update succeeded. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
	UpdateClientToken *string `pulumi:"updateClientToken"`
}

// A collection of values returned by getAddons.
type GetAddonsResult struct {
	// The collection of addon query.
	Addons            []GetAddonsAddon `pulumi:"addons"`
	ClusterIds        []string         `pulumi:"clusterIds"`
	CreateClientToken *string          `pulumi:"createClientToken"`
	DeployModes       []string         `pulumi:"deployModes"`
	DeployNodeTypes   []string         `pulumi:"deployNodeTypes"`
	// The provider-assigned unique ID for this managed resource.
	Id         string            `pulumi:"id"`
	NameRegex  *string           `pulumi:"nameRegex"`
	Names      []string          `pulumi:"names"`
	OutputFile *string           `pulumi:"outputFile"`
	Statuses   []GetAddonsStatus `pulumi:"statuses"`
	// The total count of addon query.
	TotalCount        int     `pulumi:"totalCount"`
	UpdateClientToken *string `pulumi:"updateClientToken"`
}

func GetAddonsOutput(ctx *pulumi.Context, args GetAddonsOutputArgs, opts ...pulumi.InvokeOption) GetAddonsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAddonsResult, error) {
			args := v.(GetAddonsArgs)
			r, err := GetAddons(ctx, &args, opts...)
			var s GetAddonsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAddonsResultOutput)
}

// A collection of arguments for invoking getAddons.
type GetAddonsOutputArgs struct {
	// The IDs of Cluster.
	ClusterIds pulumi.StringArrayInput `pulumi:"clusterIds"`
	// ClientToken when the addon is created successfully. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
	CreateClientToken pulumi.StringPtrInput `pulumi:"createClientToken"`
	// The deploy model, the value is `Managed` or `Unmanaged`.
	DeployModes pulumi.StringArrayInput `pulumi:"deployModes"`
	// The deploy node types, the value is `Node` or `VirtualNode`. Only effected when deployMode is `Unmanaged`.
	DeployNodeTypes pulumi.StringArrayInput `pulumi:"deployNodeTypes"`
	// A Name Regex of addon.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// The Names of addons.
	Names pulumi.StringArrayInput `pulumi:"names"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Array of addon states to filter.
	Statuses GetAddonsStatusArrayInput `pulumi:"statuses"`
	// The ClientToken when the last addon update succeeded. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
	UpdateClientToken pulumi.StringPtrInput `pulumi:"updateClientToken"`
}

func (GetAddonsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddonsArgs)(nil)).Elem()
}

// A collection of values returned by getAddons.
type GetAddonsResultOutput struct{ *pulumi.OutputState }

func (GetAddonsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAddonsResult)(nil)).Elem()
}

func (o GetAddonsResultOutput) ToGetAddonsResultOutput() GetAddonsResultOutput {
	return o
}

func (o GetAddonsResultOutput) ToGetAddonsResultOutputWithContext(ctx context.Context) GetAddonsResultOutput {
	return o
}

// The collection of addon query.
func (o GetAddonsResultOutput) Addons() GetAddonsAddonArrayOutput {
	return o.ApplyT(func(v GetAddonsResult) []GetAddonsAddon { return v.Addons }).(GetAddonsAddonArrayOutput)
}

func (o GetAddonsResultOutput) ClusterIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAddonsResult) []string { return v.ClusterIds }).(pulumi.StringArrayOutput)
}

func (o GetAddonsResultOutput) CreateClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddonsResult) *string { return v.CreateClientToken }).(pulumi.StringPtrOutput)
}

func (o GetAddonsResultOutput) DeployModes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAddonsResult) []string { return v.DeployModes }).(pulumi.StringArrayOutput)
}

func (o GetAddonsResultOutput) DeployNodeTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAddonsResult) []string { return v.DeployNodeTypes }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAddonsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAddonsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAddonsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddonsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetAddonsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAddonsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetAddonsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddonsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetAddonsResultOutput) Statuses() GetAddonsStatusArrayOutput {
	return o.ApplyT(func(v GetAddonsResult) []GetAddonsStatus { return v.Statuses }).(GetAddonsStatusArrayOutput)
}

// The total count of addon query.
func (o GetAddonsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetAddonsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func (o GetAddonsResultOutput) UpdateClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAddonsResult) *string { return v.UpdateClientToken }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAddonsResultOutput{})
}
