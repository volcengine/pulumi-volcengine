// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Ecs.Inputs
{

    public sealed class InstanceCpuOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of subnuma in socket, only support for ebm. `1` indicates disabling SNC/NPS function. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("numaPerSocket")]
        public Input<int>? NumaPerSocket { get; set; }

        /// <summary>
        /// The per core of threads, only support for ebm. `1` indicates disabling hyper threading function.
        /// </summary>
        [Input("threadsPerCore")]
        public Input<int>? ThreadsPerCore { get; set; }

        public InstanceCpuOptionsArgs()
        {
        }
        public static new InstanceCpuOptionsArgs Empty => new InstanceCpuOptionsArgs();
    }
}