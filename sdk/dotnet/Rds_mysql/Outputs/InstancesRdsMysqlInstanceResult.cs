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
    public sealed class InstancesRdsMysqlInstanceResult
    {
        /// <summary>
        /// The version of allow list.
        /// </summary>
        public readonly string AllowListVersion;
        /// <summary>
        /// The instance has used backup space. Unit: GB.
        /// </summary>
        public readonly int BackupUse;
        /// <summary>
        /// Payment methods.
        /// </summary>
        public readonly Outputs.InstancesRdsMysqlInstanceChargeDetailResult ChargeDetail;
        /// <summary>
        /// Node creation local time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Data synchronization mode.
        /// </summary>
        public readonly string DataSyncMode;
        /// <summary>
        /// The version of the RDS instance.
        /// </summary>
        public readonly string DbEngineVersion;
        /// <summary>
        /// The endpoint info of the RDS instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesRdsMysqlInstanceEndpointResult> Endpoints;
        /// <summary>
        /// The ID of the RDS instance.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The id of the RDS instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The name of the RDS instance.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// The status of the RDS instance.
        /// </summary>
        public readonly string InstanceStatus;
        /// <summary>
        /// Whether the table name is case sensitive, the default value is 1.
        /// Ranges:
        /// 0: Table names are stored as fixed and table names are case-sensitive.
        /// 1: Table names will be stored in lowercase and table names are not case sensitive.
        /// </summary>
        public readonly string LowerCaseTableNames;
        /// <summary>
        /// Maintenance Window.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesRdsMysqlInstanceMaintenanceWindowResult> MaintenanceWindows;
        /// <summary>
        /// Memory size in GB.
        /// </summary>
        public readonly int Memory;
        /// <summary>
        /// The number of nodes.
        /// </summary>
        public readonly int NodeNumber;
        /// <summary>
        /// General instance type, different from Custom instance type.
        /// </summary>
        public readonly string NodeSpec;
        /// <summary>
        /// Instance node information.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesRdsMysqlInstanceNodeResult> Nodes;
        /// <summary>
        /// The project name of the RDS instance.
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// The region of the RDS instance.
        /// </summary>
        public readonly string RegionId;
        /// <summary>
        /// Total instance storage space. Unit: GB.
        /// </summary>
        public readonly int StorageSpace;
        /// <summary>
        /// Instance storage type.
        /// </summary>
        public readonly string StorageType;
        /// <summary>
        /// The instance has used storage space. Unit: GB.
        /// </summary>
        public readonly int StorageUse;
        /// <summary>
        /// The subnet ID of the RDS instance.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesRdsMysqlInstanceTagResult> Tags;
        /// <summary>
        /// Time zone.
        /// </summary>
        public readonly string TimeZone;
        /// <summary>
        /// The update time of the RDS instance.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// CPU size.
        /// </summary>
        public readonly int VCpu;
        /// <summary>
        /// The vpc ID of the RDS instance.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// The available zone of the RDS instance.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private InstancesRdsMysqlInstanceResult(
            string allowListVersion,

            int backupUse,

            Outputs.InstancesRdsMysqlInstanceChargeDetailResult chargeDetail,

            string createTime,

            string dataSyncMode,

            string dbEngineVersion,

            ImmutableArray<Outputs.InstancesRdsMysqlInstanceEndpointResult> endpoints,

            string id,

            string instanceId,

            string instanceName,

            string instanceStatus,

            string lowerCaseTableNames,

            ImmutableArray<Outputs.InstancesRdsMysqlInstanceMaintenanceWindowResult> maintenanceWindows,

            int memory,

            int nodeNumber,

            string nodeSpec,

            ImmutableArray<Outputs.InstancesRdsMysqlInstanceNodeResult> nodes,

            string projectName,

            string regionId,

            int storageSpace,

            string storageType,

            int storageUse,

            string subnetId,

            ImmutableArray<Outputs.InstancesRdsMysqlInstanceTagResult> tags,

            string timeZone,

            string updateTime,

            int vCpu,

            string vpcId,

            string zoneId)
        {
            AllowListVersion = allowListVersion;
            BackupUse = backupUse;
            ChargeDetail = chargeDetail;
            CreateTime = createTime;
            DataSyncMode = dataSyncMode;
            DbEngineVersion = dbEngineVersion;
            Endpoints = endpoints;
            Id = id;
            InstanceId = instanceId;
            InstanceName = instanceName;
            InstanceStatus = instanceStatus;
            LowerCaseTableNames = lowerCaseTableNames;
            MaintenanceWindows = maintenanceWindows;
            Memory = memory;
            NodeNumber = nodeNumber;
            NodeSpec = nodeSpec;
            Nodes = nodes;
            ProjectName = projectName;
            RegionId = regionId;
            StorageSpace = storageSpace;
            StorageType = storageType;
            StorageUse = storageUse;
            SubnetId = subnetId;
            Tags = tags;
            TimeZone = timeZone;
            UpdateTime = updateTime;
            VCpu = vCpu;
            VpcId = vpcId;
            ZoneId = zoneId;
        }
    }
}