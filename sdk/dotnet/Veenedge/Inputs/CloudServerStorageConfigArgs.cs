// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Veenedge.Inputs
{

    public sealed class CloudServerStorageConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("dataDiskLists")]
        private InputList<Inputs.CloudServerStorageConfigDataDiskListArgs>? _dataDiskLists;

        /// <summary>
        /// The disk list info of data.
        /// </summary>
        public InputList<Inputs.CloudServerStorageConfigDataDiskListArgs> DataDiskLists
        {
            get => _dataDiskLists ?? (_dataDiskLists = new InputList<Inputs.CloudServerStorageConfigDataDiskListArgs>());
            set => _dataDiskLists = value;
        }

        /// <summary>
        /// The disk info of system.
        /// </summary>
        [Input("systemDisk", required: true)]
        public Input<Inputs.CloudServerStorageConfigSystemDiskArgs> SystemDisk { get; set; } = null!;

        public CloudServerStorageConfigArgs()
        {
        }
        public static new CloudServerStorageConfigArgs Empty => new CloudServerStorageConfigArgs();
    }
}