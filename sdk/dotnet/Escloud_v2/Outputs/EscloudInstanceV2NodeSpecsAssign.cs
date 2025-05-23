// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Escloud_v2.Outputs
{

    [OutputType]
    public sealed class EscloudInstanceV2NodeSpecsAssign
    {
        /// <summary>
        /// The extra performance of FlexPL storage spec.
        /// </summary>
        public readonly Outputs.EscloudInstanceV2NodeSpecsAssignExtraPerformance? ExtraPerformance;
        /// <summary>
        /// The number of node.
        /// </summary>
        public readonly int Number;
        /// <summary>
        /// The name of compute resource spec.
        /// </summary>
        public readonly string ResourceSpecName;
        /// <summary>
        /// The size of storage. Unit: GiB. the adjustment step size is 10GiB. Default is 100 GiB. Kibana NodeSpecsAssign should specify this field to 0.
        /// </summary>
        public readonly int StorageSize;
        /// <summary>
        /// The name of storage spec. Kibana NodeSpecsAssign should specify this field to ``.
        /// </summary>
        public readonly string StorageSpecName;
        /// <summary>
        /// The type of node, valid values: `Master`, `Hot`, `Cold`, `Warm`, `Kibana`, `Coordinator`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private EscloudInstanceV2NodeSpecsAssign(
            Outputs.EscloudInstanceV2NodeSpecsAssignExtraPerformance? extraPerformance,

            int number,

            string resourceSpecName,

            int storageSize,

            string storageSpecName,

            string type)
        {
            ExtraPerformance = extraPerformance;
            Number = number;
            ResourceSpecName = resourceSpecName;
            StorageSize = storageSize;
            StorageSpecName = storageSpecName;
            Type = type;
        }
    }
}
