// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage transit router
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.transit_router.TransitRouter("foo", {
 *     description: "acc-test",
 *     transitRouterName: "acc-test-tr",
 * });
 * ```
 *
 * ## Import
 *
 * TransitRouter can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:transit_router/transitRouter:TransitRouter default tr-2d6fr7mzya2gw58ozfes5g2oh
 * ```
 */
export class TransitRouter extends pulumi.CustomResource {
    /**
     * Get an existing TransitRouter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TransitRouterState, opts?: pulumi.CustomResourceOptions): TransitRouter {
        return new TransitRouter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:transit_router/transitRouter:TransitRouter';

    /**
     * Returns true if the given object is an instance of TransitRouter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransitRouter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransitRouter.__pulumiType;
    }

    /**
     * The ID of account.
     */
    public /*out*/ readonly accountId!: pulumi.Output<string>;
    /**
     * The business status of the transit router.
     */
    public /*out*/ readonly businessStatus!: pulumi.Output<string>;
    /**
     * The create time.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The description of the transit router.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The overdue time.
     */
    public /*out*/ readonly overdueTime!: pulumi.Output<string>;
    /**
     * The status of the transit router.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The attachments of transit router.
     */
    public /*out*/ readonly transitRouterAttachments!: pulumi.Output<outputs.transit_router.TransitRouterTransitRouterAttachment[]>;
    /**
     * The ID of the transit router.
     */
    public /*out*/ readonly transitRouterId!: pulumi.Output<string>;
    /**
     * The name of the transit router.
     */
    public readonly transitRouterName!: pulumi.Output<string>;
    /**
     * The update time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a TransitRouter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TransitRouterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TransitRouterArgs | TransitRouterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TransitRouterState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["businessStatus"] = state ? state.businessStatus : undefined;
            resourceInputs["creationTime"] = state ? state.creationTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["overdueTime"] = state ? state.overdueTime : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["transitRouterAttachments"] = state ? state.transitRouterAttachments : undefined;
            resourceInputs["transitRouterId"] = state ? state.transitRouterId : undefined;
            resourceInputs["transitRouterName"] = state ? state.transitRouterName : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as TransitRouterArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["transitRouterName"] = args ? args.transitRouterName : undefined;
            resourceInputs["accountId"] = undefined /*out*/;
            resourceInputs["businessStatus"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["overdueTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["transitRouterAttachments"] = undefined /*out*/;
            resourceInputs["transitRouterId"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TransitRouter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TransitRouter resources.
 */
export interface TransitRouterState {
    /**
     * The ID of account.
     */
    accountId?: pulumi.Input<string>;
    /**
     * The business status of the transit router.
     */
    businessStatus?: pulumi.Input<string>;
    /**
     * The create time.
     */
    creationTime?: pulumi.Input<string>;
    /**
     * The description of the transit router.
     */
    description?: pulumi.Input<string>;
    /**
     * The overdue time.
     */
    overdueTime?: pulumi.Input<string>;
    /**
     * The status of the transit router.
     */
    status?: pulumi.Input<string>;
    /**
     * The attachments of transit router.
     */
    transitRouterAttachments?: pulumi.Input<pulumi.Input<inputs.transit_router.TransitRouterTransitRouterAttachment>[]>;
    /**
     * The ID of the transit router.
     */
    transitRouterId?: pulumi.Input<string>;
    /**
     * The name of the transit router.
     */
    transitRouterName?: pulumi.Input<string>;
    /**
     * The update time.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TransitRouter resource.
 */
export interface TransitRouterArgs {
    /**
     * The description of the transit router.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the transit router.
     */
    transitRouterName?: pulumi.Input<string>;
}