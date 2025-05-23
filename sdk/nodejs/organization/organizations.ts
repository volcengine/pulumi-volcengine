// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of organizations
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.organization.getOrganizations({});
 * ```
 */
/** @deprecated volcengine.organization.Organizations has been deprecated in favor of volcengine.organization.getOrganizations */
export function organizations(args?: OrganizationsArgs, opts?: pulumi.InvokeOptions): Promise<OrganizationsResult> {
    pulumi.log.warn("organizations is deprecated: volcengine.organization.Organizations has been deprecated in favor of volcengine.organization.getOrganizations")
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:organization/organizations:Organizations", {
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking Organizations.
 */
export interface OrganizationsArgs {
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
 * A collection of values returned by Organizations.
 */
export interface OrganizationsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly nameRegex?: string;
    /**
     * The collection of query.
     */
    readonly organizations: outputs.organization.OrganizationsOrganization[];
    readonly outputFile?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of organizations
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.organization.getOrganizations({});
 * ```
 */
/** @deprecated volcengine.organization.Organizations has been deprecated in favor of volcengine.organization.getOrganizations */
export function organizationsOutput(args?: OrganizationsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<OrganizationsResult> {
    return pulumi.output(args).apply((a: any) => organizations(a, opts))
}

/**
 * A collection of arguments for invoking Organizations.
 */
export interface OrganizationsOutputArgs {
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
