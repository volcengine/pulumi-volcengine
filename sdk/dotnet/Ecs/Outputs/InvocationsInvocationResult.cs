// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Ecs.Outputs
{

    [OutputType]
    public sealed class InvocationsInvocationResult
    {
        /// <summary>
        /// The base64 encoded content of the ecs command.
        /// </summary>
        public readonly string CommandContent;
        /// <summary>
        /// The description of the ecs command.
        /// </summary>
        public readonly string CommandDescription;
        /// <summary>
        /// The id of ecs command.
        /// </summary>
        public readonly string CommandId;
        /// <summary>
        /// The name of ecs command. This field support fuzzy query.
        /// </summary>
        public readonly string CommandName;
        /// <summary>
        /// The provider of the ecs command.
        /// </summary>
        public readonly string CommandProvider;
        /// <summary>
        /// The type of ecs command. Valid values: `Shell`.
        /// </summary>
        public readonly string CommandType;
        /// <summary>
        /// The end time of the ecs invocation.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// The frequency of the ecs invocation.
        /// </summary>
        public readonly string Frequency;
        /// <summary>
        /// The id of the ecs invocation.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of ECS instance IDs.
        /// </summary>
        public readonly ImmutableArray<string> InstanceIds;
        /// <summary>
        /// The instance number of the ecs invocation.
        /// </summary>
        public readonly int InstanceNumber;
        /// <summary>
        /// The description of the ecs invocation.
        /// </summary>
        public readonly string InvocationDescription;
        /// <summary>
        /// The id of ecs invocation.
        /// </summary>
        public readonly string InvocationId;
        /// <summary>
        /// The name of ecs invocation. This field support fuzzy query.
        /// </summary>
        public readonly string InvocationName;
        /// <summary>
        /// The list of status of ecs invocation. Valid values: `Pending`, `Scheduled`, `Running`, `Success`, `Failed`, `Stopped`, `PartialFailed`, `Finished`.
        /// </summary>
        public readonly string InvocationStatus;
        /// <summary>
        /// The launch time of the ecs invocation.
        /// </summary>
        public readonly string LaunchTime;
        /// <summary>
        /// The recurrence end time of the ecs invocation.
        /// </summary>
        public readonly string RecurrenceEndTime;
        /// <summary>
        /// The repeat mode of ecs invocation. Valid values: `Once`, `Rate`, `Fixed`.
        /// </summary>
        public readonly string RepeatMode;
        /// <summary>
        /// The start time of the ecs invocation.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// The timeout of the ecs command.
        /// </summary>
        public readonly int Timeout;
        /// <summary>
        /// The username of the ecs command.
        /// </summary>
        public readonly string Username;
        /// <summary>
        /// The working directory of the ecs command.
        /// </summary>
        public readonly string WorkingDir;

        [OutputConstructor]
        private InvocationsInvocationResult(
            string commandContent,

            string commandDescription,

            string commandId,

            string commandName,

            string commandProvider,

            string commandType,

            string endTime,

            string frequency,

            string id,

            ImmutableArray<string> instanceIds,

            int instanceNumber,

            string invocationDescription,

            string invocationId,

            string invocationName,

            string invocationStatus,

            string launchTime,

            string recurrenceEndTime,

            string repeatMode,

            string startTime,

            int timeout,

            string username,

            string workingDir)
        {
            CommandContent = commandContent;
            CommandDescription = commandDescription;
            CommandId = commandId;
            CommandName = commandName;
            CommandProvider = commandProvider;
            CommandType = commandType;
            EndTime = endTime;
            Frequency = frequency;
            Id = id;
            InstanceIds = instanceIds;
            InstanceNumber = instanceNumber;
            InvocationDescription = invocationDescription;
            InvocationId = invocationId;
            InvocationName = invocationName;
            InvocationStatus = invocationStatus;
            LaunchTime = launchTime;
            RecurrenceEndTime = recurrenceEndTime;
            RepeatMode = repeatMode;
            StartTime = startTime;
            Timeout = timeout;
            Username = username;
            WorkingDir = workingDir;
        }
    }
}