// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage waf bot analyse protect rule
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.waf.BotAnalyseProtectRule("foo", {
 *     accurateGroup: {
 *         accurateRules: [
 *             {
 *                 httpObj: "request.uri",
 *                 objType: 1,
 *                 opretar: 2,
 *                 property: 0,
 *                 valueString: "tf",
 *             },
 *             {
 *                 httpObj: "request.schema",
 *                 objType: 0,
 *                 opretar: 2,
 *                 property: 0,
 *                 valueString: "tf-2",
 *             },
 *         ],
 *         logic: 2,
 *     },
 *     actionAfterVerification: 1,
 *     actionType: 1,
 *     effectTime: 1000,
 *     enable: 1,
 *     exemptionTime: 60,
 *     field: "HEADER:User-Agent",
 *     host: "www.tf-test.com",
 *     path: "/mod",
 *     pathThreshold: 1000,
 *     projectName: "default",
 *     rulePriority: 3,
 *     singleProportion: 0.25,
 *     singleThreshold: 100,
 *     statisticalDuration: 50,
 *     statisticalType: 2,
 * });
 * ```
 *
 * ## Import
 *
 * WafBotAnalyseProtectRule can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:waf/botAnalyseProtectRule:BotAnalyseProtectRule default resource_id:bot_space:host
 * ```
 */
export class BotAnalyseProtectRule extends pulumi.CustomResource {
    /**
     * Get an existing BotAnalyseProtectRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BotAnalyseProtectRuleState, opts?: pulumi.CustomResourceOptions): BotAnalyseProtectRule {
        return new BotAnalyseProtectRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:waf/botAnalyseProtectRule:BotAnalyseProtectRule';

    /**
     * Returns true if the given object is an instance of BotAnalyseProtectRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BotAnalyseProtectRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BotAnalyseProtectRule.__pulumiType;
    }

    /**
     * Advanced conditions.
     */
    public readonly accurateGroup!: pulumi.Output<outputs.waf.BotAnalyseProtectRuleAccurateGroup | undefined>;
    /**
     * Perform the action after verification/challenge.
     */
    public readonly actionAfterVerification!: pulumi.Output<number | undefined>;
    /**
     * perform the action.
     */
    public readonly actionType!: pulumi.Output<number>;
    /**
     * Limit the duration.
     */
    public readonly effectTime!: pulumi.Output<number>;
    /**
     * Whether to enable the rules.
     */
    public readonly enable!: pulumi.Output<number>;
    /**
     * The number of statistical protection rules enabled under the current domain name.
     */
    public /*out*/ readonly enableCount!: pulumi.Output<number>;
    /**
     * Exemption time takes effect when the execution action is human-machine challenge /JS/ Proof of work.
     */
    public readonly exemptionTime!: pulumi.Output<number | undefined>;
    /**
     * Statistical objects, with multiple objects separated by commas.
     */
    public readonly field!: pulumi.Output<string>;
    /**
     * Website domain names that require the setting of protection rules.
     */
    public readonly host!: pulumi.Output<string>;
    /**
     * The name of rule.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The requested path.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * The path access frequency threshold is enabled when StatisticalType=1.
     */
    public readonly pathThreshold!: pulumi.Output<number | undefined>;
    /**
     * The Name of the affiliated project resource.
     */
    public readonly projectName!: pulumi.Output<string | undefined>;
    /**
     * Details of the rule group.
     */
    public /*out*/ readonly ruleGroups!: pulumi.Output<outputs.waf.BotAnalyseProtectRuleRuleGroup[]>;
    /**
     * Priority of rule effectiveness.
     */
    public readonly rulePriority!: pulumi.Output<number>;
    /**
     * The IP proportion of the same statistical object needs to be configured when StatisticalType=3.
     */
    public readonly singleProportion!: pulumi.Output<number | undefined>;
    /**
     * The maximum number of ips of the same statistical object is enabled when StatisticalType=2.
     */
    public readonly singleThreshold!: pulumi.Output<number>;
    /**
     * The duration of statistics.
     */
    public readonly statisticalDuration!: pulumi.Output<number>;
    /**
     * Statistical content and methods.
     */
    public readonly statisticalType!: pulumi.Output<number>;
    /**
     * The total number of statistical protection rules under the current domain name.
     */
    public /*out*/ readonly totalCount!: pulumi.Output<number>;

    /**
     * Create a BotAnalyseProtectRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BotAnalyseProtectRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BotAnalyseProtectRuleArgs | BotAnalyseProtectRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BotAnalyseProtectRuleState | undefined;
            resourceInputs["accurateGroup"] = state ? state.accurateGroup : undefined;
            resourceInputs["actionAfterVerification"] = state ? state.actionAfterVerification : undefined;
            resourceInputs["actionType"] = state ? state.actionType : undefined;
            resourceInputs["effectTime"] = state ? state.effectTime : undefined;
            resourceInputs["enable"] = state ? state.enable : undefined;
            resourceInputs["enableCount"] = state ? state.enableCount : undefined;
            resourceInputs["exemptionTime"] = state ? state.exemptionTime : undefined;
            resourceInputs["field"] = state ? state.field : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["pathThreshold"] = state ? state.pathThreshold : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["ruleGroups"] = state ? state.ruleGroups : undefined;
            resourceInputs["rulePriority"] = state ? state.rulePriority : undefined;
            resourceInputs["singleProportion"] = state ? state.singleProportion : undefined;
            resourceInputs["singleThreshold"] = state ? state.singleThreshold : undefined;
            resourceInputs["statisticalDuration"] = state ? state.statisticalDuration : undefined;
            resourceInputs["statisticalType"] = state ? state.statisticalType : undefined;
            resourceInputs["totalCount"] = state ? state.totalCount : undefined;
        } else {
            const args = argsOrState as BotAnalyseProtectRuleArgs | undefined;
            if ((!args || args.actionType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'actionType'");
            }
            if ((!args || args.effectTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'effectTime'");
            }
            if ((!args || args.enable === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enable'");
            }
            if ((!args || args.field === undefined) && !opts.urn) {
                throw new Error("Missing required property 'field'");
            }
            if ((!args || args.host === undefined) && !opts.urn) {
                throw new Error("Missing required property 'host'");
            }
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            if ((!args || args.rulePriority === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rulePriority'");
            }
            if ((!args || args.singleThreshold === undefined) && !opts.urn) {
                throw new Error("Missing required property 'singleThreshold'");
            }
            if ((!args || args.statisticalDuration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'statisticalDuration'");
            }
            if ((!args || args.statisticalType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'statisticalType'");
            }
            resourceInputs["accurateGroup"] = args ? args.accurateGroup : undefined;
            resourceInputs["actionAfterVerification"] = args ? args.actionAfterVerification : undefined;
            resourceInputs["actionType"] = args ? args.actionType : undefined;
            resourceInputs["effectTime"] = args ? args.effectTime : undefined;
            resourceInputs["enable"] = args ? args.enable : undefined;
            resourceInputs["exemptionTime"] = args ? args.exemptionTime : undefined;
            resourceInputs["field"] = args ? args.field : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["pathThreshold"] = args ? args.pathThreshold : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["rulePriority"] = args ? args.rulePriority : undefined;
            resourceInputs["singleProportion"] = args ? args.singleProportion : undefined;
            resourceInputs["singleThreshold"] = args ? args.singleThreshold : undefined;
            resourceInputs["statisticalDuration"] = args ? args.statisticalDuration : undefined;
            resourceInputs["statisticalType"] = args ? args.statisticalType : undefined;
            resourceInputs["enableCount"] = undefined /*out*/;
            resourceInputs["ruleGroups"] = undefined /*out*/;
            resourceInputs["totalCount"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BotAnalyseProtectRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BotAnalyseProtectRule resources.
 */
export interface BotAnalyseProtectRuleState {
    /**
     * Advanced conditions.
     */
    accurateGroup?: pulumi.Input<inputs.waf.BotAnalyseProtectRuleAccurateGroup>;
    /**
     * Perform the action after verification/challenge.
     */
    actionAfterVerification?: pulumi.Input<number>;
    /**
     * perform the action.
     */
    actionType?: pulumi.Input<number>;
    /**
     * Limit the duration.
     */
    effectTime?: pulumi.Input<number>;
    /**
     * Whether to enable the rules.
     */
    enable?: pulumi.Input<number>;
    /**
     * The number of statistical protection rules enabled under the current domain name.
     */
    enableCount?: pulumi.Input<number>;
    /**
     * Exemption time takes effect when the execution action is human-machine challenge /JS/ Proof of work.
     */
    exemptionTime?: pulumi.Input<number>;
    /**
     * Statistical objects, with multiple objects separated by commas.
     */
    field?: pulumi.Input<string>;
    /**
     * Website domain names that require the setting of protection rules.
     */
    host?: pulumi.Input<string>;
    /**
     * The name of rule.
     */
    name?: pulumi.Input<string>;
    /**
     * The requested path.
     */
    path?: pulumi.Input<string>;
    /**
     * The path access frequency threshold is enabled when StatisticalType=1.
     */
    pathThreshold?: pulumi.Input<number>;
    /**
     * The Name of the affiliated project resource.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Details of the rule group.
     */
    ruleGroups?: pulumi.Input<pulumi.Input<inputs.waf.BotAnalyseProtectRuleRuleGroup>[]>;
    /**
     * Priority of rule effectiveness.
     */
    rulePriority?: pulumi.Input<number>;
    /**
     * The IP proportion of the same statistical object needs to be configured when StatisticalType=3.
     */
    singleProportion?: pulumi.Input<number>;
    /**
     * The maximum number of ips of the same statistical object is enabled when StatisticalType=2.
     */
    singleThreshold?: pulumi.Input<number>;
    /**
     * The duration of statistics.
     */
    statisticalDuration?: pulumi.Input<number>;
    /**
     * Statistical content and methods.
     */
    statisticalType?: pulumi.Input<number>;
    /**
     * The total number of statistical protection rules under the current domain name.
     */
    totalCount?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a BotAnalyseProtectRule resource.
 */
export interface BotAnalyseProtectRuleArgs {
    /**
     * Advanced conditions.
     */
    accurateGroup?: pulumi.Input<inputs.waf.BotAnalyseProtectRuleAccurateGroup>;
    /**
     * Perform the action after verification/challenge.
     */
    actionAfterVerification?: pulumi.Input<number>;
    /**
     * perform the action.
     */
    actionType: pulumi.Input<number>;
    /**
     * Limit the duration.
     */
    effectTime: pulumi.Input<number>;
    /**
     * Whether to enable the rules.
     */
    enable: pulumi.Input<number>;
    /**
     * Exemption time takes effect when the execution action is human-machine challenge /JS/ Proof of work.
     */
    exemptionTime?: pulumi.Input<number>;
    /**
     * Statistical objects, with multiple objects separated by commas.
     */
    field: pulumi.Input<string>;
    /**
     * Website domain names that require the setting of protection rules.
     */
    host: pulumi.Input<string>;
    /**
     * The name of rule.
     */
    name?: pulumi.Input<string>;
    /**
     * The requested path.
     */
    path: pulumi.Input<string>;
    /**
     * The path access frequency threshold is enabled when StatisticalType=1.
     */
    pathThreshold?: pulumi.Input<number>;
    /**
     * The Name of the affiliated project resource.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Priority of rule effectiveness.
     */
    rulePriority: pulumi.Input<number>;
    /**
     * The IP proportion of the same statistical object needs to be configured when StatisticalType=3.
     */
    singleProportion?: pulumi.Input<number>;
    /**
     * The maximum number of ips of the same statistical object is enabled when StatisticalType=2.
     */
    singleThreshold: pulumi.Input<number>;
    /**
     * The duration of statistics.
     */
    statisticalDuration: pulumi.Input<number>;
    /**
     * Statistical content and methods.
     */
    statisticalType: pulumi.Input<number>;
}
