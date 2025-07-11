// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of apig gateways
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.apig.getGateways({
 *     ids: [
 *         "gd13d8c6eq1emkiunq6p0",
 *         "gd07fq7pte3scmnaj1b1g",
 *     ],
 * });
 * ```
 */
/** @deprecated volcengine.apig.Gateways has been deprecated in favor of volcengine.apig.getGateways */
export function gateways(args?: GatewaysArgs, opts?: pulumi.InvokeOptions): Promise<GatewaysResult> {
    pulumi.log.warn("gateways is deprecated: volcengine.apig.Gateways has been deprecated in favor of volcengine.apig.getGateways")
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:apig/gateways:Gateways", {
        "ids": args.ids,
        "name": args.name,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "projectName": args.projectName,
        "status": args.status,
        "tags": args.tags,
        "type": args.type,
        "vpcIds": args.vpcIds,
    }, opts);
}

/**
 * A collection of arguments for invoking Gateways.
 */
export interface GatewaysArgs {
    /**
     * A list of api gateway IDs.
     */
    ids?: string[];
    /**
     * The name of api gateway. This field support fuzzy query.
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
     * The project name of api gateway.
     */
    projectName?: string;
    /**
     * The status of api gateway.
     */
    status?: string;
    /**
     * Tags.
     */
    tags?: inputs.apig.GatewaysTag[];
    /**
     * The type of api gateway.
     */
    type?: string;
    /**
     * A list of vpc IDs.
     */
    vpcIds?: string[];
}

/**
 * A collection of values returned by Gateways.
 */
export interface GatewaysResult {
    /**
     * The collection of query.
     */
    readonly gateways: outputs.apig.GatewaysGateway[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The name of the api gateway.
     */
    readonly name?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The project name of the api gateway.
     */
    readonly projectName?: string;
    /**
     * The status of the api gateway.
     */
    readonly status?: string;
    /**
     * Tags.
     */
    readonly tags?: outputs.apig.GatewaysTag[];
    /**
     * The total count of query.
     */
    readonly totalCount: number;
    /**
     * The type of the api gateway.
     */
    readonly type?: string;
    readonly vpcIds?: string[];
}
/**
 * Use this data source to query detailed information of apig gateways
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.apig.getGateways({
 *     ids: [
 *         "gd13d8c6eq1emkiunq6p0",
 *         "gd07fq7pte3scmnaj1b1g",
 *     ],
 * });
 * ```
 */
/** @deprecated volcengine.apig.Gateways has been deprecated in favor of volcengine.apig.getGateways */
export function gatewaysOutput(args?: GatewaysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GatewaysResult> {
    return pulumi.output(args).apply((a: any) => gateways(a, opts))
}

/**
 * A collection of arguments for invoking Gateways.
 */
export interface GatewaysOutputArgs {
    /**
     * A list of api gateway IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of api gateway. This field support fuzzy query.
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
     * The project name of api gateway.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The status of api gateway.
     */
    status?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.apig.GatewaysTagArgs>[]>;
    /**
     * The type of api gateway.
     */
    type?: pulumi.Input<string>;
    /**
     * A list of vpc IDs.
     */
    vpcIds?: pulumi.Input<pulumi.Input<string>[]>;
}
