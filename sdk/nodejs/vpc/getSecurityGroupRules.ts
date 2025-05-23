// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of security group rules
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vpc.getSecurityGroupRules({
 *     securityGroupId: "sg-13f2nau7x93wg3n6nu3z5sxib",
 * });
 * ```
 */
export function getSecurityGroupRules(args: GetSecurityGroupRulesArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityGroupRulesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:vpc/getSecurityGroupRules:getSecurityGroupRules", {
        "cidrIp": args.cidrIp,
        "direction": args.direction,
        "outputFile": args.outputFile,
        "protocol": args.protocol,
        "securityGroupId": args.securityGroupId,
        "sourceGroupId": args.sourceGroupId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecurityGroupRules.
 */
export interface GetSecurityGroupRulesArgs {
    /**
     * Cidr ip of egress/ingress Rule.
     */
    cidrIp?: string;
    /**
     * Direction of rule, ingress (inbound) or egress (outbound).
     */
    direction?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all`.
     */
    protocol?: string;
    /**
     * SecurityGroup ID.
     */
    securityGroupId: string;
    /**
     * ID of the source security group whose access permission you want to set.
     */
    sourceGroupId?: string;
}

/**
 * A collection of values returned by getSecurityGroupRules.
 */
export interface GetSecurityGroupRulesResult {
    /**
     * Cidr ip of egress/ingress Rule.
     */
    readonly cidrIp?: string;
    /**
     * Direction of rule, ingress (inbound) or egress (outbound).
     */
    readonly direction?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
    /**
     * Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all`.
     */
    readonly protocol?: string;
    /**
     * Id of SecurityGroup.
     */
    readonly securityGroupId: string;
    /**
     * The collection of SecurityGroup query.
     */
    readonly securityGroupRules: outputs.vpc.GetSecurityGroupRulesSecurityGroupRule[];
    /**
     * ID of the source security group whose access permission you want to set.
     */
    readonly sourceGroupId?: string;
}
/**
 * Use this data source to query detailed information of security group rules
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vpc.getSecurityGroupRules({
 *     securityGroupId: "sg-13f2nau7x93wg3n6nu3z5sxib",
 * });
 * ```
 */
export function getSecurityGroupRulesOutput(args: GetSecurityGroupRulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecurityGroupRulesResult> {
    return pulumi.output(args).apply((a: any) => getSecurityGroupRules(a, opts))
}

/**
 * A collection of arguments for invoking getSecurityGroupRules.
 */
export interface GetSecurityGroupRulesOutputArgs {
    /**
     * Cidr ip of egress/ingress Rule.
     */
    cidrIp?: pulumi.Input<string>;
    /**
     * Direction of rule, ingress (inbound) or egress (outbound).
     */
    direction?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * SecurityGroup ID.
     */
    securityGroupId: pulumi.Input<string>;
    /**
     * ID of the source security group whose access permission you want to set.
     */
    sourceGroupId?: pulumi.Input<string>;
}
