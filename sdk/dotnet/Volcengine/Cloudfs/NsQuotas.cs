// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Cloudfs
{
    public static class NsQuotas
    {
        /// <summary>
        /// Use this data source to query detailed information of cloudfs ns quotas
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
        ///     var @default = Volcengine.Cloudfs.NsQuotas.Invoke(new()
        ///     {
        ///         FsNames = new[]
        ///         {
        ///             "tffile",
        ///             "tftest2",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<NsQuotasResult> InvokeAsync(NsQuotasArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<NsQuotasResult>("volcengine:cloudfs/nsQuotas:NsQuotas", args ?? new NsQuotasArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cloudfs ns quotas
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
        ///     var @default = Volcengine.Cloudfs.NsQuotas.Invoke(new()
        ///     {
        ///         FsNames = new[]
        ///         {
        ///             "tffile",
        ///             "tftest2",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<NsQuotasResult> Invoke(NsQuotasInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<NsQuotasResult>("volcengine:cloudfs/nsQuotas:NsQuotas", args ?? new NsQuotasInvokeArgs(), options.WithDefaults());
    }


    public sealed class NsQuotasArgs : global::Pulumi.InvokeArgs
    {
        [Input("fsNames", required: true)]
        private List<string>? _fsNames;

        /// <summary>
        /// A list of fs name.
        /// </summary>
        public List<string> FsNames
        {
            get => _fsNames ?? (_fsNames = new List<string>());
            set => _fsNames = value;
        }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public NsQuotasArgs()
        {
        }
        public static new NsQuotasArgs Empty => new NsQuotasArgs();
    }

    public sealed class NsQuotasInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("fsNames", required: true)]
        private InputList<string>? _fsNames;

        /// <summary>
        /// A list of fs name.
        /// </summary>
        public InputList<string> FsNames
        {
            get => _fsNames ?? (_fsNames = new InputList<string>());
            set => _fsNames = value;
        }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public NsQuotasInvokeArgs()
        {
        }
        public static new NsQuotasInvokeArgs Empty => new NsQuotasInvokeArgs();
    }


    [OutputType]
    public sealed class NsQuotasResult
    {
        public readonly ImmutableArray<string> FsNames;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OutputFile;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.NsQuotasQuotaResult> Quotas;
        /// <summary>
        /// The total count of cloud fs quota query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private NsQuotasResult(
            ImmutableArray<string> fsNames,

            string id,

            string? outputFile,

            ImmutableArray<Outputs.NsQuotasQuotaResult> quotas,

            int totalCount)
        {
            FsNames = fsNames;
            Id = id;
            OutputFile = outputFile;
            Quotas = quotas;
            TotalCount = totalCount;
        }
    }
}