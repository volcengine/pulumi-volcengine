// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kafka

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of kafka consumed topics
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
//			fooGroup, err := kafka.NewGroup(ctx, "fooGroup", &kafka.GroupArgs{
//				InstanceId:  fooInstance.ID(),
//				GroupId:     pulumi.String("acc-test-group"),
//				Description: pulumi.String("tf-test"),
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
//			fooTopic, err := kafka.NewTopic(ctx, "fooTopic", &kafka.TopicArgs{
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
//			_ = kafka.GetConsumedTopicsOutput(ctx, kafka.GetConsumedTopicsOutputArgs{
//				InstanceId: fooInstance.ID(),
//				GroupId:    fooGroup.GroupId,
//				TopicName:  fooTopic.TopicName,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: volcengine.kafka.ConsumedTopics has been deprecated in favor of volcengine.kafka.getConsumedTopics
func ConsumedTopics(ctx *pulumi.Context, args *ConsumedTopicsArgs, opts ...pulumi.InvokeOption) (*ConsumedTopicsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ConsumedTopicsResult
	err := ctx.Invoke("volcengine:kafka/consumedTopics:ConsumedTopics", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking ConsumedTopics.
type ConsumedTopicsArgs struct {
	// The id of kafka group.
	GroupId string `pulumi:"groupId"`
	// The id of kafka instance.
	InstanceId string `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The name of kafka topic. This field supports fuzzy query.
	TopicName *string `pulumi:"topicName"`
}

// A collection of values returned by ConsumedTopics.
type ConsumedTopicsResult struct {
	// The collection of query.
	ConsumedTopics []ConsumedTopicsConsumedTopic `pulumi:"consumedTopics"`
	GroupId        string                        `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstanceId string  `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
	// The name of kafka topic.
	TopicName *string `pulumi:"topicName"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func ConsumedTopicsOutput(ctx *pulumi.Context, args ConsumedTopicsOutputArgs, opts ...pulumi.InvokeOption) ConsumedTopicsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ConsumedTopicsResult, error) {
			args := v.(ConsumedTopicsArgs)
			r, err := ConsumedTopics(ctx, &args, opts...)
			var s ConsumedTopicsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ConsumedTopicsResultOutput)
}

// A collection of arguments for invoking ConsumedTopics.
type ConsumedTopicsOutputArgs struct {
	// The id of kafka group.
	GroupId pulumi.StringInput `pulumi:"groupId"`
	// The id of kafka instance.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The name of kafka topic. This field supports fuzzy query.
	TopicName pulumi.StringPtrInput `pulumi:"topicName"`
}

func (ConsumedTopicsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsumedTopicsArgs)(nil)).Elem()
}

// A collection of values returned by ConsumedTopics.
type ConsumedTopicsResultOutput struct{ *pulumi.OutputState }

func (ConsumedTopicsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsumedTopicsResult)(nil)).Elem()
}

func (o ConsumedTopicsResultOutput) ToConsumedTopicsResultOutput() ConsumedTopicsResultOutput {
	return o
}

func (o ConsumedTopicsResultOutput) ToConsumedTopicsResultOutputWithContext(ctx context.Context) ConsumedTopicsResultOutput {
	return o
}

// The collection of query.
func (o ConsumedTopicsResultOutput) ConsumedTopics() ConsumedTopicsConsumedTopicArrayOutput {
	return o.ApplyT(func(v ConsumedTopicsResult) []ConsumedTopicsConsumedTopic { return v.ConsumedTopics }).(ConsumedTopicsConsumedTopicArrayOutput)
}

func (o ConsumedTopicsResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v ConsumedTopicsResult) string { return v.GroupId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o ConsumedTopicsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ConsumedTopicsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ConsumedTopicsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v ConsumedTopicsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o ConsumedTopicsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsumedTopicsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The name of kafka topic.
func (o ConsumedTopicsResultOutput) TopicName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConsumedTopicsResult) *string { return v.TopicName }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o ConsumedTopicsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v ConsumedTopicsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(ConsumedTopicsResultOutput{})
}
