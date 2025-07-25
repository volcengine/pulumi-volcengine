// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of apig upstreams
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.apig.getUpstreams({
 *     gatewayId: "gd13d8c6eq1emkiunq6p0",
 *     ids: [
 *         "ud18p5krj5ce3htvrd0v0",
 *         "ud18ouitrjp6fhvu61n7g",
 *     ],
 * });
 * ```
 */
export function getUpstreams(args?: GetUpstreamsArgs, opts?: pulumi.InvokeOptions): Promise<GetUpstreamsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:apig/getUpstreams:getUpstreams", {
        "gatewayId": args.gatewayId,
        "ids": args.ids,
        "name": args.name,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "resourceType": args.resourceType,
        "sourceType": args.sourceType,
    }, opts);
}

/**
 * A collection of arguments for invoking getUpstreams.
 */
export interface GetUpstreamsArgs {
    /**
     * The id of api gateway.
     */
    gatewayId?: string;
    /**
     * A list of apig upstream IDs.
     */
    ids?: string[];
    /**
     * The name of apig upstream. This field support fuzzy query.
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
     * The resource type of apig upstream. Valid values: `Console`, `Ingress`.
     */
    resourceType?: string;
    /**
     * The source type of apig upstream. Valid values: `VeFaas`, `ECS`, `FixedIP`, `K8S`, `Nacos`, `Domain`, `AIProvider`, `VeMLP`.
     */
    sourceType?: string;
}

/**
 * A collection of values returned by getUpstreams.
 */
export interface GetUpstreamsResult {
    /**
     * The id of api gateway.
     */
    readonly gatewayId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The name of apig upstream version.
     */
    readonly name?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The resource type of apig upstream.
     */
    readonly resourceType?: string;
    /**
     * The source type of apig upstream.
     */
    readonly sourceType?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
    /**
     * The collection of query.
     */
    readonly upstreams: outputs.apig.GetUpstreamsUpstream[];
}
/**
 * Use this data source to query detailed information of apig upstreams
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.apig.getUpstreams({
 *     gatewayId: "gd13d8c6eq1emkiunq6p0",
 *     ids: [
 *         "ud18p5krj5ce3htvrd0v0",
 *         "ud18ouitrjp6fhvu61n7g",
 *     ],
 * });
 * ```
 */
export function getUpstreamsOutput(args?: GetUpstreamsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUpstreamsResult> {
    return pulumi.output(args).apply((a: any) => getUpstreams(a, opts))
}

/**
 * A collection of arguments for invoking getUpstreams.
 */
export interface GetUpstreamsOutputArgs {
    /**
     * The id of api gateway.
     */
    gatewayId?: pulumi.Input<string>;
    /**
     * A list of apig upstream IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of apig upstream. This field support fuzzy query.
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
     * The resource type of apig upstream. Valid values: `Console`, `Ingress`.
     */
    resourceType?: pulumi.Input<string>;
    /**
     * The source type of apig upstream. Valid values: `VeFaas`, `ECS`, `FixedIP`, `K8S`, `Nacos`, `Domain`, `AIProvider`, `VeMLP`.
     */
    sourceType?: pulumi.Input<string>;
}
