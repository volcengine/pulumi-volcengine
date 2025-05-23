// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds_mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of rds mysql allowlists
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/rds_mysql"
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
//			var fooAllowlist []*rds_mysql.Allowlist
//			for index := 0; index < 3; index++ {
//				key0 := index
//				val0 := index
//				__res, err := rds_mysql.NewAllowlist(ctx, fmt.Sprintf("fooAllowlist-%v", key0), &rds_mysql.AllowlistArgs{
//					AllowListName: pulumi.String(fmt.Sprintf("acc-test-allowlist-%v", val0)),
//					AllowListDesc: pulumi.String("acc-test"),
//					AllowListType: pulumi.String("IPv4"),
//					AllowLists: pulumi.StringArray{
//						pulumi.String("192.168.0.0/24"),
//						pulumi.String("192.168.1.0/24"),
//					},
//				})
//				if err != nil {
//					return err
//				}
//				fooAllowlist = append(fooAllowlist, __res)
//			}
//			var splat0 pulumi.StringArray
//			for _, val0 := range fooAllowlist {
//				splat0 = append(splat0, val0.ID())
//			}
//			fooInstance, err := rds_mysql.NewInstance(ctx, "fooInstance", &rds_mysql.InstanceArgs{
//				InstanceName:        pulumi.String("acc-test-rds-mysql"),
//				DbEngineVersion:     pulumi.String("MySQL_5_7"),
//				NodeSpec:            pulumi.String("rds.mysql.1c2g"),
//				PrimaryZoneId:       pulumi.String(fooZones.Zones[0].Id),
//				SecondaryZoneId:     pulumi.String(fooZones.Zones[0].Id),
//				StorageSpace:        pulumi.Int(80),
//				SubnetId:            fooSubnet.ID(),
//				LowerCaseTableNames: pulumi.String("1"),
//				ChargeInfo: &rds_mysql.InstanceChargeInfoArgs{
//					ChargeType: pulumi.String("PostPaid"),
//				},
//				Parameters: rds_mysql.InstanceParameterArray{
//					&rds_mysql.InstanceParameterArgs{
//						ParameterName:  pulumi.String("auto_increment_increment"),
//						ParameterValue: pulumi.String("2"),
//					},
//					&rds_mysql.InstanceParameterArgs{
//						ParameterName:  pulumi.String("auto_increment_offset"),
//						ParameterValue: pulumi.String("4"),
//					},
//				},
//				AllowListIds: splat0,
//			})
//			if err != nil {
//				return err
//			}
//			_ = rds_mysql.GetAllowlistsOutput(ctx, rds_mysql.GetAllowlistsOutputArgs{
//				InstanceId: fooInstance.ID(),
//				RegionId:   pulumi.String("cn-beijing"),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: volcengine.rds_mysql.Allowlists has been deprecated in favor of volcengine.rds_mysql.getAllowlists
func Allowlists(ctx *pulumi.Context, args *AllowlistsArgs, opts ...pulumi.InvokeOption) (*AllowlistsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv AllowlistsResult
	err := ctx.Invoke("volcengine:rds_mysql/allowlists:Allowlists", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Allowlists.
type AllowlistsArgs struct {
	// Instance ID. When an InstanceId is specified, the DescribeAllowLists interface will return the whitelist bound to the specified instance.
	InstanceId *string `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The region of the allow lists.
	RegionId string `pulumi:"regionId"`
}

// A collection of values returned by Allowlists.
type AllowlistsResult struct {
	// The list of allowed list.
	AllowLists []AllowlistsAllowList `pulumi:"allowLists"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The id of the instance.
	InstanceId *string `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
	RegionId   string  `pulumi:"regionId"`
	// The total count of Scaling Activity query.
	TotalCount int `pulumi:"totalCount"`
}

func AllowlistsOutput(ctx *pulumi.Context, args AllowlistsOutputArgs, opts ...pulumi.InvokeOption) AllowlistsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (AllowlistsResult, error) {
			args := v.(AllowlistsArgs)
			r, err := Allowlists(ctx, &args, opts...)
			var s AllowlistsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(AllowlistsResultOutput)
}

// A collection of arguments for invoking Allowlists.
type AllowlistsOutputArgs struct {
	// Instance ID. When an InstanceId is specified, the DescribeAllowLists interface will return the whitelist bound to the specified instance.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The region of the allow lists.
	RegionId pulumi.StringInput `pulumi:"regionId"`
}

func (AllowlistsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowlistsArgs)(nil)).Elem()
}

// A collection of values returned by Allowlists.
type AllowlistsResultOutput struct{ *pulumi.OutputState }

func (AllowlistsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowlistsResult)(nil)).Elem()
}

func (o AllowlistsResultOutput) ToAllowlistsResultOutput() AllowlistsResultOutput {
	return o
}

func (o AllowlistsResultOutput) ToAllowlistsResultOutputWithContext(ctx context.Context) AllowlistsResultOutput {
	return o
}

// The list of allowed list.
func (o AllowlistsResultOutput) AllowLists() AllowlistsAllowListArrayOutput {
	return o.ApplyT(func(v AllowlistsResult) []AllowlistsAllowList { return v.AllowLists }).(AllowlistsAllowListArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o AllowlistsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AllowlistsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The id of the instance.
func (o AllowlistsResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AllowlistsResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o AllowlistsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AllowlistsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o AllowlistsResultOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v AllowlistsResult) string { return v.RegionId }).(pulumi.StringOutput)
}

// The total count of Scaling Activity query.
func (o AllowlistsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v AllowlistsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(AllowlistsResultOutput{})
}
