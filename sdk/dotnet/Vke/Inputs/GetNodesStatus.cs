// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vke.Inputs
{

    public sealed class GetNodesStatusArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Type of Node Condition, the value is `Progressing` or `Ok` or `Unschedulable` or `InitilizeFailed` or `Unknown` or `NotReady` or `Security` or `Balance` or `ResourceCleanupFailed`.
        /// </summary>
        [Input("conditionsType")]
        public string? ConditionsType { get; set; }

        /// <summary>
        /// The Phase of Node, the value is `Creating` or `Running` or `Updating` or `Deleting` or `Failed` or `Starting` or `Stopping` or `Stopped`.
        /// </summary>
        [Input("phase")]
        public string? Phase { get; set; }

        public GetNodesStatusArgs()
        {
        }
        public static new GetNodesStatusArgs Empty => new GetNodesStatusArgs();
    }
}
