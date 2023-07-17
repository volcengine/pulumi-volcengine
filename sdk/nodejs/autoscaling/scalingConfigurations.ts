// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of scaling configurations
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultScalingConfigurations = pulumi.output(volcengine.autoscaling.ScalingConfigurations({
 *     ids: ["scc-ybrurj4uw6gh9zecj327"],
 * }));
 * ```
 */
export function scalingConfigurations(args?: ScalingConfigurationsArgs, opts?: pulumi.InvokeOptions): Promise<ScalingConfigurationsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:autoscaling/scalingConfigurations:ScalingConfigurations", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "scalingConfigurationNames": args.scalingConfigurationNames,
        "scalingGroupId": args.scalingGroupId,
    }, opts);
}

/**
 * A collection of arguments for invoking ScalingConfigurations.
 */
export interface ScalingConfigurationsArgs {
    /**
     * A list of scaling configuration ids.
     */
    ids?: string[];
    /**
     * A Name Regex of scaling configuration.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * A list of scaling configuration names.
     */
    scalingConfigurationNames?: string[];
    /**
     * An id of scaling group.
     */
    scalingGroupId?: string;
}

/**
 * A collection of values returned by ScalingConfigurations.
 */
export interface ScalingConfigurationsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly nameRegex?: string;
    readonly outputFile?: string;
    readonly scalingConfigurationNames?: string[];
    /**
     * The collection of scaling configuration query.
     */
    readonly scalingConfigurations: outputs.autoscaling.ScalingConfigurationsScalingConfiguration[];
    /**
     * The id of the scaling group to which the scaling configuration belongs.
     */
    readonly scalingGroupId?: string;
    /**
     * The total count of scaling configuration query.
     */
    readonly totalCount: number;
}

export function scalingConfigurationsOutput(args?: ScalingConfigurationsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ScalingConfigurationsResult> {
    return pulumi.output(args).apply(a => scalingConfigurations(a, opts))
}

/**
 * A collection of arguments for invoking ScalingConfigurations.
 */
export interface ScalingConfigurationsOutputArgs {
    /**
     * A list of scaling configuration ids.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Name Regex of scaling configuration.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * A list of scaling configuration names.
     */
    scalingConfigurationNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An id of scaling group.
     */
    scalingGroupId?: pulumi.Input<string>;
}