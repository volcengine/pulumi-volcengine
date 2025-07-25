// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vmp.Outputs
{

    [OutputType]
    public sealed class GetContactsContactResult
    {
        /// <summary>
        /// A list of contact group ids.
        /// </summary>
        public readonly ImmutableArray<string> ContactGroupIds;
        /// <summary>
        /// The create time of contact.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The ding talk bot webhook of contact.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetContactsContactDingTalkBotWebhookResult> DingTalkBotWebhooks;
        /// <summary>
        /// The email of contact.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// Whether the email of contact active.
        /// </summary>
        public readonly bool EmailActive;
        /// <summary>
        /// The ID of contact.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The lark bot webhook of contact.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetContactsContactLarkBotWebhookResult> LarkBotWebhooks;
        /// <summary>
        /// The name of contact.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Whether phone number is active.
        /// </summary>
        public readonly bool PhoneNumberActive;
        /// <summary>
        /// The phone number of contact.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetContactsContactPhoneNumberResult> PhoneNumbers;
        /// <summary>
        /// The we com bot webhook of contact.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetContactsContactWeComBotWebhookResult> WeComBotWebhooks;
        /// <summary>
        /// The webhook of contact.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetContactsContactWebhookResult> Webhooks;

        [OutputConstructor]
        private GetContactsContactResult(
            ImmutableArray<string> contactGroupIds,

            string createTime,

            ImmutableArray<Outputs.GetContactsContactDingTalkBotWebhookResult> dingTalkBotWebhooks,

            string email,

            bool emailActive,

            string id,

            ImmutableArray<Outputs.GetContactsContactLarkBotWebhookResult> larkBotWebhooks,

            string name,

            bool phoneNumberActive,

            ImmutableArray<Outputs.GetContactsContactPhoneNumberResult> phoneNumbers,

            ImmutableArray<Outputs.GetContactsContactWeComBotWebhookResult> weComBotWebhooks,

            ImmutableArray<Outputs.GetContactsContactWebhookResult> webhooks)
        {
            ContactGroupIds = contactGroupIds;
            CreateTime = createTime;
            DingTalkBotWebhooks = dingTalkBotWebhooks;
            Email = email;
            EmailActive = emailActive;
            Id = id;
            LarkBotWebhooks = larkBotWebhooks;
            Name = name;
            PhoneNumberActive = phoneNumberActive;
            PhoneNumbers = phoneNumbers;
            WeComBotWebhooks = weComBotWebhooks;
            Webhooks = webhooks;
        }
    }
}
