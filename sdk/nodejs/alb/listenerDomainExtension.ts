// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage alb listener domain extension
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooListener = new volcengine.alb.Listener("fooListener", {
 *     loadBalancerId: "alb-1iidd17v3klj474adhfrunyz9",
 *     listenerName: "acc-test-listener-1",
 *     protocol: "HTTPS",
 *     port: 6666,
 *     enabled: "on",
 *     certificateId: "cert-1iidd2pahdyio74adhfr9ajwg",
 *     caCertificateId: "cert-1iidd2r9ii0hs74adhfeodxo1",
 *     serverGroupId: "rsp-1g72w74y4umf42zbhq4k4hnln",
 *     enableHttp2: "on",
 *     enableQuic: "off",
 *     aclStatus: "on",
 *     aclType: "white",
 *     aclIds: ["acl-1g72w6z11ighs2zbhq4v3rvh4"],
 *     description: "acc test listener",
 * });
 * const fooListenerDomainExtension = new volcengine.alb.ListenerDomainExtension("fooListenerDomainExtension", {
 *     listenerId: fooListener.id,
 *     domain: "test-modify.com",
 *     certificateId: "cert-1iidd2pahdyio74adhfr9ajwg",
 * });
 * ```
 *
 * ## Import
 *
 * AlbListenerDomainExtension can be imported using the listener id and domain extension id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:alb/listenerDomainExtension:ListenerDomainExtension default listenerId:extensionId
 * ```
 */
export class ListenerDomainExtension extends pulumi.CustomResource {
    /**
     * Get an existing ListenerDomainExtension resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListenerDomainExtensionState, opts?: pulumi.CustomResourceOptions): ListenerDomainExtension {
        return new ListenerDomainExtension(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:alb/listenerDomainExtension:ListenerDomainExtension';

    /**
     * Returns true if the given object is an instance of ListenerDomainExtension.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ListenerDomainExtension {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ListenerDomainExtension.__pulumiType;
    }

    /**
     * Server certificate used for the domain name.
     */
    public readonly certificateId!: pulumi.Output<string>;
    /**
     * The domain name. The maximum number of associated domain names for an HTTPS listener is 20, with a value range of 1 to 20.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * The id of the domain extension.
     */
    public /*out*/ readonly domainExtensionId!: pulumi.Output<string>;
    /**
     * The listener id. Only HTTPS listener is effective.
     */
    public readonly listenerId!: pulumi.Output<string>;

    /**
     * Create a ListenerDomainExtension resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListenerDomainExtensionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ListenerDomainExtensionArgs | ListenerDomainExtensionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ListenerDomainExtensionState | undefined;
            resourceInputs["certificateId"] = state ? state.certificateId : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["domainExtensionId"] = state ? state.domainExtensionId : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
        } else {
            const args = argsOrState as ListenerDomainExtensionArgs | undefined;
            if ((!args || args.certificateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateId'");
            }
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.listenerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listenerId'");
            }
            resourceInputs["certificateId"] = args ? args.certificateId : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["listenerId"] = args ? args.listenerId : undefined;
            resourceInputs["domainExtensionId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ListenerDomainExtension.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ListenerDomainExtension resources.
 */
export interface ListenerDomainExtensionState {
    /**
     * Server certificate used for the domain name.
     */
    certificateId?: pulumi.Input<string>;
    /**
     * The domain name. The maximum number of associated domain names for an HTTPS listener is 20, with a value range of 1 to 20.
     */
    domain?: pulumi.Input<string>;
    /**
     * The id of the domain extension.
     */
    domainExtensionId?: pulumi.Input<string>;
    /**
     * The listener id. Only HTTPS listener is effective.
     */
    listenerId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ListenerDomainExtension resource.
 */
export interface ListenerDomainExtensionArgs {
    /**
     * Server certificate used for the domain name.
     */
    certificateId: pulumi.Input<string>;
    /**
     * The domain name. The maximum number of associated domain names for an HTTPS listener is 20, with a value range of 1 to 20.
     */
    domain: pulumi.Input<string>;
    /**
     * The listener id. Only HTTPS listener is effective.
     */
    listenerId: pulumi.Input<string>;
}