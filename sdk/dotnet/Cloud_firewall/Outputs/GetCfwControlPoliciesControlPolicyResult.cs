// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cloud_firewall.Outputs
{

    [OutputType]
    public sealed class GetCfwControlPoliciesControlPolicyResult
    {
        /// <summary>
        /// The account id of the control policy.
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// The action list of the control policy. Valid values: `accept`, `deny`, `monitor`.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// The description of the control policy. This field support fuzzy query.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The dest port of the control policy.
        /// </summary>
        public readonly string DestPort;
        /// <summary>
        /// The dest port group type of the control policy.
        /// </summary>
        public readonly string DestPortGroupType;
        /// <summary>
        /// The dest port list of the control policy.
        /// </summary>
        public readonly ImmutableArray<string> DestPortLists;
        /// <summary>
        /// The dest port type of the control policy.
        /// </summary>
        public readonly string DestPortType;
        /// <summary>
        /// The destination of the control policy. This field support fuzzy query.
        /// </summary>
        public readonly string Destination;
        /// <summary>
        /// The destination cidr list of the control policy.
        /// </summary>
        public readonly ImmutableArray<string> DestinationCidrLists;
        /// <summary>
        /// The destination group type of the control policy.
        /// </summary>
        public readonly string DestinationGroupType;
        /// <summary>
        /// The destination type of the control policy.
        /// </summary>
        public readonly string DestinationType;
        /// <summary>
        /// The direction of control policy. Valid values: `in`, `out`.
        /// </summary>
        public readonly string Direction;
        /// <summary>
        /// The effect status of the control policy. 1: Not yet effective, 2: Issued in progress, 3: Effective.
        /// </summary>
        public readonly int EffectStatus;
        /// <summary>
        /// The end time of the control policy. Unix timestamp.
        /// </summary>
        public readonly int EndTime;
        /// <summary>
        /// The hit count of the control policy.
        /// </summary>
        public readonly int HitCnt;
        /// <summary>
        /// The id of the control policy.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Whether the control policy is effected.
        /// </summary>
        public readonly bool IsEffected;
        /// <summary>
        /// The priority of the control policy.
        /// </summary>
        public readonly int Prio;
        /// <summary>
        /// The proto list of the control policy. Valid values: `TCP`, `ICMP`, `UDP`, `ANY`. When the destination_type is `domain`, The proto must be `TCP`.
        /// </summary>
        public readonly string Proto;
        /// <summary>
        /// The repeat days of the control policy.
        /// </summary>
        public readonly ImmutableArray<int> RepeatDays;
        /// <summary>
        /// The repeat end time of the control policy.
        /// </summary>
        public readonly string RepeatEndTime;
        /// <summary>
        /// The repeat start time of the control policy.
        /// </summary>
        public readonly string RepeatStartTime;
        /// <summary>
        /// The repeat type of the control policy. Valid values: `Permanent`, `Once`, `Daily`, `Weekly`, `Monthly`.
        /// </summary>
        public readonly string RepeatType;
        /// <summary>
        /// The rule id of the control policy. This field support fuzzy query.
        /// </summary>
        public readonly string RuleId;
        /// <summary>
        /// The source of the control policy. This field support fuzzy query.
        /// </summary>
        public readonly string Source;
        /// <summary>
        /// The source cidr list of the control policy.
        /// </summary>
        public readonly ImmutableArray<string> SourceCidrLists;
        /// <summary>
        /// The source group type of the control policy.
        /// </summary>
        public readonly string SourceGroupType;
        /// <summary>
        /// The source type of the control policy.
        /// </summary>
        public readonly string SourceType;
        /// <summary>
        /// The start time of the control policy. Unix timestamp.
        /// </summary>
        public readonly int StartTime;
        /// <summary>
        /// The enable status list of the control policy.
        /// </summary>
        public readonly bool Status;
        /// <summary>
        /// The update time of the control policy.
        /// </summary>
        public readonly int UpdateTime;
        /// <summary>
        /// The use count of the control policy.
        /// </summary>
        public readonly int UseCount;

        [OutputConstructor]
        private GetCfwControlPoliciesControlPolicyResult(
            string accountId,

            string action,

            string description,

            string destPort,

            string destPortGroupType,

            ImmutableArray<string> destPortLists,

            string destPortType,

            string destination,

            ImmutableArray<string> destinationCidrLists,

            string destinationGroupType,

            string destinationType,

            string direction,

            int effectStatus,

            int endTime,

            int hitCnt,

            string id,

            bool isEffected,

            int prio,

            string proto,

            ImmutableArray<int> repeatDays,

            string repeatEndTime,

            string repeatStartTime,

            string repeatType,

            string ruleId,

            string source,

            ImmutableArray<string> sourceCidrLists,

            string sourceGroupType,

            string sourceType,

            int startTime,

            bool status,

            int updateTime,

            int useCount)
        {
            AccountId = accountId;
            Action = action;
            Description = description;
            DestPort = destPort;
            DestPortGroupType = destPortGroupType;
            DestPortLists = destPortLists;
            DestPortType = destPortType;
            Destination = destination;
            DestinationCidrLists = destinationCidrLists;
            DestinationGroupType = destinationGroupType;
            DestinationType = destinationType;
            Direction = direction;
            EffectStatus = effectStatus;
            EndTime = endTime;
            HitCnt = hitCnt;
            Id = id;
            IsEffected = isEffected;
            Prio = prio;
            Proto = proto;
            RepeatDays = repeatDays;
            RepeatEndTime = repeatEndTime;
            RepeatStartTime = repeatStartTime;
            RepeatType = repeatType;
            RuleId = ruleId;
            Source = source;
            SourceCidrLists = sourceCidrLists;
            SourceGroupType = sourceGroupType;
            SourceType = sourceType;
            StartTime = startTime;
            Status = status;
            UpdateTime = updateTime;
            UseCount = useCount;
        }
    }
}
