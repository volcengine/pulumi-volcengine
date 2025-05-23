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
    public sealed class GetInvocationResultsInvocationResultResult
    {
        /// <summary>
        /// The id of ecs command.
        /// </summary>
        public readonly string CommandId;
        /// <summary>
        /// The end time of the ecs invocation in the instance.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// The error code of the ecs invocation.
        /// </summary>
        public readonly string ErrorCode;
        /// <summary>
        /// The error message of the ecs invocation.
        /// </summary>
        public readonly string ErrorMessage;
        /// <summary>
        /// The exit code of the ecs command.
        /// </summary>
        public readonly int ExitCode;
        /// <summary>
        /// The id of the ecs invocation result.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The id of ecs instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The id of ecs invocation.
        /// </summary>
        public readonly string InvocationId;
        /// <summary>
        /// The id of the ecs invocation result.
        /// </summary>
        public readonly string InvocationResultId;
        /// <summary>
        /// The list of status of ecs invocation in a single instance. Valid values: `Pending`, `Running`, `Success`, `Failed`, `Timeout`.
        /// </summary>
        public readonly string InvocationResultStatus;
        /// <summary>
        /// The base64 encoded output message of the ecs invocation.
        /// </summary>
        public readonly string Output;
        /// <summary>
        /// The start time of the ecs invocation in the instance.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// The username of the ecs command.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GetInvocationResultsInvocationResultResult(
            string commandId,

            string endTime,

            string errorCode,

            string errorMessage,

            int exitCode,

            string id,

            string instanceId,

            string invocationId,

            string invocationResultId,

            string invocationResultStatus,

            string output,

            string startTime,

            string username)
        {
            CommandId = commandId;
            EndTime = endTime;
            ErrorCode = errorCode;
            ErrorMessage = errorMessage;
            ExitCode = exitCode;
            Id = id;
            InstanceId = instanceId;
            InvocationId = invocationId;
            InvocationResultId = invocationResultId;
            InvocationResultStatus = invocationResultStatus;
            Output = output;
            StartTime = startTime;
            Username = username;
        }
    }
}
