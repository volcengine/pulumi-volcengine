// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Alb.Inputs
{

    public sealed class AlbZoneMappingLoadBalancerAddressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Eip address of the Alb in this availability zone.
        /// </summary>
        [Input("eipAddress")]
        public Input<string>? EipAddress { get; set; }

        /// <summary>
        /// The Eip id of alb instance in this availability zone.
        /// </summary>
        [Input("eipId")]
        public Input<string>? EipId { get; set; }

        /// <summary>
        /// The Eni address of the Alb in this availability zone.
        /// </summary>
        [Input("eniAddress")]
        public Input<string>? EniAddress { get; set; }

        /// <summary>
        /// The Eni id of the Alb in this availability zone.
        /// </summary>
        [Input("eniId")]
        public Input<string>? EniId { get; set; }

        /// <summary>
        /// The Eni Ipv6 address of the Alb in this availability zone.
        /// </summary>
        [Input("eniIpv6Address")]
        public Input<string>? EniIpv6Address { get; set; }

        /// <summary>
        /// The Ipv6 Eip id of alb instance in this availability zone.
        /// </summary>
        [Input("ipv6EipId")]
        public Input<string>? Ipv6EipId { get; set; }

        public AlbZoneMappingLoadBalancerAddressArgs()
        {
        }
        public static new AlbZoneMappingLoadBalancerAddressArgs Empty => new AlbZoneMappingLoadBalancerAddressArgs();
    }
}
