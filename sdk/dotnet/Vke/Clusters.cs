// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vke
{
    public static class Clusters
    {
        /// <summary>
        /// Use this data source to query detailed information of vke clusters
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
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
        ///         ZoneId = "cn-beijing-a",
        ///         VpcId = fooVpc.Id,
        ///     });
        /// 
        ///     var fooSecurityGroup = new Volcengine.Vpc.SecurityGroup("fooSecurityGroup", new()
        ///     {
        ///         VpcId = fooVpc.Id,
        ///         SecurityGroupName = "acc-test-security-group2",
        ///     });
        /// 
        ///     var fooCluster = new Volcengine.Vke.Cluster("fooCluster", new()
        ///     {
        ///         Description = "created by terraform",
        ///         DeleteProtectionEnabled = false,
        ///         ClusterConfig = new Volcengine.Vke.Inputs.ClusterClusterConfigArgs
        ///         {
        ///             SubnetIds = new[]
        ///             {
        ///                 fooSubnet.Id,
        ///             },
        ///             ApiServerPublicAccessEnabled = true,
        ///             ApiServerPublicAccessConfig = new Volcengine.Vke.Inputs.ClusterClusterConfigApiServerPublicAccessConfigArgs
        ///             {
        ///                 PublicAccessNetworkConfig = new Volcengine.Vke.Inputs.ClusterClusterConfigApiServerPublicAccessConfigPublicAccessNetworkConfigArgs
        ///                 {
        ///                     BillingType = "PostPaidByBandwidth",
        ///                     Bandwidth = 1,
        ///                 },
        ///             },
        ///             ResourcePublicAccessDefaultEnabled = true,
        ///         },
        ///         PodsConfig = new Volcengine.Vke.Inputs.ClusterPodsConfigArgs
        ///         {
        ///             PodNetworkMode = "VpcCniShared",
        ///             VpcCniConfig = new Volcengine.Vke.Inputs.ClusterPodsConfigVpcCniConfigArgs
        ///             {
        ///                 SubnetIds = new[]
        ///                 {
        ///                     fooSubnet.Id,
        ///                 },
        ///             },
        ///         },
        ///         ServicesConfig = new Volcengine.Vke.Inputs.ClusterServicesConfigArgs
        ///         {
        ///             ServiceCidrsv4s = new[]
        ///             {
        ///                 "172.30.0.0/18",
        ///             },
        ///         },
        ///         Tags = new[]
        ///         {
        ///             new Volcengine.Vke.Inputs.ClusterTagArgs
        ///             {
        ///                 Key = "tf-k1",
        ///                 Value = "tf-v1",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooClusters = Volcengine.Vke.Clusters.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             fooCluster.Id,
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<ClustersResult> InvokeAsync(ClustersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ClustersResult>("volcengine:vke/clusters:Clusters", args ?? new ClustersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of vke clusters
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
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
        ///         ZoneId = "cn-beijing-a",
        ///         VpcId = fooVpc.Id,
        ///     });
        /// 
        ///     var fooSecurityGroup = new Volcengine.Vpc.SecurityGroup("fooSecurityGroup", new()
        ///     {
        ///         VpcId = fooVpc.Id,
        ///         SecurityGroupName = "acc-test-security-group2",
        ///     });
        /// 
        ///     var fooCluster = new Volcengine.Vke.Cluster("fooCluster", new()
        ///     {
        ///         Description = "created by terraform",
        ///         DeleteProtectionEnabled = false,
        ///         ClusterConfig = new Volcengine.Vke.Inputs.ClusterClusterConfigArgs
        ///         {
        ///             SubnetIds = new[]
        ///             {
        ///                 fooSubnet.Id,
        ///             },
        ///             ApiServerPublicAccessEnabled = true,
        ///             ApiServerPublicAccessConfig = new Volcengine.Vke.Inputs.ClusterClusterConfigApiServerPublicAccessConfigArgs
        ///             {
        ///                 PublicAccessNetworkConfig = new Volcengine.Vke.Inputs.ClusterClusterConfigApiServerPublicAccessConfigPublicAccessNetworkConfigArgs
        ///                 {
        ///                     BillingType = "PostPaidByBandwidth",
        ///                     Bandwidth = 1,
        ///                 },
        ///             },
        ///             ResourcePublicAccessDefaultEnabled = true,
        ///         },
        ///         PodsConfig = new Volcengine.Vke.Inputs.ClusterPodsConfigArgs
        ///         {
        ///             PodNetworkMode = "VpcCniShared",
        ///             VpcCniConfig = new Volcengine.Vke.Inputs.ClusterPodsConfigVpcCniConfigArgs
        ///             {
        ///                 SubnetIds = new[]
        ///                 {
        ///                     fooSubnet.Id,
        ///                 },
        ///             },
        ///         },
        ///         ServicesConfig = new Volcengine.Vke.Inputs.ClusterServicesConfigArgs
        ///         {
        ///             ServiceCidrsv4s = new[]
        ///             {
        ///                 "172.30.0.0/18",
        ///             },
        ///         },
        ///         Tags = new[]
        ///         {
        ///             new Volcengine.Vke.Inputs.ClusterTagArgs
        ///             {
        ///                 Key = "tf-k1",
        ///                 Value = "tf-v1",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooClusters = Volcengine.Vke.Clusters.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             fooCluster.Id,
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<ClustersResult> Invoke(ClustersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ClustersResult>("volcengine:vke/clusters:Clusters", args ?? new ClustersInvokeArgs(), options.WithDefaults());
    }


    public sealed class ClustersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ClientToken when the cluster is created successfully. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
        /// </summary>
        [Input("createClientToken")]
        public string? CreateClientToken { get; set; }

        /// <summary>
        /// The delete protection of the cluster, the value is `true` or `false`.
        /// </summary>
        [Input("deleteProtectionEnabled")]
        public bool? DeleteProtectionEnabled { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Cluster IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of the cluster.
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
        /// The page number of clusters query.
        /// </summary>
        [Input("pageNumber")]
        public int? PageNumber { get; set; }

        /// <summary>
        /// The page size of clusters query.
        /// </summary>
        [Input("pageSize")]
        public int? PageSize { get; set; }

        /// <summary>
        /// The container network model of the cluster, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.
        /// </summary>
        [Input("podsConfigPodNetworkMode")]
        public string? PodsConfigPodNetworkMode { get; set; }

        [Input("statuses")]
        private List<Inputs.ClustersStatusArgs>? _statuses;

        /// <summary>
        /// Array of cluster states to filter. (The elements of the array are logically ORed. A maximum of 15 state array elements can be filled at a time).
        /// </summary>
        public List<Inputs.ClustersStatusArgs> Statuses
        {
            get => _statuses ?? (_statuses = new List<Inputs.ClustersStatusArgs>());
            set => _statuses = value;
        }

        [Input("tags")]
        private List<Inputs.ClustersTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public List<Inputs.ClustersTagArgs> Tags
        {
            get => _tags ?? (_tags = new List<Inputs.ClustersTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The ClientToken when the last cluster update succeeded. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
        /// </summary>
        [Input("updateClientToken")]
        public string? UpdateClientToken { get; set; }

        public ClustersArgs()
        {
        }
        public static new ClustersArgs Empty => new ClustersArgs();
    }

    public sealed class ClustersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ClientToken when the cluster is created successfully. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
        /// </summary>
        [Input("createClientToken")]
        public Input<string>? CreateClientToken { get; set; }

        /// <summary>
        /// The delete protection of the cluster, the value is `true` or `false`.
        /// </summary>
        [Input("deleteProtectionEnabled")]
        public Input<bool>? DeleteProtectionEnabled { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Cluster IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of the cluster.
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
        /// The page number of clusters query.
        /// </summary>
        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        /// <summary>
        /// The page size of clusters query.
        /// </summary>
        [Input("pageSize")]
        public Input<int>? PageSize { get; set; }

        /// <summary>
        /// The container network model of the cluster, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.
        /// </summary>
        [Input("podsConfigPodNetworkMode")]
        public Input<string>? PodsConfigPodNetworkMode { get; set; }

        [Input("statuses")]
        private InputList<Inputs.ClustersStatusInputArgs>? _statuses;

        /// <summary>
        /// Array of cluster states to filter. (The elements of the array are logically ORed. A maximum of 15 state array elements can be filled at a time).
        /// </summary>
        public InputList<Inputs.ClustersStatusInputArgs> Statuses
        {
            get => _statuses ?? (_statuses = new InputList<Inputs.ClustersStatusInputArgs>());
            set => _statuses = value;
        }

        [Input("tags")]
        private InputList<Inputs.ClustersTagInputArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.ClustersTagInputArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ClustersTagInputArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The ClientToken when the last cluster update succeeded. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
        /// </summary>
        [Input("updateClientToken")]
        public Input<string>? UpdateClientToken { get; set; }

        public ClustersInvokeArgs()
        {
        }
        public static new ClustersInvokeArgs Empty => new ClustersInvokeArgs();
    }


    [OutputType]
    public sealed class ClustersResult
    {
        /// <summary>
        /// The collection of VkeCluster query.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClustersClusterResult> Clusters;
        public readonly string? CreateClientToken;
        /// <summary>
        /// The delete protection of the cluster, the value is `true` or `false`.
        /// </summary>
        public readonly bool? DeleteProtectionEnabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// The name of the cluster.
        /// </summary>
        public readonly string? Name;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        public readonly int PageNumber;
        public readonly int PageSize;
        public readonly string? PodsConfigPodNetworkMode;
        public readonly ImmutableArray<Outputs.ClustersStatusResult> Statuses;
        /// <summary>
        /// Tags of the Cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClustersTagResult> Tags;
        /// <summary>
        /// The total count of Cluster query.
        /// </summary>
        public readonly int TotalCount;
        public readonly string? UpdateClientToken;

        [OutputConstructor]
        private ClustersResult(
            ImmutableArray<Outputs.ClustersClusterResult> clusters,

            string? createClientToken,

            bool? deleteProtectionEnabled,

            string id,

            ImmutableArray<string> ids,

            string? name,

            string? nameRegex,

            string? outputFile,

            int pageNumber,

            int pageSize,

            string? podsConfigPodNetworkMode,

            ImmutableArray<Outputs.ClustersStatusResult> statuses,

            ImmutableArray<Outputs.ClustersTagResult> tags,

            int totalCount,

            string? updateClientToken)
        {
            Clusters = clusters;
            CreateClientToken = createClientToken;
            DeleteProtectionEnabled = deleteProtectionEnabled;
            Id = id;
            Ids = ids;
            Name = name;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            PageNumber = pageNumber;
            PageSize = pageSize;
            PodsConfigPodNetworkMode = podsConfigPodNetworkMode;
            Statuses = statuses;
            Tags = tags;
            TotalCount = totalCount;
            UpdateClientToken = updateClientToken;
        }
    }
}