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
    public sealed class InstanceParametersInstanceParameterResult
    {
        /// <summary>
        /// The checking code of parameter.
        /// </summary>
        public readonly string CheckingCode;
        /// <summary>
        /// Whether the parameter supports modifying.
        /// </summary>
        public readonly bool ForceModify;
        /// <summary>
        /// Does the new parameter value need to restart the instance to take effect after modification.
        /// </summary>
        public readonly bool ForceRestart;
        /// <summary>
        /// The default value of parameter.
        /// </summary>
        public readonly string ParameterDefaultValue;
        /// <summary>
        /// The description of parameter.
        /// </summary>
        public readonly string ParameterDescription;
        /// <summary>
        /// The name of parameter.
        /// </summary>
        public readonly string ParameterName;
        /// <summary>
        /// The node type of instance parameter, valid value contains `Node`, `Shard`, `ConfigServer`, `Mongos`.
        /// </summary>
        public readonly string ParameterRole;
        /// <summary>
        /// The type of parameter value.
        /// </summary>
        public readonly string ParameterType;
        /// <summary>
        /// The value of parameter.
        /// </summary>
        public readonly string ParameterValue;

        [OutputConstructor]
        private InstanceParametersInstanceParameterResult(
            string checkingCode,

            bool forceModify,

            bool forceRestart,

            string parameterDefaultValue,

            string parameterDescription,

            string parameterName,

            string parameterRole,

            string parameterType,

            string parameterValue)
        {
            CheckingCode = checkingCode;
            ForceModify = forceModify;
            ForceRestart = forceRestart;
            ParameterDefaultValue = parameterDefaultValue;
            ParameterDescription = parameterDescription;
            ParameterName = parameterName;
            ParameterRole = parameterRole;
            ParameterType = parameterType;
            ParameterValue = parameterValue;
        }
    }
}