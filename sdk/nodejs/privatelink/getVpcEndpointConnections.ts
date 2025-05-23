// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of privatelink vpc endpoint connections
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
 * const fooClb = new volcengine.clb.Clb("fooClb", {
 *     type: "public",
 *     subnetId: fooSubnet.id,
 *     loadBalancerSpec: "small_1",
 *     description: "acc-test-demo",
 *     loadBalancerName: "acc-test-clb",
 *     loadBalancerBillingType: "PostPaid",
 *     eipBillingConfig: {
 *         isp: "BGP",
 *         eipBillingType: "PostPaidByBandwidth",
 *         bandwidth: 1,
 *     },
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * const fooVpcEndpointService = new volcengine.privatelink.VpcEndpointService("fooVpcEndpointService", {
 *     resources: [{
 *         resourceId: fooClb.id,
 *         resourceType: "CLB",
 *     }],
 *     description: "acc-test",
 * });
 * const fooVpcEndpoint = new volcengine.privatelink.VpcEndpoint("fooVpcEndpoint", {
 *     securityGroupIds: [fooSecurityGroup.id],
 *     serviceId: fooVpcEndpointService.id,
 *     endpointName: "acc-test-ep",
 *     description: "acc-test",
 * });
 * const fooVpcEndpointZone = new volcengine.privatelink.VpcEndpointZone("fooVpcEndpointZone", {
 *     endpointId: fooVpcEndpoint.id,
 *     subnetId: fooSubnet.id,
 *     privateIpAddress: "172.16.0.251",
 * });
 * const fooVpcEndpointConnection = new volcengine.privatelink.VpcEndpointConnection("fooVpcEndpointConnection", {
 *     endpointId: fooVpcEndpoint.id,
 *     serviceId: fooVpcEndpointService.id,
 * }, {
 *     dependsOn: [fooVpcEndpointZone],
 * });
 * const fooVpcEndpointConnections = volcengine.privatelink.getVpcEndpointConnectionsOutput({
 *     endpointId: fooVpcEndpointConnection.endpointId,
 *     serviceId: fooVpcEndpointConnection.serviceId,
 * });
 * ```
 */
export function getVpcEndpointConnections(args: GetVpcEndpointConnectionsArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcEndpointConnectionsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:privatelink/getVpcEndpointConnections:getVpcEndpointConnections", {
        "endpointId": args.endpointId,
        "endpointOwnerAccountId": args.endpointOwnerAccountId,
        "outputFile": args.outputFile,
        "serviceId": args.serviceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcEndpointConnections.
 */
export interface GetVpcEndpointConnectionsArgs {
    /**
     * The id of the vpc endpoint.
     */
    endpointId?: string;
    /**
     * The account id of the vpc endpoint.
     */
    endpointOwnerAccountId?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The id of the vpc endpoint service.
     */
    serviceId: string;
}

/**
 * A collection of values returned by getVpcEndpointConnections.
 */
export interface GetVpcEndpointConnectionsResult {
    /**
     * The list of query.
     */
    readonly connections: outputs.privatelink.GetVpcEndpointConnectionsConnection[];
    /**
     * The id of the vpc endpoint.
     */
    readonly endpointId?: string;
    /**
     * The account id of the vpc endpoint.
     */
    readonly endpointOwnerAccountId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
    /**
     * The id of the vpc endpoint service.
     */
    readonly serviceId: string;
    /**
     * Returns the total amount of the data list.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of privatelink vpc endpoint connections
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
 * const fooClb = new volcengine.clb.Clb("fooClb", {
 *     type: "public",
 *     subnetId: fooSubnet.id,
 *     loadBalancerSpec: "small_1",
 *     description: "acc-test-demo",
 *     loadBalancerName: "acc-test-clb",
 *     loadBalancerBillingType: "PostPaid",
 *     eipBillingConfig: {
 *         isp: "BGP",
 *         eipBillingType: "PostPaidByBandwidth",
 *         bandwidth: 1,
 *     },
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * const fooVpcEndpointService = new volcengine.privatelink.VpcEndpointService("fooVpcEndpointService", {
 *     resources: [{
 *         resourceId: fooClb.id,
 *         resourceType: "CLB",
 *     }],
 *     description: "acc-test",
 * });
 * const fooVpcEndpoint = new volcengine.privatelink.VpcEndpoint("fooVpcEndpoint", {
 *     securityGroupIds: [fooSecurityGroup.id],
 *     serviceId: fooVpcEndpointService.id,
 *     endpointName: "acc-test-ep",
 *     description: "acc-test",
 * });
 * const fooVpcEndpointZone = new volcengine.privatelink.VpcEndpointZone("fooVpcEndpointZone", {
 *     endpointId: fooVpcEndpoint.id,
 *     subnetId: fooSubnet.id,
 *     privateIpAddress: "172.16.0.251",
 * });
 * const fooVpcEndpointConnection = new volcengine.privatelink.VpcEndpointConnection("fooVpcEndpointConnection", {
 *     endpointId: fooVpcEndpoint.id,
 *     serviceId: fooVpcEndpointService.id,
 * }, {
 *     dependsOn: [fooVpcEndpointZone],
 * });
 * const fooVpcEndpointConnections = volcengine.privatelink.getVpcEndpointConnectionsOutput({
 *     endpointId: fooVpcEndpointConnection.endpointId,
 *     serviceId: fooVpcEndpointConnection.serviceId,
 * });
 * ```
 */
export function getVpcEndpointConnectionsOutput(args: GetVpcEndpointConnectionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcEndpointConnectionsResult> {
    return pulumi.output(args).apply((a: any) => getVpcEndpointConnections(a, opts))
}

/**
 * A collection of arguments for invoking getVpcEndpointConnections.
 */
export interface GetVpcEndpointConnectionsOutputArgs {
    /**
     * The id of the vpc endpoint.
     */
    endpointId?: pulumi.Input<string>;
    /**
     * The account id of the vpc endpoint.
     */
    endpointOwnerAccountId?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The id of the vpc endpoint service.
     */
    serviceId: pulumi.Input<string>;
}
