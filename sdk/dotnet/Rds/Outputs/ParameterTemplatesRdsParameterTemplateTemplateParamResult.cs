// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds.Outputs
{

    [OutputType]
    public sealed class ParameterTemplatesRdsParameterTemplateTemplateParamResult
    {
        /// <summary>
        /// Parameter default value.
        /// </summary>
        public readonly string DefaultValue;
        /// <summary>
        /// Parameter description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Parameter name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Whether the modified parameters need to be restarted to take effect.
        /// </summary>
        public readonly bool Restart;
        /// <summary>
        /// Parameter running value.
        /// </summary>
        public readonly string RunningValue;
        /// <summary>
        /// Parameter value range.
        /// </summary>
        public readonly string ValueRange;

        [OutputConstructor]
        private ParameterTemplatesRdsParameterTemplateTemplateParamResult(
            string defaultValue,

            string description,

            string name,

            bool restart,

            string runningValue,

            string valueRange)
        {
            DefaultValue = defaultValue;
            Description = description;
            Name = name;
            Restart = restart;
            RunningValue = runningValue;
            ValueRange = valueRange;
        }
    }
}