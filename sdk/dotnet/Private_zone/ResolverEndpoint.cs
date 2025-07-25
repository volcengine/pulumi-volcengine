// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Private_zone
{
    /// <summary>
    /// Provides a resource to manage private zone resolver endpoint
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
    ///     var foo = new Volcengine.Private_zone.ResolverEndpoint("foo", new()
    ///     {
    ///         IpConfigs = new[]
    ///         {
    ///             new Volcengine.Private_zone.Inputs.ResolverEndpointIpConfigArgs
    ///             {
    ///                 AzId = "cn-beijing-a",
    ///                 Ip = "172.16.0.2",
    ///                 SubnetId = "subnet-mj2o4co2m2v45smt1bx1****",
    ///             },
    ///             new Volcengine.Private_zone.Inputs.ResolverEndpointIpConfigArgs
    ///             {
    ///                 AzId = "cn-beijing-a",
    ///                 Ip = "172.16.0.3",
    ///                 SubnetId = "subnet-mj2o4co2m2v45smt1bx1****",
    ///             },
    ///             new Volcengine.Private_zone.Inputs.ResolverEndpointIpConfigArgs
    ///             {
    ///                 AzId = "cn-beijing-a",
    ///                 Ip = "172.16.0.4",
    ///                 SubnetId = "subnet-mj2o4co2m2v45smt1bx1****",
    ///             },
    ///             new Volcengine.Private_zone.Inputs.ResolverEndpointIpConfigArgs
    ///             {
    ///                 AzId = "cn-beijing-a",
    ///                 Ip = "172.16.0.5",
    ///                 SubnetId = "subnet-mj2o4co2m2v45smt1bx1****",
    ///             },
    ///         },
    ///         ProjectName = "default",
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Private_zone.Inputs.ResolverEndpointTagArgs
    ///             {
    ///                 Key = "k1",
    ///                 Value = "v1",
    ///             },
    ///         },
    ///         VpcId = "vpc-13f9uuuqfdjb43n6nu5p1****",
    ///         VpcRegion = "cn-beijing",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// PrivateZoneResolverEndpoint can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:private_zone/resolverEndpoint:ResolverEndpoint default resource_id
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:private_zone/resolverEndpoint:ResolverEndpoint")]
    public partial class ResolverEndpoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// DNS request forwarding direction for terminal nodes. OUTBOUND: (default) Outbound terminal nodes forward DNS query requests from within the VPC to external DNS servers. INBOUND: Inbound terminal nodes forward DNS query requests from external sources to resolvers.
        /// </summary>
        [Output("direction")]
        public Output<string?> Direction { get; private set; } = null!;

        /// <summary>
        /// Availability zones, subnets, and IP configurations of terminal nodes.
        /// </summary>
        [Output("ipConfigs")]
        public Output<ImmutableArray<Outputs.ResolverEndpointIpConfig>> IpConfigs { get; private set; } = null!;

        /// <summary>
        /// The name of the private zone resolver endpoint.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project name of the private zone resolver endpoint.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The security group ID of the endpoint.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ResolverEndpointTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The VPC ID of the endpoint.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The VPC region of the endpoint.
        /// </summary>
        [Output("vpcRegion")]
        public Output<string> VpcRegion { get; private set; } = null!;

        /// <summary>
        /// The vpc trns of the private zone resolver endpoint. Format: trn:vpc:region:accountId:vpc/vpcId. This field is only effected when creating resource. 
        /// When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Output("vpcTrns")]
        public Output<ImmutableArray<string>> VpcTrns { get; private set; } = null!;


        /// <summary>
        /// Create a ResolverEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResolverEndpoint(string name, ResolverEndpointArgs args, CustomResourceOptions? options = null)
            : base("volcengine:private_zone/resolverEndpoint:ResolverEndpoint", name, args ?? new ResolverEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResolverEndpoint(string name, Input<string> id, ResolverEndpointState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:private_zone/resolverEndpoint:ResolverEndpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResolverEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResolverEndpoint Get(string name, Input<string> id, ResolverEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new ResolverEndpoint(name, id, state, options);
        }
    }

    public sealed class ResolverEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS request forwarding direction for terminal nodes. OUTBOUND: (default) Outbound terminal nodes forward DNS query requests from within the VPC to external DNS servers. INBOUND: Inbound terminal nodes forward DNS query requests from external sources to resolvers.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        [Input("ipConfigs", required: true)]
        private InputList<Inputs.ResolverEndpointIpConfigArgs>? _ipConfigs;

        /// <summary>
        /// Availability zones, subnets, and IP configurations of terminal nodes.
        /// </summary>
        public InputList<Inputs.ResolverEndpointIpConfigArgs> IpConfigs
        {
            get => _ipConfigs ?? (_ipConfigs = new InputList<Inputs.ResolverEndpointIpConfigArgs>());
            set => _ipConfigs = value;
        }

        /// <summary>
        /// The name of the private zone resolver endpoint.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project name of the private zone resolver endpoint.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The security group ID of the endpoint.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        [Input("tags")]
        private InputList<Inputs.ResolverEndpointTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.ResolverEndpointTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ResolverEndpointTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The VPC ID of the endpoint.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        /// <summary>
        /// The VPC region of the endpoint.
        /// </summary>
        [Input("vpcRegion", required: true)]
        public Input<string> VpcRegion { get; set; } = null!;

        [Input("vpcTrns")]
        private InputList<string>? _vpcTrns;

        /// <summary>
        /// The vpc trns of the private zone resolver endpoint. Format: trn:vpc:region:accountId:vpc/vpcId. This field is only effected when creating resource. 
        /// When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        public InputList<string> VpcTrns
        {
            get => _vpcTrns ?? (_vpcTrns = new InputList<string>());
            set => _vpcTrns = value;
        }

        public ResolverEndpointArgs()
        {
        }
        public static new ResolverEndpointArgs Empty => new ResolverEndpointArgs();
    }

    public sealed class ResolverEndpointState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS request forwarding direction for terminal nodes. OUTBOUND: (default) Outbound terminal nodes forward DNS query requests from within the VPC to external DNS servers. INBOUND: Inbound terminal nodes forward DNS query requests from external sources to resolvers.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        [Input("ipConfigs")]
        private InputList<Inputs.ResolverEndpointIpConfigGetArgs>? _ipConfigs;

        /// <summary>
        /// Availability zones, subnets, and IP configurations of terminal nodes.
        /// </summary>
        public InputList<Inputs.ResolverEndpointIpConfigGetArgs> IpConfigs
        {
            get => _ipConfigs ?? (_ipConfigs = new InputList<Inputs.ResolverEndpointIpConfigGetArgs>());
            set => _ipConfigs = value;
        }

        /// <summary>
        /// The name of the private zone resolver endpoint.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project name of the private zone resolver endpoint.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The security group ID of the endpoint.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        [Input("tags")]
        private InputList<Inputs.ResolverEndpointTagGetArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.ResolverEndpointTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ResolverEndpointTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The VPC ID of the endpoint.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The VPC region of the endpoint.
        /// </summary>
        [Input("vpcRegion")]
        public Input<string>? VpcRegion { get; set; }

        [Input("vpcTrns")]
        private InputList<string>? _vpcTrns;

        /// <summary>
        /// The vpc trns of the private zone resolver endpoint. Format: trn:vpc:region:accountId:vpc/vpcId. This field is only effected when creating resource. 
        /// When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        public InputList<string> VpcTrns
        {
            get => _vpcTrns ?? (_vpcTrns = new InputList<string>());
            set => _vpcTrns = value;
        }

        public ResolverEndpointState()
        {
        }
        public static new ResolverEndpointState Empty => new ResolverEndpointState();
    }
}
