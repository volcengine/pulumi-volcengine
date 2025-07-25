// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Apig.Inputs
{

    public sealed class ApigUpstreamSourceSourceSpecK8sSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cluster id of k8s source.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The cluster type of k8s source.
        /// </summary>
        [Input("clusterType")]
        public Input<string>? ClusterType { get; set; }

        public ApigUpstreamSourceSourceSpecK8sSourceArgs()
        {
        }
        public static new ApigUpstreamSourceSourceSpecK8sSourceArgs Empty => new ApigUpstreamSourceSourceSpecK8sSourceArgs();
    }
}
