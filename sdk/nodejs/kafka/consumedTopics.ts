// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of kafka consumed topics
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooZones = volcengine.ecs.Zones({});
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
 *     ],
 * });
 * const fooGroup = new volcengine.kafka.Group("fooGroup", {
 *     instanceId: fooInstance.id,
 *     groupId: "acc-test-group",
 *     description: "tf-test",
 * });
 * const fooSaslUser = new volcengine.kafka.SaslUser("fooSaslUser", {
 *     userName: "acc-test-user",
 *     instanceId: fooInstance.id,
 *     userPassword: "suqsnis123!",
 *     description: "tf-test",
 *     allAuthority: true,
 *     passwordType: "Scram",
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
 *     accessPolicies: [{
 *         userName: fooSaslUser.userName,
 *         accessPolicy: "Pub",
 *     }],
 * });
 * const default = volcengine.kafka.ConsumedTopicsOutput({
 *     instanceId: fooInstance.id,
 *     groupId: fooGroup.groupId,
 *     topicName: fooTopic.topicName,
 * });
 * ```
 */
export function consumedTopics(args: ConsumedTopicsArgs, opts?: pulumi.InvokeOptions): Promise<ConsumedTopicsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:kafka/consumedTopics:ConsumedTopics", {
        "groupId": args.groupId,
        "instanceId": args.instanceId,
        "outputFile": args.outputFile,
        "topicName": args.topicName,
    }, opts);
}

/**
 * A collection of arguments for invoking ConsumedTopics.
 */
export interface ConsumedTopicsArgs {
    /**
     * The id of kafka group.
     */
    groupId: string;
    /**
     * The id of kafka instance.
     */
    instanceId: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The name of kafka topic. This field supports fuzzy query.
     */
    topicName?: string;
}

/**
 * A collection of values returned by ConsumedTopics.
 */
export interface ConsumedTopicsResult {
    /**
     * The collection of query.
     */
    readonly consumedTopics: outputs.kafka.ConsumedTopicsConsumedTopic[];
    readonly groupId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly outputFile?: string;
    /**
     * The name of kafka topic.
     */
    readonly topicName?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of kafka consumed topics
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooZones = volcengine.ecs.Zones({});
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
 *     ],
 * });
 * const fooGroup = new volcengine.kafka.Group("fooGroup", {
 *     instanceId: fooInstance.id,
 *     groupId: "acc-test-group",
 *     description: "tf-test",
 * });
 * const fooSaslUser = new volcengine.kafka.SaslUser("fooSaslUser", {
 *     userName: "acc-test-user",
 *     instanceId: fooInstance.id,
 *     userPassword: "suqsnis123!",
 *     description: "tf-test",
 *     allAuthority: true,
 *     passwordType: "Scram",
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
 *     accessPolicies: [{
 *         userName: fooSaslUser.userName,
 *         accessPolicy: "Pub",
 *     }],
 * });
 * const default = volcengine.kafka.ConsumedTopicsOutput({
 *     instanceId: fooInstance.id,
 *     groupId: fooGroup.groupId,
 *     topicName: fooTopic.topicName,
 * });
 * ```
 */
export function consumedTopicsOutput(args: ConsumedTopicsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ConsumedTopicsResult> {
    return pulumi.output(args).apply((a: any) => consumedTopics(a, opts))
}

/**
 * A collection of arguments for invoking ConsumedTopics.
 */
export interface ConsumedTopicsOutputArgs {
    /**
     * The id of kafka group.
     */
    groupId: pulumi.Input<string>;
    /**
     * The id of kafka instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The name of kafka topic. This field supports fuzzy query.
     */
    topicName?: pulumi.Input<string>;
}