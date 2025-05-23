// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_mysql
{
    public static class GetAccounts
    {
        /// <summary>
        /// Use this data source to query detailed information of rds mysql accounts
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
        ///     var fooInstance = new Volcengine.Rds_mysql.Instance("fooInstance", new()
        ///     {
        ///         InstanceName = "acc-test-rds-mysql",
        ///         DbEngineVersion = "MySQL_5_7",
        ///         NodeSpec = "rds.mysql.1c2g",
        ///         PrimaryZoneId = fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///         SecondaryZoneId = fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///         StorageSpace = 80,
        ///         SubnetId = fooSubnet.Id,
        ///         LowerCaseTableNames = "1",
        ///         ChargeInfo = new Volcengine.Rds_mysql.Inputs.InstanceChargeInfoArgs
        ///         {
        ///             ChargeType = "PostPaid",
        ///         },
        ///         Parameters = new[]
        ///         {
        ///             new Volcengine.Rds_mysql.Inputs.InstanceParameterArgs
        ///             {
        ///                 ParameterName = "auto_increment_increment",
        ///                 ParameterValue = "2",
        ///             },
        ///             new Volcengine.Rds_mysql.Inputs.InstanceParameterArgs
        ///             {
        ///                 ParameterName = "auto_increment_offset",
        ///                 ParameterValue = "4",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooDatabase = new Volcengine.Rds_mysql.Database("fooDatabase", new()
        ///     {
        ///         DbName = "acc-test-db",
        ///         InstanceId = fooInstance.Id,
        ///     });
        /// 
        ///     var fooAccount = new Volcengine.Rds_mysql.Account("fooAccount", new()
        ///     {
        ///         AccountName = "acc-test-account",
        ///         AccountPassword = "93f0cb0614Aab12",
        ///         AccountType = "Normal",
        ///         InstanceId = fooInstance.Id,
        ///         AccountPrivileges = new[]
        ///         {
        ///             new Volcengine.Rds_mysql.Inputs.AccountAccountPrivilegeArgs
        ///             {
        ///                 DbName = fooDatabase.DbName,
        ///                 AccountPrivilege = "Custom",
        ///                 AccountPrivilegeDetail = "SELECT,INSERT",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooAccounts = Volcengine.Rds_mysql.GetAccounts.Invoke(new()
        ///     {
        ///         InstanceId = fooInstance.Id,
        ///         AccountName = fooAccount.AccountName,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAccountsResult> InvokeAsync(GetAccountsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccountsResult>("volcengine:rds_mysql/getAccounts:getAccounts", args ?? new GetAccountsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of rds mysql accounts
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
        ///     var fooInstance = new Volcengine.Rds_mysql.Instance("fooInstance", new()
        ///     {
        ///         InstanceName = "acc-test-rds-mysql",
        ///         DbEngineVersion = "MySQL_5_7",
        ///         NodeSpec = "rds.mysql.1c2g",
        ///         PrimaryZoneId = fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///         SecondaryZoneId = fooZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
        ///         StorageSpace = 80,
        ///         SubnetId = fooSubnet.Id,
        ///         LowerCaseTableNames = "1",
        ///         ChargeInfo = new Volcengine.Rds_mysql.Inputs.InstanceChargeInfoArgs
        ///         {
        ///             ChargeType = "PostPaid",
        ///         },
        ///         Parameters = new[]
        ///         {
        ///             new Volcengine.Rds_mysql.Inputs.InstanceParameterArgs
        ///             {
        ///                 ParameterName = "auto_increment_increment",
        ///                 ParameterValue = "2",
        ///             },
        ///             new Volcengine.Rds_mysql.Inputs.InstanceParameterArgs
        ///             {
        ///                 ParameterName = "auto_increment_offset",
        ///                 ParameterValue = "4",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooDatabase = new Volcengine.Rds_mysql.Database("fooDatabase", new()
        ///     {
        ///         DbName = "acc-test-db",
        ///         InstanceId = fooInstance.Id,
        ///     });
        /// 
        ///     var fooAccount = new Volcengine.Rds_mysql.Account("fooAccount", new()
        ///     {
        ///         AccountName = "acc-test-account",
        ///         AccountPassword = "93f0cb0614Aab12",
        ///         AccountType = "Normal",
        ///         InstanceId = fooInstance.Id,
        ///         AccountPrivileges = new[]
        ///         {
        ///             new Volcengine.Rds_mysql.Inputs.AccountAccountPrivilegeArgs
        ///             {
        ///                 DbName = fooDatabase.DbName,
        ///                 AccountPrivilege = "Custom",
        ///                 AccountPrivilegeDetail = "SELECT,INSERT",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooAccounts = Volcengine.Rds_mysql.GetAccounts.Invoke(new()
        ///     {
        ///         InstanceId = fooInstance.Id,
        ///         AccountName = fooAccount.AccountName,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAccountsResult> Invoke(GetAccountsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountsResult>("volcengine:rds_mysql/getAccounts:getAccounts", args ?? new GetAccountsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccountsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the database account. This field supports fuzzy query.
        /// </summary>
        [Input("accountName")]
        public string? AccountName { get; set; }

        /// <summary>
        /// The id of the RDS instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// A Name Regex of database account.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetAccountsArgs()
        {
        }
        public static new GetAccountsArgs Empty => new GetAccountsArgs();
    }

    public sealed class GetAccountsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the database account. This field supports fuzzy query.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// The id of the RDS instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// A Name Regex of database account.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetAccountsInvokeArgs()
        {
        }
        public static new GetAccountsInvokeArgs Empty => new GetAccountsInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccountsResult
    {
        /// <summary>
        /// The name of the database account.
        /// </summary>
        public readonly string? AccountName;
        /// <summary>
        /// The collection of RDS instance account query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccountsAccountResult> Accounts;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of database account query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetAccountsResult(
            string? accountName,

            ImmutableArray<Outputs.GetAccountsAccountResult> accounts,

            string id,

            string instanceId,

            string? nameRegex,

            string? outputFile,

            int totalCount)
        {
            AccountName = accountName;
            Accounts = accounts;
            Id = id;
            InstanceId = instanceId;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
