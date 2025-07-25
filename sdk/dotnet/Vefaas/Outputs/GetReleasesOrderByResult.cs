// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vefaas.Outputs
{

    [OutputType]
    public sealed class GetReleasesOrderByResult
    {
        /// <summary>
        /// Whether the sorting result is sorted in ascending order.
        /// </summary>
        public readonly bool? Ascend;
        /// <summary>
        /// Key names used for sorting.
        /// </summary>
        public readonly string? Key;

        [OutputConstructor]
        private GetReleasesOrderByResult(
            bool? ascend,

            string? key)
        {
            Ascend = ascend;
            Key = key;
        }
    }
}
