// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vke.Inputs
{

    public sealed class NodePoolNodeConfigSystemVolumeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Size of SystemVolume, the value range in 20~2048.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The Type of SystemVolume, the value can be `PTSSD` or `ESSD_PL0` or `ESSD_FlexPL`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public NodePoolNodeConfigSystemVolumeGetArgs()
        {
        }
        public static new NodePoolNodeConfigSystemVolumeGetArgs Empty => new NodePoolNodeConfigSystemVolumeGetArgs();
    }
}