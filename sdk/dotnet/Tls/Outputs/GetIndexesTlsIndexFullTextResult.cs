// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls.Outputs
{

    [OutputType]
    public sealed class GetIndexesTlsIndexFullTextResult
    {
        /// <summary>
        /// Whether the value is case sensitive.
        /// </summary>
        public readonly bool CaseSensitive;
        /// <summary>
        /// The delimiter of the value.
        /// </summary>
        public readonly string Delimiter;
        /// <summary>
        /// Whether the value include chinese.
        /// </summary>
        public readonly bool IncludeChinese;

        [OutputConstructor]
        private GetIndexesTlsIndexFullTextResult(
            bool caseSensitive,

            string delimiter,

            bool includeChinese)
        {
            CaseSensitive = caseSensitive;
            Delimiter = delimiter;
            IncludeChinese = includeChinese;
        }
    }
}
