// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vmp rules
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vmp.getRules({
 *     kind: "Recording",
 *     workspaceId: "baa02ffb-6f22-43c4-841b-ecf90ded****",
 * });
 * ```
 */
/** @deprecated volcengine.vmp.Rules has been deprecated in favor of volcengine.vmp.getRules */
export function rules(args: RulesArgs, opts?: pulumi.InvokeOptions): Promise<RulesResult> {
    pulumi.log.warn("rules is deprecated: volcengine.vmp.Rules has been deprecated in favor of volcengine.vmp.getRules")

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:vmp/rules:Rules", {
        "kind": args.kind,
        "name": args.name,
        "outputFile": args.outputFile,
        "ruleFileNames": args.ruleFileNames,
        "ruleGroupNames": args.ruleGroupNames,
        "status": args.status,
        "workspaceId": args.workspaceId,
    }, opts);
}

/**
 * A collection of arguments for invoking Rules.
 */
export interface RulesArgs {
    /**
     * The kind of rule.
     */
    kind: string;
    /**
     * The name of rule.
     */
    name?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The name of rule file.
     */
    ruleFileNames?: string[];
    /**
     * The name of rule group.
     */
    ruleGroupNames?: string[];
    /**
     * The status of rule.
     */
    status?: string;
    /**
     * The id of workspace.
     */
    workspaceId: string;
}

/**
 * A collection of values returned by Rules.
 */
export interface RulesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The kind of rule.
     */
    readonly kind: string;
    /**
     * The name of rule.
     */
    readonly name?: string;
    readonly outputFile?: string;
    readonly ruleFileNames?: string[];
    readonly ruleGroupNames?: string[];
    /**
     * The collection of query.
     */
    readonly rules: outputs.vmp.RulesRule[];
    /**
     * The status of rule.
     */
    readonly status?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
    readonly workspaceId: string;
}
/**
 * Use this data source to query detailed information of vmp rules
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vmp.getRules({
 *     kind: "Recording",
 *     workspaceId: "baa02ffb-6f22-43c4-841b-ecf90ded****",
 * });
 * ```
 */
/** @deprecated volcengine.vmp.Rules has been deprecated in favor of volcengine.vmp.getRules */
export function rulesOutput(args: RulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<RulesResult> {
    return pulumi.output(args).apply((a: any) => rules(a, opts))
}

/**
 * A collection of arguments for invoking Rules.
 */
export interface RulesOutputArgs {
    /**
     * The kind of rule.
     */
    kind: pulumi.Input<string>;
    /**
     * The name of rule.
     */
    name?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The name of rule file.
     */
    ruleFileNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of rule group.
     */
    ruleGroupNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The status of rule.
     */
    status?: pulumi.Input<string>;
    /**
     * The id of workspace.
     */
    workspaceId: pulumi.Input<string>;
}
