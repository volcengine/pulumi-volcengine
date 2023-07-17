// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage scaling group enabler
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * // 创建步骤：terraform init -> terraform plan -> terraform apply
 * // 删除步骤: terraform state rm volcengine_scaling_configuration.foo1 -> terraform destroy
 * // 创建伸缩组
 * const foo = new volcengine.autoscaling.ScalingGroup("foo", {
 *     scalingGroupName: "zzm-tf-test",
 *     subnetIds: ["subnet-2fegl9waotzi859gp67relkhv"],
 *     multiAzPolicy: "BALANCE",
 *     desireInstanceNumber: 0,
 *     minInstanceNumber: 0,
 *     maxInstanceNumber: 1,
 *     instanceTerminatePolicy: "OldestInstance",
 *     defaultCooldown: 10,
 * });
 * // 创建伸缩配置
 * const foo1 = new volcengine.autoscaling.ScalingConfiguration("foo1", {
 *     scalingConfigurationName: "terraform-test",
 *     scalingGroupId: foo.scalingGroupId,
 *     imageId: "image-ybx2d38wdfl8j1pupx7b",
 *     instanceTypes: ["ecs.g1.2xlarge"],
 *     instanceName: "tf-test",
 *     instanceDescription: "",
 *     hostName: "",
 *     password: "",
 *     keyPairName: "zktest",
 *     securityEnhancementStrategy: "InActive",
 *     volumes: [
 *         {
 *             volumeType: "ESSD_PL0",
 *             size: 20,
 *             deleteWithInstance: false,
 *         },
 *         {
 *             volumeType: "ESSD_PL0",
 *             size: 20,
 *             deleteWithInstance: true,
 *         },
 *     ],
 *     securityGroupIds: ["sg-12b8llnkn1la817q7y1be4kop"],
 *     eipBandwidth: 0,
 *     eipIsp: "ChinaMobile",
 *     eipBillingType: "PostPaidByBandwidth",
 * });
 * // 绑定伸缩配置
 * const foo2 = new volcengine.autoscaling.ScalingConfigurationAttachment("foo2", {scalingConfigurationId: foo1.scalingConfigurationId}, {
 *     dependsOn: [foo1],
 * });
 * // 启用伸缩组
 * const foo3 = new volcengine.autoscaling.ScalingGroupEnabler("foo3", {scalingGroupId: foo.scalingGroupId}, {
 *     dependsOn: [foo2],
 * });
 * ```
 *
 * ## Import
 *
 * Scaling Group enabler can be imported using the scaling_group_id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:autoscaling/scalingGroupEnabler:ScalingGroupEnabler default enable:scg-mizl7m1kqccg5smt1bdpijuj
 * ```
 */
export class ScalingGroupEnabler extends pulumi.CustomResource {
    /**
     * Get an existing ScalingGroupEnabler resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScalingGroupEnablerState, opts?: pulumi.CustomResourceOptions): ScalingGroupEnabler {
        return new ScalingGroupEnabler(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:autoscaling/scalingGroupEnabler:ScalingGroupEnabler';

    /**
     * Returns true if the given object is an instance of ScalingGroupEnabler.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScalingGroupEnabler {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScalingGroupEnabler.__pulumiType;
    }

    /**
     * The id of the scaling group.
     */
    public readonly scalingGroupId!: pulumi.Output<string>;

    /**
     * Create a ScalingGroupEnabler resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScalingGroupEnablerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScalingGroupEnablerArgs | ScalingGroupEnablerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ScalingGroupEnablerState | undefined;
            resourceInputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
        } else {
            const args = argsOrState as ScalingGroupEnablerArgs | undefined;
            if ((!args || args.scalingGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scalingGroupId'");
            }
            resourceInputs["scalingGroupId"] = args ? args.scalingGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ScalingGroupEnabler.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScalingGroupEnabler resources.
 */
export interface ScalingGroupEnablerState {
    /**
     * The id of the scaling group.
     */
    scalingGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ScalingGroupEnabler resource.
 */
export interface ScalingGroupEnablerArgs {
    /**
     * The id of the scaling group.
     */
    scalingGroupId: pulumi.Input<string>;
}