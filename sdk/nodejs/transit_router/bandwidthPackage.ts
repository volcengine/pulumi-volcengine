// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage transit router bandwidth package
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.transit_router.BandwidthPackage("foo", {
 *     bandwidth: 2,
 *     description: "acc-test",
 *     period: 1,
 *     renewType: "Manual",
 *     transitRouterBandwidthPackageName: "acc-tf-test",
 * });
 * ```
 *
 * ## Import
 *
 * TransitRouterBandwidthPackage can be imported using the Id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:transit_router/bandwidthPackage:BandwidthPackage default tbp-cd-2felfww0i6pkw59gp68bq****
 * ```
 */
export class BandwidthPackage extends pulumi.CustomResource {
    /**
     * Get an existing BandwidthPackage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BandwidthPackageState, opts?: pulumi.CustomResourceOptions): BandwidthPackage {
        return new BandwidthPackage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:transit_router/bandwidthPackage:BandwidthPackage';

    /**
     * Returns true if the given object is an instance of BandwidthPackage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BandwidthPackage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BandwidthPackage.__pulumiType;
    }

    /**
     * The detailed information on cross regional connections associated with bandwidth packets.
     */
    public /*out*/ readonly allocations!: pulumi.Output<outputs.transit_router.BandwidthPackageAllocation[]>;
    /**
     * The bandwidth peak of the transit router bandwidth package. Unit: Mbps. Valid values: 2-10000. Default is 2 Mbps.
     */
    public readonly bandwidth!: pulumi.Output<number | undefined>;
    /**
     * The business status of the transit router bandwidth package.
     */
    public /*out*/ readonly businessStatus!: pulumi.Output<string>;
    /**
     * The create time of the transit router bandwidth package.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The delete time of the transit router bandwidth package.
     */
    public /*out*/ readonly deleteTime!: pulumi.Output<string>;
    /**
     * The description of the transit router bandwidth package.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The expired time of the transit router bandwidth package.
     */
    public /*out*/ readonly expiredTime!: pulumi.Output<string>;
    /**
     * The period of the transit router bandwidth package, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.The modification of this field only takes effect when the value of the `renewType` is `Manual`.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The remaining renewal times of of the transit router bandwidth package. Valid values: -1 or 1~100. Default value is -1, means unlimited renewal.This field is only effective when the value of the `renewType` is `Auto`.
     */
    public readonly remainRenewTimes!: pulumi.Output<number | undefined>;
    /**
     * The remaining bandwidth of the transit router bandwidth package. Unit: Mbps.
     */
    public /*out*/ readonly remainingBandwidth!: pulumi.Output<number>;
    /**
     * The auto renewal period of the transit router bandwidth package. Valid values: 1,2,3,6,12. Default value is 1. Unit: Month.This field is only effective when the value of the `renewType` is `Auto`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly renewPeriod!: pulumi.Output<number | undefined>;
    /**
     * The renewal type of the transit router bandwidth package. Valid values: `Manual`, `Auto`, `NoRenew`. Default is `Manual`.This field is only effective when modifying the bandwidth package.
     */
    public readonly renewType!: pulumi.Output<string | undefined>;
    /**
     * The status of the transit router bandwidth package.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The name of the transit router bandwidth package.
     */
    public readonly transitRouterBandwidthPackageName!: pulumi.Output<string>;
    /**
     * The update time of the transit router bandwidth package.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a BandwidthPackage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: BandwidthPackageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BandwidthPackageArgs | BandwidthPackageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BandwidthPackageState | undefined;
            resourceInputs["allocations"] = state ? state.allocations : undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["businessStatus"] = state ? state.businessStatus : undefined;
            resourceInputs["creationTime"] = state ? state.creationTime : undefined;
            resourceInputs["deleteTime"] = state ? state.deleteTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["expiredTime"] = state ? state.expiredTime : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["remainRenewTimes"] = state ? state.remainRenewTimes : undefined;
            resourceInputs["remainingBandwidth"] = state ? state.remainingBandwidth : undefined;
            resourceInputs["renewPeriod"] = state ? state.renewPeriod : undefined;
            resourceInputs["renewType"] = state ? state.renewType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["transitRouterBandwidthPackageName"] = state ? state.transitRouterBandwidthPackageName : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as BandwidthPackageArgs | undefined;
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["remainRenewTimes"] = args ? args.remainRenewTimes : undefined;
            resourceInputs["renewPeriod"] = args ? args.renewPeriod : undefined;
            resourceInputs["renewType"] = args ? args.renewType : undefined;
            resourceInputs["transitRouterBandwidthPackageName"] = args ? args.transitRouterBandwidthPackageName : undefined;
            resourceInputs["allocations"] = undefined /*out*/;
            resourceInputs["businessStatus"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["deleteTime"] = undefined /*out*/;
            resourceInputs["expiredTime"] = undefined /*out*/;
            resourceInputs["remainingBandwidth"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BandwidthPackage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BandwidthPackage resources.
 */
export interface BandwidthPackageState {
    /**
     * The detailed information on cross regional connections associated with bandwidth packets.
     */
    allocations?: pulumi.Input<pulumi.Input<inputs.transit_router.BandwidthPackageAllocation>[]>;
    /**
     * The bandwidth peak of the transit router bandwidth package. Unit: Mbps. Valid values: 2-10000. Default is 2 Mbps.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * The business status of the transit router bandwidth package.
     */
    businessStatus?: pulumi.Input<string>;
    /**
     * The create time of the transit router bandwidth package.
     */
    creationTime?: pulumi.Input<string>;
    /**
     * The delete time of the transit router bandwidth package.
     */
    deleteTime?: pulumi.Input<string>;
    /**
     * The description of the transit router bandwidth package.
     */
    description?: pulumi.Input<string>;
    /**
     * The expired time of the transit router bandwidth package.
     */
    expiredTime?: pulumi.Input<string>;
    /**
     * The period of the transit router bandwidth package, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.The modification of this field only takes effect when the value of the `renewType` is `Manual`.
     */
    period?: pulumi.Input<number>;
    /**
     * The remaining renewal times of of the transit router bandwidth package. Valid values: -1 or 1~100. Default value is -1, means unlimited renewal.This field is only effective when the value of the `renewType` is `Auto`.
     */
    remainRenewTimes?: pulumi.Input<number>;
    /**
     * The remaining bandwidth of the transit router bandwidth package. Unit: Mbps.
     */
    remainingBandwidth?: pulumi.Input<number>;
    /**
     * The auto renewal period of the transit router bandwidth package. Valid values: 1,2,3,6,12. Default value is 1. Unit: Month.This field is only effective when the value of the `renewType` is `Auto`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    renewPeriod?: pulumi.Input<number>;
    /**
     * The renewal type of the transit router bandwidth package. Valid values: `Manual`, `Auto`, `NoRenew`. Default is `Manual`.This field is only effective when modifying the bandwidth package.
     */
    renewType?: pulumi.Input<string>;
    /**
     * The status of the transit router bandwidth package.
     */
    status?: pulumi.Input<string>;
    /**
     * The name of the transit router bandwidth package.
     */
    transitRouterBandwidthPackageName?: pulumi.Input<string>;
    /**
     * The update time of the transit router bandwidth package.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BandwidthPackage resource.
 */
export interface BandwidthPackageArgs {
    /**
     * The bandwidth peak of the transit router bandwidth package. Unit: Mbps. Valid values: 2-10000. Default is 2 Mbps.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * The description of the transit router bandwidth package.
     */
    description?: pulumi.Input<string>;
    /**
     * The period of the transit router bandwidth package, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.The modification of this field only takes effect when the value of the `renewType` is `Manual`.
     */
    period?: pulumi.Input<number>;
    /**
     * The remaining renewal times of of the transit router bandwidth package. Valid values: -1 or 1~100. Default value is -1, means unlimited renewal.This field is only effective when the value of the `renewType` is `Auto`.
     */
    remainRenewTimes?: pulumi.Input<number>;
    /**
     * The auto renewal period of the transit router bandwidth package. Valid values: 1,2,3,6,12. Default value is 1. Unit: Month.This field is only effective when the value of the `renewType` is `Auto`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    renewPeriod?: pulumi.Input<number>;
    /**
     * The renewal type of the transit router bandwidth package. Valid values: `Manual`, `Auto`, `NoRenew`. Default is `Manual`.This field is only effective when modifying the bandwidth package.
     */
    renewType?: pulumi.Input<string>;
    /**
     * The name of the transit router bandwidth package.
     */
    transitRouterBandwidthPackageName?: pulumi.Input<string>;
}