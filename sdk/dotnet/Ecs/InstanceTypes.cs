// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Ecs
{
    public static class InstanceTypes
    {
        /// <summary>
        /// Use this data source to query detailed information of ecs instance types
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Volcengine.Ecs.InstanceTypes.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<InstanceTypesResult> InvokeAsync(InstanceTypesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<InstanceTypesResult>("volcengine:ecs/instanceTypes:InstanceTypes", args ?? new InstanceTypesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of ecs instance types
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Volcengine.Ecs.InstanceTypes.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<InstanceTypesResult> Invoke(InstanceTypesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<InstanceTypesResult>("volcengine:ecs/instanceTypes:InstanceTypes", args ?? new InstanceTypesInvokeArgs(), options.WithDefaults());
    }


    public sealed class InstanceTypesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of instance type IDs. When the number of ids is greater than 10, only the first 10 are effective.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public InstanceTypesArgs()
        {
        }
        public static new InstanceTypesArgs Empty => new InstanceTypesArgs();
    }

    public sealed class InstanceTypesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of instance type IDs. When the number of ids is greater than 10, only the first 10 are effective.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public InstanceTypesInvokeArgs()
        {
        }
        public static new InstanceTypesInvokeArgs Empty => new InstanceTypesInvokeArgs();
    }


    [OutputType]
    public sealed class InstanceTypesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceTypesInstanceTypeResult> InstanceTypes;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private InstanceTypesResult(
            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.InstanceTypesInstanceTypeResult> instanceTypes,

            string? outputFile,

            int totalCount)
        {
            Id = id;
            Ids = ids;
            InstanceTypes = instanceTypes;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}