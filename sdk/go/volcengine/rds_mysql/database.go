// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds_mysql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage rds mysql database
// ## Example Usage
//
// ```go
// package main
//
// import (
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
//				VpcName:   pulumi.String("acc-test-project1"),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			fooSubnet, err := vpc.NewSubnet(ctx, "fooSubnet", &vpc.SubnetArgs{
//				SubnetName: pulumi.String("acc-subnet-test-2"),
//				CidrBlock:  pulumi.String("172.16.0.0/24"),
//				ZoneId:     pulumi.String(fooZones.Zones[0].Id),
//				VpcId:      fooVpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			fooInstance, err := rds_mysql.NewInstance(ctx, "fooInstance", &rds_mysql.InstanceArgs{
//				DbEngineVersion:     pulumi.String("MySQL_5_7"),
//				NodeSpec:            pulumi.String("rds.mysql.1c2g"),
//				PrimaryZoneId:       pulumi.String(fooZones.Zones[0].Id),
//				SecondaryZoneId:     pulumi.String(fooZones.Zones[0].Id),
//				StorageSpace:        pulumi.Int(80),
//				SubnetId:            fooSubnet.ID(),
//				InstanceName:        pulumi.String("acc-test"),
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
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rds_mysql.NewDatabase(ctx, "fooDatabase", &rds_mysql.DatabaseArgs{
//				DbName:     pulumi.String("acc-test"),
//				InstanceId: fooInstance.ID(),
//				DbDesc:     pulumi.String("test-update"),
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
// Database can be imported using the instanceId:dbName, e.g.
//
// ```sh
// $ pulumi import volcengine:rds_mysql/database:Database default mysql-42b38c769c4b:dbname
// ```
type Database struct {
	pulumi.CustomResourceState

	// Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
	CharacterSetName pulumi.StringPtrOutput `pulumi:"characterSetName"`
	// The description information of the database, with a length not exceeding 256 characters. This field is optional. If this field is not set, or if this field is set but the length of the description information is 0, then the description information is empty.
	DbDesc pulumi.StringPtrOutput `pulumi:"dbDesc"`
	// Name database.
	// illustrate:
	// Unique name.
	// The length is 2~64 characters.
	// Start with a letter and end with a letter or number.
	// Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
	// Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
	DbName pulumi.StringOutput `pulumi:"dbName"`
	// The ID of the RDS instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbName == nil {
		return nil, errors.New("invalid value for required argument 'DbName'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Database
	err := ctx.RegisterResource("volcengine:rds_mysql/database:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("volcengine:rds_mysql/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
	CharacterSetName *string `pulumi:"characterSetName"`
	// The description information of the database, with a length not exceeding 256 characters. This field is optional. If this field is not set, or if this field is set but the length of the description information is 0, then the description information is empty.
	DbDesc *string `pulumi:"dbDesc"`
	// Name database.
	// illustrate:
	// Unique name.
	// The length is 2~64 characters.
	// Start with a letter and end with a letter or number.
	// Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
	// Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
	DbName *string `pulumi:"dbName"`
	// The ID of the RDS instance.
	InstanceId *string `pulumi:"instanceId"`
}

type DatabaseState struct {
	// Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
	CharacterSetName pulumi.StringPtrInput
	// The description information of the database, with a length not exceeding 256 characters. This field is optional. If this field is not set, or if this field is set but the length of the description information is 0, then the description information is empty.
	DbDesc pulumi.StringPtrInput
	// Name database.
	// illustrate:
	// Unique name.
	// The length is 2~64 characters.
	// Start with a letter and end with a letter or number.
	// Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
	// Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
	DbName pulumi.StringPtrInput
	// The ID of the RDS instance.
	InstanceId pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
	CharacterSetName *string `pulumi:"characterSetName"`
	// The description information of the database, with a length not exceeding 256 characters. This field is optional. If this field is not set, or if this field is set but the length of the description information is 0, then the description information is empty.
	DbDesc *string `pulumi:"dbDesc"`
	// Name database.
	// illustrate:
	// Unique name.
	// The length is 2~64 characters.
	// Start with a letter and end with a letter or number.
	// Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
	// Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
	DbName string `pulumi:"dbName"`
	// The ID of the RDS instance.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
	CharacterSetName pulumi.StringPtrInput
	// The description information of the database, with a length not exceeding 256 characters. This field is optional. If this field is not set, or if this field is set but the length of the description information is 0, then the description information is empty.
	DbDesc pulumi.StringPtrInput
	// Name database.
	// illustrate:
	// Unique name.
	// The length is 2~64 characters.
	// Start with a letter and end with a letter or number.
	// Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
	// Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
	DbName pulumi.StringInput
	// The ID of the RDS instance.
	InstanceId pulumi.StringInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

// DatabaseArrayInput is an input type that accepts DatabaseArray and DatabaseArrayOutput values.
// You can construct a concrete instance of `DatabaseArrayInput` via:
//
//	DatabaseArray{ DatabaseArgs{...} }
type DatabaseArrayInput interface {
	pulumi.Input

	ToDatabaseArrayOutput() DatabaseArrayOutput
	ToDatabaseArrayOutputWithContext(context.Context) DatabaseArrayOutput
}

type DatabaseArray []DatabaseInput

func (DatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (i DatabaseArray) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return i.ToDatabaseArrayOutputWithContext(context.Background())
}

func (i DatabaseArray) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseArrayOutput)
}

// DatabaseMapInput is an input type that accepts DatabaseMap and DatabaseMapOutput values.
// You can construct a concrete instance of `DatabaseMapInput` via:
//
//	DatabaseMap{ "key": DatabaseArgs{...} }
type DatabaseMapInput interface {
	pulumi.Input

	ToDatabaseMapOutput() DatabaseMapOutput
	ToDatabaseMapOutputWithContext(context.Context) DatabaseMapOutput
}

type DatabaseMap map[string]DatabaseInput

func (DatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (i DatabaseMap) ToDatabaseMapOutput() DatabaseMapOutput {
	return i.ToDatabaseMapOutputWithContext(context.Background())
}

func (i DatabaseMap) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseMapOutput)
}

type DatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

// Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
func (o DatabaseOutput) CharacterSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.CharacterSetName }).(pulumi.StringPtrOutput)
}

// The description information of the database, with a length not exceeding 256 characters. This field is optional. If this field is not set, or if this field is set but the length of the description information is 0, then the description information is empty.
func (o DatabaseOutput) DbDesc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.DbDesc }).(pulumi.StringPtrOutput)
}

// Name database.
// illustrate:
// Unique name.
// The length is 2~64 characters.
// Start with a letter and end with a letter or number.
// Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
// Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
func (o DatabaseOutput) DbName() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.DbName }).(pulumi.StringOutput)
}

// The ID of the RDS instance.
func (o DatabaseOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type DatabaseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) Index(i pulumi.IntInput) DatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Database {
		return vs[0].([]*Database)[vs[1].(int)]
	}).(DatabaseOutput)
}

type DatabaseMapOutput struct{ *pulumi.OutputState }

func (DatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (o DatabaseMapOutput) ToDatabaseMapOutput() DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) MapIndex(k pulumi.StringInput) DatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Database {
		return vs[0].(map[string]*Database)[vs[1].(string)]
	}).(DatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInput)(nil)).Elem(), &Database{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseArrayInput)(nil)).Elem(), DatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseMapInput)(nil)).Elem(), DatabaseMap{})
	pulumi.RegisterOutputType(DatabaseOutput{})
	pulumi.RegisterOutputType(DatabaseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseMapOutput{})
}
