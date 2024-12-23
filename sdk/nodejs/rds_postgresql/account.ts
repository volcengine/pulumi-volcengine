// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage rds postgresql account
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.rds_postgresql.Account("foo", {
 *     accountName: "acc-test-account",
 *     accountPassword: "93c@*****!ab12",
 *     accountType: "Super",
 *     instanceId: "postgres-954*****7233",
 * });
 * const foo1 = new volcengine.rds_postgresql.Account("foo1", {
 *     accountName: "acc-test-account1",
 *     accountPassword: "9wc@****b12",
 *     accountPrivileges: "Inherit,Login,CreateRole,CreateDB",
 *     accountType: "Normal",
 *     instanceId: "postgres-95*****7233",
 * });
 * ```
 *
 * ## Import
 *
 * RDS postgresql account can be imported using the instance_id:account_name, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:rds_postgresql/account:Account default postgres-ca7b7019****:accountName
 * ```
 */
export class Account extends pulumi.CustomResource {
    /**
     * Get an existing Account resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountState, opts?: pulumi.CustomResourceOptions): Account {
        return new Account(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:rds_postgresql/account:Account';

    /**
     * Returns true if the given object is an instance of Account.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Account {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Account.__pulumiType;
    }

    /**
     * Database account name.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * The password of the database account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly accountPassword!: pulumi.Output<string>;
    /**
     * The privilege information of account. When the account type is a super account, there is no need to pass in this parameter, and all privileges are supported by default. When the account type is a normal account, this parameter can be passed in, the default values are Login and Inherit.
     */
    public readonly accountPrivileges!: pulumi.Output<string>;
    /**
     * The status of the database account.
     */
    public /*out*/ readonly accountStatus!: pulumi.Output<string>;
    /**
     * Database account type, value:
     * Super: A high-privilege account. Only one database account can be created for an instance.
     * Normal: An account with ordinary privileges.
     */
    public readonly accountType!: pulumi.Output<string>;
    /**
     * The ID of the RDS instance.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a Account resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountArgs | AccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccountState | undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["accountPassword"] = state ? state.accountPassword : undefined;
            resourceInputs["accountPrivileges"] = state ? state.accountPrivileges : undefined;
            resourceInputs["accountStatus"] = state ? state.accountStatus : undefined;
            resourceInputs["accountType"] = state ? state.accountType : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as AccountArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.accountPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountPassword'");
            }
            if ((!args || args.accountType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountType'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["accountPassword"] = args?.accountPassword ? pulumi.secret(args.accountPassword) : undefined;
            resourceInputs["accountPrivileges"] = args ? args.accountPrivileges : undefined;
            resourceInputs["accountType"] = args ? args.accountType : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["accountStatus"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accountPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Account.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Account resources.
 */
export interface AccountState {
    /**
     * Database account name.
     */
    accountName?: pulumi.Input<string>;
    /**
     * The password of the database account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    accountPassword?: pulumi.Input<string>;
    /**
     * The privilege information of account. When the account type is a super account, there is no need to pass in this parameter, and all privileges are supported by default. When the account type is a normal account, this parameter can be passed in, the default values are Login and Inherit.
     */
    accountPrivileges?: pulumi.Input<string>;
    /**
     * The status of the database account.
     */
    accountStatus?: pulumi.Input<string>;
    /**
     * Database account type, value:
     * Super: A high-privilege account. Only one database account can be created for an instance.
     * Normal: An account with ordinary privileges.
     */
    accountType?: pulumi.Input<string>;
    /**
     * The ID of the RDS instance.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Account resource.
 */
export interface AccountArgs {
    /**
     * Database account name.
     */
    accountName: pulumi.Input<string>;
    /**
     * The password of the database account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    accountPassword: pulumi.Input<string>;
    /**
     * The privilege information of account. When the account type is a super account, there is no need to pass in this parameter, and all privileges are supported by default. When the account type is a normal account, this parameter can be passed in, the default values are Login and Inherit.
     */
    accountPrivileges?: pulumi.Input<string>;
    /**
     * Database account type, value:
     * Super: A high-privilege account. Only one database account can be created for an instance.
     * Normal: An account with ordinary privileges.
     */
    accountType: pulumi.Input<string>;
    /**
     * The ID of the RDS instance.
     */
    instanceId: pulumi.Input<string>;
}
