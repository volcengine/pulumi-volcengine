// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds_mssql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage rds mssql instance
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/rds_mssql"
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
//				ZoneId:     *pulumi.String(fooZones.Zones[0].Id),
//				VpcId:      fooVpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rds_mssql.NewInstance(ctx, "fooInstance", &rds_mssql.InstanceArgs{
//				DbEngineVersion: pulumi.String("SQLServer_2019_Std"),
//				InstanceType:    pulumi.String("HA"),
//				NodeSpec:        pulumi.String("rds.mssql.se.ha.d2.2c4g"),
//				StorageSpace:    pulumi.Int(20),
//				SubnetIds: pulumi.StringArray{
//					fooSubnet.ID(),
//				},
//				SuperAccountPassword: pulumi.String("Tftest110"),
//				InstanceName:         pulumi.String("acc-test-mssql"),
//				ProjectName:          pulumi.String("default"),
//				ChargeInfo: &rds_mssql.InstanceChargeInfoArgs{
//					ChargeType: pulumi.String("PostPaid"),
//				},
//				Tags: rds_mssql.InstanceTagArray{
//					&rds_mssql.InstanceTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
//					},
//				},
//				BackupTime: pulumi.String("18:00Z-19:00Z"),
//				FullBackupPeriods: pulumi.StringArray{
//					pulumi.String("Monday"),
//					pulumi.String("Tuesday"),
//				},
//				BackupRetentionPeriod: pulumi.Int(14),
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
// Rds Mssql Instance can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:rds_mssql/instance:Instance default resource_id
//
// ```
type Instance struct {
	pulumi.CustomResourceState

	// Data backup retention days, value range: 7~30.
	// This field is valid and required when updating the backup plan of instance.
	BackupRetentionPeriod pulumi.IntPtrOutput `pulumi:"backupRetentionPeriod"`
	// The time window for starting the backup task is one hour interval.
	// This field is valid and required when updating the backup plan of instance.
	BackupTime pulumi.StringPtrOutput `pulumi:"backupTime"`
	// The charge info.
	ChargeInfo InstanceChargeInfoOutput `pulumi:"chargeInfo"`
	// The Compatible version. Valid values: `SQLServer_2019_Std`, `SQLServer_2019_Web`, `SQLServer_2019_Ent`.
	DbEngineVersion pulumi.StringOutput `pulumi:"dbEngineVersion"`
	// Full backup cycle. Multiple values separated by commas. The values are as follows: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday.
	// This field is valid and required when updating the backup plan of instance.
	FullBackupPeriods pulumi.StringArrayOutput `pulumi:"fullBackupPeriods"`
	// Name of the instance.
	InstanceName pulumi.StringPtrOutput `pulumi:"instanceName"`
	// The Instance type. When the value of the `dbEngineVersion` is `SQLServer_2019_Std`, the value of this field can be `HA` or `Basic`.When the value of the `dbEngineVersion` is `SQLServer_2019_Ent`, the value of this field can be `Cluster` or `Basic`.When the value of the `dbEngineVersion` is `SQLServer_2019_Web`, the value of this field can be `Basic`.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// The node specification.
	NodeSpec pulumi.StringOutput `pulumi:"nodeSpec"`
	// The project name.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// Storage space size, measured in GiB. The range of values is 20GiB to 4000GiB, with a step size of 10GiB.
	StorageSpace pulumi.IntOutput `pulumi:"storageSpace"`
	// The subnet id of the instance node. When creating an instance that includes primary and backup nodes and needs to deploy primary and backup nodes across availability zones, you can specify two subnet_id. By default, the first is the primary node availability zone, and the second is the backup node availability zone.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// The super account password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	SuperAccountPassword pulumi.StringOutput `pulumi:"superAccountPassword"`
	// Tags.
	Tags InstanceTagArrayOutput `pulumi:"tags"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ChargeInfo == nil {
		return nil, errors.New("invalid value for required argument 'ChargeInfo'")
	}
	if args.DbEngineVersion == nil {
		return nil, errors.New("invalid value for required argument 'DbEngineVersion'")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	if args.NodeSpec == nil {
		return nil, errors.New("invalid value for required argument 'NodeSpec'")
	}
	if args.StorageSpace == nil {
		return nil, errors.New("invalid value for required argument 'StorageSpace'")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	if args.SuperAccountPassword == nil {
		return nil, errors.New("invalid value for required argument 'SuperAccountPassword'")
	}
	if args.SuperAccountPassword != nil {
		args.SuperAccountPassword = pulumi.ToSecret(args.SuperAccountPassword).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"superAccountPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("volcengine:rds_mssql/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("volcengine:rds_mssql/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Data backup retention days, value range: 7~30.
	// This field is valid and required when updating the backup plan of instance.
	BackupRetentionPeriod *int `pulumi:"backupRetentionPeriod"`
	// The time window for starting the backup task is one hour interval.
	// This field is valid and required when updating the backup plan of instance.
	BackupTime *string `pulumi:"backupTime"`
	// The charge info.
	ChargeInfo *InstanceChargeInfo `pulumi:"chargeInfo"`
	// The Compatible version. Valid values: `SQLServer_2019_Std`, `SQLServer_2019_Web`, `SQLServer_2019_Ent`.
	DbEngineVersion *string `pulumi:"dbEngineVersion"`
	// Full backup cycle. Multiple values separated by commas. The values are as follows: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday.
	// This field is valid and required when updating the backup plan of instance.
	FullBackupPeriods []string `pulumi:"fullBackupPeriods"`
	// Name of the instance.
	InstanceName *string `pulumi:"instanceName"`
	// The Instance type. When the value of the `dbEngineVersion` is `SQLServer_2019_Std`, the value of this field can be `HA` or `Basic`.When the value of the `dbEngineVersion` is `SQLServer_2019_Ent`, the value of this field can be `Cluster` or `Basic`.When the value of the `dbEngineVersion` is `SQLServer_2019_Web`, the value of this field can be `Basic`.
	InstanceType *string `pulumi:"instanceType"`
	// The node specification.
	NodeSpec *string `pulumi:"nodeSpec"`
	// The project name.
	ProjectName *string `pulumi:"projectName"`
	// Storage space size, measured in GiB. The range of values is 20GiB to 4000GiB, with a step size of 10GiB.
	StorageSpace *int `pulumi:"storageSpace"`
	// The subnet id of the instance node. When creating an instance that includes primary and backup nodes and needs to deploy primary and backup nodes across availability zones, you can specify two subnet_id. By default, the first is the primary node availability zone, and the second is the backup node availability zone.
	SubnetIds []string `pulumi:"subnetIds"`
	// The super account password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	SuperAccountPassword *string `pulumi:"superAccountPassword"`
	// Tags.
	Tags []InstanceTag `pulumi:"tags"`
}

type InstanceState struct {
	// Data backup retention days, value range: 7~30.
	// This field is valid and required when updating the backup plan of instance.
	BackupRetentionPeriod pulumi.IntPtrInput
	// The time window for starting the backup task is one hour interval.
	// This field is valid and required when updating the backup plan of instance.
	BackupTime pulumi.StringPtrInput
	// The charge info.
	ChargeInfo InstanceChargeInfoPtrInput
	// The Compatible version. Valid values: `SQLServer_2019_Std`, `SQLServer_2019_Web`, `SQLServer_2019_Ent`.
	DbEngineVersion pulumi.StringPtrInput
	// Full backup cycle. Multiple values separated by commas. The values are as follows: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday.
	// This field is valid and required when updating the backup plan of instance.
	FullBackupPeriods pulumi.StringArrayInput
	// Name of the instance.
	InstanceName pulumi.StringPtrInput
	// The Instance type. When the value of the `dbEngineVersion` is `SQLServer_2019_Std`, the value of this field can be `HA` or `Basic`.When the value of the `dbEngineVersion` is `SQLServer_2019_Ent`, the value of this field can be `Cluster` or `Basic`.When the value of the `dbEngineVersion` is `SQLServer_2019_Web`, the value of this field can be `Basic`.
	InstanceType pulumi.StringPtrInput
	// The node specification.
	NodeSpec pulumi.StringPtrInput
	// The project name.
	ProjectName pulumi.StringPtrInput
	// Storage space size, measured in GiB. The range of values is 20GiB to 4000GiB, with a step size of 10GiB.
	StorageSpace pulumi.IntPtrInput
	// The subnet id of the instance node. When creating an instance that includes primary and backup nodes and needs to deploy primary and backup nodes across availability zones, you can specify two subnet_id. By default, the first is the primary node availability zone, and the second is the backup node availability zone.
	SubnetIds pulumi.StringArrayInput
	// The super account password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	SuperAccountPassword pulumi.StringPtrInput
	// Tags.
	Tags InstanceTagArrayInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Data backup retention days, value range: 7~30.
	// This field is valid and required when updating the backup plan of instance.
	BackupRetentionPeriod *int `pulumi:"backupRetentionPeriod"`
	// The time window for starting the backup task is one hour interval.
	// This field is valid and required when updating the backup plan of instance.
	BackupTime *string `pulumi:"backupTime"`
	// The charge info.
	ChargeInfo InstanceChargeInfo `pulumi:"chargeInfo"`
	// The Compatible version. Valid values: `SQLServer_2019_Std`, `SQLServer_2019_Web`, `SQLServer_2019_Ent`.
	DbEngineVersion string `pulumi:"dbEngineVersion"`
	// Full backup cycle. Multiple values separated by commas. The values are as follows: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday.
	// This field is valid and required when updating the backup plan of instance.
	FullBackupPeriods []string `pulumi:"fullBackupPeriods"`
	// Name of the instance.
	InstanceName *string `pulumi:"instanceName"`
	// The Instance type. When the value of the `dbEngineVersion` is `SQLServer_2019_Std`, the value of this field can be `HA` or `Basic`.When the value of the `dbEngineVersion` is `SQLServer_2019_Ent`, the value of this field can be `Cluster` or `Basic`.When the value of the `dbEngineVersion` is `SQLServer_2019_Web`, the value of this field can be `Basic`.
	InstanceType string `pulumi:"instanceType"`
	// The node specification.
	NodeSpec string `pulumi:"nodeSpec"`
	// The project name.
	ProjectName *string `pulumi:"projectName"`
	// Storage space size, measured in GiB. The range of values is 20GiB to 4000GiB, with a step size of 10GiB.
	StorageSpace int `pulumi:"storageSpace"`
	// The subnet id of the instance node. When creating an instance that includes primary and backup nodes and needs to deploy primary and backup nodes across availability zones, you can specify two subnet_id. By default, the first is the primary node availability zone, and the second is the backup node availability zone.
	SubnetIds []string `pulumi:"subnetIds"`
	// The super account password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	SuperAccountPassword string `pulumi:"superAccountPassword"`
	// Tags.
	Tags []InstanceTag `pulumi:"tags"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Data backup retention days, value range: 7~30.
	// This field is valid and required when updating the backup plan of instance.
	BackupRetentionPeriod pulumi.IntPtrInput
	// The time window for starting the backup task is one hour interval.
	// This field is valid and required when updating the backup plan of instance.
	BackupTime pulumi.StringPtrInput
	// The charge info.
	ChargeInfo InstanceChargeInfoInput
	// The Compatible version. Valid values: `SQLServer_2019_Std`, `SQLServer_2019_Web`, `SQLServer_2019_Ent`.
	DbEngineVersion pulumi.StringInput
	// Full backup cycle. Multiple values separated by commas. The values are as follows: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday.
	// This field is valid and required when updating the backup plan of instance.
	FullBackupPeriods pulumi.StringArrayInput
	// Name of the instance.
	InstanceName pulumi.StringPtrInput
	// The Instance type. When the value of the `dbEngineVersion` is `SQLServer_2019_Std`, the value of this field can be `HA` or `Basic`.When the value of the `dbEngineVersion` is `SQLServer_2019_Ent`, the value of this field can be `Cluster` or `Basic`.When the value of the `dbEngineVersion` is `SQLServer_2019_Web`, the value of this field can be `Basic`.
	InstanceType pulumi.StringInput
	// The node specification.
	NodeSpec pulumi.StringInput
	// The project name.
	ProjectName pulumi.StringPtrInput
	// Storage space size, measured in GiB. The range of values is 20GiB to 4000GiB, with a step size of 10GiB.
	StorageSpace pulumi.IntInput
	// The subnet id of the instance node. When creating an instance that includes primary and backup nodes and needs to deploy primary and backup nodes across availability zones, you can specify two subnet_id. By default, the first is the primary node availability zone, and the second is the backup node availability zone.
	SubnetIds pulumi.StringArrayInput
	// The super account password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	SuperAccountPassword pulumi.StringInput
	// Tags.
	Tags InstanceTagArrayInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// Data backup retention days, value range: 7~30.
// This field is valid and required when updating the backup plan of instance.
func (o InstanceOutput) BackupRetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.BackupRetentionPeriod }).(pulumi.IntPtrOutput)
}

// The time window for starting the backup task is one hour interval.
// This field is valid and required when updating the backup plan of instance.
func (o InstanceOutput) BackupTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.BackupTime }).(pulumi.StringPtrOutput)
}

// The charge info.
func (o InstanceOutput) ChargeInfo() InstanceChargeInfoOutput {
	return o.ApplyT(func(v *Instance) InstanceChargeInfoOutput { return v.ChargeInfo }).(InstanceChargeInfoOutput)
}

// The Compatible version. Valid values: `SQLServer_2019_Std`, `SQLServer_2019_Web`, `SQLServer_2019_Ent`.
func (o InstanceOutput) DbEngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.DbEngineVersion }).(pulumi.StringOutput)
}

// Full backup cycle. Multiple values separated by commas. The values are as follows: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday.
// This field is valid and required when updating the backup plan of instance.
func (o InstanceOutput) FullBackupPeriods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.FullBackupPeriods }).(pulumi.StringArrayOutput)
}

// Name of the instance.
func (o InstanceOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.InstanceName }).(pulumi.StringPtrOutput)
}

// The Instance type. When the value of the `dbEngineVersion` is `SQLServer_2019_Std`, the value of this field can be `HA` or `Basic`.When the value of the `dbEngineVersion` is `SQLServer_2019_Ent`, the value of this field can be `Cluster` or `Basic`.When the value of the `dbEngineVersion` is `SQLServer_2019_Web`, the value of this field can be `Basic`.
func (o InstanceOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// The node specification.
func (o InstanceOutput) NodeSpec() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.NodeSpec }).(pulumi.StringOutput)
}

// The project name.
func (o InstanceOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

// Storage space size, measured in GiB. The range of values is 20GiB to 4000GiB, with a step size of 10GiB.
func (o InstanceOutput) StorageSpace() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.StorageSpace }).(pulumi.IntOutput)
}

// The subnet id of the instance node. When creating an instance that includes primary and backup nodes and needs to deploy primary and backup nodes across availability zones, you can specify two subnet_id. By default, the first is the primary node availability zone, and the second is the backup node availability zone.
func (o InstanceOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// The super account password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
func (o InstanceOutput) SuperAccountPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.SuperAccountPassword }).(pulumi.StringOutput)
}

// Tags.
func (o InstanceOutput) Tags() InstanceTagArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceTagArrayOutput { return v.Tags }).(InstanceTagArrayOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}