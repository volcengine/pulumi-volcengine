// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_postgresql.Inputs
{

    public sealed class InstanceChargeDetailGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to automatically renew in prepaid scenarios.
        /// Autorenew_Enable
        /// Autorenew_Disable (default).
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// Billing expiry time (yearly and monthly only).
        /// </summary>
        [Input("chargeEndTime")]
        public Input<string>? ChargeEndTime { get; set; }

        /// <summary>
        /// Billing start time (pay-as-you-go &amp; monthly subscription).
        /// </summary>
        [Input("chargeStartTime")]
        public Input<string>? ChargeStartTime { get; set; }

        /// <summary>
        /// Pay status. Value:
        /// normal - normal
        /// overdue - overdue
        /// .
        /// </summary>
        [Input("chargeStatus")]
        public Input<string>? ChargeStatus { get; set; }

        /// <summary>
        /// Payment type. Value:
        /// PostPaid - Pay-As-You-Go
        /// PrePaid - Yearly and monthly (default).
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// Estimated release time when arrears are closed (pay-as-you-go &amp; monthly subscription).
        /// </summary>
        [Input("overdueReclaimTime")]
        public Input<string>? OverdueReclaimTime { get; set; }

        /// <summary>
        /// Shutdown time in arrears (pay-as-you-go &amp; monthly subscription).
        /// </summary>
        [Input("overdueTime")]
        public Input<string>? OverdueTime { get; set; }

        /// <summary>
        /// Purchase duration in prepaid scenarios. Default: 1.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The purchase cycle in the prepaid scenario.
        /// Month - monthly subscription (default)
        /// Year - Package year.
        /// </summary>
        [Input("periodUnit")]
        public Input<string>? PeriodUnit { get; set; }

        /// <summary>
        /// Temporary upgrade of restoration time.
        /// </summary>
        [Input("tempModifyEndTime")]
        public Input<string>? TempModifyEndTime { get; set; }

        /// <summary>
        /// Start time of temporary upgrade.
        /// </summary>
        [Input("tempModifyStartTime")]
        public Input<string>? TempModifyStartTime { get; set; }

        public InstanceChargeDetailGetArgs()
        {
        }
        public static new InstanceChargeDetailGetArgs Empty => new InstanceChargeDetailGetArgs();
    }
}
