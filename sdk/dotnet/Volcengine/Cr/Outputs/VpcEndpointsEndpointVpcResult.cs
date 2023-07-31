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
    public sealed class VpcEndpointsEndpointVpcResult
    {
        /// <summary>
        /// The id of the account.
        /// </summary>
        public readonly int AccountId;
        /// <summary>
        /// The creation time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The IP address of the mirror repository in the VPC.
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// The region id.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The status of the vpc endpoint.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The ID of the subnet.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// The ID of the vpc.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private VpcEndpointsEndpointVpcResult(
            int accountId,

            string createTime,

            string ip,

            string region,

            string status,

            string subnetId,

            string vpcId)
        {
            AccountId = accountId;
            CreateTime = createTime;
            Ip = ip;
            Region = region;
            Status = status;
            SubnetId = subnetId;
            VpcId = vpcId;
        }
    }
}