// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of direct connect gateways
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.direct_connect.Gateways({
 *     directConnectGatewayName: "tf-test",
 * });
 * ```
 */
export function gateways(args?: GatewaysArgs, opts?: pulumi.InvokeOptions): Promise<GatewaysResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:direct_connect/gateways:Gateways", {
        "cenId": args.cenId,
        "directConnectGatewayName": args.directConnectGatewayName,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "tagFilters": args.tagFilters,
    }, opts);
}

/**
 * A collection of arguments for invoking Gateways.
 */
export interface GatewaysArgs {
    /**
     * The CEN ID which direct connect gateway belongs.
     */
    cenId?: string;
    /**
     * The direst connect gateway name.
     */
    directConnectGatewayName?: string;
    /**
     * A list of IDs.
     */
    ids?: string[];
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The filter tag of direct connect.
     */
    tagFilters?: inputs.direct_connect.GatewaysTagFilter[];
}

/**
 * A collection of values returned by Gateways.
 */
export interface GatewaysResult {
    /**
     * The cen ID.
     */
    readonly cenId?: string;
    /**
     * The direct connect gateway name.
     */
    readonly directConnectGatewayName?: string;
    /**
     * The collection of query.
     */
    readonly directConnectGateways: outputs.direct_connect.GatewaysDirectConnectGateway[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly nameRegex?: string;
    readonly outputFile?: string;
    readonly tagFilters?: outputs.direct_connect.GatewaysTagFilter[];
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of direct connect gateways
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.direct_connect.Gateways({
 *     directConnectGatewayName: "tf-test",
 * });
 * ```
 */
export function gatewaysOutput(args?: GatewaysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GatewaysResult> {
    return pulumi.output(args).apply((a: any) => gateways(a, opts))
}

/**
 * A collection of arguments for invoking Gateways.
 */
export interface GatewaysOutputArgs {
    /**
     * The CEN ID which direct connect gateway belongs.
     */
    cenId?: pulumi.Input<string>;
    /**
     * The direst connect gateway name.
     */
    directConnectGatewayName?: pulumi.Input<string>;
    /**
     * A list of IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The filter tag of direct connect.
     */
    tagFilters?: pulumi.Input<pulumi.Input<inputs.direct_connect.GatewaysTagFilterArgs>[]>;
}