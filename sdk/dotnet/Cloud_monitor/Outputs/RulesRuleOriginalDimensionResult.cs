// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cloud_monitor.Outputs
{

    [OutputType]
    public sealed class RulesRuleOriginalDimensionResult
    {
        /// <summary>
        /// The key of the dimension.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The value of the dimension.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private RulesRuleOriginalDimensionResult(
            string key,

            ImmutableArray<string> values)
        {
            Key = key;
            Values = values;
        }
    }
}
