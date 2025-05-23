// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of organization service control policies
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.organization.getServiceControlPolicies({
 *     policyType: "Custom",
 *     query: "test",
 * });
 * ```
 */
export function getServiceControlPolicies(args?: GetServiceControlPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceControlPoliciesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:organization/getServiceControlPolicies:getServiceControlPolicies", {
        "outputFile": args.outputFile,
        "policyType": args.policyType,
        "query": args.query,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceControlPolicies.
 */
export interface GetServiceControlPoliciesArgs {
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The type of policy. The value can be System or Custom.
     */
    policyType?: string;
    /**
     * Query policies, support policy name or description.
     */
    query?: string;
}

/**
 * A collection of values returned by getServiceControlPolicies.
 */
export interface GetServiceControlPoliciesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
    /**
     * The collection of Policy query.
     */
    readonly policies: outputs.organization.GetServiceControlPoliciesPolicy[];
    /**
     * The type of the Policy.
     */
    readonly policyType?: string;
    readonly query?: string;
    /**
     * The total count of Policy query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of organization service control policies
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.organization.getServiceControlPolicies({
 *     policyType: "Custom",
 *     query: "test",
 * });
 * ```
 */
export function getServiceControlPoliciesOutput(args?: GetServiceControlPoliciesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceControlPoliciesResult> {
    return pulumi.output(args).apply((a: any) => getServiceControlPolicies(a, opts))
}

/**
 * A collection of arguments for invoking getServiceControlPolicies.
 */
export interface GetServiceControlPoliciesOutputArgs {
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The type of policy. The value can be System or Custom.
     */
    policyType?: pulumi.Input<string>;
    /**
     * Query policies, support policy name or description.
     */
    query?: pulumi.Input<string>;
}
