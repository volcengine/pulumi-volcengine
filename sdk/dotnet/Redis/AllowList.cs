// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Redis
{
    /// <summary>
    /// Provides a resource to manage redis allow list
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
    ///     var foo = new Volcengine.Redis.AllowList("foo", new()
    ///     {
    ///         AllowLists = new[]
    ///         {
    ///             "0.0.0.0/0",
    ///             "192.168.0.0/24",
    ///             "192.168.1.1",
    ///             "192.168.2.22",
    ///         },
    ///         AllowListDesc = "acctftestallowlist",
    ///         AllowListName = "acc_test_tf_allowlist_create",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Redis AllowList can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:redis/allowList:AllowList default acl-cn03wk541s55c376xxxx
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:redis/allowList:AllowList")]
    public partial class AllowList : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The type of the whitelist.
        /// </summary>
        [Output("allowListCategory")]
        public Output<string> AllowListCategory { get; private set; } = null!;

        /// <summary>
        /// Description of allow list.
        /// </summary>
        [Output("allowListDesc")]
        public Output<string?> AllowListDesc { get; private set; } = null!;

        /// <summary>
        /// Id of allow list.
        /// </summary>
        [Output("allowListId")]
        public Output<string> AllowListId { get; private set; } = null!;

        /// <summary>
        /// The IP number of allow list.
        /// </summary>
        [Output("allowListIpNum")]
        public Output<int> AllowListIpNum { get; private set; } = null!;

        /// <summary>
        /// Name of allow list.
        /// </summary>
        [Output("allowListName")]
        public Output<string> AllowListName { get; private set; } = null!;

        /// <summary>
        /// Type of allow list.
        /// </summary>
        [Output("allowListType")]
        public Output<string> AllowListType { get; private set; } = null!;

        /// <summary>
        /// Ip list of allow list.
        /// </summary>
        [Output("allowLists")]
        public Output<ImmutableArray<string>> AllowLists { get; private set; } = null!;

        /// <summary>
        /// The number of instance that associated to allow list.
        /// </summary>
        [Output("associatedInstanceNum")]
        public Output<int> AssociatedInstanceNum { get; private set; } = null!;

        /// <summary>
        /// Instances associated by this allow list.
        /// </summary>
        [Output("associatedInstances")]
        public Output<ImmutableArray<Outputs.AllowListAssociatedInstance>> AssociatedInstances { get; private set; } = null!;

        /// <summary>
        /// The name of the project to which the white list belongs.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The current whitelist is the list of security group information that has been associated.
        /// </summary>
        [Output("securityGroupBindInfos")]
        public Output<ImmutableArray<Outputs.AllowListSecurityGroupBindInfo>> SecurityGroupBindInfos { get; private set; } = null!;


        /// <summary>
        /// Create a AllowList resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AllowList(string name, AllowListArgs args, CustomResourceOptions? options = null)
            : base("volcengine:redis/allowList:AllowList", name, args ?? new AllowListArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AllowList(string name, Input<string> id, AllowListState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:redis/allowList:AllowList", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/volcengine",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AllowList resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AllowList Get(string name, Input<string> id, AllowListState? state = null, CustomResourceOptions? options = null)
        {
            return new AllowList(name, id, state, options);
        }
    }

    public sealed class AllowListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of allow list.
        /// </summary>
        [Input("allowListDesc")]
        public Input<string>? AllowListDesc { get; set; }

        /// <summary>
        /// Name of allow list.
        /// </summary>
        [Input("allowListName", required: true)]
        public Input<string> AllowListName { get; set; } = null!;

        [Input("allowLists", required: true)]
        private InputList<string>? _allowLists;

        /// <summary>
        /// Ip list of allow list.
        /// </summary>
        public InputList<string> AllowLists
        {
            get => _allowLists ?? (_allowLists = new InputList<string>());
            set => _allowLists = value;
        }

        public AllowListArgs()
        {
        }
        public static new AllowListArgs Empty => new AllowListArgs();
    }

    public sealed class AllowListState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of the whitelist.
        /// </summary>
        [Input("allowListCategory")]
        public Input<string>? AllowListCategory { get; set; }

        /// <summary>
        /// Description of allow list.
        /// </summary>
        [Input("allowListDesc")]
        public Input<string>? AllowListDesc { get; set; }

        /// <summary>
        /// Id of allow list.
        /// </summary>
        [Input("allowListId")]
        public Input<string>? AllowListId { get; set; }

        /// <summary>
        /// The IP number of allow list.
        /// </summary>
        [Input("allowListIpNum")]
        public Input<int>? AllowListIpNum { get; set; }

        /// <summary>
        /// Name of allow list.
        /// </summary>
        [Input("allowListName")]
        public Input<string>? AllowListName { get; set; }

        /// <summary>
        /// Type of allow list.
        /// </summary>
        [Input("allowListType")]
        public Input<string>? AllowListType { get; set; }

        [Input("allowLists")]
        private InputList<string>? _allowLists;

        /// <summary>
        /// Ip list of allow list.
        /// </summary>
        public InputList<string> AllowLists
        {
            get => _allowLists ?? (_allowLists = new InputList<string>());
            set => _allowLists = value;
        }

        /// <summary>
        /// The number of instance that associated to allow list.
        /// </summary>
        [Input("associatedInstanceNum")]
        public Input<int>? AssociatedInstanceNum { get; set; }

        [Input("associatedInstances")]
        private InputList<Inputs.AllowListAssociatedInstanceGetArgs>? _associatedInstances;

        /// <summary>
        /// Instances associated by this allow list.
        /// </summary>
        public InputList<Inputs.AllowListAssociatedInstanceGetArgs> AssociatedInstances
        {
            get => _associatedInstances ?? (_associatedInstances = new InputList<Inputs.AllowListAssociatedInstanceGetArgs>());
            set => _associatedInstances = value;
        }

        /// <summary>
        /// The name of the project to which the white list belongs.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        [Input("securityGroupBindInfos")]
        private InputList<Inputs.AllowListSecurityGroupBindInfoGetArgs>? _securityGroupBindInfos;

        /// <summary>
        /// The current whitelist is the list of security group information that has been associated.
        /// </summary>
        public InputList<Inputs.AllowListSecurityGroupBindInfoGetArgs> SecurityGroupBindInfos
        {
            get => _securityGroupBindInfos ?? (_securityGroupBindInfos = new InputList<Inputs.AllowListSecurityGroupBindInfoGetArgs>());
            set => _securityGroupBindInfos = value;
        }

        public AllowListState()
        {
        }
        public static new AllowListState Empty => new AllowListState();
    }
}
