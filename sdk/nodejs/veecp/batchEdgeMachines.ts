// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of veecp batch edge machines
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooBatchEdgeMachine = new volcengine.veecp.BatchEdgeMachine("fooBatchEdgeMachine", {
 *     clusterId: "ccvd7mte6t101fno98u60",
 *     nodePoolId: "pcvd90uacnsr73g6bjic0",
 *     ttlHours: 1,
 * });
 * const fooBatchEdgeMachines = volcengine.veecp.getBatchEdgeMachinesOutput({
 *     clusterIds: [fooBatchEdgeMachine.clusterId],
 *     ids: [fooBatchEdgeMachine.id],
 * });
 * ```
 */
/** @deprecated volcengine.veecp.BatchEdgeMachines has been deprecated in favor of volcengine.veecp.getBatchEdgeMachines */
export function batchEdgeMachines(args?: BatchEdgeMachinesArgs, opts?: pulumi.InvokeOptions): Promise<BatchEdgeMachinesResult> {
    pulumi.log.warn("batchEdgeMachines is deprecated: volcengine.veecp.BatchEdgeMachines has been deprecated in favor of volcengine.veecp.getBatchEdgeMachines")
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:veecp/batchEdgeMachines:BatchEdgeMachines", {
        "clusterIds": args.clusterIds,
        "createClientToken": args.createClientToken,
        "ids": args.ids,
        "ips": args.ips,
        "name": args.name,
        "needBootstrapScript": args.needBootstrapScript,
        "outputFile": args.outputFile,
        "statuses": args.statuses,
        "zoneIds": args.zoneIds,
    }, opts);
}

/**
 * A collection of arguments for invoking BatchEdgeMachines.
 */
export interface BatchEdgeMachinesArgs {
    /**
     * The ClusterIds of NodePool IDs.
     */
    clusterIds?: string[];
    /**
     * The ClientToken when successfully created.
     */
    createClientToken?: string;
    /**
     * A list of IDs.
     */
    ids?: string[];
    /**
     * The IPs.
     */
    ips?: string[];
    /**
     * The Name of NodePool.
     */
    name?: string;
    /**
     * Whether it is necessary to query the node management script.
     */
    needBootstrapScript?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The Status of NodePool.
     */
    statuses?: inputs.veecp.BatchEdgeMachinesStatus[];
    /**
     * The Zone Ids.
     */
    zoneIds?: string[];
}

/**
 * A collection of values returned by BatchEdgeMachines.
 */
export interface BatchEdgeMachinesResult {
    readonly clusterIds?: string[];
    /**
     * The ClientToken when successfully created.
     */
    readonly createClientToken?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly ips?: string[];
    /**
     * The collection of query.
     */
    readonly machines: outputs.veecp.BatchEdgeMachinesMachine[];
    /**
     * The Name of NodePool.
     */
    readonly name?: string;
    readonly needBootstrapScript?: string;
    readonly outputFile?: string;
    readonly statuses?: outputs.veecp.BatchEdgeMachinesStatus[];
    /**
     * The total count of query.
     */
    readonly totalCount: number;
    readonly zoneIds?: string[];
}
/**
 * Use this data source to query detailed information of veecp batch edge machines
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooBatchEdgeMachine = new volcengine.veecp.BatchEdgeMachine("fooBatchEdgeMachine", {
 *     clusterId: "ccvd7mte6t101fno98u60",
 *     nodePoolId: "pcvd90uacnsr73g6bjic0",
 *     ttlHours: 1,
 * });
 * const fooBatchEdgeMachines = volcengine.veecp.getBatchEdgeMachinesOutput({
 *     clusterIds: [fooBatchEdgeMachine.clusterId],
 *     ids: [fooBatchEdgeMachine.id],
 * });
 * ```
 */
/** @deprecated volcengine.veecp.BatchEdgeMachines has been deprecated in favor of volcengine.veecp.getBatchEdgeMachines */
export function batchEdgeMachinesOutput(args?: BatchEdgeMachinesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<BatchEdgeMachinesResult> {
    return pulumi.output(args).apply((a: any) => batchEdgeMachines(a, opts))
}

/**
 * A collection of arguments for invoking BatchEdgeMachines.
 */
export interface BatchEdgeMachinesOutputArgs {
    /**
     * The ClusterIds of NodePool IDs.
     */
    clusterIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ClientToken when successfully created.
     */
    createClientToken?: pulumi.Input<string>;
    /**
     * A list of IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The IPs.
     */
    ips?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Name of NodePool.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether it is necessary to query the node management script.
     */
    needBootstrapScript?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The Status of NodePool.
     */
    statuses?: pulumi.Input<pulumi.Input<inputs.veecp.BatchEdgeMachinesStatusArgs>[]>;
    /**
     * The Zone Ids.
     */
    zoneIds?: pulumi.Input<pulumi.Input<string>[]>;
}
