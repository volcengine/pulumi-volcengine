// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of customer gateways
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = pulumi.output(volcengine.vpn.CustomerGateways({
 *     ids: ["cgw-2d68c4zglycjk58ozfe96norh"],
 * }));
 * ```
 */
export function customerGateways(args?: CustomerGatewaysArgs, opts?: pulumi.InvokeOptions): Promise<CustomerGatewaysResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:vpn/customerGateways:CustomerGateways", {
        "customerGatewayNames": args.customerGatewayNames,
        "ids": args.ids,
        "ipAddress": args.ipAddress,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking CustomerGateways.
 */
export interface CustomerGatewaysArgs {
    /**
     * A list of customer gateway names.
     */
    customerGatewayNames?: string[];
    /**
     * A list of customer gateway ids.
     */
    ids?: string[];
    /**
     * A IP address of the customer gateway.
     */
    ipAddress?: string;
    /**
     * A Name Regex of customer gateway.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by CustomerGateways.
 */
export interface CustomerGatewaysResult {
    readonly customerGatewayNames?: string[];
    /**
     * The collection of customer gateway query.
     */
    readonly customerGateways: outputs.vpn.CustomerGatewaysCustomerGateway[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The IP address of the customer gateway.
     */
    readonly ipAddress?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The total count of customer gateway query.
     */
    readonly totalCount: number;
}

export function customerGatewaysOutput(args?: CustomerGatewaysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<CustomerGatewaysResult> {
    return pulumi.output(args).apply(a => customerGateways(a, opts))
}

/**
 * A collection of arguments for invoking CustomerGateways.
 */
export interface CustomerGatewaysOutputArgs {
    /**
     * A list of customer gateway names.
     */
    customerGatewayNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of customer gateway ids.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A IP address of the customer gateway.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * A Name Regex of customer gateway.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}