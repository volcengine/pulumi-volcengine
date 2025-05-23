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
    public sealed class GetKafkaConsumersDataResult
    {
        /// <summary>
        /// Whether allow consume.
        /// </summary>
        public readonly bool AllowConsume;
        /// <summary>
        /// The topic of consume.
        /// </summary>
        public readonly string ConsumeTopic;
        /// <summary>
        /// The ID of Topic.
        /// </summary>
        public readonly string TopicId;

        [OutputConstructor]
        private GetKafkaConsumersDataResult(
            bool allowConsume,

            string consumeTopic,

            string topicId)
        {
            AllowConsume = allowConsume;
            ConsumeTopic = consumeTopic;
            TopicId = topicId;
        }
    }
}
