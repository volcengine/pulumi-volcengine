// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage scaling lifecycle hook
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = new volcengine.autoscaling.ScalingLifecycleHook("foo", {
 *     lifecycleHookName: "tf-test",
 *     lifecycleHookPolicy: "CONTINUE",
 *     lifecycleHookTimeout: 30,
 *     lifecycleHookType: "SCALE_IN",
 *     scalingGroupId: "scg-ybru8pazhgl8j1di4tyd",
 * });
 * ```
 *
 * ## Import
 *
 * ScalingLifecycleHook can be imported using the ScalingGroupId:LifecycleHookId, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:autoscaling/scalingLifecycleHook:ScalingLifecycleHook default scg-yblfbfhy7agh9zn72iaz:sgh-ybqholahe4gso0ee88sd
 * ```
 */
export class ScalingLifecycleHook extends pulumi.CustomResource {
    /**
     * Get an existing ScalingLifecycleHook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScalingLifecycleHookState, opts?: pulumi.CustomResourceOptions): ScalingLifecycleHook {
        return new ScalingLifecycleHook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:autoscaling/scalingLifecycleHook:ScalingLifecycleHook';

    /**
     * Returns true if the given object is an instance of ScalingLifecycleHook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScalingLifecycleHook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScalingLifecycleHook.__pulumiType;
    }

    /**
     * The id of the lifecycle hook.
     */
    public /*out*/ readonly lifecycleHookId!: pulumi.Output<string>;
    /**
     * The name of the lifecycle hook.
     */
    public readonly lifecycleHookName!: pulumi.Output<string>;
    /**
     * The policy of the lifecycle hook. Valid values: CONTINUE, REJECT.
     */
    public readonly lifecycleHookPolicy!: pulumi.Output<string>;
    /**
     * The timeout of the lifecycle hook.
     */
    public readonly lifecycleHookTimeout!: pulumi.Output<number>;
    /**
     * The type of the lifecycle hook. Valid values: SCALE_IN, SCALE_OUT.
     */
    public readonly lifecycleHookType!: pulumi.Output<string>;
    /**
     * The id of the scaling group.
     */
    public readonly scalingGroupId!: pulumi.Output<string>;

    /**
     * Create a ScalingLifecycleHook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScalingLifecycleHookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScalingLifecycleHookArgs | ScalingLifecycleHookState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ScalingLifecycleHookState | undefined;
            resourceInputs["lifecycleHookId"] = state ? state.lifecycleHookId : undefined;
            resourceInputs["lifecycleHookName"] = state ? state.lifecycleHookName : undefined;
            resourceInputs["lifecycleHookPolicy"] = state ? state.lifecycleHookPolicy : undefined;
            resourceInputs["lifecycleHookTimeout"] = state ? state.lifecycleHookTimeout : undefined;
            resourceInputs["lifecycleHookType"] = state ? state.lifecycleHookType : undefined;
            resourceInputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
        } else {
            const args = argsOrState as ScalingLifecycleHookArgs | undefined;
            if ((!args || args.lifecycleHookName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lifecycleHookName'");
            }
            if ((!args || args.lifecycleHookPolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lifecycleHookPolicy'");
            }
            if ((!args || args.lifecycleHookTimeout === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lifecycleHookTimeout'");
            }
            if ((!args || args.lifecycleHookType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lifecycleHookType'");
            }
            if ((!args || args.scalingGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scalingGroupId'");
            }
            resourceInputs["lifecycleHookName"] = args ? args.lifecycleHookName : undefined;
            resourceInputs["lifecycleHookPolicy"] = args ? args.lifecycleHookPolicy : undefined;
            resourceInputs["lifecycleHookTimeout"] = args ? args.lifecycleHookTimeout : undefined;
            resourceInputs["lifecycleHookType"] = args ? args.lifecycleHookType : undefined;
            resourceInputs["scalingGroupId"] = args ? args.scalingGroupId : undefined;
            resourceInputs["lifecycleHookId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ScalingLifecycleHook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScalingLifecycleHook resources.
 */
export interface ScalingLifecycleHookState {
    /**
     * The id of the lifecycle hook.
     */
    lifecycleHookId?: pulumi.Input<string>;
    /**
     * The name of the lifecycle hook.
     */
    lifecycleHookName?: pulumi.Input<string>;
    /**
     * The policy of the lifecycle hook. Valid values: CONTINUE, REJECT.
     */
    lifecycleHookPolicy?: pulumi.Input<string>;
    /**
     * The timeout of the lifecycle hook.
     */
    lifecycleHookTimeout?: pulumi.Input<number>;
    /**
     * The type of the lifecycle hook. Valid values: SCALE_IN, SCALE_OUT.
     */
    lifecycleHookType?: pulumi.Input<string>;
    /**
     * The id of the scaling group.
     */
    scalingGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ScalingLifecycleHook resource.
 */
export interface ScalingLifecycleHookArgs {
    /**
     * The name of the lifecycle hook.
     */
    lifecycleHookName: pulumi.Input<string>;
    /**
     * The policy of the lifecycle hook. Valid values: CONTINUE, REJECT.
     */
    lifecycleHookPolicy: pulumi.Input<string>;
    /**
     * The timeout of the lifecycle hook.
     */
    lifecycleHookTimeout: pulumi.Input<number>;
    /**
     * The type of the lifecycle hook. Valid values: SCALE_IN, SCALE_OUT.
     */
    lifecycleHookType: pulumi.Input<string>;
    /**
     * The id of the scaling group.
     */
    scalingGroupId: pulumi.Input<string>;
}