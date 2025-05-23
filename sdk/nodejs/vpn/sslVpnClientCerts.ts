// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of ssl vpn client certs
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
 * const fooGateway = new volcengine.vpn.Gateway("fooGateway", {
 *     vpcId: fooVpc.id,
 *     subnetId: fooSubnet.id,
 *     bandwidth: 5,
 *     vpnGatewayName: "acc-test1",
 *     description: "acc-test1",
 *     period: 7,
 *     projectName: "default",
 *     sslEnabled: true,
 *     sslMaxConnections: 5,
 * });
 * const fooSslVpnServer = new volcengine.vpn.SslVpnServer("fooSslVpnServer", {
 *     vpnGatewayId: fooGateway.id,
 *     localSubnets: [fooSubnet.cidrBlock],
 *     clientIpPool: "172.16.2.0/24",
 *     sslVpnServerName: "acc-test-ssl",
 *     description: "acc-test",
 *     protocol: "UDP",
 *     cipher: "AES-128-CBC",
 *     auth: "SHA1",
 *     compress: true,
 * });
 * const fooSslVpnClientCert: volcengine.vpn.SslVpnClientCert[] = [];
 * for (const range = {value: 0}; range.value < 5; range.value++) {
 *     fooSslVpnClientCert.push(new volcengine.vpn.SslVpnClientCert(`fooSslVpnClientCert-${range.value}`, {
 *         sslVpnServerId: fooSslVpnServer.id,
 *         sslVpnClientCertName: `acc-test-client-cert-${range.value}`,
 *         description: "acc-test",
 *     }));
 * }
 * const fooSslVpnClientCerts = volcengine.vpn.getSslVpnClientCertsOutput({
 *     ids: fooSslVpnClientCert.map(__item => __item.id),
 * });
 * ```
 */
/** @deprecated volcengine.vpn.SslVpnClientCerts has been deprecated in favor of volcengine.vpn.getSslVpnClientCerts */
export function sslVpnClientCerts(args?: SslVpnClientCertsArgs, opts?: pulumi.InvokeOptions): Promise<SslVpnClientCertsResult> {
    pulumi.log.warn("sslVpnClientCerts is deprecated: volcengine.vpn.SslVpnClientCerts has been deprecated in favor of volcengine.vpn.getSslVpnClientCerts")
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:vpn/sslVpnClientCerts:SslVpnClientCerts", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "sslVpnClientCertName": args.sslVpnClientCertName,
        "sslVpnServerId": args.sslVpnServerId,
    }, opts);
}

/**
 * A collection of arguments for invoking SslVpnClientCerts.
 */
export interface SslVpnClientCertsArgs {
    /**
     * The ids list of ssl vpn client cert.
     */
    ids?: string[];
    /**
     * A Name Regex of ssl vpn client cert.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The name of the ssl vpn client cert.
     */
    sslVpnClientCertName?: string;
    /**
     * The id of the ssl vpn server.
     */
    sslVpnServerId?: string;
}

/**
 * A collection of values returned by SslVpnClientCerts.
 */
export interface SslVpnClientCertsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The name of the ssl vpn client cert.
     */
    readonly sslVpnClientCertName?: string;
    /**
     * The collection of of ssl vpn client certs.
     */
    readonly sslVpnClientCerts: outputs.vpn.SslVpnClientCertsSslVpnClientCert[];
    /**
     * The id of the ssl vpn server.
     */
    readonly sslVpnServerId?: string;
    /**
     * The total count of ssl vpn client cert query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of ssl vpn client certs
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
 * const fooGateway = new volcengine.vpn.Gateway("fooGateway", {
 *     vpcId: fooVpc.id,
 *     subnetId: fooSubnet.id,
 *     bandwidth: 5,
 *     vpnGatewayName: "acc-test1",
 *     description: "acc-test1",
 *     period: 7,
 *     projectName: "default",
 *     sslEnabled: true,
 *     sslMaxConnections: 5,
 * });
 * const fooSslVpnServer = new volcengine.vpn.SslVpnServer("fooSslVpnServer", {
 *     vpnGatewayId: fooGateway.id,
 *     localSubnets: [fooSubnet.cidrBlock],
 *     clientIpPool: "172.16.2.0/24",
 *     sslVpnServerName: "acc-test-ssl",
 *     description: "acc-test",
 *     protocol: "UDP",
 *     cipher: "AES-128-CBC",
 *     auth: "SHA1",
 *     compress: true,
 * });
 * const fooSslVpnClientCert: volcengine.vpn.SslVpnClientCert[] = [];
 * for (const range = {value: 0}; range.value < 5; range.value++) {
 *     fooSslVpnClientCert.push(new volcengine.vpn.SslVpnClientCert(`fooSslVpnClientCert-${range.value}`, {
 *         sslVpnServerId: fooSslVpnServer.id,
 *         sslVpnClientCertName: `acc-test-client-cert-${range.value}`,
 *         description: "acc-test",
 *     }));
 * }
 * const fooSslVpnClientCerts = volcengine.vpn.getSslVpnClientCertsOutput({
 *     ids: fooSslVpnClientCert.map(__item => __item.id),
 * });
 * ```
 */
/** @deprecated volcengine.vpn.SslVpnClientCerts has been deprecated in favor of volcengine.vpn.getSslVpnClientCerts */
export function sslVpnClientCertsOutput(args?: SslVpnClientCertsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<SslVpnClientCertsResult> {
    return pulumi.output(args).apply((a: any) => sslVpnClientCerts(a, opts))
}

/**
 * A collection of arguments for invoking SslVpnClientCerts.
 */
export interface SslVpnClientCertsOutputArgs {
    /**
     * The ids list of ssl vpn client cert.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Name Regex of ssl vpn client cert.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The name of the ssl vpn client cert.
     */
    sslVpnClientCertName?: pulumi.Input<string>;
    /**
     * The id of the ssl vpn server.
     */
    sslVpnServerId?: pulumi.Input<string>;
}
