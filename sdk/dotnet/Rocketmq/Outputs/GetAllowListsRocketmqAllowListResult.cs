// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rocketmq.Outputs
{

    [OutputType]
    public sealed class GetAllowListsRocketmqAllowListResult
    {
        /// <summary>
        /// The description of the rocketmq allow list.
        /// </summary>
        public readonly string AllowListDesc;
        /// <summary>
        /// The id of the rocketmq allow list.
        /// </summary>
        public readonly string AllowListId;
        /// <summary>
        /// The number of ip address in the rocketmq allow list.
        /// </summary>
        public readonly int AllowListIpNum;
        /// <summary>
        /// The name of the rocketmq allow list.
        /// </summary>
        public readonly string AllowListName;
        /// <summary>
        /// The type of the rocketmq allow list.
        /// </summary>
        public readonly string AllowListType;
        /// <summary>
        /// The IP address or a range of IP addresses in CIDR format of the allow list.
        /// </summary>
        public readonly ImmutableArray<string> AllowLists;
        /// <summary>
        /// The number of the rocketmq instances associated with the allow list.
        /// </summary>
        public readonly int AssociatedInstanceNum;
        /// <summary>
        /// The associated instance information of the allow list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAllowListsRocketmqAllowListAssociatedInstanceResult> AssociatedInstances;
        /// <summary>
        /// The id of the rocketmq allow list.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAllowListsRocketmqAllowListResult(
            string allowListDesc,

            string allowListId,

            int allowListIpNum,

            string allowListName,

            string allowListType,

            ImmutableArray<string> allowLists,

            int associatedInstanceNum,

            ImmutableArray<Outputs.GetAllowListsRocketmqAllowListAssociatedInstanceResult> associatedInstances,

            string id)
        {
            AllowListDesc = allowListDesc;
            AllowListId = allowListId;
            AllowListIpNum = allowListIpNum;
            AllowListName = allowListName;
            AllowListType = allowListType;
            AllowLists = allowLists;
            AssociatedInstanceNum = associatedInstanceNum;
            AssociatedInstances = associatedInstances;
            Id = id;
        }
    }
}
