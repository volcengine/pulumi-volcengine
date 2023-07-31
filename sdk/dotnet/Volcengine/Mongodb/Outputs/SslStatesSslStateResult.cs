// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Mongodb.Outputs
{

    [OutputType]
    public sealed class SslStatesSslStateResult
    {
        /// <summary>
        /// The mongodb instance ID to query.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// Whetehr SSL is valid.
        /// </summary>
        public readonly bool IsValid;
        /// <summary>
        /// Whether SSL is enabled.
        /// </summary>
        public readonly bool SslEnable;
        /// <summary>
        /// The expire time of SSL.
        /// </summary>
        public readonly string SslExpiredTime;

        [OutputConstructor]
        private SslStatesSslStateResult(
            string instanceId,

            bool isValid,

            bool sslEnable,

            string sslExpiredTime)
        {
            InstanceId = instanceId;
            IsValid = isValid;
            SslEnable = sslEnable;
            SslExpiredTime = sslExpiredTime;
        }
    }
}