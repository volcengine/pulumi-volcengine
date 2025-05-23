// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_mysql.Outputs
{

    [OutputType]
    public sealed class EndpointsEndpointResult
    {
        /// <summary>
        /// Address list.
        /// </summary>
        public readonly ImmutableArray<Outputs.EndpointsEndpointAddressResult> Addresses;
        /// <summary>
        /// When the terminal type is read-write terminal or read-only terminal, it supports setting whether new nodes are automatically added.
        /// </summary>
        public readonly string AutoAddNewNodes;
        /// <summary>
        /// The description of the mysql endpoint.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Whether global read-only is enabled, value: Enable: Enable. Disable: Disabled.
        /// </summary>
        public readonly string EnableReadOnly;
        /// <summary>
        /// Whether read-write separation is enabled, value: Enable: Enable. Disable: Disabled.
        /// </summary>
        public readonly string EnableReadWriteSplitting;
        /// <summary>
        /// The id of the mysql endpoint.
        /// </summary>
        public readonly string EndpointId;
        /// <summary>
        /// The name of the mysql endpoint.
        /// </summary>
        public readonly string EndpointName;
        /// <summary>
        /// The endpoint type of the mysql endpoint.
        /// </summary>
        public readonly string EndpointType;
        /// <summary>
        /// The id of the mysql endpoint.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of nodes configured by the connection terminal and the corresponding read-only weights.
        /// </summary>
        public readonly ImmutableArray<Outputs.EndpointsEndpointReadOnlyNodeWeightResult> ReadOnlyNodeWeights;
        /// <summary>
        /// The read write mode.
        /// </summary>
        public readonly string ReadWriteMode;

        [OutputConstructor]
        private EndpointsEndpointResult(
            ImmutableArray<Outputs.EndpointsEndpointAddressResult> addresses,

            string autoAddNewNodes,

            string description,

            string enableReadOnly,

            string enableReadWriteSplitting,

            string endpointId,

            string endpointName,

            string endpointType,

            string id,

            ImmutableArray<Outputs.EndpointsEndpointReadOnlyNodeWeightResult> readOnlyNodeWeights,

            string readWriteMode)
        {
            Addresses = addresses;
            AutoAddNewNodes = autoAddNewNodes;
            Description = description;
            EnableReadOnly = enableReadOnly;
            EnableReadWriteSplitting = enableReadWriteSplitting;
            EndpointId = endpointId;
            EndpointName = endpointName;
            EndpointType = endpointType;
            Id = id;
            ReadOnlyNodeWeights = readOnlyNodeWeights;
            ReadWriteMode = readWriteMode;
        }
    }
}
