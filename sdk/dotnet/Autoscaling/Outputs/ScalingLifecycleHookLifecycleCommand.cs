// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Autoscaling.Outputs
{

    [OutputType]
    public sealed class ScalingLifecycleHookLifecycleCommand
    {
        /// <summary>
        /// Batch job command ID, which indicates the batch job command to be executed after triggering the lifecycle hook and installed in the instance.
        /// </summary>
        public readonly string CommandId;
        /// <summary>
        /// Parameters and parameter values in batch job commands.
        /// The number of parameters ranges from 0 to 60.
        /// </summary>
        public readonly string? Parameters;

        [OutputConstructor]
        private ScalingLifecycleHookLifecycleCommand(
            string commandId,

            string? parameters)
        {
            CommandId = commandId;
            Parameters = parameters;
        }
    }
}