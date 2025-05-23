// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cr
{
    public static class GetRepositories
    {
        /// <summary>
        /// Use this data source to query detailed information of cr repositories
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
        ///     var foo = Volcengine.Cr.GetRepositories.Invoke(new()
        ///     {
        ///         Names = new[]
        ///         {
        ///             "repo*",
        ///         },
        ///         Registry = "tf-1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetRepositoriesResult> InvokeAsync(GetRepositoriesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepositoriesResult>("volcengine:cr/getRepositories:getRepositories", args ?? new GetRepositoriesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cr repositories
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
        ///     var foo = Volcengine.Cr.GetRepositories.Invoke(new()
        ///     {
        ///         Names = new[]
        ///         {
        ///             "repo*",
        ///         },
        ///         Registry = "tf-1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRepositoriesResult> Invoke(GetRepositoriesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoriesResult>("volcengine:cr/getRepositories:getRepositories", args ?? new GetRepositoriesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoriesArgs : global::Pulumi.InvokeArgs
    {
        [Input("accessLevels")]
        private List<string>? _accessLevels;

        /// <summary>
        /// The list of instance access level.
        /// </summary>
        public List<string> AccessLevels
        {
            get => _accessLevels ?? (_accessLevels = new List<string>());
            set => _accessLevels = value;
        }

        [Input("names")]
        private List<string>? _names;

        /// <summary>
        /// The list of instance names.
        /// </summary>
        public List<string> Names
        {
            get => _names ?? (_names = new List<string>());
            set => _names = value;
        }

        [Input("namespaces")]
        private List<string>? _namespaces;

        /// <summary>
        /// The list of instance namespace.
        /// </summary>
        public List<string> Namespaces
        {
            get => _namespaces ?? (_namespaces = new List<string>());
            set => _namespaces = value;
        }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The CR instance name.
        /// </summary>
        [Input("registry", required: true)]
        public string Registry { get; set; } = null!;

        public GetRepositoriesArgs()
        {
        }
        public static new GetRepositoriesArgs Empty => new GetRepositoriesArgs();
    }

    public sealed class GetRepositoriesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("accessLevels")]
        private InputList<string>? _accessLevels;

        /// <summary>
        /// The list of instance access level.
        /// </summary>
        public InputList<string> AccessLevels
        {
            get => _accessLevels ?? (_accessLevels = new InputList<string>());
            set => _accessLevels = value;
        }

        [Input("names")]
        private InputList<string>? _names;

        /// <summary>
        /// The list of instance names.
        /// </summary>
        public InputList<string> Names
        {
            get => _names ?? (_names = new InputList<string>());
            set => _names = value;
        }

        [Input("namespaces")]
        private InputList<string>? _namespaces;

        /// <summary>
        /// The list of instance namespace.
        /// </summary>
        public InputList<string> Namespaces
        {
            get => _namespaces ?? (_namespaces = new InputList<string>());
            set => _namespaces = value;
        }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The CR instance name.
        /// </summary>
        [Input("registry", required: true)]
        public Input<string> Registry { get; set; } = null!;

        public GetRepositoriesInvokeArgs()
        {
        }
        public static new GetRepositoriesInvokeArgs Empty => new GetRepositoriesInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoriesResult
    {
        public readonly ImmutableArray<string> AccessLevels;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Names;
        public readonly ImmutableArray<string> Namespaces;
        public readonly string? OutputFile;
        public readonly string Registry;
        /// <summary>
        /// The collection of repository query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRepositoriesRepositoryResult> Repositories;
        /// <summary>
        /// The total count of instance query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetRepositoriesResult(
            ImmutableArray<string> accessLevels,

            string id,

            ImmutableArray<string> names,

            ImmutableArray<string> namespaces,

            string? outputFile,

            string registry,

            ImmutableArray<Outputs.GetRepositoriesRepositoryResult> repositories,

            int totalCount)
        {
            AccessLevels = accessLevels;
            Id = id;
            Names = names;
            Namespaces = namespaces;
            OutputFile = outputFile;
            Registry = registry;
            Repositories = repositories;
            TotalCount = totalCount;
        }
    }
}
