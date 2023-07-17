// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vpn gateway routes
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultGatewayRoutes = pulumi.output(volcengine.vpn.GatewayRoutes({
 *     ids: ["vgr-2byssu52dktts2dx0ee90r5hp]"],
 * }));
 * ```
 */
export function gatewayRoutes(args?: GatewayRoutesArgs, opts?: pulumi.InvokeOptions): Promise<GatewayRoutesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:vpn/gatewayRoutes:GatewayRoutes", {
        "destinationCidrBlock": args.destinationCidrBlock,
        "ids": args.ids,
        "nextHopId": args.nextHopId,
        "outputFile": args.outputFile,
        "vpnGatewayId": args.vpnGatewayId,
    }, opts);
}

/**
 * A collection of arguments for invoking GatewayRoutes.
 */
export interface GatewayRoutesArgs {
    /**
     * A destination cidr block.
     */
    destinationCidrBlock?: string;
    /**
     * A list of VPN gateway route ids.
     */
    ids?: string[];
    /**
     * An ID of next hop.
     */
    nextHopId?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * An ID of VPN gateway.
     */
    vpnGatewayId?: string;
}

/**
 * A collection of values returned by GatewayRoutes.
 */
export interface GatewayRoutesResult {
    /**
     * The destination cidr block of the VPN gateway route.
     */
    readonly destinationCidrBlock?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The next hop id of the VPN gateway route.
     */
    readonly nextHopId?: string;
    readonly outputFile?: string;
    /**
     * The total count of VPN gateway route query.
     */
    readonly totalCount: number;
    /**
     * The ID of the VPN gateway of the VPN gateway route.
     */
    readonly vpnGatewayId?: string;
    /**
     * The collection of VPN gateway route query.
     */
    readonly vpnGatewayRoutes: outputs.vpn.GatewayRoutesVpnGatewayRoute[];
}

export function gatewayRoutesOutput(args?: GatewayRoutesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GatewayRoutesResult> {
    return pulumi.output(args).apply(a => gatewayRoutes(a, opts))
}

/**
 * A collection of arguments for invoking GatewayRoutes.
 */
export interface GatewayRoutesOutputArgs {
    /**
     * A destination cidr block.
     */
    destinationCidrBlock?: pulumi.Input<string>;
    /**
     * A list of VPN gateway route ids.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An ID of next hop.
     */
    nextHopId?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * An ID of VPN gateway.
     */
    vpnGatewayId?: pulumi.Input<string>;
}