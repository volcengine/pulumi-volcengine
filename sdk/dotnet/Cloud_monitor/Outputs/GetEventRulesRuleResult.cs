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
    public sealed class GetEventRulesRuleResult
    {
        /// <summary>
        /// The id of the account.
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// When the alarm notification method is phone, SMS, or email, the triggered alarm contact group ID.
        /// </summary>
        public readonly ImmutableArray<string> ContactGroupIds;
        /// <summary>
        /// List of contact methods.
        /// </summary>
        public readonly ImmutableArray<string> ContactMethods;
        /// <summary>
        /// The create time.
        /// </summary>
        public readonly int CreatedAt;
        /// <summary>
        /// The description of the rule.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The end time of the rule.
        /// </summary>
        public readonly string EffectEndAt;
        /// <summary>
        /// The start time of the rule.
        /// </summary>
        public readonly string EffectStartAt;
        /// <summary>
        /// When the alarm notification method is alarm callback, it triggers the callback address.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// The name of the event bus.
        /// </summary>
        public readonly string EventBusName;
        /// <summary>
        /// The source of the event.
        /// </summary>
        public readonly string EventSource;
        /// <summary>
        /// The event type.
        /// </summary>
        public readonly ImmutableArray<string> EventTypes;
        /// <summary>
        /// Filter mode, also known as event matching rules. Custom matching rules are not currently supported.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetEventRulesRuleFilterPatternResult> FilterPatterns;
        /// <summary>
        /// The id of the rule.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The level of the rule.
        /// </summary>
        public readonly string Level;
        /// <summary>
        /// The triggered message queue when the alarm notification method is Kafka message queue.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetEventRulesRuleMessageQueueResult> MessageQueues;
        /// <summary>
        /// The name of the region.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The id of the rule.
        /// </summary>
        public readonly string RuleId;
        /// <summary>
        /// Rule name, search rules by name using fuzzy search.
        /// </summary>
        public readonly string RuleName;
        /// <summary>
        /// Enable the state of the rule.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The alarm method for log service triggers the configuration of the log service.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetEventRulesRuleTlsTargetResult> TlsTargets;
        /// <summary>
        /// The updated time.
        /// </summary>
        public readonly int UpdatedAt;

        [OutputConstructor]
        private GetEventRulesRuleResult(
            string accountId,

            ImmutableArray<string> contactGroupIds,

            ImmutableArray<string> contactMethods,

            int createdAt,

            string description,

            string effectEndAt,

            string effectStartAt,

            string endpoint,

            string eventBusName,

            string eventSource,

            ImmutableArray<string> eventTypes,

            ImmutableArray<Outputs.GetEventRulesRuleFilterPatternResult> filterPatterns,

            string id,

            string level,

            ImmutableArray<Outputs.GetEventRulesRuleMessageQueueResult> messageQueues,

            string region,

            string ruleId,

            string ruleName,

            string status,

            ImmutableArray<Outputs.GetEventRulesRuleTlsTargetResult> tlsTargets,

            int updatedAt)
        {
            AccountId = accountId;
            ContactGroupIds = contactGroupIds;
            ContactMethods = contactMethods;
            CreatedAt = createdAt;
            Description = description;
            EffectEndAt = effectEndAt;
            EffectStartAt = effectStartAt;
            Endpoint = endpoint;
            EventBusName = eventBusName;
            EventSource = eventSource;
            EventTypes = eventTypes;
            FilterPatterns = filterPatterns;
            Id = id;
            Level = level;
            MessageQueues = messageQueues;
            Region = region;
            RuleId = ruleId;
            RuleName = ruleName;
            Status = status;
            TlsTargets = tlsTargets;
            UpdatedAt = updatedAt;
        }
    }
}
