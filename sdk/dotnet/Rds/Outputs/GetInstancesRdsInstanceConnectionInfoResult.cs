// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds.Outputs
{

    [OutputType]
    public sealed class GetInstancesRdsInstanceConnectionInfoResult
    {
        /// <summary>
        /// Whether global read-only is enabled.
        /// </summary>
        public readonly string EnableReadOnly;
        /// <summary>
        /// Whether read-write separation is enabled.
        /// </summary>
        public readonly string EnableReadWriteSplitting;
        /// <summary>
        /// The internal domain of the RDS instance.
        /// </summary>
        public readonly string InternalDomain;
        /// <summary>
        /// The interval port of the RDS instance.
        /// </summary>
        public readonly string InternalPort;
        /// <summary>
        /// The public domain of the RDS instance.
        /// </summary>
        public readonly string PublicDomain;
        /// <summary>
        /// The public port of the RDS instance.
        /// </summary>
        public readonly string PublicPort;

        [OutputConstructor]
        private GetInstancesRdsInstanceConnectionInfoResult(
            string enableReadOnly,

            string enableReadWriteSplitting,

            string internalDomain,

            string internalPort,

            string publicDomain,

            string publicPort)
        {
            EnableReadOnly = enableReadOnly;
            EnableReadWriteSplitting = enableReadWriteSplitting;
            InternalDomain = internalDomain;
            InternalPort = internalPort;
            PublicDomain = publicDomain;
            PublicPort = publicPort;
        }
    }
}
