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
    ///     // ipv4 public clb
    ///     var publicClb = new Volcengine.Clb.Clb("publicClb", new()
    ///     {
    ///         Type = "public",
    ///         SubnetId = fooSubnet.Id,
    ///         LoadBalancerName = "acc-test-clb-public",
    ///         LoadBalancerSpec = "small_1",
    ///         Description = "acc-test-demo",
    ///         ProjectName = "default",
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
    ///     // ipv4 private clb
    ///     var privateClb = new Volcengine.Clb.Clb("privateClb", new()
    ///     {
    ///         Type = "private",
    ///         SubnetId = fooSubnet.Id,
    ///         LoadBalancerName = "acc-test-clb-private",
    ///         LoadBalancerSpec = "small_1",
    ///         Description = "acc-test-demo",
    ///         ProjectName = "default",
    ///     });
    /// 
    ///     var eip = new Volcengine.Eip.Address("eip", new()
    ///     {
    ///         BillingType = "PostPaidByBandwidth",
    ///         Bandwidth = 1,
    ///         Isp = "BGP",
    ///         Description = "tf-test",
    ///         ProjectName = "default",
    ///     });
    /// 
    ///     var associate = new Volcengine.Eip.Associate("associate", new()
    ///     {
    ///         AllocationId = eip.Id,
    ///         InstanceId = privateClb.Id,
    ///         InstanceType = "ClbInstance",
    ///     });
    /// 
    ///     // ipv6 private clb
    ///     var vpcIpv6 = new Volcengine.Vpc.Vpc("vpcIpv6", new()
    ///     {
    ///         VpcName = "acc-test-vpc-ipv6",
    ///         CidrBlock = "172.16.0.0/16",
    ///         EnableIpv6 = true,
    ///     });
    /// 
    ///     var subnetIpv6 = new Volcengine.Vpc.Subnet("subnetIpv6", new()
    ///     {
    ///         SubnetName = "acc-test-subnet-ipv6",
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[1]?.Id),
    ///         VpcId = vpcIpv6.Id,
    ///         Ipv6CidrBlock = 1,
    ///     });
    /// 
    ///     var privateClbIpv6 = new Volcengine.Clb.Clb("privateClbIpv6", new()
    ///     {
    ///         Type = "private",
    ///         SubnetId = subnetIpv6.Id,
    ///         LoadBalancerName = "acc-test-clb-ipv6",
    ///         LoadBalancerSpec = "small_1",
    ///         Description = "acc-test-demo",
    ///         ProjectName = "default",
    ///         AddressIpVersion = "DualStack",
    ///     });
    /// 
    ///     var ipv6Gateway = new Volcengine.Vpc.Ipv6Gateway("ipv6Gateway", new()
    ///     {
    ///         VpcId = vpcIpv6.Id,
    ///     });
    /// 
    ///     var fooIpv6AddressBandwidth = new Volcengine.Vpc.Ipv6AddressBandwidth("fooIpv6AddressBandwidth", new()
    ///     {
    ///         Ipv6Address = privateClbIpv6.EniIpv6Address,
    ///         BillingType = "PostPaidByBandwidth",
    ///         Bandwidth = 5,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             ipv6Gateway,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CLB can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:clb/clb:Clb default clb-273y2ok6ets007fap8txvf6us
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:clb/clb:Clb")]
    public partial class Clb : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The address ip version of the Clb. Valid values: `ipv4`, `DualStack`. Default is `ipv4`.
        /// When the value of this field is `DualStack`, the type of the CLB must be `private`, and suggest using a combination of resource `volcengine.vpc.Ipv6Gateway` and `volcengine.vpc.Ipv6AddressBandwidth` to achieve ipv6 public network access function.
        /// </summary>
        [Output("addressIpVersion")]
        public Output<string?> AddressIpVersion { get; private set; } = null!;

        /// <summary>
        /// The description of the CLB.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Eip address of the Clb.
        /// </summary>
        [Output("eipAddress")]
        public Output<string> EipAddress { get; private set; } = null!;

        /// <summary>
        /// The billing configuration of the EIP which automatically associated to CLB. This field is valid when the type of CLB is `public`.When the type of the CLB is `private`, suggest using a combination of resource `volcengine.eip.Address` and `volcengine.eip.Associate` to achieve public network access function.
        /// </summary>
        [Output("eipBillingConfig")]
        public Output<Outputs.ClbEipBillingConfig> EipBillingConfig { get; private set; } = null!;

        /// <summary>
        /// The Eip ID of the Clb.
        /// </summary>
        [Output("eipId")]
        public Output<string> EipId { get; private set; } = null!;

        /// <summary>
        /// The eni address of the CLB.
        /// </summary>
        [Output("eniAddress")]
        public Output<string> EniAddress { get; private set; } = null!;

        /// <summary>
        /// The eni ipv6 address of the Clb.
        /// </summary>
        [Output("eniIpv6Address")]
        public Output<string> EniIpv6Address { get; private set; } = null!;

        /// <summary>
        /// The Ipv6 Eip ID of the Clb.
        /// </summary>
        [Output("ipv6EipId")]
        public Output<string> Ipv6EipId { get; private set; } = null!;

        /// <summary>
        /// The billing type of the CLB, the value can be `PostPaid` or `PrePaid`.
        /// </summary>
        [Output("loadBalancerBillingType")]
        public Output<string> LoadBalancerBillingType { get; private set; } = null!;

        /// <summary>
        /// The name of the CLB.
        /// </summary>
        [Output("loadBalancerName")]
        public Output<string> LoadBalancerName { get; private set; } = null!;

        /// <summary>
        /// The specification of the CLB, the value can be `small_1`, `small_2`, `medium_1`, `medium_2`, `large_1`, `large_2`.
        /// </summary>
        [Output("loadBalancerSpec")]
        public Output<string> LoadBalancerSpec { get; private set; } = null!;

        /// <summary>
        /// The master zone ID of the CLB.
        /// </summary>
        [Output("masterZoneId")]
        public Output<string> MasterZoneId { get; private set; } = null!;

        /// <summary>
        /// The reason of the console modification protection.
        /// </summary>
        [Output("modificationProtectionReason")]
        public Output<string?> ModificationProtectionReason { get; private set; } = null!;

        /// <summary>
        /// The status of the console modification protection, the value can be `NonProtection` or `ConsoleProtection`.
        /// </summary>
        [Output("modificationProtectionStatus")]
        public Output<string?> ModificationProtectionStatus { get; private set; } = null!;

        /// <summary>
        /// The period of the NatGateway, the valid value range in 1~9 or 12 or 24 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid NatGateway. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// The ProjectName of the CLB.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The region of the request.
        /// </summary>
        [Output("regionId")]
        public Output<string> RegionId { get; private set; } = null!;

        /// <summary>
        /// The renew type of the CLB. When the value of the load_balancer_billing_type is `PrePaid`, the query returns this field.
        /// </summary>
        [Output("renewType")]
        public Output<string> RenewType { get; private set; } = null!;

        /// <summary>
        /// The slave zone ID of the CLB.
        /// </summary>
        [Output("slaveZoneId")]
        public Output<string> SlaveZoneId { get; private set; } = null!;

        /// <summary>
        /// The id of the Subnet.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ClbTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the CLB. And optional choice contains `public` or `private`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The id of the VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Clb resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Clb(string name, ClbArgs args, CustomResourceOptions? options = null)
            : base("volcengine:clb/clb:Clb", name, args ?? new ClbArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Clb(string name, Input<string> id, ClbState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:clb/clb:Clb", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Clb resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Clb Get(string name, Input<string> id, ClbState? state = null, CustomResourceOptions? options = null)
        {
            return new Clb(name, id, state, options);
        }
    }

    public sealed class ClbArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The address ip version of the Clb. Valid values: `ipv4`, `DualStack`. Default is `ipv4`.
        /// When the value of this field is `DualStack`, the type of the CLB must be `private`, and suggest using a combination of resource `volcengine.vpc.Ipv6Gateway` and `volcengine.vpc.Ipv6AddressBandwidth` to achieve ipv6 public network access function.
        /// </summary>
        [Input("addressIpVersion")]
        public Input<string>? AddressIpVersion { get; set; }

        /// <summary>
        /// The description of the CLB.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The billing configuration of the EIP which automatically associated to CLB. This field is valid when the type of CLB is `public`.When the type of the CLB is `private`, suggest using a combination of resource `volcengine.eip.Address` and `volcengine.eip.Associate` to achieve public network access function.
        /// </summary>
        [Input("eipBillingConfig")]
        public Input<Inputs.ClbEipBillingConfigArgs>? EipBillingConfig { get; set; }

        /// <summary>
        /// The eni address of the CLB.
        /// </summary>
        [Input("eniAddress")]
        public Input<string>? EniAddress { get; set; }

        /// <summary>
        /// The eni ipv6 address of the Clb.
        /// </summary>
        [Input("eniIpv6Address")]
        public Input<string>? EniIpv6Address { get; set; }

        /// <summary>
        /// The billing type of the CLB, the value can be `PostPaid` or `PrePaid`.
        /// </summary>
        [Input("loadBalancerBillingType")]
        public Input<string>? LoadBalancerBillingType { get; set; }

        /// <summary>
        /// The name of the CLB.
        /// </summary>
        [Input("loadBalancerName")]
        public Input<string>? LoadBalancerName { get; set; }

        /// <summary>
        /// The specification of the CLB, the value can be `small_1`, `small_2`, `medium_1`, `medium_2`, `large_1`, `large_2`.
        /// </summary>
        [Input("loadBalancerSpec", required: true)]
        public Input<string> LoadBalancerSpec { get; set; } = null!;

        /// <summary>
        /// The master zone ID of the CLB.
        /// </summary>
        [Input("masterZoneId")]
        public Input<string>? MasterZoneId { get; set; }

        /// <summary>
        /// The reason of the console modification protection.
        /// </summary>
        [Input("modificationProtectionReason")]
        public Input<string>? ModificationProtectionReason { get; set; }

        /// <summary>
        /// The status of the console modification protection, the value can be `NonProtection` or `ConsoleProtection`.
        /// </summary>
        [Input("modificationProtectionStatus")]
        public Input<string>? ModificationProtectionStatus { get; set; }

        /// <summary>
        /// The period of the NatGateway, the valid value range in 1~9 or 12 or 24 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid NatGateway. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The ProjectName of the CLB.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The region of the request.
        /// </summary>
        [Input("regionId")]
        public Input<string>? RegionId { get; set; }

        /// <summary>
        /// The slave zone ID of the CLB.
        /// </summary>
        [Input("slaveZoneId")]
        public Input<string>? SlaveZoneId { get; set; }

        /// <summary>
        /// The id of the Subnet.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.ClbTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.ClbTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ClbTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the CLB. And optional choice contains `public` or `private`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The id of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public ClbArgs()
        {
        }
        public static new ClbArgs Empty => new ClbArgs();
    }

    public sealed class ClbState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The address ip version of the Clb. Valid values: `ipv4`, `DualStack`. Default is `ipv4`.
        /// When the value of this field is `DualStack`, the type of the CLB must be `private`, and suggest using a combination of resource `volcengine.vpc.Ipv6Gateway` and `volcengine.vpc.Ipv6AddressBandwidth` to achieve ipv6 public network access function.
        /// </summary>
        [Input("addressIpVersion")]
        public Input<string>? AddressIpVersion { get; set; }

        /// <summary>
        /// The description of the CLB.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Eip address of the Clb.
        /// </summary>
        [Input("eipAddress")]
        public Input<string>? EipAddress { get; set; }

        /// <summary>
        /// The billing configuration of the EIP which automatically associated to CLB. This field is valid when the type of CLB is `public`.When the type of the CLB is `private`, suggest using a combination of resource `volcengine.eip.Address` and `volcengine.eip.Associate` to achieve public network access function.
        /// </summary>
        [Input("eipBillingConfig")]
        public Input<Inputs.ClbEipBillingConfigGetArgs>? EipBillingConfig { get; set; }

        /// <summary>
        /// The Eip ID of the Clb.
        /// </summary>
        [Input("eipId")]
        public Input<string>? EipId { get; set; }

        /// <summary>
        /// The eni address of the CLB.
        /// </summary>
        [Input("eniAddress")]
        public Input<string>? EniAddress { get; set; }

        /// <summary>
        /// The eni ipv6 address of the Clb.
        /// </summary>
        [Input("eniIpv6Address")]
        public Input<string>? EniIpv6Address { get; set; }

        /// <summary>
        /// The Ipv6 Eip ID of the Clb.
        /// </summary>
        [Input("ipv6EipId")]
        public Input<string>? Ipv6EipId { get; set; }

        /// <summary>
        /// The billing type of the CLB, the value can be `PostPaid` or `PrePaid`.
        /// </summary>
        [Input("loadBalancerBillingType")]
        public Input<string>? LoadBalancerBillingType { get; set; }

        /// <summary>
        /// The name of the CLB.
        /// </summary>
        [Input("loadBalancerName")]
        public Input<string>? LoadBalancerName { get; set; }

        /// <summary>
        /// The specification of the CLB, the value can be `small_1`, `small_2`, `medium_1`, `medium_2`, `large_1`, `large_2`.
        /// </summary>
        [Input("loadBalancerSpec")]
        public Input<string>? LoadBalancerSpec { get; set; }

        /// <summary>
        /// The master zone ID of the CLB.
        /// </summary>
        [Input("masterZoneId")]
        public Input<string>? MasterZoneId { get; set; }

        /// <summary>
        /// The reason of the console modification protection.
        /// </summary>
        [Input("modificationProtectionReason")]
        public Input<string>? ModificationProtectionReason { get; set; }

        /// <summary>
        /// The status of the console modification protection, the value can be `NonProtection` or `ConsoleProtection`.
        /// </summary>
        [Input("modificationProtectionStatus")]
        public Input<string>? ModificationProtectionStatus { get; set; }

        /// <summary>
        /// The period of the NatGateway, the valid value range in 1~9 or 12 or 24 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid NatGateway. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The ProjectName of the CLB.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The region of the request.
        /// </summary>
        [Input("regionId")]
        public Input<string>? RegionId { get; set; }

        /// <summary>
        /// The renew type of the CLB. When the value of the load_balancer_billing_type is `PrePaid`, the query returns this field.
        /// </summary>
        [Input("renewType")]
        public Input<string>? RenewType { get; set; }

        /// <summary>
        /// The slave zone ID of the CLB.
        /// </summary>
        [Input("slaveZoneId")]
        public Input<string>? SlaveZoneId { get; set; }

        /// <summary>
        /// The id of the Subnet.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputList<Inputs.ClbTagGetArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.ClbTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ClbTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the CLB. And optional choice contains `public` or `private`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The id of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public ClbState()
        {
        }
        public static new ClbState Empty => new ClbState();
    }
}