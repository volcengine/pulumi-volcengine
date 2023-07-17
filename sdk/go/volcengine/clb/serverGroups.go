// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of server groups
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/clb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := clb.ServerGroups(ctx, &clb.ServerGroupsArgs{
//				Ids: []string{
//					"rsp-273yv0kir1vk07fap8tt9jtwg",
//					"rsp-273yxuqfova4g7fap8tyemn6t",
//					"rsp-273z9pt9lpdds7fap8sqdvfrf",
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
func ServerGroups(ctx *pulumi.Context, args *ServerGroupsArgs, opts ...pulumi.InvokeOption) (*ServerGroupsResult, error) {
	var rv ServerGroupsResult
	err := ctx.Invoke("volcengine:clb/serverGroups:ServerGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking ServerGroups.
type ServerGroupsArgs struct {
	// A list of ServerGroup IDs.
	Ids []string `pulumi:"ids"`
	// The id of the Clb.
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	// A Name Regex of ServerGroup.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The name of the ServerGroup.
	ServerGroupName *string `pulumi:"serverGroupName"`
}

// A collection of values returned by ServerGroups.
type ServerGroupsResult struct {
	// The collection of ServerGroup query.
	Groups []ServerGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id             string   `pulumi:"id"`
	Ids            []string `pulumi:"ids"`
	LoadBalancerId *string  `pulumi:"loadBalancerId"`
	NameRegex      *string  `pulumi:"nameRegex"`
	OutputFile     *string  `pulumi:"outputFile"`
	// The name of the ServerGroup.
	ServerGroupName *string `pulumi:"serverGroupName"`
	// The total count of ServerGroup query.
	TotalCount int `pulumi:"totalCount"`
}

func ServerGroupsOutput(ctx *pulumi.Context, args ServerGroupsOutputArgs, opts ...pulumi.InvokeOption) ServerGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ServerGroupsResult, error) {
			args := v.(ServerGroupsArgs)
			r, err := ServerGroups(ctx, &args, opts...)
			var s ServerGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ServerGroupsResultOutput)
}

// A collection of arguments for invoking ServerGroups.
type ServerGroupsOutputArgs struct {
	// A list of ServerGroup IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The id of the Clb.
	LoadBalancerId pulumi.StringPtrInput `pulumi:"loadBalancerId"`
	// A Name Regex of ServerGroup.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The name of the ServerGroup.
	ServerGroupName pulumi.StringPtrInput `pulumi:"serverGroupName"`
}

func (ServerGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerGroupsArgs)(nil)).Elem()
}

// A collection of values returned by ServerGroups.
type ServerGroupsResultOutput struct{ *pulumi.OutputState }

func (ServerGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerGroupsResult)(nil)).Elem()
}

func (o ServerGroupsResultOutput) ToServerGroupsResultOutput() ServerGroupsResultOutput {
	return o
}

func (o ServerGroupsResultOutput) ToServerGroupsResultOutputWithContext(ctx context.Context) ServerGroupsResultOutput {
	return o
}

// The collection of ServerGroup query.
func (o ServerGroupsResultOutput) Groups() ServerGroupsGroupArrayOutput {
	return o.ApplyT(func(v ServerGroupsResult) []ServerGroupsGroup { return v.Groups }).(ServerGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o ServerGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ServerGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ServerGroupsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServerGroupsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o ServerGroupsResultOutput) LoadBalancerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerGroupsResult) *string { return v.LoadBalancerId }).(pulumi.StringPtrOutput)
}

func (o ServerGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o ServerGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The name of the ServerGroup.
func (o ServerGroupsResultOutput) ServerGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerGroupsResult) *string { return v.ServerGroupName }).(pulumi.StringPtrOutput)
}

// The total count of ServerGroup query.
func (o ServerGroupsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v ServerGroupsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerGroupsResultOutput{})
}