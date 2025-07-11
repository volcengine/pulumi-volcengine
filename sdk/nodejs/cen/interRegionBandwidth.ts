// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage cen inter region bandwidth
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooCen = new volcengine.cen.Cen("fooCen", {
 *     cenName: "acc-test-cen",
 *     description: "acc-test",
 *     projectName: "default",
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * const fooBandwidthPackage = new volcengine.cen.BandwidthPackage("fooBandwidthPackage", {
 *     localGeographicRegionSetId: "China",
 *     peerGeographicRegionSetId: "China",
 *     bandwidth: 5,
 *     cenBandwidthPackageName: "acc-test-cen-bp",
 *     description: "acc-test",
 *     billingType: "PrePaid",
 *     periodUnit: "Month",
 *     period: 1,
 *     projectName: "default",
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * const fooBandwidthPackageAssociate = new volcengine.cen.BandwidthPackageAssociate("fooBandwidthPackageAssociate", {
 *     cenBandwidthPackageId: fooBandwidthPackage.id,
 *     cenId: fooCen.id,
 * });
 * const fooInterRegionBandwidth = new volcengine.cen.InterRegionBandwidth("fooInterRegionBandwidth", {
 *     cenId: fooCen.id,
 *     localRegionId: "cn-beijing",
 *     peerRegionId: "cn-shanghai",
 *     bandwidth: 2,
 * }, {
 *     dependsOn: [fooBandwidthPackageAssociate],
 * });
 * ```
 *
 * ## Import
 *
 * CenInterRegionBandwidth can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:cen/interRegionBandwidth:InterRegionBandwidth default cirb-3tex2x1cwd4c6c0v****
 * ```
 */
export class InterRegionBandwidth extends pulumi.CustomResource {
    /**
     * Get an existing InterRegionBandwidth resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InterRegionBandwidthState, opts?: pulumi.CustomResourceOptions): InterRegionBandwidth {
        return new InterRegionBandwidth(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:cen/interRegionBandwidth:InterRegionBandwidth';

    /**
     * Returns true if the given object is an instance of InterRegionBandwidth.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InterRegionBandwidth {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InterRegionBandwidth.__pulumiType;
    }

    /**
     * The bandwidth of the cen inter region bandwidth.
     */
    public readonly bandwidth!: pulumi.Output<number>;
    /**
     * The cen bandwidth package id of the cen inter region bandwidth.
     */
    public readonly cenBandwidthPackageId!: pulumi.Output<string>;
    /**
     * The cen ID of the cen inter region bandwidth.
     */
    public readonly cenId!: pulumi.Output<string>;
    /**
     * The create time of the cen inter region bandwidth.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The ID of the cen inter region bandwidth.
     */
    public /*out*/ readonly interRegionBandwidthId!: pulumi.Output<string>;
    /**
     * The local region id of the cen inter region bandwidth.
     */
    public readonly localRegionId!: pulumi.Output<string>;
    /**
     * The peer region id of the cen inter region bandwidth.
     */
    public readonly peerRegionId!: pulumi.Output<string>;
    /**
     * The status of the cen inter region bandwidth.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The update time of the cen inter region bandwidth.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a InterRegionBandwidth resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InterRegionBandwidthArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InterRegionBandwidthArgs | InterRegionBandwidthState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InterRegionBandwidthState | undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["cenBandwidthPackageId"] = state ? state.cenBandwidthPackageId : undefined;
            resourceInputs["cenId"] = state ? state.cenId : undefined;
            resourceInputs["creationTime"] = state ? state.creationTime : undefined;
            resourceInputs["interRegionBandwidthId"] = state ? state.interRegionBandwidthId : undefined;
            resourceInputs["localRegionId"] = state ? state.localRegionId : undefined;
            resourceInputs["peerRegionId"] = state ? state.peerRegionId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as InterRegionBandwidthArgs | undefined;
            if ((!args || args.bandwidth === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidth'");
            }
            if ((!args || args.cenId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cenId'");
            }
            if ((!args || args.localRegionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'localRegionId'");
            }
            if ((!args || args.peerRegionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'peerRegionId'");
            }
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["cenBandwidthPackageId"] = args ? args.cenBandwidthPackageId : undefined;
            resourceInputs["cenId"] = args ? args.cenId : undefined;
            resourceInputs["localRegionId"] = args ? args.localRegionId : undefined;
            resourceInputs["peerRegionId"] = args ? args.peerRegionId : undefined;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["interRegionBandwidthId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InterRegionBandwidth.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InterRegionBandwidth resources.
 */
export interface InterRegionBandwidthState {
    /**
     * The bandwidth of the cen inter region bandwidth.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * The cen bandwidth package id of the cen inter region bandwidth.
     */
    cenBandwidthPackageId?: pulumi.Input<string>;
    /**
     * The cen ID of the cen inter region bandwidth.
     */
    cenId?: pulumi.Input<string>;
    /**
     * The create time of the cen inter region bandwidth.
     */
    creationTime?: pulumi.Input<string>;
    /**
     * The ID of the cen inter region bandwidth.
     */
    interRegionBandwidthId?: pulumi.Input<string>;
    /**
     * The local region id of the cen inter region bandwidth.
     */
    localRegionId?: pulumi.Input<string>;
    /**
     * The peer region id of the cen inter region bandwidth.
     */
    peerRegionId?: pulumi.Input<string>;
    /**
     * The status of the cen inter region bandwidth.
     */
    status?: pulumi.Input<string>;
    /**
     * The update time of the cen inter region bandwidth.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InterRegionBandwidth resource.
 */
export interface InterRegionBandwidthArgs {
    /**
     * The bandwidth of the cen inter region bandwidth.
     */
    bandwidth: pulumi.Input<number>;
    /**
     * The cen bandwidth package id of the cen inter region bandwidth.
     */
    cenBandwidthPackageId?: pulumi.Input<string>;
    /**
     * The cen ID of the cen inter region bandwidth.
     */
    cenId: pulumi.Input<string>;
    /**
     * The local region id of the cen inter region bandwidth.
     */
    localRegionId: pulumi.Input<string>;
    /**
     * The peer region id of the cen inter region bandwidth.
     */
    peerRegionId: pulumi.Input<string>;
}
