// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of nas permission groups
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.nas.PermissionGroup("foo", {
 *     permissionGroupName: "acc-test",
 *     description: "acctest",
 *     permissionRules: [
 *         {
 *             cidrIp: "*",
 *             rwMode: "RW",
 *             useMode: "All_squash",
 *         },
 *         {
 *             cidrIp: "192.168.0.0",
 *             rwMode: "RO",
 *             useMode: "All_squash",
 *         },
 *     ],
 * });
 * const default = volcengine.nas.getPermissionGroupsOutput({
 *     filters: [{
 *         key: "PermissionGroupId",
 *         value: foo.id,
 *     }],
 * });
 * ```
 */
export function getPermissionGroups(args?: GetPermissionGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetPermissionGroupsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:nas/getPermissionGroups:getPermissionGroups", {
        "filters": args.filters,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getPermissionGroups.
 */
export interface GetPermissionGroupsArgs {
    /**
     * Filter permission groups for specified characteristics.
     */
    filters?: inputs.nas.GetPermissionGroupsFilter[];
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getPermissionGroups.
 */
export interface GetPermissionGroupsResult {
    readonly filters?: outputs.nas.GetPermissionGroupsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
    /**
     * The list of permissions groups.
     */
    readonly permissionGroups: outputs.nas.GetPermissionGroupsPermissionGroup[];
    /**
     * The total count of nas permission groups query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of nas permission groups
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.nas.PermissionGroup("foo", {
 *     permissionGroupName: "acc-test",
 *     description: "acctest",
 *     permissionRules: [
 *         {
 *             cidrIp: "*",
 *             rwMode: "RW",
 *             useMode: "All_squash",
 *         },
 *         {
 *             cidrIp: "192.168.0.0",
 *             rwMode: "RO",
 *             useMode: "All_squash",
 *         },
 *     ],
 * });
 * const default = volcengine.nas.getPermissionGroupsOutput({
 *     filters: [{
 *         key: "PermissionGroupId",
 *         value: foo.id,
 *     }],
 * });
 * ```
 */
export function getPermissionGroupsOutput(args?: GetPermissionGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPermissionGroupsResult> {
    return pulumi.output(args).apply((a: any) => getPermissionGroups(a, opts))
}

/**
 * A collection of arguments for invoking getPermissionGroups.
 */
export interface GetPermissionGroupsOutputArgs {
    /**
     * Filter permission groups for specified characteristics.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.nas.GetPermissionGroupsFilterArgs>[]>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
