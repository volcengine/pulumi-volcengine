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
    public sealed class BackupBackupPolicy
    {
        /// <summary>
        /// Data backup retention period, value: 7 to 30 days.
        /// </summary>
        public readonly int BackupRetentionPeriod;
        /// <summary>
        /// The time for executing the backup task has an interval window of 2 hours and must be an even-hour time. Format: HH:mmZ-HH:mmZ (UTC time).
        /// </summary>
        public readonly string BackupTime;
        /// <summary>
        /// Full backup period. It is recommended to select at least 2 days per week for full backup. Multiple values are separated by English commas (,). Values: Monday: Monday. Tuesday: Tuesday. Wednesday: Wednesday. Thursday: Thursday. Friday: Friday. Saturday: Saturday. Sunday: Sunday.
        /// </summary>
        public readonly string FullBackupPeriod;

        [OutputConstructor]
        private BackupBackupPolicy(
            int backupRetentionPeriod,

            string backupTime,

            string fullBackupPeriod)
        {
            BackupRetentionPeriod = backupRetentionPeriod;
            BackupTime = backupTime;
            FullBackupPeriod = fullBackupPeriod;
        }
    }
}