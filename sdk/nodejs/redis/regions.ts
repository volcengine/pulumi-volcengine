// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of redis regions
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.redis.getRegions({
 *     regionId: "cn-north-3",
 * });
 * ```
 */
/** @deprecated volcengine.redis.Regions has been deprecated in favor of volcengine.redis.getRegions */
export function regions(args?: RegionsArgs, opts?: pulumi.InvokeOptions): Promise<RegionsResult> {
    pulumi.log.warn("regions is deprecated: volcengine.redis.Regions has been deprecated in favor of volcengine.redis.getRegions")
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:redis/regions:Regions", {
        "outputFile": args.outputFile,
        "regionId": args.regionId,
    }, opts);
}

/**
 * A collection of arguments for invoking Regions.
 */
export interface RegionsArgs {
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * Target region info.
     */
    regionId?: string;
}

/**
 * A collection of values returned by Regions.
 */
export interface RegionsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
    /**
     * The id of the region.
     */
    readonly regionId?: string;
    /**
     * The collection of region query.
     */
    readonly regions: outputs.redis.RegionsRegion[];
    /**
     * The total count of region query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of redis regions
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.redis.getRegions({
 *     regionId: "cn-north-3",
 * });
 * ```
 */
/** @deprecated volcengine.redis.Regions has been deprecated in favor of volcengine.redis.getRegions */
export function regionsOutput(args?: RegionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<RegionsResult> {
    return pulumi.output(args).apply((a: any) => regions(a, opts))
}

/**
 * A collection of arguments for invoking Regions.
 */
export interface RegionsOutputArgs {
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * Target region info.
     */
    regionId?: pulumi.Input<string>;
}
