// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vmp.Inputs
{

    public sealed class NotifyPolicyLevelGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("channels", required: true)]
        private InputList<string>? _channels;

        /// <summary>
        /// The alarm notification method of the alarm notification policy, the optional value can be `Email`, `Webhook`, `LarkBotWebhook`, `DingTalkBotWebhook`, `WeComBotWebhook`.
        /// </summary>
        public InputList<string> Channels
        {
            get => _channels ?? (_channels = new InputList<string>());
            set => _channels = value;
        }

        [Input("contactGroupIds", required: true)]
        private InputList<string>? _contactGroupIds;

        /// <summary>
        /// The contact group for the alarm notification policy.
        /// </summary>
        public InputList<string> ContactGroupIds
        {
            get => _contactGroupIds ?? (_contactGroupIds = new InputList<string>());
            set => _contactGroupIds = value;
        }

        /// <summary>
        /// The level of the policy, the value can be one of the following: `P0`, `P1`, `P2`.
        /// </summary>
        [Input("level", required: true)]
        public Input<string> Level { get; set; } = null!;

        [Input("resolvedChannels")]
        private InputList<string>? _resolvedChannels;

        /// <summary>
        /// The resolved alarm notification method of the alarm notification policy, the optional value can be `Email`, `Webhook`, `LarkBotWebhook`, `DingTalkBotWebhook`, `WeComBotWebhook`.
        /// </summary>
        public InputList<string> ResolvedChannels
        {
            get => _resolvedChannels ?? (_resolvedChannels = new InputList<string>());
            set => _resolvedChannels = value;
        }

        public NotifyPolicyLevelGetArgs()
        {
        }
        public static new NotifyPolicyLevelGetArgs Empty => new NotifyPolicyLevelGetArgs();
    }
}
