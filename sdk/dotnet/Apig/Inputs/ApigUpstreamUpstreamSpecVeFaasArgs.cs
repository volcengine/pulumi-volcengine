// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Apig.Inputs
{

    public sealed class ApigUpstreamUpstreamSpecVeFaasArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The function id of vefaas.
        /// </summary>
        [Input("functionId", required: true)]
        public Input<string> FunctionId { get; set; } = null!;

        public ApigUpstreamUpstreamSpecVeFaasArgs()
        {
        }
        public static new ApigUpstreamUpstreamSpecVeFaasArgs Empty => new ApigUpstreamUpstreamSpecVeFaasArgs();
    }
}
