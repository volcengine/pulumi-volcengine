// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vpn gateways
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultGateways = pulumi.output(volcengine.vpn.Gateways({
 *     ids: ["vgw-2c012ea9fm5mo2dx0efxg46qi"],
 * }));
 * ```
 */
export function gateways(args?: GatewaysArgs, opts?: pulumi.InvokeOptions): Promise<GatewaysResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:vpn/gateways:Gateways", {
        "ids": args.ids,
        "ipAddress": args.ipAddress,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "subnetId": args.subnetId,
        "tags": args.tags,
        "vpcId": args.vpcId,
        "vpnGatewayNames": args.vpnGatewayNames,
    }, opts);
}

/**
 * A collection of arguments for invoking Gateways.
 */
export interface GatewaysArgs {
    /**
     * A list of VPN gateway ids.
     */
    ids?: string[];
    /**
     * A IP address of the VPN gateway.
     */
    ipAddress?: string;
    /**
     * A Name Regex of VPN gateway.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * A subnet ID of the VPN gateway.
     */
    subnetId?: string;
    /**
     * Tags.
     */
    tags?: inputs.vpn.GatewaysTag[];
    /**
     * A VPC ID of the VPN gateway.
     */
    vpcId?: string;
    /**
     * A list of VPN gateway names.
     */
    vpnGatewayNames?: string[];
}

/**
 * A collection of values returned by Gateways.
 */
export interface GatewaysResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The IP address of the VPN gateway.
     */
    readonly ipAddress?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    readonly subnetId?: string;
    /**
     * Tags.
     */
    readonly tags?: outputs.vpn.GatewaysTag[];
    /**
     * The total count of VPN gateway query.
     */
    readonly totalCount: number;
    /**
     * The VPC ID of the VPN gateway.
     */
    readonly vpcId?: string;
    readonly vpnGatewayNames?: string[];
    /**
     * The collection of VPN gateway query.
     */
    readonly vpnGateways: outputs.vpn.GatewaysVpnGateway[];
}

export function gatewaysOutput(args?: GatewaysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GatewaysResult> {
    return pulumi.output(args).apply(a => gateways(a, opts))
}

/**
 * A collection of arguments for invoking Gateways.
 */
export interface GatewaysOutputArgs {
    /**
     * A list of VPN gateway ids.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A IP address of the VPN gateway.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * A Name Regex of VPN gateway.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * A subnet ID of the VPN gateway.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.vpn.GatewaysTagArgs>[]>;
    /**
     * A VPC ID of the VPN gateway.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * A list of VPN gateway names.
     */
    vpnGatewayNames?: pulumi.Input<pulumi.Input<string>[]>;
}