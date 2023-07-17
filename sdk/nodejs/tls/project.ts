// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage tls project
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = new volcengine.tls.Project("foo", {
 *     description: "tf-desc",
 *     iamProjectName: "default",
 *     projectName: "tf-test",
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Tls Project can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:tls/project:Project default e020c978-4f05-40e1-9167-0113d3ef****
 * ```
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:tls/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * The create time of the tls project.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The description of the tls project.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The IAM project name of the tls project.
     */
    public readonly iamProjectName!: pulumi.Output<string>;
    /**
     * The inner net domain of the tls project.
     */
    public /*out*/ readonly innerNetDomain!: pulumi.Output<string>;
    /**
     * The name of the tls project.
     */
    public readonly projectName!: pulumi.Output<string>;
    /**
     * Tags.
     */
    public readonly tags!: pulumi.Output<outputs.tls.ProjectTag[] | undefined>;
    /**
     * The count of topics in the tls project.
     */
    public /*out*/ readonly topicCount!: pulumi.Output<number>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["iamProjectName"] = state ? state.iamProjectName : undefined;
            resourceInputs["innerNetDomain"] = state ? state.innerNetDomain : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["topicCount"] = state ? state.topicCount : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            if ((!args || args.projectName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["iamProjectName"] = args ? args.iamProjectName : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["innerNetDomain"] = undefined /*out*/;
            resourceInputs["topicCount"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Project.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * The create time of the tls project.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The description of the tls project.
     */
    description?: pulumi.Input<string>;
    /**
     * The IAM project name of the tls project.
     */
    iamProjectName?: pulumi.Input<string>;
    /**
     * The inner net domain of the tls project.
     */
    innerNetDomain?: pulumi.Input<string>;
    /**
     * The name of the tls project.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.tls.ProjectTag>[]>;
    /**
     * The count of topics in the tls project.
     */
    topicCount?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * The description of the tls project.
     */
    description?: pulumi.Input<string>;
    /**
     * The IAM project name of the tls project.
     */
    iamProjectName?: pulumi.Input<string>;
    /**
     * The name of the tls project.
     */
    projectName: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.tls.ProjectTag>[]>;
}