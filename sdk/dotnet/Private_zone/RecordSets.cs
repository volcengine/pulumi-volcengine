// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Private_zone
{
    public static class RecordSets
    {
        /// <summary>
        /// Use this data source to query detailed information of private zone record sets
        /// </summary>
        public static Task<RecordSetsResult> InvokeAsync(RecordSetsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<RecordSetsResult>("volcengine:private_zone/recordSets:RecordSets", args ?? new RecordSetsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of private zone record sets
        /// </summary>
        public static Output<RecordSetsResult> Invoke(RecordSetsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<RecordSetsResult>("volcengine:private_zone/recordSets:RecordSets", args ?? new RecordSetsInvokeArgs(), options.WithDefaults());
    }


    public sealed class RecordSetsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The host of Private Zone Record Set.
        /// </summary>
        [Input("host")]
        public string? Host { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The id of Private Zone Record Set.
        /// </summary>
        [Input("recordSetId")]
        public string? RecordSetId { get; set; }

        /// <summary>
        /// The search mode of query `host`. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
        /// </summary>
        [Input("searchMode")]
        public string? SearchMode { get; set; }

        /// <summary>
        /// The zid of Private Zone.
        /// </summary>
        [Input("zid", required: true)]
        public int Zid { get; set; }

        public RecordSetsArgs()
        {
        }
        public static new RecordSetsArgs Empty => new RecordSetsArgs();
    }

    public sealed class RecordSetsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The host of Private Zone Record Set.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The id of Private Zone Record Set.
        /// </summary>
        [Input("recordSetId")]
        public Input<string>? RecordSetId { get; set; }

        /// <summary>
        /// The search mode of query `host`. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
        /// </summary>
        [Input("searchMode")]
        public Input<string>? SearchMode { get; set; }

        /// <summary>
        /// The zid of Private Zone.
        /// </summary>
        [Input("zid", required: true)]
        public Input<int> Zid { get; set; } = null!;

        public RecordSetsInvokeArgs()
        {
        }
        public static new RecordSetsInvokeArgs Empty => new RecordSetsInvokeArgs();
    }


    [OutputType]
    public sealed class RecordSetsResult
    {
        /// <summary>
        /// The host of the private zone record.
        /// </summary>
        public readonly string? Host;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OutputFile;
        /// <summary>
        /// The id of the private zone record set.
        /// </summary>
        public readonly string? RecordSetId;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.RecordSetsRecordSetResult> RecordSets;
        public readonly string? SearchMode;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;
        public readonly int Zid;

        [OutputConstructor]
        private RecordSetsResult(
            string? host,

            string id,

            string? outputFile,

            string? recordSetId,

            ImmutableArray<Outputs.RecordSetsRecordSetResult> recordSets,

            string? searchMode,

            int totalCount,

            int zid)
        {
            Host = host;
            Id = id;
            OutputFile = outputFile;
            RecordSetId = recordSetId;
            RecordSets = recordSets;
            SearchMode = searchMode;
            TotalCount = totalCount;
            Zid = zid;
        }
    }
}