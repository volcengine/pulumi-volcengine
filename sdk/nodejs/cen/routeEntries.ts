// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cen route entries
 */
/** @deprecated volcengine.cen.RouteEntries has been deprecated in favor of volcengine.cen.getRouteEntries */
export function routeEntries(args: RouteEntriesArgs, opts?: pulumi.InvokeOptions): Promise<RouteEntriesResult> {
    pulumi.log.warn("routeEntries is deprecated: volcengine.cen.RouteEntries has been deprecated in favor of volcengine.cen.getRouteEntries")

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:cen/routeEntries:RouteEntries", {
        "cenId": args.cenId,
        "destinationCidrBlock": args.destinationCidrBlock,
        "instanceId": args.instanceId,
        "instanceRegionId": args.instanceRegionId,
        "instanceType": args.instanceType,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking RouteEntries.
 */
export interface RouteEntriesArgs {
    /**
     * A cen ID.
     */
    cenId: string;
    /**
     * A destination cidr block.
     */
    destinationCidrBlock?: string;
    /**
     * An instance ID.
     */
    instanceId?: string;
    /**
     * An instance region ID.
     */
    instanceRegionId?: string;
    /**
     * An instance type.
     */
    instanceType?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by RouteEntries.
 */
export interface RouteEntriesResult {
    /**
     * The cen ID of the cen route entry.
     */
    readonly cenId: string;
    /**
     * The collection of cen route entry query.
     */
    readonly cenRouteEntries: outputs.cen.RouteEntriesCenRouteEntry[];
    /**
     * The destination cidr block of the cen route entry.
     */
    readonly destinationCidrBlock?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The instance id of the next hop of the cen route entry.
     */
    readonly instanceId?: string;
    /**
     * The instance region id of the next hop of the cen route entry.
     */
    readonly instanceRegionId?: string;
    /**
     * The instance type of the next hop of the cen route entry.
     */
    readonly instanceType?: string;
    readonly outputFile?: string;
    /**
     * The total count of cen route entry.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of cen route entries
 */
/** @deprecated volcengine.cen.RouteEntries has been deprecated in favor of volcengine.cen.getRouteEntries */
export function routeEntriesOutput(args: RouteEntriesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<RouteEntriesResult> {
    return pulumi.output(args).apply((a: any) => routeEntries(a, opts))
}

/**
 * A collection of arguments for invoking RouteEntries.
 */
export interface RouteEntriesOutputArgs {
    /**
     * A cen ID.
     */
    cenId: pulumi.Input<string>;
    /**
     * A destination cidr block.
     */
    destinationCidrBlock?: pulumi.Input<string>;
    /**
     * An instance ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * An instance region ID.
     */
    instanceRegionId?: pulumi.Input<string>;
    /**
     * An instance type.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
