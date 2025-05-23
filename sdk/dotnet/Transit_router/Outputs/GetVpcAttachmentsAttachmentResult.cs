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
    public sealed class GetVpcAttachmentsAttachmentResult
    {
        /// <summary>
        /// The collection of attach points.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVpcAttachmentsAttachmentAttachPointResult> AttachPoints;
        /// <summary>
        /// Whether to auto publish route of the transit router to vpc instance.
        /// </summary>
        public readonly bool AutoPublishRouteEnabled;
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
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVpcAttachmentsAttachmentTagResult> Tags;
        /// <summary>
        /// The id of the transit router attachment.
        /// </summary>
        public readonly string TransitRouterAttachmentId;
        /// <summary>
        /// The name of the transit router attachment.
        /// </summary>
        public readonly string TransitRouterAttachmentName;
        /// <summary>
        /// The id of transit router.
        /// </summary>
        public readonly string TransitRouterId;
        /// <summary>
        /// The update time.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// The id of vpc.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private GetVpcAttachmentsAttachmentResult(
            ImmutableArray<Outputs.GetVpcAttachmentsAttachmentAttachPointResult> attachPoints,

            bool autoPublishRouteEnabled,

            string creationTime,

            string description,

            string status,

            ImmutableArray<Outputs.GetVpcAttachmentsAttachmentTagResult> tags,

            string transitRouterAttachmentId,

            string transitRouterAttachmentName,

            string transitRouterId,

            string updateTime,

            string vpcId)
        {
            AttachPoints = attachPoints;
            AutoPublishRouteEnabled = autoPublishRouteEnabled;
            CreationTime = creationTime;
            Description = description;
            Status = status;
            Tags = tags;
            TransitRouterAttachmentId = transitRouterAttachmentId;
            TransitRouterAttachmentName = transitRouterAttachmentName;
            TransitRouterId = transitRouterId;
            UpdateTime = updateTime;
            VpcId = vpcId;
        }
    }
}
