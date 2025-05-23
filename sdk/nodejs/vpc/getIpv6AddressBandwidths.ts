// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vpc ipv6 address bandwidths
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vpc.getIpv6AddressBandwidths({
 *     ids: ["eip-in2y2duvtlhc8gbssyfnhfre"],
 * });
 * ```
 */
export function getIpv6AddressBandwidths(args?: GetIpv6AddressBandwidthsArgs, opts?: pulumi.InvokeOptions): Promise<GetIpv6AddressBandwidthsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:vpc/getIpv6AddressBandwidths:getIpv6AddressBandwidths", {
        "associatedInstanceId": args.associatedInstanceId,
        "associatedInstanceType": args.associatedInstanceType,
        "ids": args.ids,
        "ipv6Addresses": args.ipv6Addresses,
        "isp": args.isp,
        "networkType": args.networkType,
        "outputFile": args.outputFile,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpv6AddressBandwidths.
 */
export interface GetIpv6AddressBandwidthsArgs {
    /**
     * The ID of the associated instance.
     */
    associatedInstanceId?: string;
    /**
     * The type of the associated instance.
     */
    associatedInstanceType?: string;
    /**
     * Allocation IDs of the Ipv6 address width.
     */
    ids?: string[];
    /**
     * The ipv6 addresses.
     */
    ipv6Addresses?: string[];
    /**
     * ISP of the ipv6 address.
     */
    isp?: string;
    /**
     * The network type of the ipv6 address.
     */
    networkType?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The ID of Vpc the ipv6 address in.
     */
    vpcId?: string;
}

/**
 * A collection of values returned by getIpv6AddressBandwidths.
 */
export interface GetIpv6AddressBandwidthsResult {
    readonly associatedInstanceId?: string;
    readonly associatedInstanceType?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The collection of Ipv6AddressBandwidth query.
     */
    readonly ipv6AddressBandwidths: outputs.vpc.GetIpv6AddressBandwidthsIpv6AddressBandwidth[];
    readonly ipv6Addresses?: string[];
    /**
     * The ISP of the Ipv6AddressBandwidth.
     */
    readonly isp?: string;
    /**
     * The network type of the Ipv6AddressBandwidth.
     */
    readonly networkType?: string;
    readonly outputFile?: string;
    /**
     * The total count of Ipv6AddressBandwidth query.
     */
    readonly totalCount: number;
    readonly vpcId?: string;
}
/**
 * Use this data source to query detailed information of vpc ipv6 address bandwidths
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vpc.getIpv6AddressBandwidths({
 *     ids: ["eip-in2y2duvtlhc8gbssyfnhfre"],
 * });
 * ```
 */
export function getIpv6AddressBandwidthsOutput(args?: GetIpv6AddressBandwidthsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIpv6AddressBandwidthsResult> {
    return pulumi.output(args).apply((a: any) => getIpv6AddressBandwidths(a, opts))
}

/**
 * A collection of arguments for invoking getIpv6AddressBandwidths.
 */
export interface GetIpv6AddressBandwidthsOutputArgs {
    /**
     * The ID of the associated instance.
     */
    associatedInstanceId?: pulumi.Input<string>;
    /**
     * The type of the associated instance.
     */
    associatedInstanceType?: pulumi.Input<string>;
    /**
     * Allocation IDs of the Ipv6 address width.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ipv6 addresses.
     */
    ipv6Addresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ISP of the ipv6 address.
     */
    isp?: pulumi.Input<string>;
    /**
     * The network type of the ipv6 address.
     */
    networkType?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The ID of Vpc the ipv6 address in.
     */
    vpcId?: pulumi.Input<string>;
}
