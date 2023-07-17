// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage redis backup
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultBackup = new volcengine.redis.Backup("default", {
 *     instanceId: "redis-cnlfvrv4qye6u4lpa",
 * });
 * ```
 *
 * ## Import
 *
 * Redis Backup can be imported using the instanceId:backupId, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:redis/backup:Backup default redis-cn02aqusft7ws****:b-cn02xmmrp751i9cdzcphjmk4****
 * ```
 */
export class Backup extends pulumi.CustomResource {
    /**
     * Get an existing Backup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackupState, opts?: pulumi.CustomResourceOptions): Backup {
        return new Backup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:redis/backup:Backup';

    /**
     * Returns true if the given object is an instance of Backup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Backup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Backup.__pulumiType;
    }

    /**
     * The id of backup point.
     */
    public /*out*/ readonly backupPointId!: pulumi.Output<string>;
    /**
     * Backup strategy.
     */
    public /*out*/ readonly backupStrategy!: pulumi.Output<string>;
    /**
     * Backup type.
     */
    public /*out*/ readonly backupType!: pulumi.Output<string>;
    /**
     * End time of backup.
     */
    public /*out*/ readonly endTime!: pulumi.Output<string>;
    /**
     * Information of instance.
     */
    public /*out*/ readonly instanceDetails!: pulumi.Output<outputs.redis.BackupInstanceDetail[]>;
    /**
     * Id of instance to create backup.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Size in MiB.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * Start time of backup.
     */
    public /*out*/ readonly startTime!: pulumi.Output<string>;
    /**
     * Status of backup (Creating/Available/Unavailable/Deleting).
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Backup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackupArgs | BackupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackupState | undefined;
            resourceInputs["backupPointId"] = state ? state.backupPointId : undefined;
            resourceInputs["backupStrategy"] = state ? state.backupStrategy : undefined;
            resourceInputs["backupType"] = state ? state.backupType : undefined;
            resourceInputs["endTime"] = state ? state.endTime : undefined;
            resourceInputs["instanceDetails"] = state ? state.instanceDetails : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as BackupArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["backupPointId"] = undefined /*out*/;
            resourceInputs["backupStrategy"] = undefined /*out*/;
            resourceInputs["backupType"] = undefined /*out*/;
            resourceInputs["endTime"] = undefined /*out*/;
            resourceInputs["instanceDetails"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["startTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Backup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Backup resources.
 */
export interface BackupState {
    /**
     * The id of backup point.
     */
    backupPointId?: pulumi.Input<string>;
    /**
     * Backup strategy.
     */
    backupStrategy?: pulumi.Input<string>;
    /**
     * Backup type.
     */
    backupType?: pulumi.Input<string>;
    /**
     * End time of backup.
     */
    endTime?: pulumi.Input<string>;
    /**
     * Information of instance.
     */
    instanceDetails?: pulumi.Input<pulumi.Input<inputs.redis.BackupInstanceDetail>[]>;
    /**
     * Id of instance to create backup.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Size in MiB.
     */
    size?: pulumi.Input<number>;
    /**
     * Start time of backup.
     */
    startTime?: pulumi.Input<string>;
    /**
     * Status of backup (Creating/Available/Unavailable/Deleting).
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Backup resource.
 */
export interface BackupArgs {
    /**
     * Id of instance to create backup.
     */
    instanceId: pulumi.Input<string>;
}