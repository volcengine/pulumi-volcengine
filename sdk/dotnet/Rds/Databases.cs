// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds
{
    public static class Databases
    {
        /// <summary>
        /// (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds databases
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = Volcengine.Rds.Databases.Invoke(new()
        ///     {
        ///         InstanceId = "mysql-0fdd3bab2e7c",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<DatabasesResult> InvokeAsync(DatabasesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<DatabasesResult>("volcengine:rds/databases:Databases", args ?? new DatabasesArgs(), options.WithDefaults());

        /// <summary>
        /// (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds databases
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = Volcengine.Rds.Databases.Invoke(new()
        ///     {
        ///         InstanceId = "mysql-0fdd3bab2e7c",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<DatabasesResult> Invoke(DatabasesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<DatabasesResult>("volcengine:rds/databases:Databases", args ?? new DatabasesInvokeArgs(), options.WithDefaults());
    }


    public sealed class DatabasesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The status of the RDS database.
        /// </summary>
        [Input("dbStatus")]
        public string? DbStatus { get; set; }

        /// <summary>
        /// The id of the RDS instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// A Name Regex of RDS database.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public DatabasesArgs()
        {
        }
        public static new DatabasesArgs Empty => new DatabasesArgs();
    }

    public sealed class DatabasesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The status of the RDS database.
        /// </summary>
        [Input("dbStatus")]
        public Input<string>? DbStatus { get; set; }

        /// <summary>
        /// The id of the RDS instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// A Name Regex of RDS database.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public DatabasesInvokeArgs()
        {
        }
        public static new DatabasesInvokeArgs Empty => new DatabasesInvokeArgs();
    }


    [OutputType]
    public sealed class DatabasesResult
    {
        /// <summary>
        /// The status of the RDS database.
        /// </summary>
        public readonly string? DbStatus;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The collection of RDS instance account query.
        /// </summary>
        public readonly ImmutableArray<Outputs.DatabasesRdsDatabaseResult> RdsDatabases;
        /// <summary>
        /// The total count of RDS database query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private DatabasesResult(
            string? dbStatus,

            string id,

            string instanceId,

            string? nameRegex,

            string? outputFile,

            ImmutableArray<Outputs.DatabasesRdsDatabaseResult> rdsDatabases,

            int totalCount)
        {
            DbStatus = dbStatus;
            Id = id;
            InstanceId = instanceId;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            RdsDatabases = rdsDatabases;
            TotalCount = totalCount;
        }
    }
}