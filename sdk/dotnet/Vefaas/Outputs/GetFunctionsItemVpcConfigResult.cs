// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vefaas.Outputs
{

    [OutputType]
    public sealed class GetFunctionsItemVpcConfigResult
    {
        /// <summary>
        /// Function access to the public network switch.
        /// </summary>
        public readonly bool EnableSharedInternetAccess;
        /// <summary>
        /// Whether the function enables private network access.
        /// </summary>
        public readonly bool EnableVpc;
        /// <summary>
        /// The ID of security group.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// The ID of subnet.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        /// <summary>
        /// The ID of VPC.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private GetFunctionsItemVpcConfigResult(
            bool enableSharedInternetAccess,

            bool enableVpc,

            ImmutableArray<string> securityGroupIds,

            ImmutableArray<string> subnetIds,

            string vpcId)
        {
            EnableSharedInternetAccess = enableSharedInternetAccess;
            EnableVpc = enableVpc;
            SecurityGroupIds = securityGroupIds;
            SubnetIds = subnetIds;
            VpcId = vpcId;
        }
    }
}
