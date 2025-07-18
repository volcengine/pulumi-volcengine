// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Kms.Inputs
{

    public sealed class KeyMultiRegionConfigurationReplicaKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The region id of multi-region key.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Input("trn")]
        public Input<string>? Trn { get; set; }

        public KeyMultiRegionConfigurationReplicaKeyArgs()
        {
        }
        public static new KeyMultiRegionConfigurationReplicaKeyArgs Empty => new KeyMultiRegionConfigurationReplicaKeyArgs();
    }
}
