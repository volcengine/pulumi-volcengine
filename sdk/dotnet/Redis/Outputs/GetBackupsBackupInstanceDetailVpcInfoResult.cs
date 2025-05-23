// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Redis.Outputs
{

    [OutputType]
    public sealed class GetBackupsBackupInstanceDetailVpcInfoResult
    {
        /// <summary>
        /// Id of vpc.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of vpc.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetBackupsBackupInstanceDetailVpcInfoResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}
