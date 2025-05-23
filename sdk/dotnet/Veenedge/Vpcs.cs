// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Veenedge
{
    [Obsolete(@"volcengine.veenedge.Vpcs has been deprecated in favor of volcengine.veenedge.getVpcs")]
    public static class Vpcs
    {
        /// <summary>
        /// Use this data source to query detailed information of veenedge vpcs
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
        ///     var foo = Volcengine.Veenedge.GetVpcs.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<VpcsResult> InvokeAsync(VpcsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<VpcsResult>("volcengine:veenedge/vpcs:Vpcs", args ?? new VpcsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of veenedge vpcs
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
        ///     var foo = Volcengine.Veenedge.GetVpcs.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<VpcsResult> Invoke(VpcsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<VpcsResult>("volcengine:veenedge/vpcs:Vpcs", args ?? new VpcsInvokeArgs(), options.WithDefaults());
    }


    public sealed class VpcsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of vpc IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of Vpc.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public VpcsArgs()
        {
        }
        public static new VpcsArgs Empty => new VpcsArgs();
    }

    public sealed class VpcsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of vpc IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of Vpc.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public VpcsInvokeArgs()
        {
        }
        public static new VpcsInvokeArgs Empty => new VpcsInvokeArgs();
    }


    [OutputType]
    public sealed class VpcsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of Vpc query.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The collection of Vpc query.
        /// </summary>
        public readonly ImmutableArray<Outputs.VpcsVpcInstanceResult> VpcInstances;

        [OutputConstructor]
        private VpcsResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            string? outputFile,

            int totalCount,

            ImmutableArray<Outputs.VpcsVpcInstanceResult> vpcInstances)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            TotalCount = totalCount;
            VpcInstances = vpcInstances;
        }
    }
}
