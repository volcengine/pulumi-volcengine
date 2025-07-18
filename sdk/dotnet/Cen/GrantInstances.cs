// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cen
{
    [Obsolete(@"volcengine.cen.GrantInstances has been deprecated in favor of volcengine.cen.getGrantInstances")]
    public static class GrantInstances
    {
        /// <summary>
        /// Use this data source to query detailed information of cen grant instances
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
        ///     var foo = Volcengine.Cen.GetGrantInstances.Invoke(new()
        ///     {
        ///         InstanceId = "vpc-2bysvq1xx543k2dx0eeul****",
        ///         InstanceRegionId = "cn-beijing",
        ///         InstanceType = "VPC",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GrantInstancesResult> InvokeAsync(GrantInstancesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GrantInstancesResult>("volcengine:cen/grantInstances:GrantInstances", args ?? new GrantInstancesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cen grant instances
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
        ///     var foo = Volcengine.Cen.GetGrantInstances.Invoke(new()
        ///     {
        ///         InstanceId = "vpc-2bysvq1xx543k2dx0eeul****",
        ///         InstanceRegionId = "cn-beijing",
        ///         InstanceType = "VPC",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GrantInstancesResult> Invoke(GrantInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GrantInstancesResult>("volcengine:cen/grantInstances:GrantInstances", args ?? new GrantInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GrantInstancesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// The region ID of the instance.
        /// </summary>
        [Input("instanceRegionId")]
        public string? InstanceRegionId { get; set; }

        /// <summary>
        /// The type of the instance. Valid values: `VPC`, `DCGW`.
        /// </summary>
        [Input("instanceType")]
        public string? InstanceType { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GrantInstancesArgs()
        {
        }
        public static new GrantInstancesArgs Empty => new GrantInstancesArgs();
    }

    public sealed class GrantInstancesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The region ID of the instance.
        /// </summary>
        [Input("instanceRegionId")]
        public Input<string>? InstanceRegionId { get; set; }

        /// <summary>
        /// The type of the instance. Valid values: `VPC`, `DCGW`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GrantInstancesInvokeArgs()
        {
        }
        public static new GrantInstancesInvokeArgs Empty => new GrantInstancesInvokeArgs();
    }


    [OutputType]
    public sealed class GrantInstancesResult
    {
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GrantInstancesGrantRuleResult> GrantRules;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the instance.
        /// </summary>
        public readonly string? InstanceId;
        /// <summary>
        /// The region ID of the instance.
        /// </summary>
        public readonly string? InstanceRegionId;
        /// <summary>
        /// The type of the instance.
        /// </summary>
        public readonly string? InstanceType;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GrantInstancesResult(
            ImmutableArray<Outputs.GrantInstancesGrantRuleResult> grantRules,

            string id,

            string? instanceId,

            string? instanceRegionId,

            string? instanceType,

            string? outputFile,

            int totalCount)
        {
            GrantRules = grantRules;
            Id = id;
            InstanceId = instanceId;
            InstanceRegionId = instanceRegionId;
            InstanceType = instanceType;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
