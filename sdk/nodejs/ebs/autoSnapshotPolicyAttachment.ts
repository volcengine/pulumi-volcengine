// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage ebs auto snapshot policy attachment
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooZones = volcengine.ecs.Zones({});
 * const fooVolume = new volcengine.ebs.Volume("fooVolume", {
 *     volumeName: "acc-test-volume",
 *     volumeType: "ESSD_PL0",
 *     description: "acc-test",
 *     kind: "data",
 *     size: 500,
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *     volumeChargeType: "PostPaid",
 *     projectName: "default",
 * });
 * const fooAutoSnapshotPolicy = new volcengine.ebs.AutoSnapshotPolicy("fooAutoSnapshotPolicy", {
 *     autoSnapshotPolicyName: "acc-test-auto-snapshot-policy",
 *     timePoints: [
 *         "1",
 *         "5",
 *         "9",
 *     ],
 *     retentionDays: -1,
 *     repeatWeekdays: [
 *         "2",
 *         "6",
 *     ],
 *     projectName: "default",
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * const fooAutoSnapshotPolicyAttachment = new volcengine.ebs.AutoSnapshotPolicyAttachment("fooAutoSnapshotPolicyAttachment", {
 *     autoSnapshotPolicyId: fooAutoSnapshotPolicy.id,
 *     volumeId: fooVolume.id,
 * });
 * ```
 *
 * ## Import
 *
 * EbsAutoSnapshotPolicyAttachment can be imported using the auto_snapshot_policy_id:volume_id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:ebs/autoSnapshotPolicyAttachment:AutoSnapshotPolicyAttachment default resource_id
 * ```
 */
export class AutoSnapshotPolicyAttachment extends pulumi.CustomResource {
    /**
     * Get an existing AutoSnapshotPolicyAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutoSnapshotPolicyAttachmentState, opts?: pulumi.CustomResourceOptions): AutoSnapshotPolicyAttachment {
        return new AutoSnapshotPolicyAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:ebs/autoSnapshotPolicyAttachment:AutoSnapshotPolicyAttachment';

    /**
     * Returns true if the given object is an instance of AutoSnapshotPolicyAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AutoSnapshotPolicyAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AutoSnapshotPolicyAttachment.__pulumiType;
    }

    /**
     * The id of the auto snapshot policy.
     */
    public readonly autoSnapshotPolicyId!: pulumi.Output<string>;
    /**
     * The id of the volume.
     */
    public readonly volumeId!: pulumi.Output<string>;

    /**
     * Create a AutoSnapshotPolicyAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AutoSnapshotPolicyAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AutoSnapshotPolicyAttachmentArgs | AutoSnapshotPolicyAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AutoSnapshotPolicyAttachmentState | undefined;
            resourceInputs["autoSnapshotPolicyId"] = state ? state.autoSnapshotPolicyId : undefined;
            resourceInputs["volumeId"] = state ? state.volumeId : undefined;
        } else {
            const args = argsOrState as AutoSnapshotPolicyAttachmentArgs | undefined;
            if ((!args || args.autoSnapshotPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoSnapshotPolicyId'");
            }
            if ((!args || args.volumeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'volumeId'");
            }
            resourceInputs["autoSnapshotPolicyId"] = args ? args.autoSnapshotPolicyId : undefined;
            resourceInputs["volumeId"] = args ? args.volumeId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AutoSnapshotPolicyAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AutoSnapshotPolicyAttachment resources.
 */
export interface AutoSnapshotPolicyAttachmentState {
    /**
     * The id of the auto snapshot policy.
     */
    autoSnapshotPolicyId?: pulumi.Input<string>;
    /**
     * The id of the volume.
     */
    volumeId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AutoSnapshotPolicyAttachment resource.
 */
export interface AutoSnapshotPolicyAttachmentArgs {
    /**
     * The id of the auto snapshot policy.
     */
    autoSnapshotPolicyId: pulumi.Input<string>;
    /**
     * The id of the volume.
     */
    volumeId: pulumi.Input<string>;
}