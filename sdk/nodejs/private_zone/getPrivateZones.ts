// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of private zones
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.private_zone.getPrivateZones({
 *     lineMode: 3,
 *     recursionMode: true,
 *     searchMode: "EXACT",
 *     zid: 770000,
 *     zoneName: "volces.com",
 * });
 * ```
 */
export function getPrivateZones(args?: GetPrivateZonesArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivateZonesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:private_zone/getPrivateZones:getPrivateZones", {
        "keyWord": args.keyWord,
        "lineMode": args.lineMode,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "projectName": args.projectName,
        "recursionMode": args.recursionMode,
        "region": args.region,
        "searchMode": args.searchMode,
        "tagFilters": args.tagFilters,
        "vpcId": args.vpcId,
        "zid": args.zid,
        "zoneName": args.zoneName,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrivateZones.
 */
export interface GetPrivateZonesArgs {
    /**
     * The keyword of zone name.
     */
    keyWord?: string;
    /**
     * The line mode of Private Zone, specified whether the intelligent mode and the load balance function is enabled.
     */
    lineMode?: number;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The project name of the private zone.
     */
    projectName?: string;
    /**
     * Whether the recursion mode of Private Zone is enabled.
     */
    recursionMode?: boolean;
    /**
     * The region of Private Zone.
     */
    region?: string;
    /**
     * The search mode of query. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
     */
    searchMode?: string;
    /**
     * List of tag filters.
     */
    tagFilters?: inputs.private_zone.GetPrivateZonesTagFilter[];
    /**
     * The vpc id associated to Private Zone.
     */
    vpcId?: string;
    /**
     * The zid of Private Zone.
     */
    zid?: number;
    /**
     * The name of Private Zone.
     */
    zoneName?: string;
}

/**
 * A collection of values returned by getPrivateZones.
 */
export interface GetPrivateZonesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keyWord?: string;
    /**
     * The line mode of the private zone, specified whether the intelligent mode and the load balance function is enabled.
     */
    readonly lineMode?: number;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The collection of query.
     */
    readonly privateZones: outputs.private_zone.GetPrivateZonesPrivateZone[];
    /**
     * The project name of the private zone.
     */
    readonly projectName?: string;
    /**
     * Whether the recursion mode of the private zone is enabled.
     */
    readonly recursionMode?: boolean;
    /**
     * The region of the private zone.
     */
    readonly region?: string;
    readonly searchMode?: string;
    readonly tagFilters?: outputs.private_zone.GetPrivateZonesTagFilter[];
    /**
     * The total count of query.
     */
    readonly totalCount: number;
    readonly vpcId?: string;
    /**
     * The id of the private zone.
     */
    readonly zid?: number;
    /**
     * The id of the private zone.
     */
    readonly zoneName?: string;
}
/**
 * Use this data source to query detailed information of private zones
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.private_zone.getPrivateZones({
 *     lineMode: 3,
 *     recursionMode: true,
 *     searchMode: "EXACT",
 *     zid: 770000,
 *     zoneName: "volces.com",
 * });
 * ```
 */
export function getPrivateZonesOutput(args?: GetPrivateZonesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPrivateZonesResult> {
    return pulumi.output(args).apply((a: any) => getPrivateZones(a, opts))
}

/**
 * A collection of arguments for invoking getPrivateZones.
 */
export interface GetPrivateZonesOutputArgs {
    /**
     * The keyword of zone name.
     */
    keyWord?: pulumi.Input<string>;
    /**
     * The line mode of Private Zone, specified whether the intelligent mode and the load balance function is enabled.
     */
    lineMode?: pulumi.Input<number>;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The project name of the private zone.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Whether the recursion mode of Private Zone is enabled.
     */
    recursionMode?: pulumi.Input<boolean>;
    /**
     * The region of Private Zone.
     */
    region?: pulumi.Input<string>;
    /**
     * The search mode of query. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
     */
    searchMode?: pulumi.Input<string>;
    /**
     * List of tag filters.
     */
    tagFilters?: pulumi.Input<pulumi.Input<inputs.private_zone.GetPrivateZonesTagFilterArgs>[]>;
    /**
     * The vpc id associated to Private Zone.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The zid of Private Zone.
     */
    zid?: pulumi.Input<number>;
    /**
     * The name of Private Zone.
     */
    zoneName?: pulumi.Input<string>;
}
