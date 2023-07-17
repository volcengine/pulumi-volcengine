// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage privatelink vpc endpoint zone
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/privatelink"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := privatelink.NewVpcEndpointZone(ctx, "foo", &privatelink.VpcEndpointZoneArgs{
//				EndpointId:       pulumi.String("ep-2byz5nlkimc5c2dx0ef9g****"),
//				PrivateIpAddress: pulumi.String("172.16.0.251"),
//				SubnetId:         pulumi.String("subnet-2bz47q19zhx4w2dx0eevn****"),
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
// VpcEndpointZone can be imported using the endpointId:subnetId, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:privatelink/vpcEndpointZone:VpcEndpointZone default ep-3rel75r081l345zsk2i59****:subnet-2bz47q19zhx4w2dx0eevn****
//
// ```
type VpcEndpointZone struct {
	pulumi.CustomResourceState

	// The endpoint id of vpc endpoint zone.
	EndpointId pulumi.StringOutput `pulumi:"endpointId"`
	// The network interface id of vpc endpoint.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The private ip address of vpc endpoint zone.
	PrivateIpAddress pulumi.StringOutput `pulumi:"privateIpAddress"`
	// The subnet id of vpc endpoint zone.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// The domain of vpc endpoint zone.
	ZoneDomain pulumi.StringOutput `pulumi:"zoneDomain"`
	// The Id of vpc endpoint zone.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
	// The status of vpc endpoint zone.
	ZoneStatus pulumi.StringOutput `pulumi:"zoneStatus"`
}

// NewVpcEndpointZone registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpointZone(ctx *pulumi.Context,
	name string, args *VpcEndpointZoneArgs, opts ...pulumi.ResourceOption) (*VpcEndpointZone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointId == nil {
		return nil, errors.New("invalid value for required argument 'EndpointId'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource VpcEndpointZone
	err := ctx.RegisterResource("volcengine:privatelink/vpcEndpointZone:VpcEndpointZone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpointZone gets an existing VpcEndpointZone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpointZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointZoneState, opts ...pulumi.ResourceOption) (*VpcEndpointZone, error) {
	var resource VpcEndpointZone
	err := ctx.ReadResource("volcengine:privatelink/vpcEndpointZone:VpcEndpointZone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpointZone resources.
type vpcEndpointZoneState struct {
	// The endpoint id of vpc endpoint zone.
	EndpointId *string `pulumi:"endpointId"`
	// The network interface id of vpc endpoint.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The private ip address of vpc endpoint zone.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// The subnet id of vpc endpoint zone.
	SubnetId *string `pulumi:"subnetId"`
	// The domain of vpc endpoint zone.
	ZoneDomain *string `pulumi:"zoneDomain"`
	// The Id of vpc endpoint zone.
	ZoneId *string `pulumi:"zoneId"`
	// The status of vpc endpoint zone.
	ZoneStatus *string `pulumi:"zoneStatus"`
}

type VpcEndpointZoneState struct {
	// The endpoint id of vpc endpoint zone.
	EndpointId pulumi.StringPtrInput
	// The network interface id of vpc endpoint.
	NetworkInterfaceId pulumi.StringPtrInput
	// The private ip address of vpc endpoint zone.
	PrivateIpAddress pulumi.StringPtrInput
	// The subnet id of vpc endpoint zone.
	SubnetId pulumi.StringPtrInput
	// The domain of vpc endpoint zone.
	ZoneDomain pulumi.StringPtrInput
	// The Id of vpc endpoint zone.
	ZoneId pulumi.StringPtrInput
	// The status of vpc endpoint zone.
	ZoneStatus pulumi.StringPtrInput
}

func (VpcEndpointZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointZoneState)(nil)).Elem()
}

type vpcEndpointZoneArgs struct {
	// The endpoint id of vpc endpoint zone.
	EndpointId string `pulumi:"endpointId"`
	// The private ip address of vpc endpoint zone.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// The subnet id of vpc endpoint zone.
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a VpcEndpointZone resource.
type VpcEndpointZoneArgs struct {
	// The endpoint id of vpc endpoint zone.
	EndpointId pulumi.StringInput
	// The private ip address of vpc endpoint zone.
	PrivateIpAddress pulumi.StringPtrInput
	// The subnet id of vpc endpoint zone.
	SubnetId pulumi.StringInput
}

func (VpcEndpointZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointZoneArgs)(nil)).Elem()
}

type VpcEndpointZoneInput interface {
	pulumi.Input

	ToVpcEndpointZoneOutput() VpcEndpointZoneOutput
	ToVpcEndpointZoneOutputWithContext(ctx context.Context) VpcEndpointZoneOutput
}

func (*VpcEndpointZone) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointZone)(nil)).Elem()
}

func (i *VpcEndpointZone) ToVpcEndpointZoneOutput() VpcEndpointZoneOutput {
	return i.ToVpcEndpointZoneOutputWithContext(context.Background())
}

func (i *VpcEndpointZone) ToVpcEndpointZoneOutputWithContext(ctx context.Context) VpcEndpointZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointZoneOutput)
}

// VpcEndpointZoneArrayInput is an input type that accepts VpcEndpointZoneArray and VpcEndpointZoneArrayOutput values.
// You can construct a concrete instance of `VpcEndpointZoneArrayInput` via:
//
//	VpcEndpointZoneArray{ VpcEndpointZoneArgs{...} }
type VpcEndpointZoneArrayInput interface {
	pulumi.Input

	ToVpcEndpointZoneArrayOutput() VpcEndpointZoneArrayOutput
	ToVpcEndpointZoneArrayOutputWithContext(context.Context) VpcEndpointZoneArrayOutput
}

type VpcEndpointZoneArray []VpcEndpointZoneInput

func (VpcEndpointZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointZone)(nil)).Elem()
}

func (i VpcEndpointZoneArray) ToVpcEndpointZoneArrayOutput() VpcEndpointZoneArrayOutput {
	return i.ToVpcEndpointZoneArrayOutputWithContext(context.Background())
}

func (i VpcEndpointZoneArray) ToVpcEndpointZoneArrayOutputWithContext(ctx context.Context) VpcEndpointZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointZoneArrayOutput)
}

// VpcEndpointZoneMapInput is an input type that accepts VpcEndpointZoneMap and VpcEndpointZoneMapOutput values.
// You can construct a concrete instance of `VpcEndpointZoneMapInput` via:
//
//	VpcEndpointZoneMap{ "key": VpcEndpointZoneArgs{...} }
type VpcEndpointZoneMapInput interface {
	pulumi.Input

	ToVpcEndpointZoneMapOutput() VpcEndpointZoneMapOutput
	ToVpcEndpointZoneMapOutputWithContext(context.Context) VpcEndpointZoneMapOutput
}

type VpcEndpointZoneMap map[string]VpcEndpointZoneInput

func (VpcEndpointZoneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointZone)(nil)).Elem()
}

func (i VpcEndpointZoneMap) ToVpcEndpointZoneMapOutput() VpcEndpointZoneMapOutput {
	return i.ToVpcEndpointZoneMapOutputWithContext(context.Background())
}

func (i VpcEndpointZoneMap) ToVpcEndpointZoneMapOutputWithContext(ctx context.Context) VpcEndpointZoneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointZoneMapOutput)
}

type VpcEndpointZoneOutput struct{ *pulumi.OutputState }

func (VpcEndpointZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointZone)(nil)).Elem()
}

func (o VpcEndpointZoneOutput) ToVpcEndpointZoneOutput() VpcEndpointZoneOutput {
	return o
}

func (o VpcEndpointZoneOutput) ToVpcEndpointZoneOutputWithContext(ctx context.Context) VpcEndpointZoneOutput {
	return o
}

// The endpoint id of vpc endpoint zone.
func (o VpcEndpointZoneOutput) EndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointZone) pulumi.StringOutput { return v.EndpointId }).(pulumi.StringOutput)
}

// The network interface id of vpc endpoint.
func (o VpcEndpointZoneOutput) NetworkInterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointZone) pulumi.StringOutput { return v.NetworkInterfaceId }).(pulumi.StringOutput)
}

// The private ip address of vpc endpoint zone.
func (o VpcEndpointZoneOutput) PrivateIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointZone) pulumi.StringOutput { return v.PrivateIpAddress }).(pulumi.StringOutput)
}

// The subnet id of vpc endpoint zone.
func (o VpcEndpointZoneOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointZone) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// The domain of vpc endpoint zone.
func (o VpcEndpointZoneOutput) ZoneDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointZone) pulumi.StringOutput { return v.ZoneDomain }).(pulumi.StringOutput)
}

// The Id of vpc endpoint zone.
func (o VpcEndpointZoneOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointZone) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

// The status of vpc endpoint zone.
func (o VpcEndpointZoneOutput) ZoneStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointZone) pulumi.StringOutput { return v.ZoneStatus }).(pulumi.StringOutput)
}

type VpcEndpointZoneArrayOutput struct{ *pulumi.OutputState }

func (VpcEndpointZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointZone)(nil)).Elem()
}

func (o VpcEndpointZoneArrayOutput) ToVpcEndpointZoneArrayOutput() VpcEndpointZoneArrayOutput {
	return o
}

func (o VpcEndpointZoneArrayOutput) ToVpcEndpointZoneArrayOutputWithContext(ctx context.Context) VpcEndpointZoneArrayOutput {
	return o
}

func (o VpcEndpointZoneArrayOutput) Index(i pulumi.IntInput) VpcEndpointZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcEndpointZone {
		return vs[0].([]*VpcEndpointZone)[vs[1].(int)]
	}).(VpcEndpointZoneOutput)
}

type VpcEndpointZoneMapOutput struct{ *pulumi.OutputState }

func (VpcEndpointZoneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointZone)(nil)).Elem()
}

func (o VpcEndpointZoneMapOutput) ToVpcEndpointZoneMapOutput() VpcEndpointZoneMapOutput {
	return o
}

func (o VpcEndpointZoneMapOutput) ToVpcEndpointZoneMapOutputWithContext(ctx context.Context) VpcEndpointZoneMapOutput {
	return o
}

func (o VpcEndpointZoneMapOutput) MapIndex(k pulumi.StringInput) VpcEndpointZoneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcEndpointZone {
		return vs[0].(map[string]*VpcEndpointZone)[vs[1].(string)]
	}).(VpcEndpointZoneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointZoneInput)(nil)).Elem(), &VpcEndpointZone{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointZoneArrayInput)(nil)).Elem(), VpcEndpointZoneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointZoneMapInput)(nil)).Elem(), VpcEndpointZoneMap{})
	pulumi.RegisterOutputType(VpcEndpointZoneOutput{})
	pulumi.RegisterOutputType(VpcEndpointZoneArrayOutput{})
	pulumi.RegisterOutputType(VpcEndpointZoneMapOutput{})
}