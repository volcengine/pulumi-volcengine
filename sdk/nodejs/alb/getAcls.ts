// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of alb acls
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.alb.getAcls({
 *     ids: ["acl-1g72w6z11ighs2zbhq4v3rvh4"],
 *     projectName: "default",
 * });
 * ```
 */
export function getAcls(args?: GetAclsArgs, opts?: pulumi.InvokeOptions): Promise<GetAclsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:alb/getAcls:getAcls", {
        "aclName": args.aclName,
        "ids": args.ids,
        "outputFile": args.outputFile,
        "projectName": args.projectName,
    }, opts);
}

/**
 * A collection of arguments for invoking getAcls.
 */
export interface GetAclsArgs {
    /**
     * The name of acl.
     */
    aclName?: string;
    /**
     * A list of Acl IDs.
     */
    ids?: string[];
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The name of project.
     */
    projectName?: string;
}

/**
 * A collection of values returned by getAcls.
 */
export interface GetAclsResult {
    /**
     * The Name of Acl.
     */
    readonly aclName?: string;
    /**
     * The collection of Acl query.
     */
    readonly acls: outputs.alb.GetAclsAcl[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly outputFile?: string;
    /**
     * The project name of Acl.
     */
    readonly projectName?: string;
    /**
     * The total count of Acl query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of alb acls
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.alb.getAcls({
 *     ids: ["acl-1g72w6z11ighs2zbhq4v3rvh4"],
 *     projectName: "default",
 * });
 * ```
 */
export function getAclsOutput(args?: GetAclsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAclsResult> {
    return pulumi.output(args).apply((a: any) => getAcls(a, opts))
}

/**
 * A collection of arguments for invoking getAcls.
 */
export interface GetAclsOutputArgs {
    /**
     * The name of acl.
     */
    aclName?: pulumi.Input<string>;
    /**
     * A list of Acl IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The name of project.
     */
    projectName?: pulumi.Input<string>;
}
