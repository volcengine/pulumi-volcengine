// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage rds postgresql instance readonly node
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooZones = volcengine.ecs.Zones({});
 * const fooVpc = new volcengine.vpc.Vpc("fooVpc", {
 *     vpcName: "acc-test-project1",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const fooSubnet = new volcengine.vpc.Subnet("fooSubnet", {
 *     subnetName: "acc-subnet-test-2",
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *     vpcId: fooVpc.id,
 * });
 * const fooInstance = new volcengine.rds_postgresql.Instance("fooInstance", {
 *     dbEngineVersion: "PostgreSQL_12",
 *     nodeSpec: "rds.postgres.1c2g",
 *     primaryZoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *     secondaryZoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *     storageSpace: 40,
 *     subnetId: fooSubnet.id,
 *     instanceName: "acc-test-1",
 *     chargeInfo: {
 *         chargeType: "PostPaid",
 *     },
 *     projectName: "default",
 *     tags: [{
 *         key: "tfk1",
 *         value: "tfv1",
 *     }],
 *     parameters: [
 *         {
 *             name: "auto_explain.log_analyze",
 *             value: "off",
 *         },
 *         {
 *             name: "auto_explain.log_format",
 *             value: "text",
 *         },
 *     ],
 * });
 * const fooInstanceReadonlyNode = new volcengine.rds_postgresql.InstanceReadonlyNode("fooInstanceReadonlyNode", {
 *     instanceId: fooInstance.id,
 *     nodeSpec: "rds.postgres.1c2g",
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 * });
 * ```
 *
 * ## Import
 *
 * RdsPostgresqlInstanceReadonlyNode can be imported using the instance_id:node_id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:rds_postgresql/instanceReadonlyNode:InstanceReadonlyNode default postgres-21a3333b****:postgres-ca7b7019****
 * ```
 */
export class InstanceReadonlyNode extends pulumi.CustomResource {
    /**
     * Get an existing InstanceReadonlyNode resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceReadonlyNodeState, opts?: pulumi.CustomResourceOptions): InstanceReadonlyNode {
        return new InstanceReadonlyNode(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:rds_postgresql/instanceReadonlyNode:InstanceReadonlyNode';

    /**
     * Returns true if the given object is an instance of InstanceReadonlyNode.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceReadonlyNode {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceReadonlyNode.__pulumiType;
    }

    /**
     * The RDS PostgreSQL instance id of the readonly node.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The id of the readonly node.
     */
    public /*out*/ readonly nodeId!: pulumi.Output<string>;
    /**
     * The specification of readonly node.
     */
    public readonly nodeSpec!: pulumi.Output<string>;
    /**
     * The available zone of readonly node.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a InstanceReadonlyNode resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceReadonlyNodeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceReadonlyNodeArgs | InstanceReadonlyNodeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceReadonlyNodeState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["nodeId"] = state ? state.nodeId : undefined;
            resourceInputs["nodeSpec"] = state ? state.nodeSpec : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as InstanceReadonlyNodeArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.nodeSpec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeSpec'");
            }
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["nodeSpec"] = args ? args.nodeSpec : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["nodeId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceReadonlyNode.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceReadonlyNode resources.
 */
export interface InstanceReadonlyNodeState {
    /**
     * The RDS PostgreSQL instance id of the readonly node.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The id of the readonly node.
     */
    nodeId?: pulumi.Input<string>;
    /**
     * The specification of readonly node.
     */
    nodeSpec?: pulumi.Input<string>;
    /**
     * The available zone of readonly node.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceReadonlyNode resource.
 */
export interface InstanceReadonlyNodeArgs {
    /**
     * The RDS PostgreSQL instance id of the readonly node.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The specification of readonly node.
     */
    nodeSpec: pulumi.Input<string>;
    /**
     * The available zone of readonly node.
     */
    zoneId: pulumi.Input<string>;
}