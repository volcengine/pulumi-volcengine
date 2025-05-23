// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Veenedge
{
    public static class GetInstanceTypes
    {
        /// <summary>
        /// Use this data source to query detailed information of veenedge instance types
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
        ///     var @default = Volcengine.Veenedge.GetInstanceTypes.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetInstanceTypesResult> InvokeAsync(GetInstanceTypesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceTypesResult>("volcengine:veenedge/getInstanceTypes:getInstanceTypes", args ?? new GetInstanceTypesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of veenedge instance types
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
        ///     var @default = Volcengine.Veenedge.GetInstanceTypes.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetInstanceTypesResult> Invoke(GetInstanceTypesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceTypesResult>("volcengine:veenedge/getInstanceTypes:getInstanceTypes", args ?? new GetInstanceTypesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceTypesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetInstanceTypesArgs()
        {
        }
        public static new GetInstanceTypesArgs Empty => new GetInstanceTypesArgs();
    }

    public sealed class GetInstanceTypesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetInstanceTypesInvokeArgs()
        {
        }
        public static new GetInstanceTypesInvokeArgs Empty => new GetInstanceTypesInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceTypesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The collection of instance types query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceTypesInstanceTypeConfigResult> InstanceTypeConfigs;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of instance types query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetInstanceTypesResult(
            string id,

            ImmutableArray<Outputs.GetInstanceTypesInstanceTypeConfigResult> instanceTypeConfigs,

            string? outputFile,

            int totalCount)
        {
            Id = id;
            InstanceTypeConfigs = instanceTypeConfigs;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
