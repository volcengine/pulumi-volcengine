// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of images
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.ecs.getImages({
 *     instanceTypeId: "ecs.g1.large",
 *     osType: "Linux",
 *     visibility: "public",
 * });
 * ```
 */
export function getImages(args?: GetImagesArgs, opts?: pulumi.InvokeOptions): Promise<GetImagesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:ecs/getImages:getImages", {
        "ids": args.ids,
        "imageName": args.imageName,
        "instanceTypeId": args.instanceTypeId,
        "isSupportCloudInit": args.isSupportCloudInit,
        "isTls": args.isTls,
        "nameRegex": args.nameRegex,
        "osType": args.osType,
        "outputFile": args.outputFile,
        "platform": args.platform,
        "statuses": args.statuses,
        "tags": args.tags,
        "visibility": args.visibility,
    }, opts);
}

/**
 * A collection of arguments for invoking getImages.
 */
export interface GetImagesArgs {
    /**
     * A list of Image IDs.
     */
    ids?: string[];
    /**
     * The name of Image.
     */
    imageName?: string;
    /**
     * The specification of  Instance.
     */
    instanceTypeId?: string;
    /**
     * Whether the Image support cloud-init.
     */
    isSupportCloudInit?: boolean;
    /**
     * Whether the Image maintained for a long time.
     */
    isTls?: boolean;
    /**
     * A Name Regex of Image.
     */
    nameRegex?: string;
    /**
     * The operating system type of Image.
     */
    osType?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The platform of Image.
     */
    platform?: string;
    /**
     * A list of Image status, the value can be `available` or `creating` or `error`.
     */
    statuses?: string[];
    /**
     * Tags.
     */
    tags?: inputs.ecs.GetImagesTag[];
    /**
     * The visibility of Image.
     */
    visibility?: string;
}

/**
 * A collection of values returned by getImages.
 */
export interface GetImagesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The name of Image.
     */
    readonly imageName?: string;
    /**
     * The collection of Image query.
     */
    readonly images: outputs.ecs.GetImagesImage[];
    readonly instanceTypeId?: string;
    /**
     * Whether the Image support cloud-init.
     */
    readonly isSupportCloudInit?: boolean;
    readonly isTls?: boolean;
    readonly nameRegex?: string;
    /**
     * The operating system type of Image.
     */
    readonly osType?: string;
    readonly outputFile?: string;
    /**
     * The platform of Image.
     */
    readonly platform?: string;
    /**
     * The status of Image.
     */
    readonly statuses?: string[];
    /**
     * Tags.
     */
    readonly tags?: outputs.ecs.GetImagesTag[];
    /**
     * The total count of Image query.
     */
    readonly totalCount: number;
    /**
     * The visibility of Image.
     */
    readonly visibility?: string;
}
/**
 * Use this data source to query detailed information of images
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.ecs.getImages({
 *     instanceTypeId: "ecs.g1.large",
 *     osType: "Linux",
 *     visibility: "public",
 * });
 * ```
 */
export function getImagesOutput(args?: GetImagesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetImagesResult> {
    return pulumi.output(args).apply((a: any) => getImages(a, opts))
}

/**
 * A collection of arguments for invoking getImages.
 */
export interface GetImagesOutputArgs {
    /**
     * A list of Image IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of Image.
     */
    imageName?: pulumi.Input<string>;
    /**
     * The specification of  Instance.
     */
    instanceTypeId?: pulumi.Input<string>;
    /**
     * Whether the Image support cloud-init.
     */
    isSupportCloudInit?: pulumi.Input<boolean>;
    /**
     * Whether the Image maintained for a long time.
     */
    isTls?: pulumi.Input<boolean>;
    /**
     * A Name Regex of Image.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * The operating system type of Image.
     */
    osType?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The platform of Image.
     */
    platform?: pulumi.Input<string>;
    /**
     * A list of Image status, the value can be `available` or `creating` or `error`.
     */
    statuses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.ecs.GetImagesTagArgs>[]>;
    /**
     * The visibility of Image.
     */
    visibility?: pulumi.Input<string>;
}
