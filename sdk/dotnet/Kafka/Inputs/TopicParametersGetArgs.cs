// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Kafka.Inputs
{

    public sealed class TopicParametersGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The retention hours of log. Unit: hour. Valid values: 0-2160. Default is 72.
        /// </summary>
        [Input("logRetentionHours")]
        public Input<int>? LogRetentionHours { get; set; }

        /// <summary>
        /// The max byte of message. Unit: MB. Valid values: 1-12. Default is 10.
        /// </summary>
        [Input("messageMaxByte")]
        public Input<int>? MessageMaxByte { get; set; }

        /// <summary>
        /// The min number of sync replica. The default value is the replica number minus 1.
        /// </summary>
        [Input("minInsyncReplicaNumber")]
        public Input<int>? MinInsyncReplicaNumber { get; set; }

        public TopicParametersGetArgs()
        {
        }
        public static new TopicParametersGetArgs Empty => new TopicParametersGetArgs();
    }
}