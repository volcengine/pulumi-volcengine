// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vedb mysql backups
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooZones = volcengine.ecs.getZones({});
 * const fooVpc = new volcengine.vpc.Vpc("fooVpc", {
 *     vpcName: "acc-test-vpc",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const fooSubnet = new volcengine.vpc.Subnet("fooSubnet", {
 *     subnetName: "acc-test-subnet",
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[2]?.id),
 *     vpcId: fooVpc.id,
 * });
 * const fooInstance = new volcengine.vedb_mysql.Instance("fooInstance", {
 *     chargeType: "PostPaid",
 *     storageChargeType: "PostPaid",
 *     dbEngineVersion: "MySQL_8_0",
 *     dbMinorVersion: "3.0",
 *     nodeNumber: 2,
 *     nodeSpec: "vedb.mysql.x4.large",
 *     subnetId: fooSubnet.id,
 *     instanceName: "tf-test",
 *     projectName: "testA",
 *     tags: [
 *         {
 *             key: "tftest",
 *             value: "tftest",
 *         },
 *         {
 *             key: "tftest2",
 *             value: "tftest2",
 *         },
 *     ],
 * });
 * const fooBackup = new volcengine.vedb_mysql.Backup("fooBackup", {
 *     instanceId: fooInstance.id,
 *     backupPolicy: {
 *         backupTime: "18:00Z-20:00Z",
 *         fullBackupPeriod: "Monday,Tuesday,Wednesday",
 *         backupRetentionPeriod: 8,
 *     },
 * });
 * const fooBackups = volcengine.vedb_mysql.getBackupsOutput({
 *     instanceId: fooInstance.id,
 * });
 * ```
 */
/** @deprecated volcengine.vedb_mysql.Backups has been deprecated in favor of volcengine.vedb_mysql.getBackups */
export function backups(args: BackupsArgs, opts?: pulumi.InvokeOptions): Promise<BackupsResult> {
    pulumi.log.warn("backups is deprecated: volcengine.vedb_mysql.Backups has been deprecated in favor of volcengine.vedb_mysql.getBackups")

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:vedb_mysql/backups:Backups", {
        "backupEndTime": args.backupEndTime,
        "backupMethod": args.backupMethod,
        "backupStartTime": args.backupStartTime,
        "backupStatus": args.backupStatus,
        "backupType": args.backupType,
        "instanceId": args.instanceId,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking Backups.
 */
export interface BackupsArgs {
    /**
     * The end time of the backup.
     */
    backupEndTime?: string;
    /**
     * Backup method. Currently, only physical backup is supported. The value is Physical.
     */
    backupMethod?: string;
    /**
     * The start time of the backup.
     */
    backupStartTime?: string;
    /**
     * The status of the backup.
     */
    backupStatus?: string;
    /**
     * The type of the backup.
     */
    backupType?: string;
    /**
     * The id of the instance.
     */
    instanceId: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by Backups.
 */
export interface BackupsResult {
    /**
     * The end time of the backup.
     */
    readonly backupEndTime?: string;
    /**
     * The name of the backup method.
     */
    readonly backupMethod?: string;
    /**
     * The start time of the backup.
     */
    readonly backupStartTime?: string;
    /**
     * The status of the backup.
     */
    readonly backupStatus?: string;
    /**
     * The type of the backup.
     */
    readonly backupType?: string;
    /**
     * The collection of query.
     */
    readonly backups: outputs.vedb_mysql.BackupsBackup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The id of the instance.
     */
    readonly instanceId: string;
    readonly outputFile?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of vedb mysql backups
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooZones = volcengine.ecs.getZones({});
 * const fooVpc = new volcengine.vpc.Vpc("fooVpc", {
 *     vpcName: "acc-test-vpc",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const fooSubnet = new volcengine.vpc.Subnet("fooSubnet", {
 *     subnetName: "acc-test-subnet",
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[2]?.id),
 *     vpcId: fooVpc.id,
 * });
 * const fooInstance = new volcengine.vedb_mysql.Instance("fooInstance", {
 *     chargeType: "PostPaid",
 *     storageChargeType: "PostPaid",
 *     dbEngineVersion: "MySQL_8_0",
 *     dbMinorVersion: "3.0",
 *     nodeNumber: 2,
 *     nodeSpec: "vedb.mysql.x4.large",
 *     subnetId: fooSubnet.id,
 *     instanceName: "tf-test",
 *     projectName: "testA",
 *     tags: [
 *         {
 *             key: "tftest",
 *             value: "tftest",
 *         },
 *         {
 *             key: "tftest2",
 *             value: "tftest2",
 *         },
 *     ],
 * });
 * const fooBackup = new volcengine.vedb_mysql.Backup("fooBackup", {
 *     instanceId: fooInstance.id,
 *     backupPolicy: {
 *         backupTime: "18:00Z-20:00Z",
 *         fullBackupPeriod: "Monday,Tuesday,Wednesday",
 *         backupRetentionPeriod: 8,
 *     },
 * });
 * const fooBackups = volcengine.vedb_mysql.getBackupsOutput({
 *     instanceId: fooInstance.id,
 * });
 * ```
 */
/** @deprecated volcengine.vedb_mysql.Backups has been deprecated in favor of volcengine.vedb_mysql.getBackups */
export function backupsOutput(args: BackupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<BackupsResult> {
    return pulumi.output(args).apply((a: any) => backups(a, opts))
}

/**
 * A collection of arguments for invoking Backups.
 */
export interface BackupsOutputArgs {
    /**
     * The end time of the backup.
     */
    backupEndTime?: pulumi.Input<string>;
    /**
     * Backup method. Currently, only physical backup is supported. The value is Physical.
     */
    backupMethod?: pulumi.Input<string>;
    /**
     * The start time of the backup.
     */
    backupStartTime?: pulumi.Input<string>;
    /**
     * The status of the backup.
     */
    backupStatus?: pulumi.Input<string>;
    /**
     * The type of the backup.
     */
    backupType?: pulumi.Input<string>;
    /**
     * The id of the instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
