// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Bioos.Inputs
{

    public sealed class ClusterSharedConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable a shared cluster. This value must be `true`.
        /// </summary>
        [Input("enable", required: true)]
        public Input<bool> Enable { get; set; } = null!;

        public ClusterSharedConfigGetArgs()
        {
        }
        public static new ClusterSharedConfigGetArgs Empty => new ClusterSharedConfigGetArgs();
    }
}
