// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage iam saml provider
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.iam.SamlProvider("foo", {
 *     encodedSamlMetadataDocument: "your document",
 *     samlProviderName: "terraform",
 *     ssoType: 2,
 *     status: 1,
 * });
 * ```
 *
 * ## Import
 *
 * IamSamlProvider can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:iam/samlProvider:SamlProvider default SAMLProviderName
 * ```
 */
export class SamlProvider extends pulumi.CustomResource {
    /**
     * Get an existing SamlProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SamlProviderState, opts?: pulumi.CustomResourceOptions): SamlProvider {
        return new SamlProvider(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:iam/samlProvider:SamlProvider';

    /**
     * Returns true if the given object is an instance of SamlProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SamlProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SamlProvider.__pulumiType;
    }

    /**
     * Identity provider creation time, such as 20150123T123318Z.
     */
    public /*out*/ readonly createDate!: pulumi.Output<string>;
    /**
     * The description of the SAML provider.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Metadata document, encoded in Base64.
     */
    public readonly encodedSamlMetadataDocument!: pulumi.Output<string>;
    /**
     * The name of the SAML provider.
     */
    public readonly samlProviderName!: pulumi.Output<string>;
    /**
     * SSO types, 1. Role-based SSO, 2. User-based SSO.
     */
    public readonly ssoType!: pulumi.Output<number>;
    /**
     * User SSO status, 1. Enabled, 2. Disable other console login methods after enabling, 3. Disabled, is a required field when creating user SSO.
     */
    public readonly status!: pulumi.Output<number | undefined>;
    /**
     * The format for the resource name of an identity provider is trn:iam::${accountID}:saml-provider/{$SAMLProviderName}.
     */
    public /*out*/ readonly trn!: pulumi.Output<string>;
    /**
     * Identity provider update time, such as: 20150123T123318Z.
     */
    public /*out*/ readonly updateDate!: pulumi.Output<string>;

    /**
     * Create a SamlProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SamlProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SamlProviderArgs | SamlProviderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SamlProviderState | undefined;
            resourceInputs["createDate"] = state ? state.createDate : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["encodedSamlMetadataDocument"] = state ? state.encodedSamlMetadataDocument : undefined;
            resourceInputs["samlProviderName"] = state ? state.samlProviderName : undefined;
            resourceInputs["ssoType"] = state ? state.ssoType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["trn"] = state ? state.trn : undefined;
            resourceInputs["updateDate"] = state ? state.updateDate : undefined;
        } else {
            const args = argsOrState as SamlProviderArgs | undefined;
            if ((!args || args.encodedSamlMetadataDocument === undefined) && !opts.urn) {
                throw new Error("Missing required property 'encodedSamlMetadataDocument'");
            }
            if ((!args || args.samlProviderName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'samlProviderName'");
            }
            if ((!args || args.ssoType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ssoType'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["encodedSamlMetadataDocument"] = args ? args.encodedSamlMetadataDocument : undefined;
            resourceInputs["samlProviderName"] = args ? args.samlProviderName : undefined;
            resourceInputs["ssoType"] = args ? args.ssoType : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["createDate"] = undefined /*out*/;
            resourceInputs["trn"] = undefined /*out*/;
            resourceInputs["updateDate"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SamlProvider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SamlProvider resources.
 */
export interface SamlProviderState {
    /**
     * Identity provider creation time, such as 20150123T123318Z.
     */
    createDate?: pulumi.Input<string>;
    /**
     * The description of the SAML provider.
     */
    description?: pulumi.Input<string>;
    /**
     * Metadata document, encoded in Base64.
     */
    encodedSamlMetadataDocument?: pulumi.Input<string>;
    /**
     * The name of the SAML provider.
     */
    samlProviderName?: pulumi.Input<string>;
    /**
     * SSO types, 1. Role-based SSO, 2. User-based SSO.
     */
    ssoType?: pulumi.Input<number>;
    /**
     * User SSO status, 1. Enabled, 2. Disable other console login methods after enabling, 3. Disabled, is a required field when creating user SSO.
     */
    status?: pulumi.Input<number>;
    /**
     * The format for the resource name of an identity provider is trn:iam::${accountID}:saml-provider/{$SAMLProviderName}.
     */
    trn?: pulumi.Input<string>;
    /**
     * Identity provider update time, such as: 20150123T123318Z.
     */
    updateDate?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SamlProvider resource.
 */
export interface SamlProviderArgs {
    /**
     * The description of the SAML provider.
     */
    description?: pulumi.Input<string>;
    /**
     * Metadata document, encoded in Base64.
     */
    encodedSamlMetadataDocument: pulumi.Input<string>;
    /**
     * The name of the SAML provider.
     */
    samlProviderName: pulumi.Input<string>;
    /**
     * SSO types, 1. Role-based SSO, 2. User-based SSO.
     */
    ssoType: pulumi.Input<number>;
    /**
     * User SSO status, 1. Enabled, 2. Disable other console login methods after enabling, 3. Disabled, is a required field when creating user SSO.
     */
    status?: pulumi.Input<number>;
}