// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Privatelink
{
    [Obsolete(@"volcengine.privatelink.VpcEndpointServices has been deprecated in favor of volcengine.privatelink.getVpcEndpointServices")]
    public static class VpcEndpointServices
    {
        /// <summary>
        /// Use this data source to query detailed information of privatelink vpc endpoint services
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
        ///     var fooClb = new Volcengine.Clb.Clb("fooClb", new()
        ///     {
        ///         Type = "public",
        ///         SubnetId = fooSubnet.Id,
        ///         LoadBalancerSpec = "small_1",
        ///         Description = "acc-test-demo",
        ///         LoadBalancerName = "acc-test-clb",
        ///         LoadBalancerBillingType = "PostPaid",
        ///         EipBillingConfig = new Volcengine.Clb.Inputs.ClbEipBillingConfigArgs
        ///         {
        ///             Isp = "BGP",
        ///             EipBillingType = "PostPaidByBandwidth",
        ///             Bandwidth = 1,
        ///         },
        ///         Tags = new[]
        ///         {
        ///             new Volcengine.Clb.Inputs.ClbTagArgs
        ///             {
        ///                 Key = "k1",
        ///                 Value = "v1",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooVpcEndpointService = new List&lt;Volcengine.Privatelink.VpcEndpointService&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; 2; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         fooVpcEndpointService.Add(new Volcengine.Privatelink.VpcEndpointService($"fooVpcEndpointService-{range.Value}", new()
        ///         {
        ///             Resources = new[]
        ///             {
        ///                 new Volcengine.Privatelink.Inputs.VpcEndpointServiceResourceArgs
        ///                 {
        ///                     ResourceId = fooClb.Id,
        ///                     ResourceType = "CLB",
        ///                 },
        ///             },
        ///             Description = "acc-test",
        ///             AutoAcceptEnabled = true,
        ///         }));
        ///     }
        ///     var fooVpcEndpointServices = Volcengine.Privatelink.GetVpcEndpointServices.Invoke(new()
        ///     {
        ///         Ids = fooVpcEndpointService.Select(__item =&gt; __item.Id).ToList(),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<VpcEndpointServicesResult> InvokeAsync(VpcEndpointServicesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<VpcEndpointServicesResult>("volcengine:privatelink/vpcEndpointServices:VpcEndpointServices", args ?? new VpcEndpointServicesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of privatelink vpc endpoint services
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
        ///     var fooClb = new Volcengine.Clb.Clb("fooClb", new()
        ///     {
        ///         Type = "public",
        ///         SubnetId = fooSubnet.Id,
        ///         LoadBalancerSpec = "small_1",
        ///         Description = "acc-test-demo",
        ///         LoadBalancerName = "acc-test-clb",
        ///         LoadBalancerBillingType = "PostPaid",
        ///         EipBillingConfig = new Volcengine.Clb.Inputs.ClbEipBillingConfigArgs
        ///         {
        ///             Isp = "BGP",
        ///             EipBillingType = "PostPaidByBandwidth",
        ///             Bandwidth = 1,
        ///         },
        ///         Tags = new[]
        ///         {
        ///             new Volcengine.Clb.Inputs.ClbTagArgs
        ///             {
        ///                 Key = "k1",
        ///                 Value = "v1",
        ///             },
        ///         },
        ///     });
        /// 
        ///     var fooVpcEndpointService = new List&lt;Volcengine.Privatelink.VpcEndpointService&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; 2; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         fooVpcEndpointService.Add(new Volcengine.Privatelink.VpcEndpointService($"fooVpcEndpointService-{range.Value}", new()
        ///         {
        ///             Resources = new[]
        ///             {
        ///                 new Volcengine.Privatelink.Inputs.VpcEndpointServiceResourceArgs
        ///                 {
        ///                     ResourceId = fooClb.Id,
        ///                     ResourceType = "CLB",
        ///                 },
        ///             },
        ///             Description = "acc-test",
        ///             AutoAcceptEnabled = true,
        ///         }));
        ///     }
        ///     var fooVpcEndpointServices = Volcengine.Privatelink.GetVpcEndpointServices.Invoke(new()
        ///     {
        ///         Ids = fooVpcEndpointService.Select(__item =&gt; __item.Id).ToList(),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<VpcEndpointServicesResult> Invoke(VpcEndpointServicesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<VpcEndpointServicesResult>("volcengine:privatelink/vpcEndpointServices:VpcEndpointServices", args ?? new VpcEndpointServicesInvokeArgs(), options.WithDefaults());
    }


    public sealed class VpcEndpointServicesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// The IDs of vpc endpoint service.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of vpc endpoint service.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The name of vpc endpoint service.
        /// </summary>
        [Input("serviceName")]
        public string? ServiceName { get; set; }

        public VpcEndpointServicesArgs()
        {
        }
        public static new VpcEndpointServicesArgs Empty => new VpcEndpointServicesArgs();
    }

    public sealed class VpcEndpointServicesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// The IDs of vpc endpoint service.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of vpc endpoint service.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The name of vpc endpoint service.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public VpcEndpointServicesInvokeArgs()
        {
        }
        public static new VpcEndpointServicesInvokeArgs Empty => new VpcEndpointServicesInvokeArgs();
    }


    [OutputType]
    public sealed class VpcEndpointServicesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The name of service.
        /// </summary>
        public readonly string? ServiceName;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.VpcEndpointServicesServiceResult> Services;
        /// <summary>
        /// Returns the total amount of the data list.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private VpcEndpointServicesResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            string? outputFile,

            string? serviceName,

            ImmutableArray<Outputs.VpcEndpointServicesServiceResult> services,

            int totalCount)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            ServiceName = serviceName;
            Services = services;
            TotalCount = totalCount;
        }
    }
}
