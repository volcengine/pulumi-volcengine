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
    public sealed class BotAnalyseProtectRuleRuleGroup
    {
        /// <summary>
        /// Rule group information.
        /// </summary>
        public readonly Outputs.BotAnalyseProtectRuleRuleGroupGroup? Group;
        /// <summary>
        /// Specific rule information within the rule group.
        /// </summary>
        public readonly ImmutableArray<Outputs.BotAnalyseProtectRuleRuleGroupRule> Rules;

        [OutputConstructor]
        private BotAnalyseProtectRuleRuleGroup(
            Outputs.BotAnalyseProtectRuleRuleGroupGroup? group,

            ImmutableArray<Outputs.BotAnalyseProtectRuleRuleGroupRule> rules)
        {
            Group = group;
            Rules = rules;
        }
    }
}
