// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vedb_mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of vedb mysql backups
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
//			_ = vedb_mysql.GetBackupsOutput(ctx, vedb_mysql.GetBackupsOutputArgs{
//				InstanceId: fooInstance.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func GetBackups(ctx *pulumi.Context, args *GetBackupsArgs, opts ...pulumi.InvokeOption) (*GetBackupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBackupsResult
	err := ctx.Invoke("volcengine:vedb_mysql/getBackups:getBackups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackups.
type GetBackupsArgs struct {
	// The end time of the backup.
	BackupEndTime *string `pulumi:"backupEndTime"`
	// Backup method. Currently, only physical backup is supported. The value is Physical.
	BackupMethod *string `pulumi:"backupMethod"`
	// The start time of the backup.
	BackupStartTime *string `pulumi:"backupStartTime"`
	// The status of the backup.
	BackupStatus *string `pulumi:"backupStatus"`
	// The type of the backup.
	BackupType *string `pulumi:"backupType"`
	// The id of the instance.
	InstanceId string `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getBackups.
type GetBackupsResult struct {
	// The end time of the backup.
	BackupEndTime *string `pulumi:"backupEndTime"`
	// The name of the backup method.
	BackupMethod *string `pulumi:"backupMethod"`
	// The start time of the backup.
	BackupStartTime *string `pulumi:"backupStartTime"`
	// The status of the backup.
	BackupStatus *string `pulumi:"backupStatus"`
	// The type of the backup.
	BackupType *string `pulumi:"backupType"`
	// The collection of query.
	Backups []GetBackupsBackup `pulumi:"backups"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The id of the instance.
	InstanceId string  `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func GetBackupsOutput(ctx *pulumi.Context, args GetBackupsOutputArgs, opts ...pulumi.InvokeOption) GetBackupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBackupsResult, error) {
			args := v.(GetBackupsArgs)
			r, err := GetBackups(ctx, &args, opts...)
			var s GetBackupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBackupsResultOutput)
}

// A collection of arguments for invoking getBackups.
type GetBackupsOutputArgs struct {
	// The end time of the backup.
	BackupEndTime pulumi.StringPtrInput `pulumi:"backupEndTime"`
	// Backup method. Currently, only physical backup is supported. The value is Physical.
	BackupMethod pulumi.StringPtrInput `pulumi:"backupMethod"`
	// The start time of the backup.
	BackupStartTime pulumi.StringPtrInput `pulumi:"backupStartTime"`
	// The status of the backup.
	BackupStatus pulumi.StringPtrInput `pulumi:"backupStatus"`
	// The type of the backup.
	BackupType pulumi.StringPtrInput `pulumi:"backupType"`
	// The id of the instance.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetBackupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackupsArgs)(nil)).Elem()
}

// A collection of values returned by getBackups.
type GetBackupsResultOutput struct{ *pulumi.OutputState }

func (GetBackupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackupsResult)(nil)).Elem()
}

func (o GetBackupsResultOutput) ToGetBackupsResultOutput() GetBackupsResultOutput {
	return o
}

func (o GetBackupsResultOutput) ToGetBackupsResultOutputWithContext(ctx context.Context) GetBackupsResultOutput {
	return o
}

// The end time of the backup.
func (o GetBackupsResultOutput) BackupEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupsResult) *string { return v.BackupEndTime }).(pulumi.StringPtrOutput)
}

// The name of the backup method.
func (o GetBackupsResultOutput) BackupMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupsResult) *string { return v.BackupMethod }).(pulumi.StringPtrOutput)
}

// The start time of the backup.
func (o GetBackupsResultOutput) BackupStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupsResult) *string { return v.BackupStartTime }).(pulumi.StringPtrOutput)
}

// The status of the backup.
func (o GetBackupsResultOutput) BackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupsResult) *string { return v.BackupStatus }).(pulumi.StringPtrOutput)
}

// The type of the backup.
func (o GetBackupsResultOutput) BackupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupsResult) *string { return v.BackupType }).(pulumi.StringPtrOutput)
}

// The collection of query.
func (o GetBackupsResultOutput) Backups() GetBackupsBackupArrayOutput {
	return o.ApplyT(func(v GetBackupsResult) []GetBackupsBackup { return v.Backups }).(GetBackupsBackupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBackupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The id of the instance.
func (o GetBackupsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetBackupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o GetBackupsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetBackupsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBackupsResultOutput{})
}
