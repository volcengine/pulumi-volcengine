// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_mysql.Inputs
{

    public sealed class InstanceEndpointNodeWeightGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Node ID.
        /// </summary>
        [Input("nodeId")]
        public Input<string>? NodeId { get; set; }

        /// <summary>
        /// Node type. Value: Primary: Primary node.
        /// Secondary: Standby node.
        /// ReadOnly: Read-only node.
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        /// <summary>
        /// The weight of the node.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public InstanceEndpointNodeWeightGetArgs()
        {
        }
        public static new InstanceEndpointNodeWeightGetArgs Empty => new InstanceEndpointNodeWeightGetArgs();
    }
}
