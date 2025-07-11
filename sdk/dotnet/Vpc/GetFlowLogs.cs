// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpc
{
    public static class GetFlowLogs
    {
        /// <summary>
        /// Use this data source to query detailed information of flow logs
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
        ///     var foo = Volcengine.Vpc.GetFlowLogs.Invoke(new()
        ///     {
        ///         FlowLogIds = new[]
        ///         {
        ///             "fl-13g4fqngluhog3n6nu57o****",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetFlowLogsResult> InvokeAsync(GetFlowLogsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFlowLogsResult>("volcengine:vpc/getFlowLogs:getFlowLogs", args ?? new GetFlowLogsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of flow logs
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
        ///     var foo = Volcengine.Vpc.GetFlowLogs.Invoke(new()
        ///     {
        ///         FlowLogIds = new[]
        ///         {
        ///             "fl-13g4fqngluhog3n6nu57o****",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetFlowLogsResult> Invoke(GetFlowLogsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFlowLogsResult>("volcengine:vpc/getFlowLogs:getFlowLogs", args ?? new GetFlowLogsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFlowLogsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The aggregation interval of flow log. Unit: minute. Valid values: `1`, `5`, `10`.
        /// </summary>
        [Input("aggregationInterval")]
        public int? AggregationInterval { get; set; }

        /// <summary>
        /// The description of flow log.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        [Input("flowLogIds")]
        private List<string>? _flowLogIds;

        /// <summary>
        /// A list of flow log IDs.
        /// </summary>
        public List<string> FlowLogIds
        {
            get => _flowLogIds ?? (_flowLogIds = new List<string>());
            set => _flowLogIds = value;
        }

        /// <summary>
        /// The name of flow log.
        /// </summary>
        [Input("flowLogName")]
        public string? FlowLogName { get; set; }

        /// <summary>
        /// The ID of log project.
        /// </summary>
        [Input("logProjectId")]
        public string? LogProjectId { get; set; }

        /// <summary>
        /// The ID of log topic.
        /// </summary>
        [Input("logTopicId")]
        public string? LogTopicId { get; set; }

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
        /// The project name of flow log.
        /// </summary>
        [Input("projectName")]
        public string? ProjectName { get; set; }

        /// <summary>
        /// The ID of resource.
        /// </summary>
        [Input("resourceId")]
        public string? ResourceId { get; set; }

        /// <summary>
        /// The type of resource. Valid values: `vpc`, `subnet`, `eni`.
        /// </summary>
        [Input("resourceType")]
        public string? ResourceType { get; set; }

        /// <summary>
        /// The status of flow log. Valid values: `Active`, `Pending`, `Inactive`, `Creating`, `Deleting`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        [Input("tags")]
        private List<Inputs.GetFlowLogsTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public List<Inputs.GetFlowLogsTagArgs> Tags
        {
            get => _tags ?? (_tags = new List<Inputs.GetFlowLogsTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of traffic. Valid values: `All`, `Allow`, `Drop`.
        /// </summary>
        [Input("trafficType")]
        public string? TrafficType { get; set; }

        /// <summary>
        /// The ID of VPC.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetFlowLogsArgs()
        {
        }
        public static new GetFlowLogsArgs Empty => new GetFlowLogsArgs();
    }

    public sealed class GetFlowLogsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The aggregation interval of flow log. Unit: minute. Valid values: `1`, `5`, `10`.
        /// </summary>
        [Input("aggregationInterval")]
        public Input<int>? AggregationInterval { get; set; }

        /// <summary>
        /// The description of flow log.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("flowLogIds")]
        private InputList<string>? _flowLogIds;

        /// <summary>
        /// A list of flow log IDs.
        /// </summary>
        public InputList<string> FlowLogIds
        {
            get => _flowLogIds ?? (_flowLogIds = new InputList<string>());
            set => _flowLogIds = value;
        }

        /// <summary>
        /// The name of flow log.
        /// </summary>
        [Input("flowLogName")]
        public Input<string>? FlowLogName { get; set; }

        /// <summary>
        /// The ID of log project.
        /// </summary>
        [Input("logProjectId")]
        public Input<string>? LogProjectId { get; set; }

        /// <summary>
        /// The ID of log topic.
        /// </summary>
        [Input("logTopicId")]
        public Input<string>? LogTopicId { get; set; }

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
        /// The project name of flow log.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The ID of resource.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// The type of resource. Valid values: `vpc`, `subnet`, `eni`.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        /// <summary>
        /// The status of flow log. Valid values: `Active`, `Pending`, `Inactive`, `Creating`, `Deleting`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<Inputs.GetFlowLogsTagInputArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.GetFlowLogsTagInputArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.GetFlowLogsTagInputArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of traffic. Valid values: `All`, `Allow`, `Drop`.
        /// </summary>
        [Input("trafficType")]
        public Input<string>? TrafficType { get; set; }

        /// <summary>
        /// The ID of VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GetFlowLogsInvokeArgs()
        {
        }
        public static new GetFlowLogsInvokeArgs Empty => new GetFlowLogsInvokeArgs();
    }


    [OutputType]
    public sealed class GetFlowLogsResult
    {
        /// <summary>
        /// The aggregation interval of flow log. Unit: minute. Valid values: `1`, `5`, `10`.
        /// </summary>
        public readonly int? AggregationInterval;
        /// <summary>
        /// The description of flow log.
        /// </summary>
        public readonly string? Description;
        public readonly ImmutableArray<string> FlowLogIds;
        /// <summary>
        /// The name of flow log.
        /// </summary>
        public readonly string? FlowLogName;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFlowLogsFlowLogResult> FlowLogs;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of log project.
        /// </summary>
        public readonly string? LogProjectId;
        /// <summary>
        /// The ID of log topic.
        /// </summary>
        public readonly string? LogTopicId;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The project name of flow log.
        /// </summary>
        public readonly string? ProjectName;
        /// <summary>
        /// The ID of resource.
        /// </summary>
        public readonly string? ResourceId;
        /// <summary>
        /// The type of resource. Valid values: `vpc`, `subnet`, `eni`.
        /// </summary>
        public readonly string? ResourceType;
        /// <summary>
        /// The status of flow log. Valid values: `Active`, `Pending`, `Inactive`, `Creating`, `Deleting`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFlowLogsTagResult> Tags;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The type of traffic. Valid values: `All`, `Allow`, `Drop`.
        /// </summary>
        public readonly string? TrafficType;
        /// <summary>
        /// The ID of VPC.
        /// </summary>
        public readonly string? VpcId;

        [OutputConstructor]
        private GetFlowLogsResult(
            int? aggregationInterval,

            string? description,

            ImmutableArray<string> flowLogIds,

            string? flowLogName,

            ImmutableArray<Outputs.GetFlowLogsFlowLogResult> flowLogs,

            string id,

            string? logProjectId,

            string? logTopicId,

            string? nameRegex,

            string? outputFile,

            string? projectName,

            string? resourceId,

            string? resourceType,

            string? status,

            ImmutableArray<Outputs.GetFlowLogsTagResult> tags,

            int totalCount,

            string? trafficType,

            string? vpcId)
        {
            AggregationInterval = aggregationInterval;
            Description = description;
            FlowLogIds = flowLogIds;
            FlowLogName = flowLogName;
            FlowLogs = flowLogs;
            Id = id;
            LogProjectId = logProjectId;
            LogTopicId = logTopicId;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            ProjectName = projectName;
            ResourceId = resourceId;
            ResourceType = resourceType;
            Status = status;
            Tags = tags;
            TotalCount = totalCount;
            TrafficType = trafficType;
            VpcId = vpcId;
        }
    }
}
