// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage rds mysql database
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultDatabase = new volcengine.rds_mysql.Database("default", {
 *     characterSetName: "utf8",
 *     dbName: "xxx",
 *     instanceId: "mysql-xxx",
 * });
 * ```
 *
 * ## Import
 *
 * Database can be imported using the instanceId:dbName, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:rds_mysql/database:Database default mysql-42b38c769c4b:dbname
 * ```
 */
export class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseState, opts?: pulumi.CustomResourceOptions): Database {
        return new Database(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:rds_mysql/database:Database';

    /**
     * Returns true if the given object is an instance of Database.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Database {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Database.__pulumiType;
    }

    /**
     * Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
     */
    public readonly characterSetName!: pulumi.Output<string | undefined>;
    /**
     * Name database.
     * illustrate:
     * Unique name.
     * The length is 2~64 characters.
     * Start with a letter and end with a letter or number.
     * Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
     * Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
     */
    public readonly dbName!: pulumi.Output<string>;
    /**
     * The ID of the RDS instance.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseArgs | DatabaseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseState | undefined;
            resourceInputs["characterSetName"] = state ? state.characterSetName : undefined;
            resourceInputs["dbName"] = state ? state.dbName : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as DatabaseArgs | undefined;
            if ((!args || args.dbName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbName'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["characterSetName"] = args ? args.characterSetName : undefined;
            resourceInputs["dbName"] = args ? args.dbName : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Database.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Database resources.
 */
export interface DatabaseState {
    /**
     * Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
     */
    characterSetName?: pulumi.Input<string>;
    /**
     * Name database.
     * illustrate:
     * Unique name.
     * The length is 2~64 characters.
     * Start with a letter and end with a letter or number.
     * Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
     * Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
     */
    dbName?: pulumi.Input<string>;
    /**
     * The ID of the RDS instance.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
     */
    characterSetName?: pulumi.Input<string>;
    /**
     * Name database.
     * illustrate:
     * Unique name.
     * The length is 2~64 characters.
     * Start with a letter and end with a letter or number.
     * Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
     * Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
     */
    dbName: pulumi.Input<string>;
    /**
     * The ID of the RDS instance.
     */
    instanceId: pulumi.Input<string>;
}