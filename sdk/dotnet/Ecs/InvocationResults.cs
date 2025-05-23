// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Ecs
{
    [Obsolete(@"volcengine.ecs.InvocationResults has been deprecated in favor of volcengine.ecs.getInvocationResults")]
    public static class InvocationResults
    {
        /// <summary>
        /// Use this data source to query detailed information of ecs invocation results
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
        ///     var @default = Volcengine.Ecs.GetInvocationResults.Invoke(new()
        ///     {
        ///         InvocationId = "ivk-ych9y4vujvl8j01c****",
        ///         InvocationResultStatuses = new[]
        ///         {
        ///             "Success",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<InvocationResultsResult> InvokeAsync(InvocationResultsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<InvocationResultsResult>("volcengine:ecs/invocationResults:InvocationResults", args ?? new InvocationResultsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of ecs invocation results
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
        ///     var @default = Volcengine.Ecs.GetInvocationResults.Invoke(new()
        ///     {
        ///         InvocationId = "ivk-ych9y4vujvl8j01c****",
        ///         InvocationResultStatuses = new[]
        ///         {
        ///             "Success",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<InvocationResultsResult> Invoke(InvocationResultsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<InvocationResultsResult>("volcengine:ecs/invocationResults:InvocationResults", args ?? new InvocationResultsInvokeArgs(), options.WithDefaults());
    }


    public sealed class InvocationResultsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of ecs command.
        /// </summary>
        [Input("commandId")]
        public string? CommandId { get; set; }

        /// <summary>
        /// The id of ecs instance.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// The id of ecs invocation.
        /// </summary>
        [Input("invocationId", required: true)]
        public string InvocationId { get; set; } = null!;

        [Input("invocationResultStatuses")]
        private List<string>? _invocationResultStatuses;

        /// <summary>
        /// The list of status of ecs invocation in a single instance. Valid values: `Pending`, `Running`, `Success`, `Failed`, `Timeout`.
        /// </summary>
        public List<string> InvocationResultStatuses
        {
            get => _invocationResultStatuses ?? (_invocationResultStatuses = new List<string>());
            set => _invocationResultStatuses = value;
        }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public InvocationResultsArgs()
        {
        }
        public static new InvocationResultsArgs Empty => new InvocationResultsArgs();
    }

    public sealed class InvocationResultsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of ecs command.
        /// </summary>
        [Input("commandId")]
        public Input<string>? CommandId { get; set; }

        /// <summary>
        /// The id of ecs instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The id of ecs invocation.
        /// </summary>
        [Input("invocationId", required: true)]
        public Input<string> InvocationId { get; set; } = null!;

        [Input("invocationResultStatuses")]
        private InputList<string>? _invocationResultStatuses;

        /// <summary>
        /// The list of status of ecs invocation in a single instance. Valid values: `Pending`, `Running`, `Success`, `Failed`, `Timeout`.
        /// </summary>
        public InputList<string> InvocationResultStatuses
        {
            get => _invocationResultStatuses ?? (_invocationResultStatuses = new InputList<string>());
            set => _invocationResultStatuses = value;
        }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public InvocationResultsInvokeArgs()
        {
        }
        public static new InvocationResultsInvokeArgs Empty => new InvocationResultsInvokeArgs();
    }


    [OutputType]
    public sealed class InvocationResultsResult
    {
        /// <summary>
        /// The id of the ecs command.
        /// </summary>
        public readonly string? CommandId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The id of the ecs instance.
        /// </summary>
        public readonly string? InstanceId;
        /// <summary>
        /// The id of the ecs invocation.
        /// </summary>
        public readonly string InvocationId;
        /// <summary>
        /// The status of ecs invocation in a single instance.
        /// </summary>
        public readonly ImmutableArray<string> InvocationResultStatuses;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.InvocationResultsInvocationResultResult> InvocationResults;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private InvocationResultsResult(
            string? commandId,

            string id,

            string? instanceId,

            string invocationId,

            ImmutableArray<string> invocationResultStatuses,

            ImmutableArray<Outputs.InvocationResultsInvocationResultResult> invocationResults,

            string? outputFile,

            int totalCount)
        {
            CommandId = commandId;
            Id = id;
            InstanceId = instanceId;
            InvocationId = invocationId;
            InvocationResultStatuses = invocationResultStatuses;
            InvocationResults = invocationResults;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
