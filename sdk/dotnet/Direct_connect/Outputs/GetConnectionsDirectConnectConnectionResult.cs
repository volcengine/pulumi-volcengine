// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Direct_connect.Outputs
{

    [OutputType]
    public sealed class GetConnectionsDirectConnectConnectionResult
    {
        /// <summary>
        /// The account ID which the physical leased line belongs.
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// The bandwidth of direct connect.
        /// </summary>
        public readonly int Bandwidth;
        /// <summary>
        /// The dedicated line billing type,only support `1` for yearly and monthly billing currently.
        /// </summary>
        public readonly int BillingType;
        /// <summary>
        /// The dedicated line billing status.
        /// </summary>
        public readonly string BusinessStatus;
        /// <summary>
        /// The connection type of physical leased line,valid value contains `SharedConnection`,`DedicatedConnection`.
        /// </summary>
        public readonly string ConnectionType;
        /// <summary>
        /// The creation time of direct connect.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// The dedicated line contact email.
        /// </summary>
        public readonly string CustomerContactEmail;
        /// <summary>
        /// The dedicated line contact phone.
        /// </summary>
        public readonly string CustomerContactPhone;
        /// <summary>
        /// The dedicated line contact name.
        /// </summary>
        public readonly string CustomerName;
        /// <summary>
        /// The expected resource force collection time.
        /// </summary>
        public readonly string DeletedTime;
        /// <summary>
        /// The description of direct connect connection.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The ID of the physical leased line access point.
        /// </summary>
        public readonly string DirectConnectAccessPointId;
        /// <summary>
        /// The ID of direct connect connection.
        /// </summary>
        public readonly string DirectConnectConnectionId;
        /// <summary>
        /// The name of directi connect connection.
        /// </summary>
        public readonly string DirectConnectConnectionName;
        /// <summary>
        /// The expect bandwidth of direct connect.
        /// </summary>
        public readonly int ExpectBandwidth;
        /// <summary>
        /// The expired time.
        /// </summary>
        public readonly string ExpiredTime;
        /// <summary>
        /// The operator of the physical leased line,valid value contains `ChinaTelecom`,`ChinaMobile`,`ChinaUnicom`,`ChinaOther`.
        /// </summary>
        public readonly string LineOperator;
        /// <summary>
        /// The account ID of physical leased line to which the shared leased line belongs.If the physical leased line type is an exclusive leased line,this parameter returns empty.
        /// </summary>
        public readonly string ParentConnectionAccountId;
        /// <summary>
        /// The ID of the physical leased line to which the shared leased line belongs. If the physical leased line type is an exclusive leased line, this parameter returns empty.
        /// </summary>
        public readonly string ParentConnectionId;
        /// <summary>
        /// The peer access point of the physical leased line.
        /// </summary>
        public readonly string PeerLocation;
        /// <summary>
        /// The dedicated line port spec.
        /// </summary>
        public readonly string PortSpec;
        /// <summary>
        /// The port type of direct connect.
        /// </summary>
        public readonly string PortType;
        /// <summary>
        /// The status of physical leased line.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// All tags that physical leased line added.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetConnectionsDirectConnectConnectionTagResult> Tags;
        /// <summary>
        /// The update time of direct connect.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// The vlan ID of shared connection,if `connection_type` is `DedicatedConnection`,this parameter returns 0.
        /// </summary>
        public readonly int VlanId;

        [OutputConstructor]
        private GetConnectionsDirectConnectConnectionResult(
            string accountId,

            int bandwidth,

            int billingType,

            string businessStatus,

            string connectionType,

            string creationTime,

            string customerContactEmail,

            string customerContactPhone,

            string customerName,

            string deletedTime,

            string description,

            string directConnectAccessPointId,

            string directConnectConnectionId,

            string directConnectConnectionName,

            int expectBandwidth,

            string expiredTime,

            string lineOperator,

            string parentConnectionAccountId,

            string parentConnectionId,

            string peerLocation,

            string portSpec,

            string portType,

            string status,

            ImmutableArray<Outputs.GetConnectionsDirectConnectConnectionTagResult> tags,

            string updateTime,

            int vlanId)
        {
            AccountId = accountId;
            Bandwidth = bandwidth;
            BillingType = billingType;
            BusinessStatus = businessStatus;
            ConnectionType = connectionType;
            CreationTime = creationTime;
            CustomerContactEmail = customerContactEmail;
            CustomerContactPhone = customerContactPhone;
            CustomerName = customerName;
            DeletedTime = deletedTime;
            Description = description;
            DirectConnectAccessPointId = directConnectAccessPointId;
            DirectConnectConnectionId = directConnectConnectionId;
            DirectConnectConnectionName = directConnectConnectionName;
            ExpectBandwidth = expectBandwidth;
            ExpiredTime = expiredTime;
            LineOperator = lineOperator;
            ParentConnectionAccountId = parentConnectionAccountId;
            ParentConnectionId = parentConnectionId;
            PeerLocation = peerLocation;
            PortSpec = portSpec;
            PortType = portType;
            Status = status;
            Tags = tags;
            UpdateTime = updateTime;
            VlanId = vlanId;
        }
    }
}
