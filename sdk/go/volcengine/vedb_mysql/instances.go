// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vedb_mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of vedb mysql instances
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vedb_mysql"
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
//				ZoneId:     pulumi.String(fooZones.Zones[2].Id),
//				VpcId:      fooVpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			fooInstance, err := vedb_mysql.NewInstance(ctx, "fooInstance", &vedb_mysql.InstanceArgs{
//				ChargeType:        pulumi.String("PostPaid"),
//				StorageChargeType: pulumi.String("PostPaid"),
//				DbEngineVersion:   pulumi.String("MySQL_8_0"),
//				DbMinorVersion:    pulumi.String("3.0"),
//				NodeNumber:        pulumi.Int(2),
//				NodeSpec:          pulumi.String("vedb.mysql.x4.large"),
//				SubnetId:          fooSubnet.ID(),
//				InstanceName:      pulumi.String("tf-test"),
//				ProjectName:       pulumi.String("testA"),
//				Tags: vedb_mysql.InstanceTagArray{
//					&vedb_mysql.InstanceTagArgs{
//						Key:   pulumi.String("tftest"),
//						Value: pulumi.String("tftest"),
//					},
//					&vedb_mysql.InstanceTagArgs{
//						Key:   pulumi.String("tftest2"),
//						Value: pulumi.String("tftest2"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = vedb_mysql.InstancesOutput(ctx, vedb_mysql.InstancesOutputArgs{
//				InstanceId: fooInstance.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func Instances(ctx *pulumi.Context, args *InstancesArgs, opts ...pulumi.InvokeOption) (*InstancesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv InstancesResult
	err := ctx.Invoke("volcengine:vedb_mysql/instances:Instances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Instances.
type InstancesArgs struct {
	// The charge type of the veDB Mysql instance.
	ChargeType *string `pulumi:"chargeType"`
	// The end time of creating veDB Mysql instance.
	CreateTimeEnd *string `pulumi:"createTimeEnd"`
	// The start time of creating veDB Mysql instance.
	CreateTimeStart *string `pulumi:"createTimeStart"`
	// The version of the veDB Mysql instance.
	DbEngineVersion *string `pulumi:"dbEngineVersion"`
	// The id of the veDB Mysql instance.
	InstanceId *string `pulumi:"instanceId"`
	// The name of the veDB Mysql instance.
	InstanceName *string `pulumi:"instanceName"`
	// The status of the veDB Mysql instance.
	InstanceStatus *string `pulumi:"instanceStatus"`
	// A Name Regex of veDB mysql instance.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The project name of the veDB Mysql instance.
	ProjectName *string `pulumi:"projectName"`
	// Tags.
	Tags []InstancesTag `pulumi:"tags"`
	// The available zone of the veDB Mysql instance.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by Instances.
type InstancesResult struct {
	// Calculate the billing type. Values:
	// PostPaid: Pay-as-you-go (postpaid).
	// PrePaid: Monthly/yearly subscription (prepaid).
	ChargeType      *string `pulumi:"chargeType"`
	CreateTimeEnd   *string `pulumi:"createTimeEnd"`
	CreateTimeStart *string `pulumi:"createTimeStart"`
	// The engine version of the veDB Mysql instance.
	DbEngineVersion *string `pulumi:"dbEngineVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ID of the veDB Mysql instance.
	InstanceId *string `pulumi:"instanceId"`
	// The name of the veDB Mysql instance.
	InstanceName *string `pulumi:"instanceName"`
	// The status of the veDB Mysql instance.
	InstanceStatus *string `pulumi:"instanceStatus"`
	// The collection of query.
	Instances  []InstancesInstance `pulumi:"instances"`
	NameRegex  *string             `pulumi:"nameRegex"`
	OutputFile *string             `pulumi:"outputFile"`
	// The project name of the veDB Mysql instance.
	ProjectName *string `pulumi:"projectName"`
	// Tags.
	Tags []InstancesTag `pulumi:"tags"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
	// The zone id.
	ZoneId *string `pulumi:"zoneId"`
}

func InstancesOutput(ctx *pulumi.Context, args InstancesOutputArgs, opts ...pulumi.InvokeOption) InstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (InstancesResult, error) {
			args := v.(InstancesArgs)
			r, err := Instances(ctx, &args, opts...)
			var s InstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(InstancesResultOutput)
}

// A collection of arguments for invoking Instances.
type InstancesOutputArgs struct {
	// The charge type of the veDB Mysql instance.
	ChargeType pulumi.StringPtrInput `pulumi:"chargeType"`
	// The end time of creating veDB Mysql instance.
	CreateTimeEnd pulumi.StringPtrInput `pulumi:"createTimeEnd"`
	// The start time of creating veDB Mysql instance.
	CreateTimeStart pulumi.StringPtrInput `pulumi:"createTimeStart"`
	// The version of the veDB Mysql instance.
	DbEngineVersion pulumi.StringPtrInput `pulumi:"dbEngineVersion"`
	// The id of the veDB Mysql instance.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// The name of the veDB Mysql instance.
	InstanceName pulumi.StringPtrInput `pulumi:"instanceName"`
	// The status of the veDB Mysql instance.
	InstanceStatus pulumi.StringPtrInput `pulumi:"instanceStatus"`
	// A Name Regex of veDB mysql instance.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The project name of the veDB Mysql instance.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
	// Tags.
	Tags InstancesTagArrayInput `pulumi:"tags"`
	// The available zone of the veDB Mysql instance.
	ZoneId pulumi.StringPtrInput `pulumi:"zoneId"`
}

func (InstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesArgs)(nil)).Elem()
}

// A collection of values returned by Instances.
type InstancesResultOutput struct{ *pulumi.OutputState }

func (InstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesResult)(nil)).Elem()
}

func (o InstancesResultOutput) ToInstancesResultOutput() InstancesResultOutput {
	return o
}

func (o InstancesResultOutput) ToInstancesResultOutputWithContext(ctx context.Context) InstancesResultOutput {
	return o
}

// Calculate the billing type. Values:
// PostPaid: Pay-as-you-go (postpaid).
// PrePaid: Monthly/yearly subscription (prepaid).
func (o InstancesResultOutput) ChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.ChargeType }).(pulumi.StringPtrOutput)
}

func (o InstancesResultOutput) CreateTimeEnd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.CreateTimeEnd }).(pulumi.StringPtrOutput)
}

func (o InstancesResultOutput) CreateTimeStart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.CreateTimeStart }).(pulumi.StringPtrOutput)
}

// The engine version of the veDB Mysql instance.
func (o InstancesResultOutput) DbEngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.DbEngineVersion }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o InstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of the veDB Mysql instance.
func (o InstancesResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

// The name of the veDB Mysql instance.
func (o InstancesResultOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.InstanceName }).(pulumi.StringPtrOutput)
}

// The status of the veDB Mysql instance.
func (o InstancesResultOutput) InstanceStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.InstanceStatus }).(pulumi.StringPtrOutput)
}

// The collection of query.
func (o InstancesResultOutput) Instances() InstancesInstanceArrayOutput {
	return o.ApplyT(func(v InstancesResult) []InstancesInstance { return v.Instances }).(InstancesInstanceArrayOutput)
}

func (o InstancesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o InstancesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The project name of the veDB Mysql instance.
func (o InstancesResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// Tags.
func (o InstancesResultOutput) Tags() InstancesTagArrayOutput {
	return o.ApplyT(func(v InstancesResult) []InstancesTag { return v.Tags }).(InstancesTagArrayOutput)
}

// The total count of query.
func (o InstancesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The zone id.
func (o InstancesResultOutput) ZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.ZoneId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(InstancesResultOutput{})
}