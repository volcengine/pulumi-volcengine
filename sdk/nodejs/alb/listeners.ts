// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of alb listeners
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.alb.getListeners({});
 * ```
 */
/** @deprecated volcengine.alb.Listeners has been deprecated in favor of volcengine.alb.getListeners */
export function listeners(args?: ListenersArgs, opts?: pulumi.InvokeOptions): Promise<ListenersResult> {
    pulumi.log.warn("listeners is deprecated: volcengine.alb.Listeners has been deprecated in favor of volcengine.alb.getListeners")
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:alb/listeners:Listeners", {
        "ids": args.ids,
        "listenerName": args.listenerName,
        "loadBalancerId": args.loadBalancerId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "projectName": args.projectName,
    }, opts);
}

/**
 * A collection of arguments for invoking Listeners.
 */
export interface ListenersArgs {
    /**
     * A list of Listener IDs.
     */
    ids?: string[];
    /**
     * The name of the Listener.
     */
    listenerName?: string;
    /**
     * The id of the Alb.
     */
    loadBalancerId?: string;
    /**
     * A Name Regex of Listener.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The project name of the listener.
     */
    projectName?: string;
}

/**
 * A collection of values returned by Listeners.
 */
export interface ListenersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The name of the Listener.
     */
    readonly listenerName?: string;
    /**
     * The collection of Listener query.
     */
    readonly listeners: outputs.alb.ListenersListener[];
    /**
     * The load balancer ID that the listener belongs to.
     */
    readonly loadBalancerId?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The project name of the listener.
     */
    readonly projectName?: string;
    /**
     * The total count of Listener query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of alb listeners
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.alb.getListeners({});
 * ```
 */
/** @deprecated volcengine.alb.Listeners has been deprecated in favor of volcengine.alb.getListeners */
export function listenersOutput(args?: ListenersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListenersResult> {
    return pulumi.output(args).apply((a: any) => listeners(a, opts))
}

/**
 * A collection of arguments for invoking Listeners.
 */
export interface ListenersOutputArgs {
    /**
     * A list of Listener IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Listener.
     */
    listenerName?: pulumi.Input<string>;
    /**
     * The id of the Alb.
     */
    loadBalancerId?: pulumi.Input<string>;
    /**
     * A Name Regex of Listener.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The project name of the listener.
     */
    projectName?: pulumi.Input<string>;
}
