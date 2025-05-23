// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cr endpoints
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.cr.getEndpoints({
 *     registry: "tf-1",
 * });
 * ```
 */
export function getEndpoints(args: GetEndpointsArgs, opts?: pulumi.InvokeOptions): Promise<GetEndpointsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:cr/getEndpoints:getEndpoints", {
        "outputFile": args.outputFile,
        "registry": args.registry,
    }, opts);
}

/**
 * A collection of arguments for invoking getEndpoints.
 */
export interface GetEndpointsArgs {
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The CR instance name.
     */
    registry: string;
}

/**
 * A collection of values returned by getEndpoints.
 */
export interface GetEndpointsResult {
    /**
     * The collection of endpoint query.
     */
    readonly endpoints: outputs.cr.GetEndpointsEndpoint[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
    /**
     * The name of CR instance.
     */
    readonly registry: string;
    /**
     * The total count of tag query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of cr endpoints
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.cr.getEndpoints({
 *     registry: "tf-1",
 * });
 * ```
 */
export function getEndpointsOutput(args: GetEndpointsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEndpointsResult> {
    return pulumi.output(args).apply((a: any) => getEndpoints(a, opts))
}

/**
 * A collection of arguments for invoking getEndpoints.
 */
export interface GetEndpointsOutputArgs {
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The CR instance name.
     */
    registry: pulumi.Input<string>;
}
