// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage waf ip group
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.waf.IpGroup("foo", {
 *     addType: "List",
 *     ipLists: [
 *         "1.1.1.1",
 *         "1.1.1.2",
 *         "1.1.1.3",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * WafIpGroup can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:waf/ipGroup:IpGroup default resource_id
 * ```
 */
export class IpGroup extends pulumi.CustomResource {
    /**
     * Get an existing IpGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpGroupState, opts?: pulumi.CustomResourceOptions): IpGroup {
        return new IpGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:waf/ipGroup:IpGroup';

    /**
     * Returns true if the given object is an instance of IpGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpGroup.__pulumiType;
    }

    /**
     * The way of addition.
     */
    public readonly addType!: pulumi.Output<string>;
    /**
     * The number of IP addresses within the address group.
     */
    public /*out*/ readonly ipCount!: pulumi.Output<number>;
    /**
     * The ID of the ip group.
     */
    public /*out*/ readonly ipGroupId!: pulumi.Output<number>;
    /**
     * The IP address to be added.
     */
    public readonly ipLists!: pulumi.Output<string[]>;
    /**
     * The name of ip group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The list of associated rules.
     */
    public /*out*/ readonly relatedRules!: pulumi.Output<outputs.waf.IpGroupRelatedRule[]>;
    /**
     * ip group update time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a IpGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpGroupArgs | IpGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpGroupState | undefined;
            resourceInputs["addType"] = state ? state.addType : undefined;
            resourceInputs["ipCount"] = state ? state.ipCount : undefined;
            resourceInputs["ipGroupId"] = state ? state.ipGroupId : undefined;
            resourceInputs["ipLists"] = state ? state.ipLists : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["relatedRules"] = state ? state.relatedRules : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as IpGroupArgs | undefined;
            if ((!args || args.addType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addType'");
            }
            if ((!args || args.ipLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipLists'");
            }
            resourceInputs["addType"] = args ? args.addType : undefined;
            resourceInputs["ipLists"] = args ? args.ipLists : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ipCount"] = undefined /*out*/;
            resourceInputs["ipGroupId"] = undefined /*out*/;
            resourceInputs["relatedRules"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IpGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpGroup resources.
 */
export interface IpGroupState {
    /**
     * The way of addition.
     */
    addType?: pulumi.Input<string>;
    /**
     * The number of IP addresses within the address group.
     */
    ipCount?: pulumi.Input<number>;
    /**
     * The ID of the ip group.
     */
    ipGroupId?: pulumi.Input<number>;
    /**
     * The IP address to be added.
     */
    ipLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of ip group.
     */
    name?: pulumi.Input<string>;
    /**
     * The list of associated rules.
     */
    relatedRules?: pulumi.Input<pulumi.Input<inputs.waf.IpGroupRelatedRule>[]>;
    /**
     * ip group update time.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IpGroup resource.
 */
export interface IpGroupArgs {
    /**
     * The way of addition.
     */
    addType: pulumi.Input<string>;
    /**
     * The IP address to be added.
     */
    ipLists: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of ip group.
     */
    name?: pulumi.Input<string>;
}
