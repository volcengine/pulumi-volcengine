// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage nas auto snapshot policy apply
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooZones = volcengine.nas.getZones({});
 * const fooFileSystem = new volcengine.nas.FileSystem("fooFileSystem", {
 *     fileSystemName: "acc-test-fs",
 *     description: "acc-test",
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *     capacity: 103,
 *     projectName: "default",
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * const fooAutoSnapshotPolicy = new volcengine.nas.AutoSnapshotPolicy("fooAutoSnapshotPolicy", {
 *     autoSnapshotPolicyName: "acc-test-auto_snapshot_policy",
 *     repeatWeekdays: "1,3,5,7",
 *     timePoints: "0,7,17",
 *     retentionDays: 20,
 * });
 * const fooAutoSnapshotPolicyApply = new volcengine.nas.AutoSnapshotPolicyApply("fooAutoSnapshotPolicyApply", {
 *     fileSystemId: fooFileSystem.id,
 *     autoSnapshotPolicyId: fooAutoSnapshotPolicy.id,
 * });
 * ```
 *
 * ## Import
 *
 * NasAutoSnapshotPolicyApply can be imported using the auto_snapshot_policy_id:file_system_id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:nas/autoSnapshotPolicyApply:AutoSnapshotPolicyApply default resource_id
 * ```
 */
export class AutoSnapshotPolicyApply extends pulumi.CustomResource {
    /**
     * Get an existing AutoSnapshotPolicyApply resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutoSnapshotPolicyApplyState, opts?: pulumi.CustomResourceOptions): AutoSnapshotPolicyApply {
        return new AutoSnapshotPolicyApply(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:nas/autoSnapshotPolicyApply:AutoSnapshotPolicyApply';

    /**
     * Returns true if the given object is an instance of AutoSnapshotPolicyApply.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AutoSnapshotPolicyApply {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AutoSnapshotPolicyApply.__pulumiType;
    }

    /**
     * The id of auto snapshot policy.
     */
    public readonly autoSnapshotPolicyId!: pulumi.Output<string>;
    /**
     * The id of file system.
     */
    public readonly fileSystemId!: pulumi.Output<string>;

    /**
     * Create a AutoSnapshotPolicyApply resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AutoSnapshotPolicyApplyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AutoSnapshotPolicyApplyArgs | AutoSnapshotPolicyApplyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AutoSnapshotPolicyApplyState | undefined;
            resourceInputs["autoSnapshotPolicyId"] = state ? state.autoSnapshotPolicyId : undefined;
            resourceInputs["fileSystemId"] = state ? state.fileSystemId : undefined;
        } else {
            const args = argsOrState as AutoSnapshotPolicyApplyArgs | undefined;
            if ((!args || args.autoSnapshotPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoSnapshotPolicyId'");
            }
            if ((!args || args.fileSystemId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileSystemId'");
            }
            resourceInputs["autoSnapshotPolicyId"] = args ? args.autoSnapshotPolicyId : undefined;
            resourceInputs["fileSystemId"] = args ? args.fileSystemId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AutoSnapshotPolicyApply.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AutoSnapshotPolicyApply resources.
 */
export interface AutoSnapshotPolicyApplyState {
    /**
     * The id of auto snapshot policy.
     */
    autoSnapshotPolicyId?: pulumi.Input<string>;
    /**
     * The id of file system.
     */
    fileSystemId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AutoSnapshotPolicyApply resource.
 */
export interface AutoSnapshotPolicyApplyArgs {
    /**
     * The id of auto snapshot policy.
     */
    autoSnapshotPolicyId: pulumi.Input<string>;
    /**
     * The id of file system.
     */
    fileSystemId: pulumi.Input<string>;
}
