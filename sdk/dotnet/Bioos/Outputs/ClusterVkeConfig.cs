// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Bioos.Outputs
{

    [OutputType]
    public sealed class ClusterVkeConfig
    {
        /// <summary>
        /// The id of the vke cluster.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// The name of the StorageClass that the vke cluster has installed.
        /// </summary>
        public readonly string StorageClass;

        [OutputConstructor]
        private ClusterVkeConfig(
            string clusterId,

            string storageClass)
        {
            ClusterId = clusterId;
            StorageClass = storageClass;
        }
    }
}