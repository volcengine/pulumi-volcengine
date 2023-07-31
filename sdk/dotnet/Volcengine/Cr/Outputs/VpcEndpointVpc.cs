// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Cr.Outputs
{

    [OutputType]
    public sealed class VpcEndpointVpc
    {
        /// <summary>
        /// The id of the account. When you need to expose the Enterprise Edition instance to a VPC under another primary account, you need to specify the ID of the primary account to which the VPC belongs.
        /// </summary>
        public readonly int? AccountId;
        /// <summary>
        /// The id of the subnet. If not specified, the subnet with the most remaining IPs under the VPC will be automatically selected.
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// The id of the vpc.
        /// </summary>
        public readonly string? VpcId;

        [OutputConstructor]
        private VpcEndpointVpc(
            int? accountId,

            string? subnetId,

            string? vpcId)
        {
            AccountId = accountId;
            SubnetId = subnetId;
            VpcId = vpcId;
        }
    }
}