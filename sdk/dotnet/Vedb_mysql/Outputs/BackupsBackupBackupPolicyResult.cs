// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vedb_mysql.Outputs
{

    [OutputType]
    public sealed class BackupsBackupBackupPolicyResult
    {
        /// <summary>
        /// Data backup retention period, value: 7 to 30 days.
        /// </summary>
        public readonly int BackupRetentionPeriod;
        /// <summary>
        /// The time for executing the backup task. The interval window is two hours. Format: HH:mmZ-HH:mmZ (UTC time).
        /// </summary>
        public readonly string BackupTime;
        /// <summary>
        /// Whether to enable continuous backup. The value is fixed as true.
        /// </summary>
        public readonly bool ContinueBackup;
        /// <summary>
        /// Full backup period. Multiple values are separated by English commas (,). Values:
        /// Monday: Monday.
        /// Tuesday: Tuesday.
        /// Wednesday: Wednesday.
        /// Thursday: Thursday.
        /// Friday: Friday.
        /// Saturday: Saturday.
        /// Sunday: Sunday.
        /// </summary>
        public readonly string FullBackupPeriod;
        /// <summary>
        /// The id of the instance.
        /// </summary>
        public readonly string InstanceId;

        [OutputConstructor]
        private BackupsBackupBackupPolicyResult(
            int backupRetentionPeriod,

            string backupTime,

            bool continueBackup,

            string fullBackupPeriod,

            string instanceId)
        {
            BackupRetentionPeriod = backupRetentionPeriod;
            BackupTime = backupTime;
            ContinueBackup = continueBackup;
            FullBackupPeriod = fullBackupPeriod;
            InstanceId = instanceId;
        }
    }
}