// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Ebs.Outputs
{

    [OutputType]
    public sealed class GetSnapshotGroupsSnapshotGroupResult
    {
        /// <summary>
        /// The creation time of the snapshot.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// The description of the snapshot.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The id of the snapshot group.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The image id of the snapshot.
        /// </summary>
        public readonly string ImageId;
        /// <summary>
        /// The instance id of snapshot group.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The name of snapshot group.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The project name of snapshot group.
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// The id of the snapshot group.
        /// </summary>
        public readonly string SnapshotGroupId;
        /// <summary>
        /// The snapshots of the snapshot group.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSnapshotGroupsSnapshotGroupSnapshotResult> Snapshots;
        /// <summary>
        /// A list of snapshot group status. Valid values: `creating`, `available`, `failed`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSnapshotGroupsSnapshotGroupTagResult> Tags;

        [OutputConstructor]
        private GetSnapshotGroupsSnapshotGroupResult(
            string creationTime,

            string description,

            string id,

            string imageId,

            string instanceId,

            string name,

            string projectName,

            string snapshotGroupId,

            ImmutableArray<Outputs.GetSnapshotGroupsSnapshotGroupSnapshotResult> snapshots,

            string status,

            ImmutableArray<Outputs.GetSnapshotGroupsSnapshotGroupTagResult> tags)
        {
            CreationTime = creationTime;
            Description = description;
            Id = id;
            ImageId = imageId;
            InstanceId = instanceId;
            Name = name;
            ProjectName = projectName;
            SnapshotGroupId = snapshotGroupId;
            Snapshots = snapshots;
            Status = status;
            Tags = tags;
        }
    }
}
