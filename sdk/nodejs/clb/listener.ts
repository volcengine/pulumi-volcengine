// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage listener
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = new volcengine.clb.Listener("foo", {
 *     enabled: "on",
 *     healthCheck: {
 *         domain: "volcengine.com",
 *         enabled: "on",
 *         healthyThreshold: 5,
 *         httpCode: "http_2xx",
 *         interval: 10,
 *         method: "GET",
 *         timeout: 3,
 *         unHealthyThreshold: 2,
 *         uri: "/",
 *     },
 *     listenerName: "Demo-HTTP-90",
 *     loadBalancerId: "clb-274xltt3rfmyo7fap8sv1jq39",
 *     port: 90,
 *     protocol: "HTTP",
 *     serverGroupId: "rsp-274xltv2sjoxs7fap8tlv3q3s",
 * });
 * const bar = new volcengine.clb.Listener("bar", {
 *     enabled: "on",
 *     healthCheck: {
 *         domain: "volcengine.com",
 *         enabled: "on",
 *         healthyThreshold: 5,
 *         httpCode: "http_2xx",
 *         interval: 10,
 *         method: "GET",
 *         timeout: 3,
 *         unHealthyThreshold: 2,
 *         uri: "/",
 *     },
 *     listenerName: "Demo-HTTP-91",
 *     loadBalancerId: "clb-274xltt3rfmyo7fap8sv1jq39",
 *     port: 91,
 *     protocol: "HTTP",
 *     serverGroupId: "rsp-274xltv2sjoxs7fap8tlv3q3s",
 * });
 * const demo = new volcengine.clb.Listener("demo", {
 *     enabled: "on",
 *     establishedTimeout: 10,
 *     healthCheck: {
 *         enabled: "on",
 *         healthyThreshold: 5,
 *         interval: 10,
 *         timeout: 3,
 *         unHealthyThreshold: 2,
 *     },
 *     loadBalancerId: "clb-274xltt3rfmyo7fap8sv1jq39",
 *     port: 92,
 *     protocol: "TCP",
 *     serverGroupId: "rsp-274xltv2sjoxs7fap8tlv3q3s",
 * });
 * ```
 *
 * ## Import
 *
 * Listener can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:clb/listener:Listener default lsn-273yv0mhs5xj47fap8sehiiso
 * ```
 */
export class Listener extends pulumi.CustomResource {
    /**
     * Get an existing Listener resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListenerState, opts?: pulumi.CustomResourceOptions): Listener {
        return new Listener(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:clb/listener:Listener';

    /**
     * Returns true if the given object is an instance of Listener.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Listener {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Listener.__pulumiType;
    }

    /**
     * The id list of the Acl.
     */
    public readonly aclIds!: pulumi.Output<string[] | undefined>;
    /**
     * The enable status of Acl. Optional choice contains `on`, `off`.
     */
    public readonly aclStatus!: pulumi.Output<string>;
    /**
     * The type of the Acl. Optional choice contains `white`, `black`.
     */
    public readonly aclType!: pulumi.Output<string>;
    /**
     * The certificate id associated with the listener.
     */
    public readonly certificateId!: pulumi.Output<string | undefined>;
    /**
     * The description of the Listener.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The enable status of the Listener. Optional choice contains `on`, `off`.
     */
    public readonly enabled!: pulumi.Output<string>;
    /**
     * The connection timeout of the Listener.
     */
    public readonly establishedTimeout!: pulumi.Output<number>;
    /**
     * The config of health check.
     */
    public readonly healthCheck!: pulumi.Output<outputs.clb.ListenerHealthCheck>;
    /**
     * The ID of the Listener.
     */
    public /*out*/ readonly listenerId!: pulumi.Output<string>;
    /**
     * The name of the Listener.
     */
    public readonly listenerName!: pulumi.Output<string>;
    /**
     * The region of the request.
     */
    public readonly loadBalancerId!: pulumi.Output<string>;
    /**
     * The port receiving request of the Listener, the value range in 1~65535.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * The protocol of the Listener. Optional choice contains `TCP`, `UDP`, `HTTP`, `HTTPS`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The scheduling algorithm of the Listener. Optional choice contains `wrr`, `wlc`, `sh`.
     */
    public readonly scheduler!: pulumi.Output<string>;
    /**
     * The server group id associated with the listener.
     */
    public readonly serverGroupId!: pulumi.Output<string>;

    /**
     * Create a Listener resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListenerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ListenerArgs | ListenerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ListenerState | undefined;
            resourceInputs["aclIds"] = state ? state.aclIds : undefined;
            resourceInputs["aclStatus"] = state ? state.aclStatus : undefined;
            resourceInputs["aclType"] = state ? state.aclType : undefined;
            resourceInputs["certificateId"] = state ? state.certificateId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["establishedTimeout"] = state ? state.establishedTimeout : undefined;
            resourceInputs["healthCheck"] = state ? state.healthCheck : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["listenerName"] = state ? state.listenerName : undefined;
            resourceInputs["loadBalancerId"] = state ? state.loadBalancerId : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["scheduler"] = state ? state.scheduler : undefined;
            resourceInputs["serverGroupId"] = state ? state.serverGroupId : undefined;
        } else {
            const args = argsOrState as ListenerArgs | undefined;
            if ((!args || args.loadBalancerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loadBalancerId'");
            }
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.serverGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverGroupId'");
            }
            resourceInputs["aclIds"] = args ? args.aclIds : undefined;
            resourceInputs["aclStatus"] = args ? args.aclStatus : undefined;
            resourceInputs["aclType"] = args ? args.aclType : undefined;
            resourceInputs["certificateId"] = args ? args.certificateId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["establishedTimeout"] = args ? args.establishedTimeout : undefined;
            resourceInputs["healthCheck"] = args ? args.healthCheck : undefined;
            resourceInputs["listenerName"] = args ? args.listenerName : undefined;
            resourceInputs["loadBalancerId"] = args ? args.loadBalancerId : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["scheduler"] = args ? args.scheduler : undefined;
            resourceInputs["serverGroupId"] = args ? args.serverGroupId : undefined;
            resourceInputs["listenerId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Listener.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Listener resources.
 */
export interface ListenerState {
    /**
     * The id list of the Acl.
     */
    aclIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The enable status of Acl. Optional choice contains `on`, `off`.
     */
    aclStatus?: pulumi.Input<string>;
    /**
     * The type of the Acl. Optional choice contains `white`, `black`.
     */
    aclType?: pulumi.Input<string>;
    /**
     * The certificate id associated with the listener.
     */
    certificateId?: pulumi.Input<string>;
    /**
     * The description of the Listener.
     */
    description?: pulumi.Input<string>;
    /**
     * The enable status of the Listener. Optional choice contains `on`, `off`.
     */
    enabled?: pulumi.Input<string>;
    /**
     * The connection timeout of the Listener.
     */
    establishedTimeout?: pulumi.Input<number>;
    /**
     * The config of health check.
     */
    healthCheck?: pulumi.Input<inputs.clb.ListenerHealthCheck>;
    /**
     * The ID of the Listener.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * The name of the Listener.
     */
    listenerName?: pulumi.Input<string>;
    /**
     * The region of the request.
     */
    loadBalancerId?: pulumi.Input<string>;
    /**
     * The port receiving request of the Listener, the value range in 1~65535.
     */
    port?: pulumi.Input<number>;
    /**
     * The protocol of the Listener. Optional choice contains `TCP`, `UDP`, `HTTP`, `HTTPS`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The scheduling algorithm of the Listener. Optional choice contains `wrr`, `wlc`, `sh`.
     */
    scheduler?: pulumi.Input<string>;
    /**
     * The server group id associated with the listener.
     */
    serverGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Listener resource.
 */
export interface ListenerArgs {
    /**
     * The id list of the Acl.
     */
    aclIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The enable status of Acl. Optional choice contains `on`, `off`.
     */
    aclStatus?: pulumi.Input<string>;
    /**
     * The type of the Acl. Optional choice contains `white`, `black`.
     */
    aclType?: pulumi.Input<string>;
    /**
     * The certificate id associated with the listener.
     */
    certificateId?: pulumi.Input<string>;
    /**
     * The description of the Listener.
     */
    description?: pulumi.Input<string>;
    /**
     * The enable status of the Listener. Optional choice contains `on`, `off`.
     */
    enabled?: pulumi.Input<string>;
    /**
     * The connection timeout of the Listener.
     */
    establishedTimeout?: pulumi.Input<number>;
    /**
     * The config of health check.
     */
    healthCheck?: pulumi.Input<inputs.clb.ListenerHealthCheck>;
    /**
     * The name of the Listener.
     */
    listenerName?: pulumi.Input<string>;
    /**
     * The region of the request.
     */
    loadBalancerId: pulumi.Input<string>;
    /**
     * The port receiving request of the Listener, the value range in 1~65535.
     */
    port: pulumi.Input<number>;
    /**
     * The protocol of the Listener. Optional choice contains `TCP`, `UDP`, `HTTP`, `HTTPS`.
     */
    protocol: pulumi.Input<string>;
    /**
     * The scheduling algorithm of the Listener. Optional choice contains `wrr`, `wlc`, `sh`.
     */
    scheduler?: pulumi.Input<string>;
    /**
     * The server group id associated with the listener.
     */
    serverGroupId: pulumi.Input<string>;
}