// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Ebs.Outputs
{

    [OutputType]
    public sealed class MaxExtraPerformancesPerformanceMaxExtraPerformanceCanPurchaseResult
    {
        /// <summary>
        /// The type of the extra performance.
        /// </summary>
        public readonly string ExtraPerformanceTypeId;
        /// <summary>
        /// The limit of the extra performance.
        /// </summary>
        public readonly int Limit;

        [OutputConstructor]
        private MaxExtraPerformancesPerformanceMaxExtraPerformanceCanPurchaseResult(
            string extraPerformanceTypeId,

            int limit)
        {
            ExtraPerformanceTypeId = extraPerformanceTypeId;
            Limit = limit;
        }
    }
}
