// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Privatelink
{
    public static class VpcEndpointZones
    {
        /// <summary>
        /// Use this data source to query detailed information of privatelink vpc endpoint zones
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
        ///         var @default = Output.Create(Volcengine.Privatelink.VpcEndpointZones.InvokeAsync(new Volcengine.Privatelink.VpcEndpointZonesArgs
        ///         {
        ///             EndpointId = "ep-2byz5npiuu1hc2dx0efkv****",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<VpcEndpointZonesResult> InvokeAsync(VpcEndpointZonesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<VpcEndpointZonesResult>("volcengine:privatelink/vpcEndpointZones:VpcEndpointZones", args ?? new VpcEndpointZonesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of privatelink vpc endpoint zones
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
        ///         var @default = Output.Create(Volcengine.Privatelink.VpcEndpointZones.InvokeAsync(new Volcengine.Privatelink.VpcEndpointZonesArgs
        ///         {
        ///             EndpointId = "ep-2byz5npiuu1hc2dx0efkv****",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<VpcEndpointZonesResult> Invoke(VpcEndpointZonesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<VpcEndpointZonesResult>("volcengine:privatelink/vpcEndpointZones:VpcEndpointZones", args ?? new VpcEndpointZonesInvokeArgs(), options.WithDefaults());
    }


    public sealed class VpcEndpointZonesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The endpoint id of query.
        /// </summary>
        [Input("endpointId")]
        public string? EndpointId { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public VpcEndpointZonesArgs()
        {
        }
    }

    public sealed class VpcEndpointZonesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The endpoint id of query.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public VpcEndpointZonesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class VpcEndpointZonesResult
    {
        public readonly string? EndpointId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OutputFile;
        /// <summary>
        /// Returns the total amount of the data list.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.VpcEndpointZonesVpcEndpointZoneResult> VpcEndpointZones;

        [OutputConstructor]
        private VpcEndpointZonesResult(
            string? endpointId,

            string id,

            string? outputFile,

            int totalCount,

            ImmutableArray<Outputs.VpcEndpointZonesVpcEndpointZoneResult> vpcEndpointZones)
        {
            EndpointId = endpointId;
            Id = id;
            OutputFile = outputFile;
            TotalCount = totalCount;
            VpcEndpointZones = vpcEndpointZones;
        }
    }
}