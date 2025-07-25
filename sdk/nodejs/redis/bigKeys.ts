// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of redis big keys
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
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *     vpcId: fooVpc.id,
 * });
 * const fooInstance = new volcengine.redis.Instance("fooInstance", {
 *     zoneIds: [fooZones.then(fooZones => fooZones.zones?.[0]?.id)],
 *     instanceName: "acc-test-tf-redis",
 *     shardedCluster: 1,
 *     password: "1qaz!QAZ12",
 *     nodeNumber: 2,
 *     shardCapacity: 1024,
 *     shardNumber: 2,
 *     engineVersion: "5.0",
 *     subnetId: fooSubnet.id,
 *     deletionProtection: "disabled",
 *     vpcAuthMode: "close",
 *     chargeType: "PostPaid",
 *     port: 6381,
 *     projectName: "default",
 * });
 * const fooBigKeys = volcengine.redis.getBigKeysOutput({
 *     instanceId: fooInstance.id,
 * });
 * ```
 */
/** @deprecated volcengine.redis.BigKeys has been deprecated in favor of volcengine.redis.getBigKeys */
export function bigKeys(args: BigKeysArgs, opts?: pulumi.InvokeOptions): Promise<BigKeysResult> {
    pulumi.log.warn("bigKeys is deprecated: volcengine.redis.BigKeys has been deprecated in favor of volcengine.redis.getBigKeys")

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:redis/bigKeys:BigKeys", {
        "instanceId": args.instanceId,
        "keyType": args.keyType,
        "nameRegex": args.nameRegex,
        "orderBy": args.orderBy,
        "outputFile": args.outputFile,
        "queryEndTime": args.queryEndTime,
        "queryStartTime": args.queryStartTime,
    }, opts);
}

/**
 * A collection of arguments for invoking BigKeys.
 */
export interface BigKeysArgs {
    /**
     * The ID of Instance.
     */
    instanceId: string;
    /**
     * Specify the data type used to filter the query results of hot keys.
     */
    keyType?: string;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: string;
    /**
     * Specify the sorting conditions of the query results.
     */
    orderBy?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * Query the end time in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).
     */
    queryEndTime?: string;
    /**
     * Query the start time in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).
     */
    queryStartTime?: string;
}

/**
 * A collection of values returned by BigKeys.
 */
export interface BigKeysResult {
    /**
     * Details of the big Key.
     */
    readonly bigKeys: outputs.redis.BigKeysBigKey[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    /**
     * The type of big Key.
     */
    readonly keyType?: string;
    readonly nameRegex?: string;
    readonly orderBy?: string;
    readonly outputFile?: string;
    readonly queryEndTime?: string;
    readonly queryStartTime?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of redis big keys
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
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *     vpcId: fooVpc.id,
 * });
 * const fooInstance = new volcengine.redis.Instance("fooInstance", {
 *     zoneIds: [fooZones.then(fooZones => fooZones.zones?.[0]?.id)],
 *     instanceName: "acc-test-tf-redis",
 *     shardedCluster: 1,
 *     password: "1qaz!QAZ12",
 *     nodeNumber: 2,
 *     shardCapacity: 1024,
 *     shardNumber: 2,
 *     engineVersion: "5.0",
 *     subnetId: fooSubnet.id,
 *     deletionProtection: "disabled",
 *     vpcAuthMode: "close",
 *     chargeType: "PostPaid",
 *     port: 6381,
 *     projectName: "default",
 * });
 * const fooBigKeys = volcengine.redis.getBigKeysOutput({
 *     instanceId: fooInstance.id,
 * });
 * ```
 */
/** @deprecated volcengine.redis.BigKeys has been deprecated in favor of volcengine.redis.getBigKeys */
export function bigKeysOutput(args: BigKeysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<BigKeysResult> {
    return pulumi.output(args).apply((a: any) => bigKeys(a, opts))
}

/**
 * A collection of arguments for invoking BigKeys.
 */
export interface BigKeysOutputArgs {
    /**
     * The ID of Instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Specify the data type used to filter the query results of hot keys.
     */
    keyType?: pulumi.Input<string>;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * Specify the sorting conditions of the query results.
     */
    orderBy?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * Query the end time in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).
     */
    queryEndTime?: pulumi.Input<string>;
    /**
     * Query the start time in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).
     */
    queryStartTime?: pulumi.Input<string>;
}
