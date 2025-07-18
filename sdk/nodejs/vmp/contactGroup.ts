// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage vmp contact group
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooContact = new volcengine.vmp.Contact("fooContact", {
 *     email: "acctest1@tftest.com",
 *     webhook: {
 *         address: "https://www.acctest1.com",
 *     },
 *     larkBotWebhook: {
 *         address: "https://www.acctest1.com",
 *     },
 *     dingTalkBotWebhook: {
 *         address: "https://www.dingacctest1.com",
 *         atMobiles: ["18046891812"],
 *     },
 *     phoneNumber: {
 *         countryCode: "+86",
 *         number: "18310101010",
 *     },
 * });
 * const foo1 = new volcengine.vmp.Contact("foo1", {
 *     email: "acctest2@tftest.com",
 *     webhook: {
 *         address: "https://www.acctest2.com",
 *     },
 *     larkBotWebhook: {
 *         address: "https://www.acctest2.com",
 *     },
 *     dingTalkBotWebhook: {
 *         address: "https://www.dingacctest2.com",
 *         atMobiles: ["18046891813"],
 *     },
 *     phoneNumber: {
 *         countryCode: "+86",
 *         number: "18310101011",
 *     },
 * });
 * const fooContactGroup = new volcengine.vmp.ContactGroup("fooContactGroup", {contactIds: [
 *     fooContact.id,
 *     foo1.id,
 * ]});
 * ```
 *
 * ## Import
 *
 * VMP Contact Group can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:vmp/contactGroup:ContactGroup default 60dde3ca-951c-4c05-8777-e5a7caa07ad6
 * ```
 */
export class ContactGroup extends pulumi.CustomResource {
    /**
     * Get an existing ContactGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContactGroupState, opts?: pulumi.CustomResourceOptions): ContactGroup {
        return new ContactGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:vmp/contactGroup:ContactGroup';

    /**
     * Returns true if the given object is an instance of ContactGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContactGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContactGroup.__pulumiType;
    }

    /**
     * A list of contact IDs.
     */
    public readonly contactIds!: pulumi.Output<string[] | undefined>;
    /**
     * The create time of contact group.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The name of the contact group.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a ContactGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ContactGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContactGroupArgs | ContactGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContactGroupState | undefined;
            resourceInputs["contactIds"] = state ? state.contactIds : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ContactGroupArgs | undefined;
            resourceInputs["contactIds"] = args ? args.contactIds : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ContactGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContactGroup resources.
 */
export interface ContactGroupState {
    /**
     * A list of contact IDs.
     */
    contactIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The create time of contact group.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The name of the contact group.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ContactGroup resource.
 */
export interface ContactGroupArgs {
    /**
     * A list of contact IDs.
     */
    contactIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the contact group.
     */
    name?: pulumi.Input<string>;
}
