// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_mysql
{
    /// <summary>
    /// Provides a resource to manage rds mysql allowlist
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
    ///     var foo = new Volcengine.Rds_mysql.Allowlist("foo", new()
    ///     {
    ///         AllowListDesc = "acc-test",
    ///         AllowListName = "acc-test-allowlist",
    ///         AllowListType = "IPv4",
    ///         SecurityGroupBindInfos = new[]
    ///         {
    ///             new Volcengine.Rds_mysql.Inputs.AllowlistSecurityGroupBindInfoArgs
    ///             {
    ///                 BindMode = "IngressDirectionIp",
    ///                 SecurityGroupId = "sg-13fd7wyduxekg3n6nu5t9fhj7",
    ///             },
    ///             new Volcengine.Rds_mysql.Inputs.AllowlistSecurityGroupBindInfoArgs
    ///             {
    ///                 BindMode = "IngressDirectionIp",
    ///                 SecurityGroupId = "sg-mjoa9qfyzg1s5smt1a6dmc1l",
    ///             },
    ///         },
    ///         UserAllowLists = new[]
    ///         {
    ///             "192.168.0.0/24",
    ///             "192.168.1.0/24",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RDS AllowList can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:rds_mysql/allowlist:Allowlist default acl-d1fd76693bd54e658912e7337d5b****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:rds_mysql/allowlist:Allowlist")]
    public partial class Allowlist : global::Pulumi.CustomResource
    {
        /// <summary>
        /// White list category. Values:
        /// Ordinary: Ordinary white list.
        /// Default: Default white list.
        /// Description: When this parameter is used as a request parameter, the default value is Ordinary.
        /// </summary>
        [Output("allowListCategory")]
        public Output<string> AllowListCategory { get; private set; } = null!;

        /// <summary>
        /// The description of the allow list.
        /// </summary>
        [Output("allowListDesc")]
        public Output<string?> AllowListDesc { get; private set; } = null!;

        /// <summary>
        /// The id of the allow list.
        /// </summary>
        [Output("allowListId")]
        public Output<string> AllowListId { get; private set; } = null!;

        /// <summary>
        /// The name of the allow list.
        /// </summary>
        [Output("allowListName")]
        public Output<string> AllowListName { get; private set; } = null!;

        /// <summary>
        /// The type of IP address in the whitelist. Currently only IPv4 addresses are supported.
        /// </summary>
        [Output("allowListType")]
        public Output<string> AllowListType { get; private set; } = null!;

        /// <summary>
        /// Enter an IP address or a range of IP addresses in CIDR format. Please note that if you want to use security group - related parameters, do not use this field. Instead, use the user_allow_list.
        /// </summary>
        [Output("allowLists")]
        public Output<ImmutableArray<string>> AllowLists { get; private set; } = null!;

        /// <summary>
        /// Whitelist information for the associated security group.
        /// </summary>
        [Output("securityGroupBindInfos")]
        public Output<ImmutableArray<Outputs.AllowlistSecurityGroupBindInfo>> SecurityGroupBindInfos { get; private set; } = null!;

        /// <summary>
        /// The security group ids of the allow list.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// IP addresses outside the security group that need to be added to the whitelist. IP addresses or IP address segments in CIDR format can be entered. Note: This field cannot be used simultaneously with AllowList.
        /// </summary>
        [Output("userAllowLists")]
        public Output<ImmutableArray<string>> UserAllowLists { get; private set; } = null!;


        /// <summary>
        /// Create a Allowlist resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Allowlist(string name, AllowlistArgs args, CustomResourceOptions? options = null)
            : base("volcengine:rds_mysql/allowlist:Allowlist", name, args ?? new AllowlistArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Allowlist(string name, Input<string> id, AllowlistState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:rds_mysql/allowlist:Allowlist", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Allowlist resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Allowlist Get(string name, Input<string> id, AllowlistState? state = null, CustomResourceOptions? options = null)
        {
            return new Allowlist(name, id, state, options);
        }
    }

    public sealed class AllowlistArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// White list category. Values:
        /// Ordinary: Ordinary white list.
        /// Default: Default white list.
        /// Description: When this parameter is used as a request parameter, the default value is Ordinary.
        /// </summary>
        [Input("allowListCategory")]
        public Input<string>? AllowListCategory { get; set; }

        /// <summary>
        /// The description of the allow list.
        /// </summary>
        [Input("allowListDesc")]
        public Input<string>? AllowListDesc { get; set; }

        /// <summary>
        /// The name of the allow list.
        /// </summary>
        [Input("allowListName", required: true)]
        public Input<string> AllowListName { get; set; } = null!;

        /// <summary>
        /// The type of IP address in the whitelist. Currently only IPv4 addresses are supported.
        /// </summary>
        [Input("allowListType")]
        public Input<string>? AllowListType { get; set; }

        [Input("allowLists")]
        private InputList<string>? _allowLists;

        /// <summary>
        /// Enter an IP address or a range of IP addresses in CIDR format. Please note that if you want to use security group - related parameters, do not use this field. Instead, use the user_allow_list.
        /// </summary>
        public InputList<string> AllowLists
        {
            get => _allowLists ?? (_allowLists = new InputList<string>());
            set => _allowLists = value;
        }

        [Input("securityGroupBindInfos")]
        private InputList<Inputs.AllowlistSecurityGroupBindInfoArgs>? _securityGroupBindInfos;

        /// <summary>
        /// Whitelist information for the associated security group.
        /// </summary>
        public InputList<Inputs.AllowlistSecurityGroupBindInfoArgs> SecurityGroupBindInfos
        {
            get => _securityGroupBindInfos ?? (_securityGroupBindInfos = new InputList<Inputs.AllowlistSecurityGroupBindInfoArgs>());
            set => _securityGroupBindInfos = value;
        }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The security group ids of the allow list.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("userAllowLists")]
        private InputList<string>? _userAllowLists;

        /// <summary>
        /// IP addresses outside the security group that need to be added to the whitelist. IP addresses or IP address segments in CIDR format can be entered. Note: This field cannot be used simultaneously with AllowList.
        /// </summary>
        public InputList<string> UserAllowLists
        {
            get => _userAllowLists ?? (_userAllowLists = new InputList<string>());
            set => _userAllowLists = value;
        }

        public AllowlistArgs()
        {
        }
        public static new AllowlistArgs Empty => new AllowlistArgs();
    }

    public sealed class AllowlistState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// White list category. Values:
        /// Ordinary: Ordinary white list.
        /// Default: Default white list.
        /// Description: When this parameter is used as a request parameter, the default value is Ordinary.
        /// </summary>
        [Input("allowListCategory")]
        public Input<string>? AllowListCategory { get; set; }

        /// <summary>
        /// The description of the allow list.
        /// </summary>
        [Input("allowListDesc")]
        public Input<string>? AllowListDesc { get; set; }

        /// <summary>
        /// The id of the allow list.
        /// </summary>
        [Input("allowListId")]
        public Input<string>? AllowListId { get; set; }

        /// <summary>
        /// The name of the allow list.
        /// </summary>
        [Input("allowListName")]
        public Input<string>? AllowListName { get; set; }

        /// <summary>
        /// The type of IP address in the whitelist. Currently only IPv4 addresses are supported.
        /// </summary>
        [Input("allowListType")]
        public Input<string>? AllowListType { get; set; }

        [Input("allowLists")]
        private InputList<string>? _allowLists;

        /// <summary>
        /// Enter an IP address or a range of IP addresses in CIDR format. Please note that if you want to use security group - related parameters, do not use this field. Instead, use the user_allow_list.
        /// </summary>
        public InputList<string> AllowLists
        {
            get => _allowLists ?? (_allowLists = new InputList<string>());
            set => _allowLists = value;
        }

        [Input("securityGroupBindInfos")]
        private InputList<Inputs.AllowlistSecurityGroupBindInfoGetArgs>? _securityGroupBindInfos;

        /// <summary>
        /// Whitelist information for the associated security group.
        /// </summary>
        public InputList<Inputs.AllowlistSecurityGroupBindInfoGetArgs> SecurityGroupBindInfos
        {
            get => _securityGroupBindInfos ?? (_securityGroupBindInfos = new InputList<Inputs.AllowlistSecurityGroupBindInfoGetArgs>());
            set => _securityGroupBindInfos = value;
        }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The security group ids of the allow list.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("userAllowLists")]
        private InputList<string>? _userAllowLists;

        /// <summary>
        /// IP addresses outside the security group that need to be added to the whitelist. IP addresses or IP address segments in CIDR format can be entered. Note: This field cannot be used simultaneously with AllowList.
        /// </summary>
        public InputList<string> UserAllowLists
        {
            get => _userAllowLists ?? (_userAllowLists = new InputList<string>());
            set => _userAllowLists = value;
        }

        public AllowlistState()
        {
        }
        public static new AllowlistState Empty => new AllowlistState();
    }
}
