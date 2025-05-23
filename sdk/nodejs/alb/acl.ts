// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage alb acl
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.alb.Acl("foo", {
 *     aclEntries: [
 *         {
 *             description: "e1",
 *             entry: "172.20.1.0/24",
 *         },
 *         {
 *             description: "e2",
 *             entry: "172.20.3.0/24",
 *         },
 *     ],
 *     aclName: "tf-test-1",
 *     description: "tftest",
 * });
 * ```
 *
 * ## Import
 *
 * Acl can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:alb/acl:Acl default acl-mizl7m1kqccg5smt1bdpijuj
 * ```
 */
export class Acl extends pulumi.CustomResource {
    /**
     * Get an existing Acl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AclState, opts?: pulumi.CustomResourceOptions): Acl {
        return new Acl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:alb/acl:Acl';

    /**
     * Returns true if the given object is an instance of Acl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Acl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Acl.__pulumiType;
    }

    /**
     * The acl entry set of the Acl.
     */
    public readonly aclEntries!: pulumi.Output<outputs.alb.AclAclEntry[] | undefined>;
    /**
     * The name of Acl.
     */
    public readonly aclName!: pulumi.Output<string>;
    /**
     * Create time of Acl.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The description of the Acl.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The project name of the Acl.
     */
    public readonly projectName!: pulumi.Output<string>;

    /**
     * Create a Acl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AclArgs | AclState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AclState | undefined;
            resourceInputs["aclEntries"] = state ? state.aclEntries : undefined;
            resourceInputs["aclName"] = state ? state.aclName : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
        } else {
            const args = argsOrState as AclArgs | undefined;
            resourceInputs["aclEntries"] = args ? args.aclEntries : undefined;
            resourceInputs["aclName"] = args ? args.aclName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Acl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Acl resources.
 */
export interface AclState {
    /**
     * The acl entry set of the Acl.
     */
    aclEntries?: pulumi.Input<pulumi.Input<inputs.alb.AclAclEntry>[]>;
    /**
     * The name of Acl.
     */
    aclName?: pulumi.Input<string>;
    /**
     * Create time of Acl.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The description of the Acl.
     */
    description?: pulumi.Input<string>;
    /**
     * The project name of the Acl.
     */
    projectName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Acl resource.
 */
export interface AclArgs {
    /**
     * The acl entry set of the Acl.
     */
    aclEntries?: pulumi.Input<pulumi.Input<inputs.alb.AclAclEntry>[]>;
    /**
     * The name of Acl.
     */
    aclName?: pulumi.Input<string>;
    /**
     * The description of the Acl.
     */
    description?: pulumi.Input<string>;
    /**
     * The project name of the Acl.
     */
    projectName?: pulumi.Input<string>;
}
