// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage cloud monitor event rule
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.cloud_monitor.EventRule("foo", {
 *     contactGroupIds: [
 *         "1737941730782699520",
 *         "1737940985502777344",
 *     ],
 *     contactMethods: [
 *         "Phone",
 *         "TLS",
 *         "MQ",
 *     ],
 *     effectiveTime: {
 *         endTime: "22:00",
 *         startTime: "01:00",
 *     },
 *     eventSource: "ecs",
 *     eventTypes: ["ecs:Disk:DiskError.Redeploy.Canceled"],
 *     filterPattern: {
 *         source: "ecs",
 *         types: ["ecs:Disk:DiskError.Redeploy.Canceled"],
 *     },
 *     level: "notice",
 *     messageQueues: [{
 *         instanceId: "kafka-cnoe4rfrsqfb1d64",
 *         region: "*****",
 *         topic: "tftest",
 *         type: "kafka",
 *         vpcId: "vpc-2d68hz41j7qio58ozfd6jxgtb",
 *     }],
 *     ruleName: "tftest1",
 *     status: "enable",
 *     tlsTargets: [{
 *         projectId: "17ba378d-de43-495e-8906-03ae6567b376",
 *         projectName: "tf-test",
 *         regionNameCn: "*****",
 *         regionNameEn: "*****",
 *         topicId: "7ce12237-6670-44a7-9d79-2e36961586e6",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * CloudMonitorEventRule can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:cloud_monitor/eventRule:EventRule default rule_id
 * ```
 */
export class EventRule extends pulumi.CustomResource {
    /**
     * Get an existing EventRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventRuleState, opts?: pulumi.CustomResourceOptions): EventRule {
        return new EventRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:cloud_monitor/eventRule:EventRule';

    /**
     * Returns true if the given object is an instance of EventRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventRule.__pulumiType;
    }

    /**
     * When the alarm notification method is phone, SMS, or email, the triggered alarm contact group ID.
     */
    public readonly contactGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * Alarm notification methods. Valid value: `Phone`, `Email`, `SMS`, `Webhook`: Alarm callback, `TLS`: Log Service, `MQ`: Message Queue Kafka.
     */
    public readonly contactMethods!: pulumi.Output<string[]>;
    /**
     * The description of the rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The rule takes effect at a certain time and will only be effective during this period.
     */
    public readonly effectiveTime!: pulumi.Output<outputs.cloud_monitor.EventRuleEffectiveTime>;
    /**
     * When the alarm notification method is alarm callback, it triggers the callback address.
     */
    public readonly endpoint!: pulumi.Output<string | undefined>;
    /**
     * Event source.
     */
    public readonly eventSource!: pulumi.Output<string>;
    /**
     * Event type.
     */
    public readonly eventTypes!: pulumi.Output<string[] | undefined>;
    /**
     * Filter mode, also known as event matching rules. Custom matching rules are not currently supported.
     */
    public readonly filterPattern!: pulumi.Output<outputs.cloud_monitor.EventRuleFilterPattern>;
    /**
     * Severity of alarm rules. Value can be `notice`, `warning`, `critical`.
     */
    public readonly level!: pulumi.Output<string>;
    /**
     * The triggered message queue when the alarm notification method is Kafka message queue.
     */
    public readonly messageQueues!: pulumi.Output<outputs.cloud_monitor.EventRuleMessageQueue[] | undefined>;
    /**
     * The name of the rule.
     */
    public readonly ruleName!: pulumi.Output<string>;
    /**
     * Rule status. `enable`: enable rule(default), `disable`: disable rule.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * The alarm method for log service triggers the configuration of the log service.
     */
    public readonly tlsTargets!: pulumi.Output<outputs.cloud_monitor.EventRuleTlsTarget[] | undefined>;

    /**
     * Create a EventRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventRuleArgs | EventRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventRuleState | undefined;
            resourceInputs["contactGroupIds"] = state ? state.contactGroupIds : undefined;
            resourceInputs["contactMethods"] = state ? state.contactMethods : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["effectiveTime"] = state ? state.effectiveTime : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["eventSource"] = state ? state.eventSource : undefined;
            resourceInputs["eventTypes"] = state ? state.eventTypes : undefined;
            resourceInputs["filterPattern"] = state ? state.filterPattern : undefined;
            resourceInputs["level"] = state ? state.level : undefined;
            resourceInputs["messageQueues"] = state ? state.messageQueues : undefined;
            resourceInputs["ruleName"] = state ? state.ruleName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tlsTargets"] = state ? state.tlsTargets : undefined;
        } else {
            const args = argsOrState as EventRuleArgs | undefined;
            if ((!args || args.contactMethods === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contactMethods'");
            }
            if ((!args || args.effectiveTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'effectiveTime'");
            }
            if ((!args || args.eventSource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventSource'");
            }
            if ((!args || args.filterPattern === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filterPattern'");
            }
            if ((!args || args.level === undefined) && !opts.urn) {
                throw new Error("Missing required property 'level'");
            }
            if ((!args || args.ruleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleName'");
            }
            resourceInputs["contactGroupIds"] = args ? args.contactGroupIds : undefined;
            resourceInputs["contactMethods"] = args ? args.contactMethods : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["effectiveTime"] = args ? args.effectiveTime : undefined;
            resourceInputs["endpoint"] = args ? args.endpoint : undefined;
            resourceInputs["eventSource"] = args ? args.eventSource : undefined;
            resourceInputs["eventTypes"] = args ? args.eventTypes : undefined;
            resourceInputs["filterPattern"] = args ? args.filterPattern : undefined;
            resourceInputs["level"] = args ? args.level : undefined;
            resourceInputs["messageQueues"] = args ? args.messageQueues : undefined;
            resourceInputs["ruleName"] = args ? args.ruleName : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tlsTargets"] = args ? args.tlsTargets : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EventRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventRule resources.
 */
export interface EventRuleState {
    /**
     * When the alarm notification method is phone, SMS, or email, the triggered alarm contact group ID.
     */
    contactGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Alarm notification methods. Valid value: `Phone`, `Email`, `SMS`, `Webhook`: Alarm callback, `TLS`: Log Service, `MQ`: Message Queue Kafka.
     */
    contactMethods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The rule takes effect at a certain time and will only be effective during this period.
     */
    effectiveTime?: pulumi.Input<inputs.cloud_monitor.EventRuleEffectiveTime>;
    /**
     * When the alarm notification method is alarm callback, it triggers the callback address.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * Event source.
     */
    eventSource?: pulumi.Input<string>;
    /**
     * Event type.
     */
    eventTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Filter mode, also known as event matching rules. Custom matching rules are not currently supported.
     */
    filterPattern?: pulumi.Input<inputs.cloud_monitor.EventRuleFilterPattern>;
    /**
     * Severity of alarm rules. Value can be `notice`, `warning`, `critical`.
     */
    level?: pulumi.Input<string>;
    /**
     * The triggered message queue when the alarm notification method is Kafka message queue.
     */
    messageQueues?: pulumi.Input<pulumi.Input<inputs.cloud_monitor.EventRuleMessageQueue>[]>;
    /**
     * The name of the rule.
     */
    ruleName?: pulumi.Input<string>;
    /**
     * Rule status. `enable`: enable rule(default), `disable`: disable rule.
     */
    status?: pulumi.Input<string>;
    /**
     * The alarm method for log service triggers the configuration of the log service.
     */
    tlsTargets?: pulumi.Input<pulumi.Input<inputs.cloud_monitor.EventRuleTlsTarget>[]>;
}

/**
 * The set of arguments for constructing a EventRule resource.
 */
export interface EventRuleArgs {
    /**
     * When the alarm notification method is phone, SMS, or email, the triggered alarm contact group ID.
     */
    contactGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Alarm notification methods. Valid value: `Phone`, `Email`, `SMS`, `Webhook`: Alarm callback, `TLS`: Log Service, `MQ`: Message Queue Kafka.
     */
    contactMethods: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The rule takes effect at a certain time and will only be effective during this period.
     */
    effectiveTime: pulumi.Input<inputs.cloud_monitor.EventRuleEffectiveTime>;
    /**
     * When the alarm notification method is alarm callback, it triggers the callback address.
     */
    endpoint?: pulumi.Input<string>;
    /**
     * Event source.
     */
    eventSource: pulumi.Input<string>;
    /**
     * Event type.
     */
    eventTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Filter mode, also known as event matching rules. Custom matching rules are not currently supported.
     */
    filterPattern: pulumi.Input<inputs.cloud_monitor.EventRuleFilterPattern>;
    /**
     * Severity of alarm rules. Value can be `notice`, `warning`, `critical`.
     */
    level: pulumi.Input<string>;
    /**
     * The triggered message queue when the alarm notification method is Kafka message queue.
     */
    messageQueues?: pulumi.Input<pulumi.Input<inputs.cloud_monitor.EventRuleMessageQueue>[]>;
    /**
     * The name of the rule.
     */
    ruleName: pulumi.Input<string>;
    /**
     * Rule status. `enable`: enable rule(default), `disable`: disable rule.
     */
    status?: pulumi.Input<string>;
    /**
     * The alarm method for log service triggers the configuration of the log service.
     */
    tlsTargets?: pulumi.Input<pulumi.Input<inputs.cloud_monitor.EventRuleTlsTarget>[]>;
}