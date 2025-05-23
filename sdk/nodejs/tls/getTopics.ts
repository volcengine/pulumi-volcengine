// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tls topics
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.tls.getTopics({
 *     projectId: "e020c978-4f05-40e1-9167-0113d3ef****",
 *     topicId: "edf051ed-3c46-49ba-9339-bea628fe****",
 * });
 * ```
 */
export function getTopics(args: GetTopicsArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:tls/getTopics:getTopics", {
        "isFullName": args.isFullName,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "projectId": args.projectId,
        "tags": args.tags,
        "topicId": args.topicId,
        "topicName": args.topicName,
    }, opts);
}

/**
 * A collection of arguments for invoking getTopics.
 */
export interface GetTopicsArgs {
    /**
     * Whether to match accurately when filtering based on TopicName.
     */
    isFullName?: boolean;
    /**
     * A Name Regex of tls topic.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The project id of tls topic.
     */
    projectId: string;
    /**
     * Tags.
     */
    tags?: inputs.tls.GetTopicsTag[];
    /**
     * The id of tls topic. This field supports fuzzy queries. It is not supported to specify both TopicName and TopicId at the same time.
     */
    topicId?: string;
    /**
     * The name of tls topic. This field supports fuzzy queries. It is not supported to specify both TopicName and TopicId at the same time.
     */
    topicName?: string;
}

/**
 * A collection of values returned by getTopics.
 */
export interface GetTopicsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isFullName?: boolean;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The project id of the tls topic.
     */
    readonly projectId: string;
    /**
     * Tags.
     */
    readonly tags?: outputs.tls.GetTopicsTag[];
    /**
     * The collection of tls topic query.
     */
    readonly tlsTopics: outputs.tls.GetTopicsTlsTopic[];
    /**
     * The ID of the tls topic.
     */
    readonly topicId?: string;
    /**
     * The name of the tls topic.
     */
    readonly topicName?: string;
    /**
     * The total count of tls topic query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of tls topics
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.tls.getTopics({
 *     projectId: "e020c978-4f05-40e1-9167-0113d3ef****",
 *     topicId: "edf051ed-3c46-49ba-9339-bea628fe****",
 * });
 * ```
 */
export function getTopicsOutput(args: GetTopicsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTopicsResult> {
    return pulumi.output(args).apply((a: any) => getTopics(a, opts))
}

/**
 * A collection of arguments for invoking getTopics.
 */
export interface GetTopicsOutputArgs {
    /**
     * Whether to match accurately when filtering based on TopicName.
     */
    isFullName?: pulumi.Input<boolean>;
    /**
     * A Name Regex of tls topic.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The project id of tls topic.
     */
    projectId: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.tls.GetTopicsTagArgs>[]>;
    /**
     * The id of tls topic. This field supports fuzzy queries. It is not supported to specify both TopicName and TopicId at the same time.
     */
    topicId?: pulumi.Input<string>;
    /**
     * The name of tls topic. This field supports fuzzy queries. It is not supported to specify both TopicName and TopicId at the same time.
     */
    topicName?: pulumi.Input<string>;
}
