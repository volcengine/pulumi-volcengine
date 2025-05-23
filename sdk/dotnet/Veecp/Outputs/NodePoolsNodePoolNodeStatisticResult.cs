// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Veecp.Outputs
{

    [OutputType]
    public sealed class NodePoolsNodePoolNodeStatisticResult
    {
        /// <summary>
        /// The CreatingCount of Node.
        /// </summary>
        public readonly int CreatingCount;
        /// <summary>
        /// The DeletingCount of Node.
        /// </summary>
        public readonly int DeletingCount;
        /// <summary>
        /// The FailedCount of Node.
        /// </summary>
        public readonly int FailedCount;
        /// <summary>
        /// The RunningCount of Node.
        /// </summary>
        public readonly int RunningCount;
        /// <summary>
        /// (**Deprecated**) This field has been deprecated and is not recommended for use. The StartingCount of Node.
        /// </summary>
        public readonly int StartingCount;
        /// <summary>
        /// (**Deprecated**) This field has been deprecated and is not recommended for use. The StoppedCount of Node.
        /// </summary>
        public readonly int StoppedCount;
        /// <summary>
        /// (**Deprecated**) This field has been deprecated and is not recommended for use. The StoppingCount of Node.
        /// </summary>
        public readonly int StoppingCount;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The UpdatingCount of Node.
        /// </summary>
        public readonly int UpdatingCount;

        [OutputConstructor]
        private NodePoolsNodePoolNodeStatisticResult(
            int creatingCount,

            int deletingCount,

            int failedCount,

            int runningCount,

            int startingCount,

            int stoppedCount,

            int stoppingCount,

            int totalCount,

            int updatingCount)
        {
            CreatingCount = creatingCount;
            DeletingCount = deletingCount;
            FailedCount = failedCount;
            RunningCount = runningCount;
            StartingCount = startingCount;
            StoppedCount = stoppedCount;
            StoppingCount = stoppingCount;
            TotalCount = totalCount;
            UpdatingCount = updatingCount;
        }
    }
}
