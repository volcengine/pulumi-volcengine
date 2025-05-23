// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Veecp
{
    public static class GetClusters
    {
        /// <summary>
        /// Use this data source to query detailed information of veecp clusters
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
        ///         VpcName = "acc-test-project1",
        ///         CidrBlock = "172.16.0.0/16",
        ///     });
        /// 
        ///     var fooSubnet = new Volcengine.Vpc.Subnet("fooSubnet", new()
        ///     {
        ///         SubnetName = "acc-subnet-test-2",
        ///         CidrBlock = "172.16.0.0/24",
        ///         ZoneId = fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///         VpcId = fooVpc.Id,
        ///     });
        /// 
        ///     var fooSecurityGroup = new Volcengine.Vpc.SecurityGroup("fooSecurityGroup", new()
        ///     {
        ///         VpcId = fooVpc.Id,
        ///         SecurityGroupName = "acc-test-security-group2",
        ///     });
        /// 
        ///     var fooCluster = new Volcengine.Veecp.Cluster("fooCluster", new()
        ///     {
        ///         Description = "created by terraform",
        ///         DeleteProtectionEnabled = false,
        ///         Profile = "Edge",
        ///         ClusterConfig = new Volcengine.Veecp.Inputs.ClusterClusterConfigArgs
        ///         {
        ///             SubnetIds = new[]
        ///             {
        ///                 fooSubnet.Id,
        ///             },
        ///             ApiServerPublicAccessEnabled = true,
        ///             ApiServerPublicAccessConfig = new Volcengine.Veecp.Inputs.ClusterClusterConfigApiServerPublicAccessConfigArgs
        ///             {
        ///                 PublicAccessNetworkConfig = new Volcengine.Veecp.Inputs.ClusterClusterConfigApiServerPublicAccessConfigPublicAccessNetworkConfigArgs
        ///                 {
        ///                     BillingType = "PostPaidByBandwidth",
        ///                     Bandwidth = 1,
        ///                 },
        ///             },
        ///             ResourcePublicAccessDefaultEnabled = true,
        ///         },
        ///         PodsConfig = new Volcengine.Veecp.Inputs.ClusterPodsConfigArgs
        ///         {
        ///             PodNetworkMode = "Flannel",
        ///             FlannelConfig = new Volcengine.Veecp.Inputs.ClusterPodsConfigFlannelConfigArgs
        ///             {
        ///                 PodCidrs = new[]
        ///                 {
        ///                     "172.22.224.0/20",
        ///                 },
        ///                 MaxPodsPerNode = 64,
        ///             },
        ///         },
        ///         ServicesConfig = new Volcengine.Veecp.Inputs.ClusterServicesConfigArgs
        ///         {
        ///             ServiceCidrsv4s = new[]
        ///             {
        ///                 "172.30.0.0/18",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooClusters = Volcengine.Veecp.GetClusters.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             fooCluster.Id,
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetClustersResult> InvokeAsync(GetClustersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClustersResult>("volcengine:veecp/getClusters:getClusters", args ?? new GetClustersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of veecp clusters
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
        ///         VpcName = "acc-test-project1",
        ///         CidrBlock = "172.16.0.0/16",
        ///     });
        /// 
        ///     var fooSubnet = new Volcengine.Vpc.Subnet("fooSubnet", new()
        ///     {
        ///         SubnetName = "acc-subnet-test-2",
        ///         CidrBlock = "172.16.0.0/24",
        ///         ZoneId = fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///         VpcId = fooVpc.Id,
        ///     });
        /// 
        ///     var fooSecurityGroup = new Volcengine.Vpc.SecurityGroup("fooSecurityGroup", new()
        ///     {
        ///         VpcId = fooVpc.Id,
        ///         SecurityGroupName = "acc-test-security-group2",
        ///     });
        /// 
        ///     var fooCluster = new Volcengine.Veecp.Cluster("fooCluster", new()
        ///     {
        ///         Description = "created by terraform",
        ///         DeleteProtectionEnabled = false,
        ///         Profile = "Edge",
        ///         ClusterConfig = new Volcengine.Veecp.Inputs.ClusterClusterConfigArgs
        ///         {
        ///             SubnetIds = new[]
        ///             {
        ///                 fooSubnet.Id,
        ///             },
        ///             ApiServerPublicAccessEnabled = true,
        ///             ApiServerPublicAccessConfig = new Volcengine.Veecp.Inputs.ClusterClusterConfigApiServerPublicAccessConfigArgs
        ///             {
        ///                 PublicAccessNetworkConfig = new Volcengine.Veecp.Inputs.ClusterClusterConfigApiServerPublicAccessConfigPublicAccessNetworkConfigArgs
        ///                 {
        ///                     BillingType = "PostPaidByBandwidth",
        ///                     Bandwidth = 1,
        ///                 },
        ///             },
        ///             ResourcePublicAccessDefaultEnabled = true,
        ///         },
        ///         PodsConfig = new Volcengine.Veecp.Inputs.ClusterPodsConfigArgs
        ///         {
        ///             PodNetworkMode = "Flannel",
        ///             FlannelConfig = new Volcengine.Veecp.Inputs.ClusterPodsConfigFlannelConfigArgs
        ///             {
        ///                 PodCidrs = new[]
        ///                 {
        ///                     "172.22.224.0/20",
        ///                 },
        ///                 MaxPodsPerNode = 64,
        ///             },
        ///         },
        ///         ServicesConfig = new Volcengine.Veecp.Inputs.ClusterServicesConfigArgs
        ///         {
        ///             ServiceCidrsv4s = new[]
        ///             {
        ///                 "172.30.0.0/18",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooClusters = Volcengine.Veecp.GetClusters.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             fooCluster.Id,
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetClustersResult> Invoke(GetClustersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClustersResult>("volcengine:veecp/getClusters:getClusters", args ?? new GetClustersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClustersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ClientToken when the cluster is created successfully. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
        /// </summary>
        [Input("createClientToken")]
        public string? CreateClientToken { get; set; }

        /// <summary>
        /// Cluster deletion protection. Values: true: Enable deletion protection. false: Disable deletion protection.
        /// </summary>
        [Input("deleteProtectionEnabled")]
        public bool? DeleteProtectionEnabled { get; set; }

        /// <summary>
        /// Whether to enable the edge tunnel. The value is `true` or `false`.
        /// </summary>
        [Input("edgeTunnelEnabled")]
        public bool? EdgeTunnelEnabled { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// Cluster ID. Supports exact matching. A maximum of 100 array elements can be filled in at a time. Note: When this parameter is an empty array, filtering is based on all clusters in the specified region under the account.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Cluster name.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// A Name Regex of Cluster.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The container network model of the cluster, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.
        /// </summary>
        [Input("podsConfigPodNetworkMode")]
        public string? PodsConfigPodNetworkMode { get; set; }

        [Input("profiles")]
        private List<string>? _profiles;

        /// <summary>
        /// Filter by cluster scenario: Cloud: non-edge cluster; Edge: edge cluster.
        /// </summary>
        public List<string> Profiles
        {
            get => _profiles ?? (_profiles = new List<string>());
            set => _profiles = value;
        }

        [Input("statuses")]
        private List<Inputs.GetClustersStatusArgs>? _statuses;

        /// <summary>
        /// Array of cluster states to filter. (The elements of the array are logically ORed. A maximum of 15 state array elements can be filled at a time).
        /// </summary>
        public List<Inputs.GetClustersStatusArgs> Statuses
        {
            get => _statuses ?? (_statuses = new List<Inputs.GetClustersStatusArgs>());
            set => _statuses = value;
        }

        /// <summary>
        /// The ClientToken when the last cluster update succeeded. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
        /// </summary>
        [Input("updateClientToken")]
        public string? UpdateClientToken { get; set; }

        public GetClustersArgs()
        {
        }
        public static new GetClustersArgs Empty => new GetClustersArgs();
    }

    public sealed class GetClustersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ClientToken when the cluster is created successfully. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
        /// </summary>
        [Input("createClientToken")]
        public Input<string>? CreateClientToken { get; set; }

        /// <summary>
        /// Cluster deletion protection. Values: true: Enable deletion protection. false: Disable deletion protection.
        /// </summary>
        [Input("deleteProtectionEnabled")]
        public Input<bool>? DeleteProtectionEnabled { get; set; }

        /// <summary>
        /// Whether to enable the edge tunnel. The value is `true` or `false`.
        /// </summary>
        [Input("edgeTunnelEnabled")]
        public Input<bool>? EdgeTunnelEnabled { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// Cluster ID. Supports exact matching. A maximum of 100 array elements can be filled in at a time. Note: When this parameter is an empty array, filtering is based on all clusters in the specified region under the account.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Cluster name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A Name Regex of Cluster.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The container network model of the cluster, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.
        /// </summary>
        [Input("podsConfigPodNetworkMode")]
        public Input<string>? PodsConfigPodNetworkMode { get; set; }

        [Input("profiles")]
        private InputList<string>? _profiles;

        /// <summary>
        /// Filter by cluster scenario: Cloud: non-edge cluster; Edge: edge cluster.
        /// </summary>
        public InputList<string> Profiles
        {
            get => _profiles ?? (_profiles = new InputList<string>());
            set => _profiles = value;
        }

        [Input("statuses")]
        private InputList<Inputs.GetClustersStatusInputArgs>? _statuses;

        /// <summary>
        /// Array of cluster states to filter. (The elements of the array are logically ORed. A maximum of 15 state array elements can be filled at a time).
        /// </summary>
        public InputList<Inputs.GetClustersStatusInputArgs> Statuses
        {
            get => _statuses ?? (_statuses = new InputList<Inputs.GetClustersStatusInputArgs>());
            set => _statuses = value;
        }

        /// <summary>
        /// The ClientToken when the last cluster update succeeded. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
        /// </summary>
        [Input("updateClientToken")]
        public Input<string>? UpdateClientToken { get; set; }

        public GetClustersInvokeArgs()
        {
        }
        public static new GetClustersInvokeArgs Empty => new GetClustersInvokeArgs();
    }


    [OutputType]
    public sealed class GetClustersResult
    {
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetClustersClusterResult> Clusters;
        /// <summary>
        /// ClientToken when creation is successful. ClientToken is a string that guarantees request idempotency. This string is passed in by the caller.
        /// </summary>
        public readonly string? CreateClientToken;
        /// <summary>
        /// The delete protection of the cluster, the value is `true` or `false`.
        /// </summary>
        public readonly bool? DeleteProtectionEnabled;
        public readonly bool? EdgeTunnelEnabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// Cluster name.
        /// </summary>
        public readonly string? Name;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        public readonly string? PodsConfigPodNetworkMode;
        public readonly ImmutableArray<string> Profiles;
        public readonly ImmutableArray<Outputs.GetClustersStatusResult> Statuses;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// ClientToken when the last update was successful. ClientToken is a string that guarantees request idempotency. This string is passed in by the caller.
        /// </summary>
        public readonly string? UpdateClientToken;

        [OutputConstructor]
        private GetClustersResult(
            ImmutableArray<Outputs.GetClustersClusterResult> clusters,

            string? createClientToken,

            bool? deleteProtectionEnabled,

            bool? edgeTunnelEnabled,

            string id,

            ImmutableArray<string> ids,

            string? name,

            string? nameRegex,

            string? outputFile,

            string? podsConfigPodNetworkMode,

            ImmutableArray<string> profiles,

            ImmutableArray<Outputs.GetClustersStatusResult> statuses,

            int totalCount,

            string? updateClientToken)
        {
            Clusters = clusters;
            CreateClientToken = createClientToken;
            DeleteProtectionEnabled = deleteProtectionEnabled;
            EdgeTunnelEnabled = edgeTunnelEnabled;
            Id = id;
            Ids = ids;
            Name = name;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            PodsConfigPodNetworkMode = podsConfigPodNetworkMode;
            Profiles = profiles;
            Statuses = statuses;
            TotalCount = totalCount;
            UpdateClientToken = updateClientToken;
        }
    }
}
