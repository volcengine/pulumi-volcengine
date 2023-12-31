// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of mongodb accounts
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.mongodb.Accounts({
 *     instanceId: "mongo-replica-xxx",
 * });
 * ```
 */
export function accounts(args: AccountsArgs, opts?: pulumi.InvokeOptions): Promise<AccountsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:mongodb/accounts:Accounts", {
        "accountName": args.accountName,
        "instanceId": args.instanceId,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking Accounts.
 */
export interface AccountsArgs {
    /**
     * The name of account, current support only `root`.
     */
    accountName?: string;
    /**
     * Target query mongo instance id.
     */
    instanceId: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by Accounts.
 */
export interface AccountsResult {
    /**
     * The name of account.
     */
    readonly accountName?: string;
    /**
     * The collection of accounts query.
     */
    readonly accounts: outputs.mongodb.AccountsAccount[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly outputFile?: string;
    /**
     * The total count of accounts query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of mongodb accounts
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.mongodb.Accounts({
 *     instanceId: "mongo-replica-xxx",
 * });
 * ```
 */
export function accountsOutput(args: AccountsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<AccountsResult> {
    return pulumi.output(args).apply((a: any) => accounts(a, opts))
}

/**
 * A collection of arguments for invoking Accounts.
 */
export interface AccountsOutputArgs {
    /**
     * The name of account, current support only `root`.
     */
    accountName?: pulumi.Input<string>;
    /**
     * Target query mongo instance id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
