// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls
{
    public static class GetHosts
    {
        /// <summary>
        /// Use this data source to query detailed information of tls hosts
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
        ///     var @default = Volcengine.Tls.GetHosts.Invoke(new()
        ///     {
        ///         HostGroupId = "527102e2-1e4f-45f4-a990-751152125da7",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetHostsResult> InvokeAsync(GetHostsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHostsResult>("volcengine:tls/getHosts:getHosts", args ?? new GetHostsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tls hosts
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
        ///     var @default = Volcengine.Tls.GetHosts.Invoke(new()
        ///     {
        ///         HostGroupId = "527102e2-1e4f-45f4-a990-751152125da7",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetHostsResult> Invoke(GetHostsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHostsResult>("volcengine:tls/getHosts:getHosts", args ?? new GetHostsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHostsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The the heartbeat status.
        /// </summary>
        [Input("heartbeatStatus")]
        public int? HeartbeatStatus { get; set; }

        /// <summary>
        /// The id of host group.
        /// </summary>
        [Input("hostGroupId", required: true)]
        public string HostGroupId { get; set; } = null!;

        /// <summary>
        /// The ip address.
        /// </summary>
        [Input("ip")]
        public string? Ip { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public GetHostsArgs()
        {
        }
        public static new GetHostsArgs Empty => new GetHostsArgs();
    }

    public sealed class GetHostsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The the heartbeat status.
        /// </summary>
        [Input("heartbeatStatus")]
        public Input<int>? HeartbeatStatus { get; set; }

        /// <summary>
        /// The id of host group.
        /// </summary>
        [Input("hostGroupId", required: true)]
        public Input<string> HostGroupId { get; set; } = null!;

        /// <summary>
        /// The ip address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public GetHostsInvokeArgs()
        {
        }
        public static new GetHostsInvokeArgs Empty => new GetHostsInvokeArgs();
    }


    [OutputType]
    public sealed class GetHostsResult
    {
        /// <summary>
        /// The the heartbeat status.
        /// </summary>
        public readonly int? HeartbeatStatus;
        /// <summary>
        /// The id of host group.
        /// </summary>
        public readonly string HostGroupId;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetHostsHostInfoResult> HostInfos;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ip address.
        /// </summary>
        public readonly string? Ip;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetHostsResult(
            int? heartbeatStatus,

            string hostGroupId,

            ImmutableArray<Outputs.GetHostsHostInfoResult> hostInfos,

            string id,

            string? ip,

            string? outputFile,

            int totalCount)
        {
            HeartbeatStatus = heartbeatStatus;
            HostGroupId = hostGroupId;
            HostInfos = hostInfos;
            Id = id;
            Ip = ip;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
