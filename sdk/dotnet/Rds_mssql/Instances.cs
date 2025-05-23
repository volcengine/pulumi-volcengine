// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_mssql
{
    [Obsolete(@"volcengine.rds_mssql.Instances has been deprecated in favor of volcengine.rds_mssql.getInstances")]
    public static class Instances
    {
        /// <summary>
        /// Use this data source to query detailed information of rds mssql instances
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
        ///     var foo = Volcengine.Rds_mssql.GetInstances.Invoke(new()
        ///     {
        ///         InstanceId = "mssql-d2fc5abe****",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<InstancesResult> InvokeAsync(InstancesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<InstancesResult>("volcengine:rds_mssql/instances:Instances", args ?? new InstancesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of rds mssql instances
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
        ///     var foo = Volcengine.Rds_mssql.GetInstances.Invoke(new()
        ///     {
        ///         InstanceId = "mssql-d2fc5abe****",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<InstancesResult> Invoke(InstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<InstancesResult>("volcengine:rds_mssql/instances:Instances", args ?? new InstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class InstancesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The charge type. Valid values: `PostPaid`, `PrePaid`.
        /// </summary>
        [Input("chargeType")]
        public string? ChargeType { get; set; }

        /// <summary>
        /// The end time of creating the instance, using UTC time format.
        /// </summary>
        [Input("createTimeEnd")]
        public string? CreateTimeEnd { get; set; }

        /// <summary>
        /// The start time of creating the instance, using UTC time format.
        /// </summary>
        [Input("createTimeStart")]
        public string? CreateTimeStart { get; set; }

        /// <summary>
        /// Compatible version. Valid values: `SQLServer_2019_Std`, `SQLServer_2019_Web`, `SQLServer_2019_Ent`.
        /// </summary>
        [Input("dbEngineVersion")]
        public string? DbEngineVersion { get; set; }

        /// <summary>
        /// Id of the instance.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// Name of the instance.
        /// </summary>
        [Input("instanceName")]
        public string? InstanceName { get; set; }

        /// <summary>
        /// Status of the instance.
        /// </summary>
        [Input("instanceStatus")]
        public string? InstanceStatus { get; set; }

        /// <summary>
        /// Instance type. Valid values: `HA`, `Basic`, `Cluster`.
        /// </summary>
        [Input("instanceType")]
        public string? InstanceType { get; set; }

        /// <summary>
        /// A Name Regex of RDS mssql instance.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("tags")]
        private List<Inputs.InstancesTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public List<Inputs.InstancesTagArgs> Tags
        {
            get => _tags ?? (_tags = new List<Inputs.InstancesTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The id of the zone.
        /// </summary>
        [Input("zoneId")]
        public string? ZoneId { get; set; }

        public InstancesArgs()
        {
        }
        public static new InstancesArgs Empty => new InstancesArgs();
    }

    public sealed class InstancesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The charge type. Valid values: `PostPaid`, `PrePaid`.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// The end time of creating the instance, using UTC time format.
        /// </summary>
        [Input("createTimeEnd")]
        public Input<string>? CreateTimeEnd { get; set; }

        /// <summary>
        /// The start time of creating the instance, using UTC time format.
        /// </summary>
        [Input("createTimeStart")]
        public Input<string>? CreateTimeStart { get; set; }

        /// <summary>
        /// Compatible version. Valid values: `SQLServer_2019_Std`, `SQLServer_2019_Web`, `SQLServer_2019_Ent`.
        /// </summary>
        [Input("dbEngineVersion")]
        public Input<string>? DbEngineVersion { get; set; }

        /// <summary>
        /// Id of the instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Name of the instance.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// Status of the instance.
        /// </summary>
        [Input("instanceStatus")]
        public Input<string>? InstanceStatus { get; set; }

        /// <summary>
        /// Instance type. Valid values: `HA`, `Basic`, `Cluster`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// A Name Regex of RDS mssql instance.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("tags")]
        private InputList<Inputs.InstancesTagInputArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.InstancesTagInputArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.InstancesTagInputArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The id of the zone.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public InstancesInvokeArgs()
        {
        }
        public static new InstancesInvokeArgs Empty => new InstancesInvokeArgs();
    }


    [OutputType]
    public sealed class InstancesResult
    {
        /// <summary>
        /// The charge type.
        /// </summary>
        public readonly string? ChargeType;
        public readonly string? CreateTimeEnd;
        public readonly string? CreateTimeStart;
        /// <summary>
        /// The db engine version.
        /// </summary>
        public readonly string? DbEngineVersion;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Instance ID.
        /// </summary>
        public readonly string? InstanceId;
        /// <summary>
        /// The name of the instance.
        /// </summary>
        public readonly string? InstanceName;
        /// <summary>
        /// The status of the instance.
        /// </summary>
        public readonly string? InstanceStatus;
        /// <summary>
        /// The type of the instance.
        /// </summary>
        public readonly string? InstanceType;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesInstanceResult> Instances;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesTagResult> Tags;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The zone id.
        /// </summary>
        public readonly string? ZoneId;

        [OutputConstructor]
        private InstancesResult(
            string? chargeType,

            string? createTimeEnd,

            string? createTimeStart,

            string? dbEngineVersion,

            string id,

            string? instanceId,

            string? instanceName,

            string? instanceStatus,

            string? instanceType,

            ImmutableArray<Outputs.InstancesInstanceResult> instances,

            string? nameRegex,

            string? outputFile,

            ImmutableArray<Outputs.InstancesTagResult> tags,

            int totalCount,

            string? zoneId)
        {
            ChargeType = chargeType;
            CreateTimeEnd = createTimeEnd;
            CreateTimeStart = createTimeStart;
            DbEngineVersion = dbEngineVersion;
            Id = id;
            InstanceId = instanceId;
            InstanceName = instanceName;
            InstanceStatus = instanceStatus;
            InstanceType = instanceType;
            Instances = instances;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            Tags = tags;
            TotalCount = totalCount;
            ZoneId = zoneId;
        }
    }
}
