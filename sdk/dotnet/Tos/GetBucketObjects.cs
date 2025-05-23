// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tos
{
    public static class GetBucketObjects
    {
        /// <summary>
        /// Use this data source to query detailed information of tos objects
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
        ///     var @default = Volcengine.Tos.GetBucketObjects.Invoke(new()
        ///     {
        ///         BucketName = "test",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetBucketObjectsResult> InvokeAsync(GetBucketObjectsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBucketObjectsResult>("volcengine:tos/getBucketObjects:getBucketObjects", args ?? new GetBucketObjectsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tos objects
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
        ///     var @default = Volcengine.Tos.GetBucketObjects.Invoke(new()
        ///     {
        ///         BucketName = "test",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetBucketObjectsResult> Invoke(GetBucketObjectsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBucketObjectsResult>("volcengine:tos/getBucketObjects:getBucketObjects", args ?? new GetBucketObjectsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBucketObjectsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name the TOS bucket.
        /// </summary>
        [Input("bucketName", required: true)]
        public string BucketName { get; set; } = null!;

        /// <summary>
        /// A Name Regex of TOS Object.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// The name the TOS Object.
        /// </summary>
        [Input("objectName")]
        public string? ObjectName { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetBucketObjectsArgs()
        {
        }
        public static new GetBucketObjectsArgs Empty => new GetBucketObjectsArgs();
    }

    public sealed class GetBucketObjectsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name the TOS bucket.
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        /// <summary>
        /// A Name Regex of TOS Object.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// The name the TOS Object.
        /// </summary>
        [Input("objectName")]
        public Input<string>? ObjectName { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetBucketObjectsInvokeArgs()
        {
        }
        public static new GetBucketObjectsInvokeArgs Empty => new GetBucketObjectsInvokeArgs();
    }


    [OutputType]
    public sealed class GetBucketObjectsResult
    {
        public readonly string BucketName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? NameRegex;
        public readonly string? ObjectName;
        /// <summary>
        /// The collection of TOS Object query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBucketObjectsObjectResult> Objects;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of TOS Object query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetBucketObjectsResult(
            string bucketName,

            string id,

            string? nameRegex,

            string? objectName,

            ImmutableArray<Outputs.GetBucketObjectsObjectResult> objects,

            string? outputFile,

            int totalCount)
        {
            BucketName = bucketName;
            Id = id;
            NameRegex = nameRegex;
            ObjectName = objectName;
            Objects = objects;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
