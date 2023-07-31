// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Clb.Outputs
{

    [OutputType]
    public sealed class RulesRuleResult
    {
        /// <summary>
        /// The Description of Rule.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The Domain of Rule.
        /// </summary>
        public readonly string Domain;
        /// <summary>
        /// The Id of Rule.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Id of Rule.
        /// </summary>
        public readonly string RuleId;
        /// <summary>
        /// The Id of Server Group.
        /// </summary>
        public readonly string ServerGroupId;
        /// <summary>
        /// The Url of Rule.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private RulesRuleResult(
            string description,

            string domain,

            string id,

            string ruleId,

            string serverGroupId,

            string url)
        {
            Description = description;
            Domain = domain;
            Id = id;
            RuleId = ruleId;
            ServerGroupId = serverGroupId;
            Url = url;
        }
    }
}