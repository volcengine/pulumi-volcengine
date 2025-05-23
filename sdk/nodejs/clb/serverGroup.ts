// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage server group
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooZones = volcengine.ecs.getZones({});
 * const fooVpc = new volcengine.vpc.Vpc("fooVpc", {
 *     vpcName: "acc-test-vpc",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const fooSubnet = new volcengine.vpc.Subnet("fooSubnet", {
 *     subnetName: "acc-test-subnet",
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *     vpcId: fooVpc.id,
 * });
 * const fooClb = new volcengine.clb.Clb("fooClb", {
 *     type: "public",
 *     subnetId: fooSubnet.id,
 *     loadBalancerSpec: "small_1",
 *     description: "acc0Demo",
 *     loadBalancerName: "acc-test-create",
 *     eipBillingConfig: {
 *         isp: "BGP",
 *         eipBillingType: "PostPaidByBandwidth",
 *         bandwidth: 1,
 *     },
 * });
 * const fooServerGroup = new volcengine.clb.ServerGroup("fooServerGroup", {
 *     loadBalancerId: fooClb.id,
 *     serverGroupName: "acc-test-create",
 *     description: "hello demo11",
 * });
 * ```
 *
 * ## Import
 *
 * ServerGroup can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:clb/serverGroup:ServerGroup default rsp-273yv0kir1vk07fap8tt9jtwg
 * ```
 */
export class ServerGroup extends pulumi.CustomResource {
    /**
     * Get an existing ServerGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerGroupState, opts?: pulumi.CustomResourceOptions): ServerGroup {
        return new ServerGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:clb/serverGroup:ServerGroup';

    /**
     * Returns true if the given object is an instance of ServerGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerGroup.__pulumiType;
    }

    /**
     * The address ip version of the ServerGroup. Valid values: `ipv4`, `ipv6`. Default is `ipv4`.
     */
    public readonly addressIpVersion!: pulumi.Output<string | undefined>;
    /**
     * The description of ServerGroup.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The ID of the Clb.
     */
    public readonly loadBalancerId!: pulumi.Output<string>;
    /**
     * The ID of the ServerGroup.
     */
    public readonly serverGroupId!: pulumi.Output<string>;
    /**
     * The name of the ServerGroup.
     */
    public readonly serverGroupName!: pulumi.Output<string>;

    /**
     * Create a ServerGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerGroupArgs | ServerGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerGroupState | undefined;
            resourceInputs["addressIpVersion"] = state ? state.addressIpVersion : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["loadBalancerId"] = state ? state.loadBalancerId : undefined;
            resourceInputs["serverGroupId"] = state ? state.serverGroupId : undefined;
            resourceInputs["serverGroupName"] = state ? state.serverGroupName : undefined;
        } else {
            const args = argsOrState as ServerGroupArgs | undefined;
            if ((!args || args.loadBalancerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loadBalancerId'");
            }
            resourceInputs["addressIpVersion"] = args ? args.addressIpVersion : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["loadBalancerId"] = args ? args.loadBalancerId : undefined;
            resourceInputs["serverGroupId"] = args ? args.serverGroupId : undefined;
            resourceInputs["serverGroupName"] = args ? args.serverGroupName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServerGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServerGroup resources.
 */
export interface ServerGroupState {
    /**
     * The address ip version of the ServerGroup. Valid values: `ipv4`, `ipv6`. Default is `ipv4`.
     */
    addressIpVersion?: pulumi.Input<string>;
    /**
     * The description of ServerGroup.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the Clb.
     */
    loadBalancerId?: pulumi.Input<string>;
    /**
     * The ID of the ServerGroup.
     */
    serverGroupId?: pulumi.Input<string>;
    /**
     * The name of the ServerGroup.
     */
    serverGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServerGroup resource.
 */
export interface ServerGroupArgs {
    /**
     * The address ip version of the ServerGroup. Valid values: `ipv4`, `ipv6`. Default is `ipv4`.
     */
    addressIpVersion?: pulumi.Input<string>;
    /**
     * The description of ServerGroup.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the Clb.
     */
    loadBalancerId: pulumi.Input<string>;
    /**
     * The ID of the ServerGroup.
     */
    serverGroupId?: pulumi.Input<string>;
    /**
     * The name of the ServerGroup.
     */
    serverGroupName?: pulumi.Input<string>;
}
