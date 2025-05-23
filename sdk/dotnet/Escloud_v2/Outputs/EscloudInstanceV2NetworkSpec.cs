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
    public sealed class EscloudInstanceV2NetworkSpec
    {
        /// <summary>
        /// The bandwidth of the eip. Unit: Mbps.
        /// </summary>
        public readonly int Bandwidth;
        /// <summary>
        /// Whether the eip is opened.
        /// </summary>
        public readonly bool IsOpen;
        /// <summary>
        /// The spec name of public network.
        /// </summary>
        public readonly string SpecName;
        /// <summary>
        /// The type of public network, valid values: `Elasticsearch`, `Kibana`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private EscloudInstanceV2NetworkSpec(
            int bandwidth,

            bool isOpen,

            string specName,

            string type)
        {
            Bandwidth = bandwidth;
            IsOpen = isOpen;
            SpecName = specName;
            Type = type;
        }
    }
}
