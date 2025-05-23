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

// Provides a resource to manage kafka group
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/kafka"
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
//			_, err = kafka.NewGroup(ctx, "fooGroup", &kafka.GroupArgs{
//				InstanceId:  fooInstance.ID(),
//				GroupId:     pulumi.String("acc-test-group"),
//				Description: pulumi.String("tf-test"),
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
// KafkaGroup can be imported using the instance_id:group_id, e.g.
//
// ```sh
// $ pulumi import volcengine:kafka/group:Group default kafka-****x:groupId
// ```
type Group struct {
	pulumi.CustomResourceState

	// The description of kafka group.
	Description pulumi.StringOutput `pulumi:"description"`
	// The id of kafka group.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The instance id of kafka group.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The state of kafka group.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Group
	err := ctx.RegisterResource("volcengine:kafka/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("volcengine:kafka/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// The description of kafka group.
	Description *string `pulumi:"description"`
	// The id of kafka group.
	GroupId *string `pulumi:"groupId"`
	// The instance id of kafka group.
	InstanceId *string `pulumi:"instanceId"`
	// The state of kafka group.
	State *string `pulumi:"state"`
}

type GroupState struct {
	// The description of kafka group.
	Description pulumi.StringPtrInput
	// The id of kafka group.
	GroupId pulumi.StringPtrInput
	// The instance id of kafka group.
	InstanceId pulumi.StringPtrInput
	// The state of kafka group.
	State pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// The description of kafka group.
	Description *string `pulumi:"description"`
	// The id of kafka group.
	GroupId string `pulumi:"groupId"`
	// The instance id of kafka group.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The description of kafka group.
	Description pulumi.StringPtrInput
	// The id of kafka group.
	GroupId pulumi.StringInput
	// The instance id of kafka group.
	InstanceId pulumi.StringInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//	GroupArray{ GroupArgs{...} }
type GroupArrayInput interface {
	pulumi.Input

	ToGroupArrayOutput() GroupArrayOutput
	ToGroupArrayOutputWithContext(context.Context) GroupArrayOutput
}

type GroupArray []GroupInput

func (GroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (i GroupArray) ToGroupArrayOutput() GroupArrayOutput {
	return i.ToGroupArrayOutputWithContext(context.Background())
}

func (i GroupArray) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupArrayOutput)
}

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//	GroupMap{ "key": GroupArgs{...} }
type GroupMapInput interface {
	pulumi.Input

	ToGroupMapOutput() GroupMapOutput
	ToGroupMapOutputWithContext(context.Context) GroupMapOutput
}

type GroupMap map[string]GroupInput

func (GroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (i GroupMap) ToGroupMapOutput() GroupMapOutput {
	return i.ToGroupMapOutputWithContext(context.Background())
}

func (i GroupMap) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMapOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

// The description of kafka group.
func (o GroupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The id of kafka group.
func (o GroupOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// The instance id of kafka group.
func (o GroupOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The state of kafka group.
func (o GroupOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Group {
		return vs[0].([]*Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Group {
		return vs[0].(map[string]*Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupArrayInput)(nil)).Elem(), GroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMapInput)(nil)).Elem(), GroupMap{})
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}
