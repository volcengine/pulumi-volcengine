// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vmp instance types
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vmp.getInstanceTypes({
 *     ids: ["vmp.standard.15d"],
 * });
 * ```
 */
export function getInstanceTypes(args?: GetInstanceTypesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceTypesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:vmp/getInstanceTypes:getInstanceTypes", {
        "ids": args.ids,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceTypes.
 */
export interface GetInstanceTypesArgs {
    /**
     * A list of Instance Type IDs.
     */
    ids?: string[];
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getInstanceTypes.
 */
export interface GetInstanceTypesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The collection of query.
     */
    readonly instanceTypes: outputs.vmp.GetInstanceTypesInstanceType[];
    readonly outputFile?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of vmp instance types
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vmp.getInstanceTypes({
 *     ids: ["vmp.standard.15d"],
 * });
 * ```
 */
export function getInstanceTypesOutput(args?: GetInstanceTypesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceTypesResult> {
    return pulumi.output(args).apply((a: any) => getInstanceTypes(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceTypes.
 */
export interface GetInstanceTypesOutputArgs {
    /**
     * A list of Instance Type IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
