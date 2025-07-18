// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpc
{
    public static class GetTrafficMirrorTargets
    {
        /// <summary>
        /// Use this data source to query detailed information of traffic mirror targets
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
        ///     var foo = Volcengine.Vpc.GetTrafficMirrorTargets.Invoke(new()
        ///     {
        ///         TrafficMirrorTargetIds = new[]
        ///         {
        ///             "tmt-rry7yljufsw0v0x58w2****",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetTrafficMirrorTargetsResult> InvokeAsync(GetTrafficMirrorTargetsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTrafficMirrorTargetsResult>("volcengine:vpc/getTrafficMirrorTargets:getTrafficMirrorTargets", args ?? new GetTrafficMirrorTargetsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of traffic mirror targets
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
        ///     var foo = Volcengine.Vpc.GetTrafficMirrorTargets.Invoke(new()
        ///     {
        ///         TrafficMirrorTargetIds = new[]
        ///         {
        ///             "tmt-rry7yljufsw0v0x58w2****",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetTrafficMirrorTargetsResult> Invoke(GetTrafficMirrorTargetsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTrafficMirrorTargetsResult>("volcengine:vpc/getTrafficMirrorTargets:getTrafficMirrorTargets", args ?? new GetTrafficMirrorTargetsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTrafficMirrorTargetsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A Name Regex of Resource.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The project name of traffic mirror target.
        /// </summary>
        [Input("projectName")]
        public string? ProjectName { get; set; }

        [Input("tags")]
        private List<Inputs.GetTrafficMirrorTargetsTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public List<Inputs.GetTrafficMirrorTargetsTagArgs> Tags
        {
            get => _tags ?? (_tags = new List<Inputs.GetTrafficMirrorTargetsTagArgs>());
            set => _tags = value;
        }

        [Input("trafficMirrorTargetIds")]
        private List<string>? _trafficMirrorTargetIds;

        /// <summary>
        /// A list of traffic mirror target IDs.
        /// </summary>
        public List<string> TrafficMirrorTargetIds
        {
            get => _trafficMirrorTargetIds ?? (_trafficMirrorTargetIds = new List<string>());
            set => _trafficMirrorTargetIds = value;
        }

        /// <summary>
        /// The name of traffic mirror target.
        /// </summary>
        [Input("trafficMirrorTargetName")]
        public string? TrafficMirrorTargetName { get; set; }

        public GetTrafficMirrorTargetsArgs()
        {
        }
        public static new GetTrafficMirrorTargetsArgs Empty => new GetTrafficMirrorTargetsArgs();
    }

    public sealed class GetTrafficMirrorTargetsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A Name Regex of Resource.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The project name of traffic mirror target.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        [Input("tags")]
        private InputList<Inputs.GetTrafficMirrorTargetsTagInputArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.GetTrafficMirrorTargetsTagInputArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.GetTrafficMirrorTargetsTagInputArgs>());
            set => _tags = value;
        }

        [Input("trafficMirrorTargetIds")]
        private InputList<string>? _trafficMirrorTargetIds;

        /// <summary>
        /// A list of traffic mirror target IDs.
        /// </summary>
        public InputList<string> TrafficMirrorTargetIds
        {
            get => _trafficMirrorTargetIds ?? (_trafficMirrorTargetIds = new InputList<string>());
            set => _trafficMirrorTargetIds = value;
        }

        /// <summary>
        /// The name of traffic mirror target.
        /// </summary>
        [Input("trafficMirrorTargetName")]
        public Input<string>? TrafficMirrorTargetName { get; set; }

        public GetTrafficMirrorTargetsInvokeArgs()
        {
        }
        public static new GetTrafficMirrorTargetsInvokeArgs Empty => new GetTrafficMirrorTargetsInvokeArgs();
    }


    [OutputType]
    public sealed class GetTrafficMirrorTargetsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The project name of traffic mirror target.
        /// </summary>
        public readonly string? ProjectName;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTrafficMirrorTargetsTagResult> Tags;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;
        public readonly ImmutableArray<string> TrafficMirrorTargetIds;
        /// <summary>
        /// The name of traffic mirror target.
        /// </summary>
        public readonly string? TrafficMirrorTargetName;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTrafficMirrorTargetsTrafficMirrorTargetResult> TrafficMirrorTargets;

        [OutputConstructor]
        private GetTrafficMirrorTargetsResult(
            string id,

            string? nameRegex,

            string? outputFile,

            string? projectName,

            ImmutableArray<Outputs.GetTrafficMirrorTargetsTagResult> tags,

            int totalCount,

            ImmutableArray<string> trafficMirrorTargetIds,

            string? trafficMirrorTargetName,

            ImmutableArray<Outputs.GetTrafficMirrorTargetsTrafficMirrorTargetResult> trafficMirrorTargets)
        {
            Id = id;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            ProjectName = projectName;
            Tags = tags;
            TotalCount = totalCount;
            TrafficMirrorTargetIds = trafficMirrorTargetIds;
            TrafficMirrorTargetName = trafficMirrorTargetName;
            TrafficMirrorTargets = trafficMirrorTargets;
        }
    }
}
