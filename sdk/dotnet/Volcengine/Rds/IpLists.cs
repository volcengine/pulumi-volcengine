// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Rds
{
    public static class IpLists
    {
        /// <summary>
        /// (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds ip lists
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var @default = Output.Create(Volcengine.Rds.IpLists.InvokeAsync(new Volcengine.Rds.IpListsArgs
        ///         {
        ///             InstanceId = "mysql-0fdd3bab2e7c",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<IpListsResult> InvokeAsync(IpListsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<IpListsResult>("volcengine:rds/ipLists:IpLists", args ?? new IpListsArgs(), options.WithDefaults());

        /// <summary>
        /// (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds ip lists
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var @default = Output.Create(Volcengine.Rds.IpLists.InvokeAsync(new Volcengine.Rds.IpListsArgs
        ///         {
        ///             InstanceId = "mysql-0fdd3bab2e7c",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<IpListsResult> Invoke(IpListsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<IpListsResult>("volcengine:rds/ipLists:IpLists", args ?? new IpListsInvokeArgs(), options.WithDefaults());
    }


    public sealed class IpListsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the RDS instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// A Name Regex of RDS ip list.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public IpListsArgs()
        {
        }
    }

    public sealed class IpListsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the RDS instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// A Name Regex of RDS ip list.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public IpListsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class IpListsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The collection of RDS ip list account query.
        /// </summary>
        public readonly ImmutableArray<Outputs.IpListsRdsIpListResult> RdsIpLists;
        /// <summary>
        /// The total count of RDS ip list query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private IpListsResult(
            string id,

            string instanceId,

            string? nameRegex,

            string? outputFile,

            ImmutableArray<Outputs.IpListsRdsIpListResult> rdsIpLists,

            int totalCount)
        {
            Id = id;
            InstanceId = instanceId;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            RdsIpLists = rdsIpLists;
            TotalCount = totalCount;
        }
    }
}