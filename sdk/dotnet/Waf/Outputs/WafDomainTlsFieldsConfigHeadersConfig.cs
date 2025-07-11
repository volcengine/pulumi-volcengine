// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Waf.Outputs
{

    [OutputType]
    public sealed class WafDomainTlsFieldsConfigHeadersConfig
    {
        /// <summary>
        /// Whether the log contains this field. Works only on modified scenes.
        /// </summary>
        public readonly int? Enable;
        /// <summary>
        /// For the use of composite fields, exclude the fields in the keyword list from the JSON of the fields. Works only on modified scenes.
        /// </summary>
        public readonly ImmutableArray<string> ExcludedKeyLists;
        /// <summary>
        /// Create statistical indexes for the fields of the list. Works only on modified scenes.
        /// </summary>
        public readonly ImmutableArray<string> StatisticalKeyLists;

        [OutputConstructor]
        private WafDomainTlsFieldsConfigHeadersConfig(
            int? enable,

            ImmutableArray<string> excludedKeyLists,

            ImmutableArray<string> statisticalKeyLists)
        {
            Enable = enable;
            ExcludedKeyLists = excludedKeyLists;
            StatisticalKeyLists = statisticalKeyLists;
        }
    }
}
