// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage apig upstream source
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
 * const foo1 = new volcengine.vpc.Subnet("foo1", {
 *     subnetName: "acc-test-subnet",
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *     vpcId: fooVpc.id,
 * });
 * const foo2 = new volcengine.vpc.Subnet("foo2", {
 *     subnetName: "acc-test-subnet",
 *     cidrBlock: "172.16.1.0/24",
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[1]?.id),
 *     vpcId: fooVpc.id,
 * });
 * const fooApigGateway = new volcengine.apig.ApigGateway("fooApigGateway", {
 *     type: "standard",
 *     comments: "acc-test",
 *     projectName: "default",
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 *     networkSpec: {
 *         vpcId: fooVpc.id,
 *         subnetIds: [
 *             foo1.id,
 *             foo2.id,
 *         ],
 *     },
 *     resourceSpec: {
 *         replicas: 2,
 *         instanceSpecCode: "1c2g",
 *         clbSpecCode: "small_1",
 *         publicNetworkBillingType: "bandwidth",
 *         publicNetworkBandwidth: 1,
 *         networkType: {
 *             enablePublicNetwork: true,
 *             enablePrivateNetwork: true,
 *         },
 *     },
 *     logSpec: {
 *         enable: true,
 *         projectId: "d3cb87c0-faeb-4074-b1ee-9bd747865a76",
 *         topicId: "d339482e-d86d-4bd8-a9bb-f270417f00a1",
 *     },
 *     monitorSpec: {
 *         enable: true,
 *         workspaceId: "4ed1caf3-279d-4c5f-8301-87ea38e92ffc",
 *     },
 * });
 * const fooApigGatewayService = new volcengine.apig.ApigGatewayService("fooApigGatewayService", {
 *     gatewayId: fooApigGateway.id,
 *     serviceName: "acc-test-apig-service",
 *     comments: "acc-test",
 *     protocols: [
 *         "HTTP",
 *         "HTTPS",
 *     ],
 *     authSpec: {
 *         enable: false,
 *     },
 * });
 * const foo_nacos = new volcengine.apig.ApigUpstreamSource("foo-nacos", {
 *     gatewayId: fooApigGateway.id,
 *     comments: "acc-test-nacos",
 *     sourceType: "Nacos",
 *     sourceSpec: {
 *         nacosSource: {
 *             nacosId: "nd197ls631meck48imm7g",
 *             authConfig: {
 *                 basic: {
 *                     username: "nacos",
 *                     password: "******",
 *                 },
 *             },
 *         },
 *     },
 * });
 * const foo_k8s = new volcengine.apig.ApigUpstreamSource("foo-k8s", {
 *     gatewayId: fooApigGateway.id,
 *     comments: "acc-test-k8s",
 *     sourceType: "K8S",
 *     sourceSpec: {
 *         k8sSource: {
 *             clusterId: "cd197sac4mpmnruh7um80",
 *         },
 *     },
 *     ingressSettings: [{
 *         enableIngress: true,
 *         updateStatus: true,
 *         ingressClasses: ["test"],
 *         watchNamespaces: ["default"],
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * ApigUpstreamSource can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:apig/apigUpstreamSource:ApigUpstreamSource default resource_id
 * ```
 */
export class ApigUpstreamSource extends pulumi.CustomResource {
    /**
     * Get an existing ApigUpstreamSource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApigUpstreamSourceState, opts?: pulumi.CustomResourceOptions): ApigUpstreamSource {
        return new ApigUpstreamSource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:apig/apigUpstreamSource:ApigUpstreamSource';

    /**
     * Returns true if the given object is an instance of ApigUpstreamSource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApigUpstreamSource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApigUpstreamSource.__pulumiType;
    }

    /**
     * The comments of the apig upstream source.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * The gateway id of the apig upstream source.
     */
    public readonly gatewayId!: pulumi.Output<string>;
    /**
     * The ingress settings of apig upstream source.
     */
    public readonly ingressSettings!: pulumi.Output<outputs.apig.ApigUpstreamSourceIngressSetting[]>;
    /**
     * The source spec of apig upstream source.
     */
    public readonly sourceSpec!: pulumi.Output<outputs.apig.ApigUpstreamSourceSourceSpec>;
    /**
     * The source type of the apig upstream. Valid values: `K8S`, `Nacos`.
     */
    public readonly sourceType!: pulumi.Output<string>;

    /**
     * Create a ApigUpstreamSource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApigUpstreamSourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApigUpstreamSourceArgs | ApigUpstreamSourceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApigUpstreamSourceState | undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["gatewayId"] = state ? state.gatewayId : undefined;
            resourceInputs["ingressSettings"] = state ? state.ingressSettings : undefined;
            resourceInputs["sourceSpec"] = state ? state.sourceSpec : undefined;
            resourceInputs["sourceType"] = state ? state.sourceType : undefined;
        } else {
            const args = argsOrState as ApigUpstreamSourceArgs | undefined;
            if ((!args || args.gatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayId'");
            }
            if ((!args || args.sourceSpec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceSpec'");
            }
            if ((!args || args.sourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceType'");
            }
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["gatewayId"] = args ? args.gatewayId : undefined;
            resourceInputs["ingressSettings"] = args ? args.ingressSettings : undefined;
            resourceInputs["sourceSpec"] = args ? args.sourceSpec : undefined;
            resourceInputs["sourceType"] = args ? args.sourceType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApigUpstreamSource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApigUpstreamSource resources.
 */
export interface ApigUpstreamSourceState {
    /**
     * The comments of the apig upstream source.
     */
    comments?: pulumi.Input<string>;
    /**
     * The gateway id of the apig upstream source.
     */
    gatewayId?: pulumi.Input<string>;
    /**
     * The ingress settings of apig upstream source.
     */
    ingressSettings?: pulumi.Input<pulumi.Input<inputs.apig.ApigUpstreamSourceIngressSetting>[]>;
    /**
     * The source spec of apig upstream source.
     */
    sourceSpec?: pulumi.Input<inputs.apig.ApigUpstreamSourceSourceSpec>;
    /**
     * The source type of the apig upstream. Valid values: `K8S`, `Nacos`.
     */
    sourceType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApigUpstreamSource resource.
 */
export interface ApigUpstreamSourceArgs {
    /**
     * The comments of the apig upstream source.
     */
    comments?: pulumi.Input<string>;
    /**
     * The gateway id of the apig upstream source.
     */
    gatewayId: pulumi.Input<string>;
    /**
     * The ingress settings of apig upstream source.
     */
    ingressSettings?: pulumi.Input<pulumi.Input<inputs.apig.ApigUpstreamSourceIngressSetting>[]>;
    /**
     * The source spec of apig upstream source.
     */
    sourceSpec: pulumi.Input<inputs.apig.ApigUpstreamSourceSourceSpec>;
    /**
     * The source type of the apig upstream. Valid values: `K8S`, `Nacos`.
     */
    sourceType: pulumi.Input<string>;
}
