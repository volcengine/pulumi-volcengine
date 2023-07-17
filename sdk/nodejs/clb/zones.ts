// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of clb zones
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultZones = pulumi.output(volcengine.clb.Zones());
 * ```
 */
export function zones(args?: ZonesArgs, opts?: pulumi.InvokeOptions): Promise<ZonesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:clb/zones:Zones", {
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking Zones.
 */
export interface ZonesArgs {
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by Zones.
 */
export interface ZonesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The master zones list.
     */
    readonly masterZones: outputs.clb.ZonesMasterZone[];
    readonly outputFile?: string;
    /**
     * The total count of zone query.
     */
    readonly totalCount: number;
}

export function zonesOutput(args?: ZonesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ZonesResult> {
    return pulumi.output(args).apply(a => zones(a, opts))
}

/**
 * A collection of arguments for invoking Zones.
 */
export interface ZonesOutputArgs {
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}