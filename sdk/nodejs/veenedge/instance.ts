// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage veenedge instance
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.veenedge.Instance("foo", {
 *     areaName: "*****",
 *     cloudserverId: "cloudserver-x92*****jcc8f",
 *     isp: "CMCC",
 * });
 * const foo1 = new volcengine.veenedge.Instance("foo1", {instanceId: "veen0*****0111112"});
 * ```
 *
 * ## Import
 *
 * Instance can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:veenedge/instance:Instance default veenn769ewmjjqyqh5dv
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:veenedge/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * The area name.
     */
    public readonly areaName!: pulumi.Output<string>;
    /**
     * The id of cloud server.
     */
    public readonly cloudserverId!: pulumi.Output<string>;
    /**
     * The name of cluster.
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * The default isp for multi line node.
     */
    public readonly defaultIsp!: pulumi.Output<string | undefined>;
    /**
     * Import an exist instance, usually for import a default instance generated with cloud server creating.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The isp info.
     */
    public readonly isp!: pulumi.Output<string>;
    /**
     * The name of instance, only effected in update scene.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The data of secret, only effected in update scene.
     */
    public readonly secretData!: pulumi.Output<string | undefined>;
    /**
     * The type of secret, only effected in update scene. The value can be `KeyPair` or `Password`.
     */
    public readonly secretType!: pulumi.Output<string | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["areaName"] = state ? state.areaName : undefined;
            resourceInputs["cloudserverId"] = state ? state.cloudserverId : undefined;
            resourceInputs["clusterName"] = state ? state.clusterName : undefined;
            resourceInputs["defaultIsp"] = state ? state.defaultIsp : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["isp"] = state ? state.isp : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["secretData"] = state ? state.secretData : undefined;
            resourceInputs["secretType"] = state ? state.secretType : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            resourceInputs["areaName"] = args ? args.areaName : undefined;
            resourceInputs["cloudserverId"] = args ? args.cloudserverId : undefined;
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["defaultIsp"] = args ? args.defaultIsp : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["isp"] = args ? args.isp : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["secretData"] = args ? args.secretData : undefined;
            resourceInputs["secretType"] = args ? args.secretType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * The area name.
     */
    areaName?: pulumi.Input<string>;
    /**
     * The id of cloud server.
     */
    cloudserverId?: pulumi.Input<string>;
    /**
     * The name of cluster.
     */
    clusterName?: pulumi.Input<string>;
    /**
     * The default isp for multi line node.
     */
    defaultIsp?: pulumi.Input<string>;
    /**
     * Import an exist instance, usually for import a default instance generated with cloud server creating.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The isp info.
     */
    isp?: pulumi.Input<string>;
    /**
     * The name of instance, only effected in update scene.
     */
    name?: pulumi.Input<string>;
    /**
     * The data of secret, only effected in update scene.
     */
    secretData?: pulumi.Input<string>;
    /**
     * The type of secret, only effected in update scene. The value can be `KeyPair` or `Password`.
     */
    secretType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The area name.
     */
    areaName?: pulumi.Input<string>;
    /**
     * The id of cloud server.
     */
    cloudserverId?: pulumi.Input<string>;
    /**
     * The name of cluster.
     */
    clusterName?: pulumi.Input<string>;
    /**
     * The default isp for multi line node.
     */
    defaultIsp?: pulumi.Input<string>;
    /**
     * Import an exist instance, usually for import a default instance generated with cloud server creating.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The isp info.
     */
    isp?: pulumi.Input<string>;
    /**
     * The name of instance, only effected in update scene.
     */
    name?: pulumi.Input<string>;
    /**
     * The data of secret, only effected in update scene.
     */
    secretData?: pulumi.Input<string>;
    /**
     * The type of secret, only effected in update scene. The value can be `KeyPair` or `Password`.
     */
    secretType?: pulumi.Input<string>;
}
