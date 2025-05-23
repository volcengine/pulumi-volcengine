// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nat

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage dnat entry
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/eip"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/nat"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooZones, err := ecs.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			fooVpc, err := vpc.NewVpc(ctx, "fooVpc", &vpc.VpcArgs{
//				VpcName:   pulumi.String("acc-test-vpc"),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			fooSubnet, err := vpc.NewSubnet(ctx, "fooSubnet", &vpc.SubnetArgs{
//				SubnetName: pulumi.String("acc-test-subnet"),
//				CidrBlock:  pulumi.String("172.16.0.0/24"),
//				ZoneId:     pulumi.String(fooZones.Zones[0].Id),
//				VpcId:      fooVpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			fooGateway, err := nat.NewGateway(ctx, "fooGateway", &nat.GatewayArgs{
//				VpcId:          fooVpc.ID(),
//				SubnetId:       fooSubnet.ID(),
//				Spec:           pulumi.String("Small"),
//				NatGatewayName: pulumi.String("acc-test-ng"),
//				Description:    pulumi.String("acc-test"),
//				BillingType:    pulumi.String("PostPaid"),
//				ProjectName:    pulumi.String("default"),
//				Tags: nat.GatewayTagArray{
//					&nat.GatewayTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			fooAddress, err := eip.NewAddress(ctx, "fooAddress", &eip.AddressArgs{
//				Description: pulumi.String("acc-test"),
//				Bandwidth:   pulumi.Int(1),
//				BillingType: pulumi.String("PostPaidByBandwidth"),
//				Isp:         pulumi.String("BGP"),
//			})
//			if err != nil {
//				return err
//			}
//			fooAssociate, err := eip.NewAssociate(ctx, "fooAssociate", &eip.AssociateArgs{
//				AllocationId: fooAddress.ID(),
//				InstanceId:   fooGateway.ID(),
//				InstanceType: pulumi.String("Nat"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = nat.NewDnatEntry(ctx, "fooDnatEntry", &nat.DnatEntryArgs{
//				DnatEntryName: pulumi.String("acc-test-dnat-entry"),
//				ExternalIp:    fooAddress.EipAddress,
//				ExternalPort:  pulumi.String("80"),
//				InternalIp:    pulumi.String("172.16.0.10"),
//				InternalPort:  pulumi.String("80"),
//				NatGatewayId:  fooGateway.ID(),
//				Protocol:      pulumi.String("tcp"),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				fooAssociate,
//			}))
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
// Dnat entry can be imported using the id, e.g.
//
// ```sh
// $ pulumi import volcengine:nat/dnatEntry:DnatEntry default dnat-3fvhk47kf56****
// ```
type DnatEntry struct {
	pulumi.CustomResourceState

	// The id of the DNAT rule.
	DnatEntryId pulumi.StringOutput `pulumi:"dnatEntryId"`
	// The name of the DNAT rule.
	DnatEntryName pulumi.StringPtrOutput `pulumi:"dnatEntryName"`
	// Provides the public IP address for public network access.
	ExternalIp pulumi.StringOutput `pulumi:"externalIp"`
	// The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
	ExternalPort pulumi.StringOutput `pulumi:"externalPort"`
	// Provides the internal IP address.
	InternalIp pulumi.StringOutput `pulumi:"internalIp"`
	// The port or port segment on which the cloud server instance provides services to the public network.
	InternalPort pulumi.StringOutput `pulumi:"internalPort"`
	// The id of the nat gateway to which the entry belongs.
	NatGatewayId pulumi.StringOutput `pulumi:"natGatewayId"`
	// The network protocol.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
}

// NewDnatEntry registers a new resource with the given unique name, arguments, and options.
func NewDnatEntry(ctx *pulumi.Context,
	name string, args *DnatEntryArgs, opts ...pulumi.ResourceOption) (*DnatEntry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExternalIp == nil {
		return nil, errors.New("invalid value for required argument 'ExternalIp'")
	}
	if args.ExternalPort == nil {
		return nil, errors.New("invalid value for required argument 'ExternalPort'")
	}
	if args.InternalIp == nil {
		return nil, errors.New("invalid value for required argument 'InternalIp'")
	}
	if args.InternalPort == nil {
		return nil, errors.New("invalid value for required argument 'InternalPort'")
	}
	if args.NatGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'NatGatewayId'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DnatEntry
	err := ctx.RegisterResource("volcengine:nat/dnatEntry:DnatEntry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnatEntry gets an existing DnatEntry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnatEntry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnatEntryState, opts ...pulumi.ResourceOption) (*DnatEntry, error) {
	var resource DnatEntry
	err := ctx.ReadResource("volcengine:nat/dnatEntry:DnatEntry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnatEntry resources.
type dnatEntryState struct {
	// The id of the DNAT rule.
	DnatEntryId *string `pulumi:"dnatEntryId"`
	// The name of the DNAT rule.
	DnatEntryName *string `pulumi:"dnatEntryName"`
	// Provides the public IP address for public network access.
	ExternalIp *string `pulumi:"externalIp"`
	// The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
	ExternalPort *string `pulumi:"externalPort"`
	// Provides the internal IP address.
	InternalIp *string `pulumi:"internalIp"`
	// The port or port segment on which the cloud server instance provides services to the public network.
	InternalPort *string `pulumi:"internalPort"`
	// The id of the nat gateway to which the entry belongs.
	NatGatewayId *string `pulumi:"natGatewayId"`
	// The network protocol.
	Protocol *string `pulumi:"protocol"`
}

type DnatEntryState struct {
	// The id of the DNAT rule.
	DnatEntryId pulumi.StringPtrInput
	// The name of the DNAT rule.
	DnatEntryName pulumi.StringPtrInput
	// Provides the public IP address for public network access.
	ExternalIp pulumi.StringPtrInput
	// The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
	ExternalPort pulumi.StringPtrInput
	// Provides the internal IP address.
	InternalIp pulumi.StringPtrInput
	// The port or port segment on which the cloud server instance provides services to the public network.
	InternalPort pulumi.StringPtrInput
	// The id of the nat gateway to which the entry belongs.
	NatGatewayId pulumi.StringPtrInput
	// The network protocol.
	Protocol pulumi.StringPtrInput
}

func (DnatEntryState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnatEntryState)(nil)).Elem()
}

type dnatEntryArgs struct {
	// The name of the DNAT rule.
	DnatEntryName *string `pulumi:"dnatEntryName"`
	// Provides the public IP address for public network access.
	ExternalIp string `pulumi:"externalIp"`
	// The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
	ExternalPort string `pulumi:"externalPort"`
	// Provides the internal IP address.
	InternalIp string `pulumi:"internalIp"`
	// The port or port segment on which the cloud server instance provides services to the public network.
	InternalPort string `pulumi:"internalPort"`
	// The id of the nat gateway to which the entry belongs.
	NatGatewayId string `pulumi:"natGatewayId"`
	// The network protocol.
	Protocol string `pulumi:"protocol"`
}

// The set of arguments for constructing a DnatEntry resource.
type DnatEntryArgs struct {
	// The name of the DNAT rule.
	DnatEntryName pulumi.StringPtrInput
	// Provides the public IP address for public network access.
	ExternalIp pulumi.StringInput
	// The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
	ExternalPort pulumi.StringInput
	// Provides the internal IP address.
	InternalIp pulumi.StringInput
	// The port or port segment on which the cloud server instance provides services to the public network.
	InternalPort pulumi.StringInput
	// The id of the nat gateway to which the entry belongs.
	NatGatewayId pulumi.StringInput
	// The network protocol.
	Protocol pulumi.StringInput
}

func (DnatEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnatEntryArgs)(nil)).Elem()
}

type DnatEntryInput interface {
	pulumi.Input

	ToDnatEntryOutput() DnatEntryOutput
	ToDnatEntryOutputWithContext(ctx context.Context) DnatEntryOutput
}

func (*DnatEntry) ElementType() reflect.Type {
	return reflect.TypeOf((**DnatEntry)(nil)).Elem()
}

func (i *DnatEntry) ToDnatEntryOutput() DnatEntryOutput {
	return i.ToDnatEntryOutputWithContext(context.Background())
}

func (i *DnatEntry) ToDnatEntryOutputWithContext(ctx context.Context) DnatEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnatEntryOutput)
}

// DnatEntryArrayInput is an input type that accepts DnatEntryArray and DnatEntryArrayOutput values.
// You can construct a concrete instance of `DnatEntryArrayInput` via:
//
//	DnatEntryArray{ DnatEntryArgs{...} }
type DnatEntryArrayInput interface {
	pulumi.Input

	ToDnatEntryArrayOutput() DnatEntryArrayOutput
	ToDnatEntryArrayOutputWithContext(context.Context) DnatEntryArrayOutput
}

type DnatEntryArray []DnatEntryInput

func (DnatEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnatEntry)(nil)).Elem()
}

func (i DnatEntryArray) ToDnatEntryArrayOutput() DnatEntryArrayOutput {
	return i.ToDnatEntryArrayOutputWithContext(context.Background())
}

func (i DnatEntryArray) ToDnatEntryArrayOutputWithContext(ctx context.Context) DnatEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnatEntryArrayOutput)
}

// DnatEntryMapInput is an input type that accepts DnatEntryMap and DnatEntryMapOutput values.
// You can construct a concrete instance of `DnatEntryMapInput` via:
//
//	DnatEntryMap{ "key": DnatEntryArgs{...} }
type DnatEntryMapInput interface {
	pulumi.Input

	ToDnatEntryMapOutput() DnatEntryMapOutput
	ToDnatEntryMapOutputWithContext(context.Context) DnatEntryMapOutput
}

type DnatEntryMap map[string]DnatEntryInput

func (DnatEntryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnatEntry)(nil)).Elem()
}

func (i DnatEntryMap) ToDnatEntryMapOutput() DnatEntryMapOutput {
	return i.ToDnatEntryMapOutputWithContext(context.Background())
}

func (i DnatEntryMap) ToDnatEntryMapOutputWithContext(ctx context.Context) DnatEntryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnatEntryMapOutput)
}

type DnatEntryOutput struct{ *pulumi.OutputState }

func (DnatEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnatEntry)(nil)).Elem()
}

func (o DnatEntryOutput) ToDnatEntryOutput() DnatEntryOutput {
	return o
}

func (o DnatEntryOutput) ToDnatEntryOutputWithContext(ctx context.Context) DnatEntryOutput {
	return o
}

// The id of the DNAT rule.
func (o DnatEntryOutput) DnatEntryId() pulumi.StringOutput {
	return o.ApplyT(func(v *DnatEntry) pulumi.StringOutput { return v.DnatEntryId }).(pulumi.StringOutput)
}

// The name of the DNAT rule.
func (o DnatEntryOutput) DnatEntryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnatEntry) pulumi.StringPtrOutput { return v.DnatEntryName }).(pulumi.StringPtrOutput)
}

// Provides the public IP address for public network access.
func (o DnatEntryOutput) ExternalIp() pulumi.StringOutput {
	return o.ApplyT(func(v *DnatEntry) pulumi.StringOutput { return v.ExternalIp }).(pulumi.StringOutput)
}

// The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
func (o DnatEntryOutput) ExternalPort() pulumi.StringOutput {
	return o.ApplyT(func(v *DnatEntry) pulumi.StringOutput { return v.ExternalPort }).(pulumi.StringOutput)
}

// Provides the internal IP address.
func (o DnatEntryOutput) InternalIp() pulumi.StringOutput {
	return o.ApplyT(func(v *DnatEntry) pulumi.StringOutput { return v.InternalIp }).(pulumi.StringOutput)
}

// The port or port segment on which the cloud server instance provides services to the public network.
func (o DnatEntryOutput) InternalPort() pulumi.StringOutput {
	return o.ApplyT(func(v *DnatEntry) pulumi.StringOutput { return v.InternalPort }).(pulumi.StringOutput)
}

// The id of the nat gateway to which the entry belongs.
func (o DnatEntryOutput) NatGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *DnatEntry) pulumi.StringOutput { return v.NatGatewayId }).(pulumi.StringOutput)
}

// The network protocol.
func (o DnatEntryOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *DnatEntry) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

type DnatEntryArrayOutput struct{ *pulumi.OutputState }

func (DnatEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnatEntry)(nil)).Elem()
}

func (o DnatEntryArrayOutput) ToDnatEntryArrayOutput() DnatEntryArrayOutput {
	return o
}

func (o DnatEntryArrayOutput) ToDnatEntryArrayOutputWithContext(ctx context.Context) DnatEntryArrayOutput {
	return o
}

func (o DnatEntryArrayOutput) Index(i pulumi.IntInput) DnatEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnatEntry {
		return vs[0].([]*DnatEntry)[vs[1].(int)]
	}).(DnatEntryOutput)
}

type DnatEntryMapOutput struct{ *pulumi.OutputState }

func (DnatEntryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnatEntry)(nil)).Elem()
}

func (o DnatEntryMapOutput) ToDnatEntryMapOutput() DnatEntryMapOutput {
	return o
}

func (o DnatEntryMapOutput) ToDnatEntryMapOutputWithContext(ctx context.Context) DnatEntryMapOutput {
	return o
}

func (o DnatEntryMapOutput) MapIndex(k pulumi.StringInput) DnatEntryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnatEntry {
		return vs[0].(map[string]*DnatEntry)[vs[1].(string)]
	}).(DnatEntryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnatEntryInput)(nil)).Elem(), &DnatEntry{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnatEntryArrayInput)(nil)).Elem(), DnatEntryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnatEntryMapInput)(nil)).Elem(), DnatEntryMap{})
	pulumi.RegisterOutputType(DnatEntryOutput{})
	pulumi.RegisterOutputType(DnatEntryArrayOutput{})
	pulumi.RegisterOutputType(DnatEntryMapOutput{})
}
