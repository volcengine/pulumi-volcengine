// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Direct_connect.Inputs
{

    public sealed class VirtualInterfacesTagFilterInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The tag key of cloud resource instance.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The tag value of cloud resource instance.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public VirtualInterfacesTagFilterInputArgs()
        {
        }
        public static new VirtualInterfacesTagFilterInputArgs Empty => new VirtualInterfacesTagFilterInputArgs();
    }
}
