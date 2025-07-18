// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage tls schedule sql task
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.tls.ScheduleSqlTask("foo", {
 *     description: "tf-test",
 *     destRegion: "cn-beijing",
 *     destTopicId: "b966e41a-d6a6-4999-bd75-39962xxxxxx",
 *     processEndTime: 1751295600,
 *     processSqlDelay: 60,
 *     processStartTime: 1751212980,
 *     processTimeWindow: "@m-15m,@m",
 *     query: "* | SELECT * limit 10000",
 *     requestCycle: {
 *         cronTab: "0 10 * * *",
 *         cronTimeZone: "GMT+08:00",
 *         time: 1,
 *         type: "CronTab",
 *     },
 *     status: 1,
 *     taskName: "tf-test",
 *     topicId: "8ba48bd7-2493-4300-b1d0-cb760bxxxxxx",
 * });
 * ```
 *
 * ## Import
 *
 * ScheduleSqlTask can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:tls/scheduleSqlTask:ScheduleSqlTask default resource_id
 * ```
 */
export class ScheduleSqlTask extends pulumi.CustomResource {
    /**
     * Get an existing ScheduleSqlTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScheduleSqlTaskState, opts?: pulumi.CustomResourceOptions): ScheduleSqlTask {
        return new ScheduleSqlTask(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:tls/scheduleSqlTask:ScheduleSqlTask';

    /**
     * Returns true if the given object is an instance of ScheduleSqlTask.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScheduleSqlTask {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScheduleSqlTask.__pulumiType;
    }

    /**
     * A simple description of the timed SQL analysis task.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The region to which the target log topic belongs. The default is the current region.
     */
    public readonly destRegion!: pulumi.Output<string | undefined>;
    /**
     * The target log topic ID used for storing the result data of timed SQL analysis.
     */
    public readonly destTopicId!: pulumi.Output<string>;
    /**
     * Schedule the end time of the timed SQL analysis task in the format of a second-level timestamp.
     */
    public readonly processEndTime!: pulumi.Output<number | undefined>;
    /**
     * The delay time of each scheduling. The value range is from 0 to 120, and the unit is seconds.
     */
    public readonly processSqlDelay!: pulumi.Output<number>;
    /**
     * The start time of the scheduled SQL analysis task, that is, the time when the first instance is created. The format is a timestamp at the second level.
     */
    public readonly processStartTime!: pulumi.Output<number>;
    /**
     * SQL time window, which refers to the time range for log retrieval and analysis when a timed SQL analysis task is running, is in a left-closed and right-open format.
     */
    public readonly processTimeWindow!: pulumi.Output<string>;
    /**
     * The retrieval and analysis statements for the regular execution of timed SQL analysis tasks should conform to the retrieval and analysis syntax of the log service.
     */
    public readonly query!: pulumi.Output<string>;
    /**
     * The scheduling cycle of timed SQL analysis tasks.
     */
    public readonly requestCycle!: pulumi.Output<outputs.tls.ScheduleSqlTaskRequestCycle>;
    /**
     * Whether to start the scheduled SQL analysis task immediately after completing the task configuration.
     */
    public readonly status!: pulumi.Output<number>;
    /**
     * The Name of timed SQL analysis task.
     */
    public readonly taskName!: pulumi.Output<string>;
    /**
     * The log topic ID where the original log to be analyzed for scheduled SQL is located.
     */
    public readonly topicId!: pulumi.Output<string>;

    /**
     * Create a ScheduleSqlTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScheduleSqlTaskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScheduleSqlTaskArgs | ScheduleSqlTaskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ScheduleSqlTaskState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destRegion"] = state ? state.destRegion : undefined;
            resourceInputs["destTopicId"] = state ? state.destTopicId : undefined;
            resourceInputs["processEndTime"] = state ? state.processEndTime : undefined;
            resourceInputs["processSqlDelay"] = state ? state.processSqlDelay : undefined;
            resourceInputs["processStartTime"] = state ? state.processStartTime : undefined;
            resourceInputs["processTimeWindow"] = state ? state.processTimeWindow : undefined;
            resourceInputs["query"] = state ? state.query : undefined;
            resourceInputs["requestCycle"] = state ? state.requestCycle : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["taskName"] = state ? state.taskName : undefined;
            resourceInputs["topicId"] = state ? state.topicId : undefined;
        } else {
            const args = argsOrState as ScheduleSqlTaskArgs | undefined;
            if ((!args || args.destTopicId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destTopicId'");
            }
            if ((!args || args.processSqlDelay === undefined) && !opts.urn) {
                throw new Error("Missing required property 'processSqlDelay'");
            }
            if ((!args || args.processStartTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'processStartTime'");
            }
            if ((!args || args.processTimeWindow === undefined) && !opts.urn) {
                throw new Error("Missing required property 'processTimeWindow'");
            }
            if ((!args || args.query === undefined) && !opts.urn) {
                throw new Error("Missing required property 'query'");
            }
            if ((!args || args.requestCycle === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requestCycle'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            if ((!args || args.taskName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskName'");
            }
            if ((!args || args.topicId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topicId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destRegion"] = args ? args.destRegion : undefined;
            resourceInputs["destTopicId"] = args ? args.destTopicId : undefined;
            resourceInputs["processEndTime"] = args ? args.processEndTime : undefined;
            resourceInputs["processSqlDelay"] = args ? args.processSqlDelay : undefined;
            resourceInputs["processStartTime"] = args ? args.processStartTime : undefined;
            resourceInputs["processTimeWindow"] = args ? args.processTimeWindow : undefined;
            resourceInputs["query"] = args ? args.query : undefined;
            resourceInputs["requestCycle"] = args ? args.requestCycle : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["taskName"] = args ? args.taskName : undefined;
            resourceInputs["topicId"] = args ? args.topicId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ScheduleSqlTask.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScheduleSqlTask resources.
 */
export interface ScheduleSqlTaskState {
    /**
     * A simple description of the timed SQL analysis task.
     */
    description?: pulumi.Input<string>;
    /**
     * The region to which the target log topic belongs. The default is the current region.
     */
    destRegion?: pulumi.Input<string>;
    /**
     * The target log topic ID used for storing the result data of timed SQL analysis.
     */
    destTopicId?: pulumi.Input<string>;
    /**
     * Schedule the end time of the timed SQL analysis task in the format of a second-level timestamp.
     */
    processEndTime?: pulumi.Input<number>;
    /**
     * The delay time of each scheduling. The value range is from 0 to 120, and the unit is seconds.
     */
    processSqlDelay?: pulumi.Input<number>;
    /**
     * The start time of the scheduled SQL analysis task, that is, the time when the first instance is created. The format is a timestamp at the second level.
     */
    processStartTime?: pulumi.Input<number>;
    /**
     * SQL time window, which refers to the time range for log retrieval and analysis when a timed SQL analysis task is running, is in a left-closed and right-open format.
     */
    processTimeWindow?: pulumi.Input<string>;
    /**
     * The retrieval and analysis statements for the regular execution of timed SQL analysis tasks should conform to the retrieval and analysis syntax of the log service.
     */
    query?: pulumi.Input<string>;
    /**
     * The scheduling cycle of timed SQL analysis tasks.
     */
    requestCycle?: pulumi.Input<inputs.tls.ScheduleSqlTaskRequestCycle>;
    /**
     * Whether to start the scheduled SQL analysis task immediately after completing the task configuration.
     */
    status?: pulumi.Input<number>;
    /**
     * The Name of timed SQL analysis task.
     */
    taskName?: pulumi.Input<string>;
    /**
     * The log topic ID where the original log to be analyzed for scheduled SQL is located.
     */
    topicId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ScheduleSqlTask resource.
 */
export interface ScheduleSqlTaskArgs {
    /**
     * A simple description of the timed SQL analysis task.
     */
    description?: pulumi.Input<string>;
    /**
     * The region to which the target log topic belongs. The default is the current region.
     */
    destRegion?: pulumi.Input<string>;
    /**
     * The target log topic ID used for storing the result data of timed SQL analysis.
     */
    destTopicId: pulumi.Input<string>;
    /**
     * Schedule the end time of the timed SQL analysis task in the format of a second-level timestamp.
     */
    processEndTime?: pulumi.Input<number>;
    /**
     * The delay time of each scheduling. The value range is from 0 to 120, and the unit is seconds.
     */
    processSqlDelay: pulumi.Input<number>;
    /**
     * The start time of the scheduled SQL analysis task, that is, the time when the first instance is created. The format is a timestamp at the second level.
     */
    processStartTime: pulumi.Input<number>;
    /**
     * SQL time window, which refers to the time range for log retrieval and analysis when a timed SQL analysis task is running, is in a left-closed and right-open format.
     */
    processTimeWindow: pulumi.Input<string>;
    /**
     * The retrieval and analysis statements for the regular execution of timed SQL analysis tasks should conform to the retrieval and analysis syntax of the log service.
     */
    query: pulumi.Input<string>;
    /**
     * The scheduling cycle of timed SQL analysis tasks.
     */
    requestCycle: pulumi.Input<inputs.tls.ScheduleSqlTaskRequestCycle>;
    /**
     * Whether to start the scheduled SQL analysis task immediately after completing the task configuration.
     */
    status: pulumi.Input<number>;
    /**
     * The Name of timed SQL analysis task.
     */
    taskName: pulumi.Input<string>;
    /**
     * The log topic ID where the original log to be analyzed for scheduled SQL is located.
     */
    topicId: pulumi.Input<string>;
}
