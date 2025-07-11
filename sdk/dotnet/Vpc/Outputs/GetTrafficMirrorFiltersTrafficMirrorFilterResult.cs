// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpc.Outputs
{

    [OutputType]
    public sealed class GetTrafficMirrorFiltersTrafficMirrorFilterResult
    {
        /// <summary>
        /// The create time of traffic mirror filter rule.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The description of traffic mirror filter rule.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The ingress filter rules of traffic mirror filter.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTrafficMirrorFiltersTrafficMirrorFilterEgressFilterRuleResult> EgressFilterRules;
        /// <summary>
        /// The ID of traffic mirror filter.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ingress filter rules of traffic mirror filter.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTrafficMirrorFiltersTrafficMirrorFilterIngressFilterRuleResult> IngressFilterRules;
        /// <summary>
        /// The project name of traffic mirror filter.
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// The status of traffic mirror filter.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTrafficMirrorFiltersTrafficMirrorFilterTagResult> Tags;
        /// <summary>
        /// The ID of traffic mirror filter.
        /// </summary>
        public readonly string TrafficMirrorFilterId;
        /// <summary>
        /// The name of traffic mirror filter.
        /// </summary>
        public readonly string TrafficMirrorFilterName;
        /// <summary>
        /// The last update time of traffic mirror filter.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetTrafficMirrorFiltersTrafficMirrorFilterResult(
            string createdAt,

            string description,

            ImmutableArray<Outputs.GetTrafficMirrorFiltersTrafficMirrorFilterEgressFilterRuleResult> egressFilterRules,

            string id,

            ImmutableArray<Outputs.GetTrafficMirrorFiltersTrafficMirrorFilterIngressFilterRuleResult> ingressFilterRules,

            string projectName,

            string status,

            ImmutableArray<Outputs.GetTrafficMirrorFiltersTrafficMirrorFilterTagResult> tags,

            string trafficMirrorFilterId,

            string trafficMirrorFilterName,

            string updatedAt)
        {
            CreatedAt = createdAt;
            Description = description;
            EgressFilterRules = egressFilterRules;
            Id = id;
            IngressFilterRules = ingressFilterRules;
            ProjectName = projectName;
            Status = status;
            Tags = tags;
            TrafficMirrorFilterId = trafficMirrorFilterId;
            TrafficMirrorFilterName = trafficMirrorFilterName;
            UpdatedAt = updatedAt;
        }
    }
}
