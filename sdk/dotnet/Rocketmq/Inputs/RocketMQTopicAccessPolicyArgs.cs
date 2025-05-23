// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rocketmq.Inputs
{

    public sealed class RocketMQTopicAccessPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access key of the rocketmq key.
        /// </summary>
        [Input("accessKey", required: true)]
        public Input<string> AccessKey { get; set; } = null!;

        /// <summary>
        /// The authority of the rocketmq key for the current topic. Valid values: `ALL`, `PUB`, `SUB`, `DENY`. Default is `DENY`.
        /// </summary>
        [Input("authority", required: true)]
        public Input<string> Authority { get; set; } = null!;

        public RocketMQTopicAccessPolicyArgs()
        {
        }
        public static new RocketMQTopicAccessPolicyArgs Empty => new RocketMQTopicAccessPolicyArgs();
    }
}
