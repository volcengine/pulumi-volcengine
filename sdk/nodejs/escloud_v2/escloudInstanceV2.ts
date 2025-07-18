// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooZones = volcengine.ecs.getZones({});
 * // create vpc
 * const fooVpc = new volcengine.vpc.Vpc("fooVpc", {
 *     vpcName: "acc-test-vpc",
 *     cidrBlock: "172.16.0.0/16",
 *     dnsServers: [
 *         "8.8.8.8",
 *         "114.114.114.114",
 *     ],
 *     projectName: "default",
 * });
 * // create subnet
 * const fooSubnet = new volcengine.vpc.Subnet("fooSubnet", {
 *     subnetName: "acc-test-subnet",
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *     vpcId: fooVpc.id,
 * });
 * // create escloud instance
 * const fooEscloudInstanceV2 = new volcengine.escloud_v2.EscloudInstanceV2("fooEscloudInstanceV2", {
 *     instanceName: "acc-test-escloud-instance",
 *     version: "V7_10",
 *     zoneIds: [
 *         fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *         fooZones.then(fooZones => fooZones.zones?.[1]?.id),
 *         fooZones.then(fooZones => fooZones.zones?.[2]?.id),
 *     ],
 *     subnetId: fooSubnet.id,
 *     enableHttps: false,
 *     adminPassword: "Password@@123",
 *     chargeType: "PostPaid",
 *     autoRenew: false,
 *     period: 1,
 *     configurationCode: "es.standard",
 *     enablePureMaster: true,
 *     deletionProtection: false,
 *     projectName: "default",
 *     nodeSpecsAssigns: [
 *         {
 *             type: "Master",
 *             number: 3,
 *             resourceSpecName: "es.x2.medium",
 *             storageSpecName: "es.volume.essd.pl0",
 *             storageSize: 20,
 *         },
 *         {
 *             type: "Hot",
 *             number: 6,
 *             resourceSpecName: "es.x2.medium",
 *             storageSpecName: "es.volume.essd.flexpl-standard",
 *             storageSize: 500,
 *             extraPerformance: {
 *                 throughput: 65,
 *             },
 *         },
 *         {
 *             type: "Kibana",
 *             number: 1,
 *             resourceSpecName: "kibana.x2.small",
 *             storageSpecName: "",
 *             storageSize: 0,
 *         },
 *     ],
 *     networkSpecs: [
 *         {
 *             type: "Elasticsearch",
 *             bandwidth: 1,
 *             isOpen: true,
 *             specName: "es.eip.bgp_fixed_bandwidth",
 *         },
 *         {
 *             type: "Kibana",
 *             bandwidth: 1,
 *             isOpen: true,
 *             specName: "es.eip.bgp_fixed_bandwidth",
 *         },
 *     ],
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * // create escloud ip white list
 * const fooEscloudIpWhiteList = new volcengine.escloud_v2.EscloudIpWhiteList("fooEscloudIpWhiteList", {
 *     instanceId: fooEscloudInstanceV2.id,
 *     type: "public",
 *     component: "es",
 *     ipLists: [
 *         "172.16.0.10",
 *         "172.16.0.11",
 *         "172.16.0.12",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * EscloudInstanceV2 can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:escloud_v2/escloudInstanceV2:EscloudInstanceV2 default resource_id
 * ```
 */
export class EscloudInstanceV2 extends pulumi.CustomResource {
    /**
     * Get an existing EscloudInstanceV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EscloudInstanceV2State, opts?: pulumi.CustomResourceOptions): EscloudInstanceV2 {
        return new EscloudInstanceV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:escloud_v2/escloudInstanceV2:EscloudInstanceV2';

    /**
     * Returns true if the given object is an instance of EscloudInstanceV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EscloudInstanceV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EscloudInstanceV2.__pulumiType;
    }

    /**
     * The password of administrator account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly adminPassword!: pulumi.Output<string>;
    /**
     * Whether to automatically renew in prepaid scenarios. Default is false.
     */
    public readonly autoRenew!: pulumi.Output<boolean | undefined>;
    /**
     * The cerebro private domain of instance.
     */
    public /*out*/ readonly cerebroPrivateDomain!: pulumi.Output<string>;
    /**
     * The cerebro public domain of instance.
     */
    public /*out*/ readonly cerebroPublicDomain!: pulumi.Output<string>;
    /**
     * The charge type of ESCloud instance, valid values: `PostPaid`, `PrePaid`.
     */
    public readonly chargeType!: pulumi.Output<string>;
    /**
     * Configuration code used for billing.
     */
    public readonly configurationCode!: pulumi.Output<string>;
    /**
     * Whether enable deletion protection for ESCloud instance. Default is false.
     */
    public readonly deletionProtection!: pulumi.Output<boolean | undefined>;
    /**
     * Whether Https access is enabled.
     */
    public readonly enableHttps!: pulumi.Output<boolean>;
    /**
     * Whether the Master node is independent.
     */
    public readonly enablePureMaster!: pulumi.Output<boolean>;
    /**
     * The eip address of instance.
     */
    public /*out*/ readonly esEip!: pulumi.Output<string>;
    /**
     * The eip id associated with the instance.
     */
    public /*out*/ readonly esEipId!: pulumi.Output<string>;
    /**
     * The es private domain of instance.
     */
    public /*out*/ readonly esPrivateDomain!: pulumi.Output<string>;
    /**
     * The es private endpoint of instance.
     */
    public /*out*/ readonly esPrivateEndpoint!: pulumi.Output<string>;
    /**
     * The whitelist of es private ip.
     */
    public /*out*/ readonly esPrivateIpWhitelist!: pulumi.Output<string>;
    /**
     * The es public domain of instance.
     */
    public /*out*/ readonly esPublicDomain!: pulumi.Output<string>;
    /**
     * The es public endpoint of instance.
     */
    public /*out*/ readonly esPublicEndpoint!: pulumi.Output<string>;
    /**
     * The whitelist of es public ip.
     */
    public /*out*/ readonly esPublicIpWhitelist!: pulumi.Output<string>;
    /**
     * The name of ESCloud instance.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * The eip address of kibana.
     */
    public /*out*/ readonly kibanaEip!: pulumi.Output<string>;
    /**
     * The eip id associated with kibana.
     */
    public /*out*/ readonly kibanaEipId!: pulumi.Output<string>;
    /**
     * The kibana private domain of instance.
     */
    public /*out*/ readonly kibanaPrivateDomain!: pulumi.Output<string>;
    /**
     * The whitelist of kibana private ip.
     */
    public /*out*/ readonly kibanaPrivateIpWhitelist!: pulumi.Output<string>;
    /**
     * The kibana public domain of instance.
     */
    public /*out*/ readonly kibanaPublicDomain!: pulumi.Output<string>;
    /**
     * The whitelist of kibana public ip.
     */
    public /*out*/ readonly kibanaPublicIpWhitelist!: pulumi.Output<string>;
    /**
     * The main zone id of instance.
     */
    public /*out*/ readonly mainZoneId!: pulumi.Output<string>;
    /**
     * The maintainable day for the instance. Valid values: `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`. Works only on modified scenes.
     */
    public readonly maintenanceDays!: pulumi.Output<string[]>;
    /**
     * The maintainable time period for the instance. Works only on modified scenes.
     */
    public readonly maintenanceTime!: pulumi.Output<string>;
    /**
     * The public network config of the ESCloud instance.
     */
    public readonly networkSpecs!: pulumi.Output<outputs.escloud_v2.EscloudInstanceV2NetworkSpec[] | undefined>;
    /**
     * The number and configuration of various ESCloud instance node. Kibana NodeSpecsAssign should not be modified.
     */
    public readonly nodeSpecsAssigns!: pulumi.Output<outputs.escloud_v2.EscloudInstanceV2NodeSpecsAssign[]>;
    /**
     * Purchase duration in prepaid scenarios. Unit: Monthly.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The project name to which the ESCloud instance belongs.
     */
    public readonly projectName!: pulumi.Output<string>;
    /**
     * The status of instance.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The id of subnet, the subnet must belong to the AZ selected.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * Tags.
     */
    public readonly tags!: pulumi.Output<outputs.escloud_v2.EscloudInstanceV2Tag[] | undefined>;
    /**
     * The version of instance. When creating ESCloud instance, the valid value is `V6_7` or `V7_10`. When creating OpenSearch instance, the valid value is `OPEN_SEARCH_2_9`.
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * The zone id of the ESCloud instance. Support specifying multiple availability zones.
     * The first zone id is the primary availability zone, while the rest are backup availability zones.
     */
    public readonly zoneIds!: pulumi.Output<string[]>;

    /**
     * Create a EscloudInstanceV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EscloudInstanceV2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EscloudInstanceV2Args | EscloudInstanceV2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EscloudInstanceV2State | undefined;
            resourceInputs["adminPassword"] = state ? state.adminPassword : undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["cerebroPrivateDomain"] = state ? state.cerebroPrivateDomain : undefined;
            resourceInputs["cerebroPublicDomain"] = state ? state.cerebroPublicDomain : undefined;
            resourceInputs["chargeType"] = state ? state.chargeType : undefined;
            resourceInputs["configurationCode"] = state ? state.configurationCode : undefined;
            resourceInputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            resourceInputs["enableHttps"] = state ? state.enableHttps : undefined;
            resourceInputs["enablePureMaster"] = state ? state.enablePureMaster : undefined;
            resourceInputs["esEip"] = state ? state.esEip : undefined;
            resourceInputs["esEipId"] = state ? state.esEipId : undefined;
            resourceInputs["esPrivateDomain"] = state ? state.esPrivateDomain : undefined;
            resourceInputs["esPrivateEndpoint"] = state ? state.esPrivateEndpoint : undefined;
            resourceInputs["esPrivateIpWhitelist"] = state ? state.esPrivateIpWhitelist : undefined;
            resourceInputs["esPublicDomain"] = state ? state.esPublicDomain : undefined;
            resourceInputs["esPublicEndpoint"] = state ? state.esPublicEndpoint : undefined;
            resourceInputs["esPublicIpWhitelist"] = state ? state.esPublicIpWhitelist : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["kibanaEip"] = state ? state.kibanaEip : undefined;
            resourceInputs["kibanaEipId"] = state ? state.kibanaEipId : undefined;
            resourceInputs["kibanaPrivateDomain"] = state ? state.kibanaPrivateDomain : undefined;
            resourceInputs["kibanaPrivateIpWhitelist"] = state ? state.kibanaPrivateIpWhitelist : undefined;
            resourceInputs["kibanaPublicDomain"] = state ? state.kibanaPublicDomain : undefined;
            resourceInputs["kibanaPublicIpWhitelist"] = state ? state.kibanaPublicIpWhitelist : undefined;
            resourceInputs["mainZoneId"] = state ? state.mainZoneId : undefined;
            resourceInputs["maintenanceDays"] = state ? state.maintenanceDays : undefined;
            resourceInputs["maintenanceTime"] = state ? state.maintenanceTime : undefined;
            resourceInputs["networkSpecs"] = state ? state.networkSpecs : undefined;
            resourceInputs["nodeSpecsAssigns"] = state ? state.nodeSpecsAssigns : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["zoneIds"] = state ? state.zoneIds : undefined;
        } else {
            const args = argsOrState as EscloudInstanceV2Args | undefined;
            if ((!args || args.adminPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adminPassword'");
            }
            if ((!args || args.chargeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'chargeType'");
            }
            if ((!args || args.configurationCode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configurationCode'");
            }
            if ((!args || args.enableHttps === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enableHttps'");
            }
            if ((!args || args.instanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceName'");
            }
            if ((!args || args.nodeSpecsAssigns === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeSpecsAssigns'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            if ((!args || args.zoneIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneIds'");
            }
            resourceInputs["adminPassword"] = args?.adminPassword ? pulumi.secret(args.adminPassword) : undefined;
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["chargeType"] = args ? args.chargeType : undefined;
            resourceInputs["configurationCode"] = args ? args.configurationCode : undefined;
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["enableHttps"] = args ? args.enableHttps : undefined;
            resourceInputs["enablePureMaster"] = args ? args.enablePureMaster : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["maintenanceDays"] = args ? args.maintenanceDays : undefined;
            resourceInputs["maintenanceTime"] = args ? args.maintenanceTime : undefined;
            resourceInputs["networkSpecs"] = args ? args.networkSpecs : undefined;
            resourceInputs["nodeSpecsAssigns"] = args ? args.nodeSpecsAssigns : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["zoneIds"] = args ? args.zoneIds : undefined;
            resourceInputs["cerebroPrivateDomain"] = undefined /*out*/;
            resourceInputs["cerebroPublicDomain"] = undefined /*out*/;
            resourceInputs["esEip"] = undefined /*out*/;
            resourceInputs["esEipId"] = undefined /*out*/;
            resourceInputs["esPrivateDomain"] = undefined /*out*/;
            resourceInputs["esPrivateEndpoint"] = undefined /*out*/;
            resourceInputs["esPrivateIpWhitelist"] = undefined /*out*/;
            resourceInputs["esPublicDomain"] = undefined /*out*/;
            resourceInputs["esPublicEndpoint"] = undefined /*out*/;
            resourceInputs["esPublicIpWhitelist"] = undefined /*out*/;
            resourceInputs["kibanaEip"] = undefined /*out*/;
            resourceInputs["kibanaEipId"] = undefined /*out*/;
            resourceInputs["kibanaPrivateDomain"] = undefined /*out*/;
            resourceInputs["kibanaPrivateIpWhitelist"] = undefined /*out*/;
            resourceInputs["kibanaPublicDomain"] = undefined /*out*/;
            resourceInputs["kibanaPublicIpWhitelist"] = undefined /*out*/;
            resourceInputs["mainZoneId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["adminPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(EscloudInstanceV2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EscloudInstanceV2 resources.
 */
export interface EscloudInstanceV2State {
    /**
     * The password of administrator account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    adminPassword?: pulumi.Input<string>;
    /**
     * Whether to automatically renew in prepaid scenarios. Default is false.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * The cerebro private domain of instance.
     */
    cerebroPrivateDomain?: pulumi.Input<string>;
    /**
     * The cerebro public domain of instance.
     */
    cerebroPublicDomain?: pulumi.Input<string>;
    /**
     * The charge type of ESCloud instance, valid values: `PostPaid`, `PrePaid`.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Configuration code used for billing.
     */
    configurationCode?: pulumi.Input<string>;
    /**
     * Whether enable deletion protection for ESCloud instance. Default is false.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * Whether Https access is enabled.
     */
    enableHttps?: pulumi.Input<boolean>;
    /**
     * Whether the Master node is independent.
     */
    enablePureMaster?: pulumi.Input<boolean>;
    /**
     * The eip address of instance.
     */
    esEip?: pulumi.Input<string>;
    /**
     * The eip id associated with the instance.
     */
    esEipId?: pulumi.Input<string>;
    /**
     * The es private domain of instance.
     */
    esPrivateDomain?: pulumi.Input<string>;
    /**
     * The es private endpoint of instance.
     */
    esPrivateEndpoint?: pulumi.Input<string>;
    /**
     * The whitelist of es private ip.
     */
    esPrivateIpWhitelist?: pulumi.Input<string>;
    /**
     * The es public domain of instance.
     */
    esPublicDomain?: pulumi.Input<string>;
    /**
     * The es public endpoint of instance.
     */
    esPublicEndpoint?: pulumi.Input<string>;
    /**
     * The whitelist of es public ip.
     */
    esPublicIpWhitelist?: pulumi.Input<string>;
    /**
     * The name of ESCloud instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The eip address of kibana.
     */
    kibanaEip?: pulumi.Input<string>;
    /**
     * The eip id associated with kibana.
     */
    kibanaEipId?: pulumi.Input<string>;
    /**
     * The kibana private domain of instance.
     */
    kibanaPrivateDomain?: pulumi.Input<string>;
    /**
     * The whitelist of kibana private ip.
     */
    kibanaPrivateIpWhitelist?: pulumi.Input<string>;
    /**
     * The kibana public domain of instance.
     */
    kibanaPublicDomain?: pulumi.Input<string>;
    /**
     * The whitelist of kibana public ip.
     */
    kibanaPublicIpWhitelist?: pulumi.Input<string>;
    /**
     * The main zone id of instance.
     */
    mainZoneId?: pulumi.Input<string>;
    /**
     * The maintainable day for the instance. Valid values: `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`. Works only on modified scenes.
     */
    maintenanceDays?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maintainable time period for the instance. Works only on modified scenes.
     */
    maintenanceTime?: pulumi.Input<string>;
    /**
     * The public network config of the ESCloud instance.
     */
    networkSpecs?: pulumi.Input<pulumi.Input<inputs.escloud_v2.EscloudInstanceV2NetworkSpec>[]>;
    /**
     * The number and configuration of various ESCloud instance node. Kibana NodeSpecsAssign should not be modified.
     */
    nodeSpecsAssigns?: pulumi.Input<pulumi.Input<inputs.escloud_v2.EscloudInstanceV2NodeSpecsAssign>[]>;
    /**
     * Purchase duration in prepaid scenarios. Unit: Monthly.
     */
    period?: pulumi.Input<number>;
    /**
     * The project name to which the ESCloud instance belongs.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The status of instance.
     */
    status?: pulumi.Input<string>;
    /**
     * The id of subnet, the subnet must belong to the AZ selected.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.escloud_v2.EscloudInstanceV2Tag>[]>;
    /**
     * The version of instance. When creating ESCloud instance, the valid value is `V6_7` or `V7_10`. When creating OpenSearch instance, the valid value is `OPEN_SEARCH_2_9`.
     */
    version?: pulumi.Input<string>;
    /**
     * The zone id of the ESCloud instance. Support specifying multiple availability zones.
     * The first zone id is the primary availability zone, while the rest are backup availability zones.
     */
    zoneIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a EscloudInstanceV2 resource.
 */
export interface EscloudInstanceV2Args {
    /**
     * The password of administrator account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    adminPassword: pulumi.Input<string>;
    /**
     * Whether to automatically renew in prepaid scenarios. Default is false.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * The charge type of ESCloud instance, valid values: `PostPaid`, `PrePaid`.
     */
    chargeType: pulumi.Input<string>;
    /**
     * Configuration code used for billing.
     */
    configurationCode: pulumi.Input<string>;
    /**
     * Whether enable deletion protection for ESCloud instance. Default is false.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * Whether Https access is enabled.
     */
    enableHttps: pulumi.Input<boolean>;
    /**
     * Whether the Master node is independent.
     */
    enablePureMaster?: pulumi.Input<boolean>;
    /**
     * The name of ESCloud instance.
     */
    instanceName: pulumi.Input<string>;
    /**
     * The maintainable day for the instance. Valid values: `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`. Works only on modified scenes.
     */
    maintenanceDays?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maintainable time period for the instance. Works only on modified scenes.
     */
    maintenanceTime?: pulumi.Input<string>;
    /**
     * The public network config of the ESCloud instance.
     */
    networkSpecs?: pulumi.Input<pulumi.Input<inputs.escloud_v2.EscloudInstanceV2NetworkSpec>[]>;
    /**
     * The number and configuration of various ESCloud instance node. Kibana NodeSpecsAssign should not be modified.
     */
    nodeSpecsAssigns: pulumi.Input<pulumi.Input<inputs.escloud_v2.EscloudInstanceV2NodeSpecsAssign>[]>;
    /**
     * Purchase duration in prepaid scenarios. Unit: Monthly.
     */
    period?: pulumi.Input<number>;
    /**
     * The project name to which the ESCloud instance belongs.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The id of subnet, the subnet must belong to the AZ selected.
     */
    subnetId: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.escloud_v2.EscloudInstanceV2Tag>[]>;
    /**
     * The version of instance. When creating ESCloud instance, the valid value is `V6_7` or `V7_10`. When creating OpenSearch instance, the valid value is `OPEN_SEARCH_2_9`.
     */
    version: pulumi.Input<string>;
    /**
     * The zone id of the ESCloud instance. Support specifying multiple availability zones.
     * The first zone id is the primary availability zone, while the rest are backup availability zones.
     */
    zoneIds: pulumi.Input<pulumi.Input<string>[]>;
}
