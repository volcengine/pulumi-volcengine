// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_mysql
{
    /// <summary>
    /// Provides a resource to manage rds mysql backup
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Volcengine = Pulumi.Volcengine;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var foo = new Volcengine.Rds_mysql.Backup("foo", new()
    ///     {
    ///         BackupMetas = new[]
    ///         {
    ///             new Volcengine.Rds_mysql.Inputs.BackupBackupMetaArgs
    ///             {
    ///                 DbName = "order",
    ///             },
    ///         },
    ///         BackupMethod = "Logical",
    ///         InstanceId = "mysql-c8c3f45c4b07",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RdsMysqlBackup can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:rds_mysql/backup:Backup default instanceId:backupId
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:rds_mysql/backup:Backup")]
    public partial class Backup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The id of the backup.
        /// </summary>
        [Output("backupId")]
        public Output<string> BackupId { get; private set; } = null!;

        /// <summary>
        /// When creating a library table backup of logical backup type, it is used to specify the library table information to be backed up.
        /// Prerequisite: When the value of BackupMethod is Logical, and the BackupType field is not passed.
        /// Mutual exclusion situation: When the value of the BackupType field is DumpAll, this field is not effective.
        /// Quantity limit: When creating a specified library table backup, the upper limit of the number of libraries is 5000, and the upper limit of the number of tables in each library is 5000.
        /// </summary>
        [Output("backupMetas")]
        public Output<ImmutableArray<Outputs.BackupBackupMeta>> BackupMetas { get; private set; } = null!;

        /// <summary>
        /// Backup method. Value range: Full, full backup under physical backup type. Default value. DumpAll: full database backup under logical backup type. Prerequisite: If you need to create a full database backup of logical backup type, that is, when the value of BackupType is DumpAll, the backup type should be set to logical backup, that is, the value of BackupMethod should be Logical. If you need to create a database table backup of logical backup type, you do not need to pass in this field. You only need to specify the database and table to be backed up in the BackupMeta field.
        /// </summary>
        [Output("backupMethod")]
        public Output<string> BackupMethod { get; private set; } = null!;

        /// <summary>
        /// Backup type. Currently, only full backup is supported. The value is Full.
        /// </summary>
        [Output("backupType")]
        public Output<string> BackupType { get; private set; } = null!;

        /// <summary>
        /// The id of the instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a Backup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Backup(string name, BackupArgs args, CustomResourceOptions? options = null)
            : base("volcengine:rds_mysql/backup:Backup", name, args ?? new BackupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Backup(string name, Input<string> id, BackupState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:rds_mysql/backup:Backup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/volcengine",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Backup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Backup Get(string name, Input<string> id, BackupState? state = null, CustomResourceOptions? options = null)
        {
            return new Backup(name, id, state, options);
        }
    }

    public sealed class BackupArgs : global::Pulumi.ResourceArgs
    {
        [Input("backupMetas")]
        private InputList<Inputs.BackupBackupMetaArgs>? _backupMetas;

        /// <summary>
        /// When creating a library table backup of logical backup type, it is used to specify the library table information to be backed up.
        /// Prerequisite: When the value of BackupMethod is Logical, and the BackupType field is not passed.
        /// Mutual exclusion situation: When the value of the BackupType field is DumpAll, this field is not effective.
        /// Quantity limit: When creating a specified library table backup, the upper limit of the number of libraries is 5000, and the upper limit of the number of tables in each library is 5000.
        /// </summary>
        public InputList<Inputs.BackupBackupMetaArgs> BackupMetas
        {
            get => _backupMetas ?? (_backupMetas = new InputList<Inputs.BackupBackupMetaArgs>());
            set => _backupMetas = value;
        }

        /// <summary>
        /// Backup method. Value range: Full, full backup under physical backup type. Default value. DumpAll: full database backup under logical backup type. Prerequisite: If you need to create a full database backup of logical backup type, that is, when the value of BackupType is DumpAll, the backup type should be set to logical backup, that is, the value of BackupMethod should be Logical. If you need to create a database table backup of logical backup type, you do not need to pass in this field. You only need to specify the database and table to be backed up in the BackupMeta field.
        /// </summary>
        [Input("backupMethod")]
        public Input<string>? BackupMethod { get; set; }

        /// <summary>
        /// Backup type. Currently, only full backup is supported. The value is Full.
        /// </summary>
        [Input("backupType")]
        public Input<string>? BackupType { get; set; }

        /// <summary>
        /// The id of the instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public BackupArgs()
        {
        }
        public static new BackupArgs Empty => new BackupArgs();
    }

    public sealed class BackupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the backup.
        /// </summary>
        [Input("backupId")]
        public Input<string>? BackupId { get; set; }

        [Input("backupMetas")]
        private InputList<Inputs.BackupBackupMetaGetArgs>? _backupMetas;

        /// <summary>
        /// When creating a library table backup of logical backup type, it is used to specify the library table information to be backed up.
        /// Prerequisite: When the value of BackupMethod is Logical, and the BackupType field is not passed.
        /// Mutual exclusion situation: When the value of the BackupType field is DumpAll, this field is not effective.
        /// Quantity limit: When creating a specified library table backup, the upper limit of the number of libraries is 5000, and the upper limit of the number of tables in each library is 5000.
        /// </summary>
        public InputList<Inputs.BackupBackupMetaGetArgs> BackupMetas
        {
            get => _backupMetas ?? (_backupMetas = new InputList<Inputs.BackupBackupMetaGetArgs>());
            set => _backupMetas = value;
        }

        /// <summary>
        /// Backup method. Value range: Full, full backup under physical backup type. Default value. DumpAll: full database backup under logical backup type. Prerequisite: If you need to create a full database backup of logical backup type, that is, when the value of BackupType is DumpAll, the backup type should be set to logical backup, that is, the value of BackupMethod should be Logical. If you need to create a database table backup of logical backup type, you do not need to pass in this field. You only need to specify the database and table to be backed up in the BackupMeta field.
        /// </summary>
        [Input("backupMethod")]
        public Input<string>? BackupMethod { get; set; }

        /// <summary>
        /// Backup type. Currently, only full backup is supported. The value is Full.
        /// </summary>
        [Input("backupType")]
        public Input<string>? BackupType { get; set; }

        /// <summary>
        /// The id of the instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public BackupState()
        {
        }
        public static new BackupState Empty => new BackupState();
    }
}
