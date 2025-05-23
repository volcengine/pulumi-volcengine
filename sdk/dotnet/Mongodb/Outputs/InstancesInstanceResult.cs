// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Mongodb.Outputs
{

    [OutputType]
    public sealed class InstancesInstanceResult
    {
        /// <summary>
        /// Whether to enable automatic renewal.
        /// </summary>
        public readonly bool AutoRenew;
        /// <summary>
        /// The charge status.
        /// </summary>
        public readonly string ChargeStatus;
        /// <summary>
        /// The charge type of instance.
        /// </summary>
        public readonly string ChargeType;
        /// <summary>
        /// The planned close time.
        /// </summary>
        public readonly string ClosedTime;
        /// <summary>
        /// The list of config servers.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesInstanceConfigServerResult> ConfigServers;
        /// <summary>
        /// The ID of config servers.
        /// </summary>
        public readonly string ConfigServersId;
        /// <summary>
        /// The creation time of instance.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The db engine to query, valid value contains `MongoDB`.
        /// </summary>
        public readonly string DbEngine;
        /// <summary>
        /// The version of db engine to query, valid value contains `MongoDB_4_0`.
        /// </summary>
        public readonly string DbEngineVersion;
        /// <summary>
        /// The version string of database engine.
        /// </summary>
        public readonly string DbEngineVersionStr;
        /// <summary>
        /// The expired time of instance.
        /// </summary>
        public readonly string ExpiredTime;
        /// <summary>
        /// The instance ID to query.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The instance name to query.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// The instance status to query.
        /// </summary>
        public readonly string InstanceStatus;
        /// <summary>
        /// The type of instance to query, the valid value contains `ReplicaSet` or `ShardedCluster`.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// The list of mongos.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesInstanceMongoResult> Mongos;
        /// <summary>
        /// The ID of mongos.
        /// </summary>
        public readonly string MongosId;
        /// <summary>
        /// The node information.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesInstanceNodeResult> Nodes;
        /// <summary>
        /// The private endpoint address of instance.
        /// </summary>
        public readonly string PrivateEndpoint;
        /// <summary>
        /// The project name to query.
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// The number of readonly node in instance.
        /// </summary>
        public readonly int ReadOnlyNodeNumber;
        /// <summary>
        /// The planned reclaim time of instance.
        /// </summary>
        public readonly string ReclaimTime;
        /// <summary>
        /// The list of shards.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesInstanceShardResult> Shards;
        /// <summary>
        /// Whether ssl enabled.
        /// </summary>
        public readonly bool SslEnable;
        /// <summary>
        /// The ssl expire time.
        /// </summary>
        public readonly string SslExpireTime;
        /// <summary>
        /// Whether ssl is valid.
        /// </summary>
        public readonly bool SslIsValid;
        /// <summary>
        /// The storage type of instance.
        /// </summary>
        public readonly string StorageType;
        /// <summary>
        /// The subnet id of instance.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesInstanceTagResult> Tags;
        /// <summary>
        /// The update time of instance.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// The vpc id of instance to query.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// The zone ID to query.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private InstancesInstanceResult(
            bool autoRenew,

            string chargeStatus,

            string chargeType,

            string closedTime,

            ImmutableArray<Outputs.InstancesInstanceConfigServerResult> configServers,

            string configServersId,

            string createTime,

            string dbEngine,

            string dbEngineVersion,

            string dbEngineVersionStr,

            string expiredTime,

            string instanceId,

            string instanceName,

            string instanceStatus,

            string instanceType,

            ImmutableArray<Outputs.InstancesInstanceMongoResult> mongos,

            string mongosId,

            ImmutableArray<Outputs.InstancesInstanceNodeResult> nodes,

            string privateEndpoint,

            string projectName,

            int readOnlyNodeNumber,

            string reclaimTime,

            ImmutableArray<Outputs.InstancesInstanceShardResult> shards,

            bool sslEnable,

            string sslExpireTime,

            bool sslIsValid,

            string storageType,

            string subnetId,

            ImmutableArray<Outputs.InstancesInstanceTagResult> tags,

            string updateTime,

            string vpcId,

            string zoneId)
        {
            AutoRenew = autoRenew;
            ChargeStatus = chargeStatus;
            ChargeType = chargeType;
            ClosedTime = closedTime;
            ConfigServers = configServers;
            ConfigServersId = configServersId;
            CreateTime = createTime;
            DbEngine = dbEngine;
            DbEngineVersion = dbEngineVersion;
            DbEngineVersionStr = dbEngineVersionStr;
            ExpiredTime = expiredTime;
            InstanceId = instanceId;
            InstanceName = instanceName;
            InstanceStatus = instanceStatus;
            InstanceType = instanceType;
            Mongos = mongos;
            MongosId = mongosId;
            Nodes = nodes;
            PrivateEndpoint = privateEndpoint;
            ProjectName = projectName;
            ReadOnlyNodeNumber = readOnlyNodeNumber;
            ReclaimTime = reclaimTime;
            Shards = shards;
            SslEnable = sslEnable;
            SslExpireTime = sslExpireTime;
            SslIsValid = sslIsValid;
            StorageType = storageType;
            SubnetId = subnetId;
            Tags = tags;
            UpdateTime = updateTime;
            VpcId = vpcId;
            ZoneId = zoneId;
        }
    }
}
