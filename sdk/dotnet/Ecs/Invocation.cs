// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Ecs
{
    /// <summary>
    /// Provides a resource to manage ecs invocation
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
    ///     var foo = new Volcengine.Ecs.Invocation("foo", new()
    ///     {
    ///         CommandId = "cmd-ychkepkhtim0tr3b****",
    ///         Frequency = "5m",
    ///         InstanceIds = new[]
    ///         {
    ///             "i-ychmz92487l8j00o****",
    ///         },
    ///         InvocationDescription = "tf",
    ///         InvocationName = "tf-test",
    ///         LaunchTime = "2023-06-20T09:48:00Z",
    ///         RecurrenceEndTime = "2023-06-20T09:59:00Z",
    ///         RepeatMode = "Rate",
    ///         Timeout = 90,
    ///         Username = "root",
    ///         WorkingDir = "/home",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// EcsInvocation can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:ecs/invocation:Invocation default ivk-ychnxnm45dl8j0mm****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:ecs/invocation:Invocation")]
    public partial class Invocation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The command id of the ecs invocation.
        /// </summary>
        [Output("commandId")]
        public Output<string> CommandId { get; private set; } = null!;

        /// <summary>
        /// The end time of the ecs invocation.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// The frequency of the ecs invocation. This field is valid and required when the value of the repeat_mode field is `Rate`.
        /// </summary>
        [Output("frequency")]
        public Output<string?> Frequency { get; private set; } = null!;

        /// <summary>
        /// The list of ECS instance IDs.
        /// </summary>
        [Output("instanceIds")]
        public Output<ImmutableArray<string>> InstanceIds { get; private set; } = null!;

        /// <summary>
        /// The description of the ecs invocation.
        /// </summary>
        [Output("invocationDescription")]
        public Output<string?> InvocationDescription { get; private set; } = null!;

        /// <summary>
        /// The name of the ecs invocation.
        /// </summary>
        [Output("invocationName")]
        public Output<string> InvocationName { get; private set; } = null!;

        /// <summary>
        /// The status of the ecs invocation.
        /// </summary>
        [Output("invocationStatus")]
        public Output<string> InvocationStatus { get; private set; } = null!;

        /// <summary>
        /// The launch time of the ecs invocation. RFC3339 format. This field is valid and required when the value of the repeat_mode field is `Rate` or `Fixed`.
        /// </summary>
        [Output("launchTime")]
        public Output<string?> LaunchTime { get; private set; } = null!;

        /// <summary>
        /// The recurrence end time of the ecs invocation. RFC3339 format. This field is valid and required when the value of the repeat_mode field is `Rate`.
        /// </summary>
        [Output("recurrenceEndTime")]
        public Output<string?> RecurrenceEndTime { get; private set; } = null!;

        /// <summary>
        /// The repeat mode of the ecs invocation. Valid values: `Once`, `Rate`, `Fixed`.
        /// </summary>
        [Output("repeatMode")]
        public Output<string?> RepeatMode { get; private set; } = null!;

        /// <summary>
        /// The start time of the ecs invocation.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;

        /// <summary>
        /// The timeout of the ecs command. Valid value range: 10-600. When this field is not specified, use the value of the field with the same name in ecs command as the default value.
        /// </summary>
        [Output("timeout")]
        public Output<int> Timeout { get; private set; } = null!;

        /// <summary>
        /// The username of the ecs command. When this field is not specified, use the value of the field with the same name in ecs command as the default value.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;

        /// <summary>
        /// The working directory of the ecs invocation. When this field is not specified, use the value of the field with the same name in ecs command as the default value.
        /// </summary>
        [Output("workingDir")]
        public Output<string> WorkingDir { get; private set; } = null!;


        /// <summary>
        /// Create a Invocation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Invocation(string name, InvocationArgs args, CustomResourceOptions? options = null)
            : base("volcengine:ecs/invocation:Invocation", name, args ?? new InvocationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Invocation(string name, Input<string> id, InvocationState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:ecs/invocation:Invocation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Invocation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Invocation Get(string name, Input<string> id, InvocationState? state = null, CustomResourceOptions? options = null)
        {
            return new Invocation(name, id, state, options);
        }
    }

    public sealed class InvocationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The command id of the ecs invocation.
        /// </summary>
        [Input("commandId", required: true)]
        public Input<string> CommandId { get; set; } = null!;

        /// <summary>
        /// The frequency of the ecs invocation. This field is valid and required when the value of the repeat_mode field is `Rate`.
        /// </summary>
        [Input("frequency")]
        public Input<string>? Frequency { get; set; }

        [Input("instanceIds", required: true)]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// The list of ECS instance IDs.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// The description of the ecs invocation.
        /// </summary>
        [Input("invocationDescription")]
        public Input<string>? InvocationDescription { get; set; }

        /// <summary>
        /// The name of the ecs invocation.
        /// </summary>
        [Input("invocationName", required: true)]
        public Input<string> InvocationName { get; set; } = null!;

        /// <summary>
        /// The launch time of the ecs invocation. RFC3339 format. This field is valid and required when the value of the repeat_mode field is `Rate` or `Fixed`.
        /// </summary>
        [Input("launchTime")]
        public Input<string>? LaunchTime { get; set; }

        /// <summary>
        /// The recurrence end time of the ecs invocation. RFC3339 format. This field is valid and required when the value of the repeat_mode field is `Rate`.
        /// </summary>
        [Input("recurrenceEndTime")]
        public Input<string>? RecurrenceEndTime { get; set; }

        /// <summary>
        /// The repeat mode of the ecs invocation. Valid values: `Once`, `Rate`, `Fixed`.
        /// </summary>
        [Input("repeatMode")]
        public Input<string>? RepeatMode { get; set; }

        /// <summary>
        /// The timeout of the ecs command. Valid value range: 10-600. When this field is not specified, use the value of the field with the same name in ecs command as the default value.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The username of the ecs command. When this field is not specified, use the value of the field with the same name in ecs command as the default value.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        /// <summary>
        /// The working directory of the ecs invocation. When this field is not specified, use the value of the field with the same name in ecs command as the default value.
        /// </summary>
        [Input("workingDir")]
        public Input<string>? WorkingDir { get; set; }

        public InvocationArgs()
        {
        }
        public static new InvocationArgs Empty => new InvocationArgs();
    }

    public sealed class InvocationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The command id of the ecs invocation.
        /// </summary>
        [Input("commandId")]
        public Input<string>? CommandId { get; set; }

        /// <summary>
        /// The end time of the ecs invocation.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// The frequency of the ecs invocation. This field is valid and required when the value of the repeat_mode field is `Rate`.
        /// </summary>
        [Input("frequency")]
        public Input<string>? Frequency { get; set; }

        [Input("instanceIds")]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// The list of ECS instance IDs.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        /// <summary>
        /// The description of the ecs invocation.
        /// </summary>
        [Input("invocationDescription")]
        public Input<string>? InvocationDescription { get; set; }

        /// <summary>
        /// The name of the ecs invocation.
        /// </summary>
        [Input("invocationName")]
        public Input<string>? InvocationName { get; set; }

        /// <summary>
        /// The status of the ecs invocation.
        /// </summary>
        [Input("invocationStatus")]
        public Input<string>? InvocationStatus { get; set; }

        /// <summary>
        /// The launch time of the ecs invocation. RFC3339 format. This field is valid and required when the value of the repeat_mode field is `Rate` or `Fixed`.
        /// </summary>
        [Input("launchTime")]
        public Input<string>? LaunchTime { get; set; }

        /// <summary>
        /// The recurrence end time of the ecs invocation. RFC3339 format. This field is valid and required when the value of the repeat_mode field is `Rate`.
        /// </summary>
        [Input("recurrenceEndTime")]
        public Input<string>? RecurrenceEndTime { get; set; }

        /// <summary>
        /// The repeat mode of the ecs invocation. Valid values: `Once`, `Rate`, `Fixed`.
        /// </summary>
        [Input("repeatMode")]
        public Input<string>? RepeatMode { get; set; }

        /// <summary>
        /// The start time of the ecs invocation.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// The timeout of the ecs command. Valid value range: 10-600. When this field is not specified, use the value of the field with the same name in ecs command as the default value.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The username of the ecs command. When this field is not specified, use the value of the field with the same name in ecs command as the default value.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// The working directory of the ecs invocation. When this field is not specified, use the value of the field with the same name in ecs command as the default value.
        /// </summary>
        [Input("workingDir")]
        public Input<string>? WorkingDir { get; set; }

        public InvocationState()
        {
        }
        public static new InvocationState Empty => new InvocationState();
    }
}
