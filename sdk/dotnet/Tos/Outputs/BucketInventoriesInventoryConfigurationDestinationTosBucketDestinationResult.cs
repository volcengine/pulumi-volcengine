// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tos.Outputs
{

    [OutputType]
    public sealed class BucketInventoriesInventoryConfigurationDestinationTosBucketDestinationResult
    {
        /// <summary>
        /// The account id of the destination tos bucket.
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// The name of the destination tos bucket.
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// The format of the bucket inventory. Valid values: `CSV`.
        /// </summary>
        public readonly string Format;
        /// <summary>
        /// The prefix matching information of the exported object. If not set, a list of all objects in the bucket will be generated by default.
        /// </summary>
        public readonly string Prefix;
        /// <summary>
        /// The role name used to grant object storage access to read all files from the source bucket and write files to the destination bucket.
        /// </summary>
        public readonly string Role;

        [OutputConstructor]
        private BucketInventoriesInventoryConfigurationDestinationTosBucketDestinationResult(
            string accountId,

            string bucket,

            string format,

            string prefix,

            string role)
        {
            AccountId = accountId;
            Bucket = bucket;
            Format = format;
            Prefix = prefix;
            Role = role;
        }
    }
}
