// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage vpc prefix list
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.vpc.PrefixList("foo", {
 *     description: "acc test description",
 *     ipVersion: "IPv4",
 *     maxEntries: 7,
 *     prefixListEntries: [
 *         {
 *             cidr: "192.168.4.0/28",
 *             description: "acc-test-1",
 *         },
 *         {
 *             cidr: "192.168.9.0/28",
 *             description: "acc-test-4",
 *         },
 *         {
 *             cidr: "192.168.8.0/28",
 *             description: "acc-test-5",
 *         },
 *     ],
 *     prefixListName: "acc-test-prefix",
 *     tags: [{
 *         key: "tf-key1",
 *         value: "tf-value1",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * VpcPrefixList can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:vpc/prefixList:PrefixList default resource_id
 * ```
 */
export class PrefixList extends pulumi.CustomResource {
    /**
     * Get an existing PrefixList resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrefixListState, opts?: pulumi.CustomResourceOptions): PrefixList {
        return new PrefixList(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:vpc/prefixList:PrefixList';

    /**
     * Returns true if the given object is an instance of PrefixList.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrefixList {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrefixList.__pulumiType;
    }

    /**
     * The description of the prefix list.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * IP version type. Possible values:
     * IPv4 (default): IPv4 type.
     * IPv6: IPv6 type.
     */
    public readonly ipVersion!: pulumi.Output<string>;
    /**
     * Maximum number of entries, which is the maximum number of entries that can be added to the prefix list. The value range is 1 to 200.
     */
    public readonly maxEntries!: pulumi.Output<number>;
    /**
     * Collection of resources associated with VPC prefix list.
     */
    public /*out*/ readonly prefixListAssociations!: pulumi.Output<outputs.vpc.PrefixListPrefixListAssociation[]>;
    /**
     * Prefix list entry list.
     */
    public readonly prefixListEntries!: pulumi.Output<outputs.vpc.PrefixListPrefixListEntry[] | undefined>;
    /**
     * The name of the prefix list.
     */
    public readonly prefixListName!: pulumi.Output<string>;
    /**
     * Tags.
     */
    public readonly tags!: pulumi.Output<outputs.vpc.PrefixListTag[] | undefined>;

    /**
     * Create a PrefixList resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrefixListArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrefixListArgs | PrefixListState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrefixListState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["ipVersion"] = state ? state.ipVersion : undefined;
            resourceInputs["maxEntries"] = state ? state.maxEntries : undefined;
            resourceInputs["prefixListAssociations"] = state ? state.prefixListAssociations : undefined;
            resourceInputs["prefixListEntries"] = state ? state.prefixListEntries : undefined;
            resourceInputs["prefixListName"] = state ? state.prefixListName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as PrefixListArgs | undefined;
            if ((!args || args.maxEntries === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxEntries'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ipVersion"] = args ? args.ipVersion : undefined;
            resourceInputs["maxEntries"] = args ? args.maxEntries : undefined;
            resourceInputs["prefixListEntries"] = args ? args.prefixListEntries : undefined;
            resourceInputs["prefixListName"] = args ? args.prefixListName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["prefixListAssociations"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrefixList.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrefixList resources.
 */
export interface PrefixListState {
    /**
     * The description of the prefix list.
     */
    description?: pulumi.Input<string>;
    /**
     * IP version type. Possible values:
     * IPv4 (default): IPv4 type.
     * IPv6: IPv6 type.
     */
    ipVersion?: pulumi.Input<string>;
    /**
     * Maximum number of entries, which is the maximum number of entries that can be added to the prefix list. The value range is 1 to 200.
     */
    maxEntries?: pulumi.Input<number>;
    /**
     * Collection of resources associated with VPC prefix list.
     */
    prefixListAssociations?: pulumi.Input<pulumi.Input<inputs.vpc.PrefixListPrefixListAssociation>[]>;
    /**
     * Prefix list entry list.
     */
    prefixListEntries?: pulumi.Input<pulumi.Input<inputs.vpc.PrefixListPrefixListEntry>[]>;
    /**
     * The name of the prefix list.
     */
    prefixListName?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.vpc.PrefixListTag>[]>;
}

/**
 * The set of arguments for constructing a PrefixList resource.
 */
export interface PrefixListArgs {
    /**
     * The description of the prefix list.
     */
    description?: pulumi.Input<string>;
    /**
     * IP version type. Possible values:
     * IPv4 (default): IPv4 type.
     * IPv6: IPv6 type.
     */
    ipVersion?: pulumi.Input<string>;
    /**
     * Maximum number of entries, which is the maximum number of entries that can be added to the prefix list. The value range is 1 to 200.
     */
    maxEntries: pulumi.Input<number>;
    /**
     * Prefix list entry list.
     */
    prefixListEntries?: pulumi.Input<pulumi.Input<inputs.vpc.PrefixListPrefixListEntry>[]>;
    /**
     * The name of the prefix list.
     */
    prefixListName?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.vpc.PrefixListTag>[]>;
}