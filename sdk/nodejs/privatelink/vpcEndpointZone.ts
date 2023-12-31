// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage privatelink vpc endpoint zone
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.privatelink.VpcEndpointZone("foo", {
 *     endpointId: "ep-2byz5nlkimc5c2dx0ef9g****",
 *     privateIpAddress: "172.16.0.251",
 *     subnetId: "subnet-2bz47q19zhx4w2dx0eevn****",
 * });
 * ```
 *
 * ## Import
 *
 * VpcEndpointZone can be imported using the endpointId:subnetId, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:privatelink/vpcEndpointZone:VpcEndpointZone default ep-3rel75r081l345zsk2i59****:subnet-2bz47q19zhx4w2dx0eevn****
 * ```
 */
export class VpcEndpointZone extends pulumi.CustomResource {
    /**
     * Get an existing VpcEndpointZone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcEndpointZoneState, opts?: pulumi.CustomResourceOptions): VpcEndpointZone {
        return new VpcEndpointZone(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:privatelink/vpcEndpointZone:VpcEndpointZone';

    /**
     * Returns true if the given object is an instance of VpcEndpointZone.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcEndpointZone {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcEndpointZone.__pulumiType;
    }

    /**
     * The endpoint id of vpc endpoint zone.
     */
    public readonly endpointId!: pulumi.Output<string>;
    /**
     * The network interface id of vpc endpoint.
     */
    public /*out*/ readonly networkInterfaceId!: pulumi.Output<string>;
    /**
     * The private ip address of vpc endpoint zone.
     */
    public readonly privateIpAddress!: pulumi.Output<string>;
    /**
     * The subnet id of vpc endpoint zone.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * The domain of vpc endpoint zone.
     */
    public /*out*/ readonly zoneDomain!: pulumi.Output<string>;
    /**
     * The Id of vpc endpoint zone.
     */
    public /*out*/ readonly zoneId!: pulumi.Output<string>;
    /**
     * The status of vpc endpoint zone.
     */
    public /*out*/ readonly zoneStatus!: pulumi.Output<string>;

    /**
     * Create a VpcEndpointZone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcEndpointZoneArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcEndpointZoneArgs | VpcEndpointZoneState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcEndpointZoneState | undefined;
            resourceInputs["endpointId"] = state ? state.endpointId : undefined;
            resourceInputs["networkInterfaceId"] = state ? state.networkInterfaceId : undefined;
            resourceInputs["privateIpAddress"] = state ? state.privateIpAddress : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["zoneDomain"] = state ? state.zoneDomain : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
            resourceInputs["zoneStatus"] = state ? state.zoneStatus : undefined;
        } else {
            const args = argsOrState as VpcEndpointZoneArgs | undefined;
            if ((!args || args.endpointId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointId'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            resourceInputs["endpointId"] = args ? args.endpointId : undefined;
            resourceInputs["privateIpAddress"] = args ? args.privateIpAddress : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["networkInterfaceId"] = undefined /*out*/;
            resourceInputs["zoneDomain"] = undefined /*out*/;
            resourceInputs["zoneId"] = undefined /*out*/;
            resourceInputs["zoneStatus"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcEndpointZone.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcEndpointZone resources.
 */
export interface VpcEndpointZoneState {
    /**
     * The endpoint id of vpc endpoint zone.
     */
    endpointId?: pulumi.Input<string>;
    /**
     * The network interface id of vpc endpoint.
     */
    networkInterfaceId?: pulumi.Input<string>;
    /**
     * The private ip address of vpc endpoint zone.
     */
    privateIpAddress?: pulumi.Input<string>;
    /**
     * The subnet id of vpc endpoint zone.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The domain of vpc endpoint zone.
     */
    zoneDomain?: pulumi.Input<string>;
    /**
     * The Id of vpc endpoint zone.
     */
    zoneId?: pulumi.Input<string>;
    /**
     * The status of vpc endpoint zone.
     */
    zoneStatus?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcEndpointZone resource.
 */
export interface VpcEndpointZoneArgs {
    /**
     * The endpoint id of vpc endpoint zone.
     */
    endpointId: pulumi.Input<string>;
    /**
     * The private ip address of vpc endpoint zone.
     */
    privateIpAddress?: pulumi.Input<string>;
    /**
     * The subnet id of vpc endpoint zone.
     */
    subnetId: pulumi.Input<string>;
}
