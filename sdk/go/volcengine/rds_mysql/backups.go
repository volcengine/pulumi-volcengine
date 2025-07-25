// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds_mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of rds mysql backups
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/rds_mysql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rds_mysql.GetBackups(ctx, &rds_mysql.GetBackupsArgs{
//				InstanceId: pulumi.StringRef("mysql-b51d37110dd1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: volcengine.rds_mysql.Backups has been deprecated in favor of volcengine.rds_mysql.getBackups
func Backups(ctx *pulumi.Context, args *BackupsArgs, opts ...pulumi.InvokeOption) (*BackupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv BackupsResult
	err := ctx.Invoke("volcengine:rds_mysql/backups:Backups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Backups.
type BackupsArgs struct {
	// The end time of the backup.
	BackupEndTime *string `pulumi:"backupEndTime"`
	// The id of the backup.
	BackupId *string `pulumi:"backupId"`
	// Backup type, value: Physical: Physical backup. Default value. Logical: Logical backup. Description: There is no default value. When this field is not passed, backups of all states under the query conditions limited by other fields are returned.
	BackupMethod *string `pulumi:"backupMethod"`
	// The start time of the backup.
	BackupStartTime *string `pulumi:"backupStartTime"`
	// Backup status, values: Success: Success. Failed: Failed. Running: In progress. Description: There is no default value. When this field is not passed, all backups in all states under the query conditions limited by other fields are returned.
	BackupStatus *string `pulumi:"backupStatus"`
	// Backup method, value: Full: Full backup under physical backup type or library table backup under logical backup type. Increment: Incremental backup under physical backup type. DumpAll: Full database backup under logical backup type. Description: There is no default value. When this field is not passed, all backups of all methods under the query conditions limited by other fields are returned.
	BackupType *string `pulumi:"backupType"`
	// Creator of backup. Values: System: System. User: User. Description: There is no default value. When this field is not passed, all types of backups under the query conditions limited by other fields are returned.
	CreateType *string `pulumi:"createType"`
	// The id of the instance.
	InstanceId *string `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by Backups.
type BackupsResult struct {
	// The end time of backup, in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
	BackupEndTime *string `pulumi:"backupEndTime"`
	// The id of the backup.
	BackupId *string `pulumi:"backupId"`
	// Backup type, value: Physical: Physical backup. Logical: Logical backup.
	BackupMethod *string `pulumi:"backupMethod"`
	// The start time of backup, in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
	BackupStartTime *string `pulumi:"backupStartTime"`
	// Backup status, values: Success. Failed. Running.
	BackupStatus *string `pulumi:"backupStatus"`
	// Backup method, values:
	// Full: Full backup under physical backup type or library table backup under logical backup type.
	// Increment: Incremental backup under physical backup type (created by the system).
	// DumpAll: Full database backup under logical backup type.
	// Description:
	// There is no default value. When this field is not passed, all types of backups under the query conditions limited by other fields are returned.
	BackupType *string `pulumi:"backupType"`
	// The collection of query.
	Backups []BackupsBackup `pulumi:"backups"`
	// Creator of backup. Values: System. User.
	CreateType *string `pulumi:"createType"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstanceId *string `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func BackupsOutput(ctx *pulumi.Context, args BackupsOutputArgs, opts ...pulumi.InvokeOption) BackupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (BackupsResult, error) {
			args := v.(BackupsArgs)
			r, err := Backups(ctx, &args, opts...)
			var s BackupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(BackupsResultOutput)
}

// A collection of arguments for invoking Backups.
type BackupsOutputArgs struct {
	// The end time of the backup.
	BackupEndTime pulumi.StringPtrInput `pulumi:"backupEndTime"`
	// The id of the backup.
	BackupId pulumi.StringPtrInput `pulumi:"backupId"`
	// Backup type, value: Physical: Physical backup. Default value. Logical: Logical backup. Description: There is no default value. When this field is not passed, backups of all states under the query conditions limited by other fields are returned.
	BackupMethod pulumi.StringPtrInput `pulumi:"backupMethod"`
	// The start time of the backup.
	BackupStartTime pulumi.StringPtrInput `pulumi:"backupStartTime"`
	// Backup status, values: Success: Success. Failed: Failed. Running: In progress. Description: There is no default value. When this field is not passed, all backups in all states under the query conditions limited by other fields are returned.
	BackupStatus pulumi.StringPtrInput `pulumi:"backupStatus"`
	// Backup method, value: Full: Full backup under physical backup type or library table backup under logical backup type. Increment: Incremental backup under physical backup type. DumpAll: Full database backup under logical backup type. Description: There is no default value. When this field is not passed, all backups of all methods under the query conditions limited by other fields are returned.
	BackupType pulumi.StringPtrInput `pulumi:"backupType"`
	// Creator of backup. Values: System: System. User: User. Description: There is no default value. When this field is not passed, all types of backups under the query conditions limited by other fields are returned.
	CreateType pulumi.StringPtrInput `pulumi:"createType"`
	// The id of the instance.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (BackupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupsArgs)(nil)).Elem()
}

// A collection of values returned by Backups.
type BackupsResultOutput struct{ *pulumi.OutputState }

func (BackupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupsResult)(nil)).Elem()
}

func (o BackupsResultOutput) ToBackupsResultOutput() BackupsResultOutput {
	return o
}

func (o BackupsResultOutput) ToBackupsResultOutputWithContext(ctx context.Context) BackupsResultOutput {
	return o
}

// The end time of backup, in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
func (o BackupsResultOutput) BackupEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupsResult) *string { return v.BackupEndTime }).(pulumi.StringPtrOutput)
}

// The id of the backup.
func (o BackupsResultOutput) BackupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupsResult) *string { return v.BackupId }).(pulumi.StringPtrOutput)
}

// Backup type, value: Physical: Physical backup. Logical: Logical backup.
func (o BackupsResultOutput) BackupMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupsResult) *string { return v.BackupMethod }).(pulumi.StringPtrOutput)
}

// The start time of backup, in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
func (o BackupsResultOutput) BackupStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupsResult) *string { return v.BackupStartTime }).(pulumi.StringPtrOutput)
}

// Backup status, values: Success. Failed. Running.
func (o BackupsResultOutput) BackupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupsResult) *string { return v.BackupStatus }).(pulumi.StringPtrOutput)
}

// Backup method, values:
// Full: Full backup under physical backup type or library table backup under logical backup type.
// Increment: Incremental backup under physical backup type (created by the system).
// DumpAll: Full database backup under logical backup type.
// Description:
// There is no default value. When this field is not passed, all types of backups under the query conditions limited by other fields are returned.
func (o BackupsResultOutput) BackupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupsResult) *string { return v.BackupType }).(pulumi.StringPtrOutput)
}

// The collection of query.
func (o BackupsResultOutput) Backups() BackupsBackupArrayOutput {
	return o.ApplyT(func(v BackupsResult) []BackupsBackup { return v.Backups }).(BackupsBackupArrayOutput)
}

// Creator of backup. Values: System. User.
func (o BackupsResultOutput) CreateType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupsResult) *string { return v.CreateType }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o BackupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v BackupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o BackupsResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupsResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o BackupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o BackupsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v BackupsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupsResultOutput{})
}
