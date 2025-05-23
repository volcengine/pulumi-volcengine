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
    public sealed class InstancesRdsMysqlInstanceEndpointResult
    {
        /// <summary>
        /// Address list.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesRdsMysqlInstanceEndpointAddressResult> Addresses;
        /// <summary>
        /// When the terminal type is read-write terminal or read-only terminal, it supports setting whether new nodes are automatically added.
        /// </summary>
        public readonly string AutoAddNewNodes;
        /// <summary>
        /// Address description.
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
        /// Instance connection terminal ID.
        /// </summary>
        public readonly string EndpointId;
        /// <summary>
        /// The instance connection terminal name.
        /// </summary>
        public readonly string EndpointName;
        /// <summary>
        /// Terminal type:
        /// Cluster: The default terminal. (created by default)
        /// Primary: Primary node terminal.
        /// Custom: Custom terminal.
        /// Direct: Direct connection to the terminal. (Only the operation and maintenance side)
        /// AllNode: All node terminals. (Only the operation and maintenance side).
        /// </summary>
        public readonly string EndpointType;
        /// <summary>
        /// Whether the idle connection reclaim function is enabled. true: Enabled. false: Disabled.
        /// </summary>
        public readonly bool IdleConnectionReclaim;
        /// <summary>
        /// The list of nodes configured by the connection terminal and the corresponding read-only weights.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstancesRdsMysqlInstanceEndpointNodeWeightResult> NodeWeights;
        /// <summary>
        /// Read and write mode:
        /// ReadWrite: read and write
        /// ReadOnly: read only (default).
        /// </summary>
        public readonly string ReadWriteMode;

        [OutputConstructor]
        private InstancesRdsMysqlInstanceEndpointResult(
            ImmutableArray<Outputs.InstancesRdsMysqlInstanceEndpointAddressResult> addresses,

            string autoAddNewNodes,

            string description,

            string enableReadOnly,

            string enableReadWriteSplitting,

            string endpointId,

            string endpointName,

            string endpointType,

            bool idleConnectionReclaim,

            ImmutableArray<Outputs.InstancesRdsMysqlInstanceEndpointNodeWeightResult> nodeWeights,

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
            IdleConnectionReclaim = idleConnectionReclaim;
            NodeWeights = nodeWeights;
            ReadWriteMode = readWriteMode;
        }
    }
}
