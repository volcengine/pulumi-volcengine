// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpc
{
    public static class GetPrefixLists
    {
        /// <summary>
        /// Use this data source to query detailed information of vpc prefix lists
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
        ///     var fooPrefixList = new Volcengine.Vpc.PrefixList("fooPrefixList", new()
        ///     {
        ///         PrefixListName = "acc-test-prefix",
        ///         MaxEntries = 3,
        ///         Description = "acc test description",
        ///         IpVersion = "IPv4",
        ///         PrefixListEntries = new[]
        ///         {
        ///             new Volcengine.Vpc.Inputs.PrefixListPrefixListEntryArgs
        ///             {
        ///                 Cidr = "192.168.4.0/28",
        ///                 Description = "acc-test-1",
        ///             },
        ///             new Volcengine.Vpc.Inputs.PrefixListPrefixListEntryArgs
        ///             {
        ///                 Cidr = "192.168.5.0/28",
        ///                 Description = "acc-test-2",
        ///             },
        ///         },
        ///         Tags = new[]
        ///         {
        ///             new Volcengine.Vpc.Inputs.PrefixListTagArgs
        ///             {
        ///                 Key = "tf-key1",
        ///                 Value = "tf-value1",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooPrefixLists = Volcengine.Vpc.GetPrefixLists.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             fooPrefixList.Id,
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetPrefixListsResult> InvokeAsync(GetPrefixListsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPrefixListsResult>("volcengine:vpc/getPrefixLists:getPrefixLists", args ?? new GetPrefixListsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of vpc prefix lists
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
        ///     var fooPrefixList = new Volcengine.Vpc.PrefixList("fooPrefixList", new()
        ///     {
        ///         PrefixListName = "acc-test-prefix",
        ///         MaxEntries = 3,
        ///         Description = "acc test description",
        ///         IpVersion = "IPv4",
        ///         PrefixListEntries = new[]
        ///         {
        ///             new Volcengine.Vpc.Inputs.PrefixListPrefixListEntryArgs
        ///             {
        ///                 Cidr = "192.168.4.0/28",
        ///                 Description = "acc-test-1",
        ///             },
        ///             new Volcengine.Vpc.Inputs.PrefixListPrefixListEntryArgs
        ///             {
        ///                 Cidr = "192.168.5.0/28",
        ///                 Description = "acc-test-2",
        ///             },
        ///         },
        ///         Tags = new[]
        ///         {
        ///             new Volcengine.Vpc.Inputs.PrefixListTagArgs
        ///             {
        ///                 Key = "tf-key1",
        ///                 Value = "tf-value1",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooPrefixLists = Volcengine.Vpc.GetPrefixLists.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             fooPrefixList.Id,
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetPrefixListsResult> Invoke(GetPrefixListsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrefixListsResult>("volcengine:vpc/getPrefixLists:getPrefixLists", args ?? new GetPrefixListsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPrefixListsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of prefix list ids.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// IP version of prefix list.
        /// </summary>
        [Input("ipVersion")]
        public string? IpVersion { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// A Name of prefix list.
        /// </summary>
        [Input("prefixListName")]
        public string? PrefixListName { get; set; }

        [Input("tagFilters")]
        private List<Inputs.GetPrefixListsTagFilterArgs>? _tagFilters;

        /// <summary>
        /// List of tag filters.
        /// </summary>
        public List<Inputs.GetPrefixListsTagFilterArgs> TagFilters
        {
            get => _tagFilters ?? (_tagFilters = new List<Inputs.GetPrefixListsTagFilterArgs>());
            set => _tagFilters = value;
        }

        public GetPrefixListsArgs()
        {
        }
        public static new GetPrefixListsArgs Empty => new GetPrefixListsArgs();
    }

    public sealed class GetPrefixListsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of prefix list ids.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// IP version of prefix list.
        /// </summary>
        [Input("ipVersion")]
        public Input<string>? IpVersion { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// A Name of prefix list.
        /// </summary>
        [Input("prefixListName")]
        public Input<string>? PrefixListName { get; set; }

        [Input("tagFilters")]
        private InputList<Inputs.GetPrefixListsTagFilterInputArgs>? _tagFilters;

        /// <summary>
        /// List of tag filters.
        /// </summary>
        public InputList<Inputs.GetPrefixListsTagFilterInputArgs> TagFilters
        {
            get => _tagFilters ?? (_tagFilters = new InputList<Inputs.GetPrefixListsTagFilterInputArgs>());
            set => _tagFilters = value;
        }

        public GetPrefixListsInvokeArgs()
        {
        }
        public static new GetPrefixListsInvokeArgs Empty => new GetPrefixListsInvokeArgs();
    }


    [OutputType]
    public sealed class GetPrefixListsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// The ip version of the prefix list.
        /// </summary>
        public readonly string? IpVersion;
        public readonly string? OutputFile;
        /// <summary>
        /// The prefix list name.
        /// </summary>
        public readonly string? PrefixListName;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPrefixListsPrefixListResult> PrefixLists;
        public readonly ImmutableArray<Outputs.GetPrefixListsTagFilterResult> TagFilters;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetPrefixListsResult(
            string id,

            ImmutableArray<string> ids,

            string? ipVersion,

            string? outputFile,

            string? prefixListName,

            ImmutableArray<Outputs.GetPrefixListsPrefixListResult> prefixLists,

            ImmutableArray<Outputs.GetPrefixListsTagFilterResult> tagFilters,

            int totalCount)
        {
            Id = id;
            Ids = ids;
            IpVersion = ipVersion;
            OutputFile = outputFile;
            PrefixListName = prefixListName;
            PrefixLists = prefixLists;
            TagFilters = tagFilters;
            TotalCount = totalCount;
        }
    }
}
