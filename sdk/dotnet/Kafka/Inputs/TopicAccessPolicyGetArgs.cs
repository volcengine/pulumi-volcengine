// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Kafka.Inputs
{

    public sealed class TopicAccessPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access policy of SASL user. Valid values: `PubSub`, `Pub`, `Sub`.
        /// </summary>
        [Input("accessPolicy", required: true)]
        public Input<string> AccessPolicy { get; set; } = null!;

        /// <summary>
        /// The name of SASL user.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public TopicAccessPolicyGetArgs()
        {
        }
        public static new TopicAccessPolicyGetArgs Empty => new TopicAccessPolicyGetArgs();
    }
}