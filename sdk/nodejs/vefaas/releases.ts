// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vefaas releases
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.vefaas.getReleases({
 *     functionId: "g79asxxx",
 * });
 * ```
 */
/** @deprecated volcengine.vefaas.Releases has been deprecated in favor of volcengine.vefaas.getReleases */
export function releases(args: ReleasesArgs, opts?: pulumi.InvokeOptions): Promise<ReleasesResult> {
    pulumi.log.warn("releases is deprecated: volcengine.vefaas.Releases has been deprecated in favor of volcengine.vefaas.getReleases")

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:vefaas/releases:Releases", {
        "filters": args.filters,
        "functionId": args.functionId,
        "nameRegex": args.nameRegex,
        "orderBies": args.orderBies,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking Releases.
 */
export interface ReleasesArgs {
    /**
     * Query the filtering conditions.
     */
    filters?: inputs.vefaas.ReleasesFilter[];
    /**
     * The ID of Function.
     */
    functionId: string;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: string;
    /**
     * Query the sorting parameters.
     */
    orderBies?: inputs.vefaas.ReleasesOrderBy[];
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by Releases.
 */
export interface ReleasesResult {
    readonly filters?: outputs.vefaas.ReleasesFilter[];
    /**
     * The ID of Function.
     */
    readonly functionId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of function publication records.
     */
    readonly items: outputs.vefaas.ReleasesItem[];
    readonly nameRegex?: string;
    readonly orderBies?: outputs.vefaas.ReleasesOrderBy[];
    readonly outputFile?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of vefaas releases
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.vefaas.getReleases({
 *     functionId: "g79asxxx",
 * });
 * ```
 */
/** @deprecated volcengine.vefaas.Releases has been deprecated in favor of volcengine.vefaas.getReleases */
export function releasesOutput(args: ReleasesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ReleasesResult> {
    return pulumi.output(args).apply((a: any) => releases(a, opts))
}

/**
 * A collection of arguments for invoking Releases.
 */
export interface ReleasesOutputArgs {
    /**
     * Query the filtering conditions.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.vefaas.ReleasesFilterArgs>[]>;
    /**
     * The ID of Function.
     */
    functionId: pulumi.Input<string>;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * Query the sorting parameters.
     */
    orderBies?: pulumi.Input<pulumi.Input<inputs.vefaas.ReleasesOrderByArgs>[]>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
