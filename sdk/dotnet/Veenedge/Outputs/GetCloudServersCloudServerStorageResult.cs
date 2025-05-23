// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Veenedge.Outputs
{

    [OutputType]
    public sealed class GetCloudServersCloudServerStorageResult
    {
        /// <summary>
        /// The disk info of data.
        /// </summary>
        public readonly Outputs.GetCloudServersCloudServerStorageDataDiskResult DataDisk;
        /// <summary>
        /// The disk list info of data.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCloudServersCloudServerStorageDataDiskListResult> DataDiskLists;
        /// <summary>
        /// The disk info of system.
        /// </summary>
        public readonly Outputs.GetCloudServersCloudServerStorageSystemDiskResult SystemDisk;

        [OutputConstructor]
        private GetCloudServersCloudServerStorageResult(
            Outputs.GetCloudServersCloudServerStorageDataDiskResult dataDisk,

            ImmutableArray<Outputs.GetCloudServersCloudServerStorageDataDiskListResult> dataDiskLists,

            Outputs.GetCloudServersCloudServerStorageSystemDiskResult systemDisk)
        {
            DataDisk = dataDisk;
            DataDiskLists = dataDiskLists;
            SystemDisk = systemDisk;
        }
    }
}
