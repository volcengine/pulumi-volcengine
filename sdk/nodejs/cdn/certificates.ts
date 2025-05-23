// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cdn certificates
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooCdnCertificate = new volcengine.cdn.CdnCertificate("fooCdnCertificate", {
 *     certificate: "",
 *     privateKey: "",
 *     desc: "tftest",
 *     source: "cdn_cert_hosting",
 * });
 * const fooCertificates = volcengine.cdn.getCertificatesOutput({
 *     source: fooCdnCertificate.source,
 * });
 * ```
 */
/** @deprecated volcengine.cdn.Certificates has been deprecated in favor of volcengine.cdn.getCertificates */
export function certificates(args: CertificatesArgs, opts?: pulumi.InvokeOptions): Promise<CertificatesResult> {
    pulumi.log.warn("certificates is deprecated: volcengine.cdn.Certificates has been deprecated in favor of volcengine.cdn.getCertificates")

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:cdn/certificates:Certificates", {
        "name": args.name,
        "outputFile": args.outputFile,
        "source": args.source,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking Certificates.
 */
export interface CertificatesArgs {
    /**
     * Specify a domain to obtain certificates that include that domain in the SAN field. The domain can be a wildcard domain. For example, specifying *.example.com will obtain certificates that include img.example.com or www.example.com in the SAN field.
     */
    name?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * Specify the location for storing the certificate. The parameter can take the following values: `volcCertCenter`: indicates that the certificate will be stored in the certificate center.`cdnCertHosting`: indicates that the certificate will be hosted on the content delivery network.
     */
    source: string;
    /**
     * Specify one or more states to retrieve certificates in those states. By default, all certificates in all states are returned. You can specify the following states. Multiple states are separated by commas. running: Retrieves certificates with a validity period greater than 30 days. expired: Retrieves certificates that have already expired. expiring_soon: Retrieves certificates with a validity period less than or equal to 30 days but have not yet expired.
     */
    status?: string;
}

/**
 * A collection of values returned by Certificates.
 */
export interface CertificatesResult {
    /**
     * The collection of query.
     */
    readonly certInfos: outputs.cdn.CertificatesCertInfo[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    readonly outputFile?: string;
    /**
     * Specify the location for storing the certificate. The parameter can take the following values: `volcCertCenter`: indicates that the certificate will be stored in the certificate center.`cdnCertHosting`: indicates that the certificate will be hosted on the content delivery network.
     */
    readonly source: string;
    /**
     * Specify one or more states to retrieve certificates in those states. By default, all certificates in all states are returned. You can specify the following states. Multiple states are separated by commas. running: Retrieves certificates with a validity period greater than 30 days. expired: Retrieves certificates that have already expired. expiring_soon: Retrieves certificates with a validity period less than or equal to 30 days but have not yet expired.
     */
    readonly status?: string;
    /**
     * The total count of query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of cdn certificates
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooCdnCertificate = new volcengine.cdn.CdnCertificate("fooCdnCertificate", {
 *     certificate: "",
 *     privateKey: "",
 *     desc: "tftest",
 *     source: "cdn_cert_hosting",
 * });
 * const fooCertificates = volcengine.cdn.getCertificatesOutput({
 *     source: fooCdnCertificate.source,
 * });
 * ```
 */
/** @deprecated volcengine.cdn.Certificates has been deprecated in favor of volcengine.cdn.getCertificates */
export function certificatesOutput(args: CertificatesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<CertificatesResult> {
    return pulumi.output(args).apply((a: any) => certificates(a, opts))
}

/**
 * A collection of arguments for invoking Certificates.
 */
export interface CertificatesOutputArgs {
    /**
     * Specify a domain to obtain certificates that include that domain in the SAN field. The domain can be a wildcard domain. For example, specifying *.example.com will obtain certificates that include img.example.com or www.example.com in the SAN field.
     */
    name?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * Specify the location for storing the certificate. The parameter can take the following values: `volcCertCenter`: indicates that the certificate will be stored in the certificate center.`cdnCertHosting`: indicates that the certificate will be hosted on the content delivery network.
     */
    source: pulumi.Input<string>;
    /**
     * Specify one or more states to retrieve certificates in those states. By default, all certificates in all states are returned. You can specify the following states. Multiple states are separated by commas. running: Retrieves certificates with a validity period greater than 30 days. expired: Retrieves certificates that have already expired. expiring_soon: Retrieves certificates with a validity period less than or equal to 30 days but have not yet expired.
     */
    status?: pulumi.Input<string>;
}
