// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Clb
{
    public static class Clbs
    {
        /// <summary>
        /// Use this data source to query detailed information of clbs
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
        ///         VpcName = "acc-test-vpc",
        ///         CidrBlock = "172.16.0.0/16",
        ///     });
        /// 
        ///     var fooSubnet = new Volcengine.Vpc.Subnet("fooSubnet", new()
        ///     {
        ///         SubnetName = "acc-test-subnet",
        ///         CidrBlock = "172.16.0.0/24",
        ///         ZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[0]?.Id),
        ///         VpcId = fooVpc.Id,
        ///     });
        /// 
        ///     var fooClb = new List&lt;Volcengine.Clb.Clb&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; 3; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         fooClb.Add(new Volcengine.Clb.Clb($"fooClb-{range.Value}", new()
        ///         {
        ///             Type = "public",
        ///             SubnetId = fooSubnet.Id,
        ///             LoadBalancerSpec = "small_1",
        ///             Description = "acc-test-demo",
        ///             LoadBalancerName = $"acc-test-clb-{range.Value}",
        ///             LoadBalancerBillingType = "PostPaid",
        ///             EipBillingConfig = new Volcengine.Clb.Inputs.ClbEipBillingConfigArgs
        ///             {
        ///                 Isp = "BGP",
        ///                 EipBillingType = "PostPaidByBandwidth",
        ///                 Bandwidth = 1,
        ///             },
        ///             Tags = new[]
        ///             {
        ///                 new Volcengine.Clb.Inputs.ClbTagArgs
        ///                 {
        ///                     Key = "k1",
        ///                     Value = "v1",
        ///                 },
        ///             },
        ///         }));
        ///     }
        ///     var fooClbs = Volcengine.Clb.Clbs.Invoke(new()
        ///     {
        ///         Ids = fooClb.Select(__item =&gt; __item.Id).ToList(),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<ClbsResult> InvokeAsync(ClbsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ClbsResult>("volcengine:clb/clbs:Clbs", args ?? new ClbsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of clbs
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
        ///         VpcName = "acc-test-vpc",
        ///         CidrBlock = "172.16.0.0/16",
        ///     });
        /// 
        ///     var fooSubnet = new Volcengine.Vpc.Subnet("fooSubnet", new()
        ///     {
        ///         SubnetName = "acc-test-subnet",
        ///         CidrBlock = "172.16.0.0/24",
        ///         ZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[0]?.Id),
        ///         VpcId = fooVpc.Id,
        ///     });
        /// 
        ///     var fooClb = new List&lt;Volcengine.Clb.Clb&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; 3; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         fooClb.Add(new Volcengine.Clb.Clb($"fooClb-{range.Value}", new()
        ///         {
        ///             Type = "public",
        ///             SubnetId = fooSubnet.Id,
        ///             LoadBalancerSpec = "small_1",
        ///             Description = "acc-test-demo",
        ///             LoadBalancerName = $"acc-test-clb-{range.Value}",
        ///             LoadBalancerBillingType = "PostPaid",
        ///             EipBillingConfig = new Volcengine.Clb.Inputs.ClbEipBillingConfigArgs
        ///             {
        ///                 Isp = "BGP",
        ///                 EipBillingType = "PostPaidByBandwidth",
        ///                 Bandwidth = 1,
        ///             },
        ///             Tags = new[]
        ///             {
        ///                 new Volcengine.Clb.Inputs.ClbTagArgs
        ///                 {
        ///                     Key = "k1",
        ///                     Value = "v1",
        ///                 },
        ///             },
        ///         }));
        ///     }
        ///     var fooClbs = Volcengine.Clb.Clbs.Invoke(new()
        ///     {
        ///         Ids = fooClb.Select(__item =&gt; __item.Id).ToList(),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<ClbsResult> Invoke(ClbsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ClbsResult>("volcengine:clb/clbs:Clbs", args ?? new ClbsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ClbsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The private ip address of the Clb.
        /// </summary>
        [Input("eniAddress")]
        public string? EniAddress { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Clb IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of the Clb.
        /// </summary>
        [Input("loadBalancerName")]
        public string? LoadBalancerName { get; set; }

        /// <summary>
        /// A Name Regex of Clb.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The ProjectName of Clb.
        /// </summary>
        [Input("projectName")]
        public string? ProjectName { get; set; }

        [Input("tags")]
        private List<Inputs.ClbsTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public List<Inputs.ClbsTagArgs> Tags
        {
            get => _tags ?? (_tags = new List<Inputs.ClbsTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The id of the VPC.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public ClbsArgs()
        {
        }
        public static new ClbsArgs Empty => new ClbsArgs();
    }

    public sealed class ClbsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The private ip address of the Clb.
        /// </summary>
        [Input("eniAddress")]
        public Input<string>? EniAddress { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Clb IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The name of the Clb.
        /// </summary>
        [Input("loadBalancerName")]
        public Input<string>? LoadBalancerName { get; set; }

        /// <summary>
        /// A Name Regex of Clb.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The ProjectName of Clb.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        [Input("tags")]
        private InputList<Inputs.ClbsTagInputArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.ClbsTagInputArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ClbsTagInputArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The id of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public ClbsInvokeArgs()
        {
        }
        public static new ClbsInvokeArgs Empty => new ClbsInvokeArgs();
    }


    [OutputType]
    public sealed class ClbsResult
    {
        /// <summary>
        /// The collection of Clb query.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClbsClbResult> Clbs;
        /// <summary>
        /// The Eni address of the Clb.
        /// </summary>
        public readonly string? EniAddress;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// The name of the Clb.
        /// </summary>
        public readonly string? LoadBalancerName;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The ProjectName of the Clb.
        /// </summary>
        public readonly string? ProjectName;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClbsTagResult> Tags;
        /// <summary>
        /// The total count of Clb query.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The vpc ID of the Clb.
        /// </summary>
        public readonly string? VpcId;

        [OutputConstructor]
        private ClbsResult(
            ImmutableArray<Outputs.ClbsClbResult> clbs,

            string? eniAddress,

            string id,

            ImmutableArray<string> ids,

            string? loadBalancerName,

            string? nameRegex,

            string? outputFile,

            string? projectName,

            ImmutableArray<Outputs.ClbsTagResult> tags,

            int totalCount,

            string? vpcId)
        {
            Clbs = clbs;
            EniAddress = eniAddress;
            Id = id;
            Ids = ids;
            LoadBalancerName = loadBalancerName;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            ProjectName = projectName;
            Tags = tags;
            TotalCount = totalCount;
            VpcId = vpcId;
        }
    }
}