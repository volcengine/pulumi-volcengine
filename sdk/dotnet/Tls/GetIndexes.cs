// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls
{
    public static class GetIndexes
    {
        /// <summary>
        /// Use this data source to query detailed information of tls indexes
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
        ///     var @default = Volcengine.Tls.GetIndexes.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "7ce12237-6670-44a7-9d79-2e36961586e6",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetIndexesResult> InvokeAsync(GetIndexesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIndexesResult>("volcengine:tls/getIndexes:getIndexes", args ?? new GetIndexesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tls indexes
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
        ///     var @default = Volcengine.Tls.GetIndexes.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "7ce12237-6670-44a7-9d79-2e36961586e6",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetIndexesResult> Invoke(GetIndexesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIndexesResult>("volcengine:tls/getIndexes:getIndexes", args ?? new GetIndexesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIndexesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids", required: true)]
        private List<string>? _ids;

        /// <summary>
        /// The list of topic id of tls index.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetIndexesArgs()
        {
        }
        public static new GetIndexesArgs Empty => new GetIndexesArgs();
    }

    public sealed class GetIndexesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids", required: true)]
        private InputList<string>? _ids;

        /// <summary>
        /// The list of topic id of tls index.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetIndexesInvokeArgs()
        {
        }
        public static new GetIndexesInvokeArgs Empty => new GetIndexesInvokeArgs();
    }


    [OutputType]
    public sealed class GetIndexesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        /// <summary>
        /// The collection of tls index query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIndexesTlsIndexResult> TlsIndexes;
        /// <summary>
        /// The total count of tls index query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetIndexesResult(
            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            ImmutableArray<Outputs.GetIndexesTlsIndexResult> tlsIndexes,

            int totalCount)
        {
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            TlsIndexes = tlsIndexes;
            TotalCount = totalCount;
        }
    }
}
