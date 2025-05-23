// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rabbitmq.Inputs
{

    public sealed class InstanceEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The endpoint type of the rabbitmq instance.
        /// </summary>
        [Input("endpointType")]
        public Input<string>? EndpointType { get; set; }

        /// <summary>
        /// The internal endpoint of the rabbitmq instance.
        /// </summary>
        [Input("internalEndpoint")]
        public Input<string>? InternalEndpoint { get; set; }

        /// <summary>
        /// The network type of the rabbitmq instance.
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// The public endpoint of the rabbitmq instance.
        /// </summary>
        [Input("publicEndpoint")]
        public Input<string>? PublicEndpoint { get; set; }

        public InstanceEndpointArgs()
        {
        }
        public static new InstanceEndpointArgs Empty => new InstanceEndpointArgs();
    }
}
