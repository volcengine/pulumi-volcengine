// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Bandwidth_package.Outputs
{

    [OutputType]
    public sealed class BandwidthPackagesPackageResult
    {
        /// <summary>
        /// The bandwidth of the bandwidth package.
        /// </summary>
        public readonly int Bandwidth;
        /// <summary>
        /// The id of the bandwidth package.
        /// </summary>
        public readonly string BandwidthPackageId;
        /// <summary>
        /// Shared bandwidth package name to be queried.
        /// </summary>
        public readonly string BandwidthPackageName;
        /// <summary>
        /// The billing type of the bandwidth package.
        /// </summary>
        public readonly string BillingType;
        /// <summary>
        /// The business status of the bandwidth package.
        /// </summary>
        public readonly string BusinessStatus;
        /// <summary>
        /// The creation time of the bandwidth package.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// The deleted time of the bandwidth package.
        /// </summary>
        public readonly string DeletedTime;
        /// <summary>
        /// List of public IP information included in the shared bandwidth package.
        /// </summary>
        public readonly ImmutableArray<Outputs.BandwidthPackagesPackageEipAddressResult> EipAddresses;
        /// <summary>
        /// The expiration time of the bandwidth package.
        /// </summary>
        public readonly string ExpiredTime;
        /// <summary>
        /// The id of the bandwidth package.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Line types for shared bandwidth packages.
        /// </summary>
        public readonly string Isp;
        /// <summary>
        /// The overdue time of the bandwidth package.
        /// </summary>
        public readonly string OverdueTime;
        /// <summary>
        /// The project name of the bandwidth package to be queried.
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// The IP protocol values for shared bandwidth packages are as follows: `IPv4`: IPv4 protocol. `IPv6`: IPv6 protocol.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// Security protection types for shared bandwidth packages. Parameter - N: Indicates the number of security protection types, currently only supports taking 1. Value: `AntiDDoS_Enhanced`.
        /// </summary>
        public readonly ImmutableArray<string> SecurityProtectionTypes;
        /// <summary>
        /// The status of the bandwidth package.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.BandwidthPackagesPackageTagResult> Tags;
        /// <summary>
        /// The update time of the bandwidth package.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private BandwidthPackagesPackageResult(
            int bandwidth,

            string bandwidthPackageId,

            string bandwidthPackageName,

            string billingType,

            string businessStatus,

            string creationTime,

            string deletedTime,

            ImmutableArray<Outputs.BandwidthPackagesPackageEipAddressResult> eipAddresses,

            string expiredTime,

            string id,

            string isp,

            string overdueTime,

            string projectName,

            string protocol,

            ImmutableArray<string> securityProtectionTypes,

            string status,

            ImmutableArray<Outputs.BandwidthPackagesPackageTagResult> tags,

            string updateTime)
        {
            Bandwidth = bandwidth;
            BandwidthPackageId = bandwidthPackageId;
            BandwidthPackageName = bandwidthPackageName;
            BillingType = billingType;
            BusinessStatus = businessStatus;
            CreationTime = creationTime;
            DeletedTime = deletedTime;
            EipAddresses = eipAddresses;
            ExpiredTime = expiredTime;
            Id = id;
            Isp = isp;
            OverdueTime = overdueTime;
            ProjectName = projectName;
            Protocol = protocol;
            SecurityProtectionTypes = securityProtectionTypes;
            Status = status;
            Tags = tags;
            UpdateTime = updateTime;
        }
    }
}