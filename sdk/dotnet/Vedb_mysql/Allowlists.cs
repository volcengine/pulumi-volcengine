// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vedb_mysql
{
    public static class Allowlists
    {
        /// <summary>
        /// Use this data source to query detailed information of vedb mysql allowlists
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
        ///     var fooAllowlist = new Volcengine.Vedb_mysql.Allowlist("fooAllowlist", new()
        ///     {
        ///         AllowListName = "acc-test-allowlist",
        ///         AllowListDesc = "acc-test",
        ///         AllowListType = "IPv4",
        ///         AllowLists = new[]
        ///         {
        ///             "192.168.0.0/24",
        ///             "192.168.1.0/24",
        ///             "192.168.2.0/24",
        ///         },
        ///     });
        /// 
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
        ///         ZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[2]?.Id),
        ///         VpcId = fooVpc.Id,
        ///     });
        /// 
        ///     var fooInstance = new Volcengine.Vedb_mysql.Instance("fooInstance", new()
        ///     {
        ///         ChargeType = "PostPaid",
        ///         StorageChargeType = "PostPaid",
        ///         DbEngineVersion = "MySQL_8_0",
        ///         DbMinorVersion = "3.0",
        ///         NodeNumber = 2,
        ///         NodeSpec = "vedb.mysql.x4.large",
        ///         SubnetId = fooSubnet.Id,
        ///         InstanceName = "tf-test",
        ///         ProjectName = "testA",
        ///         Tags = new[]
        ///         {
        ///             new Volcengine.Vedb_mysql.Inputs.InstanceTagArgs
        ///             {
        ///                 Key = "tftest",
        ///                 Value = "tftest",
        ///             },
        ///             new Volcengine.Vedb_mysql.Inputs.InstanceTagArgs
        ///             {
        ///                 Key = "tftest2",
        ///                 Value = "tftest2",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooAllowlistAssociate = new Volcengine.Vedb_mysql.AllowlistAssociate("fooAllowlistAssociate", new()
        ///     {
        ///         AllowListId = fooAllowlist.Id,
        ///         InstanceId = fooInstance.Id,
        ///     });
        /// 
        ///     var fooAllowlists = Volcengine.Vedb_mysql.Allowlists.Invoke(new()
        ///     {
        ///         InstanceId = fooInstance.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<AllowlistsResult> InvokeAsync(AllowlistsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<AllowlistsResult>("volcengine:vedb_mysql/allowlists:Allowlists", args ?? new AllowlistsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of vedb mysql allowlists
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
        ///     var fooAllowlist = new Volcengine.Vedb_mysql.Allowlist("fooAllowlist", new()
        ///     {
        ///         AllowListName = "acc-test-allowlist",
        ///         AllowListDesc = "acc-test",
        ///         AllowListType = "IPv4",
        ///         AllowLists = new[]
        ///         {
        ///             "192.168.0.0/24",
        ///             "192.168.1.0/24",
        ///             "192.168.2.0/24",
        ///         },
        ///     });
        /// 
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
        ///         ZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[2]?.Id),
        ///         VpcId = fooVpc.Id,
        ///     });
        /// 
        ///     var fooInstance = new Volcengine.Vedb_mysql.Instance("fooInstance", new()
        ///     {
        ///         ChargeType = "PostPaid",
        ///         StorageChargeType = "PostPaid",
        ///         DbEngineVersion = "MySQL_8_0",
        ///         DbMinorVersion = "3.0",
        ///         NodeNumber = 2,
        ///         NodeSpec = "vedb.mysql.x4.large",
        ///         SubnetId = fooSubnet.Id,
        ///         InstanceName = "tf-test",
        ///         ProjectName = "testA",
        ///         Tags = new[]
        ///         {
        ///             new Volcengine.Vedb_mysql.Inputs.InstanceTagArgs
        ///             {
        ///                 Key = "tftest",
        ///                 Value = "tftest",
        ///             },
        ///             new Volcengine.Vedb_mysql.Inputs.InstanceTagArgs
        ///             {
        ///                 Key = "tftest2",
        ///                 Value = "tftest2",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooAllowlistAssociate = new Volcengine.Vedb_mysql.AllowlistAssociate("fooAllowlistAssociate", new()
        ///     {
        ///         AllowListId = fooAllowlist.Id,
        ///         InstanceId = fooInstance.Id,
        ///     });
        /// 
        ///     var fooAllowlists = Volcengine.Vedb_mysql.Allowlists.Invoke(new()
        ///     {
        ///         InstanceId = fooInstance.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<AllowlistsResult> Invoke(AllowlistsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<AllowlistsResult>("volcengine:vedb_mysql/allowlists:Allowlists", args ?? new AllowlistsInvokeArgs(), options.WithDefaults());
    }


    public sealed class AllowlistsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Instance ID. When an InstanceId is specified, the DescribeAllowLists interface will return the whitelist bound to the specified instance.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// A Name Regex of Resource.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The region of the allow lists.
        /// </summary>
        [Input("regionId", required: true)]
        public string RegionId { get; set; } = null!;

        public AllowlistsArgs()
        {
        }
        public static new AllowlistsArgs Empty => new AllowlistsArgs();
    }

    public sealed class AllowlistsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Instance ID. When an InstanceId is specified, the DescribeAllowLists interface will return the whitelist bound to the specified instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// A Name Regex of Resource.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The region of the allow lists.
        /// </summary>
        [Input("regionId", required: true)]
        public Input<string> RegionId { get; set; } = null!;

        public AllowlistsInvokeArgs()
        {
        }
        public static new AllowlistsInvokeArgs Empty => new AllowlistsInvokeArgs();
    }


    [OutputType]
    public sealed class AllowlistsResult
    {
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.AllowlistsAllowListResult> AllowLists;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The id of the instance.
        /// </summary>
        public readonly string? InstanceId;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        public readonly string RegionId;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private AllowlistsResult(
            ImmutableArray<Outputs.AllowlistsAllowListResult> allowLists,

            string id,

            string? instanceId,

            string? nameRegex,

            string? outputFile,

            string regionId,

            int totalCount)
        {
            AllowLists = allowLists;
            Id = id;
            InstanceId = instanceId;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            RegionId = regionId;
            TotalCount = totalCount;
        }
    }
}