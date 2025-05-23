// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of eip addresses
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooAddress = new volcengine.eip.Address("fooAddress", {billingType: "PostPaidByTraffic"});
 * const fooAddresses = volcengine.eip.getAddressesOutput({
 *     ids: [fooAddress.id],
 * });
 * ```
 */
export function getAddresses(args?: GetAddressesArgs, opts?: pulumi.InvokeOptions): Promise<GetAddressesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:eip/getAddresses:getAddresses", {
        "associatedInstanceId": args.associatedInstanceId,
        "associatedInstanceType": args.associatedInstanceType,
        "eipAddresses": args.eipAddresses,
        "ids": args.ids,
        "isp": args.isp,
        "name": args.name,
        "outputFile": args.outputFile,
        "projectName": args.projectName,
        "status": args.status,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getAddresses.
 */
export interface GetAddressesArgs {
    /**
     * An id of associated instance.
     */
    associatedInstanceId?: string;
    /**
     * A type of associated instance, the value can be `Nat`, `NetworkInterface`, `ClbInstance`, `AlbInstance`, `HaVip` or `EcsInstance`.
     */
    associatedInstanceType?: string;
    /**
     * A list of EIP ip address that you want to query.
     */
    eipAddresses?: string[];
    /**
     * A list of EIP allocation ids.
     */
    ids?: string[];
    /**
     * An ISP of EIP Address, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom`.
     */
    isp?: string;
    /**
     * A name of EIP.
     */
    name?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The ProjectName of EIP.
     */
    projectName?: string;
    /**
     * A status of EIP, the value can be `Attaching` or `Detaching` or `Attached` or `Available`.
     */
    status?: string;
    /**
     * Tags.
     */
    tags?: inputs.eip.GetAddressesTag[];
}

/**
 * A collection of values returned by getAddresses.
 */
export interface GetAddressesResult {
    /**
     * The collection of EIP addresses.
     */
    readonly addresses: outputs.eip.GetAddressesAddress[];
    readonly associatedInstanceId?: string;
    readonly associatedInstanceType?: string;
    readonly eipAddresses?: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The ISP of EIP Address.
     */
    readonly isp?: string;
    /**
     * The name of the EIP.
     */
    readonly name?: string;
    readonly outputFile?: string;
    /**
     * The ProjectName of the EIP.
     */
    readonly projectName?: string;
    /**
     * The status of the EIP.
     */
    readonly status?: string;
    /**
     * Tags.
     */
    readonly tags?: outputs.eip.GetAddressesTag[];
    /**
     * The total count of EIP addresses query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of eip addresses
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooAddress = new volcengine.eip.Address("fooAddress", {billingType: "PostPaidByTraffic"});
 * const fooAddresses = volcengine.eip.getAddressesOutput({
 *     ids: [fooAddress.id],
 * });
 * ```
 */
export function getAddressesOutput(args?: GetAddressesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAddressesResult> {
    return pulumi.output(args).apply((a: any) => getAddresses(a, opts))
}

/**
 * A collection of arguments for invoking getAddresses.
 */
export interface GetAddressesOutputArgs {
    /**
     * An id of associated instance.
     */
    associatedInstanceId?: pulumi.Input<string>;
    /**
     * A type of associated instance, the value can be `Nat`, `NetworkInterface`, `ClbInstance`, `AlbInstance`, `HaVip` or `EcsInstance`.
     */
    associatedInstanceType?: pulumi.Input<string>;
    /**
     * A list of EIP ip address that you want to query.
     */
    eipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of EIP allocation ids.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An ISP of EIP Address, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom`.
     */
    isp?: pulumi.Input<string>;
    /**
     * A name of EIP.
     */
    name?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The ProjectName of EIP.
     */
    projectName?: pulumi.Input<string>;
    /**
     * A status of EIP, the value can be `Attaching` or `Detaching` or `Attached` or `Available`.
     */
    status?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.eip.GetAddressesTagArgs>[]>;
}
