// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cloud_monitor.Inputs
{

    public sealed class EventRuleTlsTargetGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The project id.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The project name.
        /// </summary>
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        /// <summary>
        /// The Chinese region name.
        /// </summary>
        [Input("regionNameCn", required: true)]
        public Input<string> RegionNameCn { get; set; } = null!;

        /// <summary>
        /// The English region name.
        /// </summary>
        [Input("regionNameEn", required: true)]
        public Input<string> RegionNameEn { get; set; } = null!;

        /// <summary>
        /// The topic id.
        /// </summary>
        [Input("topicId", required: true)]
        public Input<string> TopicId { get; set; } = null!;

        public EventRuleTlsTargetGetArgs()
        {
        }
        public static new EventRuleTlsTargetGetArgs Empty => new EventRuleTlsTargetGetArgs();
    }
}