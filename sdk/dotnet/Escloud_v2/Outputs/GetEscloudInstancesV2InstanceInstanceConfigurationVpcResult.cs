// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Escloud_v2.Outputs
{

    [OutputType]
    public sealed class GetEscloudInstancesV2InstanceInstanceConfigurationVpcResult
    {
        /// <summary>
        /// The id of vpc.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// The name of vpc.
        /// </summary>
        public readonly string VpcName;

        [OutputConstructor]
        private GetEscloudInstancesV2InstanceInstanceConfigurationVpcResult(
            string vpcId,

            string vpcName)
        {
            VpcId = vpcId;
            VpcName = vpcName;
        }
    }
}
