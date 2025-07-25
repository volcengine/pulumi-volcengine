// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Apig.Inputs
{

    public sealed class ApigUpstreamUpstreamSpecVeMlpK8sServiceClusterInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account id of k8s service.
        /// </summary>
        [Input("accountId", required: true)]
        public Input<int> AccountId { get; set; } = null!;

        /// <summary>
        /// The cluster name of k8s service.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        public ApigUpstreamUpstreamSpecVeMlpK8sServiceClusterInfoArgs()
        {
        }
        public static new ApigUpstreamUpstreamSpecVeMlpK8sServiceClusterInfoArgs Empty => new ApigUpstreamUpstreamSpecVeMlpK8sServiceClusterInfoArgs();
    }
}
