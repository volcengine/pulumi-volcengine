// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Nat.Outputs
{

    [OutputType]
    public sealed class GatewaysNatGatewayEipAddressResult
    {
        /// <summary>
        /// The ID of Eip.
        /// </summary>
        public readonly string AllocationId;
        /// <summary>
        /// The address of Eip.
        /// </summary>
        public readonly string EipAddress;
        /// <summary>
        /// The using status of Eip.
        /// </summary>
        public readonly string UsingStatus;

        [OutputConstructor]
        private GatewaysNatGatewayEipAddressResult(
            string allocationId,

            string eipAddress,

            string usingStatus)
        {
            AllocationId = allocationId;
            EipAddress = eipAddress;
            UsingStatus = usingStatus;
        }
    }
}