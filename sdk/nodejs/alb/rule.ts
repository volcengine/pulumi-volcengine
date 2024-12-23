// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage alb rule
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.alb.Rule("foo", {
 *     description: "test",
 *     domain: "www.test.com",
 *     listenerId: "lsn-1iidd19u4oni874adhezjkyj3",
 *     redirectConfig: {
 *         redirectDomain: "www.testtest.com",
 *         redirectHttpCode: "302",
 *         redirectPort: "555",
 *         redirectUri: "/testtest",
 *     },
 *     rewriteConfig: {
 *         rewritePath: "/test",
 *     },
 *     rewriteEnabled: "off",
 *     ruleAction: "Redirect",
 *     serverGroupId: "rsp-1g72w74y4umf42zbhq4k4hnln",
 *     trafficLimitEnabled: "off",
 *     trafficLimitQps: 100,
 *     url: "/test",
 * });
 * ```
 *
 * ## Import
 *
 * AlbRule can be imported using the listener id and rule id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:alb/rule:Rule default lsn-273yv0mhs5xj47fap8sehiiso:rule-****
 * ```
 */
export class Rule extends pulumi.CustomResource {
    /**
     * Get an existing Rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleState, opts?: pulumi.CustomResourceOptions): Rule {
        return new Rule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:alb/rule:Rule';

    /**
     * Returns true if the given object is an instance of Rule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rule.__pulumiType;
    }

    /**
     * The description of the Rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The domain of Rule.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * The ID of listener.
     */
    public readonly listenerId!: pulumi.Output<string>;
    /**
     * The redirect related configuration.
     */
    public readonly redirectConfig!: pulumi.Output<outputs.alb.RuleRedirectConfig | undefined>;
    /**
     * The list of rewrite configurations.
     */
    public readonly rewriteConfig!: pulumi.Output<outputs.alb.RuleRewriteConfig | undefined>;
    /**
     * Rewrite configuration switch for forwarding rules, only allows configuration and takes effect when RuleAction is empty (i.e., forwarding to server group). Only available for whitelist users, please submit an application to experience. Supported values are as follows:
     * on: enable.
     * off: disable.
     */
    public readonly rewriteEnabled!: pulumi.Output<string | undefined>;
    /**
     * The forwarding rule action, if this parameter is empty(`""`), forward to server group, if value is `Redirect`, will redirect.
     */
    public readonly ruleAction!: pulumi.Output<string>;
    /**
     * The ID of rule.
     */
    public /*out*/ readonly ruleId!: pulumi.Output<string>;
    /**
     * Server group ID, this parameter is required if `ruleAction` is empty.
     */
    public readonly serverGroupId!: pulumi.Output<string | undefined>;
    /**
     * Forwarding rule QPS rate limiting switch:
     * on: enable.
     * off: disable (default).
     */
    public readonly trafficLimitEnabled!: pulumi.Output<string | undefined>;
    /**
     * When Rules.N.TrafficLimitEnabled is turned on, this field is required. Requests per second. Valid values are between 100 and 100000.
     */
    public readonly trafficLimitQps!: pulumi.Output<number | undefined>;
    /**
     * The Url of Rule.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a Rule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleArgs | RuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["redirectConfig"] = state ? state.redirectConfig : undefined;
            resourceInputs["rewriteConfig"] = state ? state.rewriteConfig : undefined;
            resourceInputs["rewriteEnabled"] = state ? state.rewriteEnabled : undefined;
            resourceInputs["ruleAction"] = state ? state.ruleAction : undefined;
            resourceInputs["ruleId"] = state ? state.ruleId : undefined;
            resourceInputs["serverGroupId"] = state ? state.serverGroupId : undefined;
            resourceInputs["trafficLimitEnabled"] = state ? state.trafficLimitEnabled : undefined;
            resourceInputs["trafficLimitQps"] = state ? state.trafficLimitQps : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as RuleArgs | undefined;
            if ((!args || args.listenerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listenerId'");
            }
            if ((!args || args.ruleAction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleAction'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["listenerId"] = args ? args.listenerId : undefined;
            resourceInputs["redirectConfig"] = args ? args.redirectConfig : undefined;
            resourceInputs["rewriteConfig"] = args ? args.rewriteConfig : undefined;
            resourceInputs["rewriteEnabled"] = args ? args.rewriteEnabled : undefined;
            resourceInputs["ruleAction"] = args ? args.ruleAction : undefined;
            resourceInputs["serverGroupId"] = args ? args.serverGroupId : undefined;
            resourceInputs["trafficLimitEnabled"] = args ? args.trafficLimitEnabled : undefined;
            resourceInputs["trafficLimitQps"] = args ? args.trafficLimitQps : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["ruleId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Rule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Rule resources.
 */
export interface RuleState {
    /**
     * The description of the Rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The domain of Rule.
     */
    domain?: pulumi.Input<string>;
    /**
     * The ID of listener.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * The redirect related configuration.
     */
    redirectConfig?: pulumi.Input<inputs.alb.RuleRedirectConfig>;
    /**
     * The list of rewrite configurations.
     */
    rewriteConfig?: pulumi.Input<inputs.alb.RuleRewriteConfig>;
    /**
     * Rewrite configuration switch for forwarding rules, only allows configuration and takes effect when RuleAction is empty (i.e., forwarding to server group). Only available for whitelist users, please submit an application to experience. Supported values are as follows:
     * on: enable.
     * off: disable.
     */
    rewriteEnabled?: pulumi.Input<string>;
    /**
     * The forwarding rule action, if this parameter is empty(`""`), forward to server group, if value is `Redirect`, will redirect.
     */
    ruleAction?: pulumi.Input<string>;
    /**
     * The ID of rule.
     */
    ruleId?: pulumi.Input<string>;
    /**
     * Server group ID, this parameter is required if `ruleAction` is empty.
     */
    serverGroupId?: pulumi.Input<string>;
    /**
     * Forwarding rule QPS rate limiting switch:
     * on: enable.
     * off: disable (default).
     */
    trafficLimitEnabled?: pulumi.Input<string>;
    /**
     * When Rules.N.TrafficLimitEnabled is turned on, this field is required. Requests per second. Valid values are between 100 and 100000.
     */
    trafficLimitQps?: pulumi.Input<number>;
    /**
     * The Url of Rule.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Rule resource.
 */
export interface RuleArgs {
    /**
     * The description of the Rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The domain of Rule.
     */
    domain?: pulumi.Input<string>;
    /**
     * The ID of listener.
     */
    listenerId: pulumi.Input<string>;
    /**
     * The redirect related configuration.
     */
    redirectConfig?: pulumi.Input<inputs.alb.RuleRedirectConfig>;
    /**
     * The list of rewrite configurations.
     */
    rewriteConfig?: pulumi.Input<inputs.alb.RuleRewriteConfig>;
    /**
     * Rewrite configuration switch for forwarding rules, only allows configuration and takes effect when RuleAction is empty (i.e., forwarding to server group). Only available for whitelist users, please submit an application to experience. Supported values are as follows:
     * on: enable.
     * off: disable.
     */
    rewriteEnabled?: pulumi.Input<string>;
    /**
     * The forwarding rule action, if this parameter is empty(`""`), forward to server group, if value is `Redirect`, will redirect.
     */
    ruleAction: pulumi.Input<string>;
    /**
     * Server group ID, this parameter is required if `ruleAction` is empty.
     */
    serverGroupId?: pulumi.Input<string>;
    /**
     * Forwarding rule QPS rate limiting switch:
     * on: enable.
     * off: disable (default).
     */
    trafficLimitEnabled?: pulumi.Input<string>;
    /**
     * When Rules.N.TrafficLimitEnabled is turned on, this field is required. Requests per second. Valid values are between 100 and 100000.
     */
    trafficLimitQps?: pulumi.Input<number>;
    /**
     * The Url of Rule.
     */
    url?: pulumi.Input<string>;
}
