// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpc.Outputs
{

    [OutputType]
    public sealed class NetworkAclIngressAclEntry
    {
        /// <summary>
        /// The description of entry.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The id of entry.
        /// </summary>
        public readonly string? NetworkAclEntryId;
        /// <summary>
        /// The name of entry.
        /// </summary>
        public readonly string? NetworkAclEntryName;
        /// <summary>
        /// The policy of entry, default is `accept`. The value can be `accept` or `drop`.
        /// </summary>
        public readonly string? Policy;
        /// <summary>
        /// The port of entry. Default is `-1/-1`. When Protocol is `all`, `icmp` or `gre`, the port range is `-1/-1`, which means no port restriction. When the Protocol is `tcp` or `udp`, the port range is `1~65535`, and the format is `1/200`, `80/80`, which means port 1 to port 200, port 80.
        /// </summary>
        public readonly string? Port;
        /// <summary>
        /// The priority of entry.
        /// </summary>
        public readonly int? Priority;
        /// <summary>
        /// The protocol of entry, default is `all`. The value can be `icmp` or `gre` or `tcp` or `udp` or `all`.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// The SourceCidrIp of entry.
        /// </summary>
        public readonly string? SourceCidrIp;

        [OutputConstructor]
        private NetworkAclIngressAclEntry(
            string? description,

            string? networkAclEntryId,

            string? networkAclEntryName,

            string? policy,

            string? port,

            int? priority,

            string? protocol,

            string? sourceCidrIp)
        {
            Description = description;
            NetworkAclEntryId = networkAclEntryId;
            NetworkAclEntryName = networkAclEntryName;
            Policy = policy;
            Port = port;
            Priority = priority;
            Protocol = protocol;
            SourceCidrIp = sourceCidrIp;
        }
    }
}
