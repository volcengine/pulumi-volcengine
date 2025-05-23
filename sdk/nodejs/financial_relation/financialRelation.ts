// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage financial relation
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.financial_relation.FinancialRelation("foo", {
 *     accountAlias: "acc-test-financial",
 *     authLists: [
 *         1,
 *         2,
 *         3,
 *     ],
 *     relation: 4,
 *     subAccountId: 2100260000,
 * });
 * ```
 *
 * ## Import
 *
 * FinancialRelation can be imported using the sub_account_id:relation:relation_id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:financial_relation/financialRelation:FinancialRelation default resource_id
 * ```
 */
export class FinancialRelation extends pulumi.CustomResource {
    /**
     * Get an existing FinancialRelation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FinancialRelationState, opts?: pulumi.CustomResourceOptions): FinancialRelation {
        return new FinancialRelation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:financial_relation/financialRelation:FinancialRelation';

    /**
     * Returns true if the given object is an instance of FinancialRelation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FinancialRelation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FinancialRelation.__pulumiType;
    }

    /**
     * The display name of the sub account.
     */
    public readonly accountAlias!: pulumi.Output<string>;
    /**
     * The authorization list of financial management. This field is valid and required when the relation is 4. Valid value range is `1-5`.
     */
    public readonly authLists!: pulumi.Output<number[]>;
    /**
     * The relation of the financial. Valid values: `1`, `4`. `1` means financial custody, `4` means financial management.
     */
    public readonly relation!: pulumi.Output<number>;
    /**
     * The id of the financial relation.
     */
    public /*out*/ readonly relationId!: pulumi.Output<string>;
    /**
     * The status of the financial relation.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;
    /**
     * The sub account id.
     */
    public readonly subAccountId!: pulumi.Output<number>;

    /**
     * Create a FinancialRelation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FinancialRelationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FinancialRelationArgs | FinancialRelationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FinancialRelationState | undefined;
            resourceInputs["accountAlias"] = state ? state.accountAlias : undefined;
            resourceInputs["authLists"] = state ? state.authLists : undefined;
            resourceInputs["relation"] = state ? state.relation : undefined;
            resourceInputs["relationId"] = state ? state.relationId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subAccountId"] = state ? state.subAccountId : undefined;
        } else {
            const args = argsOrState as FinancialRelationArgs | undefined;
            if ((!args || args.subAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subAccountId'");
            }
            resourceInputs["accountAlias"] = args ? args.accountAlias : undefined;
            resourceInputs["authLists"] = args ? args.authLists : undefined;
            resourceInputs["relation"] = args ? args.relation : undefined;
            resourceInputs["subAccountId"] = args ? args.subAccountId : undefined;
            resourceInputs["relationId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FinancialRelation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FinancialRelation resources.
 */
export interface FinancialRelationState {
    /**
     * The display name of the sub account.
     */
    accountAlias?: pulumi.Input<string>;
    /**
     * The authorization list of financial management. This field is valid and required when the relation is 4. Valid value range is `1-5`.
     */
    authLists?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The relation of the financial. Valid values: `1`, `4`. `1` means financial custody, `4` means financial management.
     */
    relation?: pulumi.Input<number>;
    /**
     * The id of the financial relation.
     */
    relationId?: pulumi.Input<string>;
    /**
     * The status of the financial relation.
     */
    status?: pulumi.Input<number>;
    /**
     * The sub account id.
     */
    subAccountId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a FinancialRelation resource.
 */
export interface FinancialRelationArgs {
    /**
     * The display name of the sub account.
     */
    accountAlias?: pulumi.Input<string>;
    /**
     * The authorization list of financial management. This field is valid and required when the relation is 4. Valid value range is `1-5`.
     */
    authLists?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The relation of the financial. Valid values: `1`, `4`. `1` means financial custody, `4` means financial management.
     */
    relation?: pulumi.Input<number>;
    /**
     * The sub account id.
     */
    subAccountId: pulumi.Input<number>;
}
