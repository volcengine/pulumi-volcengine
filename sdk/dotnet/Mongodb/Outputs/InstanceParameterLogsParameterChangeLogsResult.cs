// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Mongodb.Outputs
{

    [OutputType]
    public sealed class InstanceParameterLogsParameterChangeLogsResult
    {
        /// <summary>
        /// The modifying time of parameter.
        /// </summary>
        public readonly string ModifyTime;
        /// <summary>
        /// The new parameter value.
        /// </summary>
        public readonly string NewParameterValue;
        /// <summary>
        /// The old parameter value.
        /// </summary>
        public readonly string OldParameterValue;
        /// <summary>
        /// The parameter name.
        /// </summary>
        public readonly string ParameterName;
        /// <summary>
        /// The node type to which the parameter belongs.
        /// </summary>
        public readonly string ParameterRole;
        /// <summary>
        /// The status of parameter change.
        /// </summary>
        public readonly string ParameterStatus;

        [OutputConstructor]
        private InstanceParameterLogsParameterChangeLogsResult(
            string modifyTime,

            string newParameterValue,

            string oldParameterValue,

            string parameterName,

            string parameterRole,

            string parameterStatus)
        {
            ModifyTime = modifyTime;
            NewParameterValue = newParameterValue;
            OldParameterValue = oldParameterValue;
            ParameterName = parameterName;
            ParameterRole = parameterRole;
            ParameterStatus = parameterStatus;
        }
    }
}