// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls.Outputs
{

    [OutputType]
    public sealed class GetTopicsTlsTopicResult
    {
        /// <summary>
        /// Whether to enable automatic partition splitting function of the tls topic.
        /// </summary>
        public readonly bool AutoSplit;
        /// <summary>
        /// The create time of the tls topic.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The description of the tls topic.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Whether to enable WebTracking function of the tls topic.
        /// </summary>
        public readonly bool EnableTracking;
        /// <summary>
        /// The ID of the tls topic.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The max count of shards in the tls topic.
        /// </summary>
        public readonly int MaxSplitShard;
        /// <summary>
        /// The modify time of the tls topic.
        /// </summary>
        public readonly string ModifyTime;
        /// <summary>
        /// The project id of tls topic.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// The count of shards in the tls topic.
        /// </summary>
        public readonly int ShardCount;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTopicsTlsTopicTagResult> Tags;
        /// <summary>
        /// The format of the time field.
        /// </summary>
        public readonly string TimeFormat;
        /// <summary>
        /// The name of the time field.
        /// </summary>
        public readonly string TimeKey;
        /// <summary>
        /// The id of tls topic. This field supports fuzzy queries. It is not supported to specify both TopicName and TopicId at the same time.
        /// </summary>
        public readonly string TopicId;
        /// <summary>
        /// The name of tls topic. This field supports fuzzy queries. It is not supported to specify both TopicName and TopicId at the same time.
        /// </summary>
        public readonly string TopicName;
        /// <summary>
        /// The data storage time of the tls topic. Unit: Day.
        /// </summary>
        public readonly int Ttl;

        [OutputConstructor]
        private GetTopicsTlsTopicResult(
            bool autoSplit,

            string createTime,

            string description,

            bool enableTracking,

            string id,

            int maxSplitShard,

            string modifyTime,

            string projectId,

            int shardCount,

            ImmutableArray<Outputs.GetTopicsTlsTopicTagResult> tags,

            string timeFormat,

            string timeKey,

            string topicId,

            string topicName,

            int ttl)
        {
            AutoSplit = autoSplit;
            CreateTime = createTime;
            Description = description;
            EnableTracking = enableTracking;
            Id = id;
            MaxSplitShard = maxSplitShard;
            ModifyTime = modifyTime;
            ProjectId = projectId;
            ShardCount = shardCount;
            Tags = tags;
            TimeFormat = timeFormat;
            TimeKey = timeKey;
            TopicId = topicId;
            TopicName = topicName;
            Ttl = ttl;
        }
    }
}
