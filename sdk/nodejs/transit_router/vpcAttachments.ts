// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of transit router vpc attachments
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.transit_router.VpcAttachments({
 *     transitRouterAttachmentIds: ["tr-attach-3rf2xi7ae6y9s5zsk2hm6pibt"],
 *     transitRouterId: "tr-2d6fr7f39unsw58ozfe1ow21x",
 * });
 * ```
 */
export function vpcAttachments(args: VpcAttachmentsArgs, opts?: pulumi.InvokeOptions): Promise<VpcAttachmentsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:transit_router/vpcAttachments:VpcAttachments", {
        "outputFile": args.outputFile,
        "transitRouterAttachmentIds": args.transitRouterAttachmentIds,
        "transitRouterId": args.transitRouterId,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking VpcAttachments.
 */
export interface VpcAttachmentsArgs {
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * A list of Transit Router Attachment ids.
     */
    transitRouterAttachmentIds?: string[];
    /**
     * The id of transit router.
     */
    transitRouterId: string;
    /**
     * The id of vpc.
     */
    vpcId?: string;
}

/**
 * A collection of values returned by VpcAttachments.
 */
export interface VpcAttachmentsResult {
    /**
     * The collection of query.
     */
    readonly attachments: outputs.transit_router.VpcAttachmentsAttachment[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
    readonly transitRouterAttachmentIds?: string[];
    /**
     * The id of the transit router.
     */
    readonly transitRouterId: string;
    /**
     * The ID of vpc.
     */
    readonly vpcId?: string;
}
/**
 * Use this data source to query detailed information of transit router vpc attachments
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.transit_router.VpcAttachments({
 *     transitRouterAttachmentIds: ["tr-attach-3rf2xi7ae6y9s5zsk2hm6pibt"],
 *     transitRouterId: "tr-2d6fr7f39unsw58ozfe1ow21x",
 * });
 * ```
 */
export function vpcAttachmentsOutput(args: VpcAttachmentsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<VpcAttachmentsResult> {
    return pulumi.output(args).apply((a: any) => vpcAttachments(a, opts))
}

/**
 * A collection of arguments for invoking VpcAttachments.
 */
export interface VpcAttachmentsOutputArgs {
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * A list of Transit Router Attachment ids.
     */
    transitRouterAttachmentIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The id of transit router.
     */
    transitRouterId: pulumi.Input<string>;
    /**
     * The id of vpc.
     */
    vpcId?: pulumi.Input<string>;
}