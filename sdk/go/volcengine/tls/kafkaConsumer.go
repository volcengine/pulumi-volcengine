// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage tls kafka consumer
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/tls"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tls.NewKafkaConsumer(ctx, "foo", &tls.KafkaConsumerArgs{
//				TopicId: pulumi.String("cfb5c08b-0c7a-44fa-8971-8afc12f1b123"),
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
// Tls Kafka Consumer can be imported using the kafka:topic_id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:tls/kafkaConsumer:KafkaConsumer default kafka:edf051ed-3c46-49ba-9339-bea628fedc15
//
// ```
type KafkaConsumer struct {
	pulumi.CustomResourceState

	// Whether allow consume.
	AllowConsume pulumi.BoolOutput `pulumi:"allowConsume"`
	// The topic of consume.
	ConsumeTopic pulumi.StringOutput `pulumi:"consumeTopic"`
	// The id of topic.
	TopicId pulumi.StringOutput `pulumi:"topicId"`
}

// NewKafkaConsumer registers a new resource with the given unique name, arguments, and options.
func NewKafkaConsumer(ctx *pulumi.Context,
	name string, args *KafkaConsumerArgs, opts ...pulumi.ResourceOption) (*KafkaConsumer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TopicId == nil {
		return nil, errors.New("invalid value for required argument 'TopicId'")
	}
	var resource KafkaConsumer
	err := ctx.RegisterResource("volcengine:tls/kafkaConsumer:KafkaConsumer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKafkaConsumer gets an existing KafkaConsumer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKafkaConsumer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KafkaConsumerState, opts ...pulumi.ResourceOption) (*KafkaConsumer, error) {
	var resource KafkaConsumer
	err := ctx.ReadResource("volcengine:tls/kafkaConsumer:KafkaConsumer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KafkaConsumer resources.
type kafkaConsumerState struct {
	// Whether allow consume.
	AllowConsume *bool `pulumi:"allowConsume"`
	// The topic of consume.
	ConsumeTopic *string `pulumi:"consumeTopic"`
	// The id of topic.
	TopicId *string `pulumi:"topicId"`
}

type KafkaConsumerState struct {
	// Whether allow consume.
	AllowConsume pulumi.BoolPtrInput
	// The topic of consume.
	ConsumeTopic pulumi.StringPtrInput
	// The id of topic.
	TopicId pulumi.StringPtrInput
}

func (KafkaConsumerState) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaConsumerState)(nil)).Elem()
}

type kafkaConsumerArgs struct {
	// The id of topic.
	TopicId string `pulumi:"topicId"`
}

// The set of arguments for constructing a KafkaConsumer resource.
type KafkaConsumerArgs struct {
	// The id of topic.
	TopicId pulumi.StringInput
}

func (KafkaConsumerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaConsumerArgs)(nil)).Elem()
}

type KafkaConsumerInput interface {
	pulumi.Input

	ToKafkaConsumerOutput() KafkaConsumerOutput
	ToKafkaConsumerOutputWithContext(ctx context.Context) KafkaConsumerOutput
}

func (*KafkaConsumer) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaConsumer)(nil)).Elem()
}

func (i *KafkaConsumer) ToKafkaConsumerOutput() KafkaConsumerOutput {
	return i.ToKafkaConsumerOutputWithContext(context.Background())
}

func (i *KafkaConsumer) ToKafkaConsumerOutputWithContext(ctx context.Context) KafkaConsumerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaConsumerOutput)
}

// KafkaConsumerArrayInput is an input type that accepts KafkaConsumerArray and KafkaConsumerArrayOutput values.
// You can construct a concrete instance of `KafkaConsumerArrayInput` via:
//
//	KafkaConsumerArray{ KafkaConsumerArgs{...} }
type KafkaConsumerArrayInput interface {
	pulumi.Input

	ToKafkaConsumerArrayOutput() KafkaConsumerArrayOutput
	ToKafkaConsumerArrayOutputWithContext(context.Context) KafkaConsumerArrayOutput
}

type KafkaConsumerArray []KafkaConsumerInput

func (KafkaConsumerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaConsumer)(nil)).Elem()
}

func (i KafkaConsumerArray) ToKafkaConsumerArrayOutput() KafkaConsumerArrayOutput {
	return i.ToKafkaConsumerArrayOutputWithContext(context.Background())
}

func (i KafkaConsumerArray) ToKafkaConsumerArrayOutputWithContext(ctx context.Context) KafkaConsumerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaConsumerArrayOutput)
}

// KafkaConsumerMapInput is an input type that accepts KafkaConsumerMap and KafkaConsumerMapOutput values.
// You can construct a concrete instance of `KafkaConsumerMapInput` via:
//
//	KafkaConsumerMap{ "key": KafkaConsumerArgs{...} }
type KafkaConsumerMapInput interface {
	pulumi.Input

	ToKafkaConsumerMapOutput() KafkaConsumerMapOutput
	ToKafkaConsumerMapOutputWithContext(context.Context) KafkaConsumerMapOutput
}

type KafkaConsumerMap map[string]KafkaConsumerInput

func (KafkaConsumerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaConsumer)(nil)).Elem()
}

func (i KafkaConsumerMap) ToKafkaConsumerMapOutput() KafkaConsumerMapOutput {
	return i.ToKafkaConsumerMapOutputWithContext(context.Background())
}

func (i KafkaConsumerMap) ToKafkaConsumerMapOutputWithContext(ctx context.Context) KafkaConsumerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaConsumerMapOutput)
}

type KafkaConsumerOutput struct{ *pulumi.OutputState }

func (KafkaConsumerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaConsumer)(nil)).Elem()
}

func (o KafkaConsumerOutput) ToKafkaConsumerOutput() KafkaConsumerOutput {
	return o
}

func (o KafkaConsumerOutput) ToKafkaConsumerOutputWithContext(ctx context.Context) KafkaConsumerOutput {
	return o
}

// Whether allow consume.
func (o KafkaConsumerOutput) AllowConsume() pulumi.BoolOutput {
	return o.ApplyT(func(v *KafkaConsumer) pulumi.BoolOutput { return v.AllowConsume }).(pulumi.BoolOutput)
}

// The topic of consume.
func (o KafkaConsumerOutput) ConsumeTopic() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaConsumer) pulumi.StringOutput { return v.ConsumeTopic }).(pulumi.StringOutput)
}

// The id of topic.
func (o KafkaConsumerOutput) TopicId() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaConsumer) pulumi.StringOutput { return v.TopicId }).(pulumi.StringOutput)
}

type KafkaConsumerArrayOutput struct{ *pulumi.OutputState }

func (KafkaConsumerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaConsumer)(nil)).Elem()
}

func (o KafkaConsumerArrayOutput) ToKafkaConsumerArrayOutput() KafkaConsumerArrayOutput {
	return o
}

func (o KafkaConsumerArrayOutput) ToKafkaConsumerArrayOutputWithContext(ctx context.Context) KafkaConsumerArrayOutput {
	return o
}

func (o KafkaConsumerArrayOutput) Index(i pulumi.IntInput) KafkaConsumerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KafkaConsumer {
		return vs[0].([]*KafkaConsumer)[vs[1].(int)]
	}).(KafkaConsumerOutput)
}

type KafkaConsumerMapOutput struct{ *pulumi.OutputState }

func (KafkaConsumerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaConsumer)(nil)).Elem()
}

func (o KafkaConsumerMapOutput) ToKafkaConsumerMapOutput() KafkaConsumerMapOutput {
	return o
}

func (o KafkaConsumerMapOutput) ToKafkaConsumerMapOutputWithContext(ctx context.Context) KafkaConsumerMapOutput {
	return o
}

func (o KafkaConsumerMapOutput) MapIndex(k pulumi.StringInput) KafkaConsumerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KafkaConsumer {
		return vs[0].(map[string]*KafkaConsumer)[vs[1].(string)]
	}).(KafkaConsumerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaConsumerInput)(nil)).Elem(), &KafkaConsumer{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaConsumerArrayInput)(nil)).Elem(), KafkaConsumerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaConsumerMapInput)(nil)).Elem(), KafkaConsumerMap{})
	pulumi.RegisterOutputType(KafkaConsumerOutput{})
	pulumi.RegisterOutputType(KafkaConsumerArrayOutput{})
	pulumi.RegisterOutputType(KafkaConsumerMapOutput{})
}