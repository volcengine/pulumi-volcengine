// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Vke.Outputs
{

    [OutputType]
    public sealed class ClustersClusterServicesConfigResult
    {
        /// <summary>
        /// The IPv4 private network address exposed by the service.
        /// </summary>
        public readonly ImmutableArray<string> ServiceCidrsv4s;

        [OutputConstructor]
        private ClustersClusterServicesConfigResult(ImmutableArray<string> serviceCidrsv4s)
        {
            ServiceCidrsv4s = serviceCidrsv4s;
        }
    }
}