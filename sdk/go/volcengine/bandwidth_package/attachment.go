// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bandwidth_package

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage bandwidth package attachment
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/bandwidth_package"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/eip"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooAddress, err := eip.NewAddress(ctx, "fooAddress", &eip.AddressArgs{
//				BillingType: pulumi.String("PostPaidByBandwidth"),
//				Bandwidth:   pulumi.Int(1),
//				Isp:         pulumi.String("BGP"),
//				Description: pulumi.String("acc-test"),
//				ProjectName: pulumi.String("default"),
//			})
//			if err != nil {
//				return err
//			}
//			ipv4BandwidthPackage, err := bandwidth_package.NewBandwidthPackage(ctx, "ipv4BandwidthPackage", &bandwidth_package.BandwidthPackageArgs{
//				BandwidthPackageName: pulumi.String("acc-test-bp"),
//				BillingType:          pulumi.String("PostPaidByBandwidth"),
//				Isp:                  pulumi.String("BGP"),
//				Description:          pulumi.String("acc-test"),
//				Bandwidth:            pulumi.Int(2),
//				Protocol:             pulumi.String("IPv4"),
//				Tags: bandwidth_package.BandwidthPackageTagArray{
//					&bandwidth_package.BandwidthPackageTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = bandwidth_package.NewAttachment(ctx, "ipv4Attachment", &bandwidth_package.AttachmentArgs{
//				AllocationId:       fooAddress.ID(),
//				BandwidthPackageId: ipv4BandwidthPackage.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			fooZones, err := ecs.Zones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			fooImages, err := ecs.Images(ctx, &ecs.ImagesArgs{
//				OsType:         pulumi.StringRef("Linux"),
//				Visibility:     pulumi.StringRef("public"),
//				InstanceTypeId: pulumi.StringRef("ecs.g1.large"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			fooVpc, err := vpc.NewVpc(ctx, "fooVpc", &vpc.VpcArgs{
//				VpcName:    pulumi.String("acc-test-vpc"),
//				CidrBlock:  pulumi.String("172.16.0.0/16"),
//				EnableIpv6: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			fooSubnet, err := vpc.NewSubnet(ctx, "fooSubnet", &vpc.SubnetArgs{
//				SubnetName:    pulumi.String("acc-test-subnet"),
//				CidrBlock:     pulumi.String("172.16.0.0/24"),
//				ZoneId:        *pulumi.String(fooZones.Zones[0].Id),
//				VpcId:         fooVpc.ID(),
//				Ipv6CidrBlock: pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			fooSecurityGroup, err := vpc.NewSecurityGroup(ctx, "fooSecurityGroup", &vpc.SecurityGroupArgs{
//				VpcId:             fooVpc.ID(),
//				SecurityGroupName: pulumi.String("acc-test-security-group"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewIpv6Gateway(ctx, "fooIpv6Gateway", &vpc.Ipv6GatewayArgs{
//				VpcId:       fooVpc.ID(),
//				Description: pulumi.String("test"),
//			})
//			if err != nil {
//				return err
//			}
//			fooInstance, err := ecs.NewInstance(ctx, "fooInstance", &ecs.InstanceArgs{
//				ImageId:            *pulumi.String(fooImages.Images[0].ImageId),
//				InstanceType:       pulumi.String("ecs.g1.large"),
//				InstanceName:       pulumi.String("acc-test-ecs-name"),
//				Password:           pulumi.String("93f0cb0614Aab12"),
//				InstanceChargeType: pulumi.String("PostPaid"),
//				SystemVolumeType:   pulumi.String("ESSD_PL0"),
//				SystemVolumeSize:   pulumi.Int(40),
//				SubnetId:           fooSubnet.ID(),
//				SecurityGroupIds: pulumi.StringArray{
//					fooSecurityGroup.ID(),
//				},
//				Ipv6AddressCount: pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			fooIpv6Addresses := vpc.Ipv6AddressesOutput(ctx, vpc.Ipv6AddressesOutputArgs{
//				AssociatedInstanceId: fooInstance.ID(),
//			}, nil)
//			fooIpv6AddressBandwidth, err := vpc.NewIpv6AddressBandwidth(ctx, "fooIpv6AddressBandwidth", &vpc.Ipv6AddressBandwidthArgs{
//				Ipv6Address: fooIpv6Addresses.ApplyT(func(fooIpv6Addresses vpc.Ipv6AddressesResult) (*string, error) {
//					return &fooIpv6Addresses.Ipv6Addresses[0].Ipv6Address, nil
//				}).(pulumi.StringPtrOutput),
//				BillingType: pulumi.String("PostPaidByBandwidth"),
//				Bandwidth:   pulumi.Int(5),
//			})
//			if err != nil {
//				return err
//			}
//			ipv6, err := bandwidth_package.NewBandwidthPackage(ctx, "ipv6", &bandwidth_package.BandwidthPackageArgs{
//				BandwidthPackageName: pulumi.String("acc-test-bp"),
//				BillingType:          pulumi.String("PostPaidByBandwidth"),
//				Isp:                  pulumi.String("BGP"),
//				Description:          pulumi.String("acc-test"),
//				Bandwidth:            pulumi.Int(2),
//				Protocol:             pulumi.String("IPv6"),
//				Tags: bandwidth_package.BandwidthPackageTagArray{
//					&bandwidth_package.BandwidthPackageTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = bandwidth_package.NewAttachment(ctx, "fooAttachment", &bandwidth_package.AttachmentArgs{
//				AllocationId:       fooIpv6AddressBandwidth.ID(),
//				BandwidthPackageId: ipv6.ID(),
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
// BandwidthPackageAttachment can be imported using the bandwidth package id and eip id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:bandwidth_package/attachment:Attachment default BandwidthPackageId:EipId
//
// ```
type Attachment struct {
	pulumi.CustomResourceState

	// The ID of the public IP or IPv6 public bandwidth to be added to the shared bandwidth package instance.
	AllocationId pulumi.StringOutput `pulumi:"allocationId"`
	// The bandwidth package id.
	BandwidthPackageId pulumi.StringOutput `pulumi:"bandwidthPackageId"`
}

// NewAttachment registers a new resource with the given unique name, arguments, and options.
func NewAttachment(ctx *pulumi.Context,
	name string, args *AttachmentArgs, opts ...pulumi.ResourceOption) (*Attachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllocationId == nil {
		return nil, errors.New("invalid value for required argument 'AllocationId'")
	}
	if args.BandwidthPackageId == nil {
		return nil, errors.New("invalid value for required argument 'BandwidthPackageId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Attachment
	err := ctx.RegisterResource("volcengine:bandwidth_package/attachment:Attachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttachment gets an existing Attachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachmentState, opts ...pulumi.ResourceOption) (*Attachment, error) {
	var resource Attachment
	err := ctx.ReadResource("volcengine:bandwidth_package/attachment:Attachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Attachment resources.
type attachmentState struct {
	// The ID of the public IP or IPv6 public bandwidth to be added to the shared bandwidth package instance.
	AllocationId *string `pulumi:"allocationId"`
	// The bandwidth package id.
	BandwidthPackageId *string `pulumi:"bandwidthPackageId"`
}

type AttachmentState struct {
	// The ID of the public IP or IPv6 public bandwidth to be added to the shared bandwidth package instance.
	AllocationId pulumi.StringPtrInput
	// The bandwidth package id.
	BandwidthPackageId pulumi.StringPtrInput
}

func (AttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachmentState)(nil)).Elem()
}

type attachmentArgs struct {
	// The ID of the public IP or IPv6 public bandwidth to be added to the shared bandwidth package instance.
	AllocationId string `pulumi:"allocationId"`
	// The bandwidth package id.
	BandwidthPackageId string `pulumi:"bandwidthPackageId"`
}

// The set of arguments for constructing a Attachment resource.
type AttachmentArgs struct {
	// The ID of the public IP or IPv6 public bandwidth to be added to the shared bandwidth package instance.
	AllocationId pulumi.StringInput
	// The bandwidth package id.
	BandwidthPackageId pulumi.StringInput
}

func (AttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachmentArgs)(nil)).Elem()
}

type AttachmentInput interface {
	pulumi.Input

	ToAttachmentOutput() AttachmentOutput
	ToAttachmentOutputWithContext(ctx context.Context) AttachmentOutput
}

func (*Attachment) ElementType() reflect.Type {
	return reflect.TypeOf((**Attachment)(nil)).Elem()
}

func (i *Attachment) ToAttachmentOutput() AttachmentOutput {
	return i.ToAttachmentOutputWithContext(context.Background())
}

func (i *Attachment) ToAttachmentOutputWithContext(ctx context.Context) AttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachmentOutput)
}

// AttachmentArrayInput is an input type that accepts AttachmentArray and AttachmentArrayOutput values.
// You can construct a concrete instance of `AttachmentArrayInput` via:
//
//	AttachmentArray{ AttachmentArgs{...} }
type AttachmentArrayInput interface {
	pulumi.Input

	ToAttachmentArrayOutput() AttachmentArrayOutput
	ToAttachmentArrayOutputWithContext(context.Context) AttachmentArrayOutput
}

type AttachmentArray []AttachmentInput

func (AttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Attachment)(nil)).Elem()
}

func (i AttachmentArray) ToAttachmentArrayOutput() AttachmentArrayOutput {
	return i.ToAttachmentArrayOutputWithContext(context.Background())
}

func (i AttachmentArray) ToAttachmentArrayOutputWithContext(ctx context.Context) AttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachmentArrayOutput)
}

// AttachmentMapInput is an input type that accepts AttachmentMap and AttachmentMapOutput values.
// You can construct a concrete instance of `AttachmentMapInput` via:
//
//	AttachmentMap{ "key": AttachmentArgs{...} }
type AttachmentMapInput interface {
	pulumi.Input

	ToAttachmentMapOutput() AttachmentMapOutput
	ToAttachmentMapOutputWithContext(context.Context) AttachmentMapOutput
}

type AttachmentMap map[string]AttachmentInput

func (AttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Attachment)(nil)).Elem()
}

func (i AttachmentMap) ToAttachmentMapOutput() AttachmentMapOutput {
	return i.ToAttachmentMapOutputWithContext(context.Background())
}

func (i AttachmentMap) ToAttachmentMapOutputWithContext(ctx context.Context) AttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachmentMapOutput)
}

type AttachmentOutput struct{ *pulumi.OutputState }

func (AttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Attachment)(nil)).Elem()
}

func (o AttachmentOutput) ToAttachmentOutput() AttachmentOutput {
	return o
}

func (o AttachmentOutput) ToAttachmentOutputWithContext(ctx context.Context) AttachmentOutput {
	return o
}

// The ID of the public IP or IPv6 public bandwidth to be added to the shared bandwidth package instance.
func (o AttachmentOutput) AllocationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Attachment) pulumi.StringOutput { return v.AllocationId }).(pulumi.StringOutput)
}

// The bandwidth package id.
func (o AttachmentOutput) BandwidthPackageId() pulumi.StringOutput {
	return o.ApplyT(func(v *Attachment) pulumi.StringOutput { return v.BandwidthPackageId }).(pulumi.StringOutput)
}

type AttachmentArrayOutput struct{ *pulumi.OutputState }

func (AttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Attachment)(nil)).Elem()
}

func (o AttachmentArrayOutput) ToAttachmentArrayOutput() AttachmentArrayOutput {
	return o
}

func (o AttachmentArrayOutput) ToAttachmentArrayOutputWithContext(ctx context.Context) AttachmentArrayOutput {
	return o
}

func (o AttachmentArrayOutput) Index(i pulumi.IntInput) AttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Attachment {
		return vs[0].([]*Attachment)[vs[1].(int)]
	}).(AttachmentOutput)
}

type AttachmentMapOutput struct{ *pulumi.OutputState }

func (AttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Attachment)(nil)).Elem()
}

func (o AttachmentMapOutput) ToAttachmentMapOutput() AttachmentMapOutput {
	return o
}

func (o AttachmentMapOutput) ToAttachmentMapOutputWithContext(ctx context.Context) AttachmentMapOutput {
	return o
}

func (o AttachmentMapOutput) MapIndex(k pulumi.StringInput) AttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Attachment {
		return vs[0].(map[string]*Attachment)[vs[1].(string)]
	}).(AttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AttachmentInput)(nil)).Elem(), &Attachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttachmentArrayInput)(nil)).Elem(), AttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttachmentMapInput)(nil)).Elem(), AttachmentMap{})
	pulumi.RegisterOutputType(AttachmentOutput{})
	pulumi.RegisterOutputType(AttachmentArrayOutput{})
	pulumi.RegisterOutputType(AttachmentMapOutput{})
}