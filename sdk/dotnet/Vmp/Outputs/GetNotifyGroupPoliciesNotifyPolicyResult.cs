// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vmp.Outputs
{

    [OutputType]
    public sealed class GetNotifyGroupPoliciesNotifyPolicyResult
    {
        /// <summary>
        /// The create time of notify group policy.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The description of notify group policy.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The id of the notify group policy.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The levels of the notify group policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNotifyGroupPoliciesNotifyPolicyLevelResult> Levels;
        /// <summary>
        /// The name of notify group policy.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetNotifyGroupPoliciesNotifyPolicyResult(
            string createTime,

            string description,

            string id,

            ImmutableArray<Outputs.GetNotifyGroupPoliciesNotifyPolicyLevelResult> levels,

            string name)
        {
            CreateTime = createTime;
            Description = description;
            Id = id;
            Levels = levels;
            Name = name;
        }
    }
}
