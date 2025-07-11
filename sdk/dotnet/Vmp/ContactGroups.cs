// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vmp
{
    [Obsolete(@"volcengine.vmp.ContactGroups has been deprecated in favor of volcengine.vmp.getContactGroups")]
    public static class ContactGroups
    {
        /// <summary>
        /// Use this data source to query detailed information of vmp contact groups
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var fooContact = new Volcengine.Vmp.Contact("fooContact", new()
        ///     {
        ///         Email = "acctest1@tftest.com",
        ///         Webhook = new Volcengine.Vmp.Inputs.ContactWebhookArgs
        ///         {
        ///             Address = "https://www.acctest1.com",
        ///         },
        ///         LarkBotWebhook = new Volcengine.Vmp.Inputs.ContactLarkBotWebhookArgs
        ///         {
        ///             Address = "https://www.acctest1.com",
        ///         },
        ///         DingTalkBotWebhook = new Volcengine.Vmp.Inputs.ContactDingTalkBotWebhookArgs
        ///         {
        ///             Address = "https://www.dingacctest1.com",
        ///             AtMobiles = new[]
        ///             {
        ///                 "18046891812",
        ///             },
        ///         },
        ///         PhoneNumber = new Volcengine.Vmp.Inputs.ContactPhoneNumberArgs
        ///         {
        ///             CountryCode = "+86",
        ///             Number = "18310101010",
        ///         },
        ///     });
        /// 
        ///     var foo1 = new Volcengine.Vmp.Contact("foo1", new()
        ///     {
        ///         Email = "acctest2@tftest.com",
        ///         Webhook = new Volcengine.Vmp.Inputs.ContactWebhookArgs
        ///         {
        ///             Address = "https://www.acctest2.com",
        ///         },
        ///         LarkBotWebhook = new Volcengine.Vmp.Inputs.ContactLarkBotWebhookArgs
        ///         {
        ///             Address = "https://www.acctest2.com",
        ///         },
        ///         DingTalkBotWebhook = new Volcengine.Vmp.Inputs.ContactDingTalkBotWebhookArgs
        ///         {
        ///             Address = "https://www.dingacctest2.com",
        ///             AtMobiles = new[]
        ///             {
        ///                 "18046891813",
        ///             },
        ///         },
        ///         PhoneNumber = new Volcengine.Vmp.Inputs.ContactPhoneNumberArgs
        ///         {
        ///             CountryCode = "+86",
        ///             Number = "18310101011",
        ///         },
        ///     });
        /// 
        ///     var fooContactGroup = new Volcengine.Vmp.ContactGroup("fooContactGroup", new()
        ///     {
        ///         ContactIds = new[]
        ///         {
        ///             fooContact.Id,
        ///             foo1.Id,
        ///         },
        ///     });
        /// 
        ///     var fooContactGroups = Volcengine.Vmp.GetContactGroups.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             fooContactGroup.Id,
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<ContactGroupsResult> InvokeAsync(ContactGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ContactGroupsResult>("volcengine:vmp/contactGroups:ContactGroups", args ?? new ContactGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of vmp contact groups
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var fooContact = new Volcengine.Vmp.Contact("fooContact", new()
        ///     {
        ///         Email = "acctest1@tftest.com",
        ///         Webhook = new Volcengine.Vmp.Inputs.ContactWebhookArgs
        ///         {
        ///             Address = "https://www.acctest1.com",
        ///         },
        ///         LarkBotWebhook = new Volcengine.Vmp.Inputs.ContactLarkBotWebhookArgs
        ///         {
        ///             Address = "https://www.acctest1.com",
        ///         },
        ///         DingTalkBotWebhook = new Volcengine.Vmp.Inputs.ContactDingTalkBotWebhookArgs
        ///         {
        ///             Address = "https://www.dingacctest1.com",
        ///             AtMobiles = new[]
        ///             {
        ///                 "18046891812",
        ///             },
        ///         },
        ///         PhoneNumber = new Volcengine.Vmp.Inputs.ContactPhoneNumberArgs
        ///         {
        ///             CountryCode = "+86",
        ///             Number = "18310101010",
        ///         },
        ///     });
        /// 
        ///     var foo1 = new Volcengine.Vmp.Contact("foo1", new()
        ///     {
        ///         Email = "acctest2@tftest.com",
        ///         Webhook = new Volcengine.Vmp.Inputs.ContactWebhookArgs
        ///         {
        ///             Address = "https://www.acctest2.com",
        ///         },
        ///         LarkBotWebhook = new Volcengine.Vmp.Inputs.ContactLarkBotWebhookArgs
        ///         {
        ///             Address = "https://www.acctest2.com",
        ///         },
        ///         DingTalkBotWebhook = new Volcengine.Vmp.Inputs.ContactDingTalkBotWebhookArgs
        ///         {
        ///             Address = "https://www.dingacctest2.com",
        ///             AtMobiles = new[]
        ///             {
        ///                 "18046891813",
        ///             },
        ///         },
        ///         PhoneNumber = new Volcengine.Vmp.Inputs.ContactPhoneNumberArgs
        ///         {
        ///             CountryCode = "+86",
        ///             Number = "18310101011",
        ///         },
        ///     });
        /// 
        ///     var fooContactGroup = new Volcengine.Vmp.ContactGroup("fooContactGroup", new()
        ///     {
        ///         ContactIds = new[]
        ///         {
        ///             fooContact.Id,
        ///             foo1.Id,
        ///         },
        ///     });
        /// 
        ///     var fooContactGroups = Volcengine.Vmp.GetContactGroups.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             fooContactGroup.Id,
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<ContactGroupsResult> Invoke(ContactGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ContactGroupsResult>("volcengine:vmp/contactGroups:ContactGroups", args ?? new ContactGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ContactGroupsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of contact group ids.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of contact group.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public ContactGroupsArgs()
        {
        }
        public static new ContactGroupsArgs Empty => new ContactGroupsArgs();
    }

    public sealed class ContactGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of contact group ids.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of contact group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public ContactGroupsInvokeArgs()
        {
        }
        public static new ContactGroupsInvokeArgs Empty => new ContactGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class ContactGroupsResult
    {
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContactGroupsContactGroupResult> ContactGroups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// The name of contact group.
        /// </summary>
        public readonly string? Name;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private ContactGroupsResult(
            ImmutableArray<Outputs.ContactGroupsContactGroupResult> contactGroups,

            string id,

            ImmutableArray<string> ids,

            string? name,

            string? outputFile,

            int totalCount)
        {
            ContactGroups = contactGroups;
            Id = id;
            Ids = ids;
            Name = name;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
