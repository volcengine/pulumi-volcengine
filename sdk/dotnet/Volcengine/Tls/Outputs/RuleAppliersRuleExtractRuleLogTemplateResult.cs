// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Tls.Outputs
{

    [OutputType]
    public sealed class RuleAppliersRuleExtractRuleLogTemplateResult
    {
        /// <summary>
        /// Log template content.
        /// </summary>
        public readonly string Format;
        /// <summary>
        /// The type of the log template.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private RuleAppliersRuleExtractRuleLogTemplateResult(
            string format,

            string type)
        {
            Format = format;
            Type = type;
        }
    }
}