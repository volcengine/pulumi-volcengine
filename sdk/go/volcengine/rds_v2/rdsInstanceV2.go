// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds_v2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Provides a resource to manage rds instance v2
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/rds_v2"
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
//			_, err = rds_v2.NewRdsInstanceV2(ctx, "fooRdsInstanceV2", &rds_v2.RdsInstanceV2Args{
//				DbEngineVersion: pulumi.String("MySQL_5_7"),
//				NodeInfos: rds_v2.RdsInstanceV2NodeInfoArray{
//					&rds_v2.RdsInstanceV2NodeInfoArgs{
//						NodeType: pulumi.String("Primary"),
//						NodeSpec: pulumi.String("rds.mysql.2c4g"),
//						ZoneId:   pulumi.String(fooZones.Zones[0].Id),
//					},
//					&rds_v2.RdsInstanceV2NodeInfoArgs{
//						NodeType: pulumi.String("Secondary"),
//						NodeSpec: pulumi.String("rds.mysql.2c4g"),
//						ZoneId:   pulumi.String(fooZones.Zones[0].Id),
//					},
//				},
//				StorageType:         pulumi.String("LocalSSD"),
//				StorageSpace:        pulumi.Int(100),
//				VpcId:               fooVpc.ID(),
//				SubnetId:            fooSubnet.ID(),
//				InstanceName:        pulumi.String("tf-test-v2"),
//				LowerCaseTableNames: pulumi.String("1"),
//				ChargeInfo: &rds_v2.RdsInstanceV2ChargeInfoArgs{
//					ChargeType: pulumi.String("PostPaid"),
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
// RDS Instance can be imported using the id, e.g.
//
// ```sh
// $ pulumi import volcengine:rds_v2/rdsInstanceV2:RdsInstanceV2 default mysql-42b38c769c4b
// ```
type RdsInstanceV2 struct {
	pulumi.CustomResourceState

	// Payment methods.
	ChargeInfo RdsInstanceV2ChargeInfoOutput `pulumi:"chargeInfo"`
	// The connection info ot the RDS instance.
	ConnectionInfos RdsInstanceV2ConnectionInfoArrayOutput `pulumi:"connectionInfos"`
	// Instance type. Value:
	// MySQL_5_7
	// MySQL_8_0.
	DbEngineVersion pulumi.StringOutput `pulumi:"dbEngineVersion"`
	// Parameter template ID. It only takes effect when the database type is MySQL/PostgreSQL/SQL_Server. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	DbParamGroupId pulumi.StringPtrOutput `pulumi:"dbParamGroupId"`
	// Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	DbTimeZone pulumi.StringPtrOutput `pulumi:"dbTimeZone"`
	// Instance name. Cannot start with a number or a dash
	// Can only contain Chinese characters, letters, numbers, underscores and dashes
	// The length is limited between 1 ~ 128.
	InstanceName pulumi.StringPtrOutput `pulumi:"instanceName"`
	// The field instanceType is no longer support. The type of Instance.
	//
	// Deprecated: The field instanceType is no longer support.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// Whether the table name is case sensitive, the default value is 1.
	// Ranges:
	// 0: Table names are stored as fixed and table names are case-sensitive.
	// 1: Table names will be stored in lowercase and table names are not case sensitive.
	LowerCaseTableNames pulumi.StringPtrOutput `pulumi:"lowerCaseTableNames"`
	// Instance specification configuration. This parameter is required for RDS for MySQL, RDS for PostgreSQL and MySQL Sharding. There is one and only one Primary node, one and only one Secondary node, and 0-10 Read-Only nodes.
	NodeInfos RdsInstanceV2NodeInfoArrayOutput `pulumi:"nodeInfos"`
	// Subordinate to the project.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// Instance storage space.
	// When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, value range: [20, 3000], unit: GB, increments every 100GB.
	// When the database type is veDB_MySQL/veDB_PostgreSQL, this parameter does not need to be passed.
	StorageSpace pulumi.IntPtrOutput `pulumi:"storageSpace"`
	// Instance storage type. When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, the value is:
	// LocalSSD - local SSD disk
	// When the database type is veDB_MySQL/veDB_PostgreSQL, the value is:
	// DistributedStorage - Distributed Storage.
	StorageType pulumi.StringOutput `pulumi:"storageType"`
	// Subnet ID.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Private network (VPC) ID. You can call the DescribeVpcs query and use this parameter to specify the VPC where the instance is to be created.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewRdsInstanceV2 registers a new resource with the given unique name, arguments, and options.
func NewRdsInstanceV2(ctx *pulumi.Context,
	name string, args *RdsInstanceV2Args, opts ...pulumi.ResourceOption) (*RdsInstanceV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ChargeInfo == nil {
		return nil, errors.New("invalid value for required argument 'ChargeInfo'")
	}
	if args.DbEngineVersion == nil {
		return nil, errors.New("invalid value for required argument 'DbEngineVersion'")
	}
	if args.NodeInfos == nil {
		return nil, errors.New("invalid value for required argument 'NodeInfos'")
	}
	if args.StorageType == nil {
		return nil, errors.New("invalid value for required argument 'StorageType'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RdsInstanceV2
	err := ctx.RegisterResource("volcengine:rds_v2/rdsInstanceV2:RdsInstanceV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdsInstanceV2 gets an existing RdsInstanceV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdsInstanceV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdsInstanceV2State, opts ...pulumi.ResourceOption) (*RdsInstanceV2, error) {
	var resource RdsInstanceV2
	err := ctx.ReadResource("volcengine:rds_v2/rdsInstanceV2:RdsInstanceV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdsInstanceV2 resources.
type rdsInstanceV2State struct {
	// Payment methods.
	ChargeInfo *RdsInstanceV2ChargeInfo `pulumi:"chargeInfo"`
	// The connection info ot the RDS instance.
	ConnectionInfos []RdsInstanceV2ConnectionInfo `pulumi:"connectionInfos"`
	// Instance type. Value:
	// MySQL_5_7
	// MySQL_8_0.
	DbEngineVersion *string `pulumi:"dbEngineVersion"`
	// Parameter template ID. It only takes effect when the database type is MySQL/PostgreSQL/SQL_Server. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	DbParamGroupId *string `pulumi:"dbParamGroupId"`
	// Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	DbTimeZone *string `pulumi:"dbTimeZone"`
	// Instance name. Cannot start with a number or a dash
	// Can only contain Chinese characters, letters, numbers, underscores and dashes
	// The length is limited between 1 ~ 128.
	InstanceName *string `pulumi:"instanceName"`
	// The field instanceType is no longer support. The type of Instance.
	//
	// Deprecated: The field instanceType is no longer support.
	InstanceType *string `pulumi:"instanceType"`
	// Whether the table name is case sensitive, the default value is 1.
	// Ranges:
	// 0: Table names are stored as fixed and table names are case-sensitive.
	// 1: Table names will be stored in lowercase and table names are not case sensitive.
	LowerCaseTableNames *string `pulumi:"lowerCaseTableNames"`
	// Instance specification configuration. This parameter is required for RDS for MySQL, RDS for PostgreSQL and MySQL Sharding. There is one and only one Primary node, one and only one Secondary node, and 0-10 Read-Only nodes.
	NodeInfos []RdsInstanceV2NodeInfo `pulumi:"nodeInfos"`
	// Subordinate to the project.
	ProjectName *string `pulumi:"projectName"`
	// Instance storage space.
	// When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, value range: [20, 3000], unit: GB, increments every 100GB.
	// When the database type is veDB_MySQL/veDB_PostgreSQL, this parameter does not need to be passed.
	StorageSpace *int `pulumi:"storageSpace"`
	// Instance storage type. When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, the value is:
	// LocalSSD - local SSD disk
	// When the database type is veDB_MySQL/veDB_PostgreSQL, the value is:
	// DistributedStorage - Distributed Storage.
	StorageType *string `pulumi:"storageType"`
	// Subnet ID.
	SubnetId *string `pulumi:"subnetId"`
	// Private network (VPC) ID. You can call the DescribeVpcs query and use this parameter to specify the VPC where the instance is to be created.
	VpcId *string `pulumi:"vpcId"`
}

type RdsInstanceV2State struct {
	// Payment methods.
	ChargeInfo RdsInstanceV2ChargeInfoPtrInput
	// The connection info ot the RDS instance.
	ConnectionInfos RdsInstanceV2ConnectionInfoArrayInput
	// Instance type. Value:
	// MySQL_5_7
	// MySQL_8_0.
	DbEngineVersion pulumi.StringPtrInput
	// Parameter template ID. It only takes effect when the database type is MySQL/PostgreSQL/SQL_Server. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	DbParamGroupId pulumi.StringPtrInput
	// Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	DbTimeZone pulumi.StringPtrInput
	// Instance name. Cannot start with a number or a dash
	// Can only contain Chinese characters, letters, numbers, underscores and dashes
	// The length is limited between 1 ~ 128.
	InstanceName pulumi.StringPtrInput
	// The field instanceType is no longer support. The type of Instance.
	//
	// Deprecated: The field instanceType is no longer support.
	InstanceType pulumi.StringPtrInput
	// Whether the table name is case sensitive, the default value is 1.
	// Ranges:
	// 0: Table names are stored as fixed and table names are case-sensitive.
	// 1: Table names will be stored in lowercase and table names are not case sensitive.
	LowerCaseTableNames pulumi.StringPtrInput
	// Instance specification configuration. This parameter is required for RDS for MySQL, RDS for PostgreSQL and MySQL Sharding. There is one and only one Primary node, one and only one Secondary node, and 0-10 Read-Only nodes.
	NodeInfos RdsInstanceV2NodeInfoArrayInput
	// Subordinate to the project.
	ProjectName pulumi.StringPtrInput
	// Instance storage space.
	// When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, value range: [20, 3000], unit: GB, increments every 100GB.
	// When the database type is veDB_MySQL/veDB_PostgreSQL, this parameter does not need to be passed.
	StorageSpace pulumi.IntPtrInput
	// Instance storage type. When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, the value is:
	// LocalSSD - local SSD disk
	// When the database type is veDB_MySQL/veDB_PostgreSQL, the value is:
	// DistributedStorage - Distributed Storage.
	StorageType pulumi.StringPtrInput
	// Subnet ID.
	SubnetId pulumi.StringPtrInput
	// Private network (VPC) ID. You can call the DescribeVpcs query and use this parameter to specify the VPC where the instance is to be created.
	VpcId pulumi.StringPtrInput
}

func (RdsInstanceV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsInstanceV2State)(nil)).Elem()
}

type rdsInstanceV2Args struct {
	// Payment methods.
	ChargeInfo RdsInstanceV2ChargeInfo `pulumi:"chargeInfo"`
	// Instance type. Value:
	// MySQL_5_7
	// MySQL_8_0.
	DbEngineVersion string `pulumi:"dbEngineVersion"`
	// Parameter template ID. It only takes effect when the database type is MySQL/PostgreSQL/SQL_Server. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	DbParamGroupId *string `pulumi:"dbParamGroupId"`
	// Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	DbTimeZone *string `pulumi:"dbTimeZone"`
	// Instance name. Cannot start with a number or a dash
	// Can only contain Chinese characters, letters, numbers, underscores and dashes
	// The length is limited between 1 ~ 128.
	InstanceName *string `pulumi:"instanceName"`
	// The field instanceType is no longer support. The type of Instance.
	//
	// Deprecated: The field instanceType is no longer support.
	InstanceType *string `pulumi:"instanceType"`
	// Whether the table name is case sensitive, the default value is 1.
	// Ranges:
	// 0: Table names are stored as fixed and table names are case-sensitive.
	// 1: Table names will be stored in lowercase and table names are not case sensitive.
	LowerCaseTableNames *string `pulumi:"lowerCaseTableNames"`
	// Instance specification configuration. This parameter is required for RDS for MySQL, RDS for PostgreSQL and MySQL Sharding. There is one and only one Primary node, one and only one Secondary node, and 0-10 Read-Only nodes.
	NodeInfos []RdsInstanceV2NodeInfo `pulumi:"nodeInfos"`
	// Subordinate to the project.
	ProjectName *string `pulumi:"projectName"`
	// Instance storage space.
	// When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, value range: [20, 3000], unit: GB, increments every 100GB.
	// When the database type is veDB_MySQL/veDB_PostgreSQL, this parameter does not need to be passed.
	StorageSpace *int `pulumi:"storageSpace"`
	// Instance storage type. When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, the value is:
	// LocalSSD - local SSD disk
	// When the database type is veDB_MySQL/veDB_PostgreSQL, the value is:
	// DistributedStorage - Distributed Storage.
	StorageType string `pulumi:"storageType"`
	// Subnet ID.
	SubnetId string `pulumi:"subnetId"`
	// Private network (VPC) ID. You can call the DescribeVpcs query and use this parameter to specify the VPC where the instance is to be created.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a RdsInstanceV2 resource.
type RdsInstanceV2Args struct {
	// Payment methods.
	ChargeInfo RdsInstanceV2ChargeInfoInput
	// Instance type. Value:
	// MySQL_5_7
	// MySQL_8_0.
	DbEngineVersion pulumi.StringInput
	// Parameter template ID. It only takes effect when the database type is MySQL/PostgreSQL/SQL_Server. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	DbParamGroupId pulumi.StringPtrInput
	// Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	DbTimeZone pulumi.StringPtrInput
	// Instance name. Cannot start with a number or a dash
	// Can only contain Chinese characters, letters, numbers, underscores and dashes
	// The length is limited between 1 ~ 128.
	InstanceName pulumi.StringPtrInput
	// The field instanceType is no longer support. The type of Instance.
	//
	// Deprecated: The field instanceType is no longer support.
	InstanceType pulumi.StringPtrInput
	// Whether the table name is case sensitive, the default value is 1.
	// Ranges:
	// 0: Table names are stored as fixed and table names are case-sensitive.
	// 1: Table names will be stored in lowercase and table names are not case sensitive.
	LowerCaseTableNames pulumi.StringPtrInput
	// Instance specification configuration. This parameter is required for RDS for MySQL, RDS for PostgreSQL and MySQL Sharding. There is one and only one Primary node, one and only one Secondary node, and 0-10 Read-Only nodes.
	NodeInfos RdsInstanceV2NodeInfoArrayInput
	// Subordinate to the project.
	ProjectName pulumi.StringPtrInput
	// Instance storage space.
	// When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, value range: [20, 3000], unit: GB, increments every 100GB.
	// When the database type is veDB_MySQL/veDB_PostgreSQL, this parameter does not need to be passed.
	StorageSpace pulumi.IntPtrInput
	// Instance storage type. When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, the value is:
	// LocalSSD - local SSD disk
	// When the database type is veDB_MySQL/veDB_PostgreSQL, the value is:
	// DistributedStorage - Distributed Storage.
	StorageType pulumi.StringInput
	// Subnet ID.
	SubnetId pulumi.StringInput
	// Private network (VPC) ID. You can call the DescribeVpcs query and use this parameter to specify the VPC where the instance is to be created.
	VpcId pulumi.StringInput
}

func (RdsInstanceV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*rdsInstanceV2Args)(nil)).Elem()
}

type RdsInstanceV2Input interface {
	pulumi.Input

	ToRdsInstanceV2Output() RdsInstanceV2Output
	ToRdsInstanceV2OutputWithContext(ctx context.Context) RdsInstanceV2Output
}

func (*RdsInstanceV2) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsInstanceV2)(nil)).Elem()
}

func (i *RdsInstanceV2) ToRdsInstanceV2Output() RdsInstanceV2Output {
	return i.ToRdsInstanceV2OutputWithContext(context.Background())
}

func (i *RdsInstanceV2) ToRdsInstanceV2OutputWithContext(ctx context.Context) RdsInstanceV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(RdsInstanceV2Output)
}

// RdsInstanceV2ArrayInput is an input type that accepts RdsInstanceV2Array and RdsInstanceV2ArrayOutput values.
// You can construct a concrete instance of `RdsInstanceV2ArrayInput` via:
//
//	RdsInstanceV2Array{ RdsInstanceV2Args{...} }
type RdsInstanceV2ArrayInput interface {
	pulumi.Input

	ToRdsInstanceV2ArrayOutput() RdsInstanceV2ArrayOutput
	ToRdsInstanceV2ArrayOutputWithContext(context.Context) RdsInstanceV2ArrayOutput
}

type RdsInstanceV2Array []RdsInstanceV2Input

func (RdsInstanceV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsInstanceV2)(nil)).Elem()
}

func (i RdsInstanceV2Array) ToRdsInstanceV2ArrayOutput() RdsInstanceV2ArrayOutput {
	return i.ToRdsInstanceV2ArrayOutputWithContext(context.Background())
}

func (i RdsInstanceV2Array) ToRdsInstanceV2ArrayOutputWithContext(ctx context.Context) RdsInstanceV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsInstanceV2ArrayOutput)
}

// RdsInstanceV2MapInput is an input type that accepts RdsInstanceV2Map and RdsInstanceV2MapOutput values.
// You can construct a concrete instance of `RdsInstanceV2MapInput` via:
//
//	RdsInstanceV2Map{ "key": RdsInstanceV2Args{...} }
type RdsInstanceV2MapInput interface {
	pulumi.Input

	ToRdsInstanceV2MapOutput() RdsInstanceV2MapOutput
	ToRdsInstanceV2MapOutputWithContext(context.Context) RdsInstanceV2MapOutput
}

type RdsInstanceV2Map map[string]RdsInstanceV2Input

func (RdsInstanceV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsInstanceV2)(nil)).Elem()
}

func (i RdsInstanceV2Map) ToRdsInstanceV2MapOutput() RdsInstanceV2MapOutput {
	return i.ToRdsInstanceV2MapOutputWithContext(context.Background())
}

func (i RdsInstanceV2Map) ToRdsInstanceV2MapOutputWithContext(ctx context.Context) RdsInstanceV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdsInstanceV2MapOutput)
}

type RdsInstanceV2Output struct{ *pulumi.OutputState }

func (RdsInstanceV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**RdsInstanceV2)(nil)).Elem()
}

func (o RdsInstanceV2Output) ToRdsInstanceV2Output() RdsInstanceV2Output {
	return o
}

func (o RdsInstanceV2Output) ToRdsInstanceV2OutputWithContext(ctx context.Context) RdsInstanceV2Output {
	return o
}

// Payment methods.
func (o RdsInstanceV2Output) ChargeInfo() RdsInstanceV2ChargeInfoOutput {
	return o.ApplyT(func(v *RdsInstanceV2) RdsInstanceV2ChargeInfoOutput { return v.ChargeInfo }).(RdsInstanceV2ChargeInfoOutput)
}

// The connection info ot the RDS instance.
func (o RdsInstanceV2Output) ConnectionInfos() RdsInstanceV2ConnectionInfoArrayOutput {
	return o.ApplyT(func(v *RdsInstanceV2) RdsInstanceV2ConnectionInfoArrayOutput { return v.ConnectionInfos }).(RdsInstanceV2ConnectionInfoArrayOutput)
}

// Instance type. Value:
// MySQL_5_7
// MySQL_8_0.
func (o RdsInstanceV2Output) DbEngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstanceV2) pulumi.StringOutput { return v.DbEngineVersion }).(pulumi.StringOutput)
}

// Parameter template ID. It only takes effect when the database type is MySQL/PostgreSQL/SQL_Server. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
func (o RdsInstanceV2Output) DbParamGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsInstanceV2) pulumi.StringPtrOutput { return v.DbParamGroupId }).(pulumi.StringPtrOutput)
}

// Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
func (o RdsInstanceV2Output) DbTimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsInstanceV2) pulumi.StringPtrOutput { return v.DbTimeZone }).(pulumi.StringPtrOutput)
}

// Instance name. Cannot start with a number or a dash
// Can only contain Chinese characters, letters, numbers, underscores and dashes
// The length is limited between 1 ~ 128.
func (o RdsInstanceV2Output) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsInstanceV2) pulumi.StringPtrOutput { return v.InstanceName }).(pulumi.StringPtrOutput)
}

// The field instanceType is no longer support. The type of Instance.
//
// Deprecated: The field instanceType is no longer support.
func (o RdsInstanceV2Output) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstanceV2) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// Whether the table name is case sensitive, the default value is 1.
// Ranges:
// 0: Table names are stored as fixed and table names are case-sensitive.
// 1: Table names will be stored in lowercase and table names are not case sensitive.
func (o RdsInstanceV2Output) LowerCaseTableNames() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdsInstanceV2) pulumi.StringPtrOutput { return v.LowerCaseTableNames }).(pulumi.StringPtrOutput)
}

// Instance specification configuration. This parameter is required for RDS for MySQL, RDS for PostgreSQL and MySQL Sharding. There is one and only one Primary node, one and only one Secondary node, and 0-10 Read-Only nodes.
func (o RdsInstanceV2Output) NodeInfos() RdsInstanceV2NodeInfoArrayOutput {
	return o.ApplyT(func(v *RdsInstanceV2) RdsInstanceV2NodeInfoArrayOutput { return v.NodeInfos }).(RdsInstanceV2NodeInfoArrayOutput)
}

// Subordinate to the project.
func (o RdsInstanceV2Output) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstanceV2) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

// Instance storage space.
// When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, value range: [20, 3000], unit: GB, increments every 100GB.
// When the database type is veDB_MySQL/veDB_PostgreSQL, this parameter does not need to be passed.
func (o RdsInstanceV2Output) StorageSpace() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RdsInstanceV2) pulumi.IntPtrOutput { return v.StorageSpace }).(pulumi.IntPtrOutput)
}

// Instance storage type. When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, the value is:
// LocalSSD - local SSD disk
// When the database type is veDB_MySQL/veDB_PostgreSQL, the value is:
// DistributedStorage - Distributed Storage.
func (o RdsInstanceV2Output) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstanceV2) pulumi.StringOutput { return v.StorageType }).(pulumi.StringOutput)
}

// Subnet ID.
func (o RdsInstanceV2Output) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstanceV2) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// Private network (VPC) ID. You can call the DescribeVpcs query and use this parameter to specify the VPC where the instance is to be created.
func (o RdsInstanceV2Output) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdsInstanceV2) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type RdsInstanceV2ArrayOutput struct{ *pulumi.OutputState }

func (RdsInstanceV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdsInstanceV2)(nil)).Elem()
}

func (o RdsInstanceV2ArrayOutput) ToRdsInstanceV2ArrayOutput() RdsInstanceV2ArrayOutput {
	return o
}

func (o RdsInstanceV2ArrayOutput) ToRdsInstanceV2ArrayOutputWithContext(ctx context.Context) RdsInstanceV2ArrayOutput {
	return o
}

func (o RdsInstanceV2ArrayOutput) Index(i pulumi.IntInput) RdsInstanceV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RdsInstanceV2 {
		return vs[0].([]*RdsInstanceV2)[vs[1].(int)]
	}).(RdsInstanceV2Output)
}

type RdsInstanceV2MapOutput struct{ *pulumi.OutputState }

func (RdsInstanceV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdsInstanceV2)(nil)).Elem()
}

func (o RdsInstanceV2MapOutput) ToRdsInstanceV2MapOutput() RdsInstanceV2MapOutput {
	return o
}

func (o RdsInstanceV2MapOutput) ToRdsInstanceV2MapOutputWithContext(ctx context.Context) RdsInstanceV2MapOutput {
	return o
}

func (o RdsInstanceV2MapOutput) MapIndex(k pulumi.StringInput) RdsInstanceV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RdsInstanceV2 {
		return vs[0].(map[string]*RdsInstanceV2)[vs[1].(string)]
	}).(RdsInstanceV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RdsInstanceV2Input)(nil)).Elem(), &RdsInstanceV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsInstanceV2ArrayInput)(nil)).Elem(), RdsInstanceV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdsInstanceV2MapInput)(nil)).Elem(), RdsInstanceV2Map{})
	pulumi.RegisterOutputType(RdsInstanceV2Output{})
	pulumi.RegisterOutputType(RdsInstanceV2ArrayOutput{})
	pulumi.RegisterOutputType(RdsInstanceV2MapOutput{})
}
