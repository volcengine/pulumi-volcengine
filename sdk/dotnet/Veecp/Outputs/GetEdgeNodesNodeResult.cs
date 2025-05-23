// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Veecp.Outputs
{

    [OutputType]
    public sealed class GetEdgeNodesNodeResult
    {
        /// <summary>
        /// The bootstrap script of node.
        /// </summary>
        public readonly string BootstrapScript;
        /// <summary>
        /// The cluster id of node.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// The Condition of Node.
        /// </summary>
        public readonly ImmutableArray<string> ConditionTypes;
        /// <summary>
        /// The Create Client Token.
        /// </summary>
        public readonly string CreateClientToken;
        /// <summary>
        /// The create time of Node.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The edge node type of node.
        /// </summary>
        public readonly string EdgeNodeType;
        /// <summary>
        /// The ID of Node.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The instance id of node.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The Name of Node.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The node pool id.
        /// </summary>
        public readonly string NodePoolId;
        /// <summary>
        /// The Phase of Node.
        /// </summary>
        public readonly string Phase;
        /// <summary>
        /// The profile of node. Distinguish between edge and central nodes.
        /// </summary>
        public readonly string Profile;
        /// <summary>
        /// The provider id of node.
        /// </summary>
        public readonly string ProviderId;
        /// <summary>
        /// The update time of Node.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetEdgeNodesNodeResult(
            string bootstrapScript,

            string clusterId,

            ImmutableArray<string> conditionTypes,

            string createClientToken,

            string createTime,

            string edgeNodeType,

            string id,

            string instanceId,

            string name,

            string nodePoolId,

            string phase,

            string profile,

            string providerId,

            string updateTime)
        {
            BootstrapScript = bootstrapScript;
            ClusterId = clusterId;
            ConditionTypes = conditionTypes;
            CreateClientToken = createClientToken;
            CreateTime = createTime;
            EdgeNodeType = edgeNodeType;
            Id = id;
            InstanceId = instanceId;
            Name = name;
            NodePoolId = nodePoolId;
            Phase = phase;
            Profile = profile;
            ProviderId = providerId;
            UpdateTime = updateTime;
        }
    }
}
