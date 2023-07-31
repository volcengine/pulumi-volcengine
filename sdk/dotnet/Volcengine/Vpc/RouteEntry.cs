// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Vpc
{
    /// <summary>
    /// Provides a resource to manage route entry
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Volcengine = Volcengine.PulumiPackage.Volcengine;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Volcengine.Vpc.RouteEntry("foo", new Volcengine.Vpc.RouteEntryArgs
    ///         {
    ///             Description = "tf-test-up",
    ///             DestinationCidrBlock = "0.0.0.0/2",
    ///             NextHopId = "ngw-274gwbqe340zk7fap8spkzo7x",
    ///             NextHopType = "NatGW",
    ///             RouteEntryName = "tf-test-up",
    ///             RouteTableId = "vtb-2744hslq5b7r47fap8tjomgnj",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Route entry can be imported using the route_table_id:route_entry_id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:vpc/routeEntry:RouteEntry default vtb-274e19skkuhog7fap8u4i8ird:rte-274e1g9ei4k5c7fap8sp974fq
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:vpc/routeEntry:RouteEntry")]
    public partial class RouteEntry : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the route entry.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The destination CIDR block of the route entry.
        /// </summary>
        [Output("destinationCidrBlock")]
        public Output<string> DestinationCidrBlock { get; private set; } = null!;

        /// <summary>
        /// The id of the next hop.
        /// </summary>
        [Output("nextHopId")]
        public Output<string> NextHopId { get; private set; } = null!;

        /// <summary>
        /// The type of the next hop, Optional choice contains `Instance`, `NetworkInterface`, `NatGW`, `VpnGW`, `TransitRouter`.
        /// </summary>
        [Output("nextHopType")]
        public Output<string> NextHopType { get; private set; } = null!;

        /// <summary>
        /// The id of the route entry.
        /// </summary>
        [Output("routeEntryId")]
        public Output<string> RouteEntryId { get; private set; } = null!;

        /// <summary>
        /// The name of the route entry.
        /// </summary>
        [Output("routeEntryName")]
        public Output<string?> RouteEntryName { get; private set; } = null!;

        /// <summary>
        /// The id of the route table.
        /// </summary>
        [Output("routeTableId")]
        public Output<string> RouteTableId { get; private set; } = null!;

        /// <summary>
        /// The description of the route entry.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a RouteEntry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouteEntry(string name, RouteEntryArgs args, CustomResourceOptions? options = null)
            : base("volcengine:vpc/routeEntry:RouteEntry", name, args ?? new RouteEntryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouteEntry(string name, Input<string> id, RouteEntryState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:vpc/routeEntry:RouteEntry", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RouteEntry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouteEntry Get(string name, Input<string> id, RouteEntryState? state = null, CustomResourceOptions? options = null)
        {
            return new RouteEntry(name, id, state, options);
        }
    }

    public sealed class RouteEntryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the route entry.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination CIDR block of the route entry.
        /// </summary>
        [Input("destinationCidrBlock", required: true)]
        public Input<string> DestinationCidrBlock { get; set; } = null!;

        /// <summary>
        /// The id of the next hop.
        /// </summary>
        [Input("nextHopId", required: true)]
        public Input<string> NextHopId { get; set; } = null!;

        /// <summary>
        /// The type of the next hop, Optional choice contains `Instance`, `NetworkInterface`, `NatGW`, `VpnGW`, `TransitRouter`.
        /// </summary>
        [Input("nextHopType", required: true)]
        public Input<string> NextHopType { get; set; } = null!;

        /// <summary>
        /// The name of the route entry.
        /// </summary>
        [Input("routeEntryName")]
        public Input<string>? RouteEntryName { get; set; }

        /// <summary>
        /// The id of the route table.
        /// </summary>
        [Input("routeTableId", required: true)]
        public Input<string> RouteTableId { get; set; } = null!;

        public RouteEntryArgs()
        {
        }
    }

    public sealed class RouteEntryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the route entry.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination CIDR block of the route entry.
        /// </summary>
        [Input("destinationCidrBlock")]
        public Input<string>? DestinationCidrBlock { get; set; }

        /// <summary>
        /// The id of the next hop.
        /// </summary>
        [Input("nextHopId")]
        public Input<string>? NextHopId { get; set; }

        /// <summary>
        /// The type of the next hop, Optional choice contains `Instance`, `NetworkInterface`, `NatGW`, `VpnGW`, `TransitRouter`.
        /// </summary>
        [Input("nextHopType")]
        public Input<string>? NextHopType { get; set; }

        /// <summary>
        /// The id of the route entry.
        /// </summary>
        [Input("routeEntryId")]
        public Input<string>? RouteEntryId { get; set; }

        /// <summary>
        /// The name of the route entry.
        /// </summary>
        [Input("routeEntryName")]
        public Input<string>? RouteEntryName { get; set; }

        /// <summary>
        /// The id of the route table.
        /// </summary>
        [Input("routeTableId")]
        public Input<string>? RouteTableId { get; set; }

        /// <summary>
        /// The description of the route entry.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public RouteEntryState()
        {
        }
    }
}