// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Private_zone
{
    public static class Records
    {
        /// <summary>
        /// Use this data source to query detailed information of private zone records
        /// </summary>
        public static Task<RecordsResult> InvokeAsync(RecordsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<RecordsResult>("volcengine:private_zone/records:Records", args ?? new RecordsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of private zone records
        /// </summary>
        public static Output<RecordsResult> Invoke(RecordsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<RecordsResult>("volcengine:private_zone/records:Records", args ?? new RecordsInvokeArgs(), options.WithDefaults());
    }


    public sealed class RecordsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The host of Private Zone Record.
        /// </summary>
        [Input("host")]
        public string? Host { get; set; }

        /// <summary>
        /// The last operator account id of Private Zone Record.
        /// </summary>
        [Input("lastOperator")]
        public string? LastOperator { get; set; }

        /// <summary>
        /// The subnet id of Private Zone Record. This field is only effected when the `intelligent_mode` of the private zone is true.
        /// </summary>
        [Input("line")]
        public string? Line { get; set; }

        /// <summary>
        /// The domain name of Private Zone Record.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The id of Private Zone Record.
        /// </summary>
        [Input("recordId")]
        public string? RecordId { get; set; }

        /// <summary>
        /// The search mode of query `host`. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
        /// </summary>
        [Input("searchMode")]
        public string? SearchMode { get; set; }

        /// <summary>
        /// The type of Private Zone Record.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        /// <summary>
        /// The value of Private Zone Record.
        /// </summary>
        [Input("value")]
        public string? Value { get; set; }

        /// <summary>
        /// The zid of Private Zone.
        /// </summary>
        [Input("zid")]
        public int? Zid { get; set; }

        public RecordsArgs()
        {
        }
        public static new RecordsArgs Empty => new RecordsArgs();
    }

    public sealed class RecordsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The host of Private Zone Record.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The last operator account id of Private Zone Record.
        /// </summary>
        [Input("lastOperator")]
        public Input<string>? LastOperator { get; set; }

        /// <summary>
        /// The subnet id of Private Zone Record. This field is only effected when the `intelligent_mode` of the private zone is true.
        /// </summary>
        [Input("line")]
        public Input<string>? Line { get; set; }

        /// <summary>
        /// The domain name of Private Zone Record.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The id of Private Zone Record.
        /// </summary>
        [Input("recordId")]
        public Input<string>? RecordId { get; set; }

        /// <summary>
        /// The search mode of query `host`. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
        /// </summary>
        [Input("searchMode")]
        public Input<string>? SearchMode { get; set; }

        /// <summary>
        /// The type of Private Zone Record.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The value of Private Zone Record.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// The zid of Private Zone.
        /// </summary>
        [Input("zid")]
        public Input<int>? Zid { get; set; }

        public RecordsInvokeArgs()
        {
        }
        public static new RecordsInvokeArgs Empty => new RecordsInvokeArgs();
    }


    [OutputType]
    public sealed class RecordsResult
    {
        /// <summary>
        /// The host of the private zone record.
        /// </summary>
        public readonly string? Host;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The last operator account id of the private zone record.
        /// </summary>
        public readonly string? LastOperator;
        /// <summary>
        /// The subnet id of the private zone record. This field is only effected when the `intelligent_mode` of the private zone is true.
        /// </summary>
        public readonly string? Line;
        public readonly string? Name;
        public readonly string? OutputFile;
        /// <summary>
        /// The id of the private zone record.
        /// </summary>
        public readonly string? RecordId;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.RecordsRecordResult> Records;
        public readonly string? SearchMode;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The type of the private zone record.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// The value of the private zone record.
        /// </summary>
        public readonly string? Value;
        /// <summary>
        /// The zid of the private zone record.
        /// </summary>
        public readonly int? Zid;

        [OutputConstructor]
        private RecordsResult(
            string? host,

            string id,

            string? lastOperator,

            string? line,

            string? name,

            string? outputFile,

            string? recordId,

            ImmutableArray<Outputs.RecordsRecordResult> records,

            string? searchMode,

            int totalCount,

            string? type,

            string? value,

            int? zid)
        {
            Host = host;
            Id = id;
            LastOperator = lastOperator;
            Line = line;
            Name = name;
            OutputFile = outputFile;
            RecordId = recordId;
            Records = records;
            SearchMode = searchMode;
            TotalCount = totalCount;
            Type = type;
            Value = value;
            Zid = zid;
        }
    }
}