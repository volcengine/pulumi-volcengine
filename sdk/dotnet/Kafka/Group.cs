// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Kafka
{
    /// <summary>
    /// Provides a resource to manage kafka group
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
    ///     var fooInstance = new Volcengine.Kafka.Instance("fooInstance", new()
    ///     {
    ///         InstanceName = "acc-test-kafka",
    ///         InstanceDescription = "tf-test",
    ///         Version = "2.2.2",
    ///         ComputeSpec = "kafka.20xrate.hw",
    ///         SubnetId = fooSubnet.Id,
    ///         UserName = "tf-user",
    ///         UserPassword = "tf-pass!@q1",
    ///         ChargeType = "PostPaid",
    ///         StorageSpace = 300,
    ///         PartitionNumber = 350,
    ///         ProjectName = "default",
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Kafka.Inputs.InstanceTagArgs
    ///             {
    ///                 Key = "k1",
    ///                 Value = "v1",
    ///             },
    ///         },
    ///         Parameters = new[]
    ///         {
    ///             new Volcengine.Kafka.Inputs.InstanceParameterArgs
    ///             {
    ///                 ParameterName = "MessageMaxByte",
    ///                 ParameterValue = "12",
    ///             },
    ///             new Volcengine.Kafka.Inputs.InstanceParameterArgs
    ///             {
    ///                 ParameterName = "LogRetentionHours",
    ///                 ParameterValue = "70",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var fooGroup = new Volcengine.Kafka.Group("fooGroup", new()
    ///     {
    ///         InstanceId = fooInstance.Id,
    ///         GroupId = "acc-test-group",
    ///         Description = "tf-test",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// KafkaGroup can be imported using the instance_id:group_id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:kafka/group:Group default kafka-****x:groupId
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:kafka/group:Group")]
    public partial class Group : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of kafka group.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The id of kafka group.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The instance id of kafka group.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The state of kafka group.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs args, CustomResourceOptions? options = null)
            : base("volcengine:kafka/group:Group", name, args ?? new GroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:kafka/group:Group", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Group(name, id, state, options);
        }
    }

    public sealed class GroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of kafka group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The id of kafka group.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// The instance id of kafka group.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public GroupArgs()
        {
        }
        public static new GroupArgs Empty => new GroupArgs();
    }

    public sealed class GroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of kafka group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The id of kafka group.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The instance id of kafka group.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The state of kafka group.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public GroupState()
        {
        }
        public static new GroupState Empty => new GroupState();
    }
}