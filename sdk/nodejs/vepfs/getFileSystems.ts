// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vepfs file systems
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooVpc = new volcengine.vpc.Vpc("fooVpc", {
 *     vpcName: "acc-test-vpc",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const fooSubnet = new volcengine.vpc.Subnet("fooSubnet", {
 *     subnetName: "acc-test-subnet",
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: "cn-beijing-a",
 *     vpcId: fooVpc.id,
 * });
 * const fooFileSystem = new volcengine.vepfs.FileSystem("fooFileSystem", {
 *     fileSystemName: "acc-test-file-system",
 *     subnetId: fooSubnet.id,
 *     storeType: "Advance_100",
 *     description: "tf-test",
 *     capacity: 12,
 *     project: "default",
 *     enableRestripe: false,
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * const fooFileSystems = volcengine.vepfs.getFileSystemsOutput({
 *     ids: [fooFileSystem.id],
 * });
 * ```
 */
export function getFileSystems(args?: GetFileSystemsArgs, opts?: pulumi.InvokeOptions): Promise<GetFileSystemsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:vepfs/getFileSystems:getFileSystems", {
        "fileSystemName": args.fileSystemName,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "project": args.project,
        "statuses": args.statuses,
        "storeType": args.storeType,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getFileSystems.
 */
export interface GetFileSystemsArgs {
    /**
     * The Name of Vepfs File System. This field support fuzzy query.
     */
    fileSystemName?: string;
    /**
     * A list of Vepfs File System IDs.
     */
    ids?: string[];
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The project of Vepfs File System.
     */
    project?: string;
    /**
     * The query status list of Vepfs File System.
     */
    statuses?: string[];
    /**
     * The Store Type of Vepfs File System.
     */
    storeType?: string;
    /**
     * The zone id of File System.
     */
    zoneId?: string;
}

/**
 * A collection of values returned by getFileSystems.
 */
export interface GetFileSystemsResult {
    /**
     * The name of the vepfs file system.
     */
    readonly fileSystemName?: string;
    /**
     * The collection of query.
     */
    readonly fileSystems: outputs.vepfs.GetFileSystemsFileSystem[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The project name of the vepfs file system.
     */
    readonly project?: string;
    /**
     * The status of the vepfs file system.
     */
    readonly statuses?: string[];
    /**
     * The store type of the vepfs file system.
     */
    readonly storeType?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
    /**
     * The id of the zone.
     */
    readonly zoneId?: string;
}
/**
 * Use this data source to query detailed information of vepfs file systems
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooVpc = new volcengine.vpc.Vpc("fooVpc", {
 *     vpcName: "acc-test-vpc",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const fooSubnet = new volcengine.vpc.Subnet("fooSubnet", {
 *     subnetName: "acc-test-subnet",
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: "cn-beijing-a",
 *     vpcId: fooVpc.id,
 * });
 * const fooFileSystem = new volcengine.vepfs.FileSystem("fooFileSystem", {
 *     fileSystemName: "acc-test-file-system",
 *     subnetId: fooSubnet.id,
 *     storeType: "Advance_100",
 *     description: "tf-test",
 *     capacity: 12,
 *     project: "default",
 *     enableRestripe: false,
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * const fooFileSystems = volcengine.vepfs.getFileSystemsOutput({
 *     ids: [fooFileSystem.id],
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
     * The Name of Vepfs File System. This field support fuzzy query.
     */
    fileSystemName?: pulumi.Input<string>;
    /**
     * A list of Vepfs File System IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The project of Vepfs File System.
     */
    project?: pulumi.Input<string>;
    /**
     * The query status list of Vepfs File System.
     */
    statuses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Store Type of Vepfs File System.
     */
    storeType?: pulumi.Input<string>;
    /**
     * The zone id of File System.
     */
    zoneId?: pulumi.Input<string>;
}
