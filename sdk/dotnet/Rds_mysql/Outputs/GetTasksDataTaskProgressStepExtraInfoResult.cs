// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_mysql.Outputs
{

    [OutputType]
    public sealed class GetTasksDataTaskProgressStepExtraInfoResult
    {
        /// <summary>
        /// Current stage. CostTime: The time taken for the current stage.
        /// CurDataSize: The amount of data imported currently.
        /// CurBinlog: The number of Binlog files being replayed currently.
        /// RemainCostTime: The remaining time taken.
        /// RemainDataSize: The remaining amount of data to be imported. RemainBinlog: The number of Binlog files remaining for playback.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Unit. Values:
        /// MS: Milliseconds.
        /// Bytes: Bytes.
        /// Files: Number of (files).
        /// </summary>
        public readonly string Unit;
        /// <summary>
        /// The specific value corresponding to the Type field.
        /// </summary>
        public readonly double Value;

        [OutputConstructor]
        private GetTasksDataTaskProgressStepExtraInfoResult(
            string type,

            string unit,

            double value)
        {
            Type = type;
            Unit = unit;
            Value = value;
        }
    }
}
