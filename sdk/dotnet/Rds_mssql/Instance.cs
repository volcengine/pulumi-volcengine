// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_mssql
{
    /// <summary>
    /// Provides a resource to manage rds mssql instance
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
    ///     var fooZones = Volcengine.Ecs.Zones.Invoke();
    /// 
    ///     var fooVpc = new Volcengine.Vpc.Vpc("fooVpc", new()
    ///     {
    ///         VpcName = "acc-test-vpc",
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var fooSubnet = new Volcengine.Vpc.Subnet("fooSubnet", new()
    ///     {
    ///         SubnetName = "acc-test-subnet",
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[0]?.Id),
    ///         VpcId = fooVpc.Id,
    ///     });
    /// 
    ///     var fooInstance = new Volcengine.Rds_mssql.Instance("fooInstance", new()
    ///     {
    ///         DbEngineVersion = "SQLServer_2019_Std",
    ///         InstanceType = "HA",
    ///         NodeSpec = "rds.mssql.se.ha.d2.2c4g",
    ///         StorageSpace = 20,
    ///         SubnetIds = new[]
    ///         {
    ///             fooSubnet.Id,
    ///         },
    ///         SuperAccountPassword = "Tftest110",
    ///         InstanceName = "acc-test-mssql",
    ///         ProjectName = "default",
    ///         ChargeInfo = new Volcengine.Rds_mssql.Inputs.InstanceChargeInfoArgs
    ///         {
    ///             ChargeType = "PostPaid",
    ///         },
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Rds_mssql.Inputs.InstanceTagArgs
    ///             {
    ///                 Key = "k1",
    ///                 Value = "v1",
    ///             },
    ///         },
    ///         BackupTime = "18:00Z-19:00Z",
    ///         FullBackupPeriods = new[]
    ///         {
    ///             "Monday",
    ///             "Tuesday",
    ///         },
    ///         BackupRetentionPeriod = 14,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Rds Mssql Instance can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:rds_mssql/instance:Instance default resource_id
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:rds_mssql/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Data backup retention days, value range: 7~30. 
        /// This field is valid and required when updating the backup plan of instance.
        /// </summary>
        [Output("backupRetentionPeriod")]
        public Output<int?> BackupRetentionPeriod { get; private set; } = null!;

        /// <summary>
        /// The time window for starting the backup task is one hour interval. 
        /// This field is valid and required when updating the backup plan of instance.
        /// </summary>
        [Output("backupTime")]
        public Output<string?> BackupTime { get; private set; } = null!;

        /// <summary>
        /// The charge info.
        /// </summary>
        [Output("chargeInfo")]
        public Output<Outputs.InstanceChargeInfo> ChargeInfo { get; private set; } = null!;

        /// <summary>
        /// The Compatible version. Valid values: `SQLServer_2019_Std`, `SQLServer_2019_Web`, `SQLServer_2019_Ent`.
        /// </summary>
        [Output("dbEngineVersion")]
        public Output<string> DbEngineVersion { get; private set; } = null!;

        /// <summary>
        /// Full backup cycle. Multiple values separated by commas. The values are as follows: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. 
        /// This field is valid and required when updating the backup plan of instance.
        /// </summary>
        [Output("fullBackupPeriods")]
        public Output<ImmutableArray<string>> FullBackupPeriods { get; private set; } = null!;

        /// <summary>
        /// Name of the instance.
        /// </summary>
        [Output("instanceName")]
        public Output<string?> InstanceName { get; private set; } = null!;

        /// <summary>
        /// The Instance type. When the value of the `db_engine_version` is `SQLServer_2019_Std`, the value of this field can be `HA` or `Basic`.When the value of the `db_engine_version` is `SQLServer_2019_Ent`, the value of this field can be `Cluster` or `Basic`.When the value of the `db_engine_version` is `SQLServer_2019_Web`, the value of this field can be `Basic`.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The node specification.
        /// </summary>
        [Output("nodeSpec")]
        public Output<string> NodeSpec { get; private set; } = null!;

        /// <summary>
        /// The project name.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// Storage space size, measured in GiB. The range of values is 20GiB to 4000GiB, with a step size of 10GiB.
        /// </summary>
        [Output("storageSpace")]
        public Output<int> StorageSpace { get; private set; } = null!;

        /// <summary>
        /// The subnet id of the instance node. When creating an instance that includes primary and backup nodes and needs to deploy primary and backup nodes across availability zones, you can specify two subnet_id. By default, the first is the primary node availability zone, and the second is the backup node availability zone.
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// The super account password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Output("superAccountPassword")]
        public Output<string> SuperAccountPassword { get; private set; } = null!;

        /// <summary>
        /// Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.InstanceTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("volcengine:rds_mssql/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:rds_mssql/instance:Instance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/volcengine",
                AdditionalSecretOutputs =
                {
                    "superAccountPassword",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Data backup retention days, value range: 7~30. 
        /// This field is valid and required when updating the backup plan of instance.
        /// </summary>
        [Input("backupRetentionPeriod")]
        public Input<int>? BackupRetentionPeriod { get; set; }

        /// <summary>
        /// The time window for starting the backup task is one hour interval. 
        /// This field is valid and required when updating the backup plan of instance.
        /// </summary>
        [Input("backupTime")]
        public Input<string>? BackupTime { get; set; }

        /// <summary>
        /// The charge info.
        /// </summary>
        [Input("chargeInfo", required: true)]
        public Input<Inputs.InstanceChargeInfoArgs> ChargeInfo { get; set; } = null!;

        /// <summary>
        /// The Compatible version. Valid values: `SQLServer_2019_Std`, `SQLServer_2019_Web`, `SQLServer_2019_Ent`.
        /// </summary>
        [Input("dbEngineVersion", required: true)]
        public Input<string> DbEngineVersion { get; set; } = null!;

        [Input("fullBackupPeriods")]
        private InputList<string>? _fullBackupPeriods;

        /// <summary>
        /// Full backup cycle. Multiple values separated by commas. The values are as follows: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. 
        /// This field is valid and required when updating the backup plan of instance.
        /// </summary>
        public InputList<string> FullBackupPeriods
        {
            get => _fullBackupPeriods ?? (_fullBackupPeriods = new InputList<string>());
            set => _fullBackupPeriods = value;
        }

        /// <summary>
        /// Name of the instance.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The Instance type. When the value of the `db_engine_version` is `SQLServer_2019_Std`, the value of this field can be `HA` or `Basic`.When the value of the `db_engine_version` is `SQLServer_2019_Ent`, the value of this field can be `Cluster` or `Basic`.When the value of the `db_engine_version` is `SQLServer_2019_Web`, the value of this field can be `Basic`.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// The node specification.
        /// </summary>
        [Input("nodeSpec", required: true)]
        public Input<string> NodeSpec { get; set; } = null!;

        /// <summary>
        /// The project name.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// Storage space size, measured in GiB. The range of values is 20GiB to 4000GiB, with a step size of 10GiB.
        /// </summary>
        [Input("storageSpace", required: true)]
        public Input<int> StorageSpace { get; set; } = null!;

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The subnet id of the instance node. When creating an instance that includes primary and backup nodes and needs to deploy primary and backup nodes across availability zones, you can specify two subnet_id. By default, the first is the primary node availability zone, and the second is the backup node availability zone.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("superAccountPassword", required: true)]
        private Input<string>? _superAccountPassword;

        /// <summary>
        /// The super account password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        public Input<string>? SuperAccountPassword
        {
            get => _superAccountPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _superAccountPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("tags")]
        private InputList<Inputs.InstanceTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.InstanceTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.InstanceTagArgs>());
            set => _tags = value;
        }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Data backup retention days, value range: 7~30. 
        /// This field is valid and required when updating the backup plan of instance.
        /// </summary>
        [Input("backupRetentionPeriod")]
        public Input<int>? BackupRetentionPeriod { get; set; }

        /// <summary>
        /// The time window for starting the backup task is one hour interval. 
        /// This field is valid and required when updating the backup plan of instance.
        /// </summary>
        [Input("backupTime")]
        public Input<string>? BackupTime { get; set; }

        /// <summary>
        /// The charge info.
        /// </summary>
        [Input("chargeInfo")]
        public Input<Inputs.InstanceChargeInfoGetArgs>? ChargeInfo { get; set; }

        /// <summary>
        /// The Compatible version. Valid values: `SQLServer_2019_Std`, `SQLServer_2019_Web`, `SQLServer_2019_Ent`.
        /// </summary>
        [Input("dbEngineVersion")]
        public Input<string>? DbEngineVersion { get; set; }

        [Input("fullBackupPeriods")]
        private InputList<string>? _fullBackupPeriods;

        /// <summary>
        /// Full backup cycle. Multiple values separated by commas. The values are as follows: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. 
        /// This field is valid and required when updating the backup plan of instance.
        /// </summary>
        public InputList<string> FullBackupPeriods
        {
            get => _fullBackupPeriods ?? (_fullBackupPeriods = new InputList<string>());
            set => _fullBackupPeriods = value;
        }

        /// <summary>
        /// Name of the instance.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The Instance type. When the value of the `db_engine_version` is `SQLServer_2019_Std`, the value of this field can be `HA` or `Basic`.When the value of the `db_engine_version` is `SQLServer_2019_Ent`, the value of this field can be `Cluster` or `Basic`.When the value of the `db_engine_version` is `SQLServer_2019_Web`, the value of this field can be `Basic`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The node specification.
        /// </summary>
        [Input("nodeSpec")]
        public Input<string>? NodeSpec { get; set; }

        /// <summary>
        /// The project name.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// Storage space size, measured in GiB. The range of values is 20GiB to 4000GiB, with a step size of 10GiB.
        /// </summary>
        [Input("storageSpace")]
        public Input<int>? StorageSpace { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The subnet id of the instance node. When creating an instance that includes primary and backup nodes and needs to deploy primary and backup nodes across availability zones, you can specify two subnet_id. By default, the first is the primary node availability zone, and the second is the backup node availability zone.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("superAccountPassword")]
        private Input<string>? _superAccountPassword;

        /// <summary>
        /// The super account password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        public Input<string>? SuperAccountPassword
        {
            get => _superAccountPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _superAccountPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("tags")]
        private InputList<Inputs.InstanceTagGetArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.InstanceTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.InstanceTagGetArgs>());
            set => _tags = value;
        }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}