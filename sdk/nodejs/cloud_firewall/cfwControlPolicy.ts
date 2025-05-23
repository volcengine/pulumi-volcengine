// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage cfw control policy
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooCfwAddressBook = new volcengine.cloud_firewall.CfwAddressBook("fooCfwAddressBook", {
 *     groupName: "acc-test-address-book",
 *     description: "acc-test",
 *     groupType: "ip",
 *     addressLists: [
 *         "192.168.1.1",
 *         "192.168.2.2",
 *     ],
 * });
 * const fooCfwControlPolicy = new volcengine.cloud_firewall.CfwControlPolicy("fooCfwControlPolicy", {
 *     direction: "in",
 *     action: "accept",
 *     destinationType: "group",
 *     destination: fooCfwAddressBook.id,
 *     proto: "TCP",
 *     sourceType: "net",
 *     source: "0.0.0.0/0",
 *     description: "acc-test-control-policy",
 *     destPortType: "port",
 *     destPort: "300",
 *     repeatType: "Weekly",
 *     repeatStartTime: "01:00",
 *     repeatEndTime: "11:00",
 *     repeatDays: [
 *         2,
 *         5,
 *     ],
 *     startTime: 1736092800,
 *     endTime: 1738339140,
 *     priority: 1,
 *     status: true,
 * });
 * ```
 *
 * ## Import
 *
 * ControlPolicy can be imported using the direction:rule_id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:cloud_firewall/cfwControlPolicy:CfwControlPolicy default resource_id
 * ```
 */
export class CfwControlPolicy extends pulumi.CustomResource {
    /**
     * Get an existing CfwControlPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CfwControlPolicyState, opts?: pulumi.CustomResourceOptions): CfwControlPolicy {
        return new CfwControlPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:cloud_firewall/cfwControlPolicy:CfwControlPolicy';

    /**
     * Returns true if the given object is an instance of CfwControlPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CfwControlPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CfwControlPolicy.__pulumiType;
    }

    /**
     * The account id of the control policy.
     */
    public /*out*/ readonly accountId!: pulumi.Output<string>;
    /**
     * The action of the control policy. Valid values: `accept`, `deny`, `monitor`.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * The description of the control policy.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The dest port of the control policy.
     */
    public readonly destPort!: pulumi.Output<string>;
    /**
     * The dest port type of the control policy. Valid values: `port`, `group`.
     */
    public readonly destPortType!: pulumi.Output<string>;
    /**
     * The destination of the control policy.
     */
    public readonly destination!: pulumi.Output<string>;
    /**
     * The destination type of the control policy. Valid values: `net`, `group`, `location`, `domain`.
     */
    public readonly destinationType!: pulumi.Output<string>;
    /**
     * The direction of the control policy. Valid values: `in`, `out`.
     */
    public readonly direction!: pulumi.Output<string>;
    /**
     * The effect status of the control policy. 1: Not yet effective, 2: Issued in progress, 3: Effective.
     */
    public /*out*/ readonly effectStatus!: pulumi.Output<number>;
    /**
     * The end time of the control policy. Unix timestamp, fields need to be precise to 23:59:00 of the set date.
     * When the value of repeatType is one of `Once`, `Daily`, `Weekly`, `Monthly`, this field is required.
     */
    public readonly endTime!: pulumi.Output<number | undefined>;
    /**
     * The hit count of the control policy.
     */
    public /*out*/ readonly hitCnt!: pulumi.Output<number>;
    /**
     * Whether the control policy is effected.
     */
    public /*out*/ readonly isEffected!: pulumi.Output<boolean>;
    /**
     * The priority of the control policy.
     */
    public /*out*/ readonly prio!: pulumi.Output<number>;
    /**
     * The priority of the control policy. Default is 0. This field is only effective when creating a control policy.0 means lowest priority, 1 means highest priority. The priority increases in order from 1, with lower priority indicating higher priority.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * The proto of the control policy. Valid values: `TCP`, `ICMP`, `UDP`, `ANY`. When the destinationType is `domain`, The proto must be `TCP`.
     */
    public readonly proto!: pulumi.Output<string>;
    /**
     * The repeat days of the control policy. When the value of repeatType is one of `Weekly`, `Monthly`, this field is required.
     * When the repeatType is `Weekly`, the valid value range is 0~6.
     * When the repeatType is `Monthly`, the valid value range is 1~31.
     */
    public readonly repeatDays!: pulumi.Output<number[] | undefined>;
    /**
     * The repeat end time of the control policy. Accurate to the minute, in the format of hh: mm. For example: 12:00.
     * When the value of repeatType is one of `Daily`, `Weekly`, `Monthly`, this field is required.
     */
    public readonly repeatEndTime!: pulumi.Output<string | undefined>;
    /**
     * The repeat start time of the control policy. Accurate to the minute, in the format of hh: mm. For example: 12:00.
     * When the value of repeatType is one of `Daily`, `Weekly`, `Monthly`, this field is required.
     */
    public readonly repeatStartTime!: pulumi.Output<string | undefined>;
    /**
     * The repeat type of the control policy. Valid values: `Permanent`, `Once`, `Daily`, `Weekly`, `Monthly`.
     */
    public readonly repeatType!: pulumi.Output<string>;
    /**
     * The rule id of the control policy.
     */
    public /*out*/ readonly ruleId!: pulumi.Output<string>;
    /**
     * The source of the control policy.
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * The source type of the control policy. Valid values: `net`, `group`, `location`.
     */
    public readonly sourceType!: pulumi.Output<string>;
    /**
     * The start time of the control policy. Unix timestamp, fields need to be precise to 23:59:00 of the set date.
     * When the value of repeatType is one of `Once`, `Daily`, `Weekly`, `Monthly`, this field is required.
     */
    public readonly startTime!: pulumi.Output<number | undefined>;
    /**
     * Whether to enable the control policy. Default is false.
     */
    public readonly status!: pulumi.Output<boolean>;
    /**
     * The update time of the control policy.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<number>;
    /**
     * The use count of the control policy.
     */
    public /*out*/ readonly useCount!: pulumi.Output<number>;

    /**
     * Create a CfwControlPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CfwControlPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CfwControlPolicyArgs | CfwControlPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CfwControlPolicyState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destPort"] = state ? state.destPort : undefined;
            resourceInputs["destPortType"] = state ? state.destPortType : undefined;
            resourceInputs["destination"] = state ? state.destination : undefined;
            resourceInputs["destinationType"] = state ? state.destinationType : undefined;
            resourceInputs["direction"] = state ? state.direction : undefined;
            resourceInputs["effectStatus"] = state ? state.effectStatus : undefined;
            resourceInputs["endTime"] = state ? state.endTime : undefined;
            resourceInputs["hitCnt"] = state ? state.hitCnt : undefined;
            resourceInputs["isEffected"] = state ? state.isEffected : undefined;
            resourceInputs["prio"] = state ? state.prio : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["proto"] = state ? state.proto : undefined;
            resourceInputs["repeatDays"] = state ? state.repeatDays : undefined;
            resourceInputs["repeatEndTime"] = state ? state.repeatEndTime : undefined;
            resourceInputs["repeatStartTime"] = state ? state.repeatStartTime : undefined;
            resourceInputs["repeatType"] = state ? state.repeatType : undefined;
            resourceInputs["ruleId"] = state ? state.ruleId : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["sourceType"] = state ? state.sourceType : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
            resourceInputs["useCount"] = state ? state.useCount : undefined;
        } else {
            const args = argsOrState as CfwControlPolicyArgs | undefined;
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.destination === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destination'");
            }
            if ((!args || args.destinationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationType'");
            }
            if ((!args || args.direction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'direction'");
            }
            if ((!args || args.proto === undefined) && !opts.urn) {
                throw new Error("Missing required property 'proto'");
            }
            if ((!args || args.source === undefined) && !opts.urn) {
                throw new Error("Missing required property 'source'");
            }
            if ((!args || args.sourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceType'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destPort"] = args ? args.destPort : undefined;
            resourceInputs["destPortType"] = args ? args.destPortType : undefined;
            resourceInputs["destination"] = args ? args.destination : undefined;
            resourceInputs["destinationType"] = args ? args.destinationType : undefined;
            resourceInputs["direction"] = args ? args.direction : undefined;
            resourceInputs["endTime"] = args ? args.endTime : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["proto"] = args ? args.proto : undefined;
            resourceInputs["repeatDays"] = args ? args.repeatDays : undefined;
            resourceInputs["repeatEndTime"] = args ? args.repeatEndTime : undefined;
            resourceInputs["repeatStartTime"] = args ? args.repeatStartTime : undefined;
            resourceInputs["repeatType"] = args ? args.repeatType : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["sourceType"] = args ? args.sourceType : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["accountId"] = undefined /*out*/;
            resourceInputs["effectStatus"] = undefined /*out*/;
            resourceInputs["hitCnt"] = undefined /*out*/;
            resourceInputs["isEffected"] = undefined /*out*/;
            resourceInputs["prio"] = undefined /*out*/;
            resourceInputs["ruleId"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["useCount"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CfwControlPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CfwControlPolicy resources.
 */
export interface CfwControlPolicyState {
    /**
     * The account id of the control policy.
     */
    accountId?: pulumi.Input<string>;
    /**
     * The action of the control policy. Valid values: `accept`, `deny`, `monitor`.
     */
    action?: pulumi.Input<string>;
    /**
     * The description of the control policy.
     */
    description?: pulumi.Input<string>;
    /**
     * The dest port of the control policy.
     */
    destPort?: pulumi.Input<string>;
    /**
     * The dest port type of the control policy. Valid values: `port`, `group`.
     */
    destPortType?: pulumi.Input<string>;
    /**
     * The destination of the control policy.
     */
    destination?: pulumi.Input<string>;
    /**
     * The destination type of the control policy. Valid values: `net`, `group`, `location`, `domain`.
     */
    destinationType?: pulumi.Input<string>;
    /**
     * The direction of the control policy. Valid values: `in`, `out`.
     */
    direction?: pulumi.Input<string>;
    /**
     * The effect status of the control policy. 1: Not yet effective, 2: Issued in progress, 3: Effective.
     */
    effectStatus?: pulumi.Input<number>;
    /**
     * The end time of the control policy. Unix timestamp, fields need to be precise to 23:59:00 of the set date.
     * When the value of repeatType is one of `Once`, `Daily`, `Weekly`, `Monthly`, this field is required.
     */
    endTime?: pulumi.Input<number>;
    /**
     * The hit count of the control policy.
     */
    hitCnt?: pulumi.Input<number>;
    /**
     * Whether the control policy is effected.
     */
    isEffected?: pulumi.Input<boolean>;
    /**
     * The priority of the control policy.
     */
    prio?: pulumi.Input<number>;
    /**
     * The priority of the control policy. Default is 0. This field is only effective when creating a control policy.0 means lowest priority, 1 means highest priority. The priority increases in order from 1, with lower priority indicating higher priority.
     */
    priority?: pulumi.Input<number>;
    /**
     * The proto of the control policy. Valid values: `TCP`, `ICMP`, `UDP`, `ANY`. When the destinationType is `domain`, The proto must be `TCP`.
     */
    proto?: pulumi.Input<string>;
    /**
     * The repeat days of the control policy. When the value of repeatType is one of `Weekly`, `Monthly`, this field is required.
     * When the repeatType is `Weekly`, the valid value range is 0~6.
     * When the repeatType is `Monthly`, the valid value range is 1~31.
     */
    repeatDays?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The repeat end time of the control policy. Accurate to the minute, in the format of hh: mm. For example: 12:00.
     * When the value of repeatType is one of `Daily`, `Weekly`, `Monthly`, this field is required.
     */
    repeatEndTime?: pulumi.Input<string>;
    /**
     * The repeat start time of the control policy. Accurate to the minute, in the format of hh: mm. For example: 12:00.
     * When the value of repeatType is one of `Daily`, `Weekly`, `Monthly`, this field is required.
     */
    repeatStartTime?: pulumi.Input<string>;
    /**
     * The repeat type of the control policy. Valid values: `Permanent`, `Once`, `Daily`, `Weekly`, `Monthly`.
     */
    repeatType?: pulumi.Input<string>;
    /**
     * The rule id of the control policy.
     */
    ruleId?: pulumi.Input<string>;
    /**
     * The source of the control policy.
     */
    source?: pulumi.Input<string>;
    /**
     * The source type of the control policy. Valid values: `net`, `group`, `location`.
     */
    sourceType?: pulumi.Input<string>;
    /**
     * The start time of the control policy. Unix timestamp, fields need to be precise to 23:59:00 of the set date.
     * When the value of repeatType is one of `Once`, `Daily`, `Weekly`, `Monthly`, this field is required.
     */
    startTime?: pulumi.Input<number>;
    /**
     * Whether to enable the control policy. Default is false.
     */
    status?: pulumi.Input<boolean>;
    /**
     * The update time of the control policy.
     */
    updateTime?: pulumi.Input<number>;
    /**
     * The use count of the control policy.
     */
    useCount?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a CfwControlPolicy resource.
 */
export interface CfwControlPolicyArgs {
    /**
     * The action of the control policy. Valid values: `accept`, `deny`, `monitor`.
     */
    action: pulumi.Input<string>;
    /**
     * The description of the control policy.
     */
    description?: pulumi.Input<string>;
    /**
     * The dest port of the control policy.
     */
    destPort?: pulumi.Input<string>;
    /**
     * The dest port type of the control policy. Valid values: `port`, `group`.
     */
    destPortType?: pulumi.Input<string>;
    /**
     * The destination of the control policy.
     */
    destination: pulumi.Input<string>;
    /**
     * The destination type of the control policy. Valid values: `net`, `group`, `location`, `domain`.
     */
    destinationType: pulumi.Input<string>;
    /**
     * The direction of the control policy. Valid values: `in`, `out`.
     */
    direction: pulumi.Input<string>;
    /**
     * The end time of the control policy. Unix timestamp, fields need to be precise to 23:59:00 of the set date.
     * When the value of repeatType is one of `Once`, `Daily`, `Weekly`, `Monthly`, this field is required.
     */
    endTime?: pulumi.Input<number>;
    /**
     * The priority of the control policy. Default is 0. This field is only effective when creating a control policy.0 means lowest priority, 1 means highest priority. The priority increases in order from 1, with lower priority indicating higher priority.
     */
    priority?: pulumi.Input<number>;
    /**
     * The proto of the control policy. Valid values: `TCP`, `ICMP`, `UDP`, `ANY`. When the destinationType is `domain`, The proto must be `TCP`.
     */
    proto: pulumi.Input<string>;
    /**
     * The repeat days of the control policy. When the value of repeatType is one of `Weekly`, `Monthly`, this field is required.
     * When the repeatType is `Weekly`, the valid value range is 0~6.
     * When the repeatType is `Monthly`, the valid value range is 1~31.
     */
    repeatDays?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The repeat end time of the control policy. Accurate to the minute, in the format of hh: mm. For example: 12:00.
     * When the value of repeatType is one of `Daily`, `Weekly`, `Monthly`, this field is required.
     */
    repeatEndTime?: pulumi.Input<string>;
    /**
     * The repeat start time of the control policy. Accurate to the minute, in the format of hh: mm. For example: 12:00.
     * When the value of repeatType is one of `Daily`, `Weekly`, `Monthly`, this field is required.
     */
    repeatStartTime?: pulumi.Input<string>;
    /**
     * The repeat type of the control policy. Valid values: `Permanent`, `Once`, `Daily`, `Weekly`, `Monthly`.
     */
    repeatType?: pulumi.Input<string>;
    /**
     * The source of the control policy.
     */
    source: pulumi.Input<string>;
    /**
     * The source type of the control policy. Valid values: `net`, `group`, `location`.
     */
    sourceType: pulumi.Input<string>;
    /**
     * The start time of the control policy. Unix timestamp, fields need to be precise to 23:59:00 of the set date.
     * When the value of repeatType is one of `Once`, `Daily`, `Weekly`, `Monthly`, this field is required.
     */
    startTime?: pulumi.Input<number>;
    /**
     * Whether to enable the control policy. Default is false.
     */
    status?: pulumi.Input<boolean>;
}
