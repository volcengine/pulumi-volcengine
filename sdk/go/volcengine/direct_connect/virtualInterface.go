// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package direct_connect

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage direct connect virtual interface
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
//			_, err := direct_connect.NewVirtualInterface(ctx, "foo", &direct_connect.VirtualInterfaceArgs{
//				Description:               pulumi.String("tf-test"),
//				DirectConnectConnectionId: pulumi.String("dcc-rtkzeotzst1cu3numzi****"),
//				DirectConnectGatewayId:    pulumi.String("dcg-638x4bjvjawwn3gd5xw****"),
//				EnableBfd:                 pulumi.Bool(false),
//				LocalIp:                   pulumi.String("**.**.**.**/**"),
//				PeerIp:                    pulumi.String("**.**.**.**/**"),
//				RouteType:                 pulumi.String("Static"),
//				Tags: direct_connect.VirtualInterfaceTagArray{
//					&direct_connect.VirtualInterfaceTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
//					},
//				},
//				VirtualInterfaceName: pulumi.String("tf-test-vi"),
//				VlanId:               pulumi.Int(2),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// DirectConnectVirtualInterface can be imported using the id, e.g.
//
// ```sh
// $ pulumi import volcengine:direct_connect/virtualInterface:VirtualInterface default resource_id
// ```
type VirtualInterface struct {
	pulumi.CustomResourceState

	// The band width limit of virtual interface,in Mbps.
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// The BFD detect interval.
	BfdDetectInterval pulumi.IntPtrOutput `pulumi:"bfdDetectInterval"`
	// The BFD detect times.
	BfdDetectMultiplier pulumi.IntPtrOutput `pulumi:"bfdDetectMultiplier"`
	// The description of virtual interface.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The direct connect connection ID which associated with.
	DirectConnectConnectionId pulumi.StringOutput `pulumi:"directConnectConnectionId"`
	// The direct connect gateway ID which associated with.
	DirectConnectGatewayId pulumi.StringOutput `pulumi:"directConnectGatewayId"`
	// Whether enable BFD detect.
	EnableBfd pulumi.BoolPtrOutput `pulumi:"enableBfd"`
	// Whether enable NQA detect.
	EnableNqa pulumi.BoolPtrOutput `pulumi:"enableNqa"`
	// The local IP that associated with.
	LocalIp pulumi.StringOutput `pulumi:"localIp"`
	// The NQA detect interval.
	NqaDetectInterval pulumi.IntPtrOutput `pulumi:"nqaDetectInterval"`
	// The NAQ detect times.
	NqaDetectMultiplier pulumi.IntPtrOutput `pulumi:"nqaDetectMultiplier"`
	// The peer IP that associated with.
	PeerIp pulumi.StringOutput `pulumi:"peerIp"`
	// The route type of virtual interface,valid value contains `Static`,`BGP`.
	RouteType pulumi.StringOutput `pulumi:"routeType"`
	// The tags that direct connect gateway added.
	Tags VirtualInterfaceTagArrayOutput `pulumi:"tags"`
	// The name of virtual interface.
	VirtualInterfaceName pulumi.StringPtrOutput `pulumi:"virtualInterfaceName"`
	// The VLAN ID used to connect to the local IDC, please ensure that this VLAN ID is not occupied, the value range: 0 ~ 2999.
	VlanId pulumi.IntOutput `pulumi:"vlanId"`
}

// NewVirtualInterface registers a new resource with the given unique name, arguments, and options.
func NewVirtualInterface(ctx *pulumi.Context,
	name string, args *VirtualInterfaceArgs, opts ...pulumi.ResourceOption) (*VirtualInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DirectConnectConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'DirectConnectConnectionId'")
	}
	if args.DirectConnectGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'DirectConnectGatewayId'")
	}
	if args.LocalIp == nil {
		return nil, errors.New("invalid value for required argument 'LocalIp'")
	}
	if args.PeerIp == nil {
		return nil, errors.New("invalid value for required argument 'PeerIp'")
	}
	if args.VlanId == nil {
		return nil, errors.New("invalid value for required argument 'VlanId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualInterface
	err := ctx.RegisterResource("volcengine:direct_connect/virtualInterface:VirtualInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualInterface gets an existing VirtualInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualInterfaceState, opts ...pulumi.ResourceOption) (*VirtualInterface, error) {
	var resource VirtualInterface
	err := ctx.ReadResource("volcengine:direct_connect/virtualInterface:VirtualInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualInterface resources.
type virtualInterfaceState struct {
	// The band width limit of virtual interface,in Mbps.
	Bandwidth *int `pulumi:"bandwidth"`
	// The BFD detect interval.
	BfdDetectInterval *int `pulumi:"bfdDetectInterval"`
	// The BFD detect times.
	BfdDetectMultiplier *int `pulumi:"bfdDetectMultiplier"`
	// The description of virtual interface.
	Description *string `pulumi:"description"`
	// The direct connect connection ID which associated with.
	DirectConnectConnectionId *string `pulumi:"directConnectConnectionId"`
	// The direct connect gateway ID which associated with.
	DirectConnectGatewayId *string `pulumi:"directConnectGatewayId"`
	// Whether enable BFD detect.
	EnableBfd *bool `pulumi:"enableBfd"`
	// Whether enable NQA detect.
	EnableNqa *bool `pulumi:"enableNqa"`
	// The local IP that associated with.
	LocalIp *string `pulumi:"localIp"`
	// The NQA detect interval.
	NqaDetectInterval *int `pulumi:"nqaDetectInterval"`
	// The NAQ detect times.
	NqaDetectMultiplier *int `pulumi:"nqaDetectMultiplier"`
	// The peer IP that associated with.
	PeerIp *string `pulumi:"peerIp"`
	// The route type of virtual interface,valid value contains `Static`,`BGP`.
	RouteType *string `pulumi:"routeType"`
	// The tags that direct connect gateway added.
	Tags []VirtualInterfaceTag `pulumi:"tags"`
	// The name of virtual interface.
	VirtualInterfaceName *string `pulumi:"virtualInterfaceName"`
	// The VLAN ID used to connect to the local IDC, please ensure that this VLAN ID is not occupied, the value range: 0 ~ 2999.
	VlanId *int `pulumi:"vlanId"`
}

type VirtualInterfaceState struct {
	// The band width limit of virtual interface,in Mbps.
	Bandwidth pulumi.IntPtrInput
	// The BFD detect interval.
	BfdDetectInterval pulumi.IntPtrInput
	// The BFD detect times.
	BfdDetectMultiplier pulumi.IntPtrInput
	// The description of virtual interface.
	Description pulumi.StringPtrInput
	// The direct connect connection ID which associated with.
	DirectConnectConnectionId pulumi.StringPtrInput
	// The direct connect gateway ID which associated with.
	DirectConnectGatewayId pulumi.StringPtrInput
	// Whether enable BFD detect.
	EnableBfd pulumi.BoolPtrInput
	// Whether enable NQA detect.
	EnableNqa pulumi.BoolPtrInput
	// The local IP that associated with.
	LocalIp pulumi.StringPtrInput
	// The NQA detect interval.
	NqaDetectInterval pulumi.IntPtrInput
	// The NAQ detect times.
	NqaDetectMultiplier pulumi.IntPtrInput
	// The peer IP that associated with.
	PeerIp pulumi.StringPtrInput
	// The route type of virtual interface,valid value contains `Static`,`BGP`.
	RouteType pulumi.StringPtrInput
	// The tags that direct connect gateway added.
	Tags VirtualInterfaceTagArrayInput
	// The name of virtual interface.
	VirtualInterfaceName pulumi.StringPtrInput
	// The VLAN ID used to connect to the local IDC, please ensure that this VLAN ID is not occupied, the value range: 0 ~ 2999.
	VlanId pulumi.IntPtrInput
}

func (VirtualInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualInterfaceState)(nil)).Elem()
}

type virtualInterfaceArgs struct {
	// The band width limit of virtual interface,in Mbps.
	Bandwidth *int `pulumi:"bandwidth"`
	// The BFD detect interval.
	BfdDetectInterval *int `pulumi:"bfdDetectInterval"`
	// The BFD detect times.
	BfdDetectMultiplier *int `pulumi:"bfdDetectMultiplier"`
	// The description of virtual interface.
	Description *string `pulumi:"description"`
	// The direct connect connection ID which associated with.
	DirectConnectConnectionId string `pulumi:"directConnectConnectionId"`
	// The direct connect gateway ID which associated with.
	DirectConnectGatewayId string `pulumi:"directConnectGatewayId"`
	// Whether enable BFD detect.
	EnableBfd *bool `pulumi:"enableBfd"`
	// Whether enable NQA detect.
	EnableNqa *bool `pulumi:"enableNqa"`
	// The local IP that associated with.
	LocalIp string `pulumi:"localIp"`
	// The NQA detect interval.
	NqaDetectInterval *int `pulumi:"nqaDetectInterval"`
	// The NAQ detect times.
	NqaDetectMultiplier *int `pulumi:"nqaDetectMultiplier"`
	// The peer IP that associated with.
	PeerIp string `pulumi:"peerIp"`
	// The route type of virtual interface,valid value contains `Static`,`BGP`.
	RouteType *string `pulumi:"routeType"`
	// The tags that direct connect gateway added.
	Tags []VirtualInterfaceTag `pulumi:"tags"`
	// The name of virtual interface.
	VirtualInterfaceName *string `pulumi:"virtualInterfaceName"`
	// The VLAN ID used to connect to the local IDC, please ensure that this VLAN ID is not occupied, the value range: 0 ~ 2999.
	VlanId int `pulumi:"vlanId"`
}

// The set of arguments for constructing a VirtualInterface resource.
type VirtualInterfaceArgs struct {
	// The band width limit of virtual interface,in Mbps.
	Bandwidth pulumi.IntPtrInput
	// The BFD detect interval.
	BfdDetectInterval pulumi.IntPtrInput
	// The BFD detect times.
	BfdDetectMultiplier pulumi.IntPtrInput
	// The description of virtual interface.
	Description pulumi.StringPtrInput
	// The direct connect connection ID which associated with.
	DirectConnectConnectionId pulumi.StringInput
	// The direct connect gateway ID which associated with.
	DirectConnectGatewayId pulumi.StringInput
	// Whether enable BFD detect.
	EnableBfd pulumi.BoolPtrInput
	// Whether enable NQA detect.
	EnableNqa pulumi.BoolPtrInput
	// The local IP that associated with.
	LocalIp pulumi.StringInput
	// The NQA detect interval.
	NqaDetectInterval pulumi.IntPtrInput
	// The NAQ detect times.
	NqaDetectMultiplier pulumi.IntPtrInput
	// The peer IP that associated with.
	PeerIp pulumi.StringInput
	// The route type of virtual interface,valid value contains `Static`,`BGP`.
	RouteType pulumi.StringPtrInput
	// The tags that direct connect gateway added.
	Tags VirtualInterfaceTagArrayInput
	// The name of virtual interface.
	VirtualInterfaceName pulumi.StringPtrInput
	// The VLAN ID used to connect to the local IDC, please ensure that this VLAN ID is not occupied, the value range: 0 ~ 2999.
	VlanId pulumi.IntInput
}

func (VirtualInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualInterfaceArgs)(nil)).Elem()
}

type VirtualInterfaceInput interface {
	pulumi.Input

	ToVirtualInterfaceOutput() VirtualInterfaceOutput
	ToVirtualInterfaceOutputWithContext(ctx context.Context) VirtualInterfaceOutput
}

func (*VirtualInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualInterface)(nil)).Elem()
}

func (i *VirtualInterface) ToVirtualInterfaceOutput() VirtualInterfaceOutput {
	return i.ToVirtualInterfaceOutputWithContext(context.Background())
}

func (i *VirtualInterface) ToVirtualInterfaceOutputWithContext(ctx context.Context) VirtualInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualInterfaceOutput)
}

// VirtualInterfaceArrayInput is an input type that accepts VirtualInterfaceArray and VirtualInterfaceArrayOutput values.
// You can construct a concrete instance of `VirtualInterfaceArrayInput` via:
//
//	VirtualInterfaceArray{ VirtualInterfaceArgs{...} }
type VirtualInterfaceArrayInput interface {
	pulumi.Input

	ToVirtualInterfaceArrayOutput() VirtualInterfaceArrayOutput
	ToVirtualInterfaceArrayOutputWithContext(context.Context) VirtualInterfaceArrayOutput
}

type VirtualInterfaceArray []VirtualInterfaceInput

func (VirtualInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualInterface)(nil)).Elem()
}

func (i VirtualInterfaceArray) ToVirtualInterfaceArrayOutput() VirtualInterfaceArrayOutput {
	return i.ToVirtualInterfaceArrayOutputWithContext(context.Background())
}

func (i VirtualInterfaceArray) ToVirtualInterfaceArrayOutputWithContext(ctx context.Context) VirtualInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualInterfaceArrayOutput)
}

// VirtualInterfaceMapInput is an input type that accepts VirtualInterfaceMap and VirtualInterfaceMapOutput values.
// You can construct a concrete instance of `VirtualInterfaceMapInput` via:
//
//	VirtualInterfaceMap{ "key": VirtualInterfaceArgs{...} }
type VirtualInterfaceMapInput interface {
	pulumi.Input

	ToVirtualInterfaceMapOutput() VirtualInterfaceMapOutput
	ToVirtualInterfaceMapOutputWithContext(context.Context) VirtualInterfaceMapOutput
}

type VirtualInterfaceMap map[string]VirtualInterfaceInput

func (VirtualInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualInterface)(nil)).Elem()
}

func (i VirtualInterfaceMap) ToVirtualInterfaceMapOutput() VirtualInterfaceMapOutput {
	return i.ToVirtualInterfaceMapOutputWithContext(context.Background())
}

func (i VirtualInterfaceMap) ToVirtualInterfaceMapOutputWithContext(ctx context.Context) VirtualInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualInterfaceMapOutput)
}

type VirtualInterfaceOutput struct{ *pulumi.OutputState }

func (VirtualInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualInterface)(nil)).Elem()
}

func (o VirtualInterfaceOutput) ToVirtualInterfaceOutput() VirtualInterfaceOutput {
	return o
}

func (o VirtualInterfaceOutput) ToVirtualInterfaceOutputWithContext(ctx context.Context) VirtualInterfaceOutput {
	return o
}

// The band width limit of virtual interface,in Mbps.
func (o VirtualInterfaceOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualInterface) pulumi.IntOutput { return v.Bandwidth }).(pulumi.IntOutput)
}

// The BFD detect interval.
func (o VirtualInterfaceOutput) BfdDetectInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualInterface) pulumi.IntPtrOutput { return v.BfdDetectInterval }).(pulumi.IntPtrOutput)
}

// The BFD detect times.
func (o VirtualInterfaceOutput) BfdDetectMultiplier() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualInterface) pulumi.IntPtrOutput { return v.BfdDetectMultiplier }).(pulumi.IntPtrOutput)
}

// The description of virtual interface.
func (o VirtualInterfaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualInterface) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The direct connect connection ID which associated with.
func (o VirtualInterfaceOutput) DirectConnectConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualInterface) pulumi.StringOutput { return v.DirectConnectConnectionId }).(pulumi.StringOutput)
}

// The direct connect gateway ID which associated with.
func (o VirtualInterfaceOutput) DirectConnectGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualInterface) pulumi.StringOutput { return v.DirectConnectGatewayId }).(pulumi.StringOutput)
}

// Whether enable BFD detect.
func (o VirtualInterfaceOutput) EnableBfd() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualInterface) pulumi.BoolPtrOutput { return v.EnableBfd }).(pulumi.BoolPtrOutput)
}

// Whether enable NQA detect.
func (o VirtualInterfaceOutput) EnableNqa() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualInterface) pulumi.BoolPtrOutput { return v.EnableNqa }).(pulumi.BoolPtrOutput)
}

// The local IP that associated with.
func (o VirtualInterfaceOutput) LocalIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualInterface) pulumi.StringOutput { return v.LocalIp }).(pulumi.StringOutput)
}

// The NQA detect interval.
func (o VirtualInterfaceOutput) NqaDetectInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualInterface) pulumi.IntPtrOutput { return v.NqaDetectInterval }).(pulumi.IntPtrOutput)
}

// The NAQ detect times.
func (o VirtualInterfaceOutput) NqaDetectMultiplier() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualInterface) pulumi.IntPtrOutput { return v.NqaDetectMultiplier }).(pulumi.IntPtrOutput)
}

// The peer IP that associated with.
func (o VirtualInterfaceOutput) PeerIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualInterface) pulumi.StringOutput { return v.PeerIp }).(pulumi.StringOutput)
}

// The route type of virtual interface,valid value contains `Static`,`BGP`.
func (o VirtualInterfaceOutput) RouteType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualInterface) pulumi.StringOutput { return v.RouteType }).(pulumi.StringOutput)
}

// The tags that direct connect gateway added.
func (o VirtualInterfaceOutput) Tags() VirtualInterfaceTagArrayOutput {
	return o.ApplyT(func(v *VirtualInterface) VirtualInterfaceTagArrayOutput { return v.Tags }).(VirtualInterfaceTagArrayOutput)
}

// The name of virtual interface.
func (o VirtualInterfaceOutput) VirtualInterfaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualInterface) pulumi.StringPtrOutput { return v.VirtualInterfaceName }).(pulumi.StringPtrOutput)
}

// The VLAN ID used to connect to the local IDC, please ensure that this VLAN ID is not occupied, the value range: 0 ~ 2999.
func (o VirtualInterfaceOutput) VlanId() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualInterface) pulumi.IntOutput { return v.VlanId }).(pulumi.IntOutput)
}

type VirtualInterfaceArrayOutput struct{ *pulumi.OutputState }

func (VirtualInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualInterface)(nil)).Elem()
}

func (o VirtualInterfaceArrayOutput) ToVirtualInterfaceArrayOutput() VirtualInterfaceArrayOutput {
	return o
}

func (o VirtualInterfaceArrayOutput) ToVirtualInterfaceArrayOutputWithContext(ctx context.Context) VirtualInterfaceArrayOutput {
	return o
}

func (o VirtualInterfaceArrayOutput) Index(i pulumi.IntInput) VirtualInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualInterface {
		return vs[0].([]*VirtualInterface)[vs[1].(int)]
	}).(VirtualInterfaceOutput)
}

type VirtualInterfaceMapOutput struct{ *pulumi.OutputState }

func (VirtualInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualInterface)(nil)).Elem()
}

func (o VirtualInterfaceMapOutput) ToVirtualInterfaceMapOutput() VirtualInterfaceMapOutput {
	return o
}

func (o VirtualInterfaceMapOutput) ToVirtualInterfaceMapOutputWithContext(ctx context.Context) VirtualInterfaceMapOutput {
	return o
}

func (o VirtualInterfaceMapOutput) MapIndex(k pulumi.StringInput) VirtualInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualInterface {
		return vs[0].(map[string]*VirtualInterface)[vs[1].(string)]
	}).(VirtualInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualInterfaceInput)(nil)).Elem(), &VirtualInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualInterfaceArrayInput)(nil)).Elem(), VirtualInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualInterfaceMapInput)(nil)).Elem(), VirtualInterfaceMap{})
	pulumi.RegisterOutputType(VirtualInterfaceOutput{})
	pulumi.RegisterOutputType(VirtualInterfaceArrayOutput{})
	pulumi.RegisterOutputType(VirtualInterfaceMapOutput{})
}
