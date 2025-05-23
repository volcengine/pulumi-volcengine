// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kafka

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of kafka groups
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
//			_ = kafka.GetGroupsOutput(ctx, kafka.GetGroupsOutputArgs{
//				InstanceId: fooGroup.InstanceId,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func GetGroups(ctx *pulumi.Context, args *GetGroupsArgs, opts ...pulumi.InvokeOption) (*GetGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGroupsResult
	err := ctx.Invoke("volcengine:kafka/getGroups:getGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroups.
type GetGroupsArgs struct {
	// The id of kafka group, support fuzzy matching.
	GroupId *string `pulumi:"groupId"`
	// The instance id of kafka group.
	InstanceId string `pulumi:"instanceId"`
	// A Name Regex of kafka group.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getGroups.
type GetGroupsResult struct {
	// The id of kafka group.
	GroupId *string `pulumi:"groupId"`
	// The collection of query.
	Groups []GetGroupsGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstanceId string  `pulumi:"instanceId"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func GetGroupsOutput(ctx *pulumi.Context, args GetGroupsOutputArgs, opts ...pulumi.InvokeOption) GetGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGroupsResult, error) {
			args := v.(GetGroupsArgs)
			r, err := GetGroups(ctx, &args, opts...)
			var s GetGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGroupsResultOutput)
}

// A collection of arguments for invoking getGroups.
type GetGroupsOutputArgs struct {
	// The id of kafka group, support fuzzy matching.
	GroupId pulumi.StringPtrInput `pulumi:"groupId"`
	// The instance id of kafka group.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// A Name Regex of kafka group.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getGroups.
type GetGroupsResultOutput struct{ *pulumi.OutputState }

func (GetGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupsResult)(nil)).Elem()
}

func (o GetGroupsResultOutput) ToGetGroupsResultOutput() GetGroupsResultOutput {
	return o
}

func (o GetGroupsResultOutput) ToGetGroupsResultOutputWithContext(ctx context.Context) GetGroupsResultOutput {
	return o
}

// The id of kafka group.
func (o GetGroupsResultOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

// The collection of query.
func (o GetGroupsResultOutput) Groups() GetGroupsGroupArrayOutput {
	return o.ApplyT(func(v GetGroupsResult) []GetGroupsGroup { return v.Groups }).(GetGroupsGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetGroupsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetGroupsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o GetGroupsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetGroupsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGroupsResultOutput{})
}
