// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.bandwidth_package.BandwidthPackage("foo", {
 *     bandwidth: 10,
 *     bandwidthPackageName: "acc-test-bp",
 *     billingType: "PostPaidByBandwidth",
 *     description: "acc-test",
 *     isp: "BGP",
 *     protocol: "IPv4",
 *     securityProtectionTypes: ["AntiDDoS_Enhanced"],
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * BandwidthPackage can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:bandwidth_package/bandwidthPackage:BandwidthPackage default bwp-2zeo05qre24nhrqpy****
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
    public static readonly __pulumiType = 'volcengine:bandwidth_package/bandwidthPackage:BandwidthPackage';

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
     * Bandwidth upper limit of shared bandwidth package, unit: Mbps. When BillingType is set to PrePaid: the value range is 5 to 5000. When BillingType is set to PostPaidByBandwidth: the value range is 2 to 5000. When BillingType is set to PostPaidByTraffic: the value range is 2 to 2000. When BillingType is set to PayBy95Peak: the value range is 2 to 5000.
     */
    public readonly bandwidth!: pulumi.Output<number>;
    /**
     * The name of the bandwidth package.
     */
    public readonly bandwidthPackageName!: pulumi.Output<string>;
    /**
     * BillingType of the bandwidth package. Valid values: `PrePaid`, `PostPaidByBandwidth`(Default), `PostPaidByTraffic`, `PayBy95Peak`. The billing method of IPv6 does not include `PrePaid`, and the billing method is only based on the `PostPaidByBandwidth`.
     */
    public readonly billingType!: pulumi.Output<string | undefined>;
    /**
     * The description of the bandwidth package.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Route type, default to BGP.
     */
    public readonly isp!: pulumi.Output<string | undefined>;
    /**
     * Duration of purchasing shared bandwidth package on an annual or monthly basis. The valid value range in 1~9 or 12, 24 or 36. Default value is 1. The period unit defaults to `Month`.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The project name of the bandwidth package.
     */
    public readonly projectName!: pulumi.Output<string>;
    /**
     * The IP protocol values for shared bandwidth packages are as follows: `IPv4`: IPv4 protocol. `IPv6`: IPv6 protocol.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Security protection types for shared bandwidth packages. Parameter - N: Indicates the number of security protection types, currently only supports taking 1. Value: `AntiDDoS_Enhanced` or left blank.If the value is `AntiDDoS_Enhanced`, then will create a shared bandwidth package with enhanced protection, which supports adding basic protection type public IP addresses.If left blank, it indicates a shared bandwidth package with basic protection, which supports the addition of public IP addresses with enhanced protection.
     */
    public readonly securityProtectionTypes!: pulumi.Output<string[] | undefined>;
    /**
     * Tags.
     */
    public readonly tags!: pulumi.Output<outputs.bandwidth_package.BandwidthPackageTag[] | undefined>;

    /**
     * Create a BandwidthPackage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BandwidthPackageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BandwidthPackageArgs | BandwidthPackageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BandwidthPackageState | undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["bandwidthPackageName"] = state ? state.bandwidthPackageName : undefined;
            resourceInputs["billingType"] = state ? state.billingType : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["isp"] = state ? state.isp : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["securityProtectionTypes"] = state ? state.securityProtectionTypes : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as BandwidthPackageArgs | undefined;
            if ((!args || args.bandwidth === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidth'");
            }
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["bandwidthPackageName"] = args ? args.bandwidthPackageName : undefined;
            resourceInputs["billingType"] = args ? args.billingType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["isp"] = args ? args.isp : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["securityProtectionTypes"] = args ? args.securityProtectionTypes : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
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
     * Bandwidth upper limit of shared bandwidth package, unit: Mbps. When BillingType is set to PrePaid: the value range is 5 to 5000. When BillingType is set to PostPaidByBandwidth: the value range is 2 to 5000. When BillingType is set to PostPaidByTraffic: the value range is 2 to 2000. When BillingType is set to PayBy95Peak: the value range is 2 to 5000.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * The name of the bandwidth package.
     */
    bandwidthPackageName?: pulumi.Input<string>;
    /**
     * BillingType of the bandwidth package. Valid values: `PrePaid`, `PostPaidByBandwidth`(Default), `PostPaidByTraffic`, `PayBy95Peak`. The billing method of IPv6 does not include `PrePaid`, and the billing method is only based on the `PostPaidByBandwidth`.
     */
    billingType?: pulumi.Input<string>;
    /**
     * The description of the bandwidth package.
     */
    description?: pulumi.Input<string>;
    /**
     * Route type, default to BGP.
     */
    isp?: pulumi.Input<string>;
    /**
     * Duration of purchasing shared bandwidth package on an annual or monthly basis. The valid value range in 1~9 or 12, 24 or 36. Default value is 1. The period unit defaults to `Month`.
     */
    period?: pulumi.Input<number>;
    /**
     * The project name of the bandwidth package.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The IP protocol values for shared bandwidth packages are as follows: `IPv4`: IPv4 protocol. `IPv6`: IPv6 protocol.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Security protection types for shared bandwidth packages. Parameter - N: Indicates the number of security protection types, currently only supports taking 1. Value: `AntiDDoS_Enhanced` or left blank.If the value is `AntiDDoS_Enhanced`, then will create a shared bandwidth package with enhanced protection, which supports adding basic protection type public IP addresses.If left blank, it indicates a shared bandwidth package with basic protection, which supports the addition of public IP addresses with enhanced protection.
     */
    securityProtectionTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.bandwidth_package.BandwidthPackageTag>[]>;
}

/**
 * The set of arguments for constructing a BandwidthPackage resource.
 */
export interface BandwidthPackageArgs {
    /**
     * Bandwidth upper limit of shared bandwidth package, unit: Mbps. When BillingType is set to PrePaid: the value range is 5 to 5000. When BillingType is set to PostPaidByBandwidth: the value range is 2 to 5000. When BillingType is set to PostPaidByTraffic: the value range is 2 to 2000. When BillingType is set to PayBy95Peak: the value range is 2 to 5000.
     */
    bandwidth: pulumi.Input<number>;
    /**
     * The name of the bandwidth package.
     */
    bandwidthPackageName?: pulumi.Input<string>;
    /**
     * BillingType of the bandwidth package. Valid values: `PrePaid`, `PostPaidByBandwidth`(Default), `PostPaidByTraffic`, `PayBy95Peak`. The billing method of IPv6 does not include `PrePaid`, and the billing method is only based on the `PostPaidByBandwidth`.
     */
    billingType?: pulumi.Input<string>;
    /**
     * The description of the bandwidth package.
     */
    description?: pulumi.Input<string>;
    /**
     * Route type, default to BGP.
     */
    isp?: pulumi.Input<string>;
    /**
     * Duration of purchasing shared bandwidth package on an annual or monthly basis. The valid value range in 1~9 or 12, 24 or 36. Default value is 1. The period unit defaults to `Month`.
     */
    period?: pulumi.Input<number>;
    /**
     * The project name of the bandwidth package.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The IP protocol values for shared bandwidth packages are as follows: `IPv4`: IPv4 protocol. `IPv6`: IPv6 protocol.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Security protection types for shared bandwidth packages. Parameter - N: Indicates the number of security protection types, currently only supports taking 1. Value: `AntiDDoS_Enhanced` or left blank.If the value is `AntiDDoS_Enhanced`, then will create a shared bandwidth package with enhanced protection, which supports adding basic protection type public IP addresses.If left blank, it indicates a shared bandwidth package with basic protection, which supports the addition of public IP addresses with enhanced protection.
     */
    securityProtectionTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.bandwidth_package.BandwidthPackageTag>[]>;
}