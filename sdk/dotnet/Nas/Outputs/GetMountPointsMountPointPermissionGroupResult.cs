// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Nas.Outputs
{

    [OutputType]
    public sealed class GetMountPointsMountPointPermissionGroupResult
    {
        /// <summary>
        /// The creation time of the permission group.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The description of the permission group.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The number of the file system.
        /// </summary>
        public readonly int FileSystemCount;
        /// <summary>
        /// The file system type of the permission group.
        /// </summary>
        public readonly string FileSystemType;
        /// <summary>
        /// The list of the mount point.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMountPointsMountPointPermissionGroupMountPointResult> MountPoints;
        /// <summary>
        /// The id of the permission group.
        /// </summary>
        public readonly string PermissionGroupId;
        /// <summary>
        /// The name of the permission group.
        /// </summary>
        public readonly string PermissionGroupName;
        /// <summary>
        /// The number of the permission rule.
        /// </summary>
        public readonly int PermissionRuleCount;

        [OutputConstructor]
        private GetMountPointsMountPointPermissionGroupResult(
            string createTime,

            string description,

            int fileSystemCount,

            string fileSystemType,

            ImmutableArray<Outputs.GetMountPointsMountPointPermissionGroupMountPointResult> mountPoints,

            string permissionGroupId,

            string permissionGroupName,

            int permissionRuleCount)
        {
            CreateTime = createTime;
            Description = description;
            FileSystemCount = fileSystemCount;
            FileSystemType = fileSystemType;
            MountPoints = mountPoints;
            PermissionGroupId = permissionGroupId;
            PermissionGroupName = permissionGroupName;
            PermissionRuleCount = permissionRuleCount;
        }
    }
}
