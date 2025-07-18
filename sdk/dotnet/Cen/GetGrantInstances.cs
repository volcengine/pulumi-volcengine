// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cen
{
    public static class GetGrantInstances
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
        public static Task<GetGrantInstancesResult> InvokeAsync(GetGrantInstancesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGrantInstancesResult>("volcengine:cen/getGrantInstances:getGrantInstances", args ?? new GetGrantInstancesArgs(), options.WithDefaults());

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
        public static Output<GetGrantInstancesResult> Invoke(GetGrantInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGrantInstancesResult>("volcengine:cen/getGrantInstances:getGrantInstances", args ?? new GetGrantInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGrantInstancesArgs : global::Pulumi.InvokeArgs
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

        public GetGrantInstancesArgs()
        {
        }
        public static new GetGrantInstancesArgs Empty => new GetGrantInstancesArgs();
    }

    public sealed class GetGrantInstancesInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetGrantInstancesInvokeArgs()
        {
        }
        public static new GetGrantInstancesInvokeArgs Empty => new GetGrantInstancesInvokeArgs();
    }


    [OutputType]
    public sealed class GetGrantInstancesResult
    {
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGrantInstancesGrantRuleResult> GrantRules;
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
        private GetGrantInstancesResult(
            ImmutableArray<Outputs.GetGrantInstancesGrantRuleResult> grantRules,

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
