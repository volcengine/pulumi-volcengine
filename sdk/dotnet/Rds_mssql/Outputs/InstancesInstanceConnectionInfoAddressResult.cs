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
    public sealed class InstancesInstanceConnectionInfoAddressResult
    {
        /// <summary>
        /// Whether to enable private to public network resolution.
        /// </summary>
        public readonly bool DnsVisibility;
        /// <summary>
        /// The domain.
        /// </summary>
        public readonly string Domain;
        /// <summary>
        /// The eip id for public address.
        /// </summary>
        public readonly string EipId;
        /// <summary>
        /// The ip address.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// The network type.
        /// </summary>
        public readonly string NetworkType;
        /// <summary>
        /// The port of the instance.
        /// </summary>
        public readonly string Port;
        /// <summary>
        /// The subnet id.
        /// </summary>
        public readonly string SubnetId;

        [OutputConstructor]
        private InstancesInstanceConnectionInfoAddressResult(
            bool dnsVisibility,

            string domain,

            string eipId,

            string ipAddress,

            string networkType,

            string port,

            string subnetId)
        {
            DnsVisibility = dnsVisibility;
            Domain = domain;
            EipId = eipId;
            IpAddress = ipAddress;
            NetworkType = networkType;
            Port = port;
            SubnetId = subnetId;
        }
    }
}