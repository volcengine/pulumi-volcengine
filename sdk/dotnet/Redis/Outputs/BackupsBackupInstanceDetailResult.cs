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
    public sealed class BackupsBackupInstanceDetailResult
    {
        /// <summary>
        /// Id of account.
        /// </summary>
        public readonly int AccountId;
        /// <summary>
        /// Arch type of instance(Standard/Cluster).
        /// </summary>
        public readonly string ArchType;
        /// <summary>
        /// Charge type of instance(Postpaid/Prepaid).
        /// </summary>
        public readonly string ChargeType;
        /// <summary>
        /// Engine version of instance.
        /// </summary>
        public readonly string EngineVersion;
        /// <summary>
        /// Expired time of instance.
        /// </summary>
        public readonly string ExpiredTime;
        /// <summary>
        /// Id of instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// Name of instance.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// The maintainable period (in UTC) of the instance.
        /// </summary>
        public readonly string MaintenanceTime;
        /// <summary>
        /// Network type of instance.
        /// </summary>
        public readonly string NetworkType;
        /// <summary>
        /// Project name of instance.
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// Id of region.
        /// </summary>
        public readonly string RegionId;
        /// <summary>
        /// Count of replica in which shard.
        /// </summary>
        public readonly int Replicas;
        /// <summary>
        /// Count of cpu cores of instance.
        /// </summary>
        public readonly int ServerCpu;
        /// <summary>
        /// Capacity of shard.
        /// </summary>
        public readonly int ShardCapacity;
        /// <summary>
        /// Count of shard.
        /// </summary>
        public readonly int ShardCount;
        /// <summary>
        /// Total capacity of instance.
        /// </summary>
        public readonly int TotalCapacity;
        /// <summary>
        /// Capacity used of this instance.
        /// </summary>
        public readonly int UsedCapacity;
        /// <summary>
        /// Information of vpc.
        /// </summary>
        public readonly ImmutableArray<Outputs.BackupsBackupInstanceDetailVpcInfoResult> VpcInfos;
        /// <summary>
        /// List of id of zone.
        /// </summary>
        public readonly ImmutableArray<string> ZoneIds;

        [OutputConstructor]
        private BackupsBackupInstanceDetailResult(
            int accountId,

            string archType,

            string chargeType,

            string engineVersion,

            string expiredTime,

            string instanceId,

            string instanceName,

            string maintenanceTime,

            string networkType,

            string projectName,

            string regionId,

            int replicas,

            int serverCpu,

            int shardCapacity,

            int shardCount,

            int totalCapacity,

            int usedCapacity,

            ImmutableArray<Outputs.BackupsBackupInstanceDetailVpcInfoResult> vpcInfos,

            ImmutableArray<string> zoneIds)
        {
            AccountId = accountId;
            ArchType = archType;
            ChargeType = chargeType;
            EngineVersion = engineVersion;
            ExpiredTime = expiredTime;
            InstanceId = instanceId;
            InstanceName = instanceName;
            MaintenanceTime = maintenanceTime;
            NetworkType = networkType;
            ProjectName = projectName;
            RegionId = regionId;
            Replicas = replicas;
            ServerCpu = serverCpu;
            ShardCapacity = shardCapacity;
            ShardCount = shardCount;
            TotalCapacity = totalCapacity;
            UsedCapacity = usedCapacity;
            VpcInfos = vpcInfos;
            ZoneIds = zoneIds;
        }
    }
}