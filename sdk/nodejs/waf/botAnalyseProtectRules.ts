// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of waf bot analyse protect rules
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.waf.getBotAnalyseProtectRules({
 *     botSpace: "BotRepeat",
 *     host: "www.tf-test.com",
 * });
 * ```
 */
/** @deprecated volcengine.waf.BotAnalyseProtectRules has been deprecated in favor of volcengine.waf.getBotAnalyseProtectRules */
export function botAnalyseProtectRules(args: BotAnalyseProtectRulesArgs, opts?: pulumi.InvokeOptions): Promise<BotAnalyseProtectRulesResult> {
    pulumi.log.warn("botAnalyseProtectRules is deprecated: volcengine.waf.BotAnalyseProtectRules has been deprecated in favor of volcengine.waf.getBotAnalyseProtectRules")

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:waf/botAnalyseProtectRules:BotAnalyseProtectRules", {
        "botSpace": args.botSpace,
        "host": args.host,
        "name": args.name,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "path": args.path,
        "projectName": args.projectName,
        "ruleTag": args.ruleTag,
    }, opts);
}

/**
 * A collection of arguments for invoking BotAnalyseProtectRules.
 */
export interface BotAnalyseProtectRulesArgs {
    /**
     * Bot protection rule type.
     */
    botSpace: string;
    /**
     * Website domain names that require the setting of protection rules.
     */
    host: string;
    /**
     * The name of the rule.
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
     * Protective path.
     */
    path?: string;
    /**
     * The name of the project to which your domain names belong.
     */
    projectName?: string;
    /**
     * Unique identification of rules.
     */
    ruleTag?: string;
}

/**
 * A collection of values returned by BotAnalyseProtectRules.
 */
export interface BotAnalyseProtectRulesResult {
    readonly botSpace: string;
    /**
     * The details of the Bot rules.
     */
    readonly datas: outputs.waf.BotAnalyseProtectRulesData[];
    /**
     * The domain name where the protection rule is located.
     */
    readonly host: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of rule.
     */
    readonly name?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * Request path.
     */
    readonly path?: string;
    readonly projectName?: string;
    /**
     * Rule label, that is, the complete rule ID.
     */
    readonly ruleTag?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of waf bot analyse protect rules
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.waf.getBotAnalyseProtectRules({
 *     botSpace: "BotRepeat",
 *     host: "www.tf-test.com",
 * });
 * ```
 */
/** @deprecated volcengine.waf.BotAnalyseProtectRules has been deprecated in favor of volcengine.waf.getBotAnalyseProtectRules */
export function botAnalyseProtectRulesOutput(args: BotAnalyseProtectRulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<BotAnalyseProtectRulesResult> {
    return pulumi.output(args).apply((a: any) => botAnalyseProtectRules(a, opts))
}

/**
 * A collection of arguments for invoking BotAnalyseProtectRules.
 */
export interface BotAnalyseProtectRulesOutputArgs {
    /**
     * Bot protection rule type.
     */
    botSpace: pulumi.Input<string>;
    /**
     * Website domain names that require the setting of protection rules.
     */
    host: pulumi.Input<string>;
    /**
     * The name of the rule.
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
     * Protective path.
     */
    path?: pulumi.Input<string>;
    /**
     * The name of the project to which your domain names belong.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Unique identification of rules.
     */
    ruleTag?: pulumi.Input<string>;
}
