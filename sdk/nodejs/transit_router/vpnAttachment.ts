// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage transit router vpn attachment
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooTransitRouter = new volcengine.transit_router.TransitRouter("fooTransitRouter", {
 *     transitRouterName: "test-tf-acc",
 *     description: "test-tf-acc",
 * });
 * const fooCustomerGateway = new volcengine.vpn.CustomerGateway("fooCustomerGateway", {
 *     ipAddress: "192.0.1.3",
 *     customerGatewayName: "acc-test",
 *     description: "acc-test",
 * });
 * const fooConnection = new volcengine.vpn.Connection("fooConnection", {
 *     vpnConnectionName: "acc-tf-test",
 *     description: "acc-tf-test",
 *     attachType: "TransitRouter",
 *     customerGatewayId: fooCustomerGateway.id,
 *     localSubnets: ["192.168.0.0/22"],
 *     remoteSubnets: ["192.161.0.0/20"],
 *     dpdAction: "none",
 *     natTraversal: true,
 *     ikeConfigPsk: "acctest@!3",
 *     ikeConfigVersion: "ikev1",
 *     ikeConfigMode: "main",
 *     ikeConfigEncAlg: "aes",
 *     ikeConfigAuthAlg: "md5",
 *     ikeConfigDhGroup: "group2",
 *     ikeConfigLifetime: 9000,
 *     ikeConfigLocalId: "acc_test",
 *     ikeConfigRemoteId: "acc_test",
 *     ipsecConfigEncAlg: "aes",
 *     ipsecConfigAuthAlg: "sha256",
 *     ipsecConfigDhGroup: "group2",
 *     ipsecConfigLifetime: 9000,
 *     logEnabled: false,
 * });
 * const fooVpnAttachment = new volcengine.transit_router.VpnAttachment("fooVpnAttachment", {
 *     zoneId: "cn-beijing-a",
 *     transitRouterAttachmentName: "tf-test-acc",
 *     description: "tf-test-acc-desc",
 *     transitRouterId: fooTransitRouter.id,
 *     vpnConnectionId: fooConnection.id,
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * TransitRouterVpnAttachment can be imported using the transitRouterId:attachmentId, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:transit_router/vpnAttachment:VpnAttachment default tr-2d6fr7mzya2gw58ozfes5g2oh:tr-attach-7qthudw0ll6jmc****
 * ```
 */
export class VpnAttachment extends pulumi.CustomResource {
    /**
     * Get an existing VpnAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpnAttachmentState, opts?: pulumi.CustomResourceOptions): VpnAttachment {
        return new VpnAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:transit_router/vpnAttachment:VpnAttachment';

    /**
     * Returns true if the given object is an instance of VpnAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpnAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpnAttachment.__pulumiType;
    }

    /**
     * The create time.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The description of the transit router vpn attachment.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The status of the transit router.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Tags.
     */
    public readonly tags!: pulumi.Output<outputs.transit_router.VpnAttachmentTag[] | undefined>;
    /**
     * The id of the transit router vpn attachment.
     */
    public /*out*/ readonly transitRouterAttachmentId!: pulumi.Output<string>;
    /**
     * The name of the transit router vpn attachment.
     */
    public readonly transitRouterAttachmentName!: pulumi.Output<string>;
    /**
     * The id of the transit router.
     */
    public readonly transitRouterId!: pulumi.Output<string>;
    /**
     * The update time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The ID of the IPSec connection.
     */
    public readonly vpnConnectionId!: pulumi.Output<string>;
    /**
     * The ID of the availability zone.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a VpnAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpnAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpnAttachmentArgs | VpnAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpnAttachmentState | undefined;
            resourceInputs["creationTime"] = state ? state.creationTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["transitRouterAttachmentId"] = state ? state.transitRouterAttachmentId : undefined;
            resourceInputs["transitRouterAttachmentName"] = state ? state.transitRouterAttachmentName : undefined;
            resourceInputs["transitRouterId"] = state ? state.transitRouterId : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
            resourceInputs["vpnConnectionId"] = state ? state.vpnConnectionId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as VpnAttachmentArgs | undefined;
            if ((!args || args.transitRouterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'transitRouterId'");
            }
            if ((!args || args.vpnConnectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpnConnectionId'");
            }
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["transitRouterAttachmentName"] = args ? args.transitRouterAttachmentName : undefined;
            resourceInputs["transitRouterId"] = args ? args.transitRouterId : undefined;
            resourceInputs["vpnConnectionId"] = args ? args.vpnConnectionId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["transitRouterAttachmentId"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpnAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpnAttachment resources.
 */
export interface VpnAttachmentState {
    /**
     * The create time.
     */
    creationTime?: pulumi.Input<string>;
    /**
     * The description of the transit router vpn attachment.
     */
    description?: pulumi.Input<string>;
    /**
     * The status of the transit router.
     */
    status?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.transit_router.VpnAttachmentTag>[]>;
    /**
     * The id of the transit router vpn attachment.
     */
    transitRouterAttachmentId?: pulumi.Input<string>;
    /**
     * The name of the transit router vpn attachment.
     */
    transitRouterAttachmentName?: pulumi.Input<string>;
    /**
     * The id of the transit router.
     */
    transitRouterId?: pulumi.Input<string>;
    /**
     * The update time.
     */
    updateTime?: pulumi.Input<string>;
    /**
     * The ID of the IPSec connection.
     */
    vpnConnectionId?: pulumi.Input<string>;
    /**
     * The ID of the availability zone.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpnAttachment resource.
 */
export interface VpnAttachmentArgs {
    /**
     * The description of the transit router vpn attachment.
     */
    description?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.transit_router.VpnAttachmentTag>[]>;
    /**
     * The name of the transit router vpn attachment.
     */
    transitRouterAttachmentName?: pulumi.Input<string>;
    /**
     * The id of the transit router.
     */
    transitRouterId: pulumi.Input<string>;
    /**
     * The ID of the IPSec connection.
     */
    vpnConnectionId: pulumi.Input<string>;
    /**
     * The ID of the availability zone.
     */
    zoneId: pulumi.Input<string>;
}
