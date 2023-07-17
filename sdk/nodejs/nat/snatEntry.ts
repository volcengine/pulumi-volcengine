// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage snat entry
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = new volcengine.nat.SnatEntry("foo", {
 *     eipId: "eip-274zlae117nr47fap8tzl24v4",
 *     natGatewayId: "ngw-2743w1f6iqby87fap8tvm9kop",
 *     snatEntryName: "tf-test-up",
 *     subnetId: "subnet-2744i7u9alnnk7fap8tkq8aft",
 * });
 * ```
 *
 * ## Import
 *
 * Snat entry can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:nat/snatEntry:SnatEntry default snat-3fvhk47kf56****
 * ```
 */
export class SnatEntry extends pulumi.CustomResource {
    /**
     * Get an existing SnatEntry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnatEntryState, opts?: pulumi.CustomResourceOptions): SnatEntry {
        return new SnatEntry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:nat/snatEntry:SnatEntry';

    /**
     * Returns true if the given object is an instance of SnatEntry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SnatEntry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SnatEntry.__pulumiType;
    }

    /**
     * The id of the public ip address used by the SNAT entry.
     */
    public readonly eipId!: pulumi.Output<string>;
    /**
     * The id of the nat gateway to which the entry belongs.
     */
    public readonly natGatewayId!: pulumi.Output<string>;
    /**
     * The name of the SNAT entry.
     */
    public readonly snatEntryName!: pulumi.Output<string>;
    /**
     * The SourceCidr of the SNAT entry. Only one of `subnet_id,source_cidr` can be specified.
     */
    public readonly sourceCidr!: pulumi.Output<string | undefined>;
    /**
     * The status of the SNAT entry.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The id of the subnet that is required to access the internet. Only one of `subnet_id,source_cidr` can be specified.
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;

    /**
     * Create a SnatEntry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnatEntryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnatEntryArgs | SnatEntryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SnatEntryState | undefined;
            resourceInputs["eipId"] = state ? state.eipId : undefined;
            resourceInputs["natGatewayId"] = state ? state.natGatewayId : undefined;
            resourceInputs["snatEntryName"] = state ? state.snatEntryName : undefined;
            resourceInputs["sourceCidr"] = state ? state.sourceCidr : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
        } else {
            const args = argsOrState as SnatEntryArgs | undefined;
            if ((!args || args.eipId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eipId'");
            }
            if ((!args || args.natGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'natGatewayId'");
            }
            resourceInputs["eipId"] = args ? args.eipId : undefined;
            resourceInputs["natGatewayId"] = args ? args.natGatewayId : undefined;
            resourceInputs["snatEntryName"] = args ? args.snatEntryName : undefined;
            resourceInputs["sourceCidr"] = args ? args.sourceCidr : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SnatEntry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SnatEntry resources.
 */
export interface SnatEntryState {
    /**
     * The id of the public ip address used by the SNAT entry.
     */
    eipId?: pulumi.Input<string>;
    /**
     * The id of the nat gateway to which the entry belongs.
     */
    natGatewayId?: pulumi.Input<string>;
    /**
     * The name of the SNAT entry.
     */
    snatEntryName?: pulumi.Input<string>;
    /**
     * The SourceCidr of the SNAT entry. Only one of `subnet_id,source_cidr` can be specified.
     */
    sourceCidr?: pulumi.Input<string>;
    /**
     * The status of the SNAT entry.
     */
    status?: pulumi.Input<string>;
    /**
     * The id of the subnet that is required to access the internet. Only one of `subnet_id,source_cidr` can be specified.
     */
    subnetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SnatEntry resource.
 */
export interface SnatEntryArgs {
    /**
     * The id of the public ip address used by the SNAT entry.
     */
    eipId: pulumi.Input<string>;
    /**
     * The id of the nat gateway to which the entry belongs.
     */
    natGatewayId: pulumi.Input<string>;
    /**
     * The name of the SNAT entry.
     */
    snatEntryName?: pulumi.Input<string>;
    /**
     * The SourceCidr of the SNAT entry. Only one of `subnet_id,source_cidr` can be specified.
     */
    sourceCidr?: pulumi.Input<string>;
    /**
     * The id of the subnet that is required to access the internet. Only one of `subnet_id,source_cidr` can be specified.
     */
    subnetId?: pulumi.Input<string>;
}