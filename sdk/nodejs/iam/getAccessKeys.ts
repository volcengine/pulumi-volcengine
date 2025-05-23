// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of iam access keys
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.iam.getAccessKeys({});
 * ```
 */
export function getAccessKeys(args?: GetAccessKeysArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessKeysResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:iam/getAccessKeys:getAccessKeys", {
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "userName": args.userName,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessKeys.
 */
export interface GetAccessKeysArgs {
    /**
     * A Name Regex of IAM.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The user names.
     */
    userName?: string;
}

/**
 * A collection of values returned by getAccessKeys.
 */
export interface GetAccessKeysResult {
    /**
     * The collection of access keys.
     */
    readonly accessKeyMetadatas: outputs.iam.GetAccessKeysAccessKeyMetadata[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The total count of user query.
     */
    readonly totalCount: number;
    /**
     * The user name.
     */
    readonly userName?: string;
}
/**
 * Use this data source to query detailed information of iam access keys
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.iam.getAccessKeys({});
 * ```
 */
export function getAccessKeysOutput(args?: GetAccessKeysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccessKeysResult> {
    return pulumi.output(args).apply((a: any) => getAccessKeys(a, opts))
}

/**
 * A collection of arguments for invoking getAccessKeys.
 */
export interface GetAccessKeysOutputArgs {
    /**
     * A Name Regex of IAM.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The user names.
     */
    userName?: pulumi.Input<string>;
}
