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
 *     dnsServers: [
 *         "8.8.8.8",
 *         "114.114.114.114",
 *     ],
 *     projectName: "default",
 * });
 * // create subnet
 * const fooSubnet = new volcengine.vpc.Subnet("fooSubnet", {
 *     subnetName: "acc-test-subnet",
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *     vpcId: fooVpc.id,
 * });
 * // create kafka instance
 * const fooInstance = new volcengine.kafka.Instance("fooInstance", {
 *     instanceName: "acc-test-kafka",
 *     instanceDescription: "tf-test",
 *     version: "2.2.2",
 *     computeSpec: "kafka.20xrate.hw",
 *     subnetId: fooSubnet.id,
 *     userName: "tf-user",
 *     userPassword: "tf-pass!@q1",
 *     chargeType: "PostPaid",
 *     storageSpace: 300,
 *     partitionNumber: 350,
 *     projectName: "default",
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 *     parameters: [
 *         {
 *             parameterName: "MessageMaxByte",
 *             parameterValue: "12",
 *         },
 *         {
 *             parameterName: "LogRetentionHours",
 *             parameterValue: "70",
 *         },
 *         {
 *             parameterName: "MessageTimestampType",
 *             parameterValue: "CreateTime",
 *         },
 *         {
 *             parameterName: "OffsetRetentionMinutes",
 *             parameterValue: "10080",
 *         },
 *         {
 *             parameterName: "AutoDeleteGroup",
 *             parameterValue: "false",
 *         },
 *     ],
 * });
 * const fooAddress = new volcengine.eip.Address("fooAddress", {
 *     billingType: "PostPaidByBandwidth",
 *     bandwidth: 1,
 *     isp: "BGP",
 *     description: "tf-test",
 *     projectName: "default",
 * });
 * const fooPublicAddress = new volcengine.kafka.PublicAddress("fooPublicAddress", {
 *     instanceId: fooInstance.id,
 *     eipId: fooAddress.id,
 * });
 * const fooGroup = new volcengine.kafka.Group("fooGroup", {
 *     instanceId: fooInstance.id,
 *     groupId: "acc-test-group",
 *     description: "tf-test",
 * });
 * const fooTopic = new volcengine.kafka.Topic("fooTopic", {
 *     topicName: "acc-test-topic",
 *     instanceId: fooInstance.id,
 *     description: "tf-test",
 *     partitionNumber: 15,
 *     replicaNumber: 3,
 *     parameters: {
 *         minInsyncReplicaNumber: 2,
 *         messageMaxByte: 10,
 *         logRetentionHours: 96,
 *     },
 *     allAuthority: false,
 * });
 * ```
 *
 * ## Import
 *
 * KafkaInstance can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:kafka/instance:Instance default kafka-insbjwbbwb
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:kafka/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * The auto renew flag of instance. Only effective when instanceChargeType is PrePaid. Default is false.
     */
    public readonly autoRenew!: pulumi.Output<boolean | undefined>;
    /**
     * The charge type of instance, the value can be `PrePaid` or `PostPaid`.
     */
    public readonly chargeType!: pulumi.Output<string>;
    /**
     * The compute spec of instance.
     */
    public readonly computeSpec!: pulumi.Output<string>;
    /**
     * The description of instance.
     */
    public readonly instanceDescription!: pulumi.Output<string>;
    /**
     * The name of instance.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * Whether enable rebalance. Only effected in modify when computeSpec field is changed.
     */
    public readonly needRebalance!: pulumi.Output<boolean | undefined>;
    /**
     * Parameter of the instance.
     */
    public readonly parameters!: pulumi.Output<outputs.kafka.InstanceParameter[] | undefined>;
    /**
     * The partition number of instance.
     */
    public readonly partitionNumber!: pulumi.Output<number>;
    /**
     * The period of instance. Only effective when instanceChargeType is PrePaid. Unit is Month.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The project name of instance.
     */
    public readonly projectName!: pulumi.Output<string>;
    /**
     * The rebalance time.
     */
    public readonly rebalanceTime!: pulumi.Output<string | undefined>;
    /**
     * The storage space of instance.
     */
    public readonly storageSpace!: pulumi.Output<number>;
    /**
     * The storage type of instance. The value can be ESSD_FlexPL or ESSD_PL0.
     */
    public readonly storageType!: pulumi.Output<string | undefined>;
    /**
     * The subnet id of instance.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * The tags of instance.
     */
    public readonly tags!: pulumi.Output<outputs.kafka.InstanceTag[] | undefined>;
    /**
     * The user name of instance. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly userName!: pulumi.Output<string>;
    /**
     * The user password of instance. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly userPassword!: pulumi.Output<string>;
    /**
     * The version of instance, the value can be `2.2.2` or `2.8.2`.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["chargeType"] = state ? state.chargeType : undefined;
            resourceInputs["computeSpec"] = state ? state.computeSpec : undefined;
            resourceInputs["instanceDescription"] = state ? state.instanceDescription : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["needRebalance"] = state ? state.needRebalance : undefined;
            resourceInputs["parameters"] = state ? state.parameters : undefined;
            resourceInputs["partitionNumber"] = state ? state.partitionNumber : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["rebalanceTime"] = state ? state.rebalanceTime : undefined;
            resourceInputs["storageSpace"] = state ? state.storageSpace : undefined;
            resourceInputs["storageType"] = state ? state.storageType : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
            resourceInputs["userPassword"] = state ? state.userPassword : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.chargeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'chargeType'");
            }
            if ((!args || args.computeSpec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'computeSpec'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            if ((!args || args.userPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userPassword'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["chargeType"] = args ? args.chargeType : undefined;
            resourceInputs["computeSpec"] = args ? args.computeSpec : undefined;
            resourceInputs["instanceDescription"] = args ? args.instanceDescription : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["needRebalance"] = args ? args.needRebalance : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["partitionNumber"] = args ? args.partitionNumber : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["rebalanceTime"] = args ? args.rebalanceTime : undefined;
            resourceInputs["storageSpace"] = args ? args.storageSpace : undefined;
            resourceInputs["storageType"] = args ? args.storageType : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["userPassword"] = args?.userPassword ? pulumi.secret(args.userPassword) : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["userPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * The auto renew flag of instance. Only effective when instanceChargeType is PrePaid. Default is false.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * The charge type of instance, the value can be `PrePaid` or `PostPaid`.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * The compute spec of instance.
     */
    computeSpec?: pulumi.Input<string>;
    /**
     * The description of instance.
     */
    instanceDescription?: pulumi.Input<string>;
    /**
     * The name of instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Whether enable rebalance. Only effected in modify when computeSpec field is changed.
     */
    needRebalance?: pulumi.Input<boolean>;
    /**
     * Parameter of the instance.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.kafka.InstanceParameter>[]>;
    /**
     * The partition number of instance.
     */
    partitionNumber?: pulumi.Input<number>;
    /**
     * The period of instance. Only effective when instanceChargeType is PrePaid. Unit is Month.
     */
    period?: pulumi.Input<number>;
    /**
     * The project name of instance.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The rebalance time.
     */
    rebalanceTime?: pulumi.Input<string>;
    /**
     * The storage space of instance.
     */
    storageSpace?: pulumi.Input<number>;
    /**
     * The storage type of instance. The value can be ESSD_FlexPL or ESSD_PL0.
     */
    storageType?: pulumi.Input<string>;
    /**
     * The subnet id of instance.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The tags of instance.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.kafka.InstanceTag>[]>;
    /**
     * The user name of instance. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    userName?: pulumi.Input<string>;
    /**
     * The user password of instance. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    userPassword?: pulumi.Input<string>;
    /**
     * The version of instance, the value can be `2.2.2` or `2.8.2`.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The auto renew flag of instance. Only effective when instanceChargeType is PrePaid. Default is false.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * The charge type of instance, the value can be `PrePaid` or `PostPaid`.
     */
    chargeType: pulumi.Input<string>;
    /**
     * The compute spec of instance.
     */
    computeSpec: pulumi.Input<string>;
    /**
     * The description of instance.
     */
    instanceDescription?: pulumi.Input<string>;
    /**
     * The name of instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Whether enable rebalance. Only effected in modify when computeSpec field is changed.
     */
    needRebalance?: pulumi.Input<boolean>;
    /**
     * Parameter of the instance.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.kafka.InstanceParameter>[]>;
    /**
     * The partition number of instance.
     */
    partitionNumber?: pulumi.Input<number>;
    /**
     * The period of instance. Only effective when instanceChargeType is PrePaid. Unit is Month.
     */
    period?: pulumi.Input<number>;
    /**
     * The project name of instance.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The rebalance time.
     */
    rebalanceTime?: pulumi.Input<string>;
    /**
     * The storage space of instance.
     */
    storageSpace?: pulumi.Input<number>;
    /**
     * The storage type of instance. The value can be ESSD_FlexPL or ESSD_PL0.
     */
    storageType?: pulumi.Input<string>;
    /**
     * The subnet id of instance.
     */
    subnetId: pulumi.Input<string>;
    /**
     * The tags of instance.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.kafka.InstanceTag>[]>;
    /**
     * The user name of instance. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    userName: pulumi.Input<string>;
    /**
     * The user password of instance. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    userPassword: pulumi.Input<string>;
    /**
     * The version of instance, the value can be `2.2.2` or `2.8.2`.
     */
    version: pulumi.Input<string>;
}
