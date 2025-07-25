// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of apig custom domains
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.apig.getCustomDomains({
 *     gatewayId: "gd13d8c6eq1emkiunq6p0",
 *     serviceId: "sd142lm6kiaj519k4l640",
 * });
 * ```
 */
export function getCustomDomains(args?: GetCustomDomainsArgs, opts?: pulumi.InvokeOptions): Promise<GetCustomDomainsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:apig/getCustomDomains:getCustomDomains", {
        "gatewayId": args.gatewayId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "resourceType": args.resourceType,
        "serviceId": args.serviceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getCustomDomains.
 */
export interface GetCustomDomainsArgs {
    /**
     * The id of api gateway.
     */
    gatewayId?: string;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The resource type of domain. Valid values: `Console`, `Ingress`.
     */
    resourceType?: string;
    /**
     * The id of api gateway service.
     */
    serviceId?: string;
}

/**
 * A collection of values returned by getCustomDomains.
 */
export interface GetCustomDomainsResult {
    /**
     * The collection of query.
     */
    readonly customDomains: outputs.apig.GetCustomDomainsCustomDomain[];
    readonly gatewayId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The resource type of domain.
     */
    readonly resourceType?: string;
    /**
     * The id of the api gateway service.
     */
    readonly serviceId?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of apig custom domains
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.apig.getCustomDomains({
 *     gatewayId: "gd13d8c6eq1emkiunq6p0",
 *     serviceId: "sd142lm6kiaj519k4l640",
 * });
 * ```
 */
export function getCustomDomainsOutput(args?: GetCustomDomainsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCustomDomainsResult> {
    return pulumi.output(args).apply((a: any) => getCustomDomains(a, opts))
}

/**
 * A collection of arguments for invoking getCustomDomains.
 */
export interface GetCustomDomainsOutputArgs {
    /**
     * The id of api gateway.
     */
    gatewayId?: pulumi.Input<string>;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The resource type of domain. Valid values: `Console`, `Ingress`.
     */
    resourceType?: pulumi.Input<string>;
    /**
     * The id of api gateway service.
     */
    serviceId?: pulumi.Input<string>;
}
