// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vmp notify group policies
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooNotifyGroupPolicy = new volcengine.vmp.NotifyGroupPolicy("fooNotifyGroupPolicy", {
 *     description: "acc-test-1",
 *     levels: [
 *         {
 *             level: "P2",
 *             groupBies: ["__rule__"],
 *             groupWait: "35",
 *             groupInterval: "30",
 *             repeatInterval: "30",
 *         },
 *         {
 *             level: "P0",
 *             groupBies: ["__rule__"],
 *             groupWait: "30",
 *             groupInterval: "30",
 *             repeatInterval: "30",
 *         },
 *         {
 *             level: "P1",
 *             groupBies: ["__rule__"],
 *             groupWait: "40",
 *             groupInterval: "45",
 *             repeatInterval: "30",
 *         },
 *     ],
 * });
 * const fooNotifyGroupPolicies = volcengine.vmp.getNotifyGroupPoliciesOutput({
 *     ids: [fooNotifyGroupPolicy.id],
 * });
 * ```
 */
export function getNotifyGroupPolicies(args?: GetNotifyGroupPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<GetNotifyGroupPoliciesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:vmp/getNotifyGroupPolicies:getNotifyGroupPolicies", {
        "ids": args.ids,
        "name": args.name,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getNotifyGroupPolicies.
 */
export interface GetNotifyGroupPoliciesArgs {
    /**
     * A list of notify group policy ids.
     */
    ids?: string[];
    /**
     * The name of notify group policy.
     */
    name?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getNotifyGroupPolicies.
 */
export interface GetNotifyGroupPoliciesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The name of notify group policy.
     */
    readonly name?: string;
    /**
     * The list of notify group policies.
     */
    readonly notifyPolicies: outputs.vmp.GetNotifyGroupPoliciesNotifyPolicy[];
    readonly outputFile?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of vmp notify group policies
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooNotifyGroupPolicy = new volcengine.vmp.NotifyGroupPolicy("fooNotifyGroupPolicy", {
 *     description: "acc-test-1",
 *     levels: [
 *         {
 *             level: "P2",
 *             groupBies: ["__rule__"],
 *             groupWait: "35",
 *             groupInterval: "30",
 *             repeatInterval: "30",
 *         },
 *         {
 *             level: "P0",
 *             groupBies: ["__rule__"],
 *             groupWait: "30",
 *             groupInterval: "30",
 *             repeatInterval: "30",
 *         },
 *         {
 *             level: "P1",
 *             groupBies: ["__rule__"],
 *             groupWait: "40",
 *             groupInterval: "45",
 *             repeatInterval: "30",
 *         },
 *     ],
 * });
 * const fooNotifyGroupPolicies = volcengine.vmp.getNotifyGroupPoliciesOutput({
 *     ids: [fooNotifyGroupPolicy.id],
 * });
 * ```
 */
export function getNotifyGroupPoliciesOutput(args?: GetNotifyGroupPoliciesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNotifyGroupPoliciesResult> {
    return pulumi.output(args).apply((a: any) => getNotifyGroupPolicies(a, opts))
}

/**
 * A collection of arguments for invoking getNotifyGroupPolicies.
 */
export interface GetNotifyGroupPoliciesOutputArgs {
    /**
     * A list of notify group policy ids.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of notify group policy.
     */
    name?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
