// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage redis allow list
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.redis.AllowList("foo", {
 *     allowLists: [
 *         "0.0.0.0/0",
 *         "192.168.0.0/24",
 *         "192.168.1.1",
 *         "192.168.2.22",
 *     ],
 *     allowListDesc: "acctftestallowlist",
 *     allowListName: "acc_test_tf_allowlist_create",
 * });
 * ```
 *
 * ## Import
 *
 * Redis AllowList can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:redis/allowList:AllowList default acl-cn03wk541s55c376xxxx
 * ```
 */
export class AllowList extends pulumi.CustomResource {
    /**
     * Get an existing AllowList resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AllowListState, opts?: pulumi.CustomResourceOptions): AllowList {
        return new AllowList(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:redis/allowList:AllowList';

    /**
     * Returns true if the given object is an instance of AllowList.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AllowList {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AllowList.__pulumiType;
    }

    /**
     * The type of the whitelist.
     */
    public /*out*/ readonly allowListCategory!: pulumi.Output<string>;
    /**
     * Description of allow list.
     */
    public readonly allowListDesc!: pulumi.Output<string | undefined>;
    /**
     * Id of allow list.
     */
    public /*out*/ readonly allowListId!: pulumi.Output<string>;
    /**
     * The IP number of allow list.
     */
    public /*out*/ readonly allowListIpNum!: pulumi.Output<number>;
    /**
     * Name of allow list.
     */
    public readonly allowListName!: pulumi.Output<string>;
    /**
     * Type of allow list.
     */
    public /*out*/ readonly allowListType!: pulumi.Output<string>;
    /**
     * Ip list of allow list.
     */
    public readonly allowLists!: pulumi.Output<string[]>;
    /**
     * The number of instance that associated to allow list.
     */
    public /*out*/ readonly associatedInstanceNum!: pulumi.Output<number>;
    /**
     * Instances associated by this allow list.
     */
    public /*out*/ readonly associatedInstances!: pulumi.Output<outputs.redis.AllowListAssociatedInstance[]>;
    /**
     * The name of the project to which the white list belongs.
     */
    public /*out*/ readonly projectName!: pulumi.Output<string>;
    /**
     * The current whitelist is the list of security group information that has been associated.
     */
    public /*out*/ readonly securityGroupBindInfos!: pulumi.Output<outputs.redis.AllowListSecurityGroupBindInfo[]>;

    /**
     * Create a AllowList resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AllowListArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AllowListArgs | AllowListState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AllowListState | undefined;
            resourceInputs["allowListCategory"] = state ? state.allowListCategory : undefined;
            resourceInputs["allowListDesc"] = state ? state.allowListDesc : undefined;
            resourceInputs["allowListId"] = state ? state.allowListId : undefined;
            resourceInputs["allowListIpNum"] = state ? state.allowListIpNum : undefined;
            resourceInputs["allowListName"] = state ? state.allowListName : undefined;
            resourceInputs["allowListType"] = state ? state.allowListType : undefined;
            resourceInputs["allowLists"] = state ? state.allowLists : undefined;
            resourceInputs["associatedInstanceNum"] = state ? state.associatedInstanceNum : undefined;
            resourceInputs["associatedInstances"] = state ? state.associatedInstances : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["securityGroupBindInfos"] = state ? state.securityGroupBindInfos : undefined;
        } else {
            const args = argsOrState as AllowListArgs | undefined;
            if ((!args || args.allowListName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allowListName'");
            }
            if ((!args || args.allowLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allowLists'");
            }
            resourceInputs["allowListDesc"] = args ? args.allowListDesc : undefined;
            resourceInputs["allowListName"] = args ? args.allowListName : undefined;
            resourceInputs["allowLists"] = args ? args.allowLists : undefined;
            resourceInputs["allowListCategory"] = undefined /*out*/;
            resourceInputs["allowListId"] = undefined /*out*/;
            resourceInputs["allowListIpNum"] = undefined /*out*/;
            resourceInputs["allowListType"] = undefined /*out*/;
            resourceInputs["associatedInstanceNum"] = undefined /*out*/;
            resourceInputs["associatedInstances"] = undefined /*out*/;
            resourceInputs["projectName"] = undefined /*out*/;
            resourceInputs["securityGroupBindInfos"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AllowList.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AllowList resources.
 */
export interface AllowListState {
    /**
     * The type of the whitelist.
     */
    allowListCategory?: pulumi.Input<string>;
    /**
     * Description of allow list.
     */
    allowListDesc?: pulumi.Input<string>;
    /**
     * Id of allow list.
     */
    allowListId?: pulumi.Input<string>;
    /**
     * The IP number of allow list.
     */
    allowListIpNum?: pulumi.Input<number>;
    /**
     * Name of allow list.
     */
    allowListName?: pulumi.Input<string>;
    /**
     * Type of allow list.
     */
    allowListType?: pulumi.Input<string>;
    /**
     * Ip list of allow list.
     */
    allowLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The number of instance that associated to allow list.
     */
    associatedInstanceNum?: pulumi.Input<number>;
    /**
     * Instances associated by this allow list.
     */
    associatedInstances?: pulumi.Input<pulumi.Input<inputs.redis.AllowListAssociatedInstance>[]>;
    /**
     * The name of the project to which the white list belongs.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The current whitelist is the list of security group information that has been associated.
     */
    securityGroupBindInfos?: pulumi.Input<pulumi.Input<inputs.redis.AllowListSecurityGroupBindInfo>[]>;
}

/**
 * The set of arguments for constructing a AllowList resource.
 */
export interface AllowListArgs {
    /**
     * Description of allow list.
     */
    allowListDesc?: pulumi.Input<string>;
    /**
     * Name of allow list.
     */
    allowListName: pulumi.Input<string>;
    /**
     * Ip list of allow list.
     */
    allowLists: pulumi.Input<pulumi.Input<string>[]>;
}
