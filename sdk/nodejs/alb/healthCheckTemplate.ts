// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage alb health check template
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.alb.HealthCheckTemplate("foo", {
 *     description: "acc-test3",
 *     healthCheckDomain: "test.com",
 *     healthCheckHttpCode: "http_2xx",
 *     healthCheckHttpVersion: "HTTP1.1",
 *     healthCheckInterval: 8,
 *     healthCheckMethod: "HEAD",
 *     healthCheckProtocol: "HTTP",
 *     healthCheckTemplateName: "acc-test-template-1",
 *     healthCheckTimeout: 11,
 *     healthCheckUri: "/",
 *     healthyThreshold: 2,
 *     unhealthyThreshold: 3,
 * });
 * ```
 *
 * ## Import
 *
 * AlbHealthCheckTemplate can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:alb/healthCheckTemplate:HealthCheckTemplate default hctpl-123*****432
 * ```
 */
export class HealthCheckTemplate extends pulumi.CustomResource {
    /**
     * Get an existing HealthCheckTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HealthCheckTemplateState, opts?: pulumi.CustomResourceOptions): HealthCheckTemplate {
        return new HealthCheckTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:alb/healthCheckTemplate:HealthCheckTemplate';

    /**
     * Returns true if the given object is an instance of HealthCheckTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HealthCheckTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HealthCheckTemplate.__pulumiType;
    }

    /**
     * The description of health check template.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The domain name to health check.
     */
    public readonly healthCheckDomain!: pulumi.Output<string>;
    /**
     * The normal HTTP status code for health check, the default is http_2xx, http_3xx, separated by commas.
     */
    public readonly healthCheckHttpCode!: pulumi.Output<string>;
    /**
     * The HTTP version of health check.
     */
    public readonly healthCheckHttpVersion!: pulumi.Output<string>;
    /**
     * The interval for performing health checks, the default value is 2, and the value is 1-300.
     */
    public readonly healthCheckInterval!: pulumi.Output<number>;
    /**
     * The health check method,default is `GET`, support `GET` and `HEAD`.
     */
    public readonly healthCheckMethod!: pulumi.Output<string>;
    /**
     * THe protocol of health check,only support HTTP.
     */
    public readonly healthCheckProtocol!: pulumi.Output<string>;
    /**
     * The health check template name.
     */
    public readonly healthCheckTemplateName!: pulumi.Output<string>;
    /**
     * The timeout of health check response,the default value is 2, and the value is 1-60.
     */
    public readonly healthCheckTimeout!: pulumi.Output<number>;
    /**
     * The uri to health check,default is `/`.
     */
    public readonly healthCheckUri!: pulumi.Output<string>;
    /**
     * The healthy threshold of the health check, the default is 3, the value is 2-10.
     */
    public readonly healthyThreshold!: pulumi.Output<number>;
    /**
     * The unhealthy threshold of the health check, the default is 3, the value is 2-10.
     */
    public readonly unhealthyThreshold!: pulumi.Output<number>;

    /**
     * Create a HealthCheckTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HealthCheckTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HealthCheckTemplateArgs | HealthCheckTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HealthCheckTemplateState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["healthCheckDomain"] = state ? state.healthCheckDomain : undefined;
            resourceInputs["healthCheckHttpCode"] = state ? state.healthCheckHttpCode : undefined;
            resourceInputs["healthCheckHttpVersion"] = state ? state.healthCheckHttpVersion : undefined;
            resourceInputs["healthCheckInterval"] = state ? state.healthCheckInterval : undefined;
            resourceInputs["healthCheckMethod"] = state ? state.healthCheckMethod : undefined;
            resourceInputs["healthCheckProtocol"] = state ? state.healthCheckProtocol : undefined;
            resourceInputs["healthCheckTemplateName"] = state ? state.healthCheckTemplateName : undefined;
            resourceInputs["healthCheckTimeout"] = state ? state.healthCheckTimeout : undefined;
            resourceInputs["healthCheckUri"] = state ? state.healthCheckUri : undefined;
            resourceInputs["healthyThreshold"] = state ? state.healthyThreshold : undefined;
            resourceInputs["unhealthyThreshold"] = state ? state.unhealthyThreshold : undefined;
        } else {
            const args = argsOrState as HealthCheckTemplateArgs | undefined;
            if ((!args || args.healthCheckTemplateName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'healthCheckTemplateName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["healthCheckDomain"] = args ? args.healthCheckDomain : undefined;
            resourceInputs["healthCheckHttpCode"] = args ? args.healthCheckHttpCode : undefined;
            resourceInputs["healthCheckHttpVersion"] = args ? args.healthCheckHttpVersion : undefined;
            resourceInputs["healthCheckInterval"] = args ? args.healthCheckInterval : undefined;
            resourceInputs["healthCheckMethod"] = args ? args.healthCheckMethod : undefined;
            resourceInputs["healthCheckProtocol"] = args ? args.healthCheckProtocol : undefined;
            resourceInputs["healthCheckTemplateName"] = args ? args.healthCheckTemplateName : undefined;
            resourceInputs["healthCheckTimeout"] = args ? args.healthCheckTimeout : undefined;
            resourceInputs["healthCheckUri"] = args ? args.healthCheckUri : undefined;
            resourceInputs["healthyThreshold"] = args ? args.healthyThreshold : undefined;
            resourceInputs["unhealthyThreshold"] = args ? args.unhealthyThreshold : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HealthCheckTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HealthCheckTemplate resources.
 */
export interface HealthCheckTemplateState {
    /**
     * The description of health check template.
     */
    description?: pulumi.Input<string>;
    /**
     * The domain name to health check.
     */
    healthCheckDomain?: pulumi.Input<string>;
    /**
     * The normal HTTP status code for health check, the default is http_2xx, http_3xx, separated by commas.
     */
    healthCheckHttpCode?: pulumi.Input<string>;
    /**
     * The HTTP version of health check.
     */
    healthCheckHttpVersion?: pulumi.Input<string>;
    /**
     * The interval for performing health checks, the default value is 2, and the value is 1-300.
     */
    healthCheckInterval?: pulumi.Input<number>;
    /**
     * The health check method,default is `GET`, support `GET` and `HEAD`.
     */
    healthCheckMethod?: pulumi.Input<string>;
    /**
     * THe protocol of health check,only support HTTP.
     */
    healthCheckProtocol?: pulumi.Input<string>;
    /**
     * The health check template name.
     */
    healthCheckTemplateName?: pulumi.Input<string>;
    /**
     * The timeout of health check response,the default value is 2, and the value is 1-60.
     */
    healthCheckTimeout?: pulumi.Input<number>;
    /**
     * The uri to health check,default is `/`.
     */
    healthCheckUri?: pulumi.Input<string>;
    /**
     * The healthy threshold of the health check, the default is 3, the value is 2-10.
     */
    healthyThreshold?: pulumi.Input<number>;
    /**
     * The unhealthy threshold of the health check, the default is 3, the value is 2-10.
     */
    unhealthyThreshold?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a HealthCheckTemplate resource.
 */
export interface HealthCheckTemplateArgs {
    /**
     * The description of health check template.
     */
    description?: pulumi.Input<string>;
    /**
     * The domain name to health check.
     */
    healthCheckDomain?: pulumi.Input<string>;
    /**
     * The normal HTTP status code for health check, the default is http_2xx, http_3xx, separated by commas.
     */
    healthCheckHttpCode?: pulumi.Input<string>;
    /**
     * The HTTP version of health check.
     */
    healthCheckHttpVersion?: pulumi.Input<string>;
    /**
     * The interval for performing health checks, the default value is 2, and the value is 1-300.
     */
    healthCheckInterval?: pulumi.Input<number>;
    /**
     * The health check method,default is `GET`, support `GET` and `HEAD`.
     */
    healthCheckMethod?: pulumi.Input<string>;
    /**
     * THe protocol of health check,only support HTTP.
     */
    healthCheckProtocol?: pulumi.Input<string>;
    /**
     * The health check template name.
     */
    healthCheckTemplateName: pulumi.Input<string>;
    /**
     * The timeout of health check response,the default value is 2, and the value is 1-60.
     */
    healthCheckTimeout?: pulumi.Input<number>;
    /**
     * The uri to health check,default is `/`.
     */
    healthCheckUri?: pulumi.Input<string>;
    /**
     * The healthy threshold of the health check, the default is 3, the value is 2-10.
     */
    healthyThreshold?: pulumi.Input<number>;
    /**
     * The unhealthy threshold of the health check, the default is 3, the value is 2-10.
     */
    unhealthyThreshold?: pulumi.Input<number>;
}