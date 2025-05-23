// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Transit_router.Outputs
{

    [OutputType]
    public sealed class GetRouteTablePropagationsPropagationResult
    {
        /// <summary>
        /// The creation time of the route table propagation.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// The status of the route table.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The ID of the network instance connection.
        /// </summary>
        public readonly string TransitRouterAttachmentId;
        /// <summary>
        /// The ID of the routing table associated with the transit router instance.
        /// </summary>
        public readonly string TransitRouterRouteTableId;

        [OutputConstructor]
        private GetRouteTablePropagationsPropagationResult(
            string creationTime,

            string status,

            string transitRouterAttachmentId,

            string transitRouterRouteTableId)
        {
            CreationTime = creationTime;
            Status = status;
            TransitRouterAttachmentId = transitRouterAttachmentId;
            TransitRouterRouteTableId = transitRouterRouteTableId;
        }
    }
}
