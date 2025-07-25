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
    public sealed class WafDomainTlsFieldsConfig
    {
        /// <summary>
        /// The configuration of Headers. Works only on modified scenes.
        /// </summary>
        public readonly Outputs.WafDomainTlsFieldsConfigHeadersConfig? HeadersConfig;

        [OutputConstructor]
        private WafDomainTlsFieldsConfig(Outputs.WafDomainTlsFieldsConfigHeadersConfig? headersConfig)
        {
            HeadersConfig = headersConfig;
        }
    }
}
