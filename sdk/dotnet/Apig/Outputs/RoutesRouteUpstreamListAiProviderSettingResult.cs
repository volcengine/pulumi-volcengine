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
    public sealed class RoutesRouteUpstreamListAiProviderSettingResult
    {
        /// <summary>
        /// The model of the ai provider.
        /// </summary>
        public readonly string Model;
        /// <summary>
        /// The target path of the ai provider.
        /// </summary>
        public readonly string TargetPath;

        [OutputConstructor]
        private RoutesRouteUpstreamListAiProviderSettingResult(
            string model,

            string targetPath)
        {
            Model = model;
            TargetPath = targetPath;
        }
    }
}
