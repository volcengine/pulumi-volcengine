// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls.Inputs
{

    public sealed class ImportTaskTargetInfoGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Log extraction rules.
        /// </summary>
        [Input("extractRule")]
        public Input<Inputs.ImportTaskTargetInfoExtractRuleGetArgs>? ExtractRule { get; set; }

        /// <summary>
        /// Log sample.
        /// </summary>
        [Input("logSample")]
        public Input<string>? LogSample { get; set; }

        /// <summary>
        /// Specify the log parsing type when importing.
        /// </summary>
        [Input("logType", required: true)]
        public Input<string> LogType { get; set; } = null!;

        /// <summary>
        /// Regional ID.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public ImportTaskTargetInfoGetArgs()
        {
        }
        public static new ImportTaskTargetInfoGetArgs Empty => new ImportTaskTargetInfoGetArgs();
    }
}
