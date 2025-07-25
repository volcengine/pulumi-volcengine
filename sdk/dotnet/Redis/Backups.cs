// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Redis
{
    [Obsolete(@"volcengine.redis.Backups has been deprecated in favor of volcengine.redis.getBackups")]
    public static class Backups
    {
        /// <summary>
        /// Use this data source to query detailed information of redis backups
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
        ///     var fooZones = Volcengine.Ecs.GetZones.Invoke();
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
        ///         ZoneId = fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///         VpcId = fooVpc.Id,
        ///     });
        /// 
        ///     var fooInstance = new Volcengine.Redis.Instance("fooInstance", new()
        ///     {
        ///         ZoneIds = new[]
        ///         {
        ///             fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///         },
        ///         InstanceName = "acc-test-tf-redis",
        ///         ShardedCluster = 1,
        ///         Password = "1qaz!QAZ12",
        ///         NodeNumber = 2,
        ///         ShardCapacity = 1024,
        ///         ShardNumber = 2,
        ///         EngineVersion = "5.0",
        ///         SubnetId = fooSubnet.Id,
        ///         DeletionProtection = "disabled",
        ///         VpcAuthMode = "close",
        ///         ChargeType = "PostPaid",
        ///         Port = 6381,
        ///         ProjectName = "default",
        ///     });
        /// 
        ///     var fooBackup = new List&lt;Volcengine.Redis.Backup&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; 3; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         fooBackup.Add(new Volcengine.Redis.Backup($"fooBackup-{range.Value}", new()
        ///         {
        ///             InstanceId = fooInstance.Id,
        ///         }));
        ///     }
        ///     var fooBackups = Volcengine.Redis.GetBackups.Invoke(new()
        ///     {
        ///         InstanceId = fooInstance.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<BackupsResult> InvokeAsync(BackupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<BackupsResult>("volcengine:redis/backups:Backups", args ?? new BackupsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of redis backups
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
        ///     var fooZones = Volcengine.Ecs.GetZones.Invoke();
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
        ///         ZoneId = fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///         VpcId = fooVpc.Id,
        ///     });
        /// 
        ///     var fooInstance = new Volcengine.Redis.Instance("fooInstance", new()
        ///     {
        ///         ZoneIds = new[]
        ///         {
        ///             fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///         },
        ///         InstanceName = "acc-test-tf-redis",
        ///         ShardedCluster = 1,
        ///         Password = "1qaz!QAZ12",
        ///         NodeNumber = 2,
        ///         ShardCapacity = 1024,
        ///         ShardNumber = 2,
        ///         EngineVersion = "5.0",
        ///         SubnetId = fooSubnet.Id,
        ///         DeletionProtection = "disabled",
        ///         VpcAuthMode = "close",
        ///         ChargeType = "PostPaid",
        ///         Port = 6381,
        ///         ProjectName = "default",
        ///     });
        /// 
        ///     var fooBackup = new List&lt;Volcengine.Redis.Backup&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; 3; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         fooBackup.Add(new Volcengine.Redis.Backup($"fooBackup-{range.Value}", new()
        ///         {
        ///             InstanceId = fooInstance.Id,
        ///         }));
        ///     }
        ///     var fooBackups = Volcengine.Redis.GetBackups.Invoke(new()
        ///     {
        ///         InstanceId = fooInstance.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<BackupsResult> Invoke(BackupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<BackupsResult>("volcengine:redis/backups:Backups", args ?? new BackupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class BackupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of backup point.
        /// </summary>
        [Input("backupPointId")]
        public string? BackupPointId { get; set; }

        /// <summary>
        /// Backup name, supporting fuzzy query.
        /// </summary>
        [Input("backupPointName")]
        public string? BackupPointName { get; set; }

        [Input("backupStrategyLists")]
        private List<string>? _backupStrategyLists;

        /// <summary>
        /// The list of backup strategy, support AutomatedBackup and ManualBackup.
        /// </summary>
        public List<string> BackupStrategyLists
        {
            get => _backupStrategyLists ?? (_backupStrategyLists = new List<string>());
            set => _backupStrategyLists = value;
        }

        /// <summary>
        /// Query end time.
        /// </summary>
        [Input("endTime")]
        public string? EndTime { get; set; }

        /// <summary>
        /// Id of instance.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Back up the project to which it belongs.
        /// </summary>
        [Input("projectName")]
        public string? ProjectName { get; set; }

        /// <summary>
        /// The query scope of the backup.
        /// </summary>
        [Input("scope")]
        public string? Scope { get; set; }

        /// <summary>
        /// Query start time.
        /// </summary>
        [Input("startTime")]
        public string? StartTime { get; set; }

        public BackupsArgs()
        {
        }
        public static new BackupsArgs Empty => new BackupsArgs();
    }

    public sealed class BackupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of backup point.
        /// </summary>
        [Input("backupPointId")]
        public Input<string>? BackupPointId { get; set; }

        /// <summary>
        /// Backup name, supporting fuzzy query.
        /// </summary>
        [Input("backupPointName")]
        public Input<string>? BackupPointName { get; set; }

        [Input("backupStrategyLists")]
        private InputList<string>? _backupStrategyLists;

        /// <summary>
        /// The list of backup strategy, support AutomatedBackup and ManualBackup.
        /// </summary>
        public InputList<string> BackupStrategyLists
        {
            get => _backupStrategyLists ?? (_backupStrategyLists = new InputList<string>());
            set => _backupStrategyLists = value;
        }

        /// <summary>
        /// Query end time.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// Id of instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Back up the project to which it belongs.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The query scope of the backup.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// Query start time.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public BackupsInvokeArgs()
        {
        }
        public static new BackupsInvokeArgs Empty => new BackupsInvokeArgs();
    }


    [OutputType]
    public sealed class BackupsResult
    {
        /// <summary>
        /// The id of backup point.
        /// </summary>
        public readonly string? BackupPointId;
        public readonly string? BackupPointName;
        public readonly ImmutableArray<string> BackupStrategyLists;
        /// <summary>
        /// Information of backups.
        /// </summary>
        public readonly ImmutableArray<Outputs.BackupsBackupResult> Backups;
        /// <summary>
        /// End time of backup.
        /// </summary>
        public readonly string? EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Id of instance.
        /// </summary>
        public readonly string? InstanceId;
        public readonly string? OutputFile;
        /// <summary>
        /// Project name of instance.
        /// </summary>
        public readonly string? ProjectName;
        public readonly string? Scope;
        /// <summary>
        /// Start time of backup.
        /// </summary>
        public readonly string? StartTime;
        /// <summary>
        /// The total count of backup query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private BackupsResult(
            string? backupPointId,

            string? backupPointName,

            ImmutableArray<string> backupStrategyLists,

            ImmutableArray<Outputs.BackupsBackupResult> backups,

            string? endTime,

            string id,

            string? instanceId,

            string? outputFile,

            string? projectName,

            string? scope,

            string? startTime,

            int totalCount)
        {
            BackupPointId = backupPointId;
            BackupPointName = backupPointName;
            BackupStrategyLists = backupStrategyLists;
            Backups = backups;
            EndTime = endTime;
            Id = id;
            InstanceId = instanceId;
            OutputFile = outputFile;
            ProjectName = projectName;
            Scope = scope;
            StartTime = startTime;
            TotalCount = totalCount;
        }
    }
}
