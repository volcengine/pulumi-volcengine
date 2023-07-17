// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage redis account
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = new volcengine.redis.Account("foo", {
 *     accountName: "test",
 *     description: "test12345",
 *     instanceId: "redis-cn0398aizj8cwmopx",
 *     password: "1qaz!QAZ",
 *     roleName: "ReadOnly",
 * });
 * ```
 *
 * ## Import
 *
 * Redis account can be imported using the instanceId:accountName, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:redis/account:Account default redis-42b38c769c4b:test
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
    public static readonly __pulumiType = 'volcengine:redis/account:Account';

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
     * Redis account name.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * The description of the redis account.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Redis instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The password of the redis account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Role type, the valid value can be `Administrator`, `ReadWrite`, `ReadOnly`, `NotDangerous`.
     */
    public readonly roleName!: pulumi.Output<string>;

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
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["roleName"] = state ? state.roleName : undefined;
        } else {
            const args = argsOrState as AccountArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.roleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleName'");
            }
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["roleName"] = args ? args.roleName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Account.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Account resources.
 */
export interface AccountState {
    /**
     * Redis account name.
     */
    accountName?: pulumi.Input<string>;
    /**
     * The description of the redis account.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the Redis instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The password of the redis account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    password?: pulumi.Input<string>;
    /**
     * Role type, the valid value can be `Administrator`, `ReadWrite`, `ReadOnly`, `NotDangerous`.
     */
    roleName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Account resource.
 */
export interface AccountArgs {
    /**
     * Redis account name.
     */
    accountName: pulumi.Input<string>;
    /**
     * The description of the redis account.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the Redis instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The password of the redis account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    password: pulumi.Input<string>;
    /**
     * Role type, the valid value can be `Administrator`, `ReadWrite`, `ReadOnly`, `NotDangerous`.
     */
    roleName: pulumi.Input<string>;
}