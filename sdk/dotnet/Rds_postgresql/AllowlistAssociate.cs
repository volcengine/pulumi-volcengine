// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_postgresql
{
    /// <summary>
    /// Provides a resource to manage rds postgresql allowlist associate
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
    ///     var fooZones = Volcengine.Ecs.Zones.Invoke();
    /// 
    ///     var fooVpc = new Volcengine.Vpc.Vpc("fooVpc", new()
    ///     {
    ///         VpcName = "acc-test-vpc",
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var fooSubnet = new Volcengine.Vpc.Subnet("fooSubnet", new()
    ///     {
    ///         SubnetName = "acc-test-subnet",
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[0]?.Id),
    ///         VpcId = fooVpc.Id,
    ///     });
    /// 
    ///     var fooInstance = new Volcengine.Rds_postgresql.Instance("fooInstance", new()
    ///     {
    ///         DbEngineVersion = "PostgreSQL_12",
    ///         NodeSpec = "rds.postgres.1c2g",
    ///         PrimaryZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[0]?.Id),
    ///         SecondaryZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[0]?.Id),
    ///         StorageSpace = 40,
    ///         SubnetId = fooSubnet.Id,
    ///         InstanceName = "acc-test-postgresql",
    ///         ChargeInfo = new Volcengine.Rds_postgresql.Inputs.InstanceChargeInfoArgs
    ///         {
    ///             ChargeType = "PostPaid",
    ///         },
    ///         ProjectName = "default",
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Rds_postgresql.Inputs.InstanceTagArgs
    ///             {
    ///                 Key = "tfk1",
    ///                 Value = "tfv1",
    ///             },
    ///         },
    ///         Parameters = new[]
    ///         {
    ///             new Volcengine.Rds_postgresql.Inputs.InstanceParameterArgs
    ///             {
    ///                 Name = "auto_explain.log_analyze",
    ///                 Value = "off",
    ///             },
    ///             new Volcengine.Rds_postgresql.Inputs.InstanceParameterArgs
    ///             {
    ///                 Name = "auto_explain.log_format",
    ///                 Value = "text",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var fooAllowlist = new Volcengine.Rds_postgresql.Allowlist("fooAllowlist", new()
    ///     {
    ///         AllowListName = "acc-test-allowlist",
    ///         AllowListDesc = "acc-test",
    ///         AllowListType = "IPv4",
    ///         AllowLists = new[]
    ///         {
    ///             "192.168.0.0/24",
    ///             "192.168.1.0/24",
    ///         },
    ///     });
    /// 
    ///     var fooAllowlistAssociate = new Volcengine.Rds_postgresql.AllowlistAssociate("fooAllowlistAssociate", new()
    ///     {
    ///         InstanceId = fooInstance.Id,
    ///         AllowListId = fooAllowlist.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RdsPostgresqlAllowlistAssociate can be imported using the instance_id:allow_list_id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:rds_postgresql/allowlistAssociate:AllowlistAssociate default resource_id
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:rds_postgresql/allowlistAssociate:AllowlistAssociate")]
    public partial class AllowlistAssociate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The id of the postgresql allow list.
        /// </summary>
        [Output("allowListId")]
        public Output<string> AllowListId { get; private set; } = null!;

        /// <summary>
        /// The id of the postgresql instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a AllowlistAssociate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AllowlistAssociate(string name, AllowlistAssociateArgs args, CustomResourceOptions? options = null)
            : base("volcengine:rds_postgresql/allowlistAssociate:AllowlistAssociate", name, args ?? new AllowlistAssociateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AllowlistAssociate(string name, Input<string> id, AllowlistAssociateState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:rds_postgresql/allowlistAssociate:AllowlistAssociate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AllowlistAssociate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AllowlistAssociate Get(string name, Input<string> id, AllowlistAssociateState? state = null, CustomResourceOptions? options = null)
        {
            return new AllowlistAssociate(name, id, state, options);
        }
    }

    public sealed class AllowlistAssociateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the postgresql allow list.
        /// </summary>
        [Input("allowListId", required: true)]
        public Input<string> AllowListId { get; set; } = null!;

        /// <summary>
        /// The id of the postgresql instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public AllowlistAssociateArgs()
        {
        }
        public static new AllowlistAssociateArgs Empty => new AllowlistAssociateArgs();
    }

    public sealed class AllowlistAssociateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the postgresql allow list.
        /// </summary>
        [Input("allowListId")]
        public Input<string>? AllowListId { get; set; }

        /// <summary>
        /// The id of the postgresql instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public AllowlistAssociateState()
        {
        }
        public static new AllowlistAssociateState Empty => new AllowlistAssociateState();
    }
}