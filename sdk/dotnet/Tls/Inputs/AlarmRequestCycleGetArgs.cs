// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls.Inputs
{

    public sealed class AlarmRequestCycleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cycle of alarm task execution, or the time point of periodic execution. The unit is minutes, and the value range is 1~1440.
        /// </summary>
        [Input("time", required: true)]
        public Input<int> Time { get; set; } = null!;

        /// <summary>
        /// Execution cycle type.
        /// Period: Periodic execution, which means executing once every certain period of time.
        /// Fixed: Regular execution, which means executing at a fixed time point every day.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public AlarmRequestCycleGetArgs()
        {
        }
        public static new AlarmRequestCycleGetArgs Empty => new AlarmRequestCycleGetArgs();
    }
}