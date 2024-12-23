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
    public sealed class CloudServersCloudServerGpuGpusResult
    {
        /// <summary>
        /// The spec of gpu.
        /// </summary>
        public readonly Outputs.CloudServersCloudServerGpuGpusGpuSpecResult GpuSpec;
        /// <summary>
        /// The number of gpu.
        /// </summary>
        public readonly int Num;

        [OutputConstructor]
        private CloudServersCloudServerGpuGpusResult(
            Outputs.CloudServersCloudServerGpuGpusGpuSpecResult gpuSpec,

            int num)
        {
            GpuSpec = gpuSpec;
            Num = num;
        }
    }
}
