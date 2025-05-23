// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_mysql.Outputs
{

    [OutputType]
    public sealed class ParameterTemplateTemplateParam
    {
        /// <summary>
        /// Instance parameter name.
        /// Description: When using CreateParameterTemplate and ModifyParameterTemplate as request parameters, only Name and RunningValue need to be passed in.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Parameter running value.
        /// Description: When making request parameters in CreateParameterTemplate and ModifyParameterTemplate, only Name and RunningValue need to be passed in.
        /// </summary>
        public readonly string RunningValue;

        [OutputConstructor]
        private ParameterTemplateTemplateParam(
            string name,

            string runningValue)
        {
            Name = name;
            RunningValue = runningValue;
        }
    }
}
