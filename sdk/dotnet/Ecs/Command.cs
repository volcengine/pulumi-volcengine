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
    /// Provides a resource to manage ecs command
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
    ///     var foo = new Volcengine.Ecs.Command("foo", new()
    ///     {
    ///         CommandContent = "IyEvYmluL2Jhc2gKCgplY2hvICJvcGVyYXRpb24gc3VjY2VzcyEi",
    ///         Description = "tf",
    ///         Timeout = 100,
    ///         Username = "root",
    ///         WorkingDir = "/home",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// EcsCommand can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:ecs/command:Command default cmd-ychkepkhtim0tr3bcsw1
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:ecs/command:Command")]
    public partial class Command : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The base64 encoded content of the ecs command.
        /// </summary>
        [Output("commandContent")]
        public Output<string> CommandContent { get; private set; } = null!;

        /// <summary>
        /// The create time of the ecs command.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The description of the ecs command.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The invocation times of the ecs command. Public commands do not display the invocation times.
        /// </summary>
        [Output("invocationTimes")]
        public Output<int> InvocationTimes { get; private set; } = null!;

        /// <summary>
        /// The name of the ecs command.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The timeout of the ecs command. Valid value range: 10-600.
        /// </summary>
        [Output("timeout")]
        public Output<int> Timeout { get; private set; } = null!;

        /// <summary>
        /// The update time of the ecs command.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The username of the ecs command.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;

        /// <summary>
        /// The working directory of the ecs command.
        /// </summary>
        [Output("workingDir")]
        public Output<string> WorkingDir { get; private set; } = null!;


        /// <summary>
        /// Create a Command resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Command(string name, CommandArgs args, CustomResourceOptions? options = null)
            : base("volcengine:ecs/command:Command", name, args ?? new CommandArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Command(string name, Input<string> id, CommandState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:ecs/command:Command", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Command resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Command Get(string name, Input<string> id, CommandState? state = null, CustomResourceOptions? options = null)
        {
            return new Command(name, id, state, options);
        }
    }

    public sealed class CommandArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The base64 encoded content of the ecs command.
        /// </summary>
        [Input("commandContent", required: true)]
        public Input<string> CommandContent { get; set; } = null!;

        /// <summary>
        /// The description of the ecs command.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the ecs command.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The timeout of the ecs command. Valid value range: 10-600.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The username of the ecs command.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// The working directory of the ecs command.
        /// </summary>
        [Input("workingDir")]
        public Input<string>? WorkingDir { get; set; }

        public CommandArgs()
        {
        }
        public static new CommandArgs Empty => new CommandArgs();
    }

    public sealed class CommandState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The base64 encoded content of the ecs command.
        /// </summary>
        [Input("commandContent")]
        public Input<string>? CommandContent { get; set; }

        /// <summary>
        /// The create time of the ecs command.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The description of the ecs command.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The invocation times of the ecs command. Public commands do not display the invocation times.
        /// </summary>
        [Input("invocationTimes")]
        public Input<int>? InvocationTimes { get; set; }

        /// <summary>
        /// The name of the ecs command.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The timeout of the ecs command. Valid value range: 10-600.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The update time of the ecs command.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The username of the ecs command.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// The working directory of the ecs command.
        /// </summary>
        [Input("workingDir")]
        public Input<string>? WorkingDir { get; set; }

        public CommandState()
        {
        }
        public static new CommandState Empty => new CommandState();
    }
}
