// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds instances
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.rds.getInstances({
 *     instanceId: "mysql-0fdd3bab2e7c",
 * });
 * ```
 */
export function getInstances(args?: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:rds/getInstances:getInstances", {
        "createEndTime": args.createEndTime,
        "createStartTime": args.createStartTime,
        "instanceId": args.instanceId,
        "instanceStatus": args.instanceStatus,
        "instanceType": args.instanceType,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "region": args.region,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    /**
     * The end time of creating RDS instance.
     */
    createEndTime?: string;
    /**
     * The start time of creating RDS instance.
     */
    createStartTime?: string;
    /**
     * The id of the RDS instance.
     */
    instanceId?: string;
    /**
     * The status of the RDS instance.
     */
    instanceStatus?: string;
    /**
     * The type of the RDS instance.
     */
    instanceType?: string;
    /**
     * A Name Regex of RDS instance.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The region of the RDS instance.
     */
    region?: string;
    /**
     * The available zone of the RDS instance.
     */
    zone?: string;
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    readonly createEndTime?: string;
    readonly createStartTime?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The ID of the RDS instance.
     */
    readonly instanceId?: string;
    /**
     * The status of the RDS instance.
     */
    readonly instanceStatus?: string;
    /**
     * The type of the RDS instance.
     */
    readonly instanceType?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The collection of RDS instance query.
     */
    readonly rdsInstances: outputs.rds.GetInstancesRdsInstance[];
    /**
     * The region of the RDS instance.
     */
    readonly region?: string;
    /**
     * The total count of RDS instance query.
     */
    readonly totalCount: number;
    /**
     * The available zone of the RDS instance.
     */
    readonly zone?: string;
}
/**
 * (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds instances
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.rds.getInstances({
 *     instanceId: "mysql-0fdd3bab2e7c",
 * });
 * ```
 */
export function getInstancesOutput(args?: GetInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstancesResult> {
    return pulumi.output(args).apply((a: any) => getInstances(a, opts))
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesOutputArgs {
    /**
     * The end time of creating RDS instance.
     */
    createEndTime?: pulumi.Input<string>;
    /**
     * The start time of creating RDS instance.
     */
    createStartTime?: pulumi.Input<string>;
    /**
     * The id of the RDS instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The status of the RDS instance.
     */
    instanceStatus?: pulumi.Input<string>;
    /**
     * The type of the RDS instance.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * A Name Regex of RDS instance.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The region of the RDS instance.
     */
    region?: pulumi.Input<string>;
    /**
     * The available zone of the RDS instance.
     */
    zone?: pulumi.Input<string>;
}
