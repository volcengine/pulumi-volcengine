// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Alb.Inputs
{

    public sealed class AlbZoneMappingArgs : global::Pulumi.ResourceArgs
    {
        [Input("loadBalancerAddresses")]
        private InputList<Inputs.AlbZoneMappingLoadBalancerAddressArgs>? _loadBalancerAddresses;

        /// <summary>
        /// The IP address information of the Alb in this availability zone.
        /// </summary>
        public InputList<Inputs.AlbZoneMappingLoadBalancerAddressArgs> LoadBalancerAddresses
        {
            get => _loadBalancerAddresses ?? (_loadBalancerAddresses = new InputList<Inputs.AlbZoneMappingLoadBalancerAddressArgs>());
            set => _loadBalancerAddresses = value;
        }

        /// <summary>
        /// The subnet id of the Alb in this availability zone.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// The availability zone id of the Alb.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public AlbZoneMappingArgs()
        {
        }
        public static new AlbZoneMappingArgs Empty => new AlbZoneMappingArgs();
    }
}
