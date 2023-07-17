// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of ecs deployment sets
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultDeploymentSets = pulumi.output(volcengine.ecs.DeploymentSets({
 *     granularity: "host",
 *     ids: ["dps-ybp1b059cb5m57n135g3"],
 * }));
 * ```
 */
export function deploymentSets(args?: DeploymentSetsArgs, opts?: pulumi.InvokeOptions): Promise<DeploymentSetsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:ecs/deploymentSets:DeploymentSets", {
        "granularity": args.granularity,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking DeploymentSets.
 */
export interface DeploymentSetsArgs {
    /**
     * The granularity of ECS DeploymentSet.Valid values: switch, host, rack.
     */
    granularity?: string;
    /**
     * A list of ECS DeploymentSet IDs.
     */
    ids?: string[];
    /**
     * A Name Regex of ECS DeploymentSet.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by DeploymentSets.
 */
export interface DeploymentSetsResult {
    /**
     * The collection of ECS DeploymentSet query.
     */
    readonly deploymentSets: outputs.ecs.DeploymentSetsDeploymentSet[];
    /**
     * The granularity of ECS DeploymentSet.
     */
    readonly granularity?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The total count of ECS DeploymentSet query.
     */
    readonly totalCount: number;
}

export function deploymentSetsOutput(args?: DeploymentSetsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<DeploymentSetsResult> {
    return pulumi.output(args).apply(a => deploymentSets(a, opts))
}

/**
 * A collection of arguments for invoking DeploymentSets.
 */
export interface DeploymentSetsOutputArgs {
    /**
     * The granularity of ECS DeploymentSet.Valid values: switch, host, rack.
     */
    granularity?: pulumi.Input<string>;
    /**
     * A list of ECS DeploymentSet IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Name Regex of ECS DeploymentSet.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}