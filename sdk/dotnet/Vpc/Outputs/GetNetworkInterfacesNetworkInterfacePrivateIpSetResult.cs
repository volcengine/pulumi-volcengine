// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpc.Outputs
{

    [OutputType]
    public sealed class GetNetworkInterfacesNetworkInterfacePrivateIpSetResult
    {
        /// <summary>
        /// The public IP that secondary private network IP associated with.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNetworkInterfacesNetworkInterfacePrivateIpSetAssociatedElasticIpResult> AssociatedElasticIps;
        /// <summary>
        /// Whether the network interface is primary IP address.
        /// </summary>
        public readonly bool Primary;
        /// <summary>
        /// The secondary private network IP address of the network interface card.
        /// </summary>
        public readonly string PrivateIpAddress;

        [OutputConstructor]
        private GetNetworkInterfacesNetworkInterfacePrivateIpSetResult(
            ImmutableArray<Outputs.GetNetworkInterfacesNetworkInterfacePrivateIpSetAssociatedElasticIpResult> associatedElasticIps,

            bool primary,

            string privateIpAddress)
        {
            AssociatedElasticIps = associatedElasticIps;
            Primary = primary;
            PrivateIpAddress = privateIpAddress;
        }
    }
}
