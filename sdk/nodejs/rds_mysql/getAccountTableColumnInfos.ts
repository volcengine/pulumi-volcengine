// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of rds mysql account table column infos
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.rds_mysql.getAccountTableColumnInfos({
 *     dbName: "ddd",
 *     instanceId: "mysql-b51d37110dd1",
 * });
 * ```
 */
export function getAccountTableColumnInfos(args: GetAccountTableColumnInfosArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountTableColumnInfosResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:rds_mysql/getAccountTableColumnInfos:getAccountTableColumnInfos", {
        "accountName": args.accountName,
        "columnName": args.columnName,
        "dbName": args.dbName,
        "host": args.host,
        "instanceId": args.instanceId,
        "outputFile": args.outputFile,
        "tableLimit": args.tableLimit,
        "tableName": args.tableName,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccountTableColumnInfos.
 */
export interface GetAccountTableColumnInfosArgs {
    /**
     * The name of the account.
     */
    accountName?: string;
    /**
     * The name of the column.
     */
    columnName?: string;
    /**
     * The name of the database.
     */
    dbName: string;
    /**
     * Specify the IP address for the account to access the database. The default value is %.
     */
    host?: string;
    /**
     * The id of the mysql instance.
     */
    instanceId: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * Specify the number of tables in the table column permission information to be returned. If it exceeds the setting, it will be truncated.
     */
    tableLimit?: number;
    /**
     * The name of the table.
     */
    tableName?: string;
}

/**
 * A collection of values returned by getAccountTableColumnInfos.
 */
export interface GetAccountTableColumnInfosResult {
    readonly accountName?: string;
    /**
     * The name of the column.
     */
    readonly columnName?: string;
    readonly dbName: string;
    readonly host?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly outputFile?: string;
    /**
     * The collection of query.
     */
    readonly tableInfos: outputs.rds_mysql.GetAccountTableColumnInfosTableInfo[];
    readonly tableLimit?: number;
    /**
     * The name of the table.
     */
    readonly tableName?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of rds mysql account table column infos
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.rds_mysql.getAccountTableColumnInfos({
 *     dbName: "ddd",
 *     instanceId: "mysql-b51d37110dd1",
 * });
 * ```
 */
export function getAccountTableColumnInfosOutput(args: GetAccountTableColumnInfosOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccountTableColumnInfosResult> {
    return pulumi.output(args).apply((a: any) => getAccountTableColumnInfos(a, opts))
}

/**
 * A collection of arguments for invoking getAccountTableColumnInfos.
 */
export interface GetAccountTableColumnInfosOutputArgs {
    /**
     * The name of the account.
     */
    accountName?: pulumi.Input<string>;
    /**
     * The name of the column.
     */
    columnName?: pulumi.Input<string>;
    /**
     * The name of the database.
     */
    dbName: pulumi.Input<string>;
    /**
     * Specify the IP address for the account to access the database. The default value is %.
     */
    host?: pulumi.Input<string>;
    /**
     * The id of the mysql instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * Specify the number of tables in the table column permission information to be returned. If it exceeds the setting, it will be truncated.
     */
    tableLimit?: pulumi.Input<number>;
    /**
     * The name of the table.
     */
    tableName?: pulumi.Input<string>;
}
