// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/** @deprecated volcengine.redis.PitrTimeWindows has been deprecated in favor of volcengine.redis.getPitrTimeWindows */
export function pitrTimeWindows(args: PitrTimeWindowsArgs, opts?: pulumi.InvokeOptions): Promise<PitrTimeWindowsResult> {
    pulumi.log.warn("pitrTimeWindows is deprecated: volcengine.redis.PitrTimeWindows has been deprecated in favor of volcengine.redis.getPitrTimeWindows")

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:redis/pitrTimeWindows:PitrTimeWindows", {
        "ids": args.ids,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking PitrTimeWindows.
 */
export interface PitrTimeWindowsArgs {
    ids: string[];
    outputFile?: string;
}

/**
 * A collection of values returned by PitrTimeWindows.
 */
export interface PitrTimeWindowsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly outputFile?: string;
    readonly periods: outputs.redis.PitrTimeWindowsPeriod[];
    readonly totalCount: number;
}
/** @deprecated volcengine.redis.PitrTimeWindows has been deprecated in favor of volcengine.redis.getPitrTimeWindows */
export function pitrTimeWindowsOutput(args: PitrTimeWindowsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<PitrTimeWindowsResult> {
    return pulumi.output(args).apply((a: any) => pitrTimeWindows(a, opts))
}

/**
 * A collection of arguments for invoking PitrTimeWindows.
 */
export interface PitrTimeWindowsOutputArgs {
    ids: pulumi.Input<pulumi.Input<string>[]>;
    outputFile?: pulumi.Input<string>;
}
