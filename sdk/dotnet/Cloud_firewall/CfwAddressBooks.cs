// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cloud_firewall
{
    [Obsolete(@"volcengine.cloud_firewall.CfwAddressBooks has been deprecated in favor of volcengine.cloud_firewall.getCfwAddressBooks")]
    public static class CfwAddressBooks
    {
        /// <summary>
        /// Use this data source to query detailed information of cfw address books
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
        ///     var foo = Volcengine.Cloud_firewall.GetCfwAddressBooks.Invoke(new()
        ///     {
        ///         GroupName = "acc-test",
        ///         GroupType = "ip",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<CfwAddressBooksResult> InvokeAsync(CfwAddressBooksArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<CfwAddressBooksResult>("volcengine:cloud_firewall/cfwAddressBooks:CfwAddressBooks", args ?? new CfwAddressBooksArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cfw address books
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
        ///     var foo = Volcengine.Cloud_firewall.GetCfwAddressBooks.Invoke(new()
        ///     {
        ///         GroupName = "acc-test",
        ///         GroupType = "ip",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<CfwAddressBooksResult> Invoke(CfwAddressBooksInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<CfwAddressBooksResult>("volcengine:cloud_firewall/cfwAddressBooks:CfwAddressBooks", args ?? new CfwAddressBooksInvokeArgs(), options.WithDefaults());
    }


    public sealed class CfwAddressBooksArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The group type of address book. This field support fuzzy query.
        /// </summary>
        [Input("address")]
        public string? Address { get; set; }

        /// <summary>
        /// The group type of address book. This field support fuzzy query.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The group name of address book. This field support fuzzy query.
        /// </summary>
        [Input("groupName")]
        public string? GroupName { get; set; }

        /// <summary>
        /// The group type of address book. Valid values: `ip`, `port`, `domain`.
        /// </summary>
        [Input("groupType")]
        public string? GroupType { get; set; }

        /// <summary>
        /// A Name Regex of Resource.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public CfwAddressBooksArgs()
        {
        }
        public static new CfwAddressBooksArgs Empty => new CfwAddressBooksArgs();
    }

    public sealed class CfwAddressBooksInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The group type of address book. This field support fuzzy query.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The group type of address book. This field support fuzzy query.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The group name of address book. This field support fuzzy query.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// The group type of address book. Valid values: `ip`, `port`, `domain`.
        /// </summary>
        [Input("groupType")]
        public Input<string>? GroupType { get; set; }

        /// <summary>
        /// A Name Regex of Resource.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public CfwAddressBooksInvokeArgs()
        {
        }
        public static new CfwAddressBooksInvokeArgs Empty => new CfwAddressBooksInvokeArgs();
    }


    [OutputType]
    public sealed class CfwAddressBooksResult
    {
        public readonly string? Address;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.CfwAddressBooksAddressBookResult> AddressBooks;
        /// <summary>
        /// The description of the address book.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The name of the address book.
        /// </summary>
        public readonly string? GroupName;
        /// <summary>
        /// The type of the address book.
        /// </summary>
        public readonly string? GroupType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private CfwAddressBooksResult(
            string? address,

            ImmutableArray<Outputs.CfwAddressBooksAddressBookResult> addressBooks,

            string? description,

            string? groupName,

            string? groupType,

            string id,

            string? nameRegex,

            string? outputFile,

            int totalCount)
        {
            Address = address;
            AddressBooks = addressBooks;
            Description = description;
            GroupName = groupName;
            GroupType = groupType;
            Id = id;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
