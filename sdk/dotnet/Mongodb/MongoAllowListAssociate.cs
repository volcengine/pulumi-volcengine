// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Mongodb
{
    /// <summary>
    /// Provides a resource to manage mongodb allow list associate
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
    ///     var fooInstance = new Volcengine.Mongodb.Instance("fooInstance", new()
    ///     {
    ///         DbEngineVersion = "MongoDB_4_0",
    ///         InstanceType = "ReplicaSet",
    ///         SuperAccountPassword = "@acc-test-123",
    ///         NodeSpec = "mongo.2c4g",
    ///         MongosNodeSpec = "mongo.mongos.2c4g",
    ///         InstanceName = "acc-test-mongo-replica",
    ///         ChargeType = "PostPaid",
    ///         ProjectName = "default",
    ///         MongosNodeNumber = 32,
    ///         ShardNumber = 3,
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
    ///     var fooMongoAllowList = new Volcengine.Mongodb.MongoAllowList("fooMongoAllowList", new()
    ///     {
    ///         AllowListName = "acc-test",
    ///         AllowListDesc = "acc-test",
    ///         AllowListType = "IPv4",
    ///         AllowList = "10.1.1.3,10.2.3.0/24,10.1.1.1",
    ///     });
    /// 
    ///     var fooMongoAllowListAssociate = new Volcengine.Mongodb.MongoAllowListAssociate("fooMongoAllowListAssociate", new()
    ///     {
    ///         AllowListId = fooMongoAllowList.Id,
    ///         InstanceId = fooInstance.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// mongodb allow list associate can be imported using the instanceId:allowListId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:mongodb/mongoAllowListAssociate:MongoAllowListAssociate default mongo-replica-e405f8e2****:acl-d1fd76693bd54e658912e7337d5b****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:mongodb/mongoAllowListAssociate:MongoAllowListAssociate")]
    public partial class MongoAllowListAssociate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Id of allow list to associate.
        /// </summary>
        [Output("allowListId")]
        public Output<string> AllowListId { get; private set; } = null!;

        /// <summary>
        /// Id of instance to associate.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a MongoAllowListAssociate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MongoAllowListAssociate(string name, MongoAllowListAssociateArgs args, CustomResourceOptions? options = null)
            : base("volcengine:mongodb/mongoAllowListAssociate:MongoAllowListAssociate", name, args ?? new MongoAllowListAssociateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MongoAllowListAssociate(string name, Input<string> id, MongoAllowListAssociateState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:mongodb/mongoAllowListAssociate:MongoAllowListAssociate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MongoAllowListAssociate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MongoAllowListAssociate Get(string name, Input<string> id, MongoAllowListAssociateState? state = null, CustomResourceOptions? options = null)
        {
            return new MongoAllowListAssociate(name, id, state, options);
        }
    }

    public sealed class MongoAllowListAssociateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of allow list to associate.
        /// </summary>
        [Input("allowListId", required: true)]
        public Input<string> AllowListId { get; set; } = null!;

        /// <summary>
        /// Id of instance to associate.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public MongoAllowListAssociateArgs()
        {
        }
        public static new MongoAllowListAssociateArgs Empty => new MongoAllowListAssociateArgs();
    }

    public sealed class MongoAllowListAssociateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of allow list to associate.
        /// </summary>
        [Input("allowListId")]
        public Input<string>? AllowListId { get; set; }

        /// <summary>
        /// Id of instance to associate.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public MongoAllowListAssociateState()
        {
        }
        public static new MongoAllowListAssociateState Empty => new MongoAllowListAssociateState();
    }
}