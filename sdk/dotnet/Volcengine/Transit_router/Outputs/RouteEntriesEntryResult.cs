// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Transit_router.Outputs
{

    [OutputType]
    public sealed class RouteEntriesEntryResult
    {
        /// <summary>
        /// The creation time of the route entry.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// Description of the transit router route entry.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The target network segment of the route entry.
        /// </summary>
        public readonly string DestinationCidrBlock;
        /// <summary>
        /// The status of the route entry.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The id of the route entry.
        /// </summary>
        public readonly string TransitRouterRouteEntryId;
        /// <summary>
        /// The name of the route entry.
        /// </summary>
        public readonly string TransitRouterRouteEntryName;
        /// <summary>
        /// The next hot id of the routing entry.
        /// </summary>
        public readonly string TransitRouterRouteEntryNextHopId;
        /// <summary>
        /// The next hop type of the routing entry. The value can be Attachment or BlackHole.
        /// </summary>
        public readonly string TransitRouterRouteEntryNextHopType;
        /// <summary>
        /// The type of the route entry.
        /// </summary>
        public readonly string TransitRouterRouteEntryType;
        /// <summary>
        /// The update time of the route entry.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private RouteEntriesEntryResult(
            string creationTime,

            string description,

            string destinationCidrBlock,

            string status,

            string transitRouterRouteEntryId,

            string transitRouterRouteEntryName,

            string transitRouterRouteEntryNextHopId,

            string transitRouterRouteEntryNextHopType,

            string transitRouterRouteEntryType,

            string updateTime)
        {
            CreationTime = creationTime;
            Description = description;
            DestinationCidrBlock = destinationCidrBlock;
            Status = status;
            TransitRouterRouteEntryId = transitRouterRouteEntryId;
            TransitRouterRouteEntryName = transitRouterRouteEntryName;
            TransitRouterRouteEntryNextHopId = transitRouterRouteEntryNextHopId;
            TransitRouterRouteEntryNextHopType = transitRouterRouteEntryNextHopType;
            TransitRouterRouteEntryType = transitRouterRouteEntryType;
            UpdateTime = updateTime;
        }
    }
}