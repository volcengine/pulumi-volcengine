// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Direct_connect
{
    public static class GetGateways
    {
        /// <summary>
        /// Use this data source to query detailed information of direct connect gateways
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
        ///     var foo = Volcengine.Direct_connect.GetGateways.Invoke(new()
        ///     {
        ///         DirectConnectGatewayName = "tf-test",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetGatewaysResult> InvokeAsync(GetGatewaysArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewaysResult>("volcengine:direct_connect/getGateways:getGateways", args ?? new GetGatewaysArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of direct connect gateways
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
        ///     var foo = Volcengine.Direct_connect.GetGateways.Invoke(new()
        ///     {
        ///         DirectConnectGatewayName = "tf-test",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetGatewaysResult> Invoke(GetGatewaysInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewaysResult>("volcengine:direct_connect/getGateways:getGateways", args ?? new GetGatewaysInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewaysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CEN ID which direct connect gateway belongs.
        /// </summary>
        [Input("cenId")]
        public string? CenId { get; set; }

        /// <summary>
        /// The direst connect gateway name.
        /// </summary>
        [Input("directConnectGatewayName")]
        public string? DirectConnectGatewayName { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

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

        [Input("tagFilters")]
        private List<Inputs.GetGatewaysTagFilterArgs>? _tagFilters;

        /// <summary>
        /// The filter tag of direct connect.
        /// </summary>
        public List<Inputs.GetGatewaysTagFilterArgs> TagFilters
        {
            get => _tagFilters ?? (_tagFilters = new List<Inputs.GetGatewaysTagFilterArgs>());
            set => _tagFilters = value;
        }

        public GetGatewaysArgs()
        {
        }
        public static new GetGatewaysArgs Empty => new GetGatewaysArgs();
    }

    public sealed class GetGatewaysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CEN ID which direct connect gateway belongs.
        /// </summary>
        [Input("cenId")]
        public Input<string>? CenId { get; set; }

        /// <summary>
        /// The direst connect gateway name.
        /// </summary>
        [Input("directConnectGatewayName")]
        public Input<string>? DirectConnectGatewayName { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

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

        [Input("tagFilters")]
        private InputList<Inputs.GetGatewaysTagFilterInputArgs>? _tagFilters;

        /// <summary>
        /// The filter tag of direct connect.
        /// </summary>
        public InputList<Inputs.GetGatewaysTagFilterInputArgs> TagFilters
        {
            get => _tagFilters ?? (_tagFilters = new InputList<Inputs.GetGatewaysTagFilterInputArgs>());
            set => _tagFilters = value;
        }

        public GetGatewaysInvokeArgs()
        {
        }
        public static new GetGatewaysInvokeArgs Empty => new GetGatewaysInvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewaysResult
    {
        /// <summary>
        /// The cen ID.
        /// </summary>
        public readonly string? CenId;
        /// <summary>
        /// The direct connect gateway name.
        /// </summary>
        public readonly string? DirectConnectGatewayName;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGatewaysDirectConnectGatewayResult> DirectConnectGateways;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.GetGatewaysTagFilterResult> TagFilters;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetGatewaysResult(
            string? cenId,

            string? directConnectGatewayName,

            ImmutableArray<Outputs.GetGatewaysDirectConnectGatewayResult> directConnectGateways,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            string? outputFile,

            ImmutableArray<Outputs.GetGatewaysTagFilterResult> tagFilters,

            int totalCount)
        {
            CenId = cenId;
            DirectConnectGatewayName = directConnectGatewayName;
            DirectConnectGateways = directConnectGateways;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            TagFilters = tagFilters;
            TotalCount = totalCount;
        }
    }
}
