// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Apig.Outputs
{

    [OutputType]
    public sealed class ApigUpstreamUpstreamSpecDomainDomainList
    {
        /// <summary>
        /// The domain of apig upstream.
        /// </summary>
        public readonly string Domain;
        /// <summary>
        /// The port of domain. Default is 80 for HTTP, 443 for HTTPS.
        /// </summary>
        public readonly int? Port;

        [OutputConstructor]
        private ApigUpstreamUpstreamSpecDomainDomainList(
            string domain,

            int? port)
        {
            Domain = domain;
            Port = port;
        }
    }
}
