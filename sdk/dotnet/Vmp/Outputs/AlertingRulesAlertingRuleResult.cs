// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vmp.Outputs
{

    [OutputType]
    public sealed class AlertingRulesAlertingRuleResult
    {
        /// <summary>
        /// The annotations of the vmp alerting rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.AlertingRulesAlertingRuleAnnotationResult> Annotations;
        /// <summary>
        /// The create time of the vmp alerting rule.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The description of the vmp alerting rule.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The id of the vmp alerting rule.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The labels of the vmp alerting rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.AlertingRulesAlertingRuleLabelResult> Labels;
        /// <summary>
        /// The alerting levels of the vmp alerting rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.AlertingRulesAlertingRuleLevelResult> Levels;
        /// <summary>
        /// The name of vmp alerting rule. This field support fuzzy query.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The notify group policy id of the vmp alerting rule.
        /// </summary>
        public readonly string NotifyGroupPolicyId;
        /// <summary>
        /// The notify policy id of the vmp alerting rule.
        /// </summary>
        public readonly string NotifyPolicyId;
        /// <summary>
        /// The alerting query of the vmp alerting rule.
        /// </summary>
        public readonly ImmutableArray<Outputs.AlertingRulesAlertingRuleQueryResult> Queries;
        /// <summary>
        /// The status of vmp alerting rule. Valid values: `Running`, `Disabled`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The type of vmp alerting rule. Valid values: `vmp/PromQL`.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The update time of the vmp alerting rule.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private AlertingRulesAlertingRuleResult(
            ImmutableArray<Outputs.AlertingRulesAlertingRuleAnnotationResult> annotations,

            string createTime,

            string description,

            string id,

            ImmutableArray<Outputs.AlertingRulesAlertingRuleLabelResult> labels,

            ImmutableArray<Outputs.AlertingRulesAlertingRuleLevelResult> levels,

            string name,

            string notifyGroupPolicyId,

            string notifyPolicyId,

            ImmutableArray<Outputs.AlertingRulesAlertingRuleQueryResult> queries,

            string status,

            string type,

            string updateTime)
        {
            Annotations = annotations;
            CreateTime = createTime;
            Description = description;
            Id = id;
            Labels = labels;
            Levels = levels;
            Name = name;
            NotifyGroupPolicyId = notifyGroupPolicyId;
            NotifyPolicyId = notifyPolicyId;
            Queries = queries;
            Status = status;
            Type = type;
            UpdateTime = updateTime;
        }
    }
}
