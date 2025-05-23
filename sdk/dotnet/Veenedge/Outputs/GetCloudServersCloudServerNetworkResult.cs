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
    public sealed class GetCloudServersCloudServerNetworkResult
    {
        /// <summary>
        /// The peak of bandwidth.
        /// </summary>
        public readonly string BandwidthPeak;
        /// <summary>
        /// Whether enable ipv6.
        /// </summary>
        public readonly bool EnableIpv6;
        /// <summary>
        /// The internal peak of bandwidth.
        /// </summary>
        public readonly string InternalBandwidthPeak;

        [OutputConstructor]
        private GetCloudServersCloudServerNetworkResult(
            string bandwidthPeak,

            bool enableIpv6,

            string internalBandwidthPeak)
        {
            BandwidthPeak = bandwidthPeak;
            EnableIpv6 = enableIpv6;
            InternalBandwidthPeak = internalBandwidthPeak;
        }
    }
}
