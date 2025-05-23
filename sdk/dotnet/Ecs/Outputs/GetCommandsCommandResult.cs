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
    public sealed class GetCommandsCommandResult
    {
        /// <summary>
        /// The base64 encoded content of the ecs command.
        /// </summary>
        public readonly string CommandContent;
        /// <summary>
        /// The id of ecs command.
        /// </summary>
        public readonly string CommandId;
        /// <summary>
        /// The provider of public command. When this field is not specified, query for custom commands.
        /// </summary>
        public readonly string CommandProvider;
        /// <summary>
        /// The create time of the ecs command.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The description of the ecs command.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The id of the ecs command.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The invocation times of the ecs command. Public commands do not display the invocation times.
        /// </summary>
        public readonly int InvocationTimes;
        /// <summary>
        /// The name of ecs command. This field support fuzzy query.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The timeout of the ecs command.
        /// </summary>
        public readonly int Timeout;
        /// <summary>
        /// The type of ecs command. Valid values: `Shell`.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The update time of the ecs command.
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// The username of the ecs command.
        /// </summary>
        public readonly string Username;
        /// <summary>
        /// The working directory of the ecs command.
        /// </summary>
        public readonly string WorkingDir;

        [OutputConstructor]
        private GetCommandsCommandResult(
            string commandContent,

            string commandId,

            string commandProvider,

            string createdAt,

            string description,

            string id,

            int invocationTimes,

            string name,

            int timeout,

            string type,

            string updatedAt,

            string username,

            string workingDir)
        {
            CommandContent = commandContent;
            CommandId = commandId;
            CommandProvider = commandProvider;
            CreatedAt = createdAt;
            Description = description;
            Id = id;
            InvocationTimes = invocationTimes;
            Name = name;
            Timeout = timeout;
            Type = type;
            UpdatedAt = updatedAt;
            Username = username;
            WorkingDir = workingDir;
        }
    }
}
