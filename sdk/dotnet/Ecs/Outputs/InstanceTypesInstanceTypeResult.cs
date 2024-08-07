// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Ecs.Outputs
{

    [OutputType]
    public sealed class InstanceTypesInstanceTypeResult
    {
        /// <summary>
        /// The CPU benchmark performance that can be provided steadily by on-demand instances is determined by the instance type.
        /// </summary>
        public readonly int BaselineCredit;
        /// <summary>
        /// The GPU device info of Instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceTypesInstanceTypeGpusResult> Gpus;
        /// <summary>
        /// The CPU credits obtained at once when creating a on-demand performance instance are fixed at 30 credits per vCPU.
        /// </summary>
        public readonly int InitialCredit;
        /// <summary>
        /// The instance type family.
        /// </summary>
        public readonly string InstanceTypeFamily;
        /// <summary>
        /// The id of the instance type.
        /// </summary>
        public readonly string InstanceTypeId;
        /// <summary>
        /// Local disk configuration information corresponding to instance specifications.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceTypesInstanceTypeLocalVolumeResult> LocalVolumes;
        /// <summary>
        /// Memory information of instance specifications.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceTypesInstanceTypeMemoryResult> Memories;
        /// <summary>
        /// Network information of instance specifications.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceTypesInstanceTypeNetworkResult> Networks;
        /// <summary>
        /// CPU information of instance specifications.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceTypesInstanceTypeProcessorResult> Processors;
        /// <summary>
        /// RDMA Specification Information.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceTypesInstanceTypeRdmaResult> Rdmas;
        /// <summary>
        /// Cloud disk information for instance specifications.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceTypesInstanceTypeVolumeResult> Volumes;

        [OutputConstructor]
        private InstanceTypesInstanceTypeResult(
            int baselineCredit,

            ImmutableArray<Outputs.InstanceTypesInstanceTypeGpusResult> gpus,

            int initialCredit,

            string instanceTypeFamily,

            string instanceTypeId,

            ImmutableArray<Outputs.InstanceTypesInstanceTypeLocalVolumeResult> localVolumes,

            ImmutableArray<Outputs.InstanceTypesInstanceTypeMemoryResult> memories,

            ImmutableArray<Outputs.InstanceTypesInstanceTypeNetworkResult> networks,

            ImmutableArray<Outputs.InstanceTypesInstanceTypeProcessorResult> processors,

            ImmutableArray<Outputs.InstanceTypesInstanceTypeRdmaResult> rdmas,

            ImmutableArray<Outputs.InstanceTypesInstanceTypeVolumeResult> volumes)
        {
            BaselineCredit = baselineCredit;
            Gpus = gpus;
            InitialCredit = initialCredit;
            InstanceTypeFamily = instanceTypeFamily;
            InstanceTypeId = instanceTypeId;
            LocalVolumes = localVolumes;
            Memories = memories;
            Networks = networks;
            Processors = processors;
            Rdmas = rdmas;
            Volumes = volumes;
        }
    }
}
