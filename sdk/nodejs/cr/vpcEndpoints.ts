// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cr vpc endpoints
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultVpcEndpoints = pulumi.output(volcengine.cr.VpcEndpoints({
 *     registry: "enterprise-1",
 *     statuses: [
 *         "Enabled",
 *         "Enabling",
 *         "Disabling",
 *         "Failed",
 *     ],
 * }));
 * ```
 */
export function vpcEndpoints(args: VpcEndpointsArgs, opts?: pulumi.InvokeOptions): Promise<VpcEndpointsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:cr/vpcEndpoints:VpcEndpoints", {
        "outputFile": args.outputFile,
        "registry": args.registry,
        "statuses": args.statuses,
    }, opts);
}

/**
 * A collection of arguments for invoking VpcEndpoints.
 */
export interface VpcEndpointsArgs {
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The CR registry name.
     */
    registry: string;
    /**
     * VPC access entry state array, used to filter out VPC access entries in the specified state. Available values are Enabling, Enabled, Disabling, Failed.
     */
    statuses?: string[];
}

/**
 * A collection of values returned by VpcEndpoints.
 */
export interface VpcEndpointsResult {
    /**
     * List of CR vpc endpoints.
     */
    readonly endpoints: outputs.cr.VpcEndpointsEndpoint[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
    /**
     * The name of CR registry.
     */
    readonly registry: string;
    readonly statuses?: string[];
    /**
     * The total count of CR vpc endpoints query.
     */
    readonly totalCount: number;
}

export function vpcEndpointsOutput(args: VpcEndpointsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<VpcEndpointsResult> {
    return pulumi.output(args).apply(a => vpcEndpoints(a, opts))
}

/**
 * A collection of arguments for invoking VpcEndpoints.
 */
export interface VpcEndpointsOutputArgs {
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The CR registry name.
     */
    registry: pulumi.Input<string>;
    /**
     * VPC access entry state array, used to filter out VPC access entries in the specified state. Available values are Enabling, Enabled, Disabling, Failed.
     */
    statuses?: pulumi.Input<pulumi.Input<string>[]>;
}