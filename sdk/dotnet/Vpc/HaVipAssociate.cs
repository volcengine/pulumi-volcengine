// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpc
{
    /// <summary>
    /// Provides a resource to manage ha vip associate
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
    ///     var fooZones = Volcengine.Ecs.GetZones.Invoke();
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
    ///         ZoneId = fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         VpcId = fooVpc.Id,
    ///     });
    /// 
    ///     var fooSecurityGroup = new Volcengine.Vpc.SecurityGroup("fooSecurityGroup", new()
    ///     {
    ///         SecurityGroupName = "acc-test-sg",
    ///         VpcId = fooVpc.Id,
    ///     });
    /// 
    ///     var fooNetworkInterface = new Volcengine.Vpc.NetworkInterface("fooNetworkInterface", new()
    ///     {
    ///         NetworkInterfaceName = "acc-test-eni",
    ///         Description = "acc-test",
    ///         SubnetId = fooSubnet.Id,
    ///         SecurityGroupIds = new[]
    ///         {
    ///             fooSecurityGroup.Id,
    ///         },
    ///         PrimaryIpAddress = "172.16.0.253",
    ///         PortSecurityEnabled = false,
    ///         PrivateIpAddresses = new[]
    ///         {
    ///             "172.16.0.2",
    ///         },
    ///         ProjectName = "default",
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Vpc.Inputs.NetworkInterfaceTagArgs
    ///             {
    ///                 Key = "k1",
    ///                 Value = "v1",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var fooHaVip = new Volcengine.Vpc.HaVip("fooHaVip", new()
    ///     {
    ///         HaVipName = "acc-test-ha-vip",
    ///         Description = "acc-test",
    ///         SubnetId = fooSubnet.Id,
    ///         IpAddress = "172.16.0.5",
    ///     });
    /// 
    ///     var fooHaVipAssociate = new Volcengine.Vpc.HaVipAssociate("fooHaVipAssociate", new()
    ///     {
    ///         HaVipId = fooHaVip.Id,
    ///         InstanceType = "NetworkInterface",
    ///         InstanceId = fooNetworkInterface.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// HaVipAssociate can be imported using the ha_vip_id:instance_id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:vpc/haVipAssociate:HaVipAssociate default havip-2byzv8icq1b7k2dx0eegb****:eni-2d5wv84h7onpc58ozfeeu****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:vpc/haVipAssociate:HaVipAssociate")]
    public partial class HaVipAssociate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The id of the Ha Vip.
        /// </summary>
        [Output("haVipId")]
        public Output<string> HaVipId { get; private set; } = null!;

        /// <summary>
        /// The id of the associated instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The type of the associated instance. Valid values: `EcsInstance`, `NetworkInterface`.
        /// </summary>
        [Output("instanceType")]
        public Output<string?> InstanceType { get; private set; } = null!;


        /// <summary>
        /// Create a HaVipAssociate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HaVipAssociate(string name, HaVipAssociateArgs args, CustomResourceOptions? options = null)
            : base("volcengine:vpc/haVipAssociate:HaVipAssociate", name, args ?? new HaVipAssociateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HaVipAssociate(string name, Input<string> id, HaVipAssociateState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:vpc/haVipAssociate:HaVipAssociate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HaVipAssociate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HaVipAssociate Get(string name, Input<string> id, HaVipAssociateState? state = null, CustomResourceOptions? options = null)
        {
            return new HaVipAssociate(name, id, state, options);
        }
    }

    public sealed class HaVipAssociateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the Ha Vip.
        /// </summary>
        [Input("haVipId", required: true)]
        public Input<string> HaVipId { get; set; } = null!;

        /// <summary>
        /// The id of the associated instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The type of the associated instance. Valid values: `EcsInstance`, `NetworkInterface`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        public HaVipAssociateArgs()
        {
        }
        public static new HaVipAssociateArgs Empty => new HaVipAssociateArgs();
    }

    public sealed class HaVipAssociateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the Ha Vip.
        /// </summary>
        [Input("haVipId")]
        public Input<string>? HaVipId { get; set; }

        /// <summary>
        /// The id of the associated instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The type of the associated instance. Valid values: `EcsInstance`, `NetworkInterface`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        public HaVipAssociateState()
        {
        }
        public static new HaVipAssociateState Empty => new HaVipAssociateState();
    }
}
