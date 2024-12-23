// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage direct connect gateway
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.direct_connect.Gateway("foo", {
 *     description: "tf-test",
 *     directConnectGatewayName: "tf-test-gateway",
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * DirectConnectGateway can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:direct_connect/gateway:Gateway default resource_id
 * ```
 */
export class Gateway extends pulumi.CustomResource {
    /**
     * Get an existing Gateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GatewayState, opts?: pulumi.CustomResourceOptions): Gateway {
        return new Gateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:direct_connect/gateway:Gateway';

    /**
     * Returns true if the given object is an instance of Gateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Gateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Gateway.__pulumiType;
    }

    /**
     * The description of direct connect gateway.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of direct connect gateway.
     */
    public readonly directConnectGatewayName!: pulumi.Output<string | undefined>;
    /**
     * The direct connect gateway tags.
     */
    public readonly tags!: pulumi.Output<outputs.direct_connect.GatewayTag[] | undefined>;

    /**
     * Create a Gateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GatewayArgs | GatewayState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GatewayState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["directConnectGatewayName"] = state ? state.directConnectGatewayName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as GatewayArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["directConnectGatewayName"] = args ? args.directConnectGatewayName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Gateway.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Gateway resources.
 */
export interface GatewayState {
    /**
     * The description of direct connect gateway.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of direct connect gateway.
     */
    directConnectGatewayName?: pulumi.Input<string>;
    /**
     * The direct connect gateway tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.direct_connect.GatewayTag>[]>;
}

/**
 * The set of arguments for constructing a Gateway resource.
 */
export interface GatewayArgs {
    /**
     * The description of direct connect gateway.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of direct connect gateway.
     */
    directConnectGatewayName?: pulumi.Input<string>;
    /**
     * The direct connect gateway tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.direct_connect.GatewayTag>[]>;
}
