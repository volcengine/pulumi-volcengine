// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rabbitmq.Outputs
{

    [OutputType]
    public sealed class GetInstancesRabbitmqInstanceChargeDetailResult
    {
        /// <summary>
        /// Whether to automatically renew in prepaid scenarios.
        /// </summary>
        public readonly bool AutoRenew;
        /// <summary>
        /// The charge end time of the rabbitmq instance.
        /// </summary>
        public readonly string ChargeEndTime;
        /// <summary>
        /// The charge expire time of the rabbitmq instance.
        /// </summary>
        public readonly string ChargeExpireTime;
        /// <summary>
        /// The charge start time of the rabbitmq instance.
        /// </summary>
        public readonly string ChargeStartTime;
        /// <summary>
        /// The charge status of the rabbitmq instance.
        /// </summary>
        public readonly string ChargeStatus;
        /// <summary>
        /// The charge type of rabbitmq instance.
        /// </summary>
        public readonly string ChargeType;
        /// <summary>
        /// The overdue reclaim time of the rabbitmq instance.
        /// </summary>
        public readonly string OverdueReclaimTime;
        /// <summary>
        /// The overdue time of the rabbitmq instance.
        /// </summary>
        public readonly string OverdueTime;

        [OutputConstructor]
        private GetInstancesRabbitmqInstanceChargeDetailResult(
            bool autoRenew,

            string chargeEndTime,

            string chargeExpireTime,

            string chargeStartTime,

            string chargeStatus,

            string chargeType,

            string overdueReclaimTime,

            string overdueTime)
        {
            AutoRenew = autoRenew;
            ChargeEndTime = chargeEndTime;
            ChargeExpireTime = chargeExpireTime;
            ChargeStartTime = chargeStartTime;
            ChargeStatus = chargeStatus;
            ChargeType = chargeType;
            OverdueReclaimTime = overdueReclaimTime;
            OverdueTime = overdueTime;
        }
    }
}
