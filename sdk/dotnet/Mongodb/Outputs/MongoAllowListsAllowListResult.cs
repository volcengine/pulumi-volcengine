// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Mongodb.Outputs
{

    [OutputType]
    public sealed class MongoAllowListsAllowListResult
    {
        /// <summary>
        /// The list of IP address in allow list.
        /// </summary>
        public readonly string AllowList;
        /// <summary>
        /// The description of allow list.
        /// </summary>
        public readonly string AllowListDesc;
        /// <summary>
        /// The ID of allow list.
        /// </summary>
        public readonly string AllowListId;
        /// <summary>
        /// The number of allow list IPs.
        /// </summary>
        public readonly int AllowListIpNum;
        /// <summary>
        /// The allow list name.
        /// </summary>
        public readonly string AllowListName;
        /// <summary>
        /// The IP address type in allow list.
        /// </summary>
        public readonly string AllowListType;
        /// <summary>
        /// The total number of instances bound under the allow list.
        /// </summary>
        public readonly int AssociatedInstanceNum;
        /// <summary>
        /// The list of associated instances.
        /// </summary>
        public readonly ImmutableArray<Outputs.MongoAllowListsAllowListAssociatedInstanceResult> AssociatedInstances;
        /// <summary>
        /// The project name of the allow list.
        /// </summary>
        public readonly string ProjectName;

        [OutputConstructor]
        private MongoAllowListsAllowListResult(
            string allowList,

            string allowListDesc,

            string allowListId,

            int allowListIpNum,

            string allowListName,

            string allowListType,

            int associatedInstanceNum,

            ImmutableArray<Outputs.MongoAllowListsAllowListAssociatedInstanceResult> associatedInstances,

            string projectName)
        {
            AllowList = allowList;
            AllowListDesc = allowListDesc;
            AllowListId = allowListId;
            AllowListIpNum = allowListIpNum;
            AllowListName = allowListName;
            AllowListType = allowListType;
            AssociatedInstanceNum = associatedInstanceNum;
            AssociatedInstances = associatedInstances;
            ProjectName = projectName;
        }
    }
}
