// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Veenedge.Inputs
{

    public sealed class CloudServerStorageConfigDataDiskListGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The capacity of storage.
        /// </summary>
        [Input("capacity", required: true)]
        public Input<string> Capacity { get; set; } = null!;

        /// <summary>
        /// The type of storage. The value can be `CloudBlockHDD` or `CloudBlockSSD`.
        /// </summary>
        [Input("storageType", required: true)]
        public Input<string> StorageType { get; set; } = null!;

        public CloudServerStorageConfigDataDiskListGetArgs()
        {
        }
    }
}