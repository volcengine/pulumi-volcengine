// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_postgresql.Inputs
{

    public sealed class AllowlistAssociatedInstanceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the postgresql instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The name of the postgresql instance.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The id of the vpc.
        /// </summary>
        [Input("vpc")]
        public Input<string>? Vpc { get; set; }

        public AllowlistAssociatedInstanceGetArgs()
        {
        }
        public static new AllowlistAssociatedInstanceGetArgs Empty => new AllowlistAssociatedInstanceGetArgs();
    }
}