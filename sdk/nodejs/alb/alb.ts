// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage alb
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = volcengine.alb.getZones({});
 * const vpcIpv6 = new volcengine.vpc.Vpc("vpcIpv6", {
 *     vpcName: "acc-test-vpc-ipv6",
 *     cidrBlock: "172.16.0.0/16",
 *     enableIpv6: true,
 * });
 * const subnetIpv61 = new volcengine.vpc.Subnet("subnetIpv61", {
 *     subnetName: "acc-test-subnet-ipv6-1",
 *     cidrBlock: "172.16.1.0/24",
 *     zoneId: foo.then(foo => foo.zones?.[0]?.id),
 *     vpcId: vpcIpv6.id,
 *     ipv6CidrBlock: 1,
 * });
 * const subnetIpv62 = new volcengine.vpc.Subnet("subnetIpv62", {
 *     subnetName: "acc-test-subnet-ipv6-2",
 *     cidrBlock: "172.16.2.0/24",
 *     zoneId: foo.then(foo => foo.zones?.[1]?.id),
 *     vpcId: vpcIpv6.id,
 *     ipv6CidrBlock: 2,
 * });
 * const ipv6Gateway = new volcengine.vpc.Ipv6Gateway("ipv6Gateway", {vpcId: vpcIpv6.id});
 * const alb_private = new volcengine.alb.Alb("alb-private", {
 *     addressIpVersion: "IPv4",
 *     type: "private",
 *     loadBalancerName: "acc-test-alb-private",
 *     description: "acc-test",
 *     subnetIds: [
 *         subnetIpv61.id,
 *         subnetIpv62.id,
 *     ],
 *     projectName: "default",
 *     deleteProtection: "off",
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * const alb_public = new volcengine.alb.Alb("alb-public", {
 *     addressIpVersion: "DualStack",
 *     type: "public",
 *     loadBalancerName: "acc-test-alb-public",
 *     description: "acc-test",
 *     subnetIds: [
 *         subnetIpv61.id,
 *         subnetIpv62.id,
 *     ],
 *     projectName: "default",
 *     deleteProtection: "off",
 *     eipBillingConfig: {
 *         isp: "BGP",
 *         eipBillingType: "PostPaidByBandwidth",
 *         bandwidth: 1,
 *     },
 *     ipv6EipBillingConfig: {
 *         isp: "BGP",
 *         billingType: "PostPaidByBandwidth",
 *         bandwidth: 1,
 *     },
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * }, {
 *     dependsOn: [ipv6Gateway],
 * });
 * ```
 *
 * ## Import
 *
 * Alb can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:alb/alb:Alb default resource_id
 * ```
 */
export class Alb extends pulumi.CustomResource {
    /**
     * Get an existing Alb resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AlbState, opts?: pulumi.CustomResourceOptions): Alb {
        return new Alb(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:alb/alb:Alb';

    /**
     * Returns true if the given object is an instance of Alb.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Alb {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Alb.__pulumiType;
    }

    /**
     * The address ip version of the Alb. Valid values: `IPv4`, `DualStack`. Default is `ipv4`.
     */
    public readonly addressIpVersion!: pulumi.Output<string | undefined>;
    /**
     * Whether to enable the delete protection function of the Alb. Valid values: `on`, `off`. Default is `off`.
     */
    public readonly deleteProtection!: pulumi.Output<string | undefined>;
    /**
     * The description of the Alb.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The DNS name.
     */
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    /**
     * The billing configuration of the EIP which automatically associated to the Alb. This field is valid when the type of the Alb is `public`.When the type of the Alb is `private`, suggest using a combination of resource `volcengine.eip.Address` and `volcengine.eip.Associate` to achieve public network access function.
     */
    public readonly eipBillingConfig!: pulumi.Output<outputs.alb.AlbEipBillingConfig>;
    /**
     * The billing configuration of the Ipv6 EIP which automatically associated to the Alb. This field is required when the type of the Alb is `public`.When the type of the Alb is `private`, suggest using a combination of resource `volcengine.vpc.Ipv6Gateway` and `volcengine.vpc.Ipv6AddressBandwidth` to achieve ipv6 public network access function.
     */
    public readonly ipv6EipBillingConfig!: pulumi.Output<outputs.alb.AlbIpv6EipBillingConfig>;
    /**
     * The name of the Alb.
     */
    public readonly loadBalancerName!: pulumi.Output<string>;
    /**
     * The local addresses of the Alb.
     */
    public /*out*/ readonly localAddresses!: pulumi.Output<string[]>;
    /**
     * The ProjectName of the Alb.
     */
    public readonly projectName!: pulumi.Output<string>;
    /**
     * The status of the Alb.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The id of the Subnet.
     */
    public readonly subnetIds!: pulumi.Output<string[]>;
    /**
     * Tags.
     */
    public readonly tags!: pulumi.Output<outputs.alb.AlbTag[] | undefined>;
    /**
     * The type of the Alb. Valid values: `public`, `private`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The vpc id of the Alb.
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;
    /**
     * Configuration information of the Alb instance in different Availability Zones.
     */
    public /*out*/ readonly zoneMappings!: pulumi.Output<outputs.alb.AlbZoneMapping[]>;

    /**
     * Create a Alb resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlbArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AlbArgs | AlbState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AlbState | undefined;
            resourceInputs["addressIpVersion"] = state ? state.addressIpVersion : undefined;
            resourceInputs["deleteProtection"] = state ? state.deleteProtection : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dnsName"] = state ? state.dnsName : undefined;
            resourceInputs["eipBillingConfig"] = state ? state.eipBillingConfig : undefined;
            resourceInputs["ipv6EipBillingConfig"] = state ? state.ipv6EipBillingConfig : undefined;
            resourceInputs["loadBalancerName"] = state ? state.loadBalancerName : undefined;
            resourceInputs["localAddresses"] = state ? state.localAddresses : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnetIds"] = state ? state.subnetIds : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["zoneMappings"] = state ? state.zoneMappings : undefined;
        } else {
            const args = argsOrState as AlbArgs | undefined;
            if ((!args || args.subnetIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetIds'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["addressIpVersion"] = args ? args.addressIpVersion : undefined;
            resourceInputs["deleteProtection"] = args ? args.deleteProtection : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["eipBillingConfig"] = args ? args.eipBillingConfig : undefined;
            resourceInputs["ipv6EipBillingConfig"] = args ? args.ipv6EipBillingConfig : undefined;
            resourceInputs["loadBalancerName"] = args ? args.loadBalancerName : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["subnetIds"] = args ? args.subnetIds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["dnsName"] = undefined /*out*/;
            resourceInputs["localAddresses"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
            resourceInputs["zoneMappings"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Alb.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Alb resources.
 */
export interface AlbState {
    /**
     * The address ip version of the Alb. Valid values: `IPv4`, `DualStack`. Default is `ipv4`.
     */
    addressIpVersion?: pulumi.Input<string>;
    /**
     * Whether to enable the delete protection function of the Alb. Valid values: `on`, `off`. Default is `off`.
     */
    deleteProtection?: pulumi.Input<string>;
    /**
     * The description of the Alb.
     */
    description?: pulumi.Input<string>;
    /**
     * The DNS name.
     */
    dnsName?: pulumi.Input<string>;
    /**
     * The billing configuration of the EIP which automatically associated to the Alb. This field is valid when the type of the Alb is `public`.When the type of the Alb is `private`, suggest using a combination of resource `volcengine.eip.Address` and `volcengine.eip.Associate` to achieve public network access function.
     */
    eipBillingConfig?: pulumi.Input<inputs.alb.AlbEipBillingConfig>;
    /**
     * The billing configuration of the Ipv6 EIP which automatically associated to the Alb. This field is required when the type of the Alb is `public`.When the type of the Alb is `private`, suggest using a combination of resource `volcengine.vpc.Ipv6Gateway` and `volcengine.vpc.Ipv6AddressBandwidth` to achieve ipv6 public network access function.
     */
    ipv6EipBillingConfig?: pulumi.Input<inputs.alb.AlbIpv6EipBillingConfig>;
    /**
     * The name of the Alb.
     */
    loadBalancerName?: pulumi.Input<string>;
    /**
     * The local addresses of the Alb.
     */
    localAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ProjectName of the Alb.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The status of the Alb.
     */
    status?: pulumi.Input<string>;
    /**
     * The id of the Subnet.
     */
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.alb.AlbTag>[]>;
    /**
     * The type of the Alb. Valid values: `public`, `private`.
     */
    type?: pulumi.Input<string>;
    /**
     * The vpc id of the Alb.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * Configuration information of the Alb instance in different Availability Zones.
     */
    zoneMappings?: pulumi.Input<pulumi.Input<inputs.alb.AlbZoneMapping>[]>;
}

/**
 * The set of arguments for constructing a Alb resource.
 */
export interface AlbArgs {
    /**
     * The address ip version of the Alb. Valid values: `IPv4`, `DualStack`. Default is `ipv4`.
     */
    addressIpVersion?: pulumi.Input<string>;
    /**
     * Whether to enable the delete protection function of the Alb. Valid values: `on`, `off`. Default is `off`.
     */
    deleteProtection?: pulumi.Input<string>;
    /**
     * The description of the Alb.
     */
    description?: pulumi.Input<string>;
    /**
     * The billing configuration of the EIP which automatically associated to the Alb. This field is valid when the type of the Alb is `public`.When the type of the Alb is `private`, suggest using a combination of resource `volcengine.eip.Address` and `volcengine.eip.Associate` to achieve public network access function.
     */
    eipBillingConfig?: pulumi.Input<inputs.alb.AlbEipBillingConfig>;
    /**
     * The billing configuration of the Ipv6 EIP which automatically associated to the Alb. This field is required when the type of the Alb is `public`.When the type of the Alb is `private`, suggest using a combination of resource `volcengine.vpc.Ipv6Gateway` and `volcengine.vpc.Ipv6AddressBandwidth` to achieve ipv6 public network access function.
     */
    ipv6EipBillingConfig?: pulumi.Input<inputs.alb.AlbIpv6EipBillingConfig>;
    /**
     * The name of the Alb.
     */
    loadBalancerName?: pulumi.Input<string>;
    /**
     * The ProjectName of the Alb.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The id of the Subnet.
     */
    subnetIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.alb.AlbTag>[]>;
    /**
     * The type of the Alb. Valid values: `public`, `private`.
     */
    type: pulumi.Input<string>;
}
