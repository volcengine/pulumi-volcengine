// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vke

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of vke node pools
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/vke"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vke.NodePools(ctx, &vke.NodePoolsArgs{
//				ClusterIds: []string{
//					"ccabe57fqtofgrbln3dog",
//				},
//				Name: pulumi.StringRef("demo"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func NodePools(ctx *pulumi.Context, args *NodePoolsArgs, opts ...pulumi.InvokeOption) (*NodePoolsResult, error) {
	var rv NodePoolsResult
	err := ctx.Invoke("volcengine:vke/nodePools:NodePools", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking NodePools.
type NodePoolsArgs struct {
	// Is enabled of AutoScaling.
	AutoScalingEnabled *bool `pulumi:"autoScalingEnabled"`
	// The ClusterId of NodePool.
	ClusterId *string `pulumi:"clusterId"`
	// The ClusterIds of NodePool IDs.
	ClusterIds []string `pulumi:"clusterIds"`
	// The ClientToken when successfully created.
	CreateClientToken *string `pulumi:"createClientToken"`
	// The IDs of NodePool.
	Ids []string `pulumi:"ids"`
	// The Name of NodePool.
	Name *string `pulumi:"name"`
	// A Name Regex of NodePool.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The Status of NodePool.
	Statuses []NodePoolsStatus `pulumi:"statuses"`
	// Tags.
	Tags []NodePoolsTag `pulumi:"tags"`
	// The ClientToken when last update was successful.
	UpdateClientToken *string `pulumi:"updateClientToken"`
}

// A collection of values returned by NodePools.
type NodePoolsResult struct {
	AutoScalingEnabled *bool `pulumi:"autoScalingEnabled"`
	// The ClusterId of NodePool.
	ClusterId  *string  `pulumi:"clusterId"`
	ClusterIds []string `pulumi:"clusterIds"`
	// The ClientToken when successfully created.
	CreateClientToken *string `pulumi:"createClientToken"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The Name of NodePool.
	Name      *string `pulumi:"name"`
	NameRegex *string `pulumi:"nameRegex"`
	// The collection of NodePools query.
	NodePools  []NodePoolsNodePool `pulumi:"nodePools"`
	OutputFile *string             `pulumi:"outputFile"`
	Statuses   []NodePoolsStatus   `pulumi:"statuses"`
	// Tags of the NodePool.
	Tags []NodePoolsTag `pulumi:"tags"`
	// Returns the total amount of the data list.
	TotalCount int `pulumi:"totalCount"`
	// The ClientToken when last update was successful.
	UpdateClientToken *string `pulumi:"updateClientToken"`
}

func NodePoolsOutput(ctx *pulumi.Context, args NodePoolsOutputArgs, opts ...pulumi.InvokeOption) NodePoolsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (NodePoolsResult, error) {
			args := v.(NodePoolsArgs)
			r, err := NodePools(ctx, &args, opts...)
			var s NodePoolsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(NodePoolsResultOutput)
}

// A collection of arguments for invoking NodePools.
type NodePoolsOutputArgs struct {
	// Is enabled of AutoScaling.
	AutoScalingEnabled pulumi.BoolPtrInput `pulumi:"autoScalingEnabled"`
	// The ClusterId of NodePool.
	ClusterId pulumi.StringPtrInput `pulumi:"clusterId"`
	// The ClusterIds of NodePool IDs.
	ClusterIds pulumi.StringArrayInput `pulumi:"clusterIds"`
	// The ClientToken when successfully created.
	CreateClientToken pulumi.StringPtrInput `pulumi:"createClientToken"`
	// The IDs of NodePool.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The Name of NodePool.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A Name Regex of NodePool.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The Status of NodePool.
	Statuses NodePoolsStatusArrayInput `pulumi:"statuses"`
	// Tags.
	Tags NodePoolsTagArrayInput `pulumi:"tags"`
	// The ClientToken when last update was successful.
	UpdateClientToken pulumi.StringPtrInput `pulumi:"updateClientToken"`
}

func (NodePoolsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodePoolsArgs)(nil)).Elem()
}

// A collection of values returned by NodePools.
type NodePoolsResultOutput struct{ *pulumi.OutputState }

func (NodePoolsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodePoolsResult)(nil)).Elem()
}

func (o NodePoolsResultOutput) ToNodePoolsResultOutput() NodePoolsResultOutput {
	return o
}

func (o NodePoolsResultOutput) ToNodePoolsResultOutputWithContext(ctx context.Context) NodePoolsResultOutput {
	return o
}

func (o NodePoolsResultOutput) AutoScalingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NodePoolsResult) *bool { return v.AutoScalingEnabled }).(pulumi.BoolPtrOutput)
}

// The ClusterId of NodePool.
func (o NodePoolsResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodePoolsResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

func (o NodePoolsResultOutput) ClusterIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NodePoolsResult) []string { return v.ClusterIds }).(pulumi.StringArrayOutput)
}

// The ClientToken when successfully created.
func (o NodePoolsResultOutput) CreateClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodePoolsResult) *string { return v.CreateClientToken }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o NodePoolsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v NodePoolsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o NodePoolsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NodePoolsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The Name of NodePool.
func (o NodePoolsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodePoolsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NodePoolsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodePoolsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// The collection of NodePools query.
func (o NodePoolsResultOutput) NodePools() NodePoolsNodePoolArrayOutput {
	return o.ApplyT(func(v NodePoolsResult) []NodePoolsNodePool { return v.NodePools }).(NodePoolsNodePoolArrayOutput)
}

func (o NodePoolsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodePoolsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o NodePoolsResultOutput) Statuses() NodePoolsStatusArrayOutput {
	return o.ApplyT(func(v NodePoolsResult) []NodePoolsStatus { return v.Statuses }).(NodePoolsStatusArrayOutput)
}

// Tags of the NodePool.
func (o NodePoolsResultOutput) Tags() NodePoolsTagArrayOutput {
	return o.ApplyT(func(v NodePoolsResult) []NodePoolsTag { return v.Tags }).(NodePoolsTagArrayOutput)
}

// Returns the total amount of the data list.
func (o NodePoolsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodePoolsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The ClientToken when last update was successful.
func (o NodePoolsResultOutput) UpdateClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodePoolsResult) *string { return v.UpdateClientToken }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NodePoolsResultOutput{})
}