// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tls shards
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.tls.getShards({
 *     topicId: "edf051ed-3c46-49ba-9339-bea628fedc15",
 * });
 * ```
 */
/** @deprecated volcengine.tls.Shards has been deprecated in favor of volcengine.tls.getShards */
export function shards(args: ShardsArgs, opts?: pulumi.InvokeOptions): Promise<ShardsResult> {
    pulumi.log.warn("shards is deprecated: volcengine.tls.Shards has been deprecated in favor of volcengine.tls.getShards")

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:tls/shards:Shards", {
        "outputFile": args.outputFile,
        "topicId": args.topicId,
    }, opts);
}

/**
 * A collection of arguments for invoking Shards.
 */
export interface ShardsArgs {
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The id of topic.
     */
    topicId: string;
}

/**
 * A collection of values returned by Shards.
 */
export interface ShardsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
    /**
     * The collection of query.
     */
    readonly shards: outputs.tls.ShardsShard[];
    /**
     * The ID of topic.
     */
    readonly topicId: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of tls shards
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.tls.getShards({
 *     topicId: "edf051ed-3c46-49ba-9339-bea628fedc15",
 * });
 * ```
 */
/** @deprecated volcengine.tls.Shards has been deprecated in favor of volcengine.tls.getShards */
export function shardsOutput(args: ShardsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ShardsResult> {
    return pulumi.output(args).apply((a: any) => shards(a, opts))
}

/**
 * A collection of arguments for invoking Shards.
 */
export interface ShardsOutputArgs {
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The id of topic.
     */
    topicId: pulumi.Input<string>;
}
