// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Nat
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
    ///     var fooGateway = new Volcengine.Nat.Gateway("fooGateway", new()
    ///     {
    ///         VpcId = fooVpc.Id,
    ///         SubnetId = fooSubnet.Id,
    ///         Spec = "Small",
    ///         NatGatewayName = "acc-test-ng",
    ///         Description = "acc-test",
    ///         BillingType = "PostPaid",
    ///         ProjectName = "default",
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Nat.Inputs.GatewayTagArgs
    ///             {
    ///                 Key = "k1",
    ///                 Value = "v1",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// NatGateway can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:nat/gateway:Gateway default ngw-vv3t043k05sm****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:nat/gateway:Gateway")]
    public partial class Gateway : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The billing type of the NatGateway, the value is `PostPaid` or `PrePaid`.
        /// </summary>
        [Output("billingType")]
        public Output<string?> BillingType { get; private set; } = null!;

        /// <summary>
        /// The description of the NatGateway.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the NatGateway.
        /// </summary>
        [Output("natGatewayName")]
        public Output<string?> NatGatewayName { get; private set; } = null!;

        /// <summary>
        /// The period of the NatGateway, the valid value range in 1~9 or 12 or 24 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid NatGateway. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// The ProjectName of the NatGateway.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The specification of the NatGateway. Optional choice contains `Small`(default), `Medium`, `Large` or leave blank.
        /// </summary>
        [Output("spec")]
        public Output<string> Spec { get; private set; } = null!;

        /// <summary>
        /// The ID of the Subnet.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.GatewayTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Gateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Gateway(string name, GatewayArgs args, CustomResourceOptions? options = null)
            : base("volcengine:nat/gateway:Gateway", name, args ?? new GatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Gateway(string name, Input<string> id, GatewayState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:nat/gateway:Gateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Gateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Gateway Get(string name, Input<string> id, GatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new Gateway(name, id, state, options);
        }
    }

    public sealed class GatewayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The billing type of the NatGateway, the value is `PostPaid` or `PrePaid`.
        /// </summary>
        [Input("billingType")]
        public Input<string>? BillingType { get; set; }

        /// <summary>
        /// The description of the NatGateway.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the NatGateway.
        /// </summary>
        [Input("natGatewayName")]
        public Input<string>? NatGatewayName { get; set; }

        /// <summary>
        /// The period of the NatGateway, the valid value range in 1~9 or 12 or 24 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid NatGateway. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The ProjectName of the NatGateway.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The specification of the NatGateway. Optional choice contains `Small`(default), `Medium`, `Large` or leave blank.
        /// </summary>
        [Input("spec")]
        public Input<string>? Spec { get; set; }

        /// <summary>
        /// The ID of the Subnet.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.GatewayTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.GatewayTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.GatewayTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the VPC.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public GatewayArgs()
        {
        }
        public static new GatewayArgs Empty => new GatewayArgs();
    }

    public sealed class GatewayState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The billing type of the NatGateway, the value is `PostPaid` or `PrePaid`.
        /// </summary>
        [Input("billingType")]
        public Input<string>? BillingType { get; set; }

        /// <summary>
        /// The description of the NatGateway.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the NatGateway.
        /// </summary>
        [Input("natGatewayName")]
        public Input<string>? NatGatewayName { get; set; }

        /// <summary>
        /// The period of the NatGateway, the valid value range in 1~9 or 12 or 24 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid NatGateway. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The ProjectName of the NatGateway.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The specification of the NatGateway. Optional choice contains `Small`(default), `Medium`, `Large` or leave blank.
        /// </summary>
        [Input("spec")]
        public Input<string>? Spec { get; set; }

        /// <summary>
        /// The ID of the Subnet.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputList<Inputs.GatewayTagGetArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.GatewayTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.GatewayTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GatewayState()
        {
        }
        public static new GatewayState Empty => new GatewayState();
    }
}
