// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cen.Outputs
{

    [OutputType]
    public sealed class GetRouteEntriesCenRouteEntryResult
    {
        /// <summary>
        /// The AS path of the cen route entry.
        /// </summary>
        public readonly ImmutableArray<string> AsPaths;
        /// <summary>
        /// A cen ID.
        /// </summary>
        public readonly string CenId;
        /// <summary>
        /// A destination cidr block.
        /// </summary>
        public readonly string DestinationCidrBlock;
        /// <summary>
        /// An instance ID.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// An instance region ID.
        /// </summary>
        public readonly string InstanceRegionId;
        /// <summary>
        /// An instance type.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// The publish status of the cen route entry.
        /// </summary>
        public readonly string PublishStatus;
        /// <summary>
        /// The status of the cen route entry.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetRouteEntriesCenRouteEntryResult(
            ImmutableArray<string> asPaths,

            string cenId,

            string destinationCidrBlock,

            string instanceId,

            string instanceRegionId,

            string instanceType,

            string publishStatus,

            string status)
        {
            AsPaths = asPaths;
            CenId = cenId;
            DestinationCidrBlock = destinationCidrBlock;
            InstanceId = instanceId;
            InstanceRegionId = instanceRegionId;
            InstanceType = instanceType;
            PublishStatus = publishStatus;
            Status = status;
        }
    }
}
