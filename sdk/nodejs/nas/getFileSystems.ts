// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of nas file systems
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooZones = volcengine.nas.getZones({});
 * const fooFileSystem: volcengine.nas.FileSystem[] = [];
 * for (const range = {value: 0}; range.value < 3; range.value++) {
 *     fooFileSystem.push(new volcengine.nas.FileSystem(`fooFileSystem-${range.value}`, {
 *         fileSystemName: `acc-test-fs-${range.value}`,
 *         description: "acc-test",
 *         zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *         capacity: 103,
 *         projectName: "default",
 *         tags: [{
 *             key: "k1",
 *             value: "v1",
 *         }],
 *     }));
 * }
 * const fooFileSystems = volcengine.nas.getFileSystemsOutput({
 *     ids: fooFileSystem.map(__item => __item.id),
 * });
 * ```
 */
export function getFileSystems(args?: GetFileSystemsArgs, opts?: pulumi.InvokeOptions): Promise<GetFileSystemsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:nas/getFileSystems:getFileSystems", {
        "chargeType": args.chargeType,
        "fileSystemName": args.fileSystemName,
        "ids": args.ids,
        "mountPointId": args.mountPointId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "permissionGroupId": args.permissionGroupId,
        "projectName": args.projectName,
        "protocolType": args.protocolType,
        "statuses": args.statuses,
        "storageType": args.storageType,
        "tags": args.tags,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getFileSystems.
 */
export interface GetFileSystemsArgs {
    /**
     * The charge type of nas file system.
     */
    chargeType?: string;
    /**
     * The name of nas file system. This field supports fuzzy queries.
     */
    fileSystemName?: string;
    /**
     * A list of nas file system ids.
     */
    ids?: string[];
    /**
     * The mount point id of nas file system.
     */
    mountPointId?: string;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The permission group id of nas file system.
     */
    permissionGroupId?: string;
    /**
     * The project name of nas file system.
     */
    projectName?: string;
    /**
     * The protocol type of nas file system.
     */
    protocolType?: string;
    /**
     * The status of nas file system.
     */
    statuses?: string[];
    /**
     * The storage type of nas file system.
     */
    storageType?: string;
    /**
     * Tags.
     */
    tags?: inputs.nas.GetFileSystemsTag[];
    /**
     * The zone id of nas file system.
     */
    zoneId?: string;
}

/**
 * A collection of values returned by getFileSystems.
 */
export interface GetFileSystemsResult {
    /**
     * The charge type of the nas file system.
     */
    readonly chargeType?: string;
    /**
     * The name of the nas file system.
     */
    readonly fileSystemName?: string;
    /**
     * The collection of query.
     */
    readonly fileSystems: outputs.nas.GetFileSystemsFileSystem[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly mountPointId?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    readonly permissionGroupId?: string;
    /**
     * The project name of the nas file system.
     */
    readonly projectName?: string;
    /**
     * The protocol type of the nas file system.
     */
    readonly protocolType?: string;
    /**
     * The status of the nas file system.
     */
    readonly statuses?: string[];
    /**
     * The storage type of the nas file system.
     */
    readonly storageType?: string;
    /**
     * Tags of the nas file system.
     */
    readonly tags?: outputs.nas.GetFileSystemsTag[];
    /**
     * The total count of query.
     */
    readonly totalCount: number;
    /**
     * The zone id of the nas file system.
     */
    readonly zoneId?: string;
}
/**
 * Use this data source to query detailed information of nas file systems
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooZones = volcengine.nas.getZones({});
 * const fooFileSystem: volcengine.nas.FileSystem[] = [];
 * for (const range = {value: 0}; range.value < 3; range.value++) {
 *     fooFileSystem.push(new volcengine.nas.FileSystem(`fooFileSystem-${range.value}`, {
 *         fileSystemName: `acc-test-fs-${range.value}`,
 *         description: "acc-test",
 *         zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *         capacity: 103,
 *         projectName: "default",
 *         tags: [{
 *             key: "k1",
 *             value: "v1",
 *         }],
 *     }));
 * }
 * const fooFileSystems = volcengine.nas.getFileSystemsOutput({
 *     ids: fooFileSystem.map(__item => __item.id),
 * });
 * ```
 */
export function getFileSystemsOutput(args?: GetFileSystemsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFileSystemsResult> {
    return pulumi.output(args).apply((a: any) => getFileSystems(a, opts))
}

/**
 * A collection of arguments for invoking getFileSystems.
 */
export interface GetFileSystemsOutputArgs {
    /**
     * The charge type of nas file system.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * The name of nas file system. This field supports fuzzy queries.
     */
    fileSystemName?: pulumi.Input<string>;
    /**
     * A list of nas file system ids.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The mount point id of nas file system.
     */
    mountPointId?: pulumi.Input<string>;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The permission group id of nas file system.
     */
    permissionGroupId?: pulumi.Input<string>;
    /**
     * The project name of nas file system.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The protocol type of nas file system.
     */
    protocolType?: pulumi.Input<string>;
    /**
     * The status of nas file system.
     */
    statuses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The storage type of nas file system.
     */
    storageType?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.nas.GetFileSystemsTagArgs>[]>;
    /**
     * The zone id of nas file system.
     */
    zoneId?: pulumi.Input<string>;
}
