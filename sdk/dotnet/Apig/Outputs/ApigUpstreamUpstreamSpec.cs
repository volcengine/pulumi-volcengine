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
    public sealed class ApigUpstreamUpstreamSpec
    {
        /// <summary>
        /// The ai provider of apig upstream.
        /// </summary>
        public readonly Outputs.ApigUpstreamUpstreamSpecAiProvider? AiProvider;
        /// <summary>
        /// The domain of apig upstream.
        /// </summary>
        public readonly Outputs.ApigUpstreamUpstreamSpecDomain? Domain;
        /// <summary>
        /// The ecs list of apig upstream.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApigUpstreamUpstreamSpecEcsList> EcsLists;
        /// <summary>
        /// The fixed ip list of apig upstream.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApigUpstreamUpstreamSpecFixedIpList> FixedIpLists;
        /// <summary>
        /// The k8s service of apig upstream.
        /// </summary>
        public readonly Outputs.ApigUpstreamUpstreamSpecK8sService? K8sService;
        /// <summary>
        /// The nacos service of apig upstream.
        /// </summary>
        public readonly Outputs.ApigUpstreamUpstreamSpecNacosService? NacosService;
        /// <summary>
        /// The vefaas of apig upstream.
        /// </summary>
        public readonly Outputs.ApigUpstreamUpstreamSpecVeFaas? VeFaas;
        /// <summary>
        /// The mlp of apig upstream.
        /// </summary>
        public readonly Outputs.ApigUpstreamUpstreamSpecVeMlp? VeMlp;

        [OutputConstructor]
        private ApigUpstreamUpstreamSpec(
            Outputs.ApigUpstreamUpstreamSpecAiProvider? aiProvider,

            Outputs.ApigUpstreamUpstreamSpecDomain? domain,

            ImmutableArray<Outputs.ApigUpstreamUpstreamSpecEcsList> ecsLists,

            ImmutableArray<Outputs.ApigUpstreamUpstreamSpecFixedIpList> fixedIpLists,

            Outputs.ApigUpstreamUpstreamSpecK8sService? k8sService,

            Outputs.ApigUpstreamUpstreamSpecNacosService? nacosService,

            Outputs.ApigUpstreamUpstreamSpecVeFaas? veFaas,

            Outputs.ApigUpstreamUpstreamSpecVeMlp? veMlp)
        {
            AiProvider = aiProvider;
            Domain = domain;
            EcsLists = ecsLists;
            FixedIpLists = fixedIpLists;
            K8sService = k8sService;
            NacosService = nacosService;
            VeFaas = veFaas;
            VeMlp = veMlp;
        }
    }
}
