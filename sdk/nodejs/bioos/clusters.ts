// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of bioos clusters
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultClusters = pulumi.output(volcengine.bioos.Clusters());
 * ```
 */
export function clusters(args?: ClustersArgs, opts?: pulumi.InvokeOptions): Promise<ClustersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:bioos/clusters:Clusters", {
        "ids": args.ids,
        "outputFile": args.outputFile,
        "public": args.public,
        "statuses": args.statuses,
        "types": args.types,
    }, opts);
}

/**
 * A collection of arguments for invoking Clusters.
 */
export interface ClustersArgs {
    /**
     * A list of cluster ids.
     */
    ids?: string[];
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * whether it is a public cluster.
     */
    public?: boolean;
    /**
     * The status of the clusters.
     */
    statuses?: string[];
    /**
     * The type of the clusters.
     */
    types?: string[];
}

/**
 * A collection of values returned by Clusters.
 */
export interface ClustersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The list of clusters.
     */
    readonly items: outputs.bioos.ClustersItem[];
    readonly outputFile?: string;
    /**
     * whether it is a public cluster.
     */
    readonly public?: boolean;
    readonly statuses?: string[];
    /**
     * The total count of Cluster query.
     */
    readonly totalCount: number;
    readonly types?: string[];
}

export function clustersOutput(args?: ClustersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ClustersResult> {
    return pulumi.output(args).apply(a => clusters(a, opts))
}

/**
 * A collection of arguments for invoking Clusters.
 */
export interface ClustersOutputArgs {
    /**
     * A list of cluster ids.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * whether it is a public cluster.
     */
    public?: pulumi.Input<boolean>;
    /**
     * The status of the clusters.
     */
    statuses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of the clusters.
     */
    types?: pulumi.Input<pulumi.Input<string>[]>;
}