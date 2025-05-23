// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Clb
{
    [Obsolete(@"volcengine.clb.Rules has been deprecated in favor of volcengine.clb.getRules")]
    public static class Rules
    {
        /// <summary>
        /// Use this data source to query detailed information of clb rules
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
        ///         Description = "acc0Demo",
        ///         LoadBalancerName = "acc-test-create",
        ///         EipBillingConfig = new Volcengine.Clb.Inputs.ClbEipBillingConfigArgs
        ///         {
        ///             Isp = "BGP",
        ///             EipBillingType = "PostPaidByBandwidth",
        ///             Bandwidth = 1,
        ///         },
        ///     });
        /// 
        ///     var fooServerGroup = new Volcengine.Clb.ServerGroup("fooServerGroup", new()
        ///     {
        ///         LoadBalancerId = fooClb.Id,
        ///         ServerGroupName = "acc-test-create",
        ///         Description = "hello demo11",
        ///     });
        /// 
        ///     var fooListener = new Volcengine.Clb.Listener("fooListener", new()
        ///     {
        ///         LoadBalancerId = fooClb.Id,
        ///         ListenerName = "acc-test-listener",
        ///         Protocol = "HTTP",
        ///         Port = 90,
        ///         ServerGroupId = fooServerGroup.Id,
        ///         HealthCheck = new Volcengine.Clb.Inputs.ListenerHealthCheckArgs
        ///         {
        ///             Enabled = "on",
        ///             Interval = 10,
        ///             Timeout = 3,
        ///             HealthyThreshold = 5,
        ///             UnHealthyThreshold = 2,
        ///             Domain = "volcengine.com",
        ///             HttpCode = "http_2xx",
        ///             Method = "GET",
        ///             Uri = "/",
        ///         },
        ///         Enabled = "on",
        ///     });
        /// 
        ///     var fooRule = new Volcengine.Clb.Rule("fooRule", new()
        ///     {
        ///         ListenerId = fooListener.Id,
        ///         ServerGroupId = fooServerGroup.Id,
        ///         Domain = "test-volc123.com",
        ///         Url = "/yyyy",
        ///     });
        /// 
        ///     var fooRules = Volcengine.Clb.GetRules.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             fooRule.Id,
        ///         },
        ///         ListenerId = fooListener.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<RulesResult> InvokeAsync(RulesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<RulesResult>("volcengine:clb/rules:Rules", args ?? new RulesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of clb rules
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
        ///         Description = "acc0Demo",
        ///         LoadBalancerName = "acc-test-create",
        ///         EipBillingConfig = new Volcengine.Clb.Inputs.ClbEipBillingConfigArgs
        ///         {
        ///             Isp = "BGP",
        ///             EipBillingType = "PostPaidByBandwidth",
        ///             Bandwidth = 1,
        ///         },
        ///     });
        /// 
        ///     var fooServerGroup = new Volcengine.Clb.ServerGroup("fooServerGroup", new()
        ///     {
        ///         LoadBalancerId = fooClb.Id,
        ///         ServerGroupName = "acc-test-create",
        ///         Description = "hello demo11",
        ///     });
        /// 
        ///     var fooListener = new Volcengine.Clb.Listener("fooListener", new()
        ///     {
        ///         LoadBalancerId = fooClb.Id,
        ///         ListenerName = "acc-test-listener",
        ///         Protocol = "HTTP",
        ///         Port = 90,
        ///         ServerGroupId = fooServerGroup.Id,
        ///         HealthCheck = new Volcengine.Clb.Inputs.ListenerHealthCheckArgs
        ///         {
        ///             Enabled = "on",
        ///             Interval = 10,
        ///             Timeout = 3,
        ///             HealthyThreshold = 5,
        ///             UnHealthyThreshold = 2,
        ///             Domain = "volcengine.com",
        ///             HttpCode = "http_2xx",
        ///             Method = "GET",
        ///             Uri = "/",
        ///         },
        ///         Enabled = "on",
        ///     });
        /// 
        ///     var fooRule = new Volcengine.Clb.Rule("fooRule", new()
        ///     {
        ///         ListenerId = fooListener.Id,
        ///         ServerGroupId = fooServerGroup.Id,
        ///         Domain = "test-volc123.com",
        ///         Url = "/yyyy",
        ///     });
        /// 
        ///     var fooRules = Volcengine.Clb.GetRules.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             fooRule.Id,
        ///         },
        ///         ListenerId = fooListener.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<RulesResult> Invoke(RulesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<RulesResult>("volcengine:clb/rules:Rules", args ?? new RulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class RulesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Rule IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The Id of listener.
        /// </summary>
        [Input("listenerId", required: true)]
        public string ListenerId { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public RulesArgs()
        {
        }
        public static new RulesArgs Empty => new RulesArgs();
    }

    public sealed class RulesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Rule IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// The Id of listener.
        /// </summary>
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public RulesInvokeArgs()
        {
        }
        public static new RulesInvokeArgs Empty => new RulesInvokeArgs();
    }


    [OutputType]
    public sealed class RulesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string ListenerId;
        public readonly string? OutputFile;
        /// <summary>
        /// The collection of Rule query.
        /// </summary>
        public readonly ImmutableArray<Outputs.RulesRuleResult> Rules;

        [OutputConstructor]
        private RulesResult(
            string id,

            ImmutableArray<string> ids,

            string listenerId,

            string? outputFile,

            ImmutableArray<Outputs.RulesRuleResult> rules)
        {
            Id = id;
            Ids = ids;
            ListenerId = listenerId;
            OutputFile = outputFile;
            Rules = rules;
        }
    }
}
