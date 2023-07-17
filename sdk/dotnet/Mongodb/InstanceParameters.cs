// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Mongodb
{
    public static class InstanceParameters
    {
        /// <summary>
        /// Use this data source to query detailed information of mongodb instance parameters
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
        ///         var foo = Output.Create(Volcengine.Mongodb.InstanceParameters.InvokeAsync(new Volcengine.Mongodb.InstanceParametersArgs
        ///         {
        ///             InstanceId = "mongo-replica-f16e9298b121",
        ///             ParameterNames = "connPoolMaxConnsPerHost",
        ///             ParameterRole = "Node",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<InstanceParametersResult> InvokeAsync(InstanceParametersArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<InstanceParametersResult>("volcengine:mongodb/instanceParameters:InstanceParameters", args ?? new InstanceParametersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of mongodb instance parameters
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
        ///         var foo = Output.Create(Volcengine.Mongodb.InstanceParameters.InvokeAsync(new Volcengine.Mongodb.InstanceParametersArgs
        ///         {
        ///             InstanceId = "mongo-replica-f16e9298b121",
        ///             ParameterNames = "connPoolMaxConnsPerHost",
        ///             ParameterRole = "Node",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<InstanceParametersResult> Invoke(InstanceParametersInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<InstanceParametersResult>("volcengine:mongodb/instanceParameters:InstanceParameters", args ?? new InstanceParametersInvokeArgs(), options.WithDefaults());
    }


    public sealed class InstanceParametersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The instance ID to query.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The parameter names, support fuzzy query, case insensitive.
        /// </summary>
        [Input("parameterNames")]
        public string? ParameterNames { get; set; }

        /// <summary>
        /// The node type of instance parameter, valid value contains `Node`, `Shard`, `ConfigServer`, `Mongos`.
        /// </summary>
        [Input("parameterRole")]
        public string? ParameterRole { get; set; }

        public InstanceParametersArgs()
        {
        }
    }

    public sealed class InstanceParametersInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The instance ID to query.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The parameter names, support fuzzy query, case insensitive.
        /// </summary>
        [Input("parameterNames")]
        public Input<string>? ParameterNames { get; set; }

        /// <summary>
        /// The node type of instance parameter, valid value contains `Node`, `Shard`, `ConfigServer`, `Mongos`.
        /// </summary>
        [Input("parameterRole")]
        public Input<string>? ParameterRole { get; set; }

        public InstanceParametersInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class InstanceParametersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The instance ID.
        /// </summary>
        public readonly string InstanceId;
        public readonly string? OutputFile;
        public readonly string? ParameterNames;
        /// <summary>
        /// The node type to which the parameter belongs.
        /// </summary>
        public readonly string? ParameterRole;
        /// <summary>
        /// The collection of parameter query.
        /// </summary>
        public readonly Outputs.InstanceParametersParametersResult Parameters;
        /// <summary>
        /// The total count of mongodb instance parameter query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private InstanceParametersResult(
            string id,

            string instanceId,

            string? outputFile,

            string? parameterNames,

            string? parameterRole,

            Outputs.InstanceParametersParametersResult parameters,

            int totalCount)
        {
            Id = id;
            InstanceId = instanceId;
            OutputFile = outputFile;
            ParameterNames = parameterNames;
            ParameterRole = parameterRole;
            Parameters = parameters;
            TotalCount = totalCount;
        }
    }
}