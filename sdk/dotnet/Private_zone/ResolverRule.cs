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
    /// Provides a resource to manage private zone resolver rule
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
    ///     var foo = new Volcengine.Private_zone.ResolverRule("foo", new()
    ///     {
    ///         EndpointId = 346,
    ///         ForwardIps = new[]
    ///         {
    ///             new Volcengine.Private_zone.Inputs.ResolverRuleForwardIpArgs
    ///             {
    ///                 Ip = "10.199.38.19",
    ///                 Port = 33,
    ///             },
    ///         },
    ///         ProjectName = "default",
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Private_zone.Inputs.ResolverRuleTagArgs
    ///             {
    ///                 Key = "k1",
    ///                 Value = "v1",
    ///             },
    ///         },
    ///         Type = "OUTBOUND",
    ///         Vpcs = new[]
    ///         {
    ///             new Volcengine.Private_zone.Inputs.ResolverRuleVpcArgs
    ///             {
    ///                 Region = "cn-beijing",
    ///                 VpcId = "vpc-13f9uuuqfdjb43n6nu5p1****",
    ///             },
    ///         },
    ///         ZoneNames = new[]
    ///         {
    ///             "www.baidu.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// PrivateZoneResolverRule can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:private_zone/resolverRule:ResolverRule default resource_id
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:private_zone/resolverRule:ResolverRule")]
    public partial class ResolverRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Terminal node ID. This parameter is only valid and required when the Type parameter is OUTBOUND.
        /// </summary>
        [Output("endpointId")]
        public Output<int?> EndpointId { get; private set; } = null!;

        /// <summary>
        /// The endpoint trn of the private zone resolver rule. Format: trn:private_zone::accountId:endpoint/endpointId. This field is only effected when creating resource. 
        /// When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Output("endpointTrn")]
        public Output<string?> EndpointTrn { get; private set; } = null!;

        /// <summary>
        /// IP address and port of external DNS server. You can add up to 10 IP addresses. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
        /// </summary>
        [Output("forwardIps")]
        public Output<ImmutableArray<Outputs.ResolverRuleForwardIp>> ForwardIps { get; private set; } = null!;

        /// <summary>
        /// The operator of the exit IP address of the recursive DNS server. This parameter is only valid when the Type parameter is LINE and is a required parameter. MOBILE, TELECOM, UNICOM.
        /// </summary>
        [Output("line")]
        public Output<string?> Line { get; private set; } = null!;

        /// <summary>
        /// The name of the rule.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project name of the private zone resolver rule.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ResolverRuleTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Forwarding rule types. OUTBOUND: Forward to external DNS servers. LINE: Set the recursive DNS server used for recursive resolution to the recursive DNS server of the Volcano Engine PublicDNS, and customize the operator's exit IP address for the recursive DNS server.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The vpc trns of the private zone resolver rule. Format: trn:vpc:region:accountId:vpc/vpcId. This field is only effected when creating resource. 
        /// When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Output("vpcTrns")]
        public Output<ImmutableArray<string>> VpcTrns { get; private set; } = null!;

        /// <summary>
        /// The parameter name &lt;region&gt; is a variable that represents the region where the VPC is located, such as cn-beijing. The parameter value can include one or more VPC IDs, such as vpc-2750bd1. For example, if you associate a VPC in the cn-beijing region with a domain name and the VPC ID is vpc-2d6si87atfh1c58ozfd0nzq8k, the parameter would be "cn-beijing":["vpc-2d6si87atfh1c58ozfd0nzq8k"]. You can add one or more regions. When the Type parameter is OUTBOUND, the VPC region must be the same as the region where the endpoint is located.
        /// </summary>
        [Output("vpcs")]
        public Output<ImmutableArray<Outputs.ResolverRuleVpc>> Vpcs { get; private set; } = null!;

        /// <summary>
        /// Domain names associated with forwarding rules. You can enter one or more domain names. Up to 500 domain names are supported. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
        /// </summary>
        [Output("zoneNames")]
        public Output<ImmutableArray<string>> ZoneNames { get; private set; } = null!;


        /// <summary>
        /// Create a ResolverRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResolverRule(string name, ResolverRuleArgs args, CustomResourceOptions? options = null)
            : base("volcengine:private_zone/resolverRule:ResolverRule", name, args ?? new ResolverRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResolverRule(string name, Input<string> id, ResolverRuleState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:private_zone/resolverRule:ResolverRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResolverRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResolverRule Get(string name, Input<string> id, ResolverRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new ResolverRule(name, id, state, options);
        }
    }

    public sealed class ResolverRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Terminal node ID. This parameter is only valid and required when the Type parameter is OUTBOUND.
        /// </summary>
        [Input("endpointId")]
        public Input<int>? EndpointId { get; set; }

        /// <summary>
        /// The endpoint trn of the private zone resolver rule. Format: trn:private_zone::accountId:endpoint/endpointId. This field is only effected when creating resource. 
        /// When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("endpointTrn")]
        public Input<string>? EndpointTrn { get; set; }

        [Input("forwardIps")]
        private InputList<Inputs.ResolverRuleForwardIpArgs>? _forwardIps;

        /// <summary>
        /// IP address and port of external DNS server. You can add up to 10 IP addresses. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
        /// </summary>
        public InputList<Inputs.ResolverRuleForwardIpArgs> ForwardIps
        {
            get => _forwardIps ?? (_forwardIps = new InputList<Inputs.ResolverRuleForwardIpArgs>());
            set => _forwardIps = value;
        }

        /// <summary>
        /// The operator of the exit IP address of the recursive DNS server. This parameter is only valid when the Type parameter is LINE and is a required parameter. MOBILE, TELECOM, UNICOM.
        /// </summary>
        [Input("line")]
        public Input<string>? Line { get; set; }

        /// <summary>
        /// The name of the rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project name of the private zone resolver rule.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        [Input("tags")]
        private InputList<Inputs.ResolverRuleTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.ResolverRuleTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ResolverRuleTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Forwarding rule types. OUTBOUND: Forward to external DNS servers. LINE: Set the recursive DNS server used for recursive resolution to the recursive DNS server of the Volcano Engine PublicDNS, and customize the operator's exit IP address for the recursive DNS server.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("vpcTrns")]
        private InputList<string>? _vpcTrns;

        /// <summary>
        /// The vpc trns of the private zone resolver rule. Format: trn:vpc:region:accountId:vpc/vpcId. This field is only effected when creating resource. 
        /// When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        public InputList<string> VpcTrns
        {
            get => _vpcTrns ?? (_vpcTrns = new InputList<string>());
            set => _vpcTrns = value;
        }

        [Input("vpcs", required: true)]
        private InputList<Inputs.ResolverRuleVpcArgs>? _vpcs;

        /// <summary>
        /// The parameter name &lt;region&gt; is a variable that represents the region where the VPC is located, such as cn-beijing. The parameter value can include one or more VPC IDs, such as vpc-2750bd1. For example, if you associate a VPC in the cn-beijing region with a domain name and the VPC ID is vpc-2d6si87atfh1c58ozfd0nzq8k, the parameter would be "cn-beijing":["vpc-2d6si87atfh1c58ozfd0nzq8k"]. You can add one or more regions. When the Type parameter is OUTBOUND, the VPC region must be the same as the region where the endpoint is located.
        /// </summary>
        public InputList<Inputs.ResolverRuleVpcArgs> Vpcs
        {
            get => _vpcs ?? (_vpcs = new InputList<Inputs.ResolverRuleVpcArgs>());
            set => _vpcs = value;
        }

        [Input("zoneNames")]
        private InputList<string>? _zoneNames;

        /// <summary>
        /// Domain names associated with forwarding rules. You can enter one or more domain names. Up to 500 domain names are supported. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
        /// </summary>
        public InputList<string> ZoneNames
        {
            get => _zoneNames ?? (_zoneNames = new InputList<string>());
            set => _zoneNames = value;
        }

        public ResolverRuleArgs()
        {
        }
        public static new ResolverRuleArgs Empty => new ResolverRuleArgs();
    }

    public sealed class ResolverRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Terminal node ID. This parameter is only valid and required when the Type parameter is OUTBOUND.
        /// </summary>
        [Input("endpointId")]
        public Input<int>? EndpointId { get; set; }

        /// <summary>
        /// The endpoint trn of the private zone resolver rule. Format: trn:private_zone::accountId:endpoint/endpointId. This field is only effected when creating resource. 
        /// When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("endpointTrn")]
        public Input<string>? EndpointTrn { get; set; }

        [Input("forwardIps")]
        private InputList<Inputs.ResolverRuleForwardIpGetArgs>? _forwardIps;

        /// <summary>
        /// IP address and port of external DNS server. You can add up to 10 IP addresses. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
        /// </summary>
        public InputList<Inputs.ResolverRuleForwardIpGetArgs> ForwardIps
        {
            get => _forwardIps ?? (_forwardIps = new InputList<Inputs.ResolverRuleForwardIpGetArgs>());
            set => _forwardIps = value;
        }

        /// <summary>
        /// The operator of the exit IP address of the recursive DNS server. This parameter is only valid when the Type parameter is LINE and is a required parameter. MOBILE, TELECOM, UNICOM.
        /// </summary>
        [Input("line")]
        public Input<string>? Line { get; set; }

        /// <summary>
        /// The name of the rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project name of the private zone resolver rule.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        [Input("tags")]
        private InputList<Inputs.ResolverRuleTagGetArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.ResolverRuleTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ResolverRuleTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Forwarding rule types. OUTBOUND: Forward to external DNS servers. LINE: Set the recursive DNS server used for recursive resolution to the recursive DNS server of the Volcano Engine PublicDNS, and customize the operator's exit IP address for the recursive DNS server.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("vpcTrns")]
        private InputList<string>? _vpcTrns;

        /// <summary>
        /// The vpc trns of the private zone resolver rule. Format: trn:vpc:region:accountId:vpc/vpcId. This field is only effected when creating resource. 
        /// When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        public InputList<string> VpcTrns
        {
            get => _vpcTrns ?? (_vpcTrns = new InputList<string>());
            set => _vpcTrns = value;
        }

        [Input("vpcs")]
        private InputList<Inputs.ResolverRuleVpcGetArgs>? _vpcs;

        /// <summary>
        /// The parameter name &lt;region&gt; is a variable that represents the region where the VPC is located, such as cn-beijing. The parameter value can include one or more VPC IDs, such as vpc-2750bd1. For example, if you associate a VPC in the cn-beijing region with a domain name and the VPC ID is vpc-2d6si87atfh1c58ozfd0nzq8k, the parameter would be "cn-beijing":["vpc-2d6si87atfh1c58ozfd0nzq8k"]. You can add one or more regions. When the Type parameter is OUTBOUND, the VPC region must be the same as the region where the endpoint is located.
        /// </summary>
        public InputList<Inputs.ResolverRuleVpcGetArgs> Vpcs
        {
            get => _vpcs ?? (_vpcs = new InputList<Inputs.ResolverRuleVpcGetArgs>());
            set => _vpcs = value;
        }

        [Input("zoneNames")]
        private InputList<string>? _zoneNames;

        /// <summary>
        /// Domain names associated with forwarding rules. You can enter one or more domain names. Up to 500 domain names are supported. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
        /// </summary>
        public InputList<string> ZoneNames
        {
            get => _zoneNames ?? (_zoneNames = new InputList<string>());
            set => _zoneNames = value;
        }

        public ResolverRuleState()
        {
        }
        public static new ResolverRuleState Empty => new ResolverRuleState();
    }
}
