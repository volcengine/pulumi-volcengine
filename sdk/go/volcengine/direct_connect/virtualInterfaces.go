// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package direct_connect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of direct connect virtual interfaces
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/direct_connect"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := direct_connect.VirtualInterfaces(ctx, &direct_connect.VirtualInterfacesArgs{
//				VirtualInterfaceName: pulumi.StringRef("tf-test"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func VirtualInterfaces(ctx *pulumi.Context, args *VirtualInterfacesArgs, opts ...pulumi.InvokeOption) (*VirtualInterfacesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv VirtualInterfacesResult
	err := ctx.Invoke("volcengine:direct_connect/virtualInterfaces:VirtualInterfaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking VirtualInterfaces.
type VirtualInterfacesArgs struct {
	// The direct connect connection ID that associated with this virtual interface.
	DirectConnectConnectionId *string `pulumi:"directConnectConnectionId"`
	// The direct connect gateway ID that associated with this virtual interface.
	DirectConnectGatewayId *string `pulumi:"directConnectGatewayId"`
	// A list of IDs.
	Ids []string `pulumi:"ids"`
	// The local IP that associated with this virtual interface.
	LocalIp *string `pulumi:"localIp"`
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The peer IP that associated with this virtual interface.
	PeerIp *string `pulumi:"peerIp"`
	// The route type of virtual interface.
	RouteType *string `pulumi:"routeType"`
	// The filter tag of direct connect virtual interface.
	TagFilters []VirtualInterfacesTagFilter `pulumi:"tagFilters"`
	// The name of virtual interface.
	VirtualInterfaceName *string `pulumi:"virtualInterfaceName"`
	// The VLAN ID of virtual interface.
	VlanId *int `pulumi:"vlanId"`
}

// A collection of values returned by VirtualInterfaces.
type VirtualInterfacesResult struct {
	// The direct connect connection ID which associated with this virtual interface.
	DirectConnectConnectionId *string `pulumi:"directConnectConnectionId"`
	// The direct connect gateway ID which associated with this virtual interface.
	DirectConnectGatewayId *string `pulumi:"directConnectGatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The local IP that associated with this virtual interface.
	LocalIp    *string `pulumi:"localIp"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The peer IP that associated with this virtual interface.
	PeerIp *string `pulumi:"peerIp"`
	// The route type of this virtual interface.
	RouteType  *string                      `pulumi:"routeType"`
	TagFilters []VirtualInterfacesTagFilter `pulumi:"tagFilters"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
	// The name of virtual interface.
	VirtualInterfaceName *string `pulumi:"virtualInterfaceName"`
	// The collection of query.
	VirtualInterfaces []VirtualInterfacesVirtualInterface `pulumi:"virtualInterfaces"`
	// The VLAN ID of virtual interface.
	VlanId *int `pulumi:"vlanId"`
}

func VirtualInterfacesOutput(ctx *pulumi.Context, args VirtualInterfacesOutputArgs, opts ...pulumi.InvokeOption) VirtualInterfacesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (VirtualInterfacesResult, error) {
			args := v.(VirtualInterfacesArgs)
			r, err := VirtualInterfaces(ctx, &args, opts...)
			var s VirtualInterfacesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(VirtualInterfacesResultOutput)
}

// A collection of arguments for invoking VirtualInterfaces.
type VirtualInterfacesOutputArgs struct {
	// The direct connect connection ID that associated with this virtual interface.
	DirectConnectConnectionId pulumi.StringPtrInput `pulumi:"directConnectConnectionId"`
	// The direct connect gateway ID that associated with this virtual interface.
	DirectConnectGatewayId pulumi.StringPtrInput `pulumi:"directConnectGatewayId"`
	// A list of IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The local IP that associated with this virtual interface.
	LocalIp pulumi.StringPtrInput `pulumi:"localIp"`
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The peer IP that associated with this virtual interface.
	PeerIp pulumi.StringPtrInput `pulumi:"peerIp"`
	// The route type of virtual interface.
	RouteType pulumi.StringPtrInput `pulumi:"routeType"`
	// The filter tag of direct connect virtual interface.
	TagFilters VirtualInterfacesTagFilterArrayInput `pulumi:"tagFilters"`
	// The name of virtual interface.
	VirtualInterfaceName pulumi.StringPtrInput `pulumi:"virtualInterfaceName"`
	// The VLAN ID of virtual interface.
	VlanId pulumi.IntPtrInput `pulumi:"vlanId"`
}

func (VirtualInterfacesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualInterfacesArgs)(nil)).Elem()
}

// A collection of values returned by VirtualInterfaces.
type VirtualInterfacesResultOutput struct{ *pulumi.OutputState }

func (VirtualInterfacesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualInterfacesResult)(nil)).Elem()
}

func (o VirtualInterfacesResultOutput) ToVirtualInterfacesResultOutput() VirtualInterfacesResultOutput {
	return o
}

func (o VirtualInterfacesResultOutput) ToVirtualInterfacesResultOutputWithContext(ctx context.Context) VirtualInterfacesResultOutput {
	return o
}

// The direct connect connection ID which associated with this virtual interface.
func (o VirtualInterfacesResultOutput) DirectConnectConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualInterfacesResult) *string { return v.DirectConnectConnectionId }).(pulumi.StringPtrOutput)
}

// The direct connect gateway ID which associated with this virtual interface.
func (o VirtualInterfacesResultOutput) DirectConnectGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualInterfacesResult) *string { return v.DirectConnectGatewayId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o VirtualInterfacesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualInterfacesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualInterfacesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualInterfacesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The local IP that associated with this virtual interface.
func (o VirtualInterfacesResultOutput) LocalIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualInterfacesResult) *string { return v.LocalIp }).(pulumi.StringPtrOutput)
}

func (o VirtualInterfacesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualInterfacesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o VirtualInterfacesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualInterfacesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The peer IP that associated with this virtual interface.
func (o VirtualInterfacesResultOutput) PeerIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualInterfacesResult) *string { return v.PeerIp }).(pulumi.StringPtrOutput)
}

// The route type of this virtual interface.
func (o VirtualInterfacesResultOutput) RouteType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualInterfacesResult) *string { return v.RouteType }).(pulumi.StringPtrOutput)
}

func (o VirtualInterfacesResultOutput) TagFilters() VirtualInterfacesTagFilterArrayOutput {
	return o.ApplyT(func(v VirtualInterfacesResult) []VirtualInterfacesTagFilter { return v.TagFilters }).(VirtualInterfacesTagFilterArrayOutput)
}

// The total count of query.
func (o VirtualInterfacesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v VirtualInterfacesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The name of virtual interface.
func (o VirtualInterfacesResultOutput) VirtualInterfaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualInterfacesResult) *string { return v.VirtualInterfaceName }).(pulumi.StringPtrOutput)
}

// The collection of query.
func (o VirtualInterfacesResultOutput) VirtualInterfaces() VirtualInterfacesVirtualInterfaceArrayOutput {
	return o.ApplyT(func(v VirtualInterfacesResult) []VirtualInterfacesVirtualInterface { return v.VirtualInterfaces }).(VirtualInterfacesVirtualInterfaceArrayOutput)
}

// The VLAN ID of virtual interface.
func (o VirtualInterfacesResultOutput) VlanId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualInterfacesResult) *int { return v.VlanId }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualInterfacesResultOutput{})
}