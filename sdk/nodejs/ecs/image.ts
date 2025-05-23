// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage image
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.ecs.Image("foo", {
 *     createWholeImage: false,
 *     description: "acc-test",
 *     imageName: "acc-test-image",
 *     instanceId: "i-ydi2q1s7wgqc6ild****",
 *     projectName: "default",
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Image can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:ecs/image:Image default resource_id
 * ```
 */
export class Image extends pulumi.CustomResource {
    /**
     * Get an existing Image resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ImageState, opts?: pulumi.CustomResourceOptions): Image {
        return new Image(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:ecs/image:Image';

    /**
     * Returns true if the given object is an instance of Image.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Image {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Image.__pulumiType;
    }

    /**
     * The architecture of Image.
     */
    public /*out*/ readonly architecture!: pulumi.Output<string>;
    /**
     * The boot mode of the custom image. Valid values: `BIOS`, `UEFI`. This field is only effective when modifying the image.
     */
    public readonly bootMode!: pulumi.Output<string>;
    /**
     * Whether to create whole image. Default is false. This field is only effective when creating a new custom image.
     */
    public readonly createWholeImage!: pulumi.Output<boolean | undefined>;
    /**
     * The create time of Image.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The description of the custom image.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the custom image.
     */
    public readonly imageName!: pulumi.Output<string>;
    /**
     * The instance id of the custom image. Only one of `instance_id, snapshot_id, snapshotGroupId` can be specified.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly instanceId!: pulumi.Output<string | undefined>;
    /**
     * Whether the Image support cloud-init.
     */
    public /*out*/ readonly isSupportCloudInit!: pulumi.Output<boolean>;
    /**
     * The name of Image operating system.
     */
    public /*out*/ readonly osName!: pulumi.Output<string>;
    /**
     * The operating system type of Image.
     */
    public /*out*/ readonly osType!: pulumi.Output<string>;
    /**
     * The platform of Image.
     */
    public /*out*/ readonly platform!: pulumi.Output<string>;
    /**
     * The platform version of Image.
     */
    public /*out*/ readonly platformVersion!: pulumi.Output<string>;
    /**
     * The project name of the custom image.
     */
    public readonly projectName!: pulumi.Output<string>;
    /**
     * The share mode of Image.
     */
    public /*out*/ readonly shareStatus!: pulumi.Output<string>;
    /**
     * The size(GiB) of Image.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * The snapshot group id of the custom image. Only one of `instance_id, snapshot_id, snapshotGroupId` can be specified.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly snapshotGroupId!: pulumi.Output<string | undefined>;
    /**
     * The snapshot id of the custom image. Only one of `instance_id, snapshot_id, snapshotGroupId` can be specified.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly snapshotId!: pulumi.Output<string | undefined>;
    /**
     * The status of Image.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Tags.
     */
    public readonly tags!: pulumi.Output<outputs.ecs.ImageTag[] | undefined>;
    /**
     * The update time of Image.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * The visibility of Image.
     */
    public /*out*/ readonly visibility!: pulumi.Output<string>;

    /**
     * Create a Image resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ImageArgs | ImageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ImageState | undefined;
            resourceInputs["architecture"] = state ? state.architecture : undefined;
            resourceInputs["bootMode"] = state ? state.bootMode : undefined;
            resourceInputs["createWholeImage"] = state ? state.createWholeImage : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["imageName"] = state ? state.imageName : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["isSupportCloudInit"] = state ? state.isSupportCloudInit : undefined;
            resourceInputs["osName"] = state ? state.osName : undefined;
            resourceInputs["osType"] = state ? state.osType : undefined;
            resourceInputs["platform"] = state ? state.platform : undefined;
            resourceInputs["platformVersion"] = state ? state.platformVersion : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["shareStatus"] = state ? state.shareStatus : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["snapshotGroupId"] = state ? state.snapshotGroupId : undefined;
            resourceInputs["snapshotId"] = state ? state.snapshotId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
        } else {
            const args = argsOrState as ImageArgs | undefined;
            if ((!args || args.imageName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageName'");
            }
            resourceInputs["bootMode"] = args ? args.bootMode : undefined;
            resourceInputs["createWholeImage"] = args ? args.createWholeImage : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["imageName"] = args ? args.imageName : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["snapshotGroupId"] = args ? args.snapshotGroupId : undefined;
            resourceInputs["snapshotId"] = args ? args.snapshotId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["architecture"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["isSupportCloudInit"] = undefined /*out*/;
            resourceInputs["osName"] = undefined /*out*/;
            resourceInputs["osType"] = undefined /*out*/;
            resourceInputs["platform"] = undefined /*out*/;
            resourceInputs["platformVersion"] = undefined /*out*/;
            resourceInputs["shareStatus"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["visibility"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Image.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Image resources.
 */
export interface ImageState {
    /**
     * The architecture of Image.
     */
    architecture?: pulumi.Input<string>;
    /**
     * The boot mode of the custom image. Valid values: `BIOS`, `UEFI`. This field is only effective when modifying the image.
     */
    bootMode?: pulumi.Input<string>;
    /**
     * Whether to create whole image. Default is false. This field is only effective when creating a new custom image.
     */
    createWholeImage?: pulumi.Input<boolean>;
    /**
     * The create time of Image.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The description of the custom image.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the custom image.
     */
    imageName?: pulumi.Input<string>;
    /**
     * The instance id of the custom image. Only one of `instance_id, snapshot_id, snapshotGroupId` can be specified.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Whether the Image support cloud-init.
     */
    isSupportCloudInit?: pulumi.Input<boolean>;
    /**
     * The name of Image operating system.
     */
    osName?: pulumi.Input<string>;
    /**
     * The operating system type of Image.
     */
    osType?: pulumi.Input<string>;
    /**
     * The platform of Image.
     */
    platform?: pulumi.Input<string>;
    /**
     * The platform version of Image.
     */
    platformVersion?: pulumi.Input<string>;
    /**
     * The project name of the custom image.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The share mode of Image.
     */
    shareStatus?: pulumi.Input<string>;
    /**
     * The size(GiB) of Image.
     */
    size?: pulumi.Input<number>;
    /**
     * The snapshot group id of the custom image. Only one of `instance_id, snapshot_id, snapshotGroupId` can be specified.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    snapshotGroupId?: pulumi.Input<string>;
    /**
     * The snapshot id of the custom image. Only one of `instance_id, snapshot_id, snapshotGroupId` can be specified.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * The status of Image.
     */
    status?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.ecs.ImageTag>[]>;
    /**
     * The update time of Image.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * The visibility of Image.
     */
    visibility?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Image resource.
 */
export interface ImageArgs {
    /**
     * The boot mode of the custom image. Valid values: `BIOS`, `UEFI`. This field is only effective when modifying the image.
     */
    bootMode?: pulumi.Input<string>;
    /**
     * Whether to create whole image. Default is false. This field is only effective when creating a new custom image.
     */
    createWholeImage?: pulumi.Input<boolean>;
    /**
     * The description of the custom image.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the custom image.
     */
    imageName: pulumi.Input<string>;
    /**
     * The instance id of the custom image. Only one of `instance_id, snapshot_id, snapshotGroupId` can be specified.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The project name of the custom image.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The snapshot group id of the custom image. Only one of `instance_id, snapshot_id, snapshotGroupId` can be specified.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    snapshotGroupId?: pulumi.Input<string>;
    /**
     * The snapshot id of the custom image. Only one of `instance_id, snapshot_id, snapshotGroupId` can be specified.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.ecs.ImageTag>[]>;
}
