// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Nas.Outputs
{

    [OutputType]
    public sealed class GetZonesZoneSaleResult
    {
        /// <summary>
        /// The type of file system.
        /// </summary>
        public readonly string FileSystemType;
        /// <summary>
        /// The type of protocol.
        /// </summary>
        public readonly string ProtocolType;
        /// <summary>
        /// The status info.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The type of storage.
        /// </summary>
        public readonly string StorageType;

        [OutputConstructor]
        private GetZonesZoneSaleResult(
            string fileSystemType,

            string protocolType,

            string status,

            string storageType)
        {
            FileSystemType = fileSystemType;
            ProtocolType = protocolType;
            Status = status;
            StorageType = storageType;
        }
    }
}
