// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of nas snapshots
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.nas.getSnapshots({
 *     fileSystemId: "enas-cnbj5c18f02afe0e",
 *     ids: [
 *         "snap-022c648fed8b",
 *         "snap-e53591b05fbd",
 *     ],
 * });
 * ```
 */
/** @deprecated volcengine.nas.Snapshots has been deprecated in favor of volcengine.nas.getSnapshots */
export function snapshots(args?: SnapshotsArgs, opts?: pulumi.InvokeOptions): Promise<SnapshotsResult> {
    pulumi.log.warn("snapshots is deprecated: volcengine.nas.Snapshots has been deprecated in favor of volcengine.nas.getSnapshots")
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:nas/snapshots:Snapshots", {
        "fileSystemId": args.fileSystemId,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "snapshotName": args.snapshotName,
        "snapshotType": args.snapshotType,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking Snapshots.
 */
export interface SnapshotsArgs {
    /**
     * The ID of file system.
     */
    fileSystemId?: string;
    /**
     * A list of Snapshot IDs.
     */
    ids?: string[];
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The name of snapshot.
     */
    snapshotName?: string;
    /**
     * The type of snapshot.
     */
    snapshotType?: string;
    /**
     * The status of snapshot.
     */
    status?: string;
}

/**
 * A collection of values returned by Snapshots.
 */
export interface SnapshotsResult {
    /**
     * The id of file system.
     */
    readonly fileSystemId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly outputFile?: string;
    /**
     * The name of snapshot.
     */
    readonly snapshotName?: string;
    /**
     * The type of snapshot.
     */
    readonly snapshotType?: string;
    /**
     * The collection of query.
     */
    readonly snapshots: outputs.nas.SnapshotsSnapshot[];
    /**
     * The status of snapshot.
     */
    readonly status?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of nas snapshots
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.nas.getSnapshots({
 *     fileSystemId: "enas-cnbj5c18f02afe0e",
 *     ids: [
 *         "snap-022c648fed8b",
 *         "snap-e53591b05fbd",
 *     ],
 * });
 * ```
 */
/** @deprecated volcengine.nas.Snapshots has been deprecated in favor of volcengine.nas.getSnapshots */
export function snapshotsOutput(args?: SnapshotsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<SnapshotsResult> {
    return pulumi.output(args).apply((a: any) => snapshots(a, opts))
}

/**
 * A collection of arguments for invoking Snapshots.
 */
export interface SnapshotsOutputArgs {
    /**
     * The ID of file system.
     */
    fileSystemId?: pulumi.Input<string>;
    /**
     * A list of Snapshot IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The name of snapshot.
     */
    snapshotName?: pulumi.Input<string>;
    /**
     * The type of snapshot.
     */
    snapshotType?: pulumi.Input<string>;
    /**
     * The status of snapshot.
     */
    status?: pulumi.Input<string>;
}
