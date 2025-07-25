// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Waf.Outputs
{

    [OutputType]
    public sealed class GetDomainsDataProtocolPortsResult
    {
        /// <summary>
        /// Ports supported by the HTTP protocol.
        /// </summary>
        public readonly ImmutableArray<int> Http;
        /// <summary>
        /// Ports supported by the HTTPs protocol.
        /// </summary>
        public readonly ImmutableArray<int> Https;

        [OutputConstructor]
        private GetDomainsDataProtocolPortsResult(
            ImmutableArray<int> http,

            ImmutableArray<int> https)
        {
            Http = http;
            Https = https;
        }
    }
}
