// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage cr tag
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * // Tag cannot be created,please import by command `terraform import volcengine_cr_tag.default registry:namespace:repository:tag`
 * const defaultTag = new volcengine.cr.Tag("default", {
 *     namespace: "langyu",
 *     registry: "enterprise-1",
 *     repository: "repo",
 * });
 * ```
 *
 * ## Import
 *
 * CR tags can be imported using the registry:namespace:repository:tag, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:cr/tag:Tag default cr-basic:namespace-1:repo-1:v1
 * ```
 */
export class Tag extends pulumi.CustomResource {
    /**
     * Get an existing Tag resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TagState, opts?: pulumi.CustomResourceOptions): Tag {
        return new Tag(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:cr/tag:Tag';

    /**
     * Returns true if the given object is an instance of Tag.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Tag {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Tag.__pulumiType;
    }

    /**
     * The chart attribute,valid when tag type is Chart.
     */
    public /*out*/ readonly chartAttributes!: pulumi.Output<outputs.cr.TagChartAttribute[]>;
    /**
     * The digest of image.
     */
    public /*out*/ readonly digest!: pulumi.Output<string>;
    /**
     * The list of image attributes,valid when tag type is Image.
     */
    public /*out*/ readonly imageAttributes!: pulumi.Output<outputs.cr.TagImageAttribute[]>;
    /**
     * The name of OCI product.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The target namespace name.
     */
    public readonly namespace!: pulumi.Output<string>;
    /**
     * The last push time of OCI product.
     */
    public /*out*/ readonly pushTime!: pulumi.Output<string>;
    /**
     * The CrRegistry name.
     */
    public readonly registry!: pulumi.Output<string>;
    /**
     * The name of repository.
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * The size of OCI product.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * The type of OCI product tag.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Tag resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TagArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TagArgs | TagState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TagState | undefined;
            resourceInputs["chartAttributes"] = state ? state.chartAttributes : undefined;
            resourceInputs["digest"] = state ? state.digest : undefined;
            resourceInputs["imageAttributes"] = state ? state.imageAttributes : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["pushTime"] = state ? state.pushTime : undefined;
            resourceInputs["registry"] = state ? state.registry : undefined;
            resourceInputs["repository"] = state ? state.repository : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as TagArgs | undefined;
            if ((!args || args.namespace === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespace'");
            }
            if ((!args || args.registry === undefined) && !opts.urn) {
                throw new Error("Missing required property 'registry'");
            }
            if ((!args || args.repository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repository'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["registry"] = args ? args.registry : undefined;
            resourceInputs["repository"] = args ? args.repository : undefined;
            resourceInputs["chartAttributes"] = undefined /*out*/;
            resourceInputs["digest"] = undefined /*out*/;
            resourceInputs["imageAttributes"] = undefined /*out*/;
            resourceInputs["pushTime"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Tag.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Tag resources.
 */
export interface TagState {
    /**
     * The chart attribute,valid when tag type is Chart.
     */
    chartAttributes?: pulumi.Input<pulumi.Input<inputs.cr.TagChartAttribute>[]>;
    /**
     * The digest of image.
     */
    digest?: pulumi.Input<string>;
    /**
     * The list of image attributes,valid when tag type is Image.
     */
    imageAttributes?: pulumi.Input<pulumi.Input<inputs.cr.TagImageAttribute>[]>;
    /**
     * The name of OCI product.
     */
    name?: pulumi.Input<string>;
    /**
     * The target namespace name.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The last push time of OCI product.
     */
    pushTime?: pulumi.Input<string>;
    /**
     * The CrRegistry name.
     */
    registry?: pulumi.Input<string>;
    /**
     * The name of repository.
     */
    repository?: pulumi.Input<string>;
    /**
     * The size of OCI product.
     */
    size?: pulumi.Input<number>;
    /**
     * The type of OCI product tag.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Tag resource.
 */
export interface TagArgs {
    /**
     * The name of OCI product.
     */
    name?: pulumi.Input<string>;
    /**
     * The target namespace name.
     */
    namespace: pulumi.Input<string>;
    /**
     * The CrRegistry name.
     */
    registry: pulumi.Input<string>;
    /**
     * The name of repository.
     */
    repository: pulumi.Input<string>;
}