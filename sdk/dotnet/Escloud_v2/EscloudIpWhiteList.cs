// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Escloud_v2
{
    /// <summary>
    /// Provides a resource to manage escloud ip white list
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
    ///     var fooZones = Volcengine.Ecs.GetZones.Invoke();
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
    ///         Description = "tfdesc",
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         VpcId = fooVpc.Id,
    ///     });
    /// 
    ///     var fooEscloudInstanceV2 = new Volcengine.Escloud_v2.EscloudInstanceV2("fooEscloudInstanceV2", new()
    ///     {
    ///         InstanceName = "acc-test-escloud-instance",
    ///         Version = "V7_10",
    ///         ZoneIds = new[]
    ///         {
    ///             fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///             fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[1]?.Id),
    ///             fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[2]?.Id),
    ///         },
    ///         SubnetId = fooSubnet.Id,
    ///         EnableHttps = false,
    ///         AdminPassword = "Password@@123",
    ///         ChargeType = "PostPaid",
    ///         AutoRenew = false,
    ///         Period = 1,
    ///         ConfigurationCode = "es.standard",
    ///         EnablePureMaster = true,
    ///         DeletionProtection = false,
    ///         ProjectName = "default",
    ///         NodeSpecsAssigns = new[]
    ///         {
    ///             new Volcengine.Escloud_v2.Inputs.EscloudInstanceV2NodeSpecsAssignArgs
    ///             {
    ///                 Type = "Master",
    ///                 Number = 3,
    ///                 ResourceSpecName = "es.x2.medium",
    ///                 StorageSpecName = "es.volume.essd.pl0",
    ///                 StorageSize = 20,
    ///             },
    ///             new Volcengine.Escloud_v2.Inputs.EscloudInstanceV2NodeSpecsAssignArgs
    ///             {
    ///                 Type = "Hot",
    ///                 Number = 6,
    ///                 ResourceSpecName = "es.x2.medium",
    ///                 StorageSpecName = "es.volume.essd.flexpl-standard",
    ///                 StorageSize = 500,
    ///                 ExtraPerformance = new Volcengine.Escloud_v2.Inputs.EscloudInstanceV2NodeSpecsAssignExtraPerformanceArgs
    ///                 {
    ///                     Throughput = 65,
    ///                 },
    ///             },
    ///             new Volcengine.Escloud_v2.Inputs.EscloudInstanceV2NodeSpecsAssignArgs
    ///             {
    ///                 Type = "Kibana",
    ///                 Number = 1,
    ///                 ResourceSpecName = "kibana.x2.small",
    ///                 StorageSpecName = "",
    ///                 StorageSize = 0,
    ///             },
    ///         },
    ///         NetworkSpecs = new[]
    ///         {
    ///             new Volcengine.Escloud_v2.Inputs.EscloudInstanceV2NetworkSpecArgs
    ///             {
    ///                 Type = "Elasticsearch",
    ///                 Bandwidth = 1,
    ///                 IsOpen = true,
    ///                 SpecName = "es.eip.bgp_fixed_bandwidth",
    ///             },
    ///             new Volcengine.Escloud_v2.Inputs.EscloudInstanceV2NetworkSpecArgs
    ///             {
    ///                 Type = "Kibana",
    ///                 Bandwidth = 1,
    ///                 IsOpen = true,
    ///                 SpecName = "es.eip.bgp_fixed_bandwidth",
    ///             },
    ///         },
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Escloud_v2.Inputs.EscloudInstanceV2TagArgs
    ///             {
    ///                 Key = "k1",
    ///                 Value = "v1",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var fooEscloudIpWhiteList = new Volcengine.Escloud_v2.EscloudIpWhiteList("fooEscloudIpWhiteList", new()
    ///     {
    ///         InstanceId = fooEscloudInstanceV2.Id,
    ///         Type = "public",
    ///         Component = "es",
    ///         IpLists = new[]
    ///         {
    ///             "172.16.0.10",
    ///             "172.16.0.11",
    ///             "172.16.0.12",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// EscloudIpWhiteList can be imported using the instance_id:type:component, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:escloud_v2/escloudIpWhiteList:EscloudIpWhiteList default resource_id
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:escloud_v2/escloudIpWhiteList:EscloudIpWhiteList")]
    public partial class EscloudIpWhiteList : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The component of the ip white list. Valid values: `es`, `kibana`.
        /// </summary>
        [Output("component")]
        public Output<string> Component { get; private set; } = null!;

        /// <summary>
        /// The id of the EsCloud instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The ip list of the ip white list.
        /// </summary>
        [Output("ipLists")]
        public Output<ImmutableArray<string>> IpLists { get; private set; } = null!;

        /// <summary>
        /// The type of the ip white list. Valid values: `private`, `public`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a EscloudIpWhiteList resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EscloudIpWhiteList(string name, EscloudIpWhiteListArgs args, CustomResourceOptions? options = null)
            : base("volcengine:escloud_v2/escloudIpWhiteList:EscloudIpWhiteList", name, args ?? new EscloudIpWhiteListArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EscloudIpWhiteList(string name, Input<string> id, EscloudIpWhiteListState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:escloud_v2/escloudIpWhiteList:EscloudIpWhiteList", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EscloudIpWhiteList resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EscloudIpWhiteList Get(string name, Input<string> id, EscloudIpWhiteListState? state = null, CustomResourceOptions? options = null)
        {
            return new EscloudIpWhiteList(name, id, state, options);
        }
    }

    public sealed class EscloudIpWhiteListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The component of the ip white list. Valid values: `es`, `kibana`.
        /// </summary>
        [Input("component", required: true)]
        public Input<string> Component { get; set; } = null!;

        /// <summary>
        /// The id of the EsCloud instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("ipLists", required: true)]
        private InputList<string>? _ipLists;

        /// <summary>
        /// The ip list of the ip white list.
        /// </summary>
        public InputList<string> IpLists
        {
            get => _ipLists ?? (_ipLists = new InputList<string>());
            set => _ipLists = value;
        }

        /// <summary>
        /// The type of the ip white list. Valid values: `private`, `public`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public EscloudIpWhiteListArgs()
        {
        }
        public static new EscloudIpWhiteListArgs Empty => new EscloudIpWhiteListArgs();
    }

    public sealed class EscloudIpWhiteListState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The component of the ip white list. Valid values: `es`, `kibana`.
        /// </summary>
        [Input("component")]
        public Input<string>? Component { get; set; }

        /// <summary>
        /// The id of the EsCloud instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("ipLists")]
        private InputList<string>? _ipLists;

        /// <summary>
        /// The ip list of the ip white list.
        /// </summary>
        public InputList<string> IpLists
        {
            get => _ipLists ?? (_ipLists = new InputList<string>());
            set => _ipLists = value;
        }

        /// <summary>
        /// The type of the ip white list. Valid values: `private`, `public`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public EscloudIpWhiteListState()
        {
        }
        public static new EscloudIpWhiteListState Empty => new EscloudIpWhiteListState();
    }
}
