// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tos.Outputs
{

    [OutputType]
    public sealed class BucketInventoriesInventoryConfigurationScheduleResult
    {
        /// <summary>
        /// The export schedule of the bucket inventory. Valid values: `Daily`, `Weekly`.
        /// </summary>
        public readonly string Frequency;

        [OutputConstructor]
        private BucketInventoriesInventoryConfigurationScheduleResult(string frequency)
        {
            Frequency = frequency;
        }
    }
}
