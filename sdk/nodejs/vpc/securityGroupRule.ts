// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage security group rule
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const g1test3 = new volcengine.vpc.SecurityGroupRule("g1test3", {
 *     cidrIp: "10.0.0.0/8",
 *     description: "tft1234",
 *     direction: "egress",
 *     portEnd: 9003,
 *     portStart: 8000,
 *     protocol: "tcp",
 *     securityGroupId: "sg-2d6722jpp55og58ozfd1sqtdb",
 * });
 * const g1test2 = new volcengine.vpc.SecurityGroupRule("g1test2", {
 *     cidrIp: "10.0.0.0/24",
 *     direction: "egress",
 *     portEnd: 9003,
 *     portStart: 8000,
 *     protocol: "tcp",
 *     securityGroupId: "sg-2d6722jpp55og58ozfd1sqtdb",
 * });
 * const g1test1 = new volcengine.vpc.SecurityGroupRule("g1test1", {
 *     cidrIp: "10.0.0.0/24",
 *     direction: "egress",
 *     portEnd: 9003,
 *     portStart: 8000,
 *     priority: 2,
 *     protocol: "tcp",
 *     securityGroupId: "sg-2d6722jpp55og58ozfd1sqtdb",
 * });
 * const g1test0 = new volcengine.vpc.SecurityGroupRule("g1test0", {
 *     cidrIp: "10.0.0.0/24",
 *     description: "tft",
 *     direction: "ingress",
 *     policy: "drop",
 *     portEnd: 80,
 *     portStart: 80,
 *     priority: 2,
 *     protocol: "tcp",
 *     securityGroupId: "sg-2d6722jpp55og58ozfd1sqtdb",
 * });
 * const g1test06 = new volcengine.vpc.SecurityGroupRule("g1test06", {
 *     description: "tft",
 *     direction: "ingress",
 *     policy: "drop",
 *     portEnd: 9003,
 *     portStart: 8000,
 *     priority: 2,
 *     protocol: "tcp",
 *     securityGroupId: "sg-2d6722jpp55og58ozfd1sqtdb",
 *     sourceGroupId: "sg-3rfe5j4xdnklc5zsk2hcw5c6q",
 * });
 * ```
 *
 * ## Import
 *
 * SecurityGroupRule can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:vpc/securityGroupRule:SecurityGroupRule default ID is a string concatenated with colons(SecurityGroupId:Protocol:PortStart:PortEnd:CidrIp:SourceGroupId:Direction:Policy:Priority)
 * ```
 */
export class SecurityGroupRule extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroupRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityGroupRuleState, opts?: pulumi.CustomResourceOptions): SecurityGroupRule {
        return new SecurityGroupRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:vpc/securityGroupRule:SecurityGroupRule';

    /**
     * Returns true if the given object is an instance of SecurityGroupRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityGroupRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityGroupRule.__pulumiType;
    }

    /**
     * Cidr ip of egress/ingress Rule.
     */
    public readonly cidrIp!: pulumi.Output<string | undefined>;
    /**
     * description of a egress rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Direction of rule, ingress (inbound) or egress (outbound).
     */
    public readonly direction!: pulumi.Output<string>;
    /**
     * Access strategy. Valid values: `accept`, `drop`. Default is `accept`.
     */
    public readonly policy!: pulumi.Output<string | undefined>;
    /**
     * Port end of egress/ingress Rule. When the `protocol` is `tcp` or `udp`, the valid value range is 1~65535. When the `protocol` is `icmp` or `all` or `icmpv6`, the valid value is -1, indicating no restriction on port values.
     */
    public readonly portEnd!: pulumi.Output<number>;
    /**
     * Port start of egress/ingress Rule. When the `protocol` is `tcp` or `udp`, the valid value range is 1~65535. When the `protocol` is `icmp` or `all` or `icmpv6`, the valid value is -1, indicating no restriction on port values.
     */
    public readonly portStart!: pulumi.Output<number>;
    /**
     * Priority of a security group rule. Valid value range: 1~100. Default is 1.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all` or `icmpv6`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Id of SecurityGroup.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * ID of the source security group whose access permission you want to set.
     */
    public readonly sourceGroupId!: pulumi.Output<string | undefined>;
    /**
     * Status of SecurityGroup.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a SecurityGroupRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityGroupRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityGroupRuleArgs | SecurityGroupRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityGroupRuleState | undefined;
            resourceInputs["cidrIp"] = state ? state.cidrIp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["direction"] = state ? state.direction : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
            resourceInputs["portEnd"] = state ? state.portEnd : undefined;
            resourceInputs["portStart"] = state ? state.portStart : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["sourceGroupId"] = state ? state.sourceGroupId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as SecurityGroupRuleArgs | undefined;
            if ((!args || args.direction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'direction'");
            }
            if ((!args || args.portEnd === undefined) && !opts.urn) {
                throw new Error("Missing required property 'portEnd'");
            }
            if ((!args || args.portStart === undefined) && !opts.urn) {
                throw new Error("Missing required property 'portStart'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.securityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            resourceInputs["cidrIp"] = args ? args.cidrIp : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["direction"] = args ? args.direction : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["portEnd"] = args ? args.portEnd : undefined;
            resourceInputs["portStart"] = args ? args.portStart : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["sourceGroupId"] = args ? args.sourceGroupId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityGroupRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityGroupRule resources.
 */
export interface SecurityGroupRuleState {
    /**
     * Cidr ip of egress/ingress Rule.
     */
    cidrIp?: pulumi.Input<string>;
    /**
     * description of a egress rule.
     */
    description?: pulumi.Input<string>;
    /**
     * Direction of rule, ingress (inbound) or egress (outbound).
     */
    direction?: pulumi.Input<string>;
    /**
     * Access strategy. Valid values: `accept`, `drop`. Default is `accept`.
     */
    policy?: pulumi.Input<string>;
    /**
     * Port end of egress/ingress Rule. When the `protocol` is `tcp` or `udp`, the valid value range is 1~65535. When the `protocol` is `icmp` or `all` or `icmpv6`, the valid value is -1, indicating no restriction on port values.
     */
    portEnd?: pulumi.Input<number>;
    /**
     * Port start of egress/ingress Rule. When the `protocol` is `tcp` or `udp`, the valid value range is 1~65535. When the `protocol` is `icmp` or `all` or `icmpv6`, the valid value is -1, indicating no restriction on port values.
     */
    portStart?: pulumi.Input<number>;
    /**
     * Priority of a security group rule. Valid value range: 1~100. Default is 1.
     */
    priority?: pulumi.Input<number>;
    /**
     * Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all` or `icmpv6`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Id of SecurityGroup.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * ID of the source security group whose access permission you want to set.
     */
    sourceGroupId?: pulumi.Input<string>;
    /**
     * Status of SecurityGroup.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityGroupRule resource.
 */
export interface SecurityGroupRuleArgs {
    /**
     * Cidr ip of egress/ingress Rule.
     */
    cidrIp?: pulumi.Input<string>;
    /**
     * description of a egress rule.
     */
    description?: pulumi.Input<string>;
    /**
     * Direction of rule, ingress (inbound) or egress (outbound).
     */
    direction: pulumi.Input<string>;
    /**
     * Access strategy. Valid values: `accept`, `drop`. Default is `accept`.
     */
    policy?: pulumi.Input<string>;
    /**
     * Port end of egress/ingress Rule. When the `protocol` is `tcp` or `udp`, the valid value range is 1~65535. When the `protocol` is `icmp` or `all` or `icmpv6`, the valid value is -1, indicating no restriction on port values.
     */
    portEnd: pulumi.Input<number>;
    /**
     * Port start of egress/ingress Rule. When the `protocol` is `tcp` or `udp`, the valid value range is 1~65535. When the `protocol` is `icmp` or `all` or `icmpv6`, the valid value is -1, indicating no restriction on port values.
     */
    portStart: pulumi.Input<number>;
    /**
     * Priority of a security group rule. Valid value range: 1~100. Default is 1.
     */
    priority?: pulumi.Input<number>;
    /**
     * Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all` or `icmpv6`.
     */
    protocol: pulumi.Input<string>;
    /**
     * Id of SecurityGroup.
     */
    securityGroupId: pulumi.Input<string>;
    /**
     * ID of the source security group whose access permission you want to set.
     */
    sourceGroupId?: pulumi.Input<string>;
}
