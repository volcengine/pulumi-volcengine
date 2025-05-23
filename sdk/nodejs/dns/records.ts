// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of dns records
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.dns.getRecords({
 *     zid: 58857,
 * });
 * ```
 */
/** @deprecated volcengine.dns.Records has been deprecated in favor of volcengine.dns.getRecords */
export function records(args: RecordsArgs, opts?: pulumi.InvokeOptions): Promise<RecordsResult> {
    pulumi.log.warn("records is deprecated: volcengine.dns.Records has been deprecated in favor of volcengine.dns.getRecords")

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:dns/records:Records", {
        "host": args.host,
        "line": args.line,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "searchMode": args.searchMode,
        "searchOrder": args.searchOrder,
        "type": args.type,
        "value": args.value,
        "zid": args.zid,
    }, opts);
}

/**
 * A collection of arguments for invoking Records.
 */
export interface RecordsArgs {
    /**
     * Domain prefix of the DNS record.
     */
    host?: string;
    /**
     * Line of the DNS record.
     */
    line?: string;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The matching mode for the Host parameter.
     */
    searchMode?: string;
    /**
     * The Method to sort the returned list of DNS records.
     */
    searchOrder?: string;
    /**
     * Type of the DNS record.
     */
    type?: string;
    /**
     * Value of the DNS record.
     */
    value?: string;
    /**
     * The ID of the domain.
     */
    zid: number;
}

/**
 * A collection of values returned by Records.
 */
export interface RecordsResult {
    /**
     * The host record included in the DNS record.
     */
    readonly host?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The line code corresponding to the DNS record.
     */
    readonly line?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The collection of query.
     */
    readonly records: outputs.dns.RecordsRecord[];
    readonly searchMode?: string;
    readonly searchOrder?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
    /**
     * The type of the DNS record.
     */
    readonly type?: string;
    /**
     * The record value contained in the DNS record.
     */
    readonly value?: string;
    readonly zid: number;
}
/**
 * Use this data source to query detailed information of dns records
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.dns.getRecords({
 *     zid: 58857,
 * });
 * ```
 */
/** @deprecated volcengine.dns.Records has been deprecated in favor of volcengine.dns.getRecords */
export function recordsOutput(args: RecordsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<RecordsResult> {
    return pulumi.output(args).apply((a: any) => records(a, opts))
}

/**
 * A collection of arguments for invoking Records.
 */
export interface RecordsOutputArgs {
    /**
     * Domain prefix of the DNS record.
     */
    host?: pulumi.Input<string>;
    /**
     * Line of the DNS record.
     */
    line?: pulumi.Input<string>;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The matching mode for the Host parameter.
     */
    searchMode?: pulumi.Input<string>;
    /**
     * The Method to sort the returned list of DNS records.
     */
    searchOrder?: pulumi.Input<string>;
    /**
     * Type of the DNS record.
     */
    type?: pulumi.Input<string>;
    /**
     * Value of the DNS record.
     */
    value?: pulumi.Input<string>;
    /**
     * The ID of the domain.
     */
    zid: pulumi.Input<number>;
}
