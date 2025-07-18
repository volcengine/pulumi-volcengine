// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_mysql.Outputs
{

    [OutputType]
    public sealed class BackupPolicyCrossBackupPolicy
    {
        /// <summary>
        /// Whether to enable cross-region backup.
        /// true: Enable.
        /// false: Disable. Default value.
        /// </summary>
        public readonly bool? BackupEnabled;
        /// <summary>
        /// The destination region ID for cross-region backup. When the value of BackupEnabled is true, this parameter is required.
        /// </summary>
        public readonly string? CrossBackupRegion;
        /// <summary>
        /// Whether to enable cross-region log backup. true: Enable. false: Disable. Default value. Description: Cross-region log backup can only be enabled when cross-region backup is enabled.
        /// </summary>
        public readonly bool? LogBackupEnabled;
        /// <summary>
        /// The number of days to retain cross - region backups, with a value range of 7 to 1825 days.
        /// </summary>
        public readonly int? Retention;

        [OutputConstructor]
        private BackupPolicyCrossBackupPolicy(
            bool? backupEnabled,

            string? crossBackupRegion,

            bool? logBackupEnabled,

            int? retention)
        {
            BackupEnabled = backupEnabled;
            CrossBackupRegion = crossBackupRegion;
            LogBackupEnabled = logBackupEnabled;
            Retention = retention;
        }
    }
}
