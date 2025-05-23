// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cloudfs
{
    public static class GetAccesses
    {
        /// <summary>
        /// Use this data source to query detailed information of cloudfs accesses
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
        ///     var @default = Volcengine.Cloudfs.GetAccesses.Invoke(new()
        ///     {
        ///         FsName = "tftest2",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAccessesResult> InvokeAsync(GetAccessesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccessesResult>("volcengine:cloudfs/getAccesses:getAccesses", args ?? new GetAccessesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cloudfs accesses
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
        ///     var @default = Volcengine.Cloudfs.GetAccesses.Invoke(new()
        ///     {
        ///         FsName = "tftest2",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAccessesResult> Invoke(GetAccessesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccessesResult>("volcengine:cloudfs/getAccesses:getAccesses", args ?? new GetAccessesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccessesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of file system.
        /// </summary>
        [Input("fsName", required: true)]
        public string FsName { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetAccessesArgs()
        {
        }
        public static new GetAccessesArgs Empty => new GetAccessesArgs();
    }

    public sealed class GetAccessesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of file system.
        /// </summary>
        [Input("fsName", required: true)]
        public Input<string> FsName { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetAccessesInvokeArgs()
        {
        }
        public static new GetAccessesInvokeArgs Empty => new GetAccessesInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccessesResult
    {
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccessesAccessResult> Accesses;
        /// <summary>
        /// The name of cloud fs.
        /// </summary>
        public readonly string FsName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetAccessesResult(
            ImmutableArray<Outputs.GetAccessesAccessResult> accesses,

            string fsName,

            string id,

            string? outputFile,

            int totalCount)
        {
            Accesses = accesses;
            FsName = fsName;
            Id = id;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
