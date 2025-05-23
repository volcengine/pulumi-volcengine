// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of veecp node pools
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooZones = volcengine.ecs.getZones({});
 * const fooVpc = new volcengine.vpc.Vpc("fooVpc", {
 *     vpcName: "acc-test-project1",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const fooSubnet = new volcengine.vpc.Subnet("fooSubnet", {
 *     subnetName: "acc-subnet-test-2",
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *     vpcId: fooVpc.id,
 * });
 * const fooSecurityGroup = new volcengine.vpc.SecurityGroup("fooSecurityGroup", {
 *     vpcId: fooVpc.id,
 *     securityGroupName: "acc-test-security-group2",
 * });
 * const fooCluster = new volcengine.veecp.Cluster("fooCluster", {
 *     description: "created by terraform",
 *     deleteProtectionEnabled: false,
 *     profile: "Edge",
 *     clusterConfig: {
 *         subnetIds: [fooSubnet.id],
 *         apiServerPublicAccessEnabled: true,
 *         apiServerPublicAccessConfig: {
 *             publicAccessNetworkConfig: {
 *                 billingType: "PostPaidByBandwidth",
 *                 bandwidth: 1,
 *             },
 *         },
 *         resourcePublicAccessDefaultEnabled: true,
 *     },
 *     podsConfig: {
 *         podNetworkMode: "Flannel",
 *         flannelConfig: {
 *             podCidrs: ["172.22.224.0/20"],
 *             maxPodsPerNode: 64,
 *         },
 *     },
 *     servicesConfig: {
 *         serviceCidrsv4s: ["172.30.0.0/18"],
 *     },
 * });
 * const fooNodePool = new volcengine.veecp.NodePool("fooNodePool", {
 *     clusterId: fooCluster.id,
 *     clientToken: "FGAHIxa23412FGAIOHioj",
 *     autoScaling: {
 *         enabled: true,
 *         minReplicas: 0,
 *         maxReplicas: 5,
 *         desiredReplicas: 0,
 *         priority: 5,
 *         subnetPolicy: "ZoneBalance",
 *     },
 *     nodeConfig: {
 *         instanceTypeIds: ["ecs.c1ie.xlarge"],
 *         subnetIds: [fooSubnet.id],
 *         imageId: "",
 *         systemVolume: {
 *             type: "ESSD_PL0",
 *             size: 80,
 *         },
 *         dataVolumes: [
 *             {
 *                 type: "ESSD_PL0",
 *                 size: 80,
 *                 mountPoint: "/tf1",
 *             },
 *             {
 *                 type: "ESSD_PL0",
 *                 size: 60,
 *                 mountPoint: "/tf2",
 *             },
 *         ],
 *         initializeScript: "ZWNobyBoZWxsbyB0ZXJyYWZvcm0h",
 *         security: {
 *             login: {
 *                 password: "UHdkMTIzNDU2",
 *             },
 *             securityStrategies: ["Hids"],
 *             securityGroupIds: [fooSecurityGroup.id],
 *         },
 *         additionalContainerStorageEnabled: false,
 *         instanceChargeType: "PostPaid",
 *         namePrefix: "acc-test",
 *         ecsTags: [{
 *             key: "ecs_k1",
 *             value: "ecs_v1",
 *         }],
 *     },
 *     kubernetesConfig: {
 *         labels: [{
 *             key: "label1",
 *             value: "value1",
 *         }],
 *         taints: [{
 *             key: "taint-key/node-type",
 *             value: "taint-value",
 *             effect: "NoSchedule",
 *         }],
 *         cordon: true,
 *     },
 * });
 * const fooNodePools = volcengine.veecp.getNodePoolsOutput({
 *     ids: [fooNodePool.id],
 * });
 * ```
 */
export function getNodePools(args?: GetNodePoolsArgs, opts?: pulumi.InvokeOptions): Promise<GetNodePoolsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:veecp/getNodePools:getNodePools", {
        "autoScalingEnabled": args.autoScalingEnabled,
        "clusterId": args.clusterId,
        "clusterIds": args.clusterIds,
        "createClientToken": args.createClientToken,
        "ids": args.ids,
        "name": args.name,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "statuses": args.statuses,
        "updateClientToken": args.updateClientToken,
    }, opts);
}

/**
 * A collection of arguments for invoking getNodePools.
 */
export interface GetNodePoolsArgs {
    /**
     * Is enabled of AutoScaling.
     */
    autoScalingEnabled?: boolean;
    /**
     * The ClusterId of NodePool.
     */
    clusterId?: string;
    /**
     * The ClusterIds of NodePool IDs.
     */
    clusterIds?: string[];
    /**
     * The ClientToken when successfully created.
     */
    createClientToken?: string;
    /**
     * The IDs of NodePool.
     */
    ids?: string[];
    /**
     * The Name of NodePool.
     */
    name?: string;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The Status of NodePool.
     */
    statuses?: inputs.veecp.GetNodePoolsStatus[];
    /**
     * The ClientToken when last update was successful.
     */
    updateClientToken?: string;
}

/**
 * A collection of values returned by getNodePools.
 */
export interface GetNodePoolsResult {
    readonly autoScalingEnabled?: boolean;
    /**
     * The ClusterId of NodePool.
     */
    readonly clusterId?: string;
    readonly clusterIds?: string[];
    /**
     * The ClientToken when successfully created.
     */
    readonly createClientToken?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The Name of NodePool.
     */
    readonly name?: string;
    readonly nameRegex?: string;
    /**
     * The collection of NodePools query.
     */
    readonly nodePools: outputs.veecp.GetNodePoolsNodePool[];
    readonly outputFile?: string;
    readonly statuses?: outputs.veecp.GetNodePoolsStatus[];
    /**
     * The total count of query.
     */
    readonly totalCount: number;
    /**
     * The ClientToken when last update was successful.
     */
    readonly updateClientToken?: string;
}
/**
 * Use this data source to query detailed information of veecp node pools
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooZones = volcengine.ecs.getZones({});
 * const fooVpc = new volcengine.vpc.Vpc("fooVpc", {
 *     vpcName: "acc-test-project1",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const fooSubnet = new volcengine.vpc.Subnet("fooSubnet", {
 *     subnetName: "acc-subnet-test-2",
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *     vpcId: fooVpc.id,
 * });
 * const fooSecurityGroup = new volcengine.vpc.SecurityGroup("fooSecurityGroup", {
 *     vpcId: fooVpc.id,
 *     securityGroupName: "acc-test-security-group2",
 * });
 * const fooCluster = new volcengine.veecp.Cluster("fooCluster", {
 *     description: "created by terraform",
 *     deleteProtectionEnabled: false,
 *     profile: "Edge",
 *     clusterConfig: {
 *         subnetIds: [fooSubnet.id],
 *         apiServerPublicAccessEnabled: true,
 *         apiServerPublicAccessConfig: {
 *             publicAccessNetworkConfig: {
 *                 billingType: "PostPaidByBandwidth",
 *                 bandwidth: 1,
 *             },
 *         },
 *         resourcePublicAccessDefaultEnabled: true,
 *     },
 *     podsConfig: {
 *         podNetworkMode: "Flannel",
 *         flannelConfig: {
 *             podCidrs: ["172.22.224.0/20"],
 *             maxPodsPerNode: 64,
 *         },
 *     },
 *     servicesConfig: {
 *         serviceCidrsv4s: ["172.30.0.0/18"],
 *     },
 * });
 * const fooNodePool = new volcengine.veecp.NodePool("fooNodePool", {
 *     clusterId: fooCluster.id,
 *     clientToken: "FGAHIxa23412FGAIOHioj",
 *     autoScaling: {
 *         enabled: true,
 *         minReplicas: 0,
 *         maxReplicas: 5,
 *         desiredReplicas: 0,
 *         priority: 5,
 *         subnetPolicy: "ZoneBalance",
 *     },
 *     nodeConfig: {
 *         instanceTypeIds: ["ecs.c1ie.xlarge"],
 *         subnetIds: [fooSubnet.id],
 *         imageId: "",
 *         systemVolume: {
 *             type: "ESSD_PL0",
 *             size: 80,
 *         },
 *         dataVolumes: [
 *             {
 *                 type: "ESSD_PL0",
 *                 size: 80,
 *                 mountPoint: "/tf1",
 *             },
 *             {
 *                 type: "ESSD_PL0",
 *                 size: 60,
 *                 mountPoint: "/tf2",
 *             },
 *         ],
 *         initializeScript: "ZWNobyBoZWxsbyB0ZXJyYWZvcm0h",
 *         security: {
 *             login: {
 *                 password: "UHdkMTIzNDU2",
 *             },
 *             securityStrategies: ["Hids"],
 *             securityGroupIds: [fooSecurityGroup.id],
 *         },
 *         additionalContainerStorageEnabled: false,
 *         instanceChargeType: "PostPaid",
 *         namePrefix: "acc-test",
 *         ecsTags: [{
 *             key: "ecs_k1",
 *             value: "ecs_v1",
 *         }],
 *     },
 *     kubernetesConfig: {
 *         labels: [{
 *             key: "label1",
 *             value: "value1",
 *         }],
 *         taints: [{
 *             key: "taint-key/node-type",
 *             value: "taint-value",
 *             effect: "NoSchedule",
 *         }],
 *         cordon: true,
 *     },
 * });
 * const fooNodePools = volcengine.veecp.getNodePoolsOutput({
 *     ids: [fooNodePool.id],
 * });
 * ```
 */
export function getNodePoolsOutput(args?: GetNodePoolsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNodePoolsResult> {
    return pulumi.output(args).apply((a: any) => getNodePools(a, opts))
}

/**
 * A collection of arguments for invoking getNodePools.
 */
export interface GetNodePoolsOutputArgs {
    /**
     * Is enabled of AutoScaling.
     */
    autoScalingEnabled?: pulumi.Input<boolean>;
    /**
     * The ClusterId of NodePool.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The ClusterIds of NodePool IDs.
     */
    clusterIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ClientToken when successfully created.
     */
    createClientToken?: pulumi.Input<string>;
    /**
     * The IDs of NodePool.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Name of NodePool.
     */
    name?: pulumi.Input<string>;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The Status of NodePool.
     */
    statuses?: pulumi.Input<pulumi.Input<inputs.veecp.GetNodePoolsStatusArgs>[]>;
    /**
     * The ClientToken when last update was successful.
     */
    updateClientToken?: pulumi.Input<string>;
}
