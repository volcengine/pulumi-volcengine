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
    public sealed class GetTransitRoutersTransitRouterResult
    {
        /// <summary>
        /// The ID of account.
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// The asn of the transit router.
        /// </summary>
        public readonly int Asn;
        /// <summary>
        /// The business status of the transit router.
        /// </summary>
        public readonly string BusinessStatus;
        /// <summary>
        /// The create time.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// The description info.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The grant status of the transit router.
        /// </summary>
        public readonly string GrantStatus;
        /// <summary>
        /// The ID of the transit router.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The overdue time.
        /// </summary>
        public readonly string OverdueTime;
        /// <summary>
        /// The ProjectName of the transit router.
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// The status of the transit router.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTransitRoutersTransitRouterTagResult> Tags;
        /// <summary>
        /// The attachments of transit router.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTransitRoutersTransitRouterTransitRouterAttachmentResult> TransitRouterAttachments;
        /// <summary>
        /// The ID of the transit router.
        /// </summary>
        public readonly string TransitRouterId;
        /// <summary>
        /// The name info.
        /// </summary>
        public readonly string TransitRouterName;
        /// <summary>
        /// The update time.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetTransitRoutersTransitRouterResult(
            string accountId,

            int asn,

            string businessStatus,

            string creationTime,

            string description,

            string grantStatus,

            string id,

            string overdueTime,

            string projectName,

            string status,

            ImmutableArray<Outputs.GetTransitRoutersTransitRouterTagResult> tags,

            ImmutableArray<Outputs.GetTransitRoutersTransitRouterTransitRouterAttachmentResult> transitRouterAttachments,

            string transitRouterId,

            string transitRouterName,

            string updateTime)
        {
            AccountId = accountId;
            Asn = asn;
            BusinessStatus = businessStatus;
            CreationTime = creationTime;
            Description = description;
            GrantStatus = grantStatus;
            Id = id;
            OverdueTime = overdueTime;
            ProjectName = projectName;
            Status = status;
            Tags = tags;
            TransitRouterAttachments = transitRouterAttachments;
            TransitRouterId = transitRouterId;
            TransitRouterName = transitRouterName;
            UpdateTime = updateTime;
        }
    }
}
