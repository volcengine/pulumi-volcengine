// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage cloud monitor contact
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const _default = new volcengine.cloud_monitor.Contact("default", {
 *     email: "192*****72@****.com",
 *     phone: "180****27812",
 * });
 * ```
 *
 * ## Import
 *
 * CloudMonitor Contact can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:cloud_monitor/contact:Contact default 145258255725730****
 * ```
 */
export class Contact extends pulumi.CustomResource {
    /**
     * Get an existing Contact resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContactState, opts?: pulumi.CustomResourceOptions): Contact {
        return new Contact(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:cloud_monitor/contact:Contact';

    /**
     * Returns true if the given object is an instance of Contact.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Contact {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Contact.__pulumiType;
    }

    /**
     * The email of contact.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * The name of contact.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The phone of contact.
     */
    public readonly phone!: pulumi.Output<string | undefined>;

    /**
     * Create a Contact resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContactArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContactArgs | ContactState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContactState | undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["phone"] = state ? state.phone : undefined;
        } else {
            const args = argsOrState as ContactArgs | undefined;
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["phone"] = args ? args.phone : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Contact.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Contact resources.
 */
export interface ContactState {
    /**
     * The email of contact.
     */
    email?: pulumi.Input<string>;
    /**
     * The name of contact.
     */
    name?: pulumi.Input<string>;
    /**
     * The phone of contact.
     */
    phone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Contact resource.
 */
export interface ContactArgs {
    /**
     * The email of contact.
     */
    email: pulumi.Input<string>;
    /**
     * The name of contact.
     */
    name?: pulumi.Input<string>;
    /**
     * The phone of contact.
     */
    phone?: pulumi.Input<string>;
}
