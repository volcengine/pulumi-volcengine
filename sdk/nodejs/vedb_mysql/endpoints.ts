// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vedb mysql endpoints
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
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[2]?.id),
 *     vpcId: fooVpc.id,
 * });
 * const fooInstance = new volcengine.vedb_mysql.Instance("fooInstance", {
 *     chargeType: "PostPaid",
 *     storageChargeType: "PostPaid",
 *     dbEngineVersion: "MySQL_8_0",
 *     dbMinorVersion: "3.0",
 *     nodeNumber: 2,
 *     nodeSpec: "vedb.mysql.x4.large",
 *     subnetId: fooSubnet.id,
 *     instanceName: "tf-test",
 *     projectName: "testA",
 *     tags: [
 *         {
 *             key: "tftest",
 *             value: "tftest",
 *         },
 *         {
 *             key: "tftest2",
 *             value: "tftest2",
 *         },
 *     ],
 * });
 * const fooInstances = volcengine.vedb_mysql.InstancesOutput({
 *     instanceId: fooInstance.id,
 * });
 * const fooEndpoint = new volcengine.vedb_mysql.Endpoint("fooEndpoint", {
 *     endpointType: "Custom",
 *     instanceId: fooInstance.id,
 *     nodeIds: [
 *         fooInstances.apply(fooInstances => fooInstances.instances?.[0]?.nodes?.[0]?.nodeId),
 *         fooInstances.apply(fooInstances => fooInstances.instances?.[0]?.nodes?.[1]?.nodeId),
 *     ],
 *     readWriteMode: "ReadWrite",
 *     endpointName: "tf-test",
 *     description: "tf test",
 *     masterAcceptReadRequests: true,
 *     distributedTransaction: true,
 *     consistLevel: "Session",
 *     consistTimeout: 100000,
 *     consistTimeoutAction: "ReadMaster",
 * });
 * const fooEndpoints = volcengine.vedb_mysql.EndpointsOutput({
 *     endpointId: fooEndpoint.endpointId,
 *     instanceId: fooInstance.id,
 * });
 * ```
 */
export function endpoints(args: EndpointsArgs, opts?: pulumi.InvokeOptions): Promise<EndpointsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:vedb_mysql/endpoints:Endpoints", {
        "endpointId": args.endpointId,
        "instanceId": args.instanceId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking Endpoints.
 */
export interface EndpointsArgs {
    /**
     * The id of the endpoint.
     */
    endpointId?: string;
    /**
     * The id of the instance.
     */
    instanceId: string;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by Endpoints.
 */
export interface EndpointsResult {
    /**
     * The id of the endpoint.
     */
    readonly endpointId?: string;
    /**
     * The collection of query.
     */
    readonly endpoints: outputs.vedb_mysql.EndpointsEndpoint[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of vedb mysql endpoints
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
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[2]?.id),
 *     vpcId: fooVpc.id,
 * });
 * const fooInstance = new volcengine.vedb_mysql.Instance("fooInstance", {
 *     chargeType: "PostPaid",
 *     storageChargeType: "PostPaid",
 *     dbEngineVersion: "MySQL_8_0",
 *     dbMinorVersion: "3.0",
 *     nodeNumber: 2,
 *     nodeSpec: "vedb.mysql.x4.large",
 *     subnetId: fooSubnet.id,
 *     instanceName: "tf-test",
 *     projectName: "testA",
 *     tags: [
 *         {
 *             key: "tftest",
 *             value: "tftest",
 *         },
 *         {
 *             key: "tftest2",
 *             value: "tftest2",
 *         },
 *     ],
 * });
 * const fooInstances = volcengine.vedb_mysql.InstancesOutput({
 *     instanceId: fooInstance.id,
 * });
 * const fooEndpoint = new volcengine.vedb_mysql.Endpoint("fooEndpoint", {
 *     endpointType: "Custom",
 *     instanceId: fooInstance.id,
 *     nodeIds: [
 *         fooInstances.apply(fooInstances => fooInstances.instances?.[0]?.nodes?.[0]?.nodeId),
 *         fooInstances.apply(fooInstances => fooInstances.instances?.[0]?.nodes?.[1]?.nodeId),
 *     ],
 *     readWriteMode: "ReadWrite",
 *     endpointName: "tf-test",
 *     description: "tf test",
 *     masterAcceptReadRequests: true,
 *     distributedTransaction: true,
 *     consistLevel: "Session",
 *     consistTimeout: 100000,
 *     consistTimeoutAction: "ReadMaster",
 * });
 * const fooEndpoints = volcengine.vedb_mysql.EndpointsOutput({
 *     endpointId: fooEndpoint.endpointId,
 *     instanceId: fooInstance.id,
 * });
 * ```
 */
export function endpointsOutput(args: EndpointsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<EndpointsResult> {
    return pulumi.output(args).apply((a: any) => endpoints(a, opts))
}

/**
 * A collection of arguments for invoking Endpoints.
 */
export interface EndpointsOutputArgs {
    /**
     * The id of the endpoint.
     */
    endpointId?: pulumi.Input<string>;
    /**
     * The id of the instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * A Name Regex of Resource.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}