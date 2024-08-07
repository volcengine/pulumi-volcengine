// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls.Inputs
{

    public sealed class IndexUserInnerKeyValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the value is case sensitive.
        /// </summary>
        [Input("caseSensitive")]
        public Input<bool>? CaseSensitive { get; set; }

        /// <summary>
        /// The delimiter of the value.
        /// </summary>
        [Input("delimiter")]
        public Input<string>? Delimiter { get; set; }

        /// <summary>
        /// Whether the value include chinese.
        /// </summary>
        [Input("includeChinese")]
        public Input<bool>? IncludeChinese { get; set; }

        [Input("jsonKeys")]
        private InputList<Inputs.IndexUserInnerKeyValueJsonKeyArgs>? _jsonKeys;

        /// <summary>
        /// The JSON subfield key value index.
        /// </summary>
        public InputList<Inputs.IndexUserInnerKeyValueJsonKeyArgs> JsonKeys
        {
            get => _jsonKeys ?? (_jsonKeys = new InputList<Inputs.IndexUserInnerKeyValueJsonKeyArgs>());
            set => _jsonKeys = value;
        }

        /// <summary>
        /// The key of the KeyValueInfo.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Whether the filed is enabled for analysis.
        /// </summary>
        [Input("sqlFlag")]
        public Input<bool>? SqlFlag { get; set; }

        /// <summary>
        /// The type of value. Valid values: `long`, `double`, `text`, `json`.
        /// </summary>
        [Input("valueType", required: true)]
        public Input<string> ValueType { get; set; } = null!;

        public IndexUserInnerKeyValueArgs()
        {
        }
        public static new IndexUserInnerKeyValueArgs Empty => new IndexUserInnerKeyValueArgs();
    }
}
