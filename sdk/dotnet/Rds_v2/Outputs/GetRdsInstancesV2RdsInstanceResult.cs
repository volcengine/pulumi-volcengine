// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_v2.Outputs
{

    [OutputType]
    public sealed class GetRdsInstancesV2RdsInstanceResult
    {
        /// <summary>
        /// Payment methods.
        /// </summary>
        public readonly Outputs.GetRdsInstancesV2RdsInstanceChargeDetailResult ChargeDetail;
        /// <summary>
        /// The connection info ot the RDS instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRdsInstancesV2RdsInstanceConnectionInfoResult> ConnectionInfos;
        /// <summary>
        /// Node creation local time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The engine of the RDS instance.
        /// </summary>
        public readonly string DbEngine;
        /// <summary>
        /// The version of the RDS instance, Value:
        /// MySQL Community:
        /// MySQL_5.7 - MySQL 5.7
        /// MySQL_8_0 - MySQL 8.0
        /// PostgreSQL Community:
        /// PostgreSQL_11 - PostgreSQL 11
        /// PostgreSQL_12 - PostgreSQL 12
        /// Microsoft SQL Server: Not available at this time
        /// SQLServer_2019 - SQL Server 2019
        /// veDB for MySQL:
        /// MySQL_8_0 - MySQL 8.0
        /// veDB for PostgreSQL:
        /// PostgreSQL_13 - PostgreSQL 13.
        /// </summary>
        public readonly string DbEngineVersion;
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
        /// The status of the RDS instance, Value:
        /// Running - running
        /// Creating - Creating
        /// Deleting - Deleting
        /// Restarting - Restarting
        /// Restoring - Restoring
        /// Updating - changing
        /// Upgrading - Upgrading
        /// Error - the error.
        /// </summary>
        public readonly string InstanceStatus;
        /// <summary>
        /// The type of the RDS instance, Value:
        /// Value:
        /// RDS for MySQL:
        /// HA - high availability version;
        /// RDS for PostgreSQL:
        /// HA - high availability version;
        /// Microsoft SQL Server: Not available at this time
        /// Enterprise - Enterprise Edition
        /// Standard - Standard Edition
        /// Web - Web version
        /// veDB for MySQL:
        /// Cluster - Cluster Edition
        /// veDB for PostgreSQL:
        /// Cluster - Cluster Edition
        /// MySQL Sharding:
        /// HA - high availability version;.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// Instance node information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRdsInstancesV2RdsInstanceNodeDetailInfoResult> NodeDetailInfos;
        /// <summary>
        /// The number of nodes.
        /// </summary>
        public readonly int NodeNumber;
        /// <summary>
        /// General instance type, different from Custom instance type.
        /// </summary>
        public readonly string NodeSpec;
        /// <summary>
        /// Instance intranet port.
        /// </summary>
        public readonly string Port;
        /// <summary>
        /// Subordinate to the project.
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// The region of the RDS instance.
        /// </summary>
        public readonly string RegionId;
        /// <summary>
        /// The number of shards.
        /// </summary>
        public readonly int ShardNumber;
        /// <summary>
        /// Total instance storage space. Unit: GB.
        /// </summary>
        public readonly int StorageSpace;
        /// <summary>
        /// Instance storage type. When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, the value is:
        /// LocalSSD - local SSD disk
        /// When the database type is veDB_MySQL/veDB_PostgreSQL, the value is:
        /// DistributedStorage - Distributed Storage.
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
        /// Time zone.
        /// </summary>
        public readonly string TimeZone;
        /// <summary>
        /// The vpc ID of the RDS instance.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// The available zone of the RDS instance.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private GetRdsInstancesV2RdsInstanceResult(
            Outputs.GetRdsInstancesV2RdsInstanceChargeDetailResult chargeDetail,

            ImmutableArray<Outputs.GetRdsInstancesV2RdsInstanceConnectionInfoResult> connectionInfos,

            string createTime,

            string dbEngine,

            string dbEngineVersion,

            string id,

            string instanceId,

            string instanceName,

            string instanceStatus,

            string instanceType,

            ImmutableArray<Outputs.GetRdsInstancesV2RdsInstanceNodeDetailInfoResult> nodeDetailInfos,

            int nodeNumber,

            string nodeSpec,

            string port,

            string projectName,

            string regionId,

            int shardNumber,

            int storageSpace,

            string storageType,

            int storageUse,

            string subnetId,

            string timeZone,

            string vpcId,

            string zoneId)
        {
            ChargeDetail = chargeDetail;
            ConnectionInfos = connectionInfos;
            CreateTime = createTime;
            DbEngine = dbEngine;
            DbEngineVersion = dbEngineVersion;
            Id = id;
            InstanceId = instanceId;
            InstanceName = instanceName;
            InstanceStatus = instanceStatus;
            InstanceType = instanceType;
            NodeDetailInfos = nodeDetailInfos;
            NodeNumber = nodeNumber;
            NodeSpec = nodeSpec;
            Port = port;
            ProjectName = projectName;
            RegionId = regionId;
            ShardNumber = shardNumber;
            StorageSpace = storageSpace;
            StorageType = storageType;
            StorageUse = storageUse;
            SubnetId = subnetId;
            TimeZone = timeZone;
            VpcId = vpcId;
            ZoneId = zoneId;
        }
    }
}
