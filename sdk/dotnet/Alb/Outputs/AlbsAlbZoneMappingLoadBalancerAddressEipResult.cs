// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Alb.Outputs
{

    [OutputType]
    public sealed class AlbsAlbZoneMappingLoadBalancerAddressEipResult
    {
        /// <summary>
        /// The association mode of the Alb. This parameter has a query value only when the type of the Eip is `anycast`.
        /// </summary>
        public readonly string AssociationMode;
        /// <summary>
        /// The peek bandwidth of the Ipv6 Eip assigned to Alb. Units: Mbps.
        /// </summary>
        public readonly int Bandwidth;
        /// <summary>
        /// The Eip address of the Alb.
        /// </summary>
        public readonly string EipAddress;
        /// <summary>
        /// The billing type of the Eip assigned to Alb. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic`.
        /// </summary>
        public readonly string EipBillingType;
        /// <summary>
        /// The Eip type of the Alb.
        /// </summary>
        public readonly string EipType;
        /// <summary>
        /// The ISP of the Ipv6 Eip assigned to Alb, the value can be `BGP`.
        /// </summary>
        public readonly string Isp;
        /// <summary>
        /// The pop locations of the Alb. This parameter has a query value only when the type of the Eip is `anycast`.
        /// </summary>
        public readonly ImmutableArray<Outputs.AlbsAlbZoneMappingLoadBalancerAddressEipPopLocationResult> PopLocations;
        /// <summary>
        /// The security protection types of the Alb.
        /// </summary>
        public readonly ImmutableArray<string> SecurityProtectionTypes;

        [OutputConstructor]
        private AlbsAlbZoneMappingLoadBalancerAddressEipResult(
            string associationMode,

            int bandwidth,

            string eipAddress,

            string eipBillingType,

            string eipType,

            string isp,

            ImmutableArray<Outputs.AlbsAlbZoneMappingLoadBalancerAddressEipPopLocationResult> popLocations,

            ImmutableArray<string> securityProtectionTypes)
        {
            AssociationMode = associationMode;
            Bandwidth = bandwidth;
            EipAddress = eipAddress;
            EipBillingType = eipBillingType;
            EipType = eipType;
            Isp = isp;
            PopLocations = popLocations;
            SecurityProtectionTypes = securityProtectionTypes;
        }
    }
}