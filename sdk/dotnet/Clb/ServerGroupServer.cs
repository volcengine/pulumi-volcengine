// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Clb
{
    /// <summary>
    /// Provides a resource to manage server group server
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
    ///         Description = "acc0Demo",
    ///         LoadBalancerName = "acc-test-create",
    ///         EipBillingConfig = new Volcengine.Clb.Inputs.ClbEipBillingConfigArgs
    ///         {
    ///             Isp = "BGP",
    ///             EipBillingType = "PostPaidByBandwidth",
    ///             Bandwidth = 1,
    ///         },
    ///     });
    /// 
    ///     var fooServerGroup = new Volcengine.Clb.ServerGroup("fooServerGroup", new()
    ///     {
    ///         LoadBalancerId = fooClb.Id,
    ///         ServerGroupName = "acc-test-create",
    ///         Description = "hello demo11",
    ///     });
    /// 
    ///     var fooSecurityGroup = new Volcengine.Vpc.SecurityGroup("fooSecurityGroup", new()
    ///     {
    ///         VpcId = fooVpc.Id,
    ///         SecurityGroupName = "acc-test-security-group",
    ///     });
    /// 
    ///     var fooInstance = new Volcengine.Ecs.Instance("fooInstance", new()
    ///     {
    ///         ImageId = "image-ycjwwciuzy5pkh54xx8f",
    ///         InstanceType = "ecs.c3i.large",
    ///         InstanceName = "acc-test-ecs-name",
    ///         Password = "93f0cb0614Aab12",
    ///         InstanceChargeType = "PostPaid",
    ///         SystemVolumeType = "ESSD_PL0",
    ///         SystemVolumeSize = 40,
    ///         SubnetId = fooSubnet.Id,
    ///         SecurityGroupIds = new[]
    ///         {
    ///             fooSecurityGroup.Id,
    ///         },
    ///     });
    /// 
    ///     var fooServerGroupServer = new Volcengine.Clb.ServerGroupServer("fooServerGroupServer", new()
    ///     {
    ///         ServerGroupId = fooServerGroup.Id,
    ///         InstanceId = fooInstance.Id,
    ///         Type = "ecs",
    ///         Weight = 100,
    ///         Port = 80,
    ///         Description = "This is a acc test server",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ServerGroupServer can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:clb/serverGroupServer:ServerGroupServer default rsp-274xltv2*****8tlv3q3s:rs-3ciynux6i1x4w****rszh49sj
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:clb/serverGroupServer:ServerGroupServer")]
    public partial class ServerGroupServer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the instance.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of ecs instance or the network card bound to ecs instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The private ip of the instance.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// The port receiving request.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// The ID of the ServerGroup.
        /// </summary>
        [Output("serverGroupId")]
        public Output<string> ServerGroupId { get; private set; } = null!;

        /// <summary>
        /// The server id of instance in ServerGroup.
        /// </summary>
        [Output("serverId")]
        public Output<string> ServerId { get; private set; } = null!;

        /// <summary>
        /// The type of instance. Optional choice contains `ecs`, `eni`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The weight of the instance, range in 0~100.
        /// </summary>
        [Output("weight")]
        public Output<int?> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a ServerGroupServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerGroupServer(string name, ServerGroupServerArgs args, CustomResourceOptions? options = null)
            : base("volcengine:clb/serverGroupServer:ServerGroupServer", name, args ?? new ServerGroupServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerGroupServer(string name, Input<string> id, ServerGroupServerState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:clb/serverGroupServer:ServerGroupServer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServerGroupServer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerGroupServer Get(string name, Input<string> id, ServerGroupServerState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerGroupServer(name, id, state, options);
        }
    }

    public sealed class ServerGroupServerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of ecs instance or the network card bound to ecs instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The private ip of the instance.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The port receiving request.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The ID of the ServerGroup.
        /// </summary>
        [Input("serverGroupId", required: true)]
        public Input<string> ServerGroupId { get; set; } = null!;

        /// <summary>
        /// The type of instance. Optional choice contains `ecs`, `eni`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The weight of the instance, range in 0~100.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public ServerGroupServerArgs()
        {
        }
        public static new ServerGroupServerArgs Empty => new ServerGroupServerArgs();
    }

    public sealed class ServerGroupServerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the instance.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of ecs instance or the network card bound to ecs instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The private ip of the instance.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The port receiving request.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The ID of the ServerGroup.
        /// </summary>
        [Input("serverGroupId")]
        public Input<string>? ServerGroupId { get; set; }

        /// <summary>
        /// The server id of instance in ServerGroup.
        /// </summary>
        [Input("serverId")]
        public Input<string>? ServerId { get; set; }

        /// <summary>
        /// The type of instance. Optional choice contains `ecs`, `eni`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The weight of the instance, range in 0~100.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public ServerGroupServerState()
        {
        }
        public static new ServerGroupServerState Empty => new ServerGroupServerState();
    }
}