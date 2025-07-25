// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Redis.Outputs
{

    [OutputType]
    public sealed class GetInstanceSpecsInstanceSpecResult
    {
        /// <summary>
        /// The architecture type of the Redis instance.
        /// </summary>
        public readonly string ArchType;
        /// <summary>
        /// The type of Redis instance.
        /// </summary>
        public readonly string? InstanceClass;
        /// <summary>
        /// The list of the number of nodes allowed to be used per shard. The number of nodes allowed for different instance types varies.
        /// </summary>
        public readonly ImmutableArray<int> NodeNumbers;
        /// <summary>
        /// The List of capacity specifications for a single shard.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceSpecsInstanceSpecShardCapacitySpecResult> ShardCapacitySpecs;
        /// <summary>
        /// The list of shards that the instance is allowed to use. The number of shards allowed for use varies among different instance architecture types.
        /// </summary>
        public readonly ImmutableArray<int> ShardNumbers;

        [OutputConstructor]
        private GetInstanceSpecsInstanceSpecResult(
            string archType,

            string? instanceClass,

            ImmutableArray<int> nodeNumbers,

            ImmutableArray<Outputs.GetInstanceSpecsInstanceSpecShardCapacitySpecResult> shardCapacitySpecs,

            ImmutableArray<int> shardNumbers)
        {
            ArchType = archType;
            InstanceClass = instanceClass;
            NodeNumbers = nodeNumbers;
            ShardCapacitySpecs = shardCapacitySpecs;
            ShardNumbers = shardNumbers;
        }
    }
}
