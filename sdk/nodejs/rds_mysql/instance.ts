// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage rds mysql instance
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = new volcengine.rds_mysql.Instance("foo", {
 *     dbEngineVersion: "MySQL_5_7",
 *     nodeSpec: "rds.mysql.1c2g",
 *     primaryZoneId: "cn-guilin-a",
 *     secondaryZoneId: "cn-guilin-b",
 *     storageSpace: 80,
 *     subnetId: "subnet-2d72yi377stts58ozfdrlk9f6",
 *     instanceName: "tf-test",
 *     lowerCaseTableNames: "1",
 *     chargeInfo: {
 *         chargeType: "PostPaid",
 *     },
 *     allowListIds: [
 *         "acl-2dd8f8317e4d4159b21630d13ae2e6ec",
 *         "acl-2eaa2a053b2a4a58b988e38ae975e81c",
 *     ],
 *     parameters: [
 *         {
 *             parameterName: "auto_increment_increment",
 *             parameterValue: "2",
 *         },
 *         {
 *             parameterName: "auto_increment_offset",
 *             parameterValue: "4",
 *         },
 *     ],
 * });
 * const readonly = new volcengine.rds_mysql.InstanceReadonlyNode("readonly", {
 *     instanceId: foo.id,
 *     nodeSpec: "rds.mysql.2c4g",
 *     zoneId: "cn-guilin-a",
 * });
 * ```
 *
 * ## Import
 *
 * Rds Mysql Instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:rds_mysql/instance:Instance default mysql-72da4258c2c7
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:rds_mysql/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Allow list Ids of the RDS instance.
     */
    public readonly allowListIds!: pulumi.Output<string[] | undefined>;
    /**
     * The version of allow list.
     */
    public /*out*/ readonly allowListVersion!: pulumi.Output<string>;
    /**
     * The instance has used backup space. Unit: GB.
     */
    public /*out*/ readonly backupUse!: pulumi.Output<number>;
    /**
     * Payment methods.
     */
    public /*out*/ readonly chargeDetails!: pulumi.Output<outputs.rds_mysql.InstanceChargeDetail[]>;
    /**
     * Payment methods.
     */
    public readonly chargeInfo!: pulumi.Output<outputs.rds_mysql.InstanceChargeInfo>;
    /**
     * Node creation local time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Data synchronization mode.
     */
    public /*out*/ readonly dataSyncMode!: pulumi.Output<string>;
    /**
     * Instance type. Value:
     * MySQL_5_7
     * MySQL_8_0.
     */
    public readonly dbEngineVersion!: pulumi.Output<string>;
    /**
     * Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly dbTimeZone!: pulumi.Output<string>;
    /**
     * The endpoint info of the RDS instance.
     */
    public /*out*/ readonly endpoints!: pulumi.Output<outputs.rds_mysql.InstanceEndpoint[]>;
    /**
     * Instance ID.
     */
    public /*out*/ readonly instanceId!: pulumi.Output<string>;
    /**
     * Instance name. Cannot start with a number or a dash
     * Can only contain Chinese characters, letters, numbers, underscores and dashes
     * The length is limited between 1 ~ 128.
     */
    public readonly instanceName!: pulumi.Output<string | undefined>;
    /**
     * The status of the RDS instance.
     */
    public /*out*/ readonly instanceStatus!: pulumi.Output<string>;
    /**
     * Whether the table name is case sensitive, the default value is 1.
     * Ranges:
     * 0: Table names are stored as fixed and table names are case-sensitive.
     * 1: Table names will be stored in lowercase and table names are not case sensitive.
     */
    public readonly lowerCaseTableNames!: pulumi.Output<string | undefined>;
    /**
     * Maintenance Window.
     */
    public /*out*/ readonly maintenanceWindows!: pulumi.Output<outputs.rds_mysql.InstanceMaintenanceWindow[]>;
    /**
     * Memory size in GB.
     */
    public /*out*/ readonly memory!: pulumi.Output<number>;
    /**
     * The number of nodes.
     */
    public /*out*/ readonly nodeNumber!: pulumi.Output<number>;
    /**
     * The specification of primary node and secondary node.
     */
    public readonly nodeSpec!: pulumi.Output<string>;
    /**
     * Instance node information.
     */
    public /*out*/ readonly nodes!: pulumi.Output<outputs.rds_mysql.InstanceNode[]>;
    /**
     * Parameter of the RDS instance. This field can only be added or modified. Deleting this field is invalid.
     */
    public readonly parameters!: pulumi.Output<outputs.rds_mysql.InstanceParameter[] | undefined>;
    /**
     * The available zone of primary node.
     */
    public readonly primaryZoneId!: pulumi.Output<string>;
    /**
     * The region of the RDS instance.
     */
    public /*out*/ readonly regionId!: pulumi.Output<string>;
    /**
     * The available zone of secondary node.
     */
    public readonly secondaryZoneId!: pulumi.Output<string>;
    /**
     * Instance storage space. Value range: [20, 3000], unit: GB, increments every 100GB. Default value: 100.
     */
    public readonly storageSpace!: pulumi.Output<number | undefined>;
    /**
     * Instance storage type.
     */
    public /*out*/ readonly storageType!: pulumi.Output<string>;
    /**
     * The instance has used storage space. Unit: GB.
     */
    public /*out*/ readonly storageUse!: pulumi.Output<number>;
    /**
     * Subnet ID of the RDS instance.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * Time zone.
     */
    public /*out*/ readonly timeZone!: pulumi.Output<string>;
    /**
     * The update time of the RDS instance.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * CPU size.
     */
    public /*out*/ readonly vCpu!: pulumi.Output<number>;
    /**
     * The vpc ID of the RDS instance.
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;
    /**
     * The available zone of the RDS instance.
     */
    public /*out*/ readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["allowListIds"] = state ? state.allowListIds : undefined;
            resourceInputs["allowListVersion"] = state ? state.allowListVersion : undefined;
            resourceInputs["backupUse"] = state ? state.backupUse : undefined;
            resourceInputs["chargeDetails"] = state ? state.chargeDetails : undefined;
            resourceInputs["chargeInfo"] = state ? state.chargeInfo : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["dataSyncMode"] = state ? state.dataSyncMode : undefined;
            resourceInputs["dbEngineVersion"] = state ? state.dbEngineVersion : undefined;
            resourceInputs["dbTimeZone"] = state ? state.dbTimeZone : undefined;
            resourceInputs["endpoints"] = state ? state.endpoints : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["instanceStatus"] = state ? state.instanceStatus : undefined;
            resourceInputs["lowerCaseTableNames"] = state ? state.lowerCaseTableNames : undefined;
            resourceInputs["maintenanceWindows"] = state ? state.maintenanceWindows : undefined;
            resourceInputs["memory"] = state ? state.memory : undefined;
            resourceInputs["nodeNumber"] = state ? state.nodeNumber : undefined;
            resourceInputs["nodeSpec"] = state ? state.nodeSpec : undefined;
            resourceInputs["nodes"] = state ? state.nodes : undefined;
            resourceInputs["parameters"] = state ? state.parameters : undefined;
            resourceInputs["primaryZoneId"] = state ? state.primaryZoneId : undefined;
            resourceInputs["regionId"] = state ? state.regionId : undefined;
            resourceInputs["secondaryZoneId"] = state ? state.secondaryZoneId : undefined;
            resourceInputs["storageSpace"] = state ? state.storageSpace : undefined;
            resourceInputs["storageType"] = state ? state.storageType : undefined;
            resourceInputs["storageUse"] = state ? state.storageUse : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["timeZone"] = state ? state.timeZone : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
            resourceInputs["vCpu"] = state ? state.vCpu : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.chargeInfo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'chargeInfo'");
            }
            if ((!args || args.dbEngineVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbEngineVersion'");
            }
            if ((!args || args.nodeSpec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeSpec'");
            }
            if ((!args || args.primaryZoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'primaryZoneId'");
            }
            if ((!args || args.secondaryZoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secondaryZoneId'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            resourceInputs["allowListIds"] = args ? args.allowListIds : undefined;
            resourceInputs["chargeInfo"] = args ? args.chargeInfo : undefined;
            resourceInputs["dbEngineVersion"] = args ? args.dbEngineVersion : undefined;
            resourceInputs["dbTimeZone"] = args ? args.dbTimeZone : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["lowerCaseTableNames"] = args ? args.lowerCaseTableNames : undefined;
            resourceInputs["nodeSpec"] = args ? args.nodeSpec : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["primaryZoneId"] = args ? args.primaryZoneId : undefined;
            resourceInputs["secondaryZoneId"] = args ? args.secondaryZoneId : undefined;
            resourceInputs["storageSpace"] = args ? args.storageSpace : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["allowListVersion"] = undefined /*out*/;
            resourceInputs["backupUse"] = undefined /*out*/;
            resourceInputs["chargeDetails"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["dataSyncMode"] = undefined /*out*/;
            resourceInputs["endpoints"] = undefined /*out*/;
            resourceInputs["instanceId"] = undefined /*out*/;
            resourceInputs["instanceStatus"] = undefined /*out*/;
            resourceInputs["maintenanceWindows"] = undefined /*out*/;
            resourceInputs["memory"] = undefined /*out*/;
            resourceInputs["nodeNumber"] = undefined /*out*/;
            resourceInputs["nodes"] = undefined /*out*/;
            resourceInputs["regionId"] = undefined /*out*/;
            resourceInputs["storageType"] = undefined /*out*/;
            resourceInputs["storageUse"] = undefined /*out*/;
            resourceInputs["timeZone"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["vCpu"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
            resourceInputs["zoneId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Allow list Ids of the RDS instance.
     */
    allowListIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The version of allow list.
     */
    allowListVersion?: pulumi.Input<string>;
    /**
     * The instance has used backup space. Unit: GB.
     */
    backupUse?: pulumi.Input<number>;
    /**
     * Payment methods.
     */
    chargeDetails?: pulumi.Input<pulumi.Input<inputs.rds_mysql.InstanceChargeDetail>[]>;
    /**
     * Payment methods.
     */
    chargeInfo?: pulumi.Input<inputs.rds_mysql.InstanceChargeInfo>;
    /**
     * Node creation local time.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Data synchronization mode.
     */
    dataSyncMode?: pulumi.Input<string>;
    /**
     * Instance type. Value:
     * MySQL_5_7
     * MySQL_8_0.
     */
    dbEngineVersion?: pulumi.Input<string>;
    /**
     * Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    dbTimeZone?: pulumi.Input<string>;
    /**
     * The endpoint info of the RDS instance.
     */
    endpoints?: pulumi.Input<pulumi.Input<inputs.rds_mysql.InstanceEndpoint>[]>;
    /**
     * Instance ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Instance name. Cannot start with a number or a dash
     * Can only contain Chinese characters, letters, numbers, underscores and dashes
     * The length is limited between 1 ~ 128.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The status of the RDS instance.
     */
    instanceStatus?: pulumi.Input<string>;
    /**
     * Whether the table name is case sensitive, the default value is 1.
     * Ranges:
     * 0: Table names are stored as fixed and table names are case-sensitive.
     * 1: Table names will be stored in lowercase and table names are not case sensitive.
     */
    lowerCaseTableNames?: pulumi.Input<string>;
    /**
     * Maintenance Window.
     */
    maintenanceWindows?: pulumi.Input<pulumi.Input<inputs.rds_mysql.InstanceMaintenanceWindow>[]>;
    /**
     * Memory size in GB.
     */
    memory?: pulumi.Input<number>;
    /**
     * The number of nodes.
     */
    nodeNumber?: pulumi.Input<number>;
    /**
     * The specification of primary node and secondary node.
     */
    nodeSpec?: pulumi.Input<string>;
    /**
     * Instance node information.
     */
    nodes?: pulumi.Input<pulumi.Input<inputs.rds_mysql.InstanceNode>[]>;
    /**
     * Parameter of the RDS instance. This field can only be added or modified. Deleting this field is invalid.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.rds_mysql.InstanceParameter>[]>;
    /**
     * The available zone of primary node.
     */
    primaryZoneId?: pulumi.Input<string>;
    /**
     * The region of the RDS instance.
     */
    regionId?: pulumi.Input<string>;
    /**
     * The available zone of secondary node.
     */
    secondaryZoneId?: pulumi.Input<string>;
    /**
     * Instance storage space. Value range: [20, 3000], unit: GB, increments every 100GB. Default value: 100.
     */
    storageSpace?: pulumi.Input<number>;
    /**
     * Instance storage type.
     */
    storageType?: pulumi.Input<string>;
    /**
     * The instance has used storage space. Unit: GB.
     */
    storageUse?: pulumi.Input<number>;
    /**
     * Subnet ID of the RDS instance.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Time zone.
     */
    timeZone?: pulumi.Input<string>;
    /**
     * The update time of the RDS instance.
     */
    updateTime?: pulumi.Input<string>;
    /**
     * CPU size.
     */
    vCpu?: pulumi.Input<number>;
    /**
     * The vpc ID of the RDS instance.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The available zone of the RDS instance.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Allow list Ids of the RDS instance.
     */
    allowListIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Payment methods.
     */
    chargeInfo: pulumi.Input<inputs.rds_mysql.InstanceChargeInfo>;
    /**
     * Instance type. Value:
     * MySQL_5_7
     * MySQL_8_0.
     */
    dbEngineVersion: pulumi.Input<string>;
    /**
     * Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    dbTimeZone?: pulumi.Input<string>;
    /**
     * Instance name. Cannot start with a number or a dash
     * Can only contain Chinese characters, letters, numbers, underscores and dashes
     * The length is limited between 1 ~ 128.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Whether the table name is case sensitive, the default value is 1.
     * Ranges:
     * 0: Table names are stored as fixed and table names are case-sensitive.
     * 1: Table names will be stored in lowercase and table names are not case sensitive.
     */
    lowerCaseTableNames?: pulumi.Input<string>;
    /**
     * The specification of primary node and secondary node.
     */
    nodeSpec: pulumi.Input<string>;
    /**
     * Parameter of the RDS instance. This field can only be added or modified. Deleting this field is invalid.
     */
    parameters?: pulumi.Input<pulumi.Input<inputs.rds_mysql.InstanceParameter>[]>;
    /**
     * The available zone of primary node.
     */
    primaryZoneId: pulumi.Input<string>;
    /**
     * The available zone of secondary node.
     */
    secondaryZoneId: pulumi.Input<string>;
    /**
     * Instance storage space. Value range: [20, 3000], unit: GB, increments every 100GB. Default value: 100.
     */
    storageSpace?: pulumi.Input<number>;
    /**
     * Subnet ID of the RDS instance.
     */
    subnetId: pulumi.Input<string>;
}