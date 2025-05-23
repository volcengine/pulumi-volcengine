// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cloudfs.Outputs
{

    [OutputType]
    public sealed class GetAccessesAccessResult
    {
        /// <summary>
        /// The account id of access.
        /// </summary>
        public readonly int AccessAccountId;
        /// <summary>
        /// The id of access.
        /// </summary>
        public readonly string AccessId;
        /// <summary>
        /// The service name of access.
        /// </summary>
        public readonly string AccessServiceName;
        /// <summary>
        /// The creation time.
        /// </summary>
        public readonly string CreatedTime;
        /// <summary>
        /// The name of file system.
        /// </summary>
        public readonly string FsName;
        /// <summary>
        /// Whether is default access.
        /// </summary>
        public readonly bool IsDefault;
        /// <summary>
        /// The id of security group.
        /// </summary>
        public readonly string SecurityGroupId;
        /// <summary>
        /// The status of access.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The id of subnet.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// The id of vpc.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// Whether to enable all vpc route.
        /// </summary>
        public readonly bool VpcRouteEnabled;

        [OutputConstructor]
        private GetAccessesAccessResult(
            int accessAccountId,

            string accessId,

            string accessServiceName,

            string createdTime,

            string fsName,

            bool isDefault,

            string securityGroupId,

            string status,

            string subnetId,

            string vpcId,

            bool vpcRouteEnabled)
        {
            AccessAccountId = accessAccountId;
            AccessId = accessId;
            AccessServiceName = accessServiceName;
            CreatedTime = createdTime;
            FsName = fsName;
            IsDefault = isDefault;
            SecurityGroupId = securityGroupId;
            Status = status;
            SubnetId = subnetId;
            VpcId = vpcId;
            VpcRouteEnabled = vpcRouteEnabled;
        }
    }
}
