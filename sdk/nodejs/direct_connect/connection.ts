// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage direct connect connection
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.direct_connect.Connection("foo", {
 *     bandwidth: 1000,
 *     customerContactEmail: "email@aaa.com",
 *     customerContactPhone: "12345678911",
 *     customerName: "tf-a",
 *     description: "tf-test",
 *     directConnectAccessPointId: "ap-cn-beijing-a",
 *     directConnectConnectionName: "tf-test-connection",
 *     lineOperator: "ChinaOther",
 *     peerLocation: "XX路XX号XX楼XX机房",
 *     portSpec: "10G",
 *     portType: "10GBase",
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * DirectConnectConnection can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:direct_connect/connection:Connection default dcc-7qthudw0ll6jmc****
 * ```
 */
export class Connection extends pulumi.CustomResource {
    /**
     * Get an existing Connection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectionState, opts?: pulumi.CustomResourceOptions): Connection {
        return new Connection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:direct_connect/connection:Connection';

    /**
     * Returns true if the given object is an instance of Connection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Connection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Connection.__pulumiType;
    }

    /**
     * The line band width,unit:Mbps.
     */
    public readonly bandwidth!: pulumi.Output<number>;
    /**
     * The dedicated line contact email.
     */
    public readonly customerContactEmail!: pulumi.Output<string>;
    /**
     * The dedicated line contact phone.
     */
    public readonly customerContactPhone!: pulumi.Output<string>;
    /**
     * The dedicated line contact name.
     */
    public readonly customerName!: pulumi.Output<string>;
    /**
     * The description of direct connect.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The direct connect access point id.
     */
    public readonly directConnectAccessPointId!: pulumi.Output<string>;
    /**
     * The name of direct connect.
     */
    public readonly directConnectConnectionName!: pulumi.Output<string>;
    /**
     * The physical leased line operator.valid value contains `ChinaTelecom`,`ChinaMobile`,`ChinaUnicom`,`ChinaOther`.
     */
    public readonly lineOperator!: pulumi.Output<string>;
    /**
     * The local IDC address.
     */
    public readonly peerLocation!: pulumi.Output<string>;
    /**
     * The physical leased line port spec.valid value contains `1G`,`10G`.
     */
    public readonly portSpec!: pulumi.Output<string>;
    /**
     * The physical leased line port type and spec.valid value contains `1000Base-T`,`10GBase-T`,`1000Base`,`10GBase`,`40GBase`,`100GBase`.
     */
    public readonly portType!: pulumi.Output<string>;
    /**
     * The physical leased line tags.
     */
    public readonly tags!: pulumi.Output<outputs.direct_connect.ConnectionTag[] | undefined>;

    /**
     * Create a Connection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectionArgs | ConnectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConnectionState | undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["customerContactEmail"] = state ? state.customerContactEmail : undefined;
            resourceInputs["customerContactPhone"] = state ? state.customerContactPhone : undefined;
            resourceInputs["customerName"] = state ? state.customerName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["directConnectAccessPointId"] = state ? state.directConnectAccessPointId : undefined;
            resourceInputs["directConnectConnectionName"] = state ? state.directConnectConnectionName : undefined;
            resourceInputs["lineOperator"] = state ? state.lineOperator : undefined;
            resourceInputs["peerLocation"] = state ? state.peerLocation : undefined;
            resourceInputs["portSpec"] = state ? state.portSpec : undefined;
            resourceInputs["portType"] = state ? state.portType : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ConnectionArgs | undefined;
            if ((!args || args.bandwidth === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidth'");
            }
            if ((!args || args.customerContactEmail === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerContactEmail'");
            }
            if ((!args || args.customerContactPhone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerContactPhone'");
            }
            if ((!args || args.customerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerName'");
            }
            if ((!args || args.directConnectAccessPointId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'directConnectAccessPointId'");
            }
            if ((!args || args.lineOperator === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lineOperator'");
            }
            if ((!args || args.peerLocation === undefined) && !opts.urn) {
                throw new Error("Missing required property 'peerLocation'");
            }
            if ((!args || args.portSpec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'portSpec'");
            }
            if ((!args || args.portType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'portType'");
            }
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["customerContactEmail"] = args ? args.customerContactEmail : undefined;
            resourceInputs["customerContactPhone"] = args ? args.customerContactPhone : undefined;
            resourceInputs["customerName"] = args ? args.customerName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["directConnectAccessPointId"] = args ? args.directConnectAccessPointId : undefined;
            resourceInputs["directConnectConnectionName"] = args ? args.directConnectConnectionName : undefined;
            resourceInputs["lineOperator"] = args ? args.lineOperator : undefined;
            resourceInputs["peerLocation"] = args ? args.peerLocation : undefined;
            resourceInputs["portSpec"] = args ? args.portSpec : undefined;
            resourceInputs["portType"] = args ? args.portType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Connection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Connection resources.
 */
export interface ConnectionState {
    /**
     * The line band width,unit:Mbps.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * The dedicated line contact email.
     */
    customerContactEmail?: pulumi.Input<string>;
    /**
     * The dedicated line contact phone.
     */
    customerContactPhone?: pulumi.Input<string>;
    /**
     * The dedicated line contact name.
     */
    customerName?: pulumi.Input<string>;
    /**
     * The description of direct connect.
     */
    description?: pulumi.Input<string>;
    /**
     * The direct connect access point id.
     */
    directConnectAccessPointId?: pulumi.Input<string>;
    /**
     * The name of direct connect.
     */
    directConnectConnectionName?: pulumi.Input<string>;
    /**
     * The physical leased line operator.valid value contains `ChinaTelecom`,`ChinaMobile`,`ChinaUnicom`,`ChinaOther`.
     */
    lineOperator?: pulumi.Input<string>;
    /**
     * The local IDC address.
     */
    peerLocation?: pulumi.Input<string>;
    /**
     * The physical leased line port spec.valid value contains `1G`,`10G`.
     */
    portSpec?: pulumi.Input<string>;
    /**
     * The physical leased line port type and spec.valid value contains `1000Base-T`,`10GBase-T`,`1000Base`,`10GBase`,`40GBase`,`100GBase`.
     */
    portType?: pulumi.Input<string>;
    /**
     * The physical leased line tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.direct_connect.ConnectionTag>[]>;
}

/**
 * The set of arguments for constructing a Connection resource.
 */
export interface ConnectionArgs {
    /**
     * The line band width,unit:Mbps.
     */
    bandwidth: pulumi.Input<number>;
    /**
     * The dedicated line contact email.
     */
    customerContactEmail: pulumi.Input<string>;
    /**
     * The dedicated line contact phone.
     */
    customerContactPhone: pulumi.Input<string>;
    /**
     * The dedicated line contact name.
     */
    customerName: pulumi.Input<string>;
    /**
     * The description of direct connect.
     */
    description?: pulumi.Input<string>;
    /**
     * The direct connect access point id.
     */
    directConnectAccessPointId: pulumi.Input<string>;
    /**
     * The name of direct connect.
     */
    directConnectConnectionName?: pulumi.Input<string>;
    /**
     * The physical leased line operator.valid value contains `ChinaTelecom`,`ChinaMobile`,`ChinaUnicom`,`ChinaOther`.
     */
    lineOperator: pulumi.Input<string>;
    /**
     * The local IDC address.
     */
    peerLocation: pulumi.Input<string>;
    /**
     * The physical leased line port spec.valid value contains `1G`,`10G`.
     */
    portSpec: pulumi.Input<string>;
    /**
     * The physical leased line port type and spec.valid value contains `1000Base-T`,`10GBase-T`,`1000Base`,`10GBase`,`40GBase`,`100GBase`.
     */
    portType: pulumi.Input<string>;
    /**
     * The physical leased line tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.direct_connect.ConnectionTag>[]>;
}
