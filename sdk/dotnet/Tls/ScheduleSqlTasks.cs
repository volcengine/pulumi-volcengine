// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls
{
    [Obsolete(@"volcengine.tls.ScheduleSqlTasks has been deprecated in favor of volcengine.tls.getScheduleSqlTasks")]
    public static class ScheduleSqlTasks
    {
        /// <summary>
        /// Use this data source to query detailed information of tls schedule sql tasks
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
        ///     var @default = Volcengine.Tls.GetScheduleSqlTasks.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<ScheduleSqlTasksResult> InvokeAsync(ScheduleSqlTasksArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ScheduleSqlTasksResult>("volcengine:tls/scheduleSqlTasks:ScheduleSqlTasks", args ?? new ScheduleSqlTasksArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tls schedule sql tasks
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
        ///     var @default = Volcengine.Tls.GetScheduleSqlTasks.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<ScheduleSqlTasksResult> Invoke(ScheduleSqlTasksInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ScheduleSqlTasksResult>("volcengine:tls/scheduleSqlTasks:ScheduleSqlTasks", args ?? new ScheduleSqlTasksInvokeArgs(), options.WithDefaults());
    }


    public sealed class ScheduleSqlTasksArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// IAM log project name.
        /// </summary>
        [Input("iamProjectName")]
        public string? IamProjectName { get; set; }

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
        /// The log project ID to which the source log topic belongs.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// The name of the log item to which the source log topic belongs.
        /// </summary>
        [Input("projectName")]
        public string? ProjectName { get; set; }

        /// <summary>
        /// Source log topic name.
        /// </summary>
        [Input("sourceTopicName")]
        public string? SourceTopicName { get; set; }

        /// <summary>
        /// Timed SQL analysis task status.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// Timed SQL analysis task ID.
        /// </summary>
        [Input("taskId")]
        public string? TaskId { get; set; }

        /// <summary>
        /// Timed SQL analysis task name.
        /// </summary>
        [Input("taskName")]
        public string? TaskName { get; set; }

        /// <summary>
        /// Source log topic ID.
        /// </summary>
        [Input("topicId")]
        public string? TopicId { get; set; }

        public ScheduleSqlTasksArgs()
        {
        }
        public static new ScheduleSqlTasksArgs Empty => new ScheduleSqlTasksArgs();
    }

    public sealed class ScheduleSqlTasksInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// IAM log project name.
        /// </summary>
        [Input("iamProjectName")]
        public Input<string>? IamProjectName { get; set; }

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
        /// The log project ID to which the source log topic belongs.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The name of the log item to which the source log topic belongs.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// Source log topic name.
        /// </summary>
        [Input("sourceTopicName")]
        public Input<string>? SourceTopicName { get; set; }

        /// <summary>
        /// Timed SQL analysis task status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Timed SQL analysis task ID.
        /// </summary>
        [Input("taskId")]
        public Input<string>? TaskId { get; set; }

        /// <summary>
        /// Timed SQL analysis task name.
        /// </summary>
        [Input("taskName")]
        public Input<string>? TaskName { get; set; }

        /// <summary>
        /// Source log topic ID.
        /// </summary>
        [Input("topicId")]
        public Input<string>? TopicId { get; set; }

        public ScheduleSqlTasksInvokeArgs()
        {
        }
        public static new ScheduleSqlTasksInvokeArgs Empty => new ScheduleSqlTasksInvokeArgs();
    }


    [OutputType]
    public sealed class ScheduleSqlTasksResult
    {
        public readonly string? IamProjectName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        public readonly string? ProjectId;
        public readonly string? ProjectName;
        /// <summary>
        /// The name of the source log topic where the original log for timed SQL analysis is located.
        /// </summary>
        public readonly string? SourceTopicName;
        /// <summary>
        /// Whether to start the scheduled SQL analysis task immediately after completing the task configuration.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Timed SQL analysis task ID.
        /// </summary>
        public readonly string? TaskId;
        /// <summary>
        /// Timed SQL analysis task name.
        /// </summary>
        public readonly string? TaskName;
        /// <summary>
        /// The List of timed SQL analysis tasks.
        /// </summary>
        public readonly ImmutableArray<Outputs.ScheduleSqlTasksTaskResult> Tasks;
        public readonly string? TopicId;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private ScheduleSqlTasksResult(
            string? iamProjectName,

            string id,

            string? nameRegex,

            string? outputFile,

            string? projectId,

            string? projectName,

            string? sourceTopicName,

            string? status,

            string? taskId,

            string? taskName,

            ImmutableArray<Outputs.ScheduleSqlTasksTaskResult> tasks,

            string? topicId,

            int totalCount)
        {
            IamProjectName = iamProjectName;
            Id = id;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            ProjectId = projectId;
            ProjectName = projectName;
            SourceTopicName = sourceTopicName;
            Status = status;
            TaskId = taskId;
            TaskName = taskName;
            Tasks = tasks;
            TopicId = topicId;
            TotalCount = totalCount;
        }
    }
}
