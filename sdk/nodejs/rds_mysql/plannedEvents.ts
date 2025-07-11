// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of rds mysql planned events
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.rds_mysql.getPlannedEvents({
 *     instanceId: "mysql-b51d37110dd1",
 * });
 * ```
 */
/** @deprecated volcengine.rds_mysql.PlannedEvents has been deprecated in favor of volcengine.rds_mysql.getPlannedEvents */
export function plannedEvents(args?: PlannedEventsArgs, opts?: pulumi.InvokeOptions): Promise<PlannedEventsResult> {
    pulumi.log.warn("plannedEvents is deprecated: volcengine.rds_mysql.PlannedEvents has been deprecated in favor of volcengine.rds_mysql.getPlannedEvents")
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:rds_mysql/plannedEvents:PlannedEvents", {
        "beginTime": args.beginTime,
        "endTime": args.endTime,
        "eventId": args.eventId,
        "eventTypes": args.eventTypes,
        "instanceId": args.instanceId,
        "outputFile": args.outputFile,
        "statuses": args.statuses,
    }, opts);
}

/**
 * A collection of arguments for invoking PlannedEvents.
 */
export interface PlannedEventsArgs {
    /**
     * The start time of the planned event.
     */
    beginTime?: string;
    /**
     * The end time of the planned event.
     */
    endTime?: string;
    /**
     * The id of the planned event.
     */
    eventId?: string;
    /**
     * The type of the planned event.
     */
    eventTypes?: string[];
    /**
     * The id of the instance.
     */
    instanceId?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The status of the planned event.
     */
    statuses?: string[];
}

/**
 * A collection of values returned by PlannedEvents.
 */
export interface PlannedEventsResult {
    readonly beginTime?: string;
    readonly endTime?: string;
    /**
     * The id of the planned event.
     */
    readonly eventId?: string;
    /**
     * The type of the planned event.
     */
    readonly eventTypes?: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The id of the instance.
     */
    readonly instanceId?: string;
    readonly outputFile?: string;
    /**
     * The collection of query.
     */
    readonly plannedEvents: outputs.rds_mysql.PlannedEventsPlannedEvent[];
    /**
     * Event status.
     */
    readonly statuses?: string[];
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of rds mysql planned events
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.rds_mysql.getPlannedEvents({
 *     instanceId: "mysql-b51d37110dd1",
 * });
 * ```
 */
/** @deprecated volcengine.rds_mysql.PlannedEvents has been deprecated in favor of volcengine.rds_mysql.getPlannedEvents */
export function plannedEventsOutput(args?: PlannedEventsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<PlannedEventsResult> {
    return pulumi.output(args).apply((a: any) => plannedEvents(a, opts))
}

/**
 * A collection of arguments for invoking PlannedEvents.
 */
export interface PlannedEventsOutputArgs {
    /**
     * The start time of the planned event.
     */
    beginTime?: pulumi.Input<string>;
    /**
     * The end time of the planned event.
     */
    endTime?: pulumi.Input<string>;
    /**
     * The id of the planned event.
     */
    eventId?: pulumi.Input<string>;
    /**
     * The type of the planned event.
     */
    eventTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The id of the instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the planned event.
     */
    statuses?: pulumi.Input<pulumi.Input<string>[]>;
}
