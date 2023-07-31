// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Redis.Outputs
{

    [OutputType]
    public sealed class AllowListsAllowListAssociatedInstanceResult
    {
        /// <summary>
        /// The Id of instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// Name of instance.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// Id of virtual private cloud.
        /// </summary>
        public readonly string Vpc;

        [OutputConstructor]
        private AllowListsAllowListAssociatedInstanceResult(
            string instanceId,

            string instanceName,

            string vpc)
        {
            InstanceId = instanceId;
            InstanceName = instanceName;
            Vpc = vpc;
        }
    }
}