// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage scaling instance attachment
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
 * const fooSecurityGroup = new volcengine.vpc.SecurityGroup("fooSecurityGroup", {
 *     securityGroupName: "acc-test-security-group",
 *     vpcId: fooVpc.id,
 * });
 * const fooImages = volcengine.ecs.getImages({
 *     osType: "Linux",
 *     visibility: "public",
 *     instanceTypeId: "ecs.g1.large",
 * });
 * const fooKeyPair = new volcengine.ecs.KeyPair("fooKeyPair", {
 *     description: "acc-test-2",
 *     keyPairName: "acc-test-key-pair-name",
 * });
 * const fooLaunchTemplate = new volcengine.ecs.LaunchTemplate("fooLaunchTemplate", {
 *     description: "acc-test-desc",
 *     eipBandwidth: 200,
 *     eipBillingType: "PostPaidByBandwidth",
 *     eipIsp: "BGP",
 *     hostName: "acc-hostname",
 *     imageId: fooImages.then(fooImages => fooImages.images?.[0]?.imageId),
 *     instanceChargeType: "PostPaid",
 *     instanceName: "acc-instance-name",
 *     instanceTypeId: "ecs.g1.large",
 *     keyPairName: fooKeyPair.keyPairName,
 *     launchTemplateName: "acc-test-template",
 *     networkInterfaces: [{
 *         subnetId: fooSubnet.id,
 *         securityGroupIds: [fooSecurityGroup.id],
 *     }],
 *     volumes: [{
 *         volumeType: "ESSD_PL0",
 *         size: 50,
 *         deleteWithInstance: true,
 *     }],
 * });
 * const fooScalingGroup = new volcengine.autoscaling.ScalingGroup("fooScalingGroup", {
 *     scalingGroupName: "acc-test-scaling-group",
 *     subnetIds: [fooSubnet.id],
 *     multiAzPolicy: "BALANCE",
 *     desireInstanceNumber: -1,
 *     minInstanceNumber: 0,
 *     maxInstanceNumber: 1,
 *     instanceTerminatePolicy: "OldestInstance",
 *     defaultCooldown: 10,
 *     launchTemplateId: fooLaunchTemplate.id,
 *     launchTemplateVersion: "Default",
 * });
 * const fooScalingGroupEnabler = new volcengine.autoscaling.ScalingGroupEnabler("fooScalingGroupEnabler", {scalingGroupId: fooScalingGroup.id});
 * const fooInstance = new volcengine.ecs.Instance("fooInstance", {
 *     instanceName: "acc-test-ecs",
 *     description: "acc-test",
 *     hostName: "tf-acc-test",
 *     imageId: fooImages.then(fooImages => fooImages.images?.[0]?.imageId),
 *     instanceType: "ecs.g1.large",
 *     password: "93f0cb0614Aab12",
 *     instanceChargeType: "PostPaid",
 *     systemVolumeType: "ESSD_PL0",
 *     systemVolumeSize: 40,
 *     subnetId: fooSubnet.id,
 *     securityGroupIds: [fooSecurityGroup.id],
 * });
 * const fooScalingInstanceAttachment = new volcengine.autoscaling.ScalingInstanceAttachment("fooScalingInstanceAttachment", {
 *     instanceId: fooInstance.id,
 *     scalingGroupId: fooScalingGroup.id,
 *     entrusted: true,
 * }, {
 *     dependsOn: [fooScalingGroupEnabler],
 * });
 * ```
 *
 * ## Import
 *
 * Scaling instance attachment can be imported using the scaling_group_id and instance_id, e.g.
 * You can choose to remove or detach the instance according to the `delete_type` field.
 *
 * ```sh
 * $ pulumi import volcengine:autoscaling/scalingInstanceAttachment:ScalingInstanceAttachment default scg-mizl7m1kqccg5smt1bdpijuj:i-l8u2ai4j0fauo6mrpgk8
 * ```
 */
export class ScalingInstanceAttachment extends pulumi.CustomResource {
    /**
     * Get an existing ScalingInstanceAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScalingInstanceAttachmentState, opts?: pulumi.CustomResourceOptions): ScalingInstanceAttachment {
        return new ScalingInstanceAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:autoscaling/scalingInstanceAttachment:ScalingInstanceAttachment';

    /**
     * Returns true if the given object is an instance of ScalingInstanceAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScalingInstanceAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScalingInstanceAttachment.__pulumiType;
    }

    /**
     * The type of delete activity. Valid values: Remove, Detach. Default value is Remove.
     */
    public readonly deleteType!: pulumi.Output<string | undefined>;
    /**
     * Whether to cancel the association of the instance with the load balancing and public network IP. Valid values: both, none. Default value is both.
     */
    public readonly detachOption!: pulumi.Output<string | undefined>;
    /**
     * Whether to host the instance to a scaling group. Default value is false.
     */
    public readonly entrusted!: pulumi.Output<boolean | undefined>;
    /**
     * The id of the instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The id of the scaling group.
     */
    public readonly scalingGroupId!: pulumi.Output<string>;

    /**
     * Create a ScalingInstanceAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScalingInstanceAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScalingInstanceAttachmentArgs | ScalingInstanceAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ScalingInstanceAttachmentState | undefined;
            resourceInputs["deleteType"] = state ? state.deleteType : undefined;
            resourceInputs["detachOption"] = state ? state.detachOption : undefined;
            resourceInputs["entrusted"] = state ? state.entrusted : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
        } else {
            const args = argsOrState as ScalingInstanceAttachmentArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.scalingGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scalingGroupId'");
            }
            resourceInputs["deleteType"] = args ? args.deleteType : undefined;
            resourceInputs["detachOption"] = args ? args.detachOption : undefined;
            resourceInputs["entrusted"] = args ? args.entrusted : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["scalingGroupId"] = args ? args.scalingGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ScalingInstanceAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScalingInstanceAttachment resources.
 */
export interface ScalingInstanceAttachmentState {
    /**
     * The type of delete activity. Valid values: Remove, Detach. Default value is Remove.
     */
    deleteType?: pulumi.Input<string>;
    /**
     * Whether to cancel the association of the instance with the load balancing and public network IP. Valid values: both, none. Default value is both.
     */
    detachOption?: pulumi.Input<string>;
    /**
     * Whether to host the instance to a scaling group. Default value is false.
     */
    entrusted?: pulumi.Input<boolean>;
    /**
     * The id of the instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The id of the scaling group.
     */
    scalingGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ScalingInstanceAttachment resource.
 */
export interface ScalingInstanceAttachmentArgs {
    /**
     * The type of delete activity. Valid values: Remove, Detach. Default value is Remove.
     */
    deleteType?: pulumi.Input<string>;
    /**
     * Whether to cancel the association of the instance with the load balancing and public network IP. Valid values: both, none. Default value is both.
     */
    detachOption?: pulumi.Input<string>;
    /**
     * Whether to host the instance to a scaling group. Default value is false.
     */
    entrusted?: pulumi.Input<boolean>;
    /**
     * The id of the instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The id of the scaling group.
     */
    scalingGroupId: pulumi.Input<string>;
}
