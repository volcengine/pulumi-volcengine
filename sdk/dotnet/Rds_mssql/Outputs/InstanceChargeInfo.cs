// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_mssql.Outputs
{

    [OutputType]
    public sealed class InstanceChargeInfo
    {
        /// <summary>
        /// Whether to enable automatic renewal in the prepaid scenario. This parameter can be set when the ChargeType is `Prepaid`.
        /// </summary>
        public readonly bool? AutoRenew;
        /// <summary>
        /// Charge end time.
        /// </summary>
        public readonly string? ChargeEndTime;
        /// <summary>
        /// Charge start time.
        /// </summary>
        public readonly string? ChargeStartTime;
        /// <summary>
        /// The charge status.
        /// </summary>
        public readonly string? ChargeStatus;
        /// <summary>
        /// The charge type. Valid values: `PostPaid`, `PrePaid`.
        /// </summary>
        public readonly string ChargeType;
        /// <summary>
        /// Expected release time when overdue fees are shut down.
        /// </summary>
        public readonly string? OverdueReclaimTime;
        /// <summary>
        /// Time for Disconnection due to Unpaid Fees.
        /// </summary>
        public readonly string? OverdueTime;
        /// <summary>
        /// Purchase duration in a prepaid scenario. This parameter is required when the ChargeType is `Prepaid`.
        /// </summary>
        public readonly int? Period;

        [OutputConstructor]
        private InstanceChargeInfo(
            bool? autoRenew,

            string? chargeEndTime,

            string? chargeStartTime,

            string? chargeStatus,

            string chargeType,

            string? overdueReclaimTime,

            string? overdueTime,

            int? period)
        {
            AutoRenew = autoRenew;
            ChargeEndTime = chargeEndTime;
            ChargeStartTime = chargeStartTime;
            ChargeStatus = chargeStatus;
            ChargeType = chargeType;
            OverdueReclaimTime = overdueReclaimTime;
            OverdueTime = overdueTime;
            Period = period;
        }
    }
}
