// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cloud monitor event rules
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.cloud_monitor.EventRules({
 *     ruleName: "tftest",
 * });
 * ```
 */
export function eventRules(args?: EventRulesArgs, opts?: pulumi.InvokeOptions): Promise<EventRulesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:cloud_monitor/eventRules:EventRules", {
        "outputFile": args.outputFile,
        "ruleName": args.ruleName,
        "source": args.source,
    }, opts);
}

/**
 * A collection of arguments for invoking EventRules.
 */
export interface EventRulesArgs {
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * Rule name, search rules by name using fuzzy search.
     */
    ruleName?: string;
    /**
     * Event source.
     */
    source?: string;
}

/**
 * A collection of values returned by EventRules.
 */
export interface EventRulesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
    /**
     * The name of the rule.
     */
    readonly ruleName?: string;
    /**
     * The collection of query.
     */
    readonly rules: outputs.cloud_monitor.EventRulesRule[];
    /**
     * Event source corresponding to pattern matching.
     */
    readonly source?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of cloud monitor event rules
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.cloud_monitor.EventRules({
 *     ruleName: "tftest",
 * });
 * ```
 */
export function eventRulesOutput(args?: EventRulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<EventRulesResult> {
    return pulumi.output(args).apply((a: any) => eventRules(a, opts))
}

/**
 * A collection of arguments for invoking EventRules.
 */
export interface EventRulesOutputArgs {
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * Rule name, search rules by name using fuzzy search.
     */
    ruleName?: pulumi.Input<string>;
    /**
     * Event source.
     */
    source?: pulumi.Input<string>;
}