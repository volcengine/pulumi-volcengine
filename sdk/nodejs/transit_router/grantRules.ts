// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of transit router grant rules
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.transit_router.getGrantRules({
 *     transitRouterId: "tr-2bzy39uy6u3282dx0efxiqyq0",
 * });
 * ```
 */
/** @deprecated volcengine.transit_router.GrantRules has been deprecated in favor of volcengine.transit_router.getGrantRules */
export function grantRules(args: GrantRulesArgs, opts?: pulumi.InvokeOptions): Promise<GrantRulesResult> {
    pulumi.log.warn("grantRules is deprecated: volcengine.transit_router.GrantRules has been deprecated in favor of volcengine.transit_router.getGrantRules")

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:transit_router/grantRules:GrantRules", {
        "grantAccountId": args.grantAccountId,
        "outputFile": args.outputFile,
        "transitRouterId": args.transitRouterId,
    }, opts);
}

/**
 * A collection of arguments for invoking GrantRules.
 */
export interface GrantRulesArgs {
    /**
     * The id of the grant account.
     */
    grantAccountId?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The id of the transit router.
     */
    transitRouterId: string;
}

/**
 * A collection of values returned by GrantRules.
 */
export interface GrantRulesResult {
    /**
     * The id of the grant account.
     */
    readonly grantAccountId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
    /**
     * The collection of query.
     */
    readonly rules: outputs.transit_router.GrantRulesRule[];
    /**
     * The total count of query.
     */
    readonly totalCount: number;
    /**
     * The id of the transaction router.
     */
    readonly transitRouterId: string;
}
/**
 * Use this data source to query detailed information of transit router grant rules
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.transit_router.getGrantRules({
 *     transitRouterId: "tr-2bzy39uy6u3282dx0efxiqyq0",
 * });
 * ```
 */
/** @deprecated volcengine.transit_router.GrantRules has been deprecated in favor of volcengine.transit_router.getGrantRules */
export function grantRulesOutput(args: GrantRulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GrantRulesResult> {
    return pulumi.output(args).apply((a: any) => grantRules(a, opts))
}

/**
 * A collection of arguments for invoking GrantRules.
 */
export interface GrantRulesOutputArgs {
    /**
     * The id of the grant account.
     */
    grantAccountId?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The id of the transit router.
     */
    transitRouterId: pulumi.Input<string>;
}
