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
    public sealed class ClbsClbEipBillingConfigResult
    {
        /// <summary>
        /// The peek bandwidth of the Ipv6 EIP assigned to CLB. Units: Mbps.
        /// </summary>
        public readonly int Bandwidth;
        /// <summary>
        /// The billing type of the EIP assigned to CLB. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.
        /// </summary>
        public readonly string EipBillingType;
        /// <summary>
        /// The ISP of the Ipv6 EIP assigned to CLB, the value can be `BGP`.
        /// </summary>
        public readonly string Isp;

        [OutputConstructor]
        private ClbsClbEipBillingConfigResult(
            int bandwidth,

            string eipBillingType,

            string isp)
        {
            Bandwidth = bandwidth;
            EipBillingType = eipBillingType;
            Isp = isp;
        }
    }
}