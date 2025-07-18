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
    public sealed class GatewayServicesGatewayServiceCustomDomainResult
    {
        /// <summary>
        /// The domain of the api gateway service.
        /// </summary>
        public readonly string Domain;
        /// <summary>
        /// The Id of the api gateway service.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GatewayServicesGatewayServiceCustomDomainResult(
            string domain,

            string id)
        {
            Domain = domain;
            Id = id;
        }
    }
}
