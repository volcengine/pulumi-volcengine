// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage cdn certificate
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.cdn.CdnCertificate("foo", {
 *     certificate: "",
 *     desc: "tftest",
 *     privateKey: "",
 *     source: "cdn_cert_hosting",
 * });
 * ```
 *
 * ## Import
 *
 * CdnCertificate can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:cdn/cdnCertificate:CdnCertificate default resource_id
 * ```
 *
 *  You can delete the certificate hosted on the content delivery network. You can configure the HTTPS module to associate the certificate and domain name through the domain_config field of volcengine_cdn_domain. If the certificate to be deleted is already associated with a domain name, the deletion will fail. To remove the association between the domain name and the certificate, you can disable the HTTPS function for the domain name in the Content Delivery Network console.
 */
export class CdnCertificate extends pulumi.CustomResource {
    /**
     * Get an existing CdnCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CdnCertificateState, opts?: pulumi.CustomResourceOptions): CdnCertificate {
        return new CdnCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:cdn/cdnCertificate:CdnCertificate';

    /**
     * Returns true if the given object is an instance of CdnCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CdnCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CdnCertificate.__pulumiType;
    }

    /**
     * Content of the specified certificate public key file. Line breaks in the content should be replaced with `\r\n`. The file extension for the certificate public key is `.crt` or `.pem`. The public key must include the complete certificate chain. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly certificate!: pulumi.Output<string>;
    /**
     * Note on the certificate.
     */
    public readonly desc!: pulumi.Output<string>;
    /**
     * The content of the specified certificate private key file. Replace line breaks in the content with `\r\n`. The file extension for the certificate private key is `.key` or `.pem`. The private key must be unencrypted. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly privateKey!: pulumi.Output<string>;
    /**
     * Specify the location for storing the certificate. The parameter can take the following values: `volcCertCenter`: indicates that the certificate will be stored in the certificate center.`cdnCertHosting`: indicates that the certificate will be hosted on the content delivery network.
     */
    public readonly source!: pulumi.Output<string>;

    /**
     * Create a CdnCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CdnCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CdnCertificateArgs | CdnCertificateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CdnCertificateState | undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["desc"] = state ? state.desc : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
        } else {
            const args = argsOrState as CdnCertificateArgs | undefined;
            if ((!args || args.certificate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificate'");
            }
            if ((!args || args.desc === undefined) && !opts.urn) {
                throw new Error("Missing required property 'desc'");
            }
            if ((!args || args.privateKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateKey'");
            }
            if ((!args || args.source === undefined) && !opts.urn) {
                throw new Error("Missing required property 'source'");
            }
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["desc"] = args ? args.desc : undefined;
            resourceInputs["privateKey"] = args ? args.privateKey : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CdnCertificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CdnCertificate resources.
 */
export interface CdnCertificateState {
    /**
     * Content of the specified certificate public key file. Line breaks in the content should be replaced with `\r\n`. The file extension for the certificate public key is `.crt` or `.pem`. The public key must include the complete certificate chain. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Note on the certificate.
     */
    desc?: pulumi.Input<string>;
    /**
     * The content of the specified certificate private key file. Replace line breaks in the content with `\r\n`. The file extension for the certificate private key is `.key` or `.pem`. The private key must be unencrypted. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * Specify the location for storing the certificate. The parameter can take the following values: `volcCertCenter`: indicates that the certificate will be stored in the certificate center.`cdnCertHosting`: indicates that the certificate will be hosted on the content delivery network.
     */
    source?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CdnCertificate resource.
 */
export interface CdnCertificateArgs {
    /**
     * Content of the specified certificate public key file. Line breaks in the content should be replaced with `\r\n`. The file extension for the certificate public key is `.crt` or `.pem`. The public key must include the complete certificate chain. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    certificate: pulumi.Input<string>;
    /**
     * Note on the certificate.
     */
    desc: pulumi.Input<string>;
    /**
     * The content of the specified certificate private key file. Replace line breaks in the content with `\r\n`. The file extension for the certificate private key is `.key` or `.pem`. The private key must be unencrypted. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    privateKey: pulumi.Input<string>;
    /**
     * Specify the location for storing the certificate. The parameter can take the following values: `volcCertCenter`: indicates that the certificate will be stored in the certificate center.`cdnCertHosting`: indicates that the certificate will be hosted on the content delivery network.
     */
    source: pulumi.Input<string>;
}