// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls
{
    /// <summary>
    /// Provides a resource to manage tls consumer group
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
    ///     var foo = new Volcengine.Tls.ConsumerGroup("foo", new()
    ///     {
    ///         ConsumerGroupName = "tf-test-consumer-group",
    ///         HeartbeatTtl = 120,
    ///         OrderedConsume = false,
    ///         ProjectId = "17ba378d-de43-495e-8906-03aexxxxxx",
    ///         TopicIdLists = new[]
    ///         {
    ///             "0ed72ac8-9531-4967-b216-ac30xxxxxx",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ConsumerGroup can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:tls/consumerGroup:ConsumerGroup default resource_id
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:tls/consumerGroup:ConsumerGroup")]
    public partial class ConsumerGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the consumer group.
        /// </summary>
        [Output("consumerGroupName")]
        public Output<string> ConsumerGroupName { get; private set; } = null!;

        /// <summary>
        /// The time of heart rate expiration, measured in seconds, has a value range of 1 to 300.
        /// </summary>
        [Output("heartbeatTtl")]
        public Output<int> HeartbeatTtl { get; private set; } = null!;

        /// <summary>
        /// Whether to consume in sequence.
        /// </summary>
        [Output("orderedConsume")]
        public Output<bool> OrderedConsume { get; private set; } = null!;

        /// <summary>
        /// The log project ID to which the consumption group belongs.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The list of log topic ids to be consumed by the consumer group.
        /// </summary>
        [Output("topicIdLists")]
        public Output<ImmutableArray<string>> TopicIdLists { get; private set; } = null!;


        /// <summary>
        /// Create a ConsumerGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConsumerGroup(string name, ConsumerGroupArgs args, CustomResourceOptions? options = null)
            : base("volcengine:tls/consumerGroup:ConsumerGroup", name, args ?? new ConsumerGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConsumerGroup(string name, Input<string> id, ConsumerGroupState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:tls/consumerGroup:ConsumerGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConsumerGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConsumerGroup Get(string name, Input<string> id, ConsumerGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ConsumerGroup(name, id, state, options);
        }
    }

    public sealed class ConsumerGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the consumer group.
        /// </summary>
        [Input("consumerGroupName", required: true)]
        public Input<string> ConsumerGroupName { get; set; } = null!;

        /// <summary>
        /// The time of heart rate expiration, measured in seconds, has a value range of 1 to 300.
        /// </summary>
        [Input("heartbeatTtl", required: true)]
        public Input<int> HeartbeatTtl { get; set; } = null!;

        /// <summary>
        /// Whether to consume in sequence.
        /// </summary>
        [Input("orderedConsume", required: true)]
        public Input<bool> OrderedConsume { get; set; } = null!;

        /// <summary>
        /// The log project ID to which the consumption group belongs.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("topicIdLists", required: true)]
        private InputList<string>? _topicIdLists;

        /// <summary>
        /// The list of log topic ids to be consumed by the consumer group.
        /// </summary>
        public InputList<string> TopicIdLists
        {
            get => _topicIdLists ?? (_topicIdLists = new InputList<string>());
            set => _topicIdLists = value;
        }

        public ConsumerGroupArgs()
        {
        }
        public static new ConsumerGroupArgs Empty => new ConsumerGroupArgs();
    }

    public sealed class ConsumerGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the consumer group.
        /// </summary>
        [Input("consumerGroupName")]
        public Input<string>? ConsumerGroupName { get; set; }

        /// <summary>
        /// The time of heart rate expiration, measured in seconds, has a value range of 1 to 300.
        /// </summary>
        [Input("heartbeatTtl")]
        public Input<int>? HeartbeatTtl { get; set; }

        /// <summary>
        /// Whether to consume in sequence.
        /// </summary>
        [Input("orderedConsume")]
        public Input<bool>? OrderedConsume { get; set; }

        /// <summary>
        /// The log project ID to which the consumption group belongs.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("topicIdLists")]
        private InputList<string>? _topicIdLists;

        /// <summary>
        /// The list of log topic ids to be consumed by the consumer group.
        /// </summary>
        public InputList<string> TopicIdLists
        {
            get => _topicIdLists ?? (_topicIdLists = new InputList<string>());
            set => _topicIdLists = value;
        }

        public ConsumerGroupState()
        {
        }
        public static new ConsumerGroupState Empty => new ConsumerGroupState();
    }
}
