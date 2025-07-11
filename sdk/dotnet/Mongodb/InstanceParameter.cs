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
    /// Provides a resource to manage mongodb instance parameter
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
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
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
    ///         ZoneId = fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
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
    ///     var fooInstanceParameter = new Volcengine.Mongodb.InstanceParameter("fooInstanceParameter", new()
    ///     {
    ///         InstanceId = fooInstance.Id,
    ///         ParameterName = "cursorTimeoutMillis",
    ///         ParameterRole = "Node",
    ///         ParameterValue = "600111",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// mongodb parameter can be imported using the param:instanceId:parameterName:parameterRole, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:mongodb/instanceParameter:InstanceParameter default param:mongo-replica-e405f8e2****:connPoolMaxConnsPerHost
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:mongodb/instanceParameter:InstanceParameter")]
    public partial class InstanceParameter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The name of parameter. The parameter resource can only be added or modified, deleting this resource will not actually execute any operation.
        /// </summary>
        [Output("parameterName")]
        public Output<string> ParameterName { get; private set; } = null!;

        /// <summary>
        /// The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
        /// </summary>
        [Output("parameterRole")]
        public Output<string> ParameterRole { get; private set; } = null!;

        /// <summary>
        /// The value of parameter.
        /// </summary>
        [Output("parameterValue")]
        public Output<string> ParameterValue { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceParameter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceParameter(string name, InstanceParameterArgs args, CustomResourceOptions? options = null)
            : base("volcengine:mongodb/instanceParameter:InstanceParameter", name, args ?? new InstanceParameterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceParameter(string name, Input<string> id, InstanceParameterState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:mongodb/instanceParameter:InstanceParameter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceParameter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceParameter Get(string name, Input<string> id, InstanceParameterState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceParameter(name, id, state, options);
        }
    }

    public sealed class InstanceParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The name of parameter. The parameter resource can only be added or modified, deleting this resource will not actually execute any operation.
        /// </summary>
        [Input("parameterName", required: true)]
        public Input<string> ParameterName { get; set; } = null!;

        /// <summary>
        /// The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
        /// </summary>
        [Input("parameterRole", required: true)]
        public Input<string> ParameterRole { get; set; } = null!;

        /// <summary>
        /// The value of parameter.
        /// </summary>
        [Input("parameterValue", required: true)]
        public Input<string> ParameterValue { get; set; } = null!;

        public InstanceParameterArgs()
        {
        }
        public static new InstanceParameterArgs Empty => new InstanceParameterArgs();
    }

    public sealed class InstanceParameterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The name of parameter. The parameter resource can only be added or modified, deleting this resource will not actually execute any operation.
        /// </summary>
        [Input("parameterName")]
        public Input<string>? ParameterName { get; set; }

        /// <summary>
        /// The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
        /// </summary>
        [Input("parameterRole")]
        public Input<string>? ParameterRole { get; set; }

        /// <summary>
        /// The value of parameter.
        /// </summary>
        [Input("parameterValue")]
        public Input<string>? ParameterValue { get; set; }

        public InstanceParameterState()
        {
        }
        public static new InstanceParameterState Empty => new InstanceParameterState();
    }
}
