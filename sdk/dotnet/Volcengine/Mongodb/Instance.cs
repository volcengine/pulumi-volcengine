// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Mongodb
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Volcengine = Volcengine.PulumiPackage.Volcengine;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Volcengine.Mongodb.Instance("foo", new Volcengine.Mongodb.InstanceArgs
    ///         {
    ///             ChargeType = "PostPaid",
    ///             DbEngineVersion = "MongoDB_4_0",
    ///             InstanceName = "mongo-replica-be9995d32e4a",
    ///             InstanceType = "ReplicaSet",
    ///             NodeSpec = "mongo.2c4g",
    ///             ProjectName = "default",
    ///             StorageSpaceGb = 20,
    ///             SubnetId = "subnet-rrx4ns6abw1sv0x57wq6h47",
    ///             SuperAccountPassword = "******",
    ///             Tags = 
    ///             {
    ///                 new Volcengine.Mongodb.Inputs.InstanceTagArgs
    ///                 {
    ///                     Key = "k1",
    ///                     Value = "v1",
    ///                 },
    ///             },
    ///             ZoneId = "cn-beijing-a",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// mongodb instance can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:mongodb/instance:Instance default mongo-replica-e405f8e2****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:mongodb/instance:Instance")]
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to enable automatic renewal.
        /// </summary>
        [Output("autoRenew")]
        public Output<bool> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// The charge type of instance, valid value contains `Prepaid` or `PostPaid`.
        /// </summary>
        [Output("chargeType")]
        public Output<string> ChargeType { get; private set; } = null!;

        /// <summary>
        /// The version of db engine, valid value contains `MongoDB_4_0`, `MongoDB_5_0`.
        /// </summary>
        [Output("dbEngineVersion")]
        public Output<string> DbEngineVersion { get; private set; } = null!;

        /// <summary>
        /// The instance name.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// The type of instance,the valid value contains `ReplicaSet` or `ShardedCluster`.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The mongos node number of shard cluster,value range is `2~23`, this parameter is required when `InstanceType` is `ShardedCluster`.
        /// </summary>
        [Output("mongosNodeNumber")]
        public Output<int?> MongosNodeNumber { get; private set; } = null!;

        /// <summary>
        /// The mongos node spec of shard cluster, this parameter is required when `InstanceType` is `ShardedCluster`.
        /// </summary>
        [Output("mongosNodeSpec")]
        public Output<string?> MongosNodeSpec { get; private set; } = null!;

        /// <summary>
        /// The spec of node.
        /// </summary>
        [Output("nodeSpec")]
        public Output<string> NodeSpec { get; private set; } = null!;

        /// <summary>
        /// The instance purchase duration,the value range is `1~3` when `PeriodUtil` is `Year`, the value range is `1~9` when `PeriodUtil` is `Month`, this parameter is required when `ChargeType` is `Prepaid`.
        /// </summary>
        [Output("period")]
        public Output<int> Period { get; private set; } = null!;

        /// <summary>
        /// The period unit,valid value contains `Year` or `Month`, this parameter is required when `ChargeType` is `Prepaid`.
        /// </summary>
        [Output("periodUnit")]
        public Output<string> PeriodUnit { get; private set; } = null!;

        /// <summary>
        /// The project name to which the instance belongs.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The number of shards in shard cluster,value range is `2~32`, this parameter is required when `InstanceType` is `ShardedCluster`.
        /// </summary>
        [Output("shardNumber")]
        public Output<int?> ShardNumber { get; private set; } = null!;

        /// <summary>
        /// The total storage space of a replica set instance, or the storage space of a single shard in a sharded cluster, in GiB.
        /// </summary>
        [Output("storageSpaceGb")]
        public Output<int> StorageSpaceGb { get; private set; } = null!;

        /// <summary>
        /// The subnet id of instance.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// The password of database account.
        /// </summary>
        [Output("superAccountPassword")]
        public Output<string?> SuperAccountPassword { get; private set; } = null!;

        /// <summary>
        /// Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.InstanceTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The vpc ID.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The zone ID of instance.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("volcengine:mongodb/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:mongodb/instance:Instance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable automatic renewal.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// The charge type of instance, valid value contains `Prepaid` or `PostPaid`.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// The version of db engine, valid value contains `MongoDB_4_0`, `MongoDB_5_0`.
        /// </summary>
        [Input("dbEngineVersion")]
        public Input<string>? DbEngineVersion { get; set; }

        /// <summary>
        /// The instance name.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The type of instance,the valid value contains `ReplicaSet` or `ShardedCluster`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The mongos node number of shard cluster,value range is `2~23`, this parameter is required when `InstanceType` is `ShardedCluster`.
        /// </summary>
        [Input("mongosNodeNumber")]
        public Input<int>? MongosNodeNumber { get; set; }

        /// <summary>
        /// The mongos node spec of shard cluster, this parameter is required when `InstanceType` is `ShardedCluster`.
        /// </summary>
        [Input("mongosNodeSpec")]
        public Input<string>? MongosNodeSpec { get; set; }

        /// <summary>
        /// The spec of node.
        /// </summary>
        [Input("nodeSpec", required: true)]
        public Input<string> NodeSpec { get; set; } = null!;

        /// <summary>
        /// The instance purchase duration,the value range is `1~3` when `PeriodUtil` is `Year`, the value range is `1~9` when `PeriodUtil` is `Month`, this parameter is required when `ChargeType` is `Prepaid`.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The period unit,valid value contains `Year` or `Month`, this parameter is required when `ChargeType` is `Prepaid`.
        /// </summary>
        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        /// <summary>
        /// The project name to which the instance belongs.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The number of shards in shard cluster,value range is `2~32`, this parameter is required when `InstanceType` is `ShardedCluster`.
        /// </summary>
        [Input("shardNumber")]
        public Input<int>? ShardNumber { get; set; }

        /// <summary>
        /// The total storage space of a replica set instance, or the storage space of a single shard in a sharded cluster, in GiB.
        /// </summary>
        [Input("storageSpaceGb", required: true)]
        public Input<int> StorageSpaceGb { get; set; } = null!;

        /// <summary>
        /// The subnet id of instance.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        /// <summary>
        /// The password of database account.
        /// </summary>
        [Input("superAccountPassword")]
        public Input<string>? SuperAccountPassword { get; set; }

        [Input("tags")]
        private InputList<Inputs.InstanceTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.InstanceTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.InstanceTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The vpc ID.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The zone ID of instance.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable automatic renewal.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// The charge type of instance, valid value contains `Prepaid` or `PostPaid`.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// The version of db engine, valid value contains `MongoDB_4_0`, `MongoDB_5_0`.
        /// </summary>
        [Input("dbEngineVersion")]
        public Input<string>? DbEngineVersion { get; set; }

        /// <summary>
        /// The instance name.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The type of instance,the valid value contains `ReplicaSet` or `ShardedCluster`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The mongos node number of shard cluster,value range is `2~23`, this parameter is required when `InstanceType` is `ShardedCluster`.
        /// </summary>
        [Input("mongosNodeNumber")]
        public Input<int>? MongosNodeNumber { get; set; }

        /// <summary>
        /// The mongos node spec of shard cluster, this parameter is required when `InstanceType` is `ShardedCluster`.
        /// </summary>
        [Input("mongosNodeSpec")]
        public Input<string>? MongosNodeSpec { get; set; }

        /// <summary>
        /// The spec of node.
        /// </summary>
        [Input("nodeSpec")]
        public Input<string>? NodeSpec { get; set; }

        /// <summary>
        /// The instance purchase duration,the value range is `1~3` when `PeriodUtil` is `Year`, the value range is `1~9` when `PeriodUtil` is `Month`, this parameter is required when `ChargeType` is `Prepaid`.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The period unit,valid value contains `Year` or `Month`, this parameter is required when `ChargeType` is `Prepaid`.
        /// </summary>
        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        /// <summary>
        /// The project name to which the instance belongs.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The number of shards in shard cluster,value range is `2~32`, this parameter is required when `InstanceType` is `ShardedCluster`.
        /// </summary>
        [Input("shardNumber")]
        public Input<int>? ShardNumber { get; set; }

        /// <summary>
        /// The total storage space of a replica set instance, or the storage space of a single shard in a sharded cluster, in GiB.
        /// </summary>
        [Input("storageSpaceGb")]
        public Input<int>? StorageSpaceGb { get; set; }

        /// <summary>
        /// The subnet id of instance.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// The password of database account.
        /// </summary>
        [Input("superAccountPassword")]
        public Input<string>? SuperAccountPassword { get; set; }

        [Input("tags")]
        private InputList<Inputs.InstanceTagGetArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.InstanceTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.InstanceTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The vpc ID.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The zone ID of instance.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public InstanceState()
        {
        }
    }
}