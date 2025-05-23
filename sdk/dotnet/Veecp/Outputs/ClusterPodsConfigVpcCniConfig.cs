// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Veecp.Outputs
{

    [OutputType]
    public sealed class ClusterPodsConfigVpcCniConfig
    {
        /// <summary>
        /// A list of Pod subnet IDs for the VPC-CNI container network.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;

        [OutputConstructor]
        private ClusterPodsConfigVpcCniConfig(ImmutableArray<string> subnetIds)
        {
            SubnetIds = subnetIds;
        }
    }
}
