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
    public sealed class GetUpstreamsUpstreamResult
    {
        /// <summary>
        /// The backend target list of apig upstream.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUpstreamsUpstreamBackendTargetListResult> BackendTargetLists;
        /// <summary>
        /// The circuit breaking settings of apig upstream.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUpstreamsUpstreamCircuitBreakingSettingResult> CircuitBreakingSettings;
        /// <summary>
        /// The comments of apig upstream.
        /// </summary>
        public readonly string Comments;
        /// <summary>
        /// The create time of apig upstream.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The id of api gateway.
        /// </summary>
        public readonly string GatewayId;
        /// <summary>
        /// The id of apig upstream.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The load balancer settings of apig upstream.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUpstreamsUpstreamLoadBalancerSettingResult> LoadBalancerSettings;
        /// <summary>
        /// The name of apig upstream. This field support fuzzy query.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The protocol of apig upstream.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// The resource type of apig upstream. Valid values: `Console`, `Ingress`.
        /// </summary>
        public readonly string ResourceType;
        /// <summary>
        /// The source type of apig upstream. Valid values: `VeFaas`, `ECS`, `FixedIP`, `K8S`, `Nacos`, `Domain`, `AIProvider`, `VeMLP`.
        /// </summary>
        public readonly string SourceType;
        /// <summary>
        /// The tls settings of apig upstream.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUpstreamsUpstreamTlsSettingResult> TlsSettings;
        /// <summary>
        /// The update time of apig upstream version.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// The upstream spec of apig upstream.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUpstreamsUpstreamUpstreamSpecResult> UpstreamSpecs;
        /// <summary>
        /// The version details of apig upstream.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUpstreamsUpstreamVersionDetailResult> VersionDetails;

        [OutputConstructor]
        private GetUpstreamsUpstreamResult(
            ImmutableArray<Outputs.GetUpstreamsUpstreamBackendTargetListResult> backendTargetLists,

            ImmutableArray<Outputs.GetUpstreamsUpstreamCircuitBreakingSettingResult> circuitBreakingSettings,

            string comments,

            string createTime,

            string gatewayId,

            string id,

            ImmutableArray<Outputs.GetUpstreamsUpstreamLoadBalancerSettingResult> loadBalancerSettings,

            string name,

            string protocol,

            string resourceType,

            string sourceType,

            ImmutableArray<Outputs.GetUpstreamsUpstreamTlsSettingResult> tlsSettings,

            string updateTime,

            ImmutableArray<Outputs.GetUpstreamsUpstreamUpstreamSpecResult> upstreamSpecs,

            ImmutableArray<Outputs.GetUpstreamsUpstreamVersionDetailResult> versionDetails)
        {
            BackendTargetLists = backendTargetLists;
            CircuitBreakingSettings = circuitBreakingSettings;
            Comments = comments;
            CreateTime = createTime;
            GatewayId = gatewayId;
            Id = id;
            LoadBalancerSettings = loadBalancerSettings;
            Name = name;
            Protocol = protocol;
            ResourceType = resourceType;
            SourceType = sourceType;
            TlsSettings = tlsSettings;
            UpdateTime = updateTime;
            UpstreamSpecs = upstreamSpecs;
            VersionDetails = versionDetails;
        }
    }
}
