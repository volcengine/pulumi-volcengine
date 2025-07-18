// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Apig.Inputs
{

    public sealed class ApigUpstreamUpstreamSpecK8sServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of k8s service.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The namespace of k8s service.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        /// <summary>
        /// The port of k8s service.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        public ApigUpstreamUpstreamSpecK8sServiceArgs()
        {
        }
        public static new ApigUpstreamUpstreamSpecK8sServiceArgs Empty => new ApigUpstreamUpstreamSpecK8sServiceArgs();
    }
}
