// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vefaas
{
    public static class GetTimers
    {
        /// <summary>
        /// Use this data source to query detailed information of vefaas timers
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
        ///     var foo = Volcengine.Vefaas.GetTimers.Invoke(new()
        ///     {
        ///         FunctionId = "g79asxxx",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetTimersResult> InvokeAsync(GetTimersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTimersResult>("volcengine:vefaas/getTimers:getTimers", args ?? new GetTimersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of vefaas timers
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
        ///     var foo = Volcengine.Vefaas.GetTimers.Invoke(new()
        ///     {
        ///         FunctionId = "g79asxxx",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetTimersResult> Invoke(GetTimersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTimersResult>("volcengine:vefaas/getTimers:getTimers", args ?? new GetTimersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTimersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of Function.
        /// </summary>
        [Input("functionId", required: true)]
        public string FunctionId { get; set; } = null!;

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

        public GetTimersArgs()
        {
        }
        public static new GetTimersArgs Empty => new GetTimersArgs();
    }

    public sealed class GetTimersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of Function.
        /// </summary>
        [Input("functionId", required: true)]
        public Input<string> FunctionId { get; set; } = null!;

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

        public GetTimersInvokeArgs()
        {
        }
        public static new GetTimersInvokeArgs Empty => new GetTimersInvokeArgs();
    }


    [OutputType]
    public sealed class GetTimersResult
    {
        /// <summary>
        /// The ID of Function.
        /// </summary>
        public readonly string FunctionId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of timer trigger.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTimersItemResult> Items;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetTimersResult(
            string functionId,

            string id,

            ImmutableArray<Outputs.GetTimersItemResult> items,

            string? nameRegex,

            string? outputFile,

            int totalCount)
        {
            FunctionId = functionId;
            Id = id;
            Items = items;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
