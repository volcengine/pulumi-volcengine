// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Transit_router
{
    /// <summary>
    /// Provides a resource to manage transit router route table association
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Volcengine = Volcengine.Pulumi.Volcengine;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var foo = new Volcengine.Transit_router.RouteTableAssociation("foo", new()
    ///     {
    ///         TransitRouterAttachmentId = "tr-attach-im73ng3n5kao8gbssz2ddpuq",
    ///         TransitRouterRouteTableId = "tr-rtb-12b7qd3fmzf2817q7y2jkbd55",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// TransitRouterRouteTableAssociation can be imported using the TransitRouterAttachmentId:TransitRouterRouteTableId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:transit_router/routeTableAssociation:RouteTableAssociation default tr-attach-13n2l4c****:tr-rt-1i5i8khf9m58gae5kcx6****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:transit_router/routeTableAssociation:RouteTableAssociation")]
    public partial class RouteTableAssociation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the network instance connection.
        /// </summary>
        [Output("transitRouterAttachmentId")]
        public Output<string> TransitRouterAttachmentId { get; private set; } = null!;

        /// <summary>
        /// The ID of the routing table associated with the transit router instance.
        /// </summary>
        [Output("transitRouterRouteTableId")]
        public Output<string> TransitRouterRouteTableId { get; private set; } = null!;


        /// <summary>
        /// Create a RouteTableAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouteTableAssociation(string name, RouteTableAssociationArgs args, CustomResourceOptions? options = null)
            : base("volcengine:transit_router/routeTableAssociation:RouteTableAssociation", name, args ?? new RouteTableAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouteTableAssociation(string name, Input<string> id, RouteTableAssociationState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:transit_router/routeTableAssociation:RouteTableAssociation", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/volcengine",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RouteTableAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouteTableAssociation Get(string name, Input<string> id, RouteTableAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new RouteTableAssociation(name, id, state, options);
        }
    }

    public sealed class RouteTableAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the network instance connection.
        /// </summary>
        [Input("transitRouterAttachmentId", required: true)]
        public Input<string> TransitRouterAttachmentId { get; set; } = null!;

        /// <summary>
        /// The ID of the routing table associated with the transit router instance.
        /// </summary>
        [Input("transitRouterRouteTableId", required: true)]
        public Input<string> TransitRouterRouteTableId { get; set; } = null!;

        public RouteTableAssociationArgs()
        {
        }
        public static new RouteTableAssociationArgs Empty => new RouteTableAssociationArgs();
    }

    public sealed class RouteTableAssociationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the network instance connection.
        /// </summary>
        [Input("transitRouterAttachmentId")]
        public Input<string>? TransitRouterAttachmentId { get; set; }

        /// <summary>
        /// The ID of the routing table associated with the transit router instance.
        /// </summary>
        [Input("transitRouterRouteTableId")]
        public Input<string>? TransitRouterRouteTableId { get; set; }

        public RouteTableAssociationState()
        {
        }
        public static new RouteTableAssociationState Empty => new RouteTableAssociationState();
    }
}