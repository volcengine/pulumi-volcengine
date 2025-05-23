// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cloud_identity.Outputs
{

    [OutputType]
    public sealed class GroupsGroupResult
    {
        /// <summary>
        /// The created time of the cloud identity group.
        /// </summary>
        public readonly string CreatedTime;
        /// <summary>
        /// The description of the cloud identity group.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The display name of cloud identity group.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The id of the cloud identity group.
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// The name of cloud identity group.
        /// </summary>
        public readonly string GroupName;
        /// <summary>
        /// The id of the cloud identity group.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The join type of cloud identity group. Valid values: `Auto`, `Manual`.
        /// </summary>
        public readonly string JoinType;
        /// <summary>
        /// The source of the cloud identity group.
        /// </summary>
        public readonly string Source;
        /// <summary>
        /// The updated time of the cloud identity group.
        /// </summary>
        public readonly string UpdatedTime;

        [OutputConstructor]
        private GroupsGroupResult(
            string createdTime,

            string description,

            string displayName,

            string groupId,

            string groupName,

            string id,

            string joinType,

            string source,

            string updatedTime)
        {
            CreatedTime = createdTime;
            Description = description;
            DisplayName = displayName;
            GroupId = groupId;
            GroupName = groupName;
            Id = id;
            JoinType = joinType;
            Source = source;
            UpdatedTime = updatedTime;
        }
    }
}
