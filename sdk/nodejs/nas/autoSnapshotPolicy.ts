// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage nas auto snapshot policy
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.nas.AutoSnapshotPolicy("foo", {
 *     autoSnapshotPolicyName: "acc-test-auto_snapshot_policy",
 *     repeatWeekdays: "1,3,5,7",
 *     retentionDays: 20,
 *     timePoints: "0,7,17",
 * });
 * ```
 *
 * ## Import
 *
 * NasAutoSnapshotPolicy can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:nas/autoSnapshotPolicy:AutoSnapshotPolicy default resource_id
 * ```
 */
export class AutoSnapshotPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AutoSnapshotPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutoSnapshotPolicyState, opts?: pulumi.CustomResourceOptions): AutoSnapshotPolicy {
        return new AutoSnapshotPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:nas/autoSnapshotPolicy:AutoSnapshotPolicy';

    /**
     * Returns true if the given object is an instance of AutoSnapshotPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AutoSnapshotPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AutoSnapshotPolicy.__pulumiType;
    }

    /**
     * The name of the auto snapshot policy.
     */
    public readonly autoSnapshotPolicyName!: pulumi.Output<string>;
    /**
     * The create time of auto snapshot policy.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The count of file system which auto snapshot policy bind.
     */
    public /*out*/ readonly fileSystemCount!: pulumi.Output<number>;
    /**
     * The repeat weekdays of the auto snapshot policy. Support setting multiple dates, separated by English commas. Valid values: `1` ~ `7`.
     */
    public readonly repeatWeekdays!: pulumi.Output<string>;
    /**
     * The retention days of the auto snapshot policy. Valid values: -1(permanent) or 1 ~ 65536. Default is 30.
     */
    public readonly retentionDays!: pulumi.Output<number>;
    /**
     * The status of auto snapshot policy.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The time points of the auto snapshot policy. Support setting multiple dates, separated by English commas. Valid values: `0` ~ `23`.
     */
    public readonly timePoints!: pulumi.Output<string>;

    /**
     * Create a AutoSnapshotPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AutoSnapshotPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AutoSnapshotPolicyArgs | AutoSnapshotPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AutoSnapshotPolicyState | undefined;
            resourceInputs["autoSnapshotPolicyName"] = state ? state.autoSnapshotPolicyName : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["fileSystemCount"] = state ? state.fileSystemCount : undefined;
            resourceInputs["repeatWeekdays"] = state ? state.repeatWeekdays : undefined;
            resourceInputs["retentionDays"] = state ? state.retentionDays : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["timePoints"] = state ? state.timePoints : undefined;
        } else {
            const args = argsOrState as AutoSnapshotPolicyArgs | undefined;
            if ((!args || args.autoSnapshotPolicyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoSnapshotPolicyName'");
            }
            if ((!args || args.repeatWeekdays === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repeatWeekdays'");
            }
            if ((!args || args.timePoints === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timePoints'");
            }
            resourceInputs["autoSnapshotPolicyName"] = args ? args.autoSnapshotPolicyName : undefined;
            resourceInputs["repeatWeekdays"] = args ? args.repeatWeekdays : undefined;
            resourceInputs["retentionDays"] = args ? args.retentionDays : undefined;
            resourceInputs["timePoints"] = args ? args.timePoints : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["fileSystemCount"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AutoSnapshotPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AutoSnapshotPolicy resources.
 */
export interface AutoSnapshotPolicyState {
    /**
     * The name of the auto snapshot policy.
     */
    autoSnapshotPolicyName?: pulumi.Input<string>;
    /**
     * The create time of auto snapshot policy.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The count of file system which auto snapshot policy bind.
     */
    fileSystemCount?: pulumi.Input<number>;
    /**
     * The repeat weekdays of the auto snapshot policy. Support setting multiple dates, separated by English commas. Valid values: `1` ~ `7`.
     */
    repeatWeekdays?: pulumi.Input<string>;
    /**
     * The retention days of the auto snapshot policy. Valid values: -1(permanent) or 1 ~ 65536. Default is 30.
     */
    retentionDays?: pulumi.Input<number>;
    /**
     * The status of auto snapshot policy.
     */
    status?: pulumi.Input<string>;
    /**
     * The time points of the auto snapshot policy. Support setting multiple dates, separated by English commas. Valid values: `0` ~ `23`.
     */
    timePoints?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AutoSnapshotPolicy resource.
 */
export interface AutoSnapshotPolicyArgs {
    /**
     * The name of the auto snapshot policy.
     */
    autoSnapshotPolicyName: pulumi.Input<string>;
    /**
     * The repeat weekdays of the auto snapshot policy. Support setting multiple dates, separated by English commas. Valid values: `1` ~ `7`.
     */
    repeatWeekdays: pulumi.Input<string>;
    /**
     * The retention days of the auto snapshot policy. Valid values: -1(permanent) or 1 ~ 65536. Default is 30.
     */
    retentionDays?: pulumi.Input<number>;
    /**
     * The time points of the auto snapshot policy. Support setting multiple dates, separated by English commas. Valid values: `0` ~ `23`.
     */
    timePoints: pulumi.Input<string>;
}
