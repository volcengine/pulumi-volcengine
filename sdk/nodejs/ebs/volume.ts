// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooZones = volcengine.ecs.getZones({});
 * // create vpc
 * const fooVpc = new volcengine.vpc.Vpc("fooVpc", {
 *     vpcName: "acc-test-vpc",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * // create subnet
 * const fooSubnet = new volcengine.vpc.Subnet("fooSubnet", {
 *     subnetName: "acc-test-subnet",
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *     vpcId: fooVpc.id,
 * });
 * // create security group
 * const fooSecurityGroup = new volcengine.vpc.SecurityGroup("fooSecurityGroup", {
 *     securityGroupName: "acc-test-security-group",
 *     vpcId: fooVpc.id,
 * });
 * const fooImages = volcengine.ecs.getImages({
 *     osType: "Linux",
 *     visibility: "public",
 *     instanceTypeId: "ecs.g3il.large",
 * });
 * // create PrePaid ecs instance
 * const fooInstance = new volcengine.ecs.Instance("fooInstance", {
 *     instanceName: "acc-test-ecs",
 *     description: "acc-test",
 *     hostName: "tf-acc-test",
 *     imageId: fooImages.then(fooImages => fooImages.images?.[0]?.imageId),
 *     instanceType: "ecs.g3il.large",
 *     password: "93f0cb0614Aab12",
 *     instanceChargeType: "PrePaid",
 *     period: 1,
 *     systemVolumeType: "ESSD_PL0",
 *     systemVolumeSize: 40,
 *     subnetId: fooSubnet.id,
 *     securityGroupIds: [fooSecurityGroup.id],
 *     projectName: "default",
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * // create PrePaid data volume
 * const preVolume = new volcengine.ebs.Volume("preVolume", {
 *     volumeName: "acc-test-volume",
 *     volumeType: "ESSD_PL0",
 *     description: "acc-test",
 *     kind: "data",
 *     size: 40,
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *     volumeChargeType: "PrePaid",
 *     instanceId: fooInstance.id,
 *     projectName: "default",
 *     deleteWithInstance: true,
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * // create PostPaid data volume
 * const postVolume = new volcengine.ebs.Volume("postVolume", {
 *     volumeName: "acc-test-volume",
 *     volumeType: "ESSD_PL0",
 *     description: "acc-test",
 *     kind: "data",
 *     size: 40,
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *     volumeChargeType: "PostPaid",
 *     projectName: "default",
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * // attach PostPaid data volume to ecs instance
 * const fooVolumeAttach = new volcengine.ebs.VolumeAttach("fooVolumeAttach", {
 *     instanceId: fooInstance.id,
 *     volumeId: postVolume.id,
 * });
 * ```
 *
 * ## Import
 *
 * Volume can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:ebs/volume:Volume default vol-mizl7m1kqccg5smt1bdpijuj
 * ```
 */
export class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeState, opts?: pulumi.CustomResourceOptions): Volume {
        return new Volume(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:ebs/volume:Volume';

    /**
     * Returns true if the given object is an instance of Volume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Volume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Volume.__pulumiType;
    }

    /**
     * Creation time of Volume.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Delete Volume with Attached Instance.
     */
    public readonly deleteWithInstance!: pulumi.Output<boolean>;
    /**
     * The description of the Volume.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The extra IOPS performance size for volume. Unit: times per second. The valid values for `Balance` and `IOPS` is 0~50000.
     */
    public readonly extraPerformanceIops!: pulumi.Output<number>;
    /**
     * The extra Throughput performance size for volume. Unit: MB/s. The valid values for ESSD FlexPL volume is 0~650.
     */
    public readonly extraPerformanceThroughputMb!: pulumi.Output<number>;
    /**
     * The type of extra performance for volume. The valid values for ESSD FlexPL volume are `Throughput`, `Balance`, `IOPS`. The valid value for TSSD_TL0 volume is `Throughput`.
     */
    public readonly extraPerformanceTypeId!: pulumi.Output<string | undefined>;
    /**
     * The ID of the instance to which the created volume is automatically attached. It is recommended to attach the PostPaid
     * volume to instance through resource `volume_attach`.When use this field to attach ecs instance, the attached volume
     * cannot be deleted by terraform, please use `terraform state rm volcengine_volume.resource_name` command to remove it
     * from terraform state file and management.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The kind of Volume, the value is `data`.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The ProjectName of the Volume.
     */
    public readonly projectName!: pulumi.Output<string>;
    /**
     * The size of Volume.
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * The id of the snapshot. When creating a volume using snapshots, this field is required.
     * When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly snapshotId!: pulumi.Output<string | undefined>;
    /**
     * Status of Volume.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Tags.
     */
    public readonly tags!: pulumi.Output<outputs.ebs.VolumeTag[] | undefined>;
    /**
     * Status of Trade.
     */
    public /*out*/ readonly tradeStatus!: pulumi.Output<number>;
    /**
     * The charge type of the Volume, the value is `PostPaid` or `PrePaid`. The `PrePaid` volume cannot be detached.
     */
    public readonly volumeChargeType!: pulumi.Output<string | undefined>;
    /**
     * The name of Volume.
     */
    public readonly volumeName!: pulumi.Output<string>;
    /**
     * The type of Volume. Valid values: `ESSD_PL0`, `ESSD_FlexPL`, `TSSD_TL0`.
     */
    public readonly volumeType!: pulumi.Output<string>;
    /**
     * The id of the Zone.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a Volume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeArgs | VolumeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VolumeState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["deleteWithInstance"] = state ? state.deleteWithInstance : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["extraPerformanceIops"] = state ? state.extraPerformanceIops : undefined;
            resourceInputs["extraPerformanceThroughputMb"] = state ? state.extraPerformanceThroughputMb : undefined;
            resourceInputs["extraPerformanceTypeId"] = state ? state.extraPerformanceTypeId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["kind"] = state ? state.kind : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["snapshotId"] = state ? state.snapshotId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tradeStatus"] = state ? state.tradeStatus : undefined;
            resourceInputs["volumeChargeType"] = state ? state.volumeChargeType : undefined;
            resourceInputs["volumeName"] = state ? state.volumeName : undefined;
            resourceInputs["volumeType"] = state ? state.volumeType : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as VolumeArgs | undefined;
            if ((!args || args.kind === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kind'");
            }
            if ((!args || args.size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'size'");
            }
            if ((!args || args.volumeName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'volumeName'");
            }
            if ((!args || args.volumeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'volumeType'");
            }
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["deleteWithInstance"] = args ? args.deleteWithInstance : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["extraPerformanceIops"] = args ? args.extraPerformanceIops : undefined;
            resourceInputs["extraPerformanceThroughputMb"] = args ? args.extraPerformanceThroughputMb : undefined;
            resourceInputs["extraPerformanceTypeId"] = args ? args.extraPerformanceTypeId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["snapshotId"] = args ? args.snapshotId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["volumeChargeType"] = args ? args.volumeChargeType : undefined;
            resourceInputs["volumeName"] = args ? args.volumeName : undefined;
            resourceInputs["volumeType"] = args ? args.volumeType : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tradeStatus"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Volume.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Volume resources.
 */
export interface VolumeState {
    /**
     * Creation time of Volume.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Delete Volume with Attached Instance.
     */
    deleteWithInstance?: pulumi.Input<boolean>;
    /**
     * The description of the Volume.
     */
    description?: pulumi.Input<string>;
    /**
     * The extra IOPS performance size for volume. Unit: times per second. The valid values for `Balance` and `IOPS` is 0~50000.
     */
    extraPerformanceIops?: pulumi.Input<number>;
    /**
     * The extra Throughput performance size for volume. Unit: MB/s. The valid values for ESSD FlexPL volume is 0~650.
     */
    extraPerformanceThroughputMb?: pulumi.Input<number>;
    /**
     * The type of extra performance for volume. The valid values for ESSD FlexPL volume are `Throughput`, `Balance`, `IOPS`. The valid value for TSSD_TL0 volume is `Throughput`.
     */
    extraPerformanceTypeId?: pulumi.Input<string>;
    /**
     * The ID of the instance to which the created volume is automatically attached. It is recommended to attach the PostPaid
     * volume to instance through resource `volume_attach`.When use this field to attach ecs instance, the attached volume
     * cannot be deleted by terraform, please use `terraform state rm volcengine_volume.resource_name` command to remove it
     * from terraform state file and management.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The kind of Volume, the value is `data`.
     */
    kind?: pulumi.Input<string>;
    /**
     * The ProjectName of the Volume.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The size of Volume.
     */
    size?: pulumi.Input<number>;
    /**
     * The id of the snapshot. When creating a volume using snapshots, this field is required.
     * When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * Status of Volume.
     */
    status?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.ebs.VolumeTag>[]>;
    /**
     * Status of Trade.
     */
    tradeStatus?: pulumi.Input<number>;
    /**
     * The charge type of the Volume, the value is `PostPaid` or `PrePaid`. The `PrePaid` volume cannot be detached.
     */
    volumeChargeType?: pulumi.Input<string>;
    /**
     * The name of Volume.
     */
    volumeName?: pulumi.Input<string>;
    /**
     * The type of Volume. Valid values: `ESSD_PL0`, `ESSD_FlexPL`, `TSSD_TL0`.
     */
    volumeType?: pulumi.Input<string>;
    /**
     * The id of the Zone.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Volume resource.
 */
export interface VolumeArgs {
    /**
     * Delete Volume with Attached Instance.
     */
    deleteWithInstance?: pulumi.Input<boolean>;
    /**
     * The description of the Volume.
     */
    description?: pulumi.Input<string>;
    /**
     * The extra IOPS performance size for volume. Unit: times per second. The valid values for `Balance` and `IOPS` is 0~50000.
     */
    extraPerformanceIops?: pulumi.Input<number>;
    /**
     * The extra Throughput performance size for volume. Unit: MB/s. The valid values for ESSD FlexPL volume is 0~650.
     */
    extraPerformanceThroughputMb?: pulumi.Input<number>;
    /**
     * The type of extra performance for volume. The valid values for ESSD FlexPL volume are `Throughput`, `Balance`, `IOPS`. The valid value for TSSD_TL0 volume is `Throughput`.
     */
    extraPerformanceTypeId?: pulumi.Input<string>;
    /**
     * The ID of the instance to which the created volume is automatically attached. It is recommended to attach the PostPaid
     * volume to instance through resource `volume_attach`.When use this field to attach ecs instance, the attached volume
     * cannot be deleted by terraform, please use `terraform state rm volcengine_volume.resource_name` command to remove it
     * from terraform state file and management.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The kind of Volume, the value is `data`.
     */
    kind: pulumi.Input<string>;
    /**
     * The ProjectName of the Volume.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The size of Volume.
     */
    size: pulumi.Input<number>;
    /**
     * The id of the snapshot. When creating a volume using snapshots, this field is required.
     * When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.ebs.VolumeTag>[]>;
    /**
     * The charge type of the Volume, the value is `PostPaid` or `PrePaid`. The `PrePaid` volume cannot be detached.
     */
    volumeChargeType?: pulumi.Input<string>;
    /**
     * The name of Volume.
     */
    volumeName: pulumi.Input<string>;
    /**
     * The type of Volume. Valid values: `ESSD_PL0`, `ESSD_FlexPL`, `TSSD_TL0`.
     */
    volumeType: pulumi.Input<string>;
    /**
     * The id of the Zone.
     */
    zoneId: pulumi.Input<string>;
}
