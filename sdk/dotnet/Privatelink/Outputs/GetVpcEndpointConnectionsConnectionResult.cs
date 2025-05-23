// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Privatelink.Outputs
{

    [OutputType]
    public sealed class GetVpcEndpointConnectionsConnectionResult
    {
        /// <summary>
        /// The status of the connection.
        /// </summary>
        public readonly string ConnectionStatus;
        /// <summary>
        /// The create time of the connection.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// The id of the vpc endpoint.
        /// </summary>
        public readonly string EndpointId;
        /// <summary>
        /// The account id of the vpc endpoint.
        /// </summary>
        public readonly string EndpointOwnerAccountId;
        /// <summary>
        /// The vpc id of the vpc endpoint.
        /// </summary>
        public readonly string EndpointVpcId;
        /// <summary>
        /// The id of the vpc endpoint service.
        /// </summary>
        public readonly string ServiceId;
        /// <summary>
        /// The update time of the connection.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// The available zones.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetVpcEndpointConnectionsConnectionZoneResult> Zones;

        [OutputConstructor]
        private GetVpcEndpointConnectionsConnectionResult(
            string connectionStatus,

            string creationTime,

            string endpointId,

            string endpointOwnerAccountId,

            string endpointVpcId,

            string serviceId,

            string updateTime,

            ImmutableArray<Outputs.GetVpcEndpointConnectionsConnectionZoneResult> zones)
        {
            ConnectionStatus = connectionStatus;
            CreationTime = creationTime;
            EndpointId = endpointId;
            EndpointOwnerAccountId = endpointOwnerAccountId;
            EndpointVpcId = endpointVpcId;
            ServiceId = serviceId;
            UpdateTime = updateTime;
            Zones = zones;
        }
    }
}
