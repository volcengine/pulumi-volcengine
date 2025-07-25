// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vmp alerts
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vmp.getAlerts({
 *     ids: ["9a4f84-0868efcb795c2ac4-73cefd4b3263****"],
 * });
 * ```
 */
/** @deprecated volcengine.vmp.Alerts has been deprecated in favor of volcengine.vmp.getAlerts */
export function alerts(args?: AlertsArgs, opts?: pulumi.InvokeOptions): Promise<AlertsResult> {
    pulumi.log.warn("alerts is deprecated: volcengine.vmp.Alerts has been deprecated in favor of volcengine.vmp.getAlerts")
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:vmp/alerts:Alerts", {
        "alertingRuleIds": args.alertingRuleIds,
        "currentPhase": args.currentPhase,
        "desc": args.desc,
        "ids": args.ids,
        "level": args.level,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking Alerts.
 */
export interface AlertsArgs {
    /**
     * A list of alerting rule IDs.
     */
    alertingRuleIds?: string[];
    /**
     * The status of vmp alert. Valid values: `Pending`, `Active`, `Resolved`, `Disabled`.
     */
    currentPhase?: string;
    /**
     * Whether to use descending sorting.
     */
    desc?: boolean;
    /**
     * A list of vmp alert IDs.
     */
    ids?: string[];
    /**
     * The level of vmp alert. Valid values: `P0`, `P1`, `P2`.
     */
    level?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by Alerts.
 */
export interface AlertsResult {
    readonly alertingRuleIds?: string[];
    /**
     * The collection of query.
     */
    readonly alerts: outputs.vmp.AlertsAlert[];
    /**
     * The status of the vmp alert.
     */
    readonly currentPhase?: string;
    readonly desc?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The level of the vmp alerting rule.
     */
    readonly level?: string;
    readonly outputFile?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of vmp alerts
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vmp.getAlerts({
 *     ids: ["9a4f84-0868efcb795c2ac4-73cefd4b3263****"],
 * });
 * ```
 */
/** @deprecated volcengine.vmp.Alerts has been deprecated in favor of volcengine.vmp.getAlerts */
export function alertsOutput(args?: AlertsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<AlertsResult> {
    return pulumi.output(args).apply((a: any) => alerts(a, opts))
}

/**
 * A collection of arguments for invoking Alerts.
 */
export interface AlertsOutputArgs {
    /**
     * A list of alerting rule IDs.
     */
    alertingRuleIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The status of vmp alert. Valid values: `Pending`, `Active`, `Resolved`, `Disabled`.
     */
    currentPhase?: pulumi.Input<string>;
    /**
     * Whether to use descending sorting.
     */
    desc?: pulumi.Input<boolean>;
    /**
     * A list of vmp alert IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The level of vmp alert. Valid values: `P0`, `P1`, `P2`.
     */
    level?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
