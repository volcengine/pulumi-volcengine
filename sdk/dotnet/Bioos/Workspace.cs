// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Bioos
{
    /// <summary>
    /// Provides a resource to manage bioos workspace
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
    ///     var foo = new Volcengine.Bioos.Workspace("foo", new()
    ///     {
    ///         CoverPath = "template-cover/pic5.png",
    ///         Description = "test-description23",
    ///     });
    /// 
    ///     //必填
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Workspace can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:bioos/workspace:Workspace default *****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:bioos/workspace:Workspace")]
    public partial class Workspace : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cover path (relative path in tos bucket).
        /// </summary>
        [Output("coverPath")]
        public Output<string> CoverPath { get; private set; } = null!;

        /// <summary>
        /// The description of the workspace.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Whether the update complete.
        /// </summary>
        [Output("updated")]
        public Output<bool> Updated { get; private set; } = null!;

        /// <summary>
        /// The id of the workspace.
        /// </summary>
        [Output("workspaceId")]
        public Output<string> WorkspaceId { get; private set; } = null!;


        /// <summary>
        /// Create a Workspace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Workspace(string name, WorkspaceArgs args, CustomResourceOptions? options = null)
            : base("volcengine:bioos/workspace:Workspace", name, args ?? new WorkspaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Workspace(string name, Input<string> id, WorkspaceState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:bioos/workspace:Workspace", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Workspace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Workspace Get(string name, Input<string> id, WorkspaceState? state = null, CustomResourceOptions? options = null)
        {
            return new Workspace(name, id, state, options);
        }
    }

    public sealed class WorkspaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cover path (relative path in tos bucket).
        /// </summary>
        [Input("coverPath")]
        public Input<string>? CoverPath { get; set; }

        /// <summary>
        /// The description of the workspace.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public WorkspaceArgs()
        {
        }
        public static new WorkspaceArgs Empty => new WorkspaceArgs();
    }

    public sealed class WorkspaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cover path (relative path in tos bucket).
        /// </summary>
        [Input("coverPath")]
        public Input<string>? CoverPath { get; set; }

        /// <summary>
        /// The description of the workspace.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether the update complete.
        /// </summary>
        [Input("updated")]
        public Input<bool>? Updated { get; set; }

        /// <summary>
        /// The id of the workspace.
        /// </summary>
        [Input("workspaceId")]
        public Input<string>? WorkspaceId { get; set; }

        public WorkspaceState()
        {
        }
        public static new WorkspaceState Empty => new WorkspaceState();
    }
}
