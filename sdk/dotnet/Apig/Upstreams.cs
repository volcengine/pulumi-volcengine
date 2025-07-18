// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Apig
{
    [Obsolete(@"volcengine.apig.Upstreams has been deprecated in favor of volcengine.apig.getUpstreams")]
    public static class Upstreams
    {
        /// <summary>
        /// Use this data source to query detailed information of apig upstreams
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
        ///     var foo = Volcengine.Apig.GetUpstreams.Invoke(new()
        ///     {
        ///         GatewayId = "gd13d8c6eq1emkiunq6p0",
        ///         Ids = new[]
        ///         {
        ///             "ud18p5krj5ce3htvrd0v0",
        ///             "ud18ouitrjp6fhvu61n7g",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<UpstreamsResult> InvokeAsync(UpstreamsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<UpstreamsResult>("volcengine:apig/upstreams:Upstreams", args ?? new UpstreamsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of apig upstreams
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
        ///     var foo = Volcengine.Apig.GetUpstreams.Invoke(new()
        ///     {
        ///         GatewayId = "gd13d8c6eq1emkiunq6p0",
        ///         Ids = new[]
        ///         {
        ///             "ud18p5krj5ce3htvrd0v0",
        ///             "ud18ouitrjp6fhvu61n7g",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<UpstreamsResult> Invoke(UpstreamsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<UpstreamsResult>("volcengine:apig/upstreams:Upstreams", args ?? new UpstreamsInvokeArgs(), options.WithDefaults());
    }


    public sealed class UpstreamsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of api gateway.
        /// </summary>
        [Input("gatewayId")]
        public string? GatewayId { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of apig upstream IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of apig upstream. This field support fuzzy query.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// A Name Regex of Resource.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The resource type of apig upstream. Valid values: `Console`, `Ingress`.
        /// </summary>
        [Input("resourceType")]
        public string? ResourceType { get; set; }

        /// <summary>
        /// The source type of apig upstream. Valid values: `VeFaas`, `ECS`, `FixedIP`, `K8S`, `Nacos`, `Domain`, `AIProvider`, `VeMLP`.
        /// </summary>
        [Input("sourceType")]
        public string? SourceType { get; set; }

        public UpstreamsArgs()
        {
        }
        public static new UpstreamsArgs Empty => new UpstreamsArgs();
    }

    public sealed class UpstreamsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of api gateway.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of apig upstream IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of apig upstream. This field support fuzzy query.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A Name Regex of Resource.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The resource type of apig upstream. Valid values: `Console`, `Ingress`.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        /// <summary>
        /// The source type of apig upstream. Valid values: `VeFaas`, `ECS`, `FixedIP`, `K8S`, `Nacos`, `Domain`, `AIProvider`, `VeMLP`.
        /// </summary>
        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        public UpstreamsInvokeArgs()
        {
        }
        public static new UpstreamsInvokeArgs Empty => new UpstreamsInvokeArgs();
    }


    [OutputType]
    public sealed class UpstreamsResult
    {
        /// <summary>
        /// The id of api gateway.
        /// </summary>
        public readonly string? GatewayId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// The name of apig upstream version.
        /// </summary>
        public readonly string? Name;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The resource type of apig upstream.
        /// </summary>
        public readonly string? ResourceType;
        /// <summary>
        /// The source type of apig upstream.
        /// </summary>
        public readonly string? SourceType;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.UpstreamsUpstreamResult> Upstreams;

        [OutputConstructor]
        private UpstreamsResult(
            string? gatewayId,

            string id,

            ImmutableArray<string> ids,

            string? name,

            string? nameRegex,

            string? outputFile,

            string? resourceType,

            string? sourceType,

            int totalCount,

            ImmutableArray<Outputs.UpstreamsUpstreamResult> upstreams)
        {
            GatewayId = gatewayId;
            Id = id;
            Ids = ids;
            Name = name;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            ResourceType = resourceType;
            SourceType = sourceType;
            TotalCount = totalCount;
            Upstreams = upstreams;
        }
    }
}
