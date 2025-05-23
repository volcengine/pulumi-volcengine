// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Clb.Outputs
{

    [OutputType]
    public sealed class GetClbsClbIpv6AddressBandwidthResult
    {
        /// <summary>
        /// The peek bandwidth of the Ipv6 EIP assigned to CLB. Units: Mbps.
        /// </summary>
        public readonly int Bandwidth;
        /// <summary>
        /// The bandwidth package id of the Ipv6 EIP assigned to CLB.
        /// </summary>
        public readonly string BandwidthPackageId;
        /// <summary>
        /// The billing type of the Ipv6 EIP assigned to CLB. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic`.
        /// </summary>
        public readonly string BillingType;
        /// <summary>
        /// The ISP of the Ipv6 EIP assigned to CLB, the value can be `BGP`.
        /// </summary>
        public readonly string Isp;
        /// <summary>
        /// The network type of the CLB Ipv6 address.
        /// </summary>
        public readonly string NetworkType;

        [OutputConstructor]
        private GetClbsClbIpv6AddressBandwidthResult(
            int bandwidth,

            string bandwidthPackageId,

            string billingType,

            string isp,

            string networkType)
        {
            Bandwidth = bandwidth;
            BandwidthPackageId = bandwidthPackageId;
            BillingType = billingType;
            Isp = isp;
            NetworkType = networkType;
        }
    }
}
