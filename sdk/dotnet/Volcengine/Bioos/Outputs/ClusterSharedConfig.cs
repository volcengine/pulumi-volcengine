// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Bioos.Outputs
{

    [OutputType]
    public sealed class ClusterSharedConfig
    {
        /// <summary>
        /// Whether to enable a shared cluster.
        /// </summary>
        public readonly bool Enable;

        [OutputConstructor]
        private ClusterSharedConfig(bool enable)
        {
            Enable = enable;
        }
    }
}