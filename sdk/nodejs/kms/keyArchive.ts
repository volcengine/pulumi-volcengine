// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage kms key archive
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooKeyring = new volcengine.kms.Keyring("fooKeyring", {
 *     keyringName: "tf-test",
 *     description: "tf-test",
 *     projectName: "default",
 * });
 * const fooKey = new volcengine.kms.Key("fooKey", {
 *     keyringName: fooKeyring.keyringName,
 *     keyName: "mrk-tf-key-mod",
 *     description: "tf test key-mod",
 *     tags: [{
 *         key: "tfkey3",
 *         value: "tfvalue3",
 *     }],
 * });
 * const fooKeyArchive = new volcengine.kms.KeyArchive("fooKeyArchive", {keyId: fooKey.id});
 * ```
 *
 * ## Import
 *
 * KmsKeyArchive can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:kms/keyArchive:KeyArchive default resource_id
 * ```
 *
 * or
 *
 * ```sh
 * $ pulumi import volcengine:kms/keyArchive:KeyArchive default key_name:keyring_name
 * ```
 */
export class KeyArchive extends pulumi.CustomResource {
    /**
     * Get an existing KeyArchive resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeyArchiveState, opts?: pulumi.CustomResourceOptions): KeyArchive {
        return new KeyArchive(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:kms/keyArchive:KeyArchive';

    /**
     * Returns true if the given object is an instance of KeyArchive.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KeyArchive {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KeyArchive.__pulumiType;
    }

    /**
     * The id of the CMK.
     */
    public readonly keyId!: pulumi.Output<string>;
    /**
     * The name of the CMK.
     */
    public readonly keyName!: pulumi.Output<string>;
    /**
     * The state of the key.
     */
    public /*out*/ readonly keyState!: pulumi.Output<string>;
    /**
     * The name of the keyring.
     */
    public readonly keyringName!: pulumi.Output<string | undefined>;

    /**
     * Create a KeyArchive resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: KeyArchiveArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyArchiveArgs | KeyArchiveState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KeyArchiveState | undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
            resourceInputs["keyName"] = state ? state.keyName : undefined;
            resourceInputs["keyState"] = state ? state.keyState : undefined;
            resourceInputs["keyringName"] = state ? state.keyringName : undefined;
        } else {
            const args = argsOrState as KeyArchiveArgs | undefined;
            resourceInputs["keyId"] = args ? args.keyId : undefined;
            resourceInputs["keyName"] = args ? args.keyName : undefined;
            resourceInputs["keyringName"] = args ? args.keyringName : undefined;
            resourceInputs["keyState"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KeyArchive.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KeyArchive resources.
 */
export interface KeyArchiveState {
    /**
     * The id of the CMK.
     */
    keyId?: pulumi.Input<string>;
    /**
     * The name of the CMK.
     */
    keyName?: pulumi.Input<string>;
    /**
     * The state of the key.
     */
    keyState?: pulumi.Input<string>;
    /**
     * The name of the keyring.
     */
    keyringName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KeyArchive resource.
 */
export interface KeyArchiveArgs {
    /**
     * The id of the CMK.
     */
    keyId?: pulumi.Input<string>;
    /**
     * The name of the CMK.
     */
    keyName?: pulumi.Input<string>;
    /**
     * The name of the keyring.
     */
    keyringName?: pulumi.Input<string>;
}
