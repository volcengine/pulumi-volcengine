// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls.Outputs
{

    [OutputType]
    public sealed class GetAlarmsAlarmResult
    {
        /// <summary>
        /// The alarm id.
        /// </summary>
        public readonly string AlarmId;
        /// <summary>
        /// The alarm name.
        /// </summary>
        public readonly string AlarmName;
        /// <summary>
        /// List of notification groups corresponding to the alarm.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAlarmsAlarmAlarmNotifyGroupResult> AlarmNotifyGroups;
        /// <summary>
        /// Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
        /// </summary>
        public readonly int AlarmPeriod;
        /// <summary>
        /// Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAlarmsAlarmAlarmPeriodDetailResult> AlarmPeriodDetails;
        /// <summary>
        /// Alarm trigger condition.
        /// </summary>
        public readonly string Condition;
        /// <summary>
        /// The create time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The modify time.
        /// </summary>
        public readonly string ModifyTime;
        /// <summary>
        /// The project id.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// Search and analyze sentences, 1~3 can be configured.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAlarmsAlarmQueryRequestResult> QueryRequests;
        /// <summary>
        /// The execution period of the alarm task.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAlarmsAlarmRequestCycleResult> RequestCycles;
        /// <summary>
        /// The status.
        /// </summary>
        public readonly bool Status;
        /// <summary>
        /// Continuous cycle. The alarm will be issued after the trigger condition is continuously met for TriggerPeriod periods; the minimum value is 1, the maximum value is 10, and the default value is 1.
        /// </summary>
        public readonly int TriggerPeriod;
        /// <summary>
        /// Customize the alarm notification content.
        /// </summary>
        public readonly string UserDefineMsg;

        [OutputConstructor]
        private GetAlarmsAlarmResult(
            string alarmId,

            string alarmName,

            ImmutableArray<Outputs.GetAlarmsAlarmAlarmNotifyGroupResult> alarmNotifyGroups,

            int alarmPeriod,

            ImmutableArray<Outputs.GetAlarmsAlarmAlarmPeriodDetailResult> alarmPeriodDetails,

            string condition,

            string createTime,

            string modifyTime,

            string projectId,

            ImmutableArray<Outputs.GetAlarmsAlarmQueryRequestResult> queryRequests,

            ImmutableArray<Outputs.GetAlarmsAlarmRequestCycleResult> requestCycles,

            bool status,

            int triggerPeriod,

            string userDefineMsg)
        {
            AlarmId = alarmId;
            AlarmName = alarmName;
            AlarmNotifyGroups = alarmNotifyGroups;
            AlarmPeriod = alarmPeriod;
            AlarmPeriodDetails = alarmPeriodDetails;
            Condition = condition;
            CreateTime = createTime;
            ModifyTime = modifyTime;
            ProjectId = projectId;
            QueryRequests = queryRequests;
            RequestCycles = requestCycles;
            Status = status;
            TriggerPeriod = triggerPeriod;
            UserDefineMsg = userDefineMsg;
        }
    }
}
