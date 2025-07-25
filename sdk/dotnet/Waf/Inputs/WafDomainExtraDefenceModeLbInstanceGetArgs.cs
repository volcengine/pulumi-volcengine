// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Waf.Inputs
{

    public sealed class WafDomainExtraDefenceModeLbInstanceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set the protection mode for exceptional ALB instances. Works only on modified scenes.
        /// </summary>
        [Input("defenceMode")]
        public Input<int>? DefenceMode { get; set; }

        /// <summary>
        /// The Id of ALB instance. Works only on modified scenes.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public WafDomainExtraDefenceModeLbInstanceGetArgs()
        {
        }
        public static new WafDomainExtraDefenceModeLbInstanceGetArgs Empty => new WafDomainExtraDefenceModeLbInstanceGetArgs();
    }
}
