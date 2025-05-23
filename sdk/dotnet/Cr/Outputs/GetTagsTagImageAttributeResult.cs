// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cr.Outputs
{

    [OutputType]
    public sealed class GetTagsTagImageAttributeResult
    {
        /// <summary>
        /// The image architecture.
        /// </summary>
        public readonly string Architecture;
        /// <summary>
        /// The image author.
        /// </summary>
        public readonly string Author;
        /// <summary>
        /// The digest of image.
        /// </summary>
        public readonly string Digest;
        /// <summary>
        /// The iamge os.
        /// </summary>
        public readonly string Os;

        [OutputConstructor]
        private GetTagsTagImageAttributeResult(
            string architecture,

            string author,

            string digest,

            string os)
        {
            Architecture = architecture;
            Author = author;
            Digest = digest;
            Os = os;
        }
    }
}
