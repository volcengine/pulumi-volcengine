// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of waf domains
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.waf.getDomains({});
 * ```
 */
export function getDomains(args?: GetDomainsArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:waf/getDomains:getDomains", {
        "accurateQuery": args.accurateQuery,
        "domain": args.domain,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDomains.
 */
export interface GetDomainsArgs {
    /**
     * Matching mode.
     */
    accurateQuery?: number;
    /**
     * The domain name of the protected website that needs to be queried.
     */
    domain?: string;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by getDomains.
 */
export interface GetDomainsResult {
    readonly accurateQuery?: number;
    /**
     * The collection of query.
     */
    readonly datas: outputs.waf.GetDomainsData[];
    /**
     * domain names that need to be protected by WAF.
     */
    readonly domain?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of waf domains
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.waf.getDomains({});
 * ```
 */
export function getDomainsOutput(args?: GetDomainsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainsResult> {
    return pulumi.output(args).apply((a: any) => getDomains(a, opts))
}

/**
 * A collection of arguments for invoking getDomains.
 */
export interface GetDomainsOutputArgs {
    /**
     * Matching mode.
     */
    accurateQuery?: pulumi.Input<number>;
    /**
     * The domain name of the protected website that needs to be queried.
     */
    domain?: pulumi.Input<string>;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
