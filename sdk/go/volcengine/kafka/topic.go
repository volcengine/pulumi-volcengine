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

// Provides a resource to manage kafka topic
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
//			fooSaslUser, err := kafka.NewSaslUser(ctx, "fooSaslUser", &kafka.SaslUserArgs{
//				UserName:     pulumi.String("acc-test-user"),
//				InstanceId:   fooInstance.ID(),
//				UserPassword: pulumi.String("suqsnis123!"),
//				Description:  pulumi.String("tf-test"),
//				AllAuthority: pulumi.Bool(true),
//				PasswordType: pulumi.String("Scram"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = kafka.NewTopic(ctx, "fooTopic", &kafka.TopicArgs{
//				TopicName:       pulumi.String("acc-test-topic"),
//				InstanceId:      fooInstance.ID(),
//				Description:     pulumi.String("tf-test"),
//				PartitionNumber: pulumi.Int(15),
//				ReplicaNumber:   pulumi.Int(3),
//				Parameters: &kafka.TopicParametersArgs{
//					MinInsyncReplicaNumber: pulumi.Int(2),
//					MessageMaxByte:         pulumi.Int(10),
//					LogRetentionHours:      pulumi.Int(96),
//				},
//				AllAuthority: pulumi.Bool(false),
//				AccessPolicies: kafka.TopicAccessPolicyArray{
//					&kafka.TopicAccessPolicyArgs{
//						UserName:     fooSaslUser.UserName,
//						AccessPolicy: pulumi.String("Pub"),
//					},
//				},
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
// KafkaTopic can be imported using the instance_id:topic_name, e.g.
//
// ```sh
// $ pulumi import volcengine:kafka/topic:Topic default kafka-cnoeeapetf4s****:topic
// ```
type Topic struct {
	pulumi.CustomResourceState

	// The access policies info of the kafka topic. This field only valid when the value of the AllAuthority is false.
	AccessPolicies TopicAccessPolicyArrayOutput `pulumi:"accessPolicies"`
	// Whether the kafka topic is configured to be accessible by all users. Default: true.
	AllAuthority pulumi.BoolPtrOutput `pulumi:"allAuthority"`
	// The description of the kafka topic.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The instance id of the kafka topic.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The parameters of the kafka topic.
	Parameters TopicParametersOutput `pulumi:"parameters"`
	// The number of partition in kafka topic. The value range in 1-300. This field can only be adjusted up but not down.
	PartitionNumber pulumi.IntOutput `pulumi:"partitionNumber"`
	// The number of replica in kafka topic. The value can be 2 or 3. Default is 3.
	ReplicaNumber pulumi.IntPtrOutput `pulumi:"replicaNumber"`
	// The name of the kafka topic.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.PartitionNumber == nil {
		return nil, errors.New("invalid value for required argument 'PartitionNumber'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Topic
	err := ctx.RegisterResource("volcengine:kafka/topic:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("volcengine:kafka/topic:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
	// The access policies info of the kafka topic. This field only valid when the value of the AllAuthority is false.
	AccessPolicies []TopicAccessPolicy `pulumi:"accessPolicies"`
	// Whether the kafka topic is configured to be accessible by all users. Default: true.
	AllAuthority *bool `pulumi:"allAuthority"`
	// The description of the kafka topic.
	Description *string `pulumi:"description"`
	// The instance id of the kafka topic.
	InstanceId *string `pulumi:"instanceId"`
	// The parameters of the kafka topic.
	Parameters *TopicParameters `pulumi:"parameters"`
	// The number of partition in kafka topic. The value range in 1-300. This field can only be adjusted up but not down.
	PartitionNumber *int `pulumi:"partitionNumber"`
	// The number of replica in kafka topic. The value can be 2 or 3. Default is 3.
	ReplicaNumber *int `pulumi:"replicaNumber"`
	// The name of the kafka topic.
	TopicName *string `pulumi:"topicName"`
}

type TopicState struct {
	// The access policies info of the kafka topic. This field only valid when the value of the AllAuthority is false.
	AccessPolicies TopicAccessPolicyArrayInput
	// Whether the kafka topic is configured to be accessible by all users. Default: true.
	AllAuthority pulumi.BoolPtrInput
	// The description of the kafka topic.
	Description pulumi.StringPtrInput
	// The instance id of the kafka topic.
	InstanceId pulumi.StringPtrInput
	// The parameters of the kafka topic.
	Parameters TopicParametersPtrInput
	// The number of partition in kafka topic. The value range in 1-300. This field can only be adjusted up but not down.
	PartitionNumber pulumi.IntPtrInput
	// The number of replica in kafka topic. The value can be 2 or 3. Default is 3.
	ReplicaNumber pulumi.IntPtrInput
	// The name of the kafka topic.
	TopicName pulumi.StringPtrInput
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// The access policies info of the kafka topic. This field only valid when the value of the AllAuthority is false.
	AccessPolicies []TopicAccessPolicy `pulumi:"accessPolicies"`
	// Whether the kafka topic is configured to be accessible by all users. Default: true.
	AllAuthority *bool `pulumi:"allAuthority"`
	// The description of the kafka topic.
	Description *string `pulumi:"description"`
	// The instance id of the kafka topic.
	InstanceId string `pulumi:"instanceId"`
	// The parameters of the kafka topic.
	Parameters *TopicParameters `pulumi:"parameters"`
	// The number of partition in kafka topic. The value range in 1-300. This field can only be adjusted up but not down.
	PartitionNumber int `pulumi:"partitionNumber"`
	// The number of replica in kafka topic. The value can be 2 or 3. Default is 3.
	ReplicaNumber *int `pulumi:"replicaNumber"`
	// The name of the kafka topic.
	TopicName string `pulumi:"topicName"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// The access policies info of the kafka topic. This field only valid when the value of the AllAuthority is false.
	AccessPolicies TopicAccessPolicyArrayInput
	// Whether the kafka topic is configured to be accessible by all users. Default: true.
	AllAuthority pulumi.BoolPtrInput
	// The description of the kafka topic.
	Description pulumi.StringPtrInput
	// The instance id of the kafka topic.
	InstanceId pulumi.StringInput
	// The parameters of the kafka topic.
	Parameters TopicParametersPtrInput
	// The number of partition in kafka topic. The value range in 1-300. This field can only be adjusted up but not down.
	PartitionNumber pulumi.IntInput
	// The number of replica in kafka topic. The value can be 2 or 3. Default is 3.
	ReplicaNumber pulumi.IntPtrInput
	// The name of the kafka topic.
	TopicName pulumi.StringInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}

type TopicInput interface {
	pulumi.Input

	ToTopicOutput() TopicOutput
	ToTopicOutputWithContext(ctx context.Context) TopicOutput
}

func (*Topic) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (i *Topic) ToTopicOutput() TopicOutput {
	return i.ToTopicOutputWithContext(context.Background())
}

func (i *Topic) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicOutput)
}

// TopicArrayInput is an input type that accepts TopicArray and TopicArrayOutput values.
// You can construct a concrete instance of `TopicArrayInput` via:
//
//	TopicArray{ TopicArgs{...} }
type TopicArrayInput interface {
	pulumi.Input

	ToTopicArrayOutput() TopicArrayOutput
	ToTopicArrayOutputWithContext(context.Context) TopicArrayOutput
}

type TopicArray []TopicInput

func (TopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Topic)(nil)).Elem()
}

func (i TopicArray) ToTopicArrayOutput() TopicArrayOutput {
	return i.ToTopicArrayOutputWithContext(context.Background())
}

func (i TopicArray) ToTopicArrayOutputWithContext(ctx context.Context) TopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicArrayOutput)
}

// TopicMapInput is an input type that accepts TopicMap and TopicMapOutput values.
// You can construct a concrete instance of `TopicMapInput` via:
//
//	TopicMap{ "key": TopicArgs{...} }
type TopicMapInput interface {
	pulumi.Input

	ToTopicMapOutput() TopicMapOutput
	ToTopicMapOutputWithContext(context.Context) TopicMapOutput
}

type TopicMap map[string]TopicInput

func (TopicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Topic)(nil)).Elem()
}

func (i TopicMap) ToTopicMapOutput() TopicMapOutput {
	return i.ToTopicMapOutputWithContext(context.Background())
}

func (i TopicMap) ToTopicMapOutputWithContext(ctx context.Context) TopicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicMapOutput)
}

type TopicOutput struct{ *pulumi.OutputState }

func (TopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (o TopicOutput) ToTopicOutput() TopicOutput {
	return o
}

func (o TopicOutput) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return o
}

// The access policies info of the kafka topic. This field only valid when the value of the AllAuthority is false.
func (o TopicOutput) AccessPolicies() TopicAccessPolicyArrayOutput {
	return o.ApplyT(func(v *Topic) TopicAccessPolicyArrayOutput { return v.AccessPolicies }).(TopicAccessPolicyArrayOutput)
}

// Whether the kafka topic is configured to be accessible by all users. Default: true.
func (o TopicOutput) AllAuthority() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.BoolPtrOutput { return v.AllAuthority }).(pulumi.BoolPtrOutput)
}

// The description of the kafka topic.
func (o TopicOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The instance id of the kafka topic.
func (o TopicOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The parameters of the kafka topic.
func (o TopicOutput) Parameters() TopicParametersOutput {
	return o.ApplyT(func(v *Topic) TopicParametersOutput { return v.Parameters }).(TopicParametersOutput)
}

// The number of partition in kafka topic. The value range in 1-300. This field can only be adjusted up but not down.
func (o TopicOutput) PartitionNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntOutput { return v.PartitionNumber }).(pulumi.IntOutput)
}

// The number of replica in kafka topic. The value can be 2 or 3. Default is 3.
func (o TopicOutput) ReplicaNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntPtrOutput { return v.ReplicaNumber }).(pulumi.IntPtrOutput)
}

// The name of the kafka topic.
func (o TopicOutput) TopicName() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.TopicName }).(pulumi.StringOutput)
}

type TopicArrayOutput struct{ *pulumi.OutputState }

func (TopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Topic)(nil)).Elem()
}

func (o TopicArrayOutput) ToTopicArrayOutput() TopicArrayOutput {
	return o
}

func (o TopicArrayOutput) ToTopicArrayOutputWithContext(ctx context.Context) TopicArrayOutput {
	return o
}

func (o TopicArrayOutput) Index(i pulumi.IntInput) TopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Topic {
		return vs[0].([]*Topic)[vs[1].(int)]
	}).(TopicOutput)
}

type TopicMapOutput struct{ *pulumi.OutputState }

func (TopicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Topic)(nil)).Elem()
}

func (o TopicMapOutput) ToTopicMapOutput() TopicMapOutput {
	return o
}

func (o TopicMapOutput) ToTopicMapOutputWithContext(ctx context.Context) TopicMapOutput {
	return o
}

func (o TopicMapOutput) MapIndex(k pulumi.StringInput) TopicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Topic {
		return vs[0].(map[string]*Topic)[vs[1].(string)]
	}).(TopicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TopicInput)(nil)).Elem(), &Topic{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicArrayInput)(nil)).Elem(), TopicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicMapInput)(nil)).Elem(), TopicMap{})
	pulumi.RegisterOutputType(TopicOutput{})
	pulumi.RegisterOutputType(TopicArrayOutput{})
	pulumi.RegisterOutputType(TopicMapOutput{})
}
