// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vke.Inputs
{

    public sealed class ClusterPodsConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Flannel network configuration.
        /// </summary>
        [Input("flannelConfig")]
        public Input<Inputs.ClusterPodsConfigFlannelConfigArgs>? FlannelConfig { get; set; }

        /// <summary>
        /// The container network model of the cluster, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.
        /// </summary>
        [Input("podNetworkMode", required: true)]
        public Input<string> PodNetworkMode { get; set; } = null!;

        /// <summary>
        /// VPC-CNI network configuration.
        /// </summary>
        [Input("vpcCniConfig")]
        public Input<Inputs.ClusterPodsConfigVpcCniConfigArgs>? VpcCniConfig { get; set; }

        public ClusterPodsConfigArgs()
        {
        }
        public static new ClusterPodsConfigArgs Empty => new ClusterPodsConfigArgs();
    }
}
