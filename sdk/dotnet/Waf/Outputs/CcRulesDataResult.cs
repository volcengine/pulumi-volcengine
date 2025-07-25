// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Waf.Outputs
{

    [OutputType]
    public sealed class CcRulesDataResult
    {
        /// <summary>
        /// The total number of enabled rules within the rule group.
        /// </summary>
        public readonly int EnableCount;
        /// <summary>
        /// The creation time of the rule group.
        /// </summary>
        public readonly string InsertTime;
        /// <summary>
        /// Details of the rule group.
        /// </summary>
        public readonly ImmutableArray<Outputs.CcRulesDataRuleGroupResult> RuleGroups;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// Fuzzy search by the requested path.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private CcRulesDataResult(
            int enableCount,

            string insertTime,

            ImmutableArray<Outputs.CcRulesDataRuleGroupResult> ruleGroups,

            int totalCount,

            string url)
        {
            EnableCount = enableCount;
            InsertTime = insertTime;
            RuleGroups = ruleGroups;
            TotalCount = totalCount;
            Url = url;
        }
    }
}
