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
    public sealed class BgpPeersBgpPeerResult
    {
        /// <summary>
        /// The id of account.
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// The key of auth.
        /// </summary>
        public readonly string AuthKey;
        /// <summary>
        /// The id of bgp peer.
        /// </summary>
        public readonly string BgpPeerId;
        /// <summary>
        /// The name of bgp peer.
        /// </summary>
        public readonly string BgpPeerName;
        /// <summary>
        /// The create time of bgp peer.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// The Description of bgp peer.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The local asn of bgp peer.
        /// </summary>
        public readonly int LocalAsn;
        /// <summary>
        /// The remote asn of bgp peer.
        /// </summary>
        public readonly int RemoteAsn;
        /// <summary>
        /// The session status of bgp peer.
        /// </summary>
        public readonly string SessionStatus;
        /// <summary>
        /// The status of bgp peer.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The update time of bgp peer.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// The id of virtual interface.
        /// </summary>
        public readonly string VirtualInterfaceId;

        [OutputConstructor]
        private BgpPeersBgpPeerResult(
            string accountId,

            string authKey,

            string bgpPeerId,

            string bgpPeerName,

            string creationTime,

            string description,

            int localAsn,

            int remoteAsn,

            string sessionStatus,

            string status,

            string updateTime,

            string virtualInterfaceId)
        {
            AccountId = accountId;
            AuthKey = authKey;
            BgpPeerId = bgpPeerId;
            BgpPeerName = bgpPeerName;
            CreationTime = creationTime;
            Description = description;
            LocalAsn = localAsn;
            RemoteAsn = remoteAsn;
            SessionStatus = sessionStatus;
            Status = status;
            UpdateTime = updateTime;
            VirtualInterfaceId = virtualInterfaceId;
        }
    }
}