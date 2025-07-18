// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cen inter region bandwidths
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
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
 * const fooInterRegionBandwidths = volcengine.cen.getInterRegionBandwidthsOutput({
 *     ids: [fooInterRegionBandwidth.id],
 * });
 * ```
 */
/** @deprecated volcengine.cen.InterRegionBandwidths has been deprecated in favor of volcengine.cen.getInterRegionBandwidths */
export function interRegionBandwidths(args?: InterRegionBandwidthsArgs, opts?: pulumi.InvokeOptions): Promise<InterRegionBandwidthsResult> {
    pulumi.log.warn("interRegionBandwidths is deprecated: volcengine.cen.InterRegionBandwidths has been deprecated in favor of volcengine.cen.getInterRegionBandwidths")
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:cen/interRegionBandwidths:InterRegionBandwidths", {
        "cenId": args.cenId,
        "ids": args.ids,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking InterRegionBandwidths.
 */
export interface InterRegionBandwidthsArgs {
    /**
     * The ID of the cen.
     */
    cenId?: string;
    /**
     * A list of cen inter region bandwidth IDs.
     */
    ids?: string[];
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by InterRegionBandwidths.
 */
export interface InterRegionBandwidthsResult {
    /**
     * The cen ID of the cen inter region bandwidth.
     */
    readonly cenId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The collection of cen inter region bandwidth query.
     */
    readonly interRegionBandwidths: outputs.cen.InterRegionBandwidthsInterRegionBandwidth[];
    readonly outputFile?: string;
    /**
     * The total count of cen inter region bandwidth query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of cen inter region bandwidths
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
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
 * const fooInterRegionBandwidths = volcengine.cen.getInterRegionBandwidthsOutput({
 *     ids: [fooInterRegionBandwidth.id],
 * });
 * ```
 */
/** @deprecated volcengine.cen.InterRegionBandwidths has been deprecated in favor of volcengine.cen.getInterRegionBandwidths */
export function interRegionBandwidthsOutput(args?: InterRegionBandwidthsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<InterRegionBandwidthsResult> {
    return pulumi.output(args).apply((a: any) => interRegionBandwidths(a, opts))
}

/**
 * A collection of arguments for invoking InterRegionBandwidths.
 */
export interface InterRegionBandwidthsOutputArgs {
    /**
     * The ID of the cen.
     */
    cenId?: pulumi.Input<string>;
    /**
     * A list of cen inter region bandwidth IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
