// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of private zone resolver endpoints
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.private_zone.ResolverEndpoints({});
 * ```
 */
export function resolverEndpoints(args?: ResolverEndpointsArgs, opts?: pulumi.InvokeOptions): Promise<ResolverEndpointsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:private_zone/resolverEndpoints:ResolverEndpoints", {
        "direction": args.direction,
        "name": args.name,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "status": args.status,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking ResolverEndpoints.
 */
export interface ResolverEndpointsArgs {
    /**
     * The direction of the private zone resolver endpoint.
     */
    direction?: string;
    /**
     * The name of the private zone resolver endpoint.
     */
    name?: string;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The status of the private zone resolver endpoint.
     */
    status?: string;
    /**
     * The vpc ID of the private zone resolver endpoint.
     */
    vpcId?: string;
}

/**
 * A collection of values returned by ResolverEndpoints.
 */
export interface ResolverEndpointsResult {
    /**
     * The direction of the endpoint.
     */
    readonly direction?: string;
    /**
     * The collection of query.
     */
    readonly endpoints: outputs.private_zone.ResolverEndpointsEndpoint[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the endpoint.
     */
    readonly name?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The status of the endpoint.
     */
    readonly status?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
    /**
     * The vpc id of the endpoint.
     */
    readonly vpcId?: string;
}
/**
 * Use this data source to query detailed information of private zone resolver endpoints
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.private_zone.ResolverEndpoints({});
 * ```
 */
export function resolverEndpointsOutput(args?: ResolverEndpointsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ResolverEndpointsResult> {
    return pulumi.output(args).apply((a: any) => resolverEndpoints(a, opts))
}

/**
 * A collection of arguments for invoking ResolverEndpoints.
 */
export interface ResolverEndpointsOutputArgs {
    /**
     * The direction of the private zone resolver endpoint.
     */
    direction?: pulumi.Input<string>;
    /**
     * The name of the private zone resolver endpoint.
     */
    name?: pulumi.Input<string>;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the private zone resolver endpoint.
     */
    status?: pulumi.Input<string>;
    /**
     * The vpc ID of the private zone resolver endpoint.
     */
    vpcId?: pulumi.Input<string>;
}