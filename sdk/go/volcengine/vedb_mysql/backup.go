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

// Provides a resource to manage vedb mysql backup
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
//			_, err = vedb_mysql.NewBackup(ctx, "fooBackup", &vedb_mysql.BackupArgs{
//				InstanceId: fooInstance.ID(),
//				BackupPolicy: &vedb_mysql.BackupBackupPolicyArgs{
//					BackupTime:            pulumi.String("18:00Z-20:00Z"),
//					FullBackupPeriod:      pulumi.String("Monday,Tuesday,Wednesday"),
//					BackupRetentionPeriod: pulumi.Int(8),
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
// VedbMysqlBackup can be imported using the instance id and backup id, e.g.
//
// ```sh
// $ pulumi import volcengine:vedb_mysql/backup:Backup default instanceID:backupId
// ```
type Backup struct {
	pulumi.CustomResourceState

	// The id of the backup.
	BackupId pulumi.StringOutput `pulumi:"backupId"`
	// Backup method. Currently, only physical backup is supported. The value is Physical.
	BackupMethod pulumi.StringOutput `pulumi:"backupMethod"`
	// Data backup strategy for instances.
	BackupPolicy BackupBackupPolicyOutput `pulumi:"backupPolicy"`
	// Backup type. Currently, only full backup is supported. The value is Full.
	BackupType pulumi.StringOutput `pulumi:"backupType"`
	// The id of the instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewBackup registers a new resource with the given unique name, arguments, and options.
func NewBackup(ctx *pulumi.Context,
	name string, args *BackupArgs, opts ...pulumi.ResourceOption) (*Backup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Backup
	err := ctx.RegisterResource("volcengine:vedb_mysql/backup:Backup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackup gets an existing Backup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupState, opts ...pulumi.ResourceOption) (*Backup, error) {
	var resource Backup
	err := ctx.ReadResource("volcengine:vedb_mysql/backup:Backup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Backup resources.
type backupState struct {
	// The id of the backup.
	BackupId *string `pulumi:"backupId"`
	// Backup method. Currently, only physical backup is supported. The value is Physical.
	BackupMethod *string `pulumi:"backupMethod"`
	// Data backup strategy for instances.
	BackupPolicy *BackupBackupPolicy `pulumi:"backupPolicy"`
	// Backup type. Currently, only full backup is supported. The value is Full.
	BackupType *string `pulumi:"backupType"`
	// The id of the instance.
	InstanceId *string `pulumi:"instanceId"`
}

type BackupState struct {
	// The id of the backup.
	BackupId pulumi.StringPtrInput
	// Backup method. Currently, only physical backup is supported. The value is Physical.
	BackupMethod pulumi.StringPtrInput
	// Data backup strategy for instances.
	BackupPolicy BackupBackupPolicyPtrInput
	// Backup type. Currently, only full backup is supported. The value is Full.
	BackupType pulumi.StringPtrInput
	// The id of the instance.
	InstanceId pulumi.StringPtrInput
}

func (BackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupState)(nil)).Elem()
}

type backupArgs struct {
	// Backup method. Currently, only physical backup is supported. The value is Physical.
	BackupMethod *string `pulumi:"backupMethod"`
	// Data backup strategy for instances.
	BackupPolicy *BackupBackupPolicy `pulumi:"backupPolicy"`
	// Backup type. Currently, only full backup is supported. The value is Full.
	BackupType *string `pulumi:"backupType"`
	// The id of the instance.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a Backup resource.
type BackupArgs struct {
	// Backup method. Currently, only physical backup is supported. The value is Physical.
	BackupMethod pulumi.StringPtrInput
	// Data backup strategy for instances.
	BackupPolicy BackupBackupPolicyPtrInput
	// Backup type. Currently, only full backup is supported. The value is Full.
	BackupType pulumi.StringPtrInput
	// The id of the instance.
	InstanceId pulumi.StringInput
}

func (BackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupArgs)(nil)).Elem()
}

type BackupInput interface {
	pulumi.Input

	ToBackupOutput() BackupOutput
	ToBackupOutputWithContext(ctx context.Context) BackupOutput
}

func (*Backup) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil)).Elem()
}

func (i *Backup) ToBackupOutput() BackupOutput {
	return i.ToBackupOutputWithContext(context.Background())
}

func (i *Backup) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupOutput)
}

// BackupArrayInput is an input type that accepts BackupArray and BackupArrayOutput values.
// You can construct a concrete instance of `BackupArrayInput` via:
//
//	BackupArray{ BackupArgs{...} }
type BackupArrayInput interface {
	pulumi.Input

	ToBackupArrayOutput() BackupArrayOutput
	ToBackupArrayOutputWithContext(context.Context) BackupArrayOutput
}

type BackupArray []BackupInput

func (BackupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Backup)(nil)).Elem()
}

func (i BackupArray) ToBackupArrayOutput() BackupArrayOutput {
	return i.ToBackupArrayOutputWithContext(context.Background())
}

func (i BackupArray) ToBackupArrayOutputWithContext(ctx context.Context) BackupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupArrayOutput)
}

// BackupMapInput is an input type that accepts BackupMap and BackupMapOutput values.
// You can construct a concrete instance of `BackupMapInput` via:
//
//	BackupMap{ "key": BackupArgs{...} }
type BackupMapInput interface {
	pulumi.Input

	ToBackupMapOutput() BackupMapOutput
	ToBackupMapOutputWithContext(context.Context) BackupMapOutput
}

type BackupMap map[string]BackupInput

func (BackupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Backup)(nil)).Elem()
}

func (i BackupMap) ToBackupMapOutput() BackupMapOutput {
	return i.ToBackupMapOutputWithContext(context.Background())
}

func (i BackupMap) ToBackupMapOutputWithContext(ctx context.Context) BackupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupMapOutput)
}

type BackupOutput struct{ *pulumi.OutputState }

func (BackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil)).Elem()
}

func (o BackupOutput) ToBackupOutput() BackupOutput {
	return o
}

func (o BackupOutput) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return o
}

// The id of the backup.
func (o BackupOutput) BackupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.BackupId }).(pulumi.StringOutput)
}

// Backup method. Currently, only physical backup is supported. The value is Physical.
func (o BackupOutput) BackupMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.BackupMethod }).(pulumi.StringOutput)
}

// Data backup strategy for instances.
func (o BackupOutput) BackupPolicy() BackupBackupPolicyOutput {
	return o.ApplyT(func(v *Backup) BackupBackupPolicyOutput { return v.BackupPolicy }).(BackupBackupPolicyOutput)
}

// Backup type. Currently, only full backup is supported. The value is Full.
func (o BackupOutput) BackupType() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.BackupType }).(pulumi.StringOutput)
}

// The id of the instance.
func (o BackupOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type BackupArrayOutput struct{ *pulumi.OutputState }

func (BackupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Backup)(nil)).Elem()
}

func (o BackupArrayOutput) ToBackupArrayOutput() BackupArrayOutput {
	return o
}

func (o BackupArrayOutput) ToBackupArrayOutputWithContext(ctx context.Context) BackupArrayOutput {
	return o
}

func (o BackupArrayOutput) Index(i pulumi.IntInput) BackupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Backup {
		return vs[0].([]*Backup)[vs[1].(int)]
	}).(BackupOutput)
}

type BackupMapOutput struct{ *pulumi.OutputState }

func (BackupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Backup)(nil)).Elem()
}

func (o BackupMapOutput) ToBackupMapOutput() BackupMapOutput {
	return o
}

func (o BackupMapOutput) ToBackupMapOutputWithContext(ctx context.Context) BackupMapOutput {
	return o
}

func (o BackupMapOutput) MapIndex(k pulumi.StringInput) BackupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Backup {
		return vs[0].(map[string]*Backup)[vs[1].(string)]
	}).(BackupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupInput)(nil)).Elem(), &Backup{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupArrayInput)(nil)).Elem(), BackupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupMapInput)(nil)).Elem(), BackupMap{})
	pulumi.RegisterOutputType(BackupOutput{})
	pulumi.RegisterOutputType(BackupArrayOutput{})
	pulumi.RegisterOutputType(BackupMapOutput{})
}
