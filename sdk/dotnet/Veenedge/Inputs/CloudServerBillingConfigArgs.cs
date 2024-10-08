// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Veenedge.Inputs
{

    public sealed class CloudServerBillingConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The method of bandwidth billing. The value can be `MonthlyP95` or `DailyPeak`.
        /// </summary>
        [Input("bandwidthBillingMethod", required: true)]
        public Input<string> BandwidthBillingMethod { get; set; } = null!;

        /// <summary>
        /// The method of computing billing. The value can be `MonthlyPeak` or `DailyPeak`.
        /// </summary>
        [Input("computingBillingMethod", required: true)]
        public Input<string> ComputingBillingMethod { get; set; } = null!;

        public CloudServerBillingConfigArgs()
        {
        }
        public static new CloudServerBillingConfigArgs Empty => new CloudServerBillingConfigArgs();
    }
}
