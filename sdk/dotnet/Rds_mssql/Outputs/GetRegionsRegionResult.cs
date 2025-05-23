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
    public sealed class GetRegionsRegionResult
    {
        /// <summary>
        /// The id of the region.
        /// </summary>
        public readonly string RegionId;
        /// <summary>
        /// The name of region.
        /// </summary>
        public readonly string RegionName;

        [OutputConstructor]
        private GetRegionsRegionResult(
            string regionId,

            string regionName)
        {
            RegionId = regionId;
            RegionName = regionName;
        }
    }
}
