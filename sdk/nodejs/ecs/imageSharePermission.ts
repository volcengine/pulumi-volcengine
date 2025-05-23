// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage image share permission
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooImage = new volcengine.ecs.Image("fooImage", {
 *     imageName: "acc-test-image",
 *     description: "acc-test",
 *     instanceId: "i-ydi2q1s7wgqc6ild****",
 *     createWholeImage: false,
 *     projectName: "default",
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * const fooImageSharePermission = new volcengine.ecs.ImageSharePermission("fooImageSharePermission", {
 *     imageId: fooImage.id,
 *     accountId: "21000*****",
 * });
 * ```
 *
 * ## Import
 *
 * ImageSharePermission can be imported using the image_id:account_id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:ecs/imageSharePermission:ImageSharePermission default resource_id
 * ```
 */
export class ImageSharePermission extends pulumi.CustomResource {
    /**
     * Get an existing ImageSharePermission resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ImageSharePermissionState, opts?: pulumi.CustomResourceOptions): ImageSharePermission {
        return new ImageSharePermission(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:ecs/imageSharePermission:ImageSharePermission';

    /**
     * Returns true if the given object is an instance of ImageSharePermission.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ImageSharePermission {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ImageSharePermission.__pulumiType;
    }

    /**
     * The share account id of the image.
     */
    public readonly accountId!: pulumi.Output<string>;
    /**
     * The id of the image.
     */
    public readonly imageId!: pulumi.Output<string>;

    /**
     * Create a ImageSharePermission resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImageSharePermissionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ImageSharePermissionArgs | ImageSharePermissionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ImageSharePermissionState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
        } else {
            const args = argsOrState as ImageSharePermissionArgs | undefined;
            if ((!args || args.accountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountId'");
            }
            if ((!args || args.imageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageId'");
            }
            resourceInputs["accountId"] = args ? args.accountId : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ImageSharePermission.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ImageSharePermission resources.
 */
export interface ImageSharePermissionState {
    /**
     * The share account id of the image.
     */
    accountId?: pulumi.Input<string>;
    /**
     * The id of the image.
     */
    imageId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ImageSharePermission resource.
 */
export interface ImageSharePermissionArgs {
    /**
     * The share account id of the image.
     */
    accountId: pulumi.Input<string>;
    /**
     * The id of the image.
     */
    imageId: pulumi.Input<string>;
}
