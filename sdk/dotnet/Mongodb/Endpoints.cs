// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Mongodb
{
    public static class Endpoints
    {
        /// <summary>
        /// Use this data source to query detailed information of mongodb endpoints
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
        ///     var fooInstance = new Volcengine.Mongodb.Instance("fooInstance", new()
        ///     {
        ///         DbEngineVersion = "MongoDB_4_0",
        ///         InstanceType = "ShardedCluster",
        ///         SuperAccountPassword = "@acc-test-123",
        ///         NodeSpec = "mongo.shard.1c2g",
        ///         MongosNodeSpec = "mongo.mongos.1c2g",
        ///         InstanceName = "acc-test-mongo-shard",
        ///         ChargeType = "PostPaid",
        ///         ProjectName = "default",
        ///         MongosNodeNumber = 2,
        ///         ShardNumber = 2,
        ///         StorageSpaceGb = 20,
        ///         SubnetId = fooSubnet.Id,
        ///         ZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[0]?.Id),
        ///         Tags = new[]
        ///         {
        ///             new Volcengine.Mongodb.Inputs.InstanceTagArgs
        ///             {
        ///                 Key = "k1",
        ///                 Value = "v1",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooAddress = new List&lt;Volcengine.Eip.Address&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; 2; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         fooAddress.Add(new Volcengine.Eip.Address($"fooAddress-{range.Value}", new()
        ///         {
        ///             BillingType = "PostPaidByBandwidth",
        ///             Bandwidth = 1,
        ///             Isp = "ChinaUnicom",
        ///             Description = "acc-test",
        ///             ProjectName = "default",
        ///         }));
        ///     }
        ///     var fooPublic = new Volcengine.Mongodb.Endpoint("fooPublic", new()
        ///     {
        ///         InstanceId = fooInstance.Id,
        ///         NetworkType = "Public",
        ///         ObjectId = fooInstance.MongosId,
        ///         MongosNodeIds = new[]
        ///         {
        ///             fooInstance.Mongos.Apply(mongos =&gt; mongos[0].MongosNodeId),
        ///             fooInstance.Mongos.Apply(mongos =&gt; mongos[1].MongosNodeId),
        ///         },
        ///         EipIds = fooAddress.Select(__item =&gt; __item.Id).ToList(),
        ///     });
        /// 
        ///     var fooPrivate = new Volcengine.Mongodb.Endpoint("fooPrivate", new()
        ///     {
        ///         InstanceId = fooInstance.Id,
        ///         NetworkType = "Private",
        ///         ObjectId = fooInstance.ConfigServersId,
        ///     });
        /// 
        ///     var fooEndpoints = Volcengine.Mongodb.Endpoints.Invoke(new()
        ///     {
        ///         InstanceId = fooInstance.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<EndpointsResult> InvokeAsync(EndpointsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<EndpointsResult>("volcengine:mongodb/endpoints:Endpoints", args ?? new EndpointsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of mongodb endpoints
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
        ///     var fooInstance = new Volcengine.Mongodb.Instance("fooInstance", new()
        ///     {
        ///         DbEngineVersion = "MongoDB_4_0",
        ///         InstanceType = "ShardedCluster",
        ///         SuperAccountPassword = "@acc-test-123",
        ///         NodeSpec = "mongo.shard.1c2g",
        ///         MongosNodeSpec = "mongo.mongos.1c2g",
        ///         InstanceName = "acc-test-mongo-shard",
        ///         ChargeType = "PostPaid",
        ///         ProjectName = "default",
        ///         MongosNodeNumber = 2,
        ///         ShardNumber = 2,
        ///         StorageSpaceGb = 20,
        ///         SubnetId = fooSubnet.Id,
        ///         ZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[0]?.Id),
        ///         Tags = new[]
        ///         {
        ///             new Volcengine.Mongodb.Inputs.InstanceTagArgs
        ///             {
        ///                 Key = "k1",
        ///                 Value = "v1",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooAddress = new List&lt;Volcengine.Eip.Address&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; 2; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         fooAddress.Add(new Volcengine.Eip.Address($"fooAddress-{range.Value}", new()
        ///         {
        ///             BillingType = "PostPaidByBandwidth",
        ///             Bandwidth = 1,
        ///             Isp = "ChinaUnicom",
        ///             Description = "acc-test",
        ///             ProjectName = "default",
        ///         }));
        ///     }
        ///     var fooPublic = new Volcengine.Mongodb.Endpoint("fooPublic", new()
        ///     {
        ///         InstanceId = fooInstance.Id,
        ///         NetworkType = "Public",
        ///         ObjectId = fooInstance.MongosId,
        ///         MongosNodeIds = new[]
        ///         {
        ///             fooInstance.Mongos.Apply(mongos =&gt; mongos[0].MongosNodeId),
        ///             fooInstance.Mongos.Apply(mongos =&gt; mongos[1].MongosNodeId),
        ///         },
        ///         EipIds = fooAddress.Select(__item =&gt; __item.Id).ToList(),
        ///     });
        /// 
        ///     var fooPrivate = new Volcengine.Mongodb.Endpoint("fooPrivate", new()
        ///     {
        ///         InstanceId = fooInstance.Id,
        ///         NetworkType = "Private",
        ///         ObjectId = fooInstance.ConfigServersId,
        ///     });
        /// 
        ///     var fooEndpoints = Volcengine.Mongodb.Endpoints.Invoke(new()
        ///     {
        ///         InstanceId = fooInstance.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<EndpointsResult> Invoke(EndpointsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<EndpointsResult>("volcengine:mongodb/endpoints:Endpoints", args ?? new EndpointsInvokeArgs(), options.WithDefaults());
    }


    public sealed class EndpointsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The instance ID to query.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public EndpointsArgs()
        {
        }
        public static new EndpointsArgs Empty => new EndpointsArgs();
    }

    public sealed class EndpointsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The instance ID to query.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public EndpointsInvokeArgs()
        {
        }
        public static new EndpointsInvokeArgs Empty => new EndpointsInvokeArgs();
    }


    [OutputType]
    public sealed class EndpointsResult
    {
        /// <summary>
        /// The collection of mongodb endpoints query.
        /// </summary>
        public readonly ImmutableArray<Outputs.EndpointsEndpointResult> Endpoints;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? InstanceId;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of mongodb endpoint query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private EndpointsResult(
            ImmutableArray<Outputs.EndpointsEndpointResult> endpoints,

            string id,

            string? instanceId,

            string? outputFile,

            int totalCount)
        {
            Endpoints = endpoints;
            Id = id;
            InstanceId = instanceId;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}