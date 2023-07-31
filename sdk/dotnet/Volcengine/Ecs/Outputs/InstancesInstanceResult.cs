// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Ecs.Outputs
{

    [OutputType]
    public sealed class InstancesInstanceResult
    {
        /// <summary>
        /// The number of ECS instance CPU cores.
        /// </summary>
        public readonly int Cpus;
        /// <summary>
        /// The create time of ECS instance.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The ID of DeploymentSet.
        /// </summary>
        public readonly string DeploymentSetId;
        /// <summary>
        /// The description of ECS instance.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The GPU device info of Instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesInstanceGpuDeviceResult> GpuDevices;
        /// <summary>
        /// The host name of ECS instance.
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// The image ID of ECS instance.
        /// </summary>
        public readonly string ImageId;
        /// <summary>
        /// The charge type of ECS instance.
        /// </summary>
        public readonly string InstanceChargeType;
        /// <summary>
        /// The ID of ECS instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The name of ECS instance.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// The spec type of ECS instance.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// The number of IPv6 addresses of the ECS instance.
        /// </summary>
        public readonly int Ipv6AddressCount;
        /// <summary>
        /// The  IPv6 address list of the ECS instance.
        /// </summary>
        public readonly ImmutableArray<string> Ipv6Addresses;
        /// <summary>
        /// The Flag of GPU instance.If the instance is GPU,The flag is true.
        /// </summary>
        public readonly bool IsGpu;
        /// <summary>
        /// The ssh key ID of ECS instance.
        /// </summary>
        public readonly string KeyPairId;
        /// <summary>
        /// The key pair name of ECS instance.
        /// </summary>
        public readonly string KeyPairName;
        /// <summary>
        /// The memory size of ECS instance.
        /// </summary>
        public readonly int MemorySize;
        /// <summary>
        /// The networkInterface detail collection of ECS instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesInstanceNetworkInterfaceResult> NetworkInterfaces;
        /// <summary>
        /// The os name of ECS instance.
        /// </summary>
        public readonly string OsName;
        /// <summary>
        /// The os type of ECS instance.
        /// </summary>
        public readonly string OsType;
        /// <summary>
        /// The ProjectName of ECS instance.
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// The spot strategy of ECS instance.
        /// </summary>
        public readonly string SpotStrategy;
        /// <summary>
        /// The status of ECS instance.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The stop mode of ECS instance.
        /// </summary>
        public readonly string StoppedMode;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesInstanceTagResult> Tags;
        /// <summary>
        /// The update time of ECS instance.
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// The volume detail collection of volume.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesInstanceVolumeResult> Volumes;
        /// <summary>
        /// The VPC ID of ECS instance.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// The available zone ID of ECS instance.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private InstancesInstanceResult(
            int cpus,

            string createdAt,

            string deploymentSetId,

            string description,

            ImmutableArray<Outputs.InstancesInstanceGpuDeviceResult> gpuDevices,

            string hostName,

            string imageId,

            string instanceChargeType,

            string instanceId,

            string instanceName,

            string instanceType,

            int ipv6AddressCount,

            ImmutableArray<string> ipv6Addresses,

            bool isGpu,

            string keyPairId,

            string keyPairName,

            int memorySize,

            ImmutableArray<Outputs.InstancesInstanceNetworkInterfaceResult> networkInterfaces,

            string osName,

            string osType,

            string projectName,

            string spotStrategy,

            string status,

            string stoppedMode,

            ImmutableArray<Outputs.InstancesInstanceTagResult> tags,

            string updatedAt,

            ImmutableArray<Outputs.InstancesInstanceVolumeResult> volumes,

            string vpcId,

            string zoneId)
        {
            Cpus = cpus;
            CreatedAt = createdAt;
            DeploymentSetId = deploymentSetId;
            Description = description;
            GpuDevices = gpuDevices;
            HostName = hostName;
            ImageId = imageId;
            InstanceChargeType = instanceChargeType;
            InstanceId = instanceId;
            InstanceName = instanceName;
            InstanceType = instanceType;
            Ipv6AddressCount = ipv6AddressCount;
            Ipv6Addresses = ipv6Addresses;
            IsGpu = isGpu;
            KeyPairId = keyPairId;
            KeyPairName = keyPairName;
            MemorySize = memorySize;
            NetworkInterfaces = networkInterfaces;
            OsName = osName;
            OsType = osType;
            ProjectName = projectName;
            SpotStrategy = spotStrategy;
            Status = status;
            StoppedMode = stoppedMode;
            Tags = tags;
            UpdatedAt = updatedAt;
            Volumes = volumes;
            VpcId = vpcId;
            ZoneId = zoneId;
        }
    }
}