// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Waf.Inputs
{

    public sealed class AclRuleAccurateGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("accurateRules", required: true)]
        private InputList<Inputs.AclRuleAccurateGroupAccurateRuleArgs>? _accurateRules;

        /// <summary>
        /// Details of advanced conditions.
        /// </summary>
        public InputList<Inputs.AclRuleAccurateGroupAccurateRuleArgs> AccurateRules
        {
            get => _accurateRules ?? (_accurateRules = new InputList<Inputs.AclRuleAccurateGroupAccurateRuleArgs>());
            set => _accurateRules = value;
        }

        /// <summary>
        /// The logical relationship of advanced conditions.
        /// </summary>
        [Input("logic", required: true)]
        public Input<int> Logic { get; set; } = null!;

        public AclRuleAccurateGroupArgs()
        {
        }
        public static new AclRuleAccurateGroupArgs Empty => new AclRuleAccurateGroupArgs();
    }
}
