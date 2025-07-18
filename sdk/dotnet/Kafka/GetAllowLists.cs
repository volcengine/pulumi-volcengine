// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Kafka
{
    public static class GetAllowLists
    {
        /// <summary>
        /// Use this data source to query detailed information of kafka allow lists
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
        ///     var foo = Volcengine.Kafka.GetAllowLists.Invoke(new()
        ///     {
        ///         InstanceId = "kafka-xxx",
        ///         RegionId = "cn-beijing",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAllowListsResult> InvokeAsync(GetAllowListsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAllowListsResult>("volcengine:kafka/getAllowLists:getAllowLists", args ?? new GetAllowListsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of kafka allow lists
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
        ///     var foo = Volcengine.Kafka.GetAllowLists.Invoke(new()
        ///     {
        ///         InstanceId = "kafka-xxx",
        ///         RegionId = "cn-beijing",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAllowListsResult> Invoke(GetAllowListsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAllowListsResult>("volcengine:kafka/getAllowLists:getAllowLists", args ?? new GetAllowListsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAllowListsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The instance ID to query.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The region ID.
        /// </summary>
        [Input("regionId", required: true)]
        public string RegionId { get; set; } = null!;

        public GetAllowListsArgs()
        {
        }
        public static new GetAllowListsArgs Empty => new GetAllowListsArgs();
    }

    public sealed class GetAllowListsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The instance ID to query.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The region ID.
        /// </summary>
        [Input("regionId", required: true)]
        public Input<string> RegionId { get; set; } = null!;

        public GetAllowListsInvokeArgs()
        {
        }
        public static new GetAllowListsInvokeArgs Empty => new GetAllowListsInvokeArgs();
    }


    [OutputType]
    public sealed class GetAllowListsResult
    {
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAllowListsAllowListResult> AllowLists;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The id of the instance.
        /// </summary>
        public readonly string? InstanceId;
        public readonly string? OutputFile;
        public readonly string RegionId;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetAllowListsResult(
            ImmutableArray<Outputs.GetAllowListsAllowListResult> allowLists,

            string id,

            string? instanceId,

            string? outputFile,

            string regionId,

            int totalCount)
        {
            AllowLists = allowLists;
            Id = id;
            InstanceId = instanceId;
            OutputFile = outputFile;
            RegionId = regionId;
            TotalCount = totalCount;
        }
    }
}
