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
    public sealed class VpnAttachmentsAttachmentResult
    {
        /// <summary>
        /// The create time.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// The description info.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The status of the transit router.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The id of the transit router attachment.
        /// </summary>
        public readonly string TransitRouterAttachmentId;
        /// <summary>
        /// The name of the transit router attachment.
        /// </summary>
        public readonly string TransitRouterAttachmentName;
        /// <summary>
        /// The id of the transit router.
        /// </summary>
        public readonly string TransitRouterId;
        /// <summary>
        /// The update time.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// The ID of the IPSec connection.
        /// </summary>
        public readonly string VpnConnectionId;
        /// <summary>
        /// The ID of the availability zone.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private VpnAttachmentsAttachmentResult(
            string creationTime,

            string description,

            string status,

            string transitRouterAttachmentId,

            string transitRouterAttachmentName,

            string transitRouterId,

            string updateTime,

            string vpnConnectionId,

            string zoneId)
        {
            CreationTime = creationTime;
            Description = description;
            Status = status;
            TransitRouterAttachmentId = transitRouterAttachmentId;
            TransitRouterAttachmentName = transitRouterAttachmentName;
            TransitRouterId = transitRouterId;
            UpdateTime = updateTime;
            VpnConnectionId = vpnConnectionId;
            ZoneId = zoneId;
        }
    }
}