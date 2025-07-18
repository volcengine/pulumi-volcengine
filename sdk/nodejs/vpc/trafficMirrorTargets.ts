// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of traffic mirror targets
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.vpc.getTrafficMirrorTargets({
 *     trafficMirrorTargetIds: ["tmt-rry7yljufsw0v0x58w2****"],
 * });
 * ```
 */
/** @deprecated volcengine.vpc.TrafficMirrorTargets has been deprecated in favor of volcengine.vpc.getTrafficMirrorTargets */
export function trafficMirrorTargets(args?: TrafficMirrorTargetsArgs, opts?: pulumi.InvokeOptions): Promise<TrafficMirrorTargetsResult> {
    pulumi.log.warn("trafficMirrorTargets is deprecated: volcengine.vpc.TrafficMirrorTargets has been deprecated in favor of volcengine.vpc.getTrafficMirrorTargets")
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:vpc/trafficMirrorTargets:TrafficMirrorTargets", {
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "projectName": args.projectName,
        "tags": args.tags,
        "trafficMirrorTargetIds": args.trafficMirrorTargetIds,
        "trafficMirrorTargetName": args.trafficMirrorTargetName,
    }, opts);
}

/**
 * A collection of arguments for invoking TrafficMirrorTargets.
 */
export interface TrafficMirrorTargetsArgs {
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The project name of traffic mirror target.
     */
    projectName?: string;
    /**
     * Tags.
     */
    tags?: inputs.vpc.TrafficMirrorTargetsTag[];
    /**
     * A list of traffic mirror target IDs.
     */
    trafficMirrorTargetIds?: string[];
    /**
     * The name of traffic mirror target.
     */
    trafficMirrorTargetName?: string;
}

/**
 * A collection of values returned by TrafficMirrorTargets.
 */
export interface TrafficMirrorTargetsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The project name of traffic mirror target.
     */
    readonly projectName?: string;
    /**
     * Tags.
     */
    readonly tags?: outputs.vpc.TrafficMirrorTargetsTag[];
    /**
     * The total count of query.
     */
    readonly totalCount: number;
    readonly trafficMirrorTargetIds?: string[];
    /**
     * The name of traffic mirror target.
     */
    readonly trafficMirrorTargetName?: string;
    /**
     * The collection of query.
     */
    readonly trafficMirrorTargets: outputs.vpc.TrafficMirrorTargetsTrafficMirrorTarget[];
}
/**
 * Use this data source to query detailed information of traffic mirror targets
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.vpc.getTrafficMirrorTargets({
 *     trafficMirrorTargetIds: ["tmt-rry7yljufsw0v0x58w2****"],
 * });
 * ```
 */
/** @deprecated volcengine.vpc.TrafficMirrorTargets has been deprecated in favor of volcengine.vpc.getTrafficMirrorTargets */
export function trafficMirrorTargetsOutput(args?: TrafficMirrorTargetsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<TrafficMirrorTargetsResult> {
    return pulumi.output(args).apply((a: any) => trafficMirrorTargets(a, opts))
}

/**
 * A collection of arguments for invoking TrafficMirrorTargets.
 */
export interface TrafficMirrorTargetsOutputArgs {
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The project name of traffic mirror target.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.vpc.TrafficMirrorTargetsTagArgs>[]>;
    /**
     * A list of traffic mirror target IDs.
     */
    trafficMirrorTargetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of traffic mirror target.
     */
    trafficMirrorTargetName?: pulumi.Input<string>;
}
