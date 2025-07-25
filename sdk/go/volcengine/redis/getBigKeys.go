// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of redis big keys
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/redis"
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
//			fooInstance, err := redis.NewInstance(ctx, "fooInstance", &redis.InstanceArgs{
//				ZoneIds: pulumi.StringArray{
//					pulumi.String(fooZones.Zones[0].Id),
//				},
//				InstanceName:       pulumi.String("acc-test-tf-redis"),
//				ShardedCluster:     pulumi.Int(1),
//				Password:           pulumi.String("1qaz!QAZ12"),
//				NodeNumber:         pulumi.Int(2),
//				ShardCapacity:      pulumi.Int(1024),
//				ShardNumber:        pulumi.Int(2),
//				EngineVersion:      pulumi.String("5.0"),
//				SubnetId:           fooSubnet.ID(),
//				DeletionProtection: pulumi.String("disabled"),
//				VpcAuthMode:        pulumi.String("close"),
//				ChargeType:         pulumi.String("PostPaid"),
//				Port:               pulumi.Int(6381),
//				ProjectName:        pulumi.String("default"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = redis.GetBigKeysOutput(ctx, redis.GetBigKeysOutputArgs{
//				InstanceId: fooInstance.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func GetBigKeys(ctx *pulumi.Context, args *GetBigKeysArgs, opts ...pulumi.InvokeOption) (*GetBigKeysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBigKeysResult
	err := ctx.Invoke("volcengine:redis/getBigKeys:getBigKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBigKeys.
type GetBigKeysArgs struct {
	// The ID of Instance.
	InstanceId string `pulumi:"instanceId"`
	// Specify the data type used to filter the query results of hot keys.
	KeyType *string `pulumi:"keyType"`
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// Specify the sorting conditions of the query results.
	OrderBy *string `pulumi:"orderBy"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// Query the end time in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).
	QueryEndTime *string `pulumi:"queryEndTime"`
	// Query the start time in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).
	QueryStartTime *string `pulumi:"queryStartTime"`
}

// A collection of values returned by getBigKeys.
type GetBigKeysResult struct {
	// Details of the big Key.
	BigKeys []GetBigKeysBigKey `pulumi:"bigKeys"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// The type of big Key.
	KeyType        *string `pulumi:"keyType"`
	NameRegex      *string `pulumi:"nameRegex"`
	OrderBy        *string `pulumi:"orderBy"`
	OutputFile     *string `pulumi:"outputFile"`
	QueryEndTime   *string `pulumi:"queryEndTime"`
	QueryStartTime *string `pulumi:"queryStartTime"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func GetBigKeysOutput(ctx *pulumi.Context, args GetBigKeysOutputArgs, opts ...pulumi.InvokeOption) GetBigKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBigKeysResult, error) {
			args := v.(GetBigKeysArgs)
			r, err := GetBigKeys(ctx, &args, opts...)
			var s GetBigKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBigKeysResultOutput)
}

// A collection of arguments for invoking getBigKeys.
type GetBigKeysOutputArgs struct {
	// The ID of Instance.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Specify the data type used to filter the query results of hot keys.
	KeyType pulumi.StringPtrInput `pulumi:"keyType"`
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// Specify the sorting conditions of the query results.
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Query the end time in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).
	QueryEndTime pulumi.StringPtrInput `pulumi:"queryEndTime"`
	// Query the start time in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).
	QueryStartTime pulumi.StringPtrInput `pulumi:"queryStartTime"`
}

func (GetBigKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBigKeysArgs)(nil)).Elem()
}

// A collection of values returned by getBigKeys.
type GetBigKeysResultOutput struct{ *pulumi.OutputState }

func (GetBigKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBigKeysResult)(nil)).Elem()
}

func (o GetBigKeysResultOutput) ToGetBigKeysResultOutput() GetBigKeysResultOutput {
	return o
}

func (o GetBigKeysResultOutput) ToGetBigKeysResultOutputWithContext(ctx context.Context) GetBigKeysResultOutput {
	return o
}

// Details of the big Key.
func (o GetBigKeysResultOutput) BigKeys() GetBigKeysBigKeyArrayOutput {
	return o.ApplyT(func(v GetBigKeysResult) []GetBigKeysBigKey { return v.BigKeys }).(GetBigKeysBigKeyArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBigKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBigKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBigKeysResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetBigKeysResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// The type of big Key.
func (o GetBigKeysResultOutput) KeyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBigKeysResult) *string { return v.KeyType }).(pulumi.StringPtrOutput)
}

func (o GetBigKeysResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBigKeysResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetBigKeysResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBigKeysResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func (o GetBigKeysResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBigKeysResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetBigKeysResultOutput) QueryEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBigKeysResult) *string { return v.QueryEndTime }).(pulumi.StringPtrOutput)
}

func (o GetBigKeysResultOutput) QueryStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBigKeysResult) *string { return v.QueryStartTime }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o GetBigKeysResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetBigKeysResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBigKeysResultOutput{})
}
