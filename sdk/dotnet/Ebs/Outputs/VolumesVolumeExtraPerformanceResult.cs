// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Ebs.Outputs
{

    [OutputType]
    public sealed class VolumesVolumeExtraPerformanceResult
    {
        /// <summary>
        /// The type of extra performance for volume.
        /// </summary>
        public readonly string ExtraPerformanceTypeId;
        /// <summary>
        /// The total IOPS performance size for volume.
        /// </summary>
        public readonly int Iops;
        /// <summary>
        /// The total Throughput performance size for volume. Unit: MB/s.
        /// </summary>
        public readonly int Throughput;

        [OutputConstructor]
        private VolumesVolumeExtraPerformanceResult(
            string extraPerformanceTypeId,

            int iops,

            int throughput)
        {
            ExtraPerformanceTypeId = extraPerformanceTypeId;
            Iops = iops;
            Throughput = throughput;
        }
    }
}