// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Privatelink
{
    /// <summary>
    /// Provides a resource to manage privatelink vpc endpoint zone
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
    ///     var fooZones = Volcengine.Ecs.Zones.Invoke();
    /// 
    ///     var fooVpc = new Volcengine.Vpc.Vpc("fooVpc", new()
    ///     {
    ///         VpcName = "acc-test-vpc",
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var fooSubnet = new Volcengine.Vpc.Subnet("fooSubnet", new()
    ///     {
    ///         SubnetName = "acc-test-subnet",
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[0]?.Id),
    ///         VpcId = fooVpc.Id,
    ///     });
    /// 
    ///     var fooSecurityGroup = new Volcengine.Vpc.SecurityGroup("fooSecurityGroup", new()
    ///     {
    ///         SecurityGroupName = "acc-test-security-group",
    ///         VpcId = fooVpc.Id,
    ///     });
    /// 
    ///     var fooClb = new Volcengine.Clb.Clb("fooClb", new()
    ///     {
    ///         Type = "public",
    ///         SubnetId = fooSubnet.Id,
    ///         LoadBalancerSpec = "small_1",
    ///         Description = "acc-test-demo",
    ///         LoadBalancerName = "acc-test-clb",
    ///         LoadBalancerBillingType = "PostPaid",
    ///         EipBillingConfig = new Volcengine.Clb.Inputs.ClbEipBillingConfigArgs
    ///         {
    ///             Isp = "BGP",
    ///             EipBillingType = "PostPaidByBandwidth",
    ///             Bandwidth = 1,
    ///         },
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Clb.Inputs.ClbTagArgs
    ///             {
    ///                 Key = "k1",
    ///                 Value = "v1",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var fooVpcEndpointService = new Volcengine.Privatelink.VpcEndpointService("fooVpcEndpointService", new()
    ///     {
    ///         Resources = new[]
    ///         {
    ///             new Volcengine.Privatelink.Inputs.VpcEndpointServiceResourceArgs
    ///             {
    ///                 ResourceId = fooClb.Id,
    ///                 ResourceType = "CLB",
    ///             },
    ///         },
    ///         Description = "acc-test",
    ///         AutoAcceptEnabled = true,
    ///     });
    /// 
    ///     var fooVpcEndpoint = new Volcengine.Privatelink.VpcEndpoint("fooVpcEndpoint", new()
    ///     {
    ///         SecurityGroupIds = new[]
    ///         {
    ///             fooSecurityGroup.Id,
    ///         },
    ///         ServiceId = fooVpcEndpointService.Id,
    ///         EndpointName = "acc-test-ep",
    ///         Description = "acc-test",
    ///     });
    /// 
    ///     var fooVpcEndpointZone = new Volcengine.Privatelink.VpcEndpointZone("fooVpcEndpointZone", new()
    ///     {
    ///         EndpointId = fooVpcEndpoint.Id,
    ///         SubnetId = fooSubnet.Id,
    ///         PrivateIpAddress = "172.16.0.251",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// VpcEndpointZone can be imported using the endpointId:subnetId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:privatelink/vpcEndpointZone:VpcEndpointZone default ep-3rel75r081l345zsk2i59****:subnet-2bz47q19zhx4w2dx0eevn****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:privatelink/vpcEndpointZone:VpcEndpointZone")]
    public partial class VpcEndpointZone : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The endpoint id of vpc endpoint zone.
        /// </summary>
        [Output("endpointId")]
        public Output<string> EndpointId { get; private set; } = null!;

        /// <summary>
        /// The network interface id of vpc endpoint.
        /// </summary>
        [Output("networkInterfaceId")]
        public Output<string> NetworkInterfaceId { get; private set; } = null!;

        /// <summary>
        /// The private ip address of vpc endpoint zone.
        /// </summary>
        [Output("privateIpAddress")]
        public Output<string> PrivateIpAddress { get; private set; } = null!;

        /// <summary>
        /// The subnet id of vpc endpoint zone.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// The domain of vpc endpoint zone.
        /// </summary>
        [Output("zoneDomain")]
        public Output<string> ZoneDomain { get; private set; } = null!;

        /// <summary>
        /// The Id of vpc endpoint zone.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;

        /// <summary>
        /// The status of vpc endpoint zone.
        /// </summary>
        [Output("zoneStatus")]
        public Output<string> ZoneStatus { get; private set; } = null!;


        /// <summary>
        /// Create a VpcEndpointZone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcEndpointZone(string name, VpcEndpointZoneArgs args, CustomResourceOptions? options = null)
            : base("volcengine:privatelink/vpcEndpointZone:VpcEndpointZone", name, args ?? new VpcEndpointZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcEndpointZone(string name, Input<string> id, VpcEndpointZoneState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:privatelink/vpcEndpointZone:VpcEndpointZone", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcEndpointZone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcEndpointZone Get(string name, Input<string> id, VpcEndpointZoneState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcEndpointZone(name, id, state, options);
        }
    }

    public sealed class VpcEndpointZoneArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The endpoint id of vpc endpoint zone.
        /// </summary>
        [Input("endpointId", required: true)]
        public Input<string> EndpointId { get; set; } = null!;

        /// <summary>
        /// The private ip address of vpc endpoint zone.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        /// <summary>
        /// The subnet id of vpc endpoint zone.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public VpcEndpointZoneArgs()
        {
        }
        public static new VpcEndpointZoneArgs Empty => new VpcEndpointZoneArgs();
    }

    public sealed class VpcEndpointZoneState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The endpoint id of vpc endpoint zone.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// The network interface id of vpc endpoint.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// The private ip address of vpc endpoint zone.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        /// <summary>
        /// The subnet id of vpc endpoint zone.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// The domain of vpc endpoint zone.
        /// </summary>
        [Input("zoneDomain")]
        public Input<string>? ZoneDomain { get; set; }

        /// <summary>
        /// The Id of vpc endpoint zone.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        /// <summary>
        /// The status of vpc endpoint zone.
        /// </summary>
        [Input("zoneStatus")]
        public Input<string>? ZoneStatus { get; set; }

        public VpcEndpointZoneState()
        {
        }
        public static new VpcEndpointZoneState Empty => new VpcEndpointZoneState();
    }
}