// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage tls shipper
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.tls.Shipper("foo", {
 *     contentInfo: {
 *         format: "json",
 *         jsonInfo: {
 *             enable: true,
 *             keys: [
 *                 "__content",
 *                 "__pod_name__",
 *                 "__path__",
 *                 "__tf-test__",
 *             ],
 *         },
 *     },
 *     shipperEndTime: 1751255700021,
 *     shipperName: "tf-test",
 *     shipperStartTime: 1750737324521,
 *     shipperType: "tos",
 *     topicId: "8ba48bd7-2493-4300-b1d0-cb7xxxxxx",
 *     tosShipperInfo: {
 *         bucket: "tf-test",
 *         compress: "snappy",
 *         interval: 100,
 *         maxSize: 100,
 *         partitionFormat: "%Y/%m/%d/%H/%M",
 *         prefix: "terraform_1.9.4_linux_amd64.zip",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Shipper can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:tls/shipper:Shipper default resource_id
 * ```
 */
export class Shipper extends pulumi.CustomResource {
    /**
     * Get an existing Shipper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ShipperState, opts?: pulumi.CustomResourceOptions): Shipper {
        return new Shipper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:tls/shipper:Shipper';

    /**
     * Returns true if the given object is an instance of Shipper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Shipper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Shipper.__pulumiType;
    }

    /**
     * Configuration of the delivery format for log content.
     */
    public readonly contentInfo!: pulumi.Output<outputs.tls.ShipperContentInfo>;
    /**
     * JSON format log content configuration.
     */
    public readonly kafkaShipperInfo!: pulumi.Output<outputs.tls.ShipperKafkaShipperInfo>;
    /**
     * Delivery end time, millisecond timestamp. If not configured, it will keep delivering. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly shipperEndTime!: pulumi.Output<number | undefined>;
    /**
     * Delivery configuration name.
     */
    public readonly shipperName!: pulumi.Output<string>;
    /**
     * Delivery start time, millisecond timestamp. If not configured, it defaults to the current time. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly shipperStartTime!: pulumi.Output<number | undefined>;
    /**
     * The type of delivery.
     */
    public readonly shipperType!: pulumi.Output<string>;
    /**
     * Whether to enable the delivery configuration. The default value is true.
     */
    public readonly status!: pulumi.Output<boolean>;
    /**
     * The log topic ID where the log to be delivered is located.
     */
    public readonly topicId!: pulumi.Output<string>;
    /**
     * Deliver the relevant configuration to the object storage (TOS).
     */
    public readonly tosShipperInfo!: pulumi.Output<outputs.tls.ShipperTosShipperInfo>;

    /**
     * Create a Shipper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ShipperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ShipperArgs | ShipperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ShipperState | undefined;
            resourceInputs["contentInfo"] = state ? state.contentInfo : undefined;
            resourceInputs["kafkaShipperInfo"] = state ? state.kafkaShipperInfo : undefined;
            resourceInputs["shipperEndTime"] = state ? state.shipperEndTime : undefined;
            resourceInputs["shipperName"] = state ? state.shipperName : undefined;
            resourceInputs["shipperStartTime"] = state ? state.shipperStartTime : undefined;
            resourceInputs["shipperType"] = state ? state.shipperType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["topicId"] = state ? state.topicId : undefined;
            resourceInputs["tosShipperInfo"] = state ? state.tosShipperInfo : undefined;
        } else {
            const args = argsOrState as ShipperArgs | undefined;
            if ((!args || args.contentInfo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contentInfo'");
            }
            if ((!args || args.shipperName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'shipperName'");
            }
            if ((!args || args.topicId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topicId'");
            }
            resourceInputs["contentInfo"] = args ? args.contentInfo : undefined;
            resourceInputs["kafkaShipperInfo"] = args ? args.kafkaShipperInfo : undefined;
            resourceInputs["shipperEndTime"] = args ? args.shipperEndTime : undefined;
            resourceInputs["shipperName"] = args ? args.shipperName : undefined;
            resourceInputs["shipperStartTime"] = args ? args.shipperStartTime : undefined;
            resourceInputs["shipperType"] = args ? args.shipperType : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["topicId"] = args ? args.topicId : undefined;
            resourceInputs["tosShipperInfo"] = args ? args.tosShipperInfo : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Shipper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Shipper resources.
 */
export interface ShipperState {
    /**
     * Configuration of the delivery format for log content.
     */
    contentInfo?: pulumi.Input<inputs.tls.ShipperContentInfo>;
    /**
     * JSON format log content configuration.
     */
    kafkaShipperInfo?: pulumi.Input<inputs.tls.ShipperKafkaShipperInfo>;
    /**
     * Delivery end time, millisecond timestamp. If not configured, it will keep delivering. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    shipperEndTime?: pulumi.Input<number>;
    /**
     * Delivery configuration name.
     */
    shipperName?: pulumi.Input<string>;
    /**
     * Delivery start time, millisecond timestamp. If not configured, it defaults to the current time. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    shipperStartTime?: pulumi.Input<number>;
    /**
     * The type of delivery.
     */
    shipperType?: pulumi.Input<string>;
    /**
     * Whether to enable the delivery configuration. The default value is true.
     */
    status?: pulumi.Input<boolean>;
    /**
     * The log topic ID where the log to be delivered is located.
     */
    topicId?: pulumi.Input<string>;
    /**
     * Deliver the relevant configuration to the object storage (TOS).
     */
    tosShipperInfo?: pulumi.Input<inputs.tls.ShipperTosShipperInfo>;
}

/**
 * The set of arguments for constructing a Shipper resource.
 */
export interface ShipperArgs {
    /**
     * Configuration of the delivery format for log content.
     */
    contentInfo: pulumi.Input<inputs.tls.ShipperContentInfo>;
    /**
     * JSON format log content configuration.
     */
    kafkaShipperInfo?: pulumi.Input<inputs.tls.ShipperKafkaShipperInfo>;
    /**
     * Delivery end time, millisecond timestamp. If not configured, it will keep delivering. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    shipperEndTime?: pulumi.Input<number>;
    /**
     * Delivery configuration name.
     */
    shipperName: pulumi.Input<string>;
    /**
     * Delivery start time, millisecond timestamp. If not configured, it defaults to the current time. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    shipperStartTime?: pulumi.Input<number>;
    /**
     * The type of delivery.
     */
    shipperType?: pulumi.Input<string>;
    /**
     * Whether to enable the delivery configuration. The default value is true.
     */
    status?: pulumi.Input<boolean>;
    /**
     * The log topic ID where the log to be delivered is located.
     */
    topicId: pulumi.Input<string>;
    /**
     * Deliver the relevant configuration to the object storage (TOS).
     */
    tosShipperInfo?: pulumi.Input<inputs.tls.ShipperTosShipperInfo>;
}
