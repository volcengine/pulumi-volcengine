// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Private_zone.Inputs
{

    public sealed class PrivateZoneVpcGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The region of the bind vpc. The default value is the region of the default provider config.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The id of the bind vpc.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public PrivateZoneVpcGetArgs()
        {
        }
        public static new PrivateZoneVpcGetArgs Empty => new PrivateZoneVpcGetArgs();
    }
}
