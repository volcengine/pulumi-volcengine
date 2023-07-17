// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Veenedge.Outputs
{

    [OutputType]
    public sealed class CloudServersCloudServerResult
    {
        /// <summary>
        /// The config of billing.
        /// </summary>
        public readonly Outputs.CloudServersCloudServerBillingConfigResult BillingConfig;
        /// <summary>
        /// The Id of cloud server.
        /// </summary>
        public readonly string CloudServerIdentity;
        /// <summary>
        /// The cpu info of cloud server.
        /// </summary>
        public readonly string Cpu;
        /// <summary>
        /// The create time info.
        /// </summary>
        public readonly int CreateTime;
        /// <summary>
        /// The config of custom data.
        /// </summary>
        public readonly Outputs.CloudServersCloudServerCustomDataResult CustomData;
        /// <summary>
        /// The config of gpu.
        /// </summary>
        public readonly Outputs.CloudServersCloudServerGpuResult Gpu;
        /// <summary>
        /// The Id of cloud server.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The config of image.
        /// </summary>
        public readonly Outputs.CloudServersCloudServerImageResult Image;
        /// <summary>
        /// The count of instance.
        /// </summary>
        public readonly int InstanceCount;
        /// <summary>
        /// The status of instances.
        /// </summary>
        public readonly ImmutableArray<Outputs.CloudServersCloudServerInstanceStatusResult> InstanceStatuses;
        /// <summary>
        /// The memory info of cloud server.
        /// </summary>
        public readonly string Mem;
        /// <summary>
        /// The name of cloud server.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The config of network.
        /// </summary>
        public readonly Outputs.CloudServersCloudServerNetworkResult Network;
        /// <summary>
        /// The config of schedule strategy.
        /// </summary>
        public readonly Outputs.CloudServersCloudServerScheduleStrategyConfigsResult ScheduleStrategyConfigs;
        /// <summary>
        /// The config of secret.
        /// </summary>
        public readonly Outputs.CloudServersCloudServerSecretConfigResult SecretConfig;
        /// <summary>
        /// The server area count number.
        /// </summary>
        public readonly int ServerAreaCount;
        /// <summary>
        /// The area level of cloud server.
        /// </summary>
        public readonly string ServerAreaLevel;
        /// <summary>
        /// The server areas info.
        /// </summary>
        public readonly ImmutableArray<Outputs.CloudServersCloudServerServerAreaResult> ServerAreas;
        /// <summary>
        /// The spec info of cloud server.
        /// </summary>
        public readonly string Spec;
        /// <summary>
        /// The Chinese spec info of cloud server.
        /// </summary>
        public readonly string SpecDisplay;
        /// <summary>
        /// The spec summary of cloud server.
        /// </summary>
        public readonly ImmutableDictionary<string, object> SpecSum;
        /// <summary>
        /// The config of storage.
        /// </summary>
        public readonly Outputs.CloudServersCloudServerStorageResult Storage;
        /// <summary>
        /// The update time info.
        /// </summary>
        public readonly int UpdateTime;

        [OutputConstructor]
        private CloudServersCloudServerResult(
            Outputs.CloudServersCloudServerBillingConfigResult billingConfig,

            string cloudServerIdentity,

            string cpu,

            int createTime,

            Outputs.CloudServersCloudServerCustomDataResult customData,

            Outputs.CloudServersCloudServerGpuResult gpu,

            string id,

            Outputs.CloudServersCloudServerImageResult image,

            int instanceCount,

            ImmutableArray<Outputs.CloudServersCloudServerInstanceStatusResult> instanceStatuses,

            string mem,

            string name,

            Outputs.CloudServersCloudServerNetworkResult network,

            Outputs.CloudServersCloudServerScheduleStrategyConfigsResult scheduleStrategyConfigs,

            Outputs.CloudServersCloudServerSecretConfigResult secretConfig,

            int serverAreaCount,

            string serverAreaLevel,

            ImmutableArray<Outputs.CloudServersCloudServerServerAreaResult> serverAreas,

            string spec,

            string specDisplay,

            ImmutableDictionary<string, object> specSum,

            Outputs.CloudServersCloudServerStorageResult storage,

            int updateTime)
        {
            BillingConfig = billingConfig;
            CloudServerIdentity = cloudServerIdentity;
            Cpu = cpu;
            CreateTime = createTime;
            CustomData = customData;
            Gpu = gpu;
            Id = id;
            Image = image;
            InstanceCount = instanceCount;
            InstanceStatuses = instanceStatuses;
            Mem = mem;
            Name = name;
            Network = network;
            ScheduleStrategyConfigs = scheduleStrategyConfigs;
            SecretConfig = secretConfig;
            ServerAreaCount = serverAreaCount;
            ServerAreaLevel = serverAreaLevel;
            ServerAreas = serverAreas;
            Spec = spec;
            SpecDisplay = specDisplay;
            SpecSum = specSum;
            Storage = storage;
            UpdateTime = updateTime;
        }
    }
}