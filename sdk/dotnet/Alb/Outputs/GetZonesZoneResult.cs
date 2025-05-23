// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Alb.Outputs
{

    [OutputType]
    public sealed class GetZonesZoneResult
    {
        /// <summary>
        /// The id of the zone.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The id of the zone.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private GetZonesZoneResult(
            string id,

            string zoneId)
        {
            Id = id;
            ZoneId = zoneId;
        }
    }
}
