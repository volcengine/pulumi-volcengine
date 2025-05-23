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
    public sealed class GetGroupsRocketmqGroupResult
    {
        /// <summary>
        /// The consumed topic information of the rocketmq group.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupsRocketmqGroupConsumedClientResult> ConsumedClients;
        /// <summary>
        /// The consumed topic information of the rocketmq group.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupsRocketmqGroupConsumedTopicResult> ConsumedTopics;
        /// <summary>
        /// The create time of the rocketmq group.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The description of the rocketmq group.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The id of rocketmq group. This field support fuzzy query.
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// The type of rocketmq group. Valid values: `TCP`.
        /// </summary>
        public readonly string GroupType;
        /// <summary>
        /// Whether the subscription relationship of consumer instance groups within the group is consistent.
        /// </summary>
        public readonly bool IsSubSame;
        /// <summary>
        /// The message delay time of the rocketmq group. The unit is milliseconds.
        /// </summary>
        public readonly string MessageDelayTime;
        /// <summary>
        /// The message model of the rocketmq group.
        /// </summary>
        public readonly string MessageModel;
        /// <summary>
        /// The status of the rocketmq group.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The total consume rate of the rocketmq group. The unit is per second.
        /// </summary>
        public readonly string TotalConsumeRate;
        /// <summary>
        /// The total amount of unconsumed messages.
        /// </summary>
        public readonly int TotalDiff;

        [OutputConstructor]
        private GetGroupsRocketmqGroupResult(
            ImmutableArray<Outputs.GetGroupsRocketmqGroupConsumedClientResult> consumedClients,

            ImmutableArray<Outputs.GetGroupsRocketmqGroupConsumedTopicResult> consumedTopics,

            string createTime,

            string description,

            string groupId,

            string groupType,

            bool isSubSame,

            string messageDelayTime,

            string messageModel,

            string status,

            string totalConsumeRate,

            int totalDiff)
        {
            ConsumedClients = consumedClients;
            ConsumedTopics = consumedTopics;
            CreateTime = createTime;
            Description = description;
            GroupId = groupId;
            GroupType = groupType;
            IsSubSame = isSubSame;
            MessageDelayTime = messageDelayTime;
            MessageModel = messageModel;
            Status = status;
            TotalConsumeRate = totalConsumeRate;
            TotalDiff = totalDiff;
        }
    }
}
