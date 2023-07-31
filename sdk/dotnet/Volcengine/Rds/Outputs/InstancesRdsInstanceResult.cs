// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Rds.Outputs
{

    [OutputType]
    public sealed class InstancesRdsInstanceResult
    {
        /// <summary>
        /// The charge status of the RDS instance.
        /// </summary>
        public readonly string ChargeStatus;
        /// <summary>
        /// The charge type of the RDS instance.
        /// </summary>
        public readonly string ChargeType;
        /// <summary>
        /// The connection info ot the RDS instance.
        /// </summary>
        public readonly Outputs.InstancesRdsInstanceConnectionInfoResult ConnectionInfo;
        /// <summary>
        /// The create time of the RDS instance.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The engine of the RDS instance.
        /// </summary>
        public readonly string DbEngine;
        /// <summary>
        /// The engine version of the RDS instance.
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
        /// The spec type detail of RDS instance.
        /// </summary>
        public readonly Outputs.InstancesRdsInstanceInstanceSpecResult InstanceSpec;
        /// <summary>
        /// The status of the RDS instance.
        /// </summary>
        public readonly string InstanceStatus;
        /// <summary>
        /// The type of the RDS instance.
        /// </summary>
        public readonly string InstanceType;
        public readonly ImmutableArray<string> ReadOnlyInstanceIds;
        /// <summary>
        /// The region of the RDS instance.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The total storage GB of the RDS instance.
        /// </summary>
        public readonly int StorageSpaceGb;
        /// <summary>
        /// The update time of the RDS instance.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// The vpc ID of the RDS instance.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// The available zone of the RDS instance.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private InstancesRdsInstanceResult(
            string chargeStatus,

            string chargeType,

            Outputs.InstancesRdsInstanceConnectionInfoResult connectionInfo,

            string createTime,

            string dbEngine,

            string dbEngineVersion,

            string id,

            string instanceId,

            string instanceName,

            Outputs.InstancesRdsInstanceInstanceSpecResult instanceSpec,

            string instanceStatus,

            string instanceType,

            ImmutableArray<string> readOnlyInstanceIds,

            string region,

            int storageSpaceGb,

            string updateTime,

            string vpcId,

            string zone)
        {
            ChargeStatus = chargeStatus;
            ChargeType = chargeType;
            ConnectionInfo = connectionInfo;
            CreateTime = createTime;
            DbEngine = dbEngine;
            DbEngineVersion = dbEngineVersion;
            Id = id;
            InstanceId = instanceId;
            InstanceName = instanceName;
            InstanceSpec = instanceSpec;
            InstanceStatus = instanceStatus;
            InstanceType = instanceType;
            ReadOnlyInstanceIds = readOnlyInstanceIds;
            Region = region;
            StorageSpaceGb = storageSpaceGb;
            UpdateTime = updateTime;
            VpcId = vpcId;
            Zone = zone;
        }
    }
}