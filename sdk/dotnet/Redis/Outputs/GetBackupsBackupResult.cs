// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Redis.Outputs
{

    [OutputType]
    public sealed class GetBackupsBackupResult
    {
        /// <summary>
        /// The download address information of the backup file to which the current backup point belongs.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBackupsBackupBackupPointDownloadUrlResult> BackupPointDownloadUrls;
        /// <summary>
        /// The id of backup point.
        /// </summary>
        public readonly string BackupPointId;
        /// <summary>
        /// Backup strategy.
        /// </summary>
        public readonly string BackupStrategy;
        /// <summary>
        /// Backup type.
        /// </summary>
        public readonly string BackupType;
        /// <summary>
        /// Query end time.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// (**Deprecated**) Replaced by instance_info. Information of instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBackupsBackupInstanceDetailResult> InstanceDetails;
        /// <summary>
        /// Id of instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// Information of instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBackupsBackupInstanceInfoResult> InstanceInfos;
        /// <summary>
        /// Back up the project to which it belongs.
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// Size in MiB.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// Query start time.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Status of backup (Creating/Available/Unavailable/Deleting).
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Backup retention days.
        /// </summary>
        public readonly int Ttl;

        [OutputConstructor]
        private GetBackupsBackupResult(
            ImmutableArray<Outputs.GetBackupsBackupBackupPointDownloadUrlResult> backupPointDownloadUrls,

            string backupPointId,

            string backupStrategy,

            string backupType,

            string endTime,

            ImmutableArray<Outputs.GetBackupsBackupInstanceDetailResult> instanceDetails,

            string instanceId,

            ImmutableArray<Outputs.GetBackupsBackupInstanceInfoResult> instanceInfos,

            string projectName,

            int size,

            string startTime,

            string status,

            int ttl)
        {
            BackupPointDownloadUrls = backupPointDownloadUrls;
            BackupPointId = backupPointId;
            BackupStrategy = backupStrategy;
            BackupType = backupType;
            EndTime = endTime;
            InstanceDetails = instanceDetails;
            InstanceId = instanceId;
            InstanceInfos = instanceInfos;
            ProjectName = projectName;
            Size = size;
            StartTime = startTime;
            Status = status;
            Ttl = ttl;
        }
    }
}
