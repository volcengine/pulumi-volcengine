// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Mongodb
{
    public static class SslStates
    {
        /// <summary>
        /// Use this data source to query detailed information of mongodb ssl states
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var foo = Output.Create(Volcengine.Mongodb.SslStates.InvokeAsync(new Volcengine.Mongodb.SslStatesArgs
        ///         {
        ///             InstanceId = "mongo-replica-f16e9298b121",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<SslStatesResult> InvokeAsync(SslStatesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<SslStatesResult>("volcengine:mongodb/sslStates:SslStates", args ?? new SslStatesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of mongodb ssl states
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var foo = Output.Create(Volcengine.Mongodb.SslStates.InvokeAsync(new Volcengine.Mongodb.SslStatesArgs
        ///         {
        ///             InstanceId = "mongo-replica-f16e9298b121",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<SslStatesResult> Invoke(SslStatesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<SslStatesResult>("volcengine:mongodb/sslStates:SslStates", args ?? new SslStatesInvokeArgs(), options.WithDefaults());
    }


    public sealed class SslStatesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The mongodb instance ID to query.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public SslStatesArgs()
        {
        }
    }

    public sealed class SslStatesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The mongodb instance ID to query.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public SslStatesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class SslStatesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The mongodb instance id.
        /// </summary>
        public readonly string InstanceId;
        public readonly string? OutputFile;
        /// <summary>
        /// The collection of mongodb ssl state query.
        /// </summary>
        public readonly ImmutableArray<Outputs.SslStatesSslStateResult> SslStates;
        /// <summary>
        /// The total count of mongodb ssl state query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private SslStatesResult(
            string id,

            string instanceId,

            string? outputFile,

            ImmutableArray<Outputs.SslStatesSslStateResult> sslStates,

            int totalCount)
        {
            Id = id;
            InstanceId = instanceId;
            OutputFile = outputFile;
            SslStates = sslStates;
            TotalCount = totalCount;
        }
    }
}