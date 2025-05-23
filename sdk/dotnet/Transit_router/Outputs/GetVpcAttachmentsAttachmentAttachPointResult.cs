// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Transit_router.Outputs
{

    [OutputType]
    public sealed class GetVpcAttachmentsAttachmentAttachPointResult
    {
        /// <summary>
        /// The ID of network interface.
        /// </summary>
        public readonly string NetworkInterfaceId;
        /// <summary>
        /// The ID of subnet.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// The ID of zone.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private GetVpcAttachmentsAttachmentAttachPointResult(
            string networkInterfaceId,

            string subnetId,

            string zoneId)
        {
            NetworkInterfaceId = networkInterfaceId;
            SubnetId = subnetId;
            ZoneId = zoneId;
        }
    }
}
