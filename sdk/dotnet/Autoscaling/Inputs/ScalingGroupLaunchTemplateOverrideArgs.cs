// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Autoscaling.Inputs
{

    public sealed class ScalingGroupLaunchTemplateOverrideArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance type.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        public ScalingGroupLaunchTemplateOverrideArgs()
        {
        }
        public static new ScalingGroupLaunchTemplateOverrideArgs Empty => new ScalingGroupLaunchTemplateOverrideArgs();
    }
}