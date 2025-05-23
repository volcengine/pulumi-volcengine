// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Kafka
{
    [Obsolete(@"volcengine.kafka.SaslUsers has been deprecated in favor of volcengine.kafka.getSaslUsers")]
    public static class SaslUsers
    {
        /// <summary>
        /// Use this data source to query detailed information of kafka sasl users
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
        ///     var fooInstance = new Volcengine.Kafka.Instance("fooInstance", new()
        ///     {
        ///         InstanceName = "acc-test-kafka",
        ///         InstanceDescription = "tf-test",
        ///         Version = "2.2.2",
        ///         ComputeSpec = "kafka.20xrate.hw",
        ///         SubnetId = fooSubnet.Id,
        ///         UserName = "tf-user",
        ///         UserPassword = "tf-pass!@q1",
        ///         ChargeType = "PostPaid",
        ///         StorageSpace = 300,
        ///         PartitionNumber = 350,
        ///         ProjectName = "default",
        ///         Tags = new[]
        ///         {
        ///             new Volcengine.Kafka.Inputs.InstanceTagArgs
        ///             {
        ///                 Key = "k1",
        ///                 Value = "v1",
        ///             },
        ///         },
        ///         Parameters = new[]
        ///         {
        ///             new Volcengine.Kafka.Inputs.InstanceParameterArgs
        ///             {
        ///                 ParameterName = "MessageMaxByte",
        ///                 ParameterValue = "12",
        ///             },
        ///             new Volcengine.Kafka.Inputs.InstanceParameterArgs
        ///             {
        ///                 ParameterName = "LogRetentionHours",
        ///                 ParameterValue = "70",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooSaslUser = new Volcengine.Kafka.SaslUser("fooSaslUser", new()
        ///     {
        ///         UserName = "acc-test-user",
        ///         InstanceId = fooInstance.Id,
        ///         UserPassword = "suqsnis123!",
        ///         Description = "tf-test",
        ///         AllAuthority = true,
        ///         PasswordType = "Scram",
        ///     });
        /// 
        ///     var @default = Volcengine.Kafka.GetSaslUsers.Invoke(new()
        ///     {
        ///         InstanceId = fooInstance.Id,
        ///         UserName = fooSaslUser.UserName,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<SaslUsersResult> InvokeAsync(SaslUsersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<SaslUsersResult>("volcengine:kafka/saslUsers:SaslUsers", args ?? new SaslUsersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of kafka sasl users
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
        ///     var fooInstance = new Volcengine.Kafka.Instance("fooInstance", new()
        ///     {
        ///         InstanceName = "acc-test-kafka",
        ///         InstanceDescription = "tf-test",
        ///         Version = "2.2.2",
        ///         ComputeSpec = "kafka.20xrate.hw",
        ///         SubnetId = fooSubnet.Id,
        ///         UserName = "tf-user",
        ///         UserPassword = "tf-pass!@q1",
        ///         ChargeType = "PostPaid",
        ///         StorageSpace = 300,
        ///         PartitionNumber = 350,
        ///         ProjectName = "default",
        ///         Tags = new[]
        ///         {
        ///             new Volcengine.Kafka.Inputs.InstanceTagArgs
        ///             {
        ///                 Key = "k1",
        ///                 Value = "v1",
        ///             },
        ///         },
        ///         Parameters = new[]
        ///         {
        ///             new Volcengine.Kafka.Inputs.InstanceParameterArgs
        ///             {
        ///                 ParameterName = "MessageMaxByte",
        ///                 ParameterValue = "12",
        ///             },
        ///             new Volcengine.Kafka.Inputs.InstanceParameterArgs
        ///             {
        ///                 ParameterName = "LogRetentionHours",
        ///                 ParameterValue = "70",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooSaslUser = new Volcengine.Kafka.SaslUser("fooSaslUser", new()
        ///     {
        ///         UserName = "acc-test-user",
        ///         InstanceId = fooInstance.Id,
        ///         UserPassword = "suqsnis123!",
        ///         Description = "tf-test",
        ///         AllAuthority = true,
        ///         PasswordType = "Scram",
        ///     });
        /// 
        ///     var @default = Volcengine.Kafka.GetSaslUsers.Invoke(new()
        ///     {
        ///         InstanceId = fooInstance.Id,
        ///         UserName = fooSaslUser.UserName,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<SaslUsersResult> Invoke(SaslUsersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<SaslUsersResult>("volcengine:kafka/saslUsers:SaslUsers", args ?? new SaslUsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class SaslUsersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The user name, support fuzzy matching.
        /// </summary>
        [Input("userName")]
        public string? UserName { get; set; }

        public SaslUsersArgs()
        {
        }
        public static new SaslUsersArgs Empty => new SaslUsersArgs();
    }

    public sealed class SaslUsersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The user name, support fuzzy matching.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public SaslUsersInvokeArgs()
        {
        }
        public static new SaslUsersInvokeArgs Empty => new SaslUsersInvokeArgs();
    }


    [OutputType]
    public sealed class SaslUsersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The name of user.
        /// </summary>
        public readonly string? UserName;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.SaslUsersUserResult> Users;

        [OutputConstructor]
        private SaslUsersResult(
            string id,

            string instanceId,

            string? outputFile,

            int totalCount,

            string? userName,

            ImmutableArray<Outputs.SaslUsersUserResult> users)
        {
            Id = id;
            InstanceId = instanceId;
            OutputFile = outputFile;
            TotalCount = totalCount;
            UserName = userName;
            Users = users;
        }
    }
}
