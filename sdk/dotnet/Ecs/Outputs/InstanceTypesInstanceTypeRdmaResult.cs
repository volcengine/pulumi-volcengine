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
    public sealed class InstanceTypesInstanceTypeRdmaResult
    {
        /// <summary>
        /// Number of RDMA network cards.
        /// </summary>
        public readonly int RdmaNetworkInterfaces;

        [OutputConstructor]
        private InstanceTypesInstanceTypeRdmaResult(int rdmaNetworkInterfaces)
        {
            RdmaNetworkInterfaces = rdmaNetworkInterfaces;
        }
    }
}