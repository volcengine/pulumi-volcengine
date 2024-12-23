// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage direct connect bgp peer
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.direct_connect.BgpPeer("foo", {
 *     description: "tf-test",
 *     remoteAsn: 2000,
 *     virtualInterfaceId: "dcv-62vi13v131tsn3gd6il****",
 * });
 * ```
 *
 * ## Import
 *
 * DirectConnectBgpPeer can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:direct_connect/bgpPeer:BgpPeer default bgp-2752hz4teko3k7fap8u4c****
 * ```
 */
export class BgpPeer extends pulumi.CustomResource {
    /**
     * Get an existing BgpPeer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BgpPeerState, opts?: pulumi.CustomResourceOptions): BgpPeer {
        return new BgpPeer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:direct_connect/bgpPeer:BgpPeer';

    /**
     * Returns true if the given object is an instance of BgpPeer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BgpPeer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BgpPeer.__pulumiType;
    }

    /**
     * The id of account.
     */
    public /*out*/ readonly accountId!: pulumi.Output<string>;
    /**
     * The auth key of bgp peer.
     */
    public readonly authKey!: pulumi.Output<string | undefined>;
    /**
     * The id of bgp peer.
     */
    public /*out*/ readonly bgpPeerId!: pulumi.Output<string>;
    /**
     * The name of bgp peer.
     */
    public readonly bgpPeerName!: pulumi.Output<string>;
    /**
     * The create time of bgp peer.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The description of bgp peer.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The local asn of bgp peer.
     */
    public /*out*/ readonly localAsn!: pulumi.Output<number>;
    /**
     * The remote asn of bgp peer.
     */
    public readonly remoteAsn!: pulumi.Output<number>;
    /**
     * The session status of bgp peer.
     */
    public /*out*/ readonly sessionStatus!: pulumi.Output<string>;
    /**
     * The status of bgp peer.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The update time of bgp peer.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The id of virtual interface.
     */
    public readonly virtualInterfaceId!: pulumi.Output<string>;

    /**
     * Create a BgpPeer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BgpPeerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BgpPeerArgs | BgpPeerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BgpPeerState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["authKey"] = state ? state.authKey : undefined;
            resourceInputs["bgpPeerId"] = state ? state.bgpPeerId : undefined;
            resourceInputs["bgpPeerName"] = state ? state.bgpPeerName : undefined;
            resourceInputs["creationTime"] = state ? state.creationTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["localAsn"] = state ? state.localAsn : undefined;
            resourceInputs["remoteAsn"] = state ? state.remoteAsn : undefined;
            resourceInputs["sessionStatus"] = state ? state.sessionStatus : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
            resourceInputs["virtualInterfaceId"] = state ? state.virtualInterfaceId : undefined;
        } else {
            const args = argsOrState as BgpPeerArgs | undefined;
            if ((!args || args.remoteAsn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remoteAsn'");
            }
            if ((!args || args.virtualInterfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualInterfaceId'");
            }
            resourceInputs["authKey"] = args ? args.authKey : undefined;
            resourceInputs["bgpPeerName"] = args ? args.bgpPeerName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["remoteAsn"] = args ? args.remoteAsn : undefined;
            resourceInputs["virtualInterfaceId"] = args ? args.virtualInterfaceId : undefined;
            resourceInputs["accountId"] = undefined /*out*/;
            resourceInputs["bgpPeerId"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["localAsn"] = undefined /*out*/;
            resourceInputs["sessionStatus"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BgpPeer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BgpPeer resources.
 */
export interface BgpPeerState {
    /**
     * The id of account.
     */
    accountId?: pulumi.Input<string>;
    /**
     * The auth key of bgp peer.
     */
    authKey?: pulumi.Input<string>;
    /**
     * The id of bgp peer.
     */
    bgpPeerId?: pulumi.Input<string>;
    /**
     * The name of bgp peer.
     */
    bgpPeerName?: pulumi.Input<string>;
    /**
     * The create time of bgp peer.
     */
    creationTime?: pulumi.Input<string>;
    /**
     * The description of bgp peer.
     */
    description?: pulumi.Input<string>;
    /**
     * The local asn of bgp peer.
     */
    localAsn?: pulumi.Input<number>;
    /**
     * The remote asn of bgp peer.
     */
    remoteAsn?: pulumi.Input<number>;
    /**
     * The session status of bgp peer.
     */
    sessionStatus?: pulumi.Input<string>;
    /**
     * The status of bgp peer.
     */
    status?: pulumi.Input<string>;
    /**
     * The update time of bgp peer.
     */
    updateTime?: pulumi.Input<string>;
    /**
     * The id of virtual interface.
     */
    virtualInterfaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BgpPeer resource.
 */
export interface BgpPeerArgs {
    /**
     * The auth key of bgp peer.
     */
    authKey?: pulumi.Input<string>;
    /**
     * The name of bgp peer.
     */
    bgpPeerName?: pulumi.Input<string>;
    /**
     * The description of bgp peer.
     */
    description?: pulumi.Input<string>;
    /**
     * The remote asn of bgp peer.
     */
    remoteAsn: pulumi.Input<number>;
    /**
     * The id of virtual interface.
     */
    virtualInterfaceId: pulumi.Input<string>;
}
