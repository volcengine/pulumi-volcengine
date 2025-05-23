// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage nas mount point
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooZones = volcengine.nas.getZones({});
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
 * const fooPermissionGroup = new volcengine.nas.PermissionGroup("fooPermissionGroup", {
 *     permissionGroupName: "acc-test",
 *     description: "acctest",
 *     permissionRules: [
 *         {
 *             cidrIp: "*",
 *             rwMode: "RW",
 *             useMode: "All_squash",
 *         },
 *         {
 *             cidrIp: "192.168.0.0",
 *             rwMode: "RO",
 *             useMode: "All_squash",
 *         },
 *     ],
 * });
 * const fooFileSystem = new volcengine.nas.FileSystem("fooFileSystem", {
 *     fileSystemName: "acc-test-fs",
 *     description: "acc-test",
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *     capacity: 103,
 *     projectName: "default",
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * const fooMountPoint = new volcengine.nas.MountPoint("fooMountPoint", {
 *     fileSystemId: fooFileSystem.id,
 *     mountPointName: "acc-test",
 *     permissionGroupId: fooPermissionGroup.id,
 *     subnetId: fooSubnet.id,
 * });
 * ```
 *
 * ## Import
 *
 * Nas Mount Point can be imported using the file system id and mount point id, e.g.
 *
 * ```sh
 * $ pulumi import volcengine:nas/mountPoint:MountPoint default enas-cnbj18bcb923****:mount-a6ee****
 * ```
 */
export class MountPoint extends pulumi.CustomResource {
    /**
     * Get an existing MountPoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MountPointState, opts?: pulumi.CustomResourceOptions): MountPoint {
        return new MountPoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:nas/mountPoint:MountPoint';

    /**
     * Returns true if the given object is an instance of MountPoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MountPoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MountPoint.__pulumiType;
    }

    /**
     * The file system id.
     */
    public readonly fileSystemId!: pulumi.Output<string>;
    /**
     * The mount point id.
     */
    public /*out*/ readonly mountPointId!: pulumi.Output<string>;
    /**
     * The mount point name.
     */
    public readonly mountPointName!: pulumi.Output<string>;
    /**
     * The permission group id.
     */
    public readonly permissionGroupId!: pulumi.Output<string>;
    /**
     * The subnet id.
     */
    public readonly subnetId!: pulumi.Output<string>;

    /**
     * Create a MountPoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MountPointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MountPointArgs | MountPointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MountPointState | undefined;
            resourceInputs["fileSystemId"] = state ? state.fileSystemId : undefined;
            resourceInputs["mountPointId"] = state ? state.mountPointId : undefined;
            resourceInputs["mountPointName"] = state ? state.mountPointName : undefined;
            resourceInputs["permissionGroupId"] = state ? state.permissionGroupId : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
        } else {
            const args = argsOrState as MountPointArgs | undefined;
            if ((!args || args.fileSystemId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fileSystemId'");
            }
            if ((!args || args.mountPointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mountPointName'");
            }
            if ((!args || args.permissionGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissionGroupId'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            resourceInputs["fileSystemId"] = args ? args.fileSystemId : undefined;
            resourceInputs["mountPointName"] = args ? args.mountPointName : undefined;
            resourceInputs["permissionGroupId"] = args ? args.permissionGroupId : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["mountPointId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MountPoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MountPoint resources.
 */
export interface MountPointState {
    /**
     * The file system id.
     */
    fileSystemId?: pulumi.Input<string>;
    /**
     * The mount point id.
     */
    mountPointId?: pulumi.Input<string>;
    /**
     * The mount point name.
     */
    mountPointName?: pulumi.Input<string>;
    /**
     * The permission group id.
     */
    permissionGroupId?: pulumi.Input<string>;
    /**
     * The subnet id.
     */
    subnetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MountPoint resource.
 */
export interface MountPointArgs {
    /**
     * The file system id.
     */
    fileSystemId: pulumi.Input<string>;
    /**
     * The mount point name.
     */
    mountPointName: pulumi.Input<string>;
    /**
     * The permission group id.
     */
    permissionGroupId: pulumi.Input<string>;
    /**
     * The subnet id.
     */
    subnetId: pulumi.Input<string>;
}
