// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kafka

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage kafka public address
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
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/kafka"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooZones, err := ecs.Zones(ctx, nil, nil)
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
//			fooInstance, err := kafka.NewInstance(ctx, "fooInstance", &kafka.InstanceArgs{
//				InstanceName:        pulumi.String("acc-test-kafka"),
//				InstanceDescription: pulumi.String("tf-test"),
//				Version:             pulumi.String("2.2.2"),
//				ComputeSpec:         pulumi.String("kafka.20xrate.hw"),
//				SubnetId:            fooSubnet.ID(),
//				UserName:            pulumi.String("tf-user"),
//				UserPassword:        pulumi.String("tf-pass!@q1"),
//				ChargeType:          pulumi.String("PostPaid"),
//				StorageSpace:        pulumi.Int(300),
//				PartitionNumber:     pulumi.Int(350),
//				ProjectName:         pulumi.String("default"),
//				Tags: kafka.InstanceTagArray{
//					&kafka.InstanceTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
//					},
//				},
//				Parameters: kafka.InstanceParameterArray{
//					&kafka.InstanceParameterArgs{
//						ParameterName:  pulumi.String("MessageMaxByte"),
//						ParameterValue: pulumi.String("12"),
//					},
//					&kafka.InstanceParameterArgs{
//						ParameterName:  pulumi.String("LogRetentionHours"),
//						ParameterValue: pulumi.String("70"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			fooAddress, err := eip.NewAddress(ctx, "fooAddress", &eip.AddressArgs{
//				BillingType: pulumi.String("PostPaidByBandwidth"),
//				Bandwidth:   pulumi.Int(1),
//				Isp:         pulumi.String("BGP"),
//				Description: pulumi.String("tf-test"),
//				ProjectName: pulumi.String("default"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = kafka.NewPublicAddress(ctx, "fooPublicAddress", &kafka.PublicAddressArgs{
//				InstanceId: fooInstance.ID(),
//				EipId:      fooAddress.ID(),
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
// KafkaPublicAddress can be imported using the instance_id:eip_id, e.g.
//
// ```sh
// $ pulumi import volcengine:kafka/publicAddress:PublicAddress default instance_id:eip_id
// ```
type PublicAddress struct {
	pulumi.CustomResourceState

	// The id of eip.
	EipId pulumi.StringOutput `pulumi:"eipId"`
	// The endpoint type of instance.
	EndpointType pulumi.StringOutput `pulumi:"endpointType"`
	// The id of kafka instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The network type of instance.
	NetworkType pulumi.StringOutput `pulumi:"networkType"`
	// The public endpoint of instance.
	PublicEndpoint pulumi.StringOutput `pulumi:"publicEndpoint"`
}

// NewPublicAddress registers a new resource with the given unique name, arguments, and options.
func NewPublicAddress(ctx *pulumi.Context,
	name string, args *PublicAddressArgs, opts ...pulumi.ResourceOption) (*PublicAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EipId == nil {
		return nil, errors.New("invalid value for required argument 'EipId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PublicAddress
	err := ctx.RegisterResource("volcengine:kafka/publicAddress:PublicAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicAddress gets an existing PublicAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicAddressState, opts ...pulumi.ResourceOption) (*PublicAddress, error) {
	var resource PublicAddress
	err := ctx.ReadResource("volcengine:kafka/publicAddress:PublicAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicAddress resources.
type publicAddressState struct {
	// The id of eip.
	EipId *string `pulumi:"eipId"`
	// The endpoint type of instance.
	EndpointType *string `pulumi:"endpointType"`
	// The id of kafka instance.
	InstanceId *string `pulumi:"instanceId"`
	// The network type of instance.
	NetworkType *string `pulumi:"networkType"`
	// The public endpoint of instance.
	PublicEndpoint *string `pulumi:"publicEndpoint"`
}

type PublicAddressState struct {
	// The id of eip.
	EipId pulumi.StringPtrInput
	// The endpoint type of instance.
	EndpointType pulumi.StringPtrInput
	// The id of kafka instance.
	InstanceId pulumi.StringPtrInput
	// The network type of instance.
	NetworkType pulumi.StringPtrInput
	// The public endpoint of instance.
	PublicEndpoint pulumi.StringPtrInput
}

func (PublicAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicAddressState)(nil)).Elem()
}

type publicAddressArgs struct {
	// The id of eip.
	EipId string `pulumi:"eipId"`
	// The id of kafka instance.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a PublicAddress resource.
type PublicAddressArgs struct {
	// The id of eip.
	EipId pulumi.StringInput
	// The id of kafka instance.
	InstanceId pulumi.StringInput
}

func (PublicAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicAddressArgs)(nil)).Elem()
}

type PublicAddressInput interface {
	pulumi.Input

	ToPublicAddressOutput() PublicAddressOutput
	ToPublicAddressOutputWithContext(ctx context.Context) PublicAddressOutput
}

func (*PublicAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicAddress)(nil)).Elem()
}

func (i *PublicAddress) ToPublicAddressOutput() PublicAddressOutput {
	return i.ToPublicAddressOutputWithContext(context.Background())
}

func (i *PublicAddress) ToPublicAddressOutputWithContext(ctx context.Context) PublicAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicAddressOutput)
}

// PublicAddressArrayInput is an input type that accepts PublicAddressArray and PublicAddressArrayOutput values.
// You can construct a concrete instance of `PublicAddressArrayInput` via:
//
//	PublicAddressArray{ PublicAddressArgs{...} }
type PublicAddressArrayInput interface {
	pulumi.Input

	ToPublicAddressArrayOutput() PublicAddressArrayOutput
	ToPublicAddressArrayOutputWithContext(context.Context) PublicAddressArrayOutput
}

type PublicAddressArray []PublicAddressInput

func (PublicAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicAddress)(nil)).Elem()
}

func (i PublicAddressArray) ToPublicAddressArrayOutput() PublicAddressArrayOutput {
	return i.ToPublicAddressArrayOutputWithContext(context.Background())
}

func (i PublicAddressArray) ToPublicAddressArrayOutputWithContext(ctx context.Context) PublicAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicAddressArrayOutput)
}

// PublicAddressMapInput is an input type that accepts PublicAddressMap and PublicAddressMapOutput values.
// You can construct a concrete instance of `PublicAddressMapInput` via:
//
//	PublicAddressMap{ "key": PublicAddressArgs{...} }
type PublicAddressMapInput interface {
	pulumi.Input

	ToPublicAddressMapOutput() PublicAddressMapOutput
	ToPublicAddressMapOutputWithContext(context.Context) PublicAddressMapOutput
}

type PublicAddressMap map[string]PublicAddressInput

func (PublicAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicAddress)(nil)).Elem()
}

func (i PublicAddressMap) ToPublicAddressMapOutput() PublicAddressMapOutput {
	return i.ToPublicAddressMapOutputWithContext(context.Background())
}

func (i PublicAddressMap) ToPublicAddressMapOutputWithContext(ctx context.Context) PublicAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicAddressMapOutput)
}

type PublicAddressOutput struct{ *pulumi.OutputState }

func (PublicAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicAddress)(nil)).Elem()
}

func (o PublicAddressOutput) ToPublicAddressOutput() PublicAddressOutput {
	return o
}

func (o PublicAddressOutput) ToPublicAddressOutputWithContext(ctx context.Context) PublicAddressOutput {
	return o
}

// The id of eip.
func (o PublicAddressOutput) EipId() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicAddress) pulumi.StringOutput { return v.EipId }).(pulumi.StringOutput)
}

// The endpoint type of instance.
func (o PublicAddressOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicAddress) pulumi.StringOutput { return v.EndpointType }).(pulumi.StringOutput)
}

// The id of kafka instance.
func (o PublicAddressOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicAddress) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The network type of instance.
func (o PublicAddressOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicAddress) pulumi.StringOutput { return v.NetworkType }).(pulumi.StringOutput)
}

// The public endpoint of instance.
func (o PublicAddressOutput) PublicEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicAddress) pulumi.StringOutput { return v.PublicEndpoint }).(pulumi.StringOutput)
}

type PublicAddressArrayOutput struct{ *pulumi.OutputState }

func (PublicAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicAddress)(nil)).Elem()
}

func (o PublicAddressArrayOutput) ToPublicAddressArrayOutput() PublicAddressArrayOutput {
	return o
}

func (o PublicAddressArrayOutput) ToPublicAddressArrayOutputWithContext(ctx context.Context) PublicAddressArrayOutput {
	return o
}

func (o PublicAddressArrayOutput) Index(i pulumi.IntInput) PublicAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PublicAddress {
		return vs[0].([]*PublicAddress)[vs[1].(int)]
	}).(PublicAddressOutput)
}

type PublicAddressMapOutput struct{ *pulumi.OutputState }

func (PublicAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicAddress)(nil)).Elem()
}

func (o PublicAddressMapOutput) ToPublicAddressMapOutput() PublicAddressMapOutput {
	return o
}

func (o PublicAddressMapOutput) ToPublicAddressMapOutputWithContext(ctx context.Context) PublicAddressMapOutput {
	return o
}

func (o PublicAddressMapOutput) MapIndex(k pulumi.StringInput) PublicAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PublicAddress {
		return vs[0].(map[string]*PublicAddress)[vs[1].(string)]
	}).(PublicAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublicAddressInput)(nil)).Elem(), &PublicAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicAddressArrayInput)(nil)).Elem(), PublicAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicAddressMapInput)(nil)).Elem(), PublicAddressMap{})
	pulumi.RegisterOutputType(PublicAddressOutput{})
	pulumi.RegisterOutputType(PublicAddressArrayOutput{})
	pulumi.RegisterOutputType(PublicAddressMapOutput{})
}