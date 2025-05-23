// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cloud identity permission set provisionings
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.cloud_identity.getPermissionSetProvisionings({
 *     targetId: "210026****",
 * });
 * ```
 */
export function getPermissionSetProvisionings(args?: GetPermissionSetProvisioningsArgs, opts?: pulumi.InvokeOptions): Promise<GetPermissionSetProvisioningsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:cloud_identity/getPermissionSetProvisionings:getPermissionSetProvisionings", {
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "permissionSetId": args.permissionSetId,
        "targetId": args.targetId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPermissionSetProvisionings.
 */
export interface GetPermissionSetProvisioningsArgs {
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The id of cloud identity permission set.
     */
    permissionSetId?: string;
    /**
     * The target account id of cloud identity permission set.
     */
    targetId?: string;
}

/**
 * A collection of values returned by getPermissionSetProvisionings.
 */
export interface GetPermissionSetProvisioningsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The collection of query.
     */
    readonly permissionProvisionings: outputs.cloud_identity.GetPermissionSetProvisioningsPermissionProvisioning[];
    /**
     * The id of the cloud identity permission set.
     */
    readonly permissionSetId?: string;
    /**
     * The target account id of the cloud identity permission set provisioning.
     */
    readonly targetId?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of cloud identity permission set provisionings
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.cloud_identity.getPermissionSetProvisionings({
 *     targetId: "210026****",
 * });
 * ```
 */
export function getPermissionSetProvisioningsOutput(args?: GetPermissionSetProvisioningsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPermissionSetProvisioningsResult> {
    return pulumi.output(args).apply((a: any) => getPermissionSetProvisionings(a, opts))
}

/**
 * A collection of arguments for invoking getPermissionSetProvisionings.
 */
export interface GetPermissionSetProvisioningsOutputArgs {
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The id of cloud identity permission set.
     */
    permissionSetId?: pulumi.Input<string>;
    /**
     * The target account id of cloud identity permission set.
     */
    targetId?: pulumi.Input<string>;
}
