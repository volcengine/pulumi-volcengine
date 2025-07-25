// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vefaas functions
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.vefaas.getFunctions({});
 * ```
 */
/** @deprecated volcengine.vefaas.Functions has been deprecated in favor of volcengine.vefaas.getFunctions */
export function functions(args?: FunctionsArgs, opts?: pulumi.InvokeOptions): Promise<FunctionsResult> {
    pulumi.log.warn("functions is deprecated: volcengine.vefaas.Functions has been deprecated in favor of volcengine.vefaas.getFunctions")
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:vefaas/functions:Functions", {
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking Functions.
 */
export interface FunctionsArgs {
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by Functions.
 */
export interface FunctionsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The collection of query.
     */
    readonly items: outputs.vefaas.FunctionsItem[];
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of vefaas functions
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.vefaas.getFunctions({});
 * ```
 */
/** @deprecated volcengine.vefaas.Functions has been deprecated in favor of volcengine.vefaas.getFunctions */
export function functionsOutput(args?: FunctionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<FunctionsResult> {
    return pulumi.output(args).apply((a: any) => functions(a, opts))
}

/**
 * A collection of arguments for invoking Functions.
 */
export interface FunctionsOutputArgs {
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
