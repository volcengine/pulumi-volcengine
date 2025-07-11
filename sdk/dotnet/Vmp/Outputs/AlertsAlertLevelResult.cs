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
    public sealed class AlertsAlertLevelResult
    {
        /// <summary>
        /// The comparator of the vmp alerting rule.
        /// </summary>
        public readonly string Comparator;
        /// <summary>
        /// The duration of the alerting rule.
        /// </summary>
        public readonly string For;
        /// <summary>
        /// The level of vmp alert. Valid values: `P0`, `P1`, `P2`.
        /// </summary>
        public readonly string Level;
        /// <summary>
        /// The threshold of the vmp alerting rule.
        /// </summary>
        public readonly double Threshold;

        [OutputConstructor]
        private AlertsAlertLevelResult(
            string comparator,

            string @for,

            string level,

            double threshold)
        {
            Comparator = comparator;
            For = @for;
            Level = level;
            Threshold = threshold;
        }
    }
}
