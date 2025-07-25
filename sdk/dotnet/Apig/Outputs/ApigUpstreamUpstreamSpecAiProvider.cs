// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Apig.Outputs
{

    [OutputType]
    public sealed class ApigUpstreamUpstreamSpecAiProvider
    {
        /// <summary>
        /// The base url of ai provider.
        /// </summary>
        public readonly string BaseUrl;
        /// <summary>
        /// The custom body params of ai provider.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? CustomBodyParams;
        /// <summary>
        /// The custom header params of ai provider.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? CustomHeaderParams;
        /// <summary>
        /// The custom model service of ai provider.
        /// </summary>
        public readonly Outputs.ApigUpstreamUpstreamSpecAiProviderCustomModelService? CustomModelService;
        /// <summary>
        /// The name of ai provider.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The token of ai provider.
        /// </summary>
        public readonly string Token;

        [OutputConstructor]
        private ApigUpstreamUpstreamSpecAiProvider(
            string baseUrl,

            ImmutableDictionary<string, string>? customBodyParams,

            ImmutableDictionary<string, string>? customHeaderParams,

            Outputs.ApigUpstreamUpstreamSpecAiProviderCustomModelService? customModelService,

            string name,

            string token)
        {
            BaseUrl = baseUrl;
            CustomBodyParams = customBodyParams;
            CustomHeaderParams = customHeaderParams;
            CustomModelService = customModelService;
            Name = name;
            Token = token;
        }
    }
}
