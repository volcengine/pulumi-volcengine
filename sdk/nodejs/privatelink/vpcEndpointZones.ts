// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of privatelink vpc endpoint zones
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultVpcEndpointZones = pulumi.output(volcengine.privatelink.VpcEndpointZones({
 *     endpointId: "ep-2byz5npiuu1hc2dx0efkv****",
 * }));
 * ```
 */
export function vpcEndpointZones(args?: VpcEndpointZonesArgs, opts?: pulumi.InvokeOptions): Promise<VpcEndpointZonesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:privatelink/vpcEndpointZones:VpcEndpointZones", {
        "endpointId": args.endpointId,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking VpcEndpointZones.
 */
export interface VpcEndpointZonesArgs {
    /**
     * The endpoint id of query.
     */
    endpointId?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by VpcEndpointZones.
 */
export interface VpcEndpointZonesResult {
    readonly endpointId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
    /**
     * Returns the total amount of the data list.
     */
    readonly totalCount: number;
    /**
     * The collection of query.
     */
    readonly vpcEndpointZones: outputs.privatelink.VpcEndpointZonesVpcEndpointZone[];
}

export function vpcEndpointZonesOutput(args?: VpcEndpointZonesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<VpcEndpointZonesResult> {
    return pulumi.output(args).apply(a => vpcEndpointZones(a, opts))
}

/**
 * A collection of arguments for invoking VpcEndpointZones.
 */
export interface VpcEndpointZonesOutputArgs {
    /**
     * The endpoint id of query.
     */
    endpointId?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}