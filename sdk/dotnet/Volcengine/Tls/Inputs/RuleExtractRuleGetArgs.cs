// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Tls.Inputs
{

    public sealed class RuleExtractRuleGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The first log line needs to match the regular expression.
        /// </summary>
        [Input("beginRegex")]
        public Input<string>? BeginRegex { get; set; }

        /// <summary>
        /// The delimiter of the log.
        /// </summary>
        [Input("delimiter")]
        public Input<string>? Delimiter { get; set; }

        [Input("filterKeyRegexes")]
        private InputList<Inputs.RuleExtractRuleFilterKeyRegexGetArgs>? _filterKeyRegexes;

        /// <summary>
        /// The filter key list.
        /// </summary>
        public InputList<Inputs.RuleExtractRuleFilterKeyRegexGetArgs> FilterKeyRegexes
        {
            get => _filterKeyRegexes ?? (_filterKeyRegexes = new InputList<Inputs.RuleExtractRuleFilterKeyRegexGetArgs>());
            set => _filterKeyRegexes = value;
        }

        [Input("keys")]
        private InputList<string>? _keys;

        /// <summary>
        /// A list of log field names (Key).
        /// </summary>
        public InputList<string> Keys
        {
            get => _keys ?? (_keys = new InputList<string>());
            set => _keys = value;
        }

        /// <summary>
        /// The entire log needs to match the regular expression.
        /// </summary>
        [Input("logRegex")]
        public Input<string>? LogRegex { get; set; }

        /// <summary>
        /// Automatically extract log fields according to the specified log template.
        /// </summary>
        [Input("logTemplate")]
        public Input<Inputs.RuleExtractRuleLogTemplateGetArgs>? LogTemplate { get; set; }

        /// <summary>
        /// Parsing format of the time field.
        /// </summary>
        [Input("timeFormat")]
        public Input<string>? TimeFormat { get; set; }

        /// <summary>
        /// The field name of the log time field.
        /// </summary>
        [Input("timeKey")]
        public Input<string>? TimeKey { get; set; }

        /// <summary>
        /// When uploading the failed log, the key name of the failed log.
        /// </summary>
        [Input("unMatchLogKey")]
        public Input<string>? UnMatchLogKey { get; set; }

        /// <summary>
        /// Whether to upload the log of parsing failure.
        /// </summary>
        [Input("unMatchUpLoadSwitch")]
        public Input<bool>? UnMatchUpLoadSwitch { get; set; }

        public RuleExtractRuleGetArgs()
        {
        }
    }
}