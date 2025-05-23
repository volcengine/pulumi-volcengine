// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Ecs.Outputs
{

    [OutputType]
    public sealed class GetDeploymentSetsDeploymentSetResult
    {
        /// <summary>
        /// The ID of ECS DeploymentSet.
        /// </summary>
        public readonly string DeploymentSetId;
        /// <summary>
        /// The name of ECS DeploymentSet.
        /// </summary>
        public readonly string DeploymentSetName;
        /// <summary>
        /// The description of ECS DeploymentSet.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The granularity of ECS DeploymentSet.Valid values: switch, host, rack.
        /// </summary>
        public readonly string Granularity;
        /// <summary>
        /// The strategy of ECS DeploymentSet.
        /// </summary>
        public readonly string Strategy;

        [OutputConstructor]
        private GetDeploymentSetsDeploymentSetResult(
            string deploymentSetId,

            string deploymentSetName,

            string description,

            string granularity,

            string strategy)
        {
            DeploymentSetId = deploymentSetId;
            DeploymentSetName = deploymentSetName;
            Description = description;
            Granularity = granularity;
            Strategy = strategy;
        }
    }
}
