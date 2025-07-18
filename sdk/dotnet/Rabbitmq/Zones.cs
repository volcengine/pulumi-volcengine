// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rabbitmq
{
    [Obsolete(@"volcengine.rabbitmq.Zones has been deprecated in favor of volcengine.rabbitmq.getZones")]
    public static class Zones
    {
        /// <summary>
        /// Use this data source to query detailed information of rabbitmq zones
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
        ///     var foo = Volcengine.Rabbitmq.GetZones.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<ZonesResult> InvokeAsync(ZonesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ZonesResult>("volcengine:rabbitmq/zones:Zones", args ?? new ZonesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of rabbitmq zones
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
        ///     var foo = Volcengine.Rabbitmq.GetZones.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<ZonesResult> Invoke(ZonesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ZonesResult>("volcengine:rabbitmq/zones:Zones", args ?? new ZonesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ZonesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public ZonesArgs()
        {
        }
        public static new ZonesArgs Empty => new ZonesArgs();
    }

    public sealed class ZonesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public ZonesInvokeArgs()
        {
        }
        public static new ZonesInvokeArgs Empty => new ZonesInvokeArgs();
    }


    [OutputType]
    public sealed class ZonesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.ZonesZoneResult> Zones;

        [OutputConstructor]
        private ZonesResult(
            string id,

            string? outputFile,

            int totalCount,

            ImmutableArray<Outputs.ZonesZoneResult> zones)
        {
            Id = id;
            OutputFile = outputFile;
            TotalCount = totalCount;
            Zones = zones;
        }
    }
}
