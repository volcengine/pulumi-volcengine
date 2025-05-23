// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of privatelink vpc endpoint services
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
 * const fooVpcEndpointService: volcengine.privatelink.VpcEndpointService[] = [];
 * for (const range = {value: 0}; range.value < 2; range.value++) {
 *     fooVpcEndpointService.push(new volcengine.privatelink.VpcEndpointService(`fooVpcEndpointService-${range.value}`, {
 *         resources: [{
 *             resourceId: fooClb.id,
 *             resourceType: "CLB",
 *         }],
 *         description: "acc-test",
 *         autoAcceptEnabled: true,
 *     }));
 * }
 * const fooVpcEndpointServices = volcengine.privatelink.getVpcEndpointServicesOutput({
 *     ids: fooVpcEndpointService.map(__item => __item.id),
 * });
 * ```
 */
/** @deprecated volcengine.privatelink.VpcEndpointServices has been deprecated in favor of volcengine.privatelink.getVpcEndpointServices */
export function vpcEndpointServices(args?: VpcEndpointServicesArgs, opts?: pulumi.InvokeOptions): Promise<VpcEndpointServicesResult> {
    pulumi.log.warn("vpcEndpointServices is deprecated: volcengine.privatelink.VpcEndpointServices has been deprecated in favor of volcengine.privatelink.getVpcEndpointServices")
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:privatelink/vpcEndpointServices:VpcEndpointServices", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking VpcEndpointServices.
 */
export interface VpcEndpointServicesArgs {
    /**
     * The IDs of vpc endpoint service.
     */
    ids?: string[];
    /**
     * A Name Regex of vpc endpoint service.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The name of vpc endpoint service.
     */
    serviceName?: string;
}

/**
 * A collection of values returned by VpcEndpointServices.
 */
export interface VpcEndpointServicesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The name of service.
     */
    readonly serviceName?: string;
    /**
     * The collection of query.
     */
    readonly services: outputs.privatelink.VpcEndpointServicesService[];
    /**
     * Returns the total amount of the data list.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of privatelink vpc endpoint services
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
 * const fooVpcEndpointService: volcengine.privatelink.VpcEndpointService[] = [];
 * for (const range = {value: 0}; range.value < 2; range.value++) {
 *     fooVpcEndpointService.push(new volcengine.privatelink.VpcEndpointService(`fooVpcEndpointService-${range.value}`, {
 *         resources: [{
 *             resourceId: fooClb.id,
 *             resourceType: "CLB",
 *         }],
 *         description: "acc-test",
 *         autoAcceptEnabled: true,
 *     }));
 * }
 * const fooVpcEndpointServices = volcengine.privatelink.getVpcEndpointServicesOutput({
 *     ids: fooVpcEndpointService.map(__item => __item.id),
 * });
 * ```
 */
/** @deprecated volcengine.privatelink.VpcEndpointServices has been deprecated in favor of volcengine.privatelink.getVpcEndpointServices */
export function vpcEndpointServicesOutput(args?: VpcEndpointServicesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<VpcEndpointServicesResult> {
    return pulumi.output(args).apply((a: any) => vpcEndpointServices(a, opts))
}

/**
 * A collection of arguments for invoking VpcEndpointServices.
 */
export interface VpcEndpointServicesOutputArgs {
    /**
     * The IDs of vpc endpoint service.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Name Regex of vpc endpoint service.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The name of vpc endpoint service.
     */
    serviceName?: pulumi.Input<string>;
}
