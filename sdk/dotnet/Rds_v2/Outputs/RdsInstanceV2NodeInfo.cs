// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_v2.Outputs
{

    [OutputType]
    public sealed class RdsInstanceV2NodeInfo
    {
        /// <summary>
        /// The ID of the node.
        /// </summary>
        public readonly string? NodeId;
        /// <summary>
        /// Masternode specs. Pass
        /// DescribeDBInstanceSpecs Query the instance specifications that can be sold.
        /// </summary>
        public readonly string NodeSpec;
        /// <summary>
        /// Node type, the value is "Primary", "Secondary", "ReadOnly".
        /// </summary>
        public readonly string NodeType;
        /// <summary>
        /// Zone ID.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private RdsInstanceV2NodeInfo(
            string? nodeId,

            string nodeSpec,

            string nodeType,

            string zoneId)
        {
            NodeId = nodeId;
            NodeSpec = nodeSpec;
            NodeType = nodeType;
            ZoneId = zoneId;
        }
    }
}
