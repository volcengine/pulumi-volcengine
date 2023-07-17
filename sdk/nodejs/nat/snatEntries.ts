// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of snat entries
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultSnatEntries = pulumi.output(volcengine.nat.SnatEntries({
 *     ids: ["snat-274zl8b1kxzb47fap8u35uune"],
 * }));
 * ```
 */
export function snatEntries(args?: SnatEntriesArgs, opts?: pulumi.InvokeOptions): Promise<SnatEntriesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:nat/snatEntries:SnatEntries", {
        "eipId": args.eipId,
        "ids": args.ids,
        "natGatewayId": args.natGatewayId,
        "outputFile": args.outputFile,
        "snatEntryName": args.snatEntryName,
        "sourceCidr": args.sourceCidr,
        "subnetId": args.subnetId,
    }, opts);
}

/**
 * A collection of arguments for invoking SnatEntries.
 */
export interface SnatEntriesArgs {
    /**
     * An id of the public ip address used by the SNAT entry.
     */
    eipId?: string;
    /**
     * A list of SNAT entry ids.
     */
    ids?: string[];
    /**
     * An id of the nat gateway to which the entry belongs.
     */
    natGatewayId?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * A name of SNAT entry.
     */
    snatEntryName?: string;
    /**
     * The SourceCidr of SNAT entry.
     */
    sourceCidr?: string;
    /**
     * An id of the subnet that is required to access the Internet.
     */
    subnetId?: string;
}

/**
 * A collection of values returned by SnatEntries.
 */
export interface SnatEntriesResult {
    /**
     * The id of the public ip address used by the SNAT entry.
     */
    readonly eipId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The id of the nat gateway to which the entry belongs.
     */
    readonly natGatewayId?: string;
    readonly outputFile?: string;
    /**
     * The collection of snat entries.
     */
    readonly snatEntries: outputs.nat.SnatEntriesSnatEntry[];
    /**
     * The name of the SNAT entry.
     */
    readonly snatEntryName?: string;
    /**
     * The SourceCidr of the SNAT entry.
     */
    readonly sourceCidr?: string;
    /**
     * The id of the subnet that is required to access the internet.
     */
    readonly subnetId?: string;
    /**
     * The total count of snat entries query.
     */
    readonly totalCount: number;
}

export function snatEntriesOutput(args?: SnatEntriesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<SnatEntriesResult> {
    return pulumi.output(args).apply(a => snatEntries(a, opts))
}

/**
 * A collection of arguments for invoking SnatEntries.
 */
export interface SnatEntriesOutputArgs {
    /**
     * An id of the public ip address used by the SNAT entry.
     */
    eipId?: pulumi.Input<string>;
    /**
     * A list of SNAT entry ids.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An id of the nat gateway to which the entry belongs.
     */
    natGatewayId?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * A name of SNAT entry.
     */
    snatEntryName?: pulumi.Input<string>;
    /**
     * The SourceCidr of SNAT entry.
     */
    sourceCidr?: pulumi.Input<string>;
    /**
     * An id of the subnet that is required to access the Internet.
     */
    subnetId?: pulumi.Input<string>;
}