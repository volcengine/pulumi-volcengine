// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage tls host group
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = new volcengine.tls.HostGroup("foo", {
 *     autoUpdate: false,
 *     hostGroupName: "tfgroup",
 *     hostGroupType: "Label",
 *     hostIdentifier: "tf-controller",
 *     serviceLogging: false,
 * });
 * ```
 *
 * ## Import
 *
 * Tls Host Group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:tls/hostGroup:HostGroup default edf052s21s*******dc15
 * ```
 */
export class HostGroup extends pulumi.CustomResource {
    /**
     * Get an existing HostGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostGroupState, opts?: pulumi.CustomResourceOptions): HostGroup {
        return new HostGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:tls/hostGroup:HostGroup';

    /**
     * Returns true if the given object is an instance of HostGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HostGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HostGroup.__pulumiType;
    }

    /**
     * The abnormal heartbeat status count of host.
     */
    public /*out*/ readonly abnormalHeartbeatStatusCount!: pulumi.Output<number>;
    /**
     * The latest version of log collector.
     */
    public /*out*/ readonly agentLatestVersion!: pulumi.Output<string>;
    /**
     * Whether enable auto update.
     */
    public readonly autoUpdate!: pulumi.Output<boolean | undefined>;
    /**
     * The create time of host group.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The count of host.
     */
    public /*out*/ readonly hostCount!: pulumi.Output<number>;
    /**
     * The name of host group.
     */
    public readonly hostGroupName!: pulumi.Output<string>;
    /**
     * The type of host group. The value can be IP or Label.
     */
    public readonly hostGroupType!: pulumi.Output<string>;
    /**
     * The identifier of host.
     */
    public readonly hostIdentifier!: pulumi.Output<string | undefined>;
    /**
     * The ip list of host group.
     */
    public readonly hostIpLists!: pulumi.Output<string[]>;
    /**
     * The project name of iam.
     */
    public readonly iamProjectName!: pulumi.Output<string | undefined>;
    /**
     * The modify time of host group.
     */
    public /*out*/ readonly modifyTime!: pulumi.Output<string>;
    /**
     * The normal heartbeat status count of host.
     */
    public /*out*/ readonly normalHeartbeatStatusCount!: pulumi.Output<number>;
    /**
     * The rule count of host.
     */
    public /*out*/ readonly ruleCount!: pulumi.Output<number>;
    /**
     * Whether enable service logging.
     */
    public readonly serviceLogging!: pulumi.Output<boolean | undefined>;
    /**
     * The update end time of log collector.
     */
    public readonly updateEndTime!: pulumi.Output<string>;
    /**
     * The update start time of log collector.
     */
    public readonly updateStartTime!: pulumi.Output<string>;

    /**
     * Create a HostGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostGroupArgs | HostGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HostGroupState | undefined;
            resourceInputs["abnormalHeartbeatStatusCount"] = state ? state.abnormalHeartbeatStatusCount : undefined;
            resourceInputs["agentLatestVersion"] = state ? state.agentLatestVersion : undefined;
            resourceInputs["autoUpdate"] = state ? state.autoUpdate : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["hostCount"] = state ? state.hostCount : undefined;
            resourceInputs["hostGroupName"] = state ? state.hostGroupName : undefined;
            resourceInputs["hostGroupType"] = state ? state.hostGroupType : undefined;
            resourceInputs["hostIdentifier"] = state ? state.hostIdentifier : undefined;
            resourceInputs["hostIpLists"] = state ? state.hostIpLists : undefined;
            resourceInputs["iamProjectName"] = state ? state.iamProjectName : undefined;
            resourceInputs["modifyTime"] = state ? state.modifyTime : undefined;
            resourceInputs["normalHeartbeatStatusCount"] = state ? state.normalHeartbeatStatusCount : undefined;
            resourceInputs["ruleCount"] = state ? state.ruleCount : undefined;
            resourceInputs["serviceLogging"] = state ? state.serviceLogging : undefined;
            resourceInputs["updateEndTime"] = state ? state.updateEndTime : undefined;
            resourceInputs["updateStartTime"] = state ? state.updateStartTime : undefined;
        } else {
            const args = argsOrState as HostGroupArgs | undefined;
            if ((!args || args.hostGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostGroupName'");
            }
            if ((!args || args.hostGroupType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostGroupType'");
            }
            resourceInputs["autoUpdate"] = args ? args.autoUpdate : undefined;
            resourceInputs["hostGroupName"] = args ? args.hostGroupName : undefined;
            resourceInputs["hostGroupType"] = args ? args.hostGroupType : undefined;
            resourceInputs["hostIdentifier"] = args ? args.hostIdentifier : undefined;
            resourceInputs["hostIpLists"] = args ? args.hostIpLists : undefined;
            resourceInputs["iamProjectName"] = args ? args.iamProjectName : undefined;
            resourceInputs["serviceLogging"] = args ? args.serviceLogging : undefined;
            resourceInputs["updateEndTime"] = args ? args.updateEndTime : undefined;
            resourceInputs["updateStartTime"] = args ? args.updateStartTime : undefined;
            resourceInputs["abnormalHeartbeatStatusCount"] = undefined /*out*/;
            resourceInputs["agentLatestVersion"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["hostCount"] = undefined /*out*/;
            resourceInputs["modifyTime"] = undefined /*out*/;
            resourceInputs["normalHeartbeatStatusCount"] = undefined /*out*/;
            resourceInputs["ruleCount"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HostGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HostGroup resources.
 */
export interface HostGroupState {
    /**
     * The abnormal heartbeat status count of host.
     */
    abnormalHeartbeatStatusCount?: pulumi.Input<number>;
    /**
     * The latest version of log collector.
     */
    agentLatestVersion?: pulumi.Input<string>;
    /**
     * Whether enable auto update.
     */
    autoUpdate?: pulumi.Input<boolean>;
    /**
     * The create time of host group.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The count of host.
     */
    hostCount?: pulumi.Input<number>;
    /**
     * The name of host group.
     */
    hostGroupName?: pulumi.Input<string>;
    /**
     * The type of host group. The value can be IP or Label.
     */
    hostGroupType?: pulumi.Input<string>;
    /**
     * The identifier of host.
     */
    hostIdentifier?: pulumi.Input<string>;
    /**
     * The ip list of host group.
     */
    hostIpLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The project name of iam.
     */
    iamProjectName?: pulumi.Input<string>;
    /**
     * The modify time of host group.
     */
    modifyTime?: pulumi.Input<string>;
    /**
     * The normal heartbeat status count of host.
     */
    normalHeartbeatStatusCount?: pulumi.Input<number>;
    /**
     * The rule count of host.
     */
    ruleCount?: pulumi.Input<number>;
    /**
     * Whether enable service logging.
     */
    serviceLogging?: pulumi.Input<boolean>;
    /**
     * The update end time of log collector.
     */
    updateEndTime?: pulumi.Input<string>;
    /**
     * The update start time of log collector.
     */
    updateStartTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HostGroup resource.
 */
export interface HostGroupArgs {
    /**
     * Whether enable auto update.
     */
    autoUpdate?: pulumi.Input<boolean>;
    /**
     * The name of host group.
     */
    hostGroupName: pulumi.Input<string>;
    /**
     * The type of host group. The value can be IP or Label.
     */
    hostGroupType: pulumi.Input<string>;
    /**
     * The identifier of host.
     */
    hostIdentifier?: pulumi.Input<string>;
    /**
     * The ip list of host group.
     */
    hostIpLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The project name of iam.
     */
    iamProjectName?: pulumi.Input<string>;
    /**
     * Whether enable service logging.
     */
    serviceLogging?: pulumi.Input<boolean>;
    /**
     * The update end time of log collector.
     */
    updateEndTime?: pulumi.Input<string>;
    /**
     * The update start time of log collector.
     */
    updateStartTime?: pulumi.Input<string>;
}