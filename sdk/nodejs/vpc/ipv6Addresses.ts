// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vpc ipv6 addresses
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vpc.getIpv6Addresses({
 *     associatedInstanceId: "i-yca53yuhj6gh9zl53kav",
 * });
 * ```
 */
/** @deprecated volcengine.vpc.Ipv6Addresses has been deprecated in favor of volcengine.vpc.getIpv6Addresses */
export function ipv6Addresses(args?: Ipv6AddressesArgs, opts?: pulumi.InvokeOptions): Promise<Ipv6AddressesResult> {
    pulumi.log.warn("ipv6Addresses is deprecated: volcengine.vpc.Ipv6Addresses has been deprecated in favor of volcengine.vpc.getIpv6Addresses")
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:vpc/ipv6Addresses:Ipv6Addresses", {
        "associatedInstanceId": args.associatedInstanceId,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking Ipv6Addresses.
 */
export interface Ipv6AddressesArgs {
    /**
     * The ID of the ECS instance that is assigned the IPv6 address.
     */
    associatedInstanceId?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by Ipv6Addresses.
 */
export interface Ipv6AddressesResult {
    readonly associatedInstanceId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The collection of Ipv6Address query.
     */
    readonly ipv6Addresses: outputs.vpc.Ipv6AddressesIpv6Address[];
    readonly outputFile?: string;
    /**
     * The total count of Ipv6Address query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of vpc ipv6 addresses
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vpc.getIpv6Addresses({
 *     associatedInstanceId: "i-yca53yuhj6gh9zl53kav",
 * });
 * ```
 */
/** @deprecated volcengine.vpc.Ipv6Addresses has been deprecated in favor of volcengine.vpc.getIpv6Addresses */
export function ipv6AddressesOutput(args?: Ipv6AddressesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<Ipv6AddressesResult> {
    return pulumi.output(args).apply((a: any) => ipv6Addresses(a, opts))
}

/**
 * A collection of arguments for invoking Ipv6Addresses.
 */
export interface Ipv6AddressesOutputArgs {
    /**
     * The ID of the ECS instance that is assigned the IPv6 address.
     */
    associatedInstanceId?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
