// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_postgresql
{
    public static class Schemas
    {
        /// <summary>
        /// Use this data source to query detailed information of rds postgresql schemas
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
        ///     var fooZones = Volcengine.Ecs.Zones.Invoke();
        /// 
        ///     var fooVpc = new Volcengine.Vpc.Vpc("fooVpc", new()
        ///     {
        ///         VpcName = "acc-test-project1",
        ///         CidrBlock = "172.16.0.0/16",
        ///     });
        /// 
        ///     var fooSubnet = new Volcengine.Vpc.Subnet("fooSubnet", new()
        ///     {
        ///         SubnetName = "acc-subnet-test-2",
        ///         CidrBlock = "172.16.0.0/24",
        ///         ZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[0]?.Id),
        ///         VpcId = fooVpc.Id,
        ///     });
        /// 
        ///     var fooInstance = new Volcengine.Rds_postgresql.Instance("fooInstance", new()
        ///     {
        ///         DbEngineVersion = "PostgreSQL_12",
        ///         NodeSpec = "rds.postgres.1c2g",
        ///         PrimaryZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[0]?.Id),
        ///         SecondaryZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[0]?.Id),
        ///         StorageSpace = 40,
        ///         SubnetId = fooSubnet.Id,
        ///         InstanceName = "acc-test-1",
        ///         ChargeInfo = new Volcengine.Rds_postgresql.Inputs.InstanceChargeInfoArgs
        ///         {
        ///             ChargeType = "PostPaid",
        ///         },
        ///         ProjectName = "default",
        ///         Tags = new[]
        ///         {
        ///             new Volcengine.Rds_postgresql.Inputs.InstanceTagArgs
        ///             {
        ///                 Key = "tfk1",
        ///                 Value = "tfv1",
        ///             },
        ///         },
        ///         Parameters = new[]
        ///         {
        ///             new Volcengine.Rds_postgresql.Inputs.InstanceParameterArgs
        ///             {
        ///                 Name = "auto_explain.log_analyze",
        ///                 Value = "off",
        ///             },
        ///             new Volcengine.Rds_postgresql.Inputs.InstanceParameterArgs
        ///             {
        ///                 Name = "auto_explain.log_format",
        ///                 Value = "text",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooDatabase = new Volcengine.Rds_postgresql.Database("fooDatabase", new()
        ///     {
        ///         DbName = "acc-test",
        ///         InstanceId = fooInstance.Id,
        ///         CType = "C",
        ///         Collate = "zh_CN.utf8",
        ///     });
        /// 
        ///     var fooAccount = new Volcengine.Rds_postgresql.Account("fooAccount", new()
        ///     {
        ///         AccountName = "acc-test-account",
        ///         AccountPassword = "9wc@********12",
        ///         AccountType = "Normal",
        ///         InstanceId = fooInstance.Id,
        ///         AccountPrivileges = "Inherit,Login,CreateRole,CreateDB",
        ///     });
        /// 
        ///     var foo1 = new Volcengine.Rds_postgresql.Account("foo1", new()
        ///     {
        ///         AccountName = "acc-test-account1",
        ///         AccountPassword = "9wc@*******12",
        ///         AccountType = "Normal",
        ///         InstanceId = fooInstance.Id,
        ///         AccountPrivileges = "Inherit,Login,CreateRole,CreateDB",
        ///     });
        /// 
        ///     var fooSchema = new Volcengine.Rds_postgresql.Schema("fooSchema", new()
        ///     {
        ///         DbName = fooDatabase.DbName,
        ///         InstanceId = fooInstance.Id,
        ///         Owner = fooAccount.AccountName,
        ///         SchemaName = "acc-test-schema",
        ///     });
        /// 
        ///     var fooSchemas = Volcengine.Rds_postgresql.Schemas.Invoke(new()
        ///     {
        ///         DbName = fooSchema.DbName,
        ///         InstanceId = fooInstance.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<SchemasResult> InvokeAsync(SchemasArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<SchemasResult>("volcengine:rds_postgresql/schemas:Schemas", args ?? new SchemasArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of rds postgresql schemas
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
        ///     var fooZones = Volcengine.Ecs.Zones.Invoke();
        /// 
        ///     var fooVpc = new Volcengine.Vpc.Vpc("fooVpc", new()
        ///     {
        ///         VpcName = "acc-test-project1",
        ///         CidrBlock = "172.16.0.0/16",
        ///     });
        /// 
        ///     var fooSubnet = new Volcengine.Vpc.Subnet("fooSubnet", new()
        ///     {
        ///         SubnetName = "acc-subnet-test-2",
        ///         CidrBlock = "172.16.0.0/24",
        ///         ZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[0]?.Id),
        ///         VpcId = fooVpc.Id,
        ///     });
        /// 
        ///     var fooInstance = new Volcengine.Rds_postgresql.Instance("fooInstance", new()
        ///     {
        ///         DbEngineVersion = "PostgreSQL_12",
        ///         NodeSpec = "rds.postgres.1c2g",
        ///         PrimaryZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[0]?.Id),
        ///         SecondaryZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[0]?.Id),
        ///         StorageSpace = 40,
        ///         SubnetId = fooSubnet.Id,
        ///         InstanceName = "acc-test-1",
        ///         ChargeInfo = new Volcengine.Rds_postgresql.Inputs.InstanceChargeInfoArgs
        ///         {
        ///             ChargeType = "PostPaid",
        ///         },
        ///         ProjectName = "default",
        ///         Tags = new[]
        ///         {
        ///             new Volcengine.Rds_postgresql.Inputs.InstanceTagArgs
        ///             {
        ///                 Key = "tfk1",
        ///                 Value = "tfv1",
        ///             },
        ///         },
        ///         Parameters = new[]
        ///         {
        ///             new Volcengine.Rds_postgresql.Inputs.InstanceParameterArgs
        ///             {
        ///                 Name = "auto_explain.log_analyze",
        ///                 Value = "off",
        ///             },
        ///             new Volcengine.Rds_postgresql.Inputs.InstanceParameterArgs
        ///             {
        ///                 Name = "auto_explain.log_format",
        ///                 Value = "text",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooDatabase = new Volcengine.Rds_postgresql.Database("fooDatabase", new()
        ///     {
        ///         DbName = "acc-test",
        ///         InstanceId = fooInstance.Id,
        ///         CType = "C",
        ///         Collate = "zh_CN.utf8",
        ///     });
        /// 
        ///     var fooAccount = new Volcengine.Rds_postgresql.Account("fooAccount", new()
        ///     {
        ///         AccountName = "acc-test-account",
        ///         AccountPassword = "9wc@********12",
        ///         AccountType = "Normal",
        ///         InstanceId = fooInstance.Id,
        ///         AccountPrivileges = "Inherit,Login,CreateRole,CreateDB",
        ///     });
        /// 
        ///     var foo1 = new Volcengine.Rds_postgresql.Account("foo1", new()
        ///     {
        ///         AccountName = "acc-test-account1",
        ///         AccountPassword = "9wc@*******12",
        ///         AccountType = "Normal",
        ///         InstanceId = fooInstance.Id,
        ///         AccountPrivileges = "Inherit,Login,CreateRole,CreateDB",
        ///     });
        /// 
        ///     var fooSchema = new Volcengine.Rds_postgresql.Schema("fooSchema", new()
        ///     {
        ///         DbName = fooDatabase.DbName,
        ///         InstanceId = fooInstance.Id,
        ///         Owner = fooAccount.AccountName,
        ///         SchemaName = "acc-test-schema",
        ///     });
        /// 
        ///     var fooSchemas = Volcengine.Rds_postgresql.Schemas.Invoke(new()
        ///     {
        ///         DbName = fooSchema.DbName,
        ///         InstanceId = fooInstance.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<SchemasResult> Invoke(SchemasInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<SchemasResult>("volcengine:rds_postgresql/schemas:Schemas", args ?? new SchemasInvokeArgs(), options.WithDefaults());
    }


    public sealed class SchemasArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("dbName")]
        public string? DbName { get; set; }

        /// <summary>
        /// The id of the instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public SchemasArgs()
        {
        }
        public static new SchemasArgs Empty => new SchemasArgs();
    }

    public sealed class SchemasInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("dbName")]
        public Input<string>? DbName { get; set; }

        /// <summary>
        /// The id of the instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public SchemasInvokeArgs()
        {
        }
        public static new SchemasInvokeArgs Empty => new SchemasInvokeArgs();
    }


    [OutputType]
    public sealed class SchemasResult
    {
        /// <summary>
        /// The name of the database.
        /// </summary>
        public readonly string? DbName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? OutputFile;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.SchemasSchemaResult> Schemas;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private SchemasResult(
            string? dbName,

            string id,

            string instanceId,

            string? outputFile,

            ImmutableArray<Outputs.SchemasSchemaResult> schemas,

            int totalCount)
        {
            DbName = dbName;
            Id = id;
            InstanceId = instanceId;
            OutputFile = outputFile;
            Schemas = schemas;
            TotalCount = totalCount;
        }
    }
}