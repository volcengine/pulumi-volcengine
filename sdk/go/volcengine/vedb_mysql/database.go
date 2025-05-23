// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vedb_mysql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage vedb mysql database
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
//			_, err = vedb_mysql.NewDatabase(ctx, "fooDatabase", &vedb_mysql.DatabaseArgs{
//				DbName:     pulumi.String("tf-table"),
//				InstanceId: fooInstance.ID(),
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
// VedbMysqlDatabase can be imported using the instance id and database name, e.g.
//
// ```sh
// $ pulumi import volcengine:vedb_mysql/database:Database default vedbm-r3xq0zdl****:testdb
// ```
type Database struct {
	pulumi.CustomResourceState

	// Database character set: utf8mb4 (default), utf8, latin1, ascii.
	CharacterSetName pulumi.StringOutput `pulumi:"characterSetName"`
	// The name of the database. Naming rules:
	// Unique name. Start with a lowercase letter and end with a letter or number. The length is within 2 to 64 characters.
	// Consist of lowercase letters, numbers, underscores (_), or hyphens (-).
	// The name cannot contain certain reserved words.
	DbName pulumi.StringOutput `pulumi:"dbName"`
	// The id of the instance.
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
	err := ctx.RegisterResource("volcengine:vedb_mysql/database:Database", name, args, &resource, opts...)
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
	err := ctx.ReadResource("volcengine:vedb_mysql/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// Database character set: utf8mb4 (default), utf8, latin1, ascii.
	CharacterSetName *string `pulumi:"characterSetName"`
	// The name of the database. Naming rules:
	// Unique name. Start with a lowercase letter and end with a letter or number. The length is within 2 to 64 characters.
	// Consist of lowercase letters, numbers, underscores (_), or hyphens (-).
	// The name cannot contain certain reserved words.
	DbName *string `pulumi:"dbName"`
	// The id of the instance.
	InstanceId *string `pulumi:"instanceId"`
}

type DatabaseState struct {
	// Database character set: utf8mb4 (default), utf8, latin1, ascii.
	CharacterSetName pulumi.StringPtrInput
	// The name of the database. Naming rules:
	// Unique name. Start with a lowercase letter and end with a letter or number. The length is within 2 to 64 characters.
	// Consist of lowercase letters, numbers, underscores (_), or hyphens (-).
	// The name cannot contain certain reserved words.
	DbName pulumi.StringPtrInput
	// The id of the instance.
	InstanceId pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// Database character set: utf8mb4 (default), utf8, latin1, ascii.
	CharacterSetName *string `pulumi:"characterSetName"`
	// The name of the database. Naming rules:
	// Unique name. Start with a lowercase letter and end with a letter or number. The length is within 2 to 64 characters.
	// Consist of lowercase letters, numbers, underscores (_), or hyphens (-).
	// The name cannot contain certain reserved words.
	DbName string `pulumi:"dbName"`
	// The id of the instance.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// Database character set: utf8mb4 (default), utf8, latin1, ascii.
	CharacterSetName pulumi.StringPtrInput
	// The name of the database. Naming rules:
	// Unique name. Start with a lowercase letter and end with a letter or number. The length is within 2 to 64 characters.
	// Consist of lowercase letters, numbers, underscores (_), or hyphens (-).
	// The name cannot contain certain reserved words.
	DbName pulumi.StringInput
	// The id of the instance.
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

// Database character set: utf8mb4 (default), utf8, latin1, ascii.
func (o DatabaseOutput) CharacterSetName() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CharacterSetName }).(pulumi.StringOutput)
}

// The name of the database. Naming rules:
// Unique name. Start with a lowercase letter and end with a letter or number. The length is within 2 to 64 characters.
// Consist of lowercase letters, numbers, underscores (_), or hyphens (-).
// The name cannot contain certain reserved words.
func (o DatabaseOutput) DbName() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.DbName }).(pulumi.StringOutput)
}

// The id of the instance.
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
