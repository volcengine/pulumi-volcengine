// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of transit router route entries
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.transit_router.getRouteEntries({
 *     ids: ["tr-rte-12b7qd5eo3h1c17q7y1sq5ixv"],
 *     transitRouterRouteTableId: "tr-rtb-12b7qd3fmzf2817q7y2jkbd55",
 * });
 * ```
 */
export function getRouteEntries(args: GetRouteEntriesArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteEntriesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:transit_router/getRouteEntries:getRouteEntries", {
        "destinationCidrBlock": args.destinationCidrBlock,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "status": args.status,
        "transitRouterRouteEntryName": args.transitRouterRouteEntryName,
        "transitRouterRouteTableId": args.transitRouterRouteTableId,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouteEntries.
 */
export interface GetRouteEntriesArgs {
    /**
     * The target network segment of the route entry.
     */
    destinationCidrBlock?: string;
    /**
     * The ids of the transit router route entry.
     */
    ids?: string[];
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The status of the route entry.
     */
    status?: string;
    /**
     * The name of the route entry.
     */
    transitRouterRouteEntryName?: string;
    /**
     * The id of the route table.
     */
    transitRouterRouteTableId: string;
}

/**
 * A collection of values returned by getRouteEntries.
 */
export interface GetRouteEntriesResult {
    /**
     * The target network segment of the route entry.
     */
    readonly destinationCidrBlock?: string;
    /**
     * The list of route entries.
     */
    readonly entries: outputs.transit_router.GetRouteEntriesEntry[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly outputFile?: string;
    /**
     * The status of the route entry.
     */
    readonly status?: string;
    /**
     * The total count of data query.
     */
    readonly totalCount: number;
    /**
     * The name of the route entry.
     */
    readonly transitRouterRouteEntryName?: string;
    readonly transitRouterRouteTableId: string;
}
/**
 * Use this data source to query detailed information of transit router route entries
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.transit_router.getRouteEntries({
 *     ids: ["tr-rte-12b7qd5eo3h1c17q7y1sq5ixv"],
 *     transitRouterRouteTableId: "tr-rtb-12b7qd3fmzf2817q7y2jkbd55",
 * });
 * ```
 */
export function getRouteEntriesOutput(args: GetRouteEntriesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouteEntriesResult> {
    return pulumi.output(args).apply((a: any) => getRouteEntries(a, opts))
}

/**
 * A collection of arguments for invoking getRouteEntries.
 */
export interface GetRouteEntriesOutputArgs {
    /**
     * The target network segment of the route entry.
     */
    destinationCidrBlock?: pulumi.Input<string>;
    /**
     * The ids of the transit router route entry.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the route entry.
     */
    status?: pulumi.Input<string>;
    /**
     * The name of the route entry.
     */
    transitRouterRouteEntryName?: pulumi.Input<string>;
    /**
     * The id of the route table.
     */
    transitRouterRouteTableId: pulumi.Input<string>;
}
