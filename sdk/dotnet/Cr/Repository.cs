// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cr
{
    /// <summary>
    /// Provides a resource to manage cr repository
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
    ///     var foo = new Volcengine.Cr.Repository("foo", new()
    ///     {
    ///         AccessLevel = "Public",
    ///         Description = "A test repository created by terraform.",
    ///         Namespace = "namespace-1",
    ///         Registry = "tf-2",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CR Repository can be imported using the registry:namespace:name, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:cr/repository:Repository default cr-basic:namespace-1:repo-1
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:cr/repository:Repository")]
    public partial class Repository : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The access level of CrRepository.
        /// </summary>
        [Output("accessLevel")]
        public Output<string?> AccessLevel { get; private set; } = null!;

        /// <summary>
        /// The creation time of repository.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The description of CrRepository.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of CrRepository.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The target namespace name.
        /// </summary>
        [Output("namespace")]
        public Output<string> Namespace { get; private set; } = null!;

        /// <summary>
        /// The CrRegistry name.
        /// </summary>
        [Output("registry")]
        public Output<string> Registry { get; private set; } = null!;

        /// <summary>
        /// The last update time of repository.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Repository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Repository(string name, RepositoryArgs args, CustomResourceOptions? options = null)
            : base("volcengine:cr/repository:Repository", name, args ?? new RepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Repository(string name, Input<string> id, RepositoryState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:cr/repository:Repository", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Repository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Repository Get(string name, Input<string> id, RepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new Repository(name, id, state, options);
        }
    }

    public sealed class RepositoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access level of CrRepository.
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        /// <summary>
        /// The description of CrRepository.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of CrRepository.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The target namespace name.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        /// <summary>
        /// The CrRegistry name.
        /// </summary>
        [Input("registry", required: true)]
        public Input<string> Registry { get; set; } = null!;

        public RepositoryArgs()
        {
        }
        public static new RepositoryArgs Empty => new RepositoryArgs();
    }

    public sealed class RepositoryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access level of CrRepository.
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        /// <summary>
        /// The creation time of repository.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The description of CrRepository.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of CrRepository.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The target namespace name.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The CrRegistry name.
        /// </summary>
        [Input("registry")]
        public Input<string>? Registry { get; set; }

        /// <summary>
        /// The last update time of repository.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public RepositoryState()
        {
        }
        public static new RepositoryState Empty => new RepositoryState();
    }
}
