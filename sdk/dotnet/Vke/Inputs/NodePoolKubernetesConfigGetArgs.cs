// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vke.Inputs
{

    public sealed class NodePoolKubernetesConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Cordon of KubernetesConfig.
        /// </summary>
        [Input("cordon", required: true)]
        public Input<bool> Cordon { get; set; } = null!;

        [Input("labels")]
        private InputList<Inputs.NodePoolKubernetesConfigLabelGetArgs>? _labels;

        /// <summary>
        /// The Labels of KubernetesConfig.
        /// </summary>
        public InputList<Inputs.NodePoolKubernetesConfigLabelGetArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.NodePoolKubernetesConfigLabelGetArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// The NamePrefix of node metadata.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("taints")]
        private InputList<Inputs.NodePoolKubernetesConfigTaintGetArgs>? _taints;

        /// <summary>
        /// The Taints of KubernetesConfig.
        /// </summary>
        public InputList<Inputs.NodePoolKubernetesConfigTaintGetArgs> Taints
        {
            get => _taints ?? (_taints = new InputList<Inputs.NodePoolKubernetesConfigTaintGetArgs>());
            set => _taints = value;
        }

        public NodePoolKubernetesConfigGetArgs()
        {
        }
        public static new NodePoolKubernetesConfigGetArgs Empty => new NodePoolKubernetesConfigGetArgs();
    }
}