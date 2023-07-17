// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls
{
    public static class RuleAppliers
    {
        /// <summary>
        /// Use this data source to query detailed information of tls rule appliers
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
        ///         var @default = Output.Create(Volcengine.Tls.RuleAppliers.InvokeAsync(new Volcengine.Tls.RuleAppliersArgs
        ///         {
        ///             HostGroupId = "fbea6619-7b0c-40f3-ac7e-45c63e3f676e",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<RuleAppliersResult> InvokeAsync(RuleAppliersArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<RuleAppliersResult>("volcengine:tls/ruleAppliers:RuleAppliers", args ?? new RuleAppliersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tls rule appliers
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
        ///         var @default = Output.Create(Volcengine.Tls.RuleAppliers.InvokeAsync(new Volcengine.Tls.RuleAppliersArgs
        ///         {
        ///             HostGroupId = "fbea6619-7b0c-40f3-ac7e-45c63e3f676e",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<RuleAppliersResult> Invoke(RuleAppliersInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<RuleAppliersResult>("volcengine:tls/ruleAppliers:RuleAppliers", args ?? new RuleAppliersInvokeArgs(), options.WithDefaults());
    }


    public sealed class RuleAppliersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The host group id.
        /// </summary>
        [Input("hostGroupId", required: true)]
        public string HostGroupId { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public RuleAppliersArgs()
        {
        }
    }

    public sealed class RuleAppliersInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The host group id.
        /// </summary>
        [Input("hostGroupId", required: true)]
        public Input<string> HostGroupId { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public RuleAppliersInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class RuleAppliersResult
    {
        public readonly string HostGroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OutputFile;
        /// <summary>
        /// The rules list.
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleAppliersRuleResult> Rules;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private RuleAppliersResult(
            string hostGroupId,

            string id,

            string? outputFile,

            ImmutableArray<Outputs.RuleAppliersRuleResult> rules,

            int totalCount)
        {
            HostGroupId = hostGroupId;
            Id = id;
            OutputFile = outputFile;
            Rules = rules;
            TotalCount = totalCount;
        }
    }
}