// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Vke.Inputs
{

    public sealed class NodeKubernetesConfigTaintArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Effect of Taints, the value can be `NoSchedule` or `NoExecute` or `PreferNoSchedule`.
        /// </summary>
        [Input("effect")]
        public Input<string>? Effect { get; set; }

        /// <summary>
        /// The Key of Taints.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The Value of Taints.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public NodeKubernetesConfigTaintArgs()
        {
        }
    }
}