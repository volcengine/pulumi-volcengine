// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Transit_router
{
    public static class RouteEntries
    {
        /// <summary>
        /// Use this data source to query detailed information of transit router route entries
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
        ///     var @default = Volcengine.Transit_router.RouteEntries.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "tr-rte-12b7qd5eo3h1c17q7y1sq5ixv",
        ///         },
        ///         TransitRouterRouteTableId = "tr-rtb-12b7qd3fmzf2817q7y2jkbd55",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<RouteEntriesResult> InvokeAsync(RouteEntriesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<RouteEntriesResult>("volcengine:transit_router/routeEntries:RouteEntries", args ?? new RouteEntriesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of transit router route entries
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
        ///     var @default = Volcengine.Transit_router.RouteEntries.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "tr-rte-12b7qd5eo3h1c17q7y1sq5ixv",
        ///         },
        ///         TransitRouterRouteTableId = "tr-rtb-12b7qd3fmzf2817q7y2jkbd55",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<RouteEntriesResult> Invoke(RouteEntriesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<RouteEntriesResult>("volcengine:transit_router/routeEntries:RouteEntries", args ?? new RouteEntriesInvokeArgs(), options.WithDefaults());
    }


    public sealed class RouteEntriesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The target network segment of the route entry.
        /// </summary>
        [Input("destinationCidrBlock")]
        public string? DestinationCidrBlock { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// The ids of the transit router route entry.
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

        /// <summary>
        /// The status of the route entry.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// The name of the route entry.
        /// </summary>
        [Input("transitRouterRouteEntryName")]
        public string? TransitRouterRouteEntryName { get; set; }

        /// <summary>
        /// The id of the route table.
        /// </summary>
        [Input("transitRouterRouteTableId", required: true)]
        public string TransitRouterRouteTableId { get; set; } = null!;

        public RouteEntriesArgs()
        {
        }
        public static new RouteEntriesArgs Empty => new RouteEntriesArgs();
    }

    public sealed class RouteEntriesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The target network segment of the route entry.
        /// </summary>
        [Input("destinationCidrBlock")]
        public Input<string>? DestinationCidrBlock { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// The ids of the transit router route entry.
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

        /// <summary>
        /// The status of the route entry.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The name of the route entry.
        /// </summary>
        [Input("transitRouterRouteEntryName")]
        public Input<string>? TransitRouterRouteEntryName { get; set; }

        /// <summary>
        /// The id of the route table.
        /// </summary>
        [Input("transitRouterRouteTableId", required: true)]
        public Input<string> TransitRouterRouteTableId { get; set; } = null!;

        public RouteEntriesInvokeArgs()
        {
        }
        public static new RouteEntriesInvokeArgs Empty => new RouteEntriesInvokeArgs();
    }


    [OutputType]
    public sealed class RouteEntriesResult
    {
        /// <summary>
        /// The target network segment of the route entry.
        /// </summary>
        public readonly string? DestinationCidrBlock;
        /// <summary>
        /// The list of route entries.
        /// </summary>
        public readonly ImmutableArray<Outputs.RouteEntriesEntryResult> Entries;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        /// <summary>
        /// The status of the route entry.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// The total count of data query.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The name of the route entry.
        /// </summary>
        public readonly string? TransitRouterRouteEntryName;
        public readonly string TransitRouterRouteTableId;

        [OutputConstructor]
        private RouteEntriesResult(
            string? destinationCidrBlock,

            ImmutableArray<Outputs.RouteEntriesEntryResult> entries,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            string? status,

            int totalCount,

            string? transitRouterRouteEntryName,

            string transitRouterRouteTableId)
        {
            DestinationCidrBlock = destinationCidrBlock;
            Entries = entries;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            Status = status;
            TotalCount = totalCount;
            TransitRouterRouteEntryName = transitRouterRouteEntryName;
            TransitRouterRouteTableId = transitRouterRouteTableId;
        }
    }
}