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
    /// Provides a resource to manage privatelink vpc endpoint service resource
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
    ///     var foo1 = new Volcengine.Clb.Clb("foo1", new()
    ///     {
    ///         Type = "public",
    ///         SubnetId = fooSubnet.Id,
    ///         LoadBalancerSpec = "small_1",
    ///         Description = "acc-test-demo",
    ///         LoadBalancerName = "acc-test-clb-new",
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
    ///     var fooVpcEndpointServiceResource = new Volcengine.Privatelink.VpcEndpointServiceResource("fooVpcEndpointServiceResource", new()
    ///     {
    ///         ServiceId = fooVpcEndpointService.Id,
    ///         ResourceId = foo1.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// VpcEndpointServiceResource can be imported using the serviceId:resourceId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource default epsvc-2fe630gurkl37k5gfuy33****:clb-bp1o94dp5i6ea****
    /// ```
    /// 
    ///  It is not recommended to use this resource for binding resources, it is recommended to use the resources field of vpc_endpoint_service for binding. If using this resource and vpc_endpoint_service jointly for operations, use lifecycle ignore_changes to suppress changes to the resources field in vpc_endpoint_service.
    /// </summary>
    [VolcengineResourceType("volcengine:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource")]
    public partial class VpcEndpointServiceResource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The id of resource. It is not recommended to use this resource for binding resources, it is recommended to use the resources field of vpc_endpoint_service for binding. If using this resource and vpc_endpoint_service jointly for operations, use lifecycle ignore_changes to suppress changes to the resources field in vpc_endpoint_service.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// The id of service.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;


        /// <summary>
        /// Create a VpcEndpointServiceResource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcEndpointServiceResource(string name, VpcEndpointServiceResourceArgs args, CustomResourceOptions? options = null)
            : base("volcengine:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource", name, args ?? new VpcEndpointServiceResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcEndpointServiceResource(string name, Input<string> id, VpcEndpointServiceResourceState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcEndpointServiceResource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcEndpointServiceResource Get(string name, Input<string> id, VpcEndpointServiceResourceState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcEndpointServiceResource(name, id, state, options);
        }
    }

    public sealed class VpcEndpointServiceResourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of resource. It is not recommended to use this resource for binding resources, it is recommended to use the resources field of vpc_endpoint_service for binding. If using this resource and vpc_endpoint_service jointly for operations, use lifecycle ignore_changes to suppress changes to the resources field in vpc_endpoint_service.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// The id of service.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public VpcEndpointServiceResourceArgs()
        {
        }
        public static new VpcEndpointServiceResourceArgs Empty => new VpcEndpointServiceResourceArgs();
    }

    public sealed class VpcEndpointServiceResourceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of resource. It is not recommended to use this resource for binding resources, it is recommended to use the resources field of vpc_endpoint_service for binding. If using this resource and vpc_endpoint_service jointly for operations, use lifecycle ignore_changes to suppress changes to the resources field in vpc_endpoint_service.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// The id of service.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        public VpcEndpointServiceResourceState()
        {
        }
        public static new VpcEndpointServiceResourceState Empty => new VpcEndpointServiceResourceState();
    }
}