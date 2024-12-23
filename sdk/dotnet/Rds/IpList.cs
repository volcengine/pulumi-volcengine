// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds
{
    /// <summary>
    /// (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Provides a resource to manage rds ip list
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
    ///     var foo = new Volcengine.Rds.IpList("foo", new()
    ///     {
    ///         GroupName = "foo",
    ///         InstanceId = "mysql-0fdd3bab2e7c",
    ///         IpLists = new[]
    ///         {
    ///             "1.1.1.1",
    ///             "2.2.2.2",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RDSIPList can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:rds/ipList:IpList default mysql-42b38c769c4b:group_name
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:rds/ipList:IpList")]
    public partial class IpList : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the RDS ip list.
        /// </summary>
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        /// <summary>
        /// The ID of the RDS instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The list of IP address.
        /// </summary>
        [Output("ipLists")]
        public Output<ImmutableArray<string>> IpLists { get; private set; } = null!;


        /// <summary>
        /// Create a IpList resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpList(string name, IpListArgs args, CustomResourceOptions? options = null)
            : base("volcengine:rds/ipList:IpList", name, args ?? new IpListArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpList(string name, Input<string> id, IpListState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:rds/ipList:IpList", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IpList resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpList Get(string name, Input<string> id, IpListState? state = null, CustomResourceOptions? options = null)
        {
            return new IpList(name, id, state, options);
        }
    }

    public sealed class IpListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the RDS ip list.
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        /// <summary>
        /// The ID of the RDS instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("ipLists", required: true)]
        private InputList<string>? _ipLists;

        /// <summary>
        /// The list of IP address.
        /// </summary>
        public InputList<string> IpLists
        {
            get => _ipLists ?? (_ipLists = new InputList<string>());
            set => _ipLists = value;
        }

        public IpListArgs()
        {
        }
        public static new IpListArgs Empty => new IpListArgs();
    }

    public sealed class IpListState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the RDS ip list.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// The ID of the RDS instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("ipLists")]
        private InputList<string>? _ipLists;

        /// <summary>
        /// The list of IP address.
        /// </summary>
        public InputList<string> IpLists
        {
            get => _ipLists ?? (_ipLists = new InputList<string>());
            set => _ipLists = value;
        }

        public IpListState()
        {
        }
        public static new IpListState Empty => new IpListState();
    }
}
