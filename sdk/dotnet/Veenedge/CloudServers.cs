// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Veenedge
{
    public static class CloudServers
    {
        /// <summary>
        /// Use this data source to query detailed information of veenedge cloud servers
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
        ///     var @default = Volcengine.Veenedge.CloudServers.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<CloudServersResult> InvokeAsync(CloudServersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<CloudServersResult>("volcengine:veenedge/cloudServers:CloudServers", args ?? new CloudServersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of veenedge cloud servers
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
        ///     var @default = Volcengine.Veenedge.CloudServers.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<CloudServersResult> Invoke(CloudServersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<CloudServersResult>("volcengine:veenedge/cloudServers:CloudServers", args ?? new CloudServersInvokeArgs(), options.WithDefaults());
    }


    public sealed class CloudServersArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of cloud server IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of Cloud Server.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public CloudServersArgs()
        {
        }
        public static new CloudServersArgs Empty => new CloudServersArgs();
    }

    public sealed class CloudServersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of cloud server IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of Cloud Server.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public CloudServersInvokeArgs()
        {
        }
        public static new CloudServersInvokeArgs Empty => new CloudServersInvokeArgs();
    }


    [OutputType]
    public sealed class CloudServersResult
    {
        /// <summary>
        /// The collection of cloud servers query.
        /// </summary>
        public readonly ImmutableArray<Outputs.CloudServersCloudServerResult> CloudServers;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of cloud servers query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private CloudServersResult(
            ImmutableArray<Outputs.CloudServersCloudServerResult> cloudServers,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            string? outputFile,

            int totalCount)
        {
            CloudServers = cloudServers;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}