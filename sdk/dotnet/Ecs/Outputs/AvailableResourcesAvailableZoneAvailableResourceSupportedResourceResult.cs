// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Ecs.Outputs
{

    [OutputType]
    public sealed class AvailableResourcesAvailableZoneAvailableResourceSupportedResourceResult
    {
        /// <summary>
        /// The resource status of the available zone. Valid values: `Available`, `SoldOut`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The value of the resource.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private AvailableResourcesAvailableZoneAvailableResourceSupportedResourceResult(
            string status,

            string value)
        {
            Status = status;
            Value = value;
        }
    }
}