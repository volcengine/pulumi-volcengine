// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage tos bucket inventory
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.tos.BucketInventory("foo", {
 *     bucketName: "terraform-demo",
 *     destination: {
 *         tosBucketDestination: {
 *             accountId: "21000*****",
 *             bucket: "terraform-demo",
 *             format: "CSV",
 *             prefix: "tf-test-prefix",
 *             role: "TosArchiveTOSInventory",
 *         },
 *     },
 *     filter: {
 *         prefix: "test-tf",
 *     },
 *     includedObjectVersions: "All",
 *     inventoryId: "acc-test-inventory",
 *     isEnabled: true,
 *     optionalFields: {
 *         fields: [
 *             "Size",
 *             "StorageClass",
 *             "CRC64",
 *         ],
 *     },
 *     schedule: {
 *         frequency: "Weekly",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * TosBucketInventory can be imported using the bucket_name:inventory_id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:tos/bucketInventory:BucketInventory default resource_id
 * ```
 */
export class BucketInventory extends pulumi.CustomResource {
    /**
     * Get an existing BucketInventory resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketInventoryState, opts?: pulumi.CustomResourceOptions): BucketInventory {
        return new BucketInventory(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:tos/bucketInventory:BucketInventory';

    /**
     * Returns true if the given object is an instance of BucketInventory.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketInventory {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketInventory.__pulumiType;
    }

    /**
     * The name of the bucket.
     */
    public readonly bucketName!: pulumi.Output<string>;
    /**
     * The destination information of the bucket inventory.
     */
    public readonly destination!: pulumi.Output<outputs.tos.BucketInventoryDestination>;
    /**
     * The filter of the bucket inventory.
     */
    public readonly filter!: pulumi.Output<outputs.tos.BucketInventoryFilter | undefined>;
    /**
     * The export version of object. Valid values: `All`, `Current`.
     */
    public readonly includedObjectVersions!: pulumi.Output<string>;
    /**
     * The name of the bucket inventory.
     */
    public readonly inventoryId!: pulumi.Output<string>;
    /**
     * Whether to enable the bucket inventory.
     */
    public readonly isEnabled!: pulumi.Output<boolean>;
    /**
     * The information exported from the bucket inventory.
     */
    public readonly optionalFields!: pulumi.Output<outputs.tos.BucketInventoryOptionalFields | undefined>;
    /**
     * The export schedule of the bucket inventory.
     */
    public readonly schedule!: pulumi.Output<outputs.tos.BucketInventorySchedule>;

    /**
     * Create a BucketInventory resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketInventoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketInventoryArgs | BucketInventoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BucketInventoryState | undefined;
            resourceInputs["bucketName"] = state ? state.bucketName : undefined;
            resourceInputs["destination"] = state ? state.destination : undefined;
            resourceInputs["filter"] = state ? state.filter : undefined;
            resourceInputs["includedObjectVersions"] = state ? state.includedObjectVersions : undefined;
            resourceInputs["inventoryId"] = state ? state.inventoryId : undefined;
            resourceInputs["isEnabled"] = state ? state.isEnabled : undefined;
            resourceInputs["optionalFields"] = state ? state.optionalFields : undefined;
            resourceInputs["schedule"] = state ? state.schedule : undefined;
        } else {
            const args = argsOrState as BucketInventoryArgs | undefined;
            if ((!args || args.bucketName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucketName'");
            }
            if ((!args || args.destination === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destination'");
            }
            if ((!args || args.includedObjectVersions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'includedObjectVersions'");
            }
            if ((!args || args.inventoryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'inventoryId'");
            }
            if ((!args || args.isEnabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'isEnabled'");
            }
            if ((!args || args.schedule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schedule'");
            }
            resourceInputs["bucketName"] = args ? args.bucketName : undefined;
            resourceInputs["destination"] = args ? args.destination : undefined;
            resourceInputs["filter"] = args ? args.filter : undefined;
            resourceInputs["includedObjectVersions"] = args ? args.includedObjectVersions : undefined;
            resourceInputs["inventoryId"] = args ? args.inventoryId : undefined;
            resourceInputs["isEnabled"] = args ? args.isEnabled : undefined;
            resourceInputs["optionalFields"] = args ? args.optionalFields : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BucketInventory.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BucketInventory resources.
 */
export interface BucketInventoryState {
    /**
     * The name of the bucket.
     */
    bucketName?: pulumi.Input<string>;
    /**
     * The destination information of the bucket inventory.
     */
    destination?: pulumi.Input<inputs.tos.BucketInventoryDestination>;
    /**
     * The filter of the bucket inventory.
     */
    filter?: pulumi.Input<inputs.tos.BucketInventoryFilter>;
    /**
     * The export version of object. Valid values: `All`, `Current`.
     */
    includedObjectVersions?: pulumi.Input<string>;
    /**
     * The name of the bucket inventory.
     */
    inventoryId?: pulumi.Input<string>;
    /**
     * Whether to enable the bucket inventory.
     */
    isEnabled?: pulumi.Input<boolean>;
    /**
     * The information exported from the bucket inventory.
     */
    optionalFields?: pulumi.Input<inputs.tos.BucketInventoryOptionalFields>;
    /**
     * The export schedule of the bucket inventory.
     */
    schedule?: pulumi.Input<inputs.tos.BucketInventorySchedule>;
}

/**
 * The set of arguments for constructing a BucketInventory resource.
 */
export interface BucketInventoryArgs {
    /**
     * The name of the bucket.
     */
    bucketName: pulumi.Input<string>;
    /**
     * The destination information of the bucket inventory.
     */
    destination: pulumi.Input<inputs.tos.BucketInventoryDestination>;
    /**
     * The filter of the bucket inventory.
     */
    filter?: pulumi.Input<inputs.tos.BucketInventoryFilter>;
    /**
     * The export version of object. Valid values: `All`, `Current`.
     */
    includedObjectVersions: pulumi.Input<string>;
    /**
     * The name of the bucket inventory.
     */
    inventoryId: pulumi.Input<string>;
    /**
     * Whether to enable the bucket inventory.
     */
    isEnabled: pulumi.Input<boolean>;
    /**
     * The information exported from the bucket inventory.
     */
    optionalFields?: pulumi.Input<inputs.tos.BucketInventoryOptionalFields>;
    /**
     * The export schedule of the bucket inventory.
     */
    schedule: pulumi.Input<inputs.tos.BucketInventorySchedule>;
}
