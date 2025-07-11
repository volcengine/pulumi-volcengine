// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of redis allow lists
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
//			fooAllowList, err := redis.NewAllowList(ctx, "fooAllowList", &redis.AllowListArgs{
//				AllowLists: pulumi.StringArray{
//					pulumi.String("192.168.0.0/24"),
//				},
//				AllowListName: pulumi.String("acc-test-allowlist"),
//			})
//			if err != nil {
//				return err
//			}
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
//			fooAllowListAssociate, err := redis.NewAllowListAssociate(ctx, "fooAllowListAssociate", &redis.AllowListAssociateArgs{
//				AllowListId: fooAllowList.ID(),
//				InstanceId:  fooInstance.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_ = redis.GetAllowListsOutput(ctx, redis.GetAllowListsOutputArgs{
//				InstanceId: fooAllowListAssociate.InstanceId,
//				RegionId:   pulumi.String("cn-beijing"),
//				NameRegex:  fooAllowList.AllowListName,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func GetAllowLists(ctx *pulumi.Context, args *GetAllowListsArgs, opts ...pulumi.InvokeOption) (*GetAllowListsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAllowListsResult
	err := ctx.Invoke("volcengine:redis/getAllowLists:getAllowLists", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAllowLists.
type GetAllowListsArgs struct {
	// The Id of instance.
	InstanceId *string `pulumi:"instanceId"`
	// Filter out the whitelist that meets the conditions based on the IP address. When using IPAddress query, it will precisely match this IP address and filter the IP address segments containing this IP address.
	IpAddress *string `pulumi:"ipAddress"`
	// Screen out the whitelist that meets the conditions based on the IP address segment. When using IPSegment queries, the IP address segment will be precisely matched for filtering.
	IpSegment *string `pulumi:"ipSegment"`
	// A Name Regex of Allow List.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The name of the project to which the white list belongs.
	ProjectName *string `pulumi:"projectName"`
	// Filter whether to query only the default whitelist based on the type of whitelist.
	QueryDefault *bool `pulumi:"queryDefault"`
	// The Id of region.
	RegionId string `pulumi:"regionId"`
}

// A collection of values returned by getAllowLists.
type GetAllowListsResult struct {
	// Information of list of allow list.
	AllowLists []GetAllowListsAllowList `pulumi:"allowLists"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Id of instance.
	InstanceId *string `pulumi:"instanceId"`
	IpAddress  *string `pulumi:"ipAddress"`
	IpSegment  *string `pulumi:"ipSegment"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The name of the project to which the white list belongs.
	ProjectName  *string `pulumi:"projectName"`
	QueryDefault *bool   `pulumi:"queryDefault"`
	RegionId     string  `pulumi:"regionId"`
	// The total count of allow list query.
	TotalCount int `pulumi:"totalCount"`
}

func GetAllowListsOutput(ctx *pulumi.Context, args GetAllowListsOutputArgs, opts ...pulumi.InvokeOption) GetAllowListsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAllowListsResult, error) {
			args := v.(GetAllowListsArgs)
			r, err := GetAllowLists(ctx, &args, opts...)
			var s GetAllowListsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAllowListsResultOutput)
}

// A collection of arguments for invoking getAllowLists.
type GetAllowListsOutputArgs struct {
	// The Id of instance.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// Filter out the whitelist that meets the conditions based on the IP address. When using IPAddress query, it will precisely match this IP address and filter the IP address segments containing this IP address.
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
	// Screen out the whitelist that meets the conditions based on the IP address segment. When using IPSegment queries, the IP address segment will be precisely matched for filtering.
	IpSegment pulumi.StringPtrInput `pulumi:"ipSegment"`
	// A Name Regex of Allow List.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The name of the project to which the white list belongs.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
	// Filter whether to query only the default whitelist based on the type of whitelist.
	QueryDefault pulumi.BoolPtrInput `pulumi:"queryDefault"`
	// The Id of region.
	RegionId pulumi.StringInput `pulumi:"regionId"`
}

func (GetAllowListsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAllowListsArgs)(nil)).Elem()
}

// A collection of values returned by getAllowLists.
type GetAllowListsResultOutput struct{ *pulumi.OutputState }

func (GetAllowListsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAllowListsResult)(nil)).Elem()
}

func (o GetAllowListsResultOutput) ToGetAllowListsResultOutput() GetAllowListsResultOutput {
	return o
}

func (o GetAllowListsResultOutput) ToGetAllowListsResultOutputWithContext(ctx context.Context) GetAllowListsResultOutput {
	return o
}

// Information of list of allow list.
func (o GetAllowListsResultOutput) AllowLists() GetAllowListsAllowListArrayOutput {
	return o.ApplyT(func(v GetAllowListsResult) []GetAllowListsAllowList { return v.AllowLists }).(GetAllowListsAllowListArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAllowListsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAllowListsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Id of instance.
func (o GetAllowListsResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAllowListsResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o GetAllowListsResultOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAllowListsResult) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o GetAllowListsResultOutput) IpSegment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAllowListsResult) *string { return v.IpSegment }).(pulumi.StringPtrOutput)
}

func (o GetAllowListsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAllowListsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetAllowListsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAllowListsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The name of the project to which the white list belongs.
func (o GetAllowListsResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAllowListsResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

func (o GetAllowListsResultOutput) QueryDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAllowListsResult) *bool { return v.QueryDefault }).(pulumi.BoolPtrOutput)
}

func (o GetAllowListsResultOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAllowListsResult) string { return v.RegionId }).(pulumi.StringOutput)
}

// The total count of allow list query.
func (o GetAllowListsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetAllowListsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAllowListsResultOutput{})
}
