// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls.Outputs
{

    [OutputType]
    public sealed class IndexesTlsIndexUserInnerKeyValueResult
    {
        /// <summary>
        /// Whether the value is case sensitive.
        /// </summary>
        public readonly bool CaseSensitive;
        /// <summary>
        /// The delimiter of the value.
        /// </summary>
        public readonly string Delimiter;
        /// <summary>
        /// Whether the value include chinese.
        /// </summary>
        public readonly bool IncludeChinese;
        /// <summary>
        /// The JSON subfield key value index.
        /// </summary>
        public readonly ImmutableArray<Outputs.IndexesTlsIndexUserInnerKeyValueJsonKeyResult> JsonKeys;
        /// <summary>
        /// The key of the KeyValue index.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Whether the filed is enabled for analysis.
        /// </summary>
        public readonly bool SqlFlag;
        /// <summary>
        /// The type of value.
        /// </summary>
        public readonly string ValueType;

        [OutputConstructor]
        private IndexesTlsIndexUserInnerKeyValueResult(
            bool caseSensitive,

            string delimiter,

            bool includeChinese,

            ImmutableArray<Outputs.IndexesTlsIndexUserInnerKeyValueJsonKeyResult> jsonKeys,

            string key,

            bool sqlFlag,

            string valueType)
        {
            CaseSensitive = caseSensitive;
            Delimiter = delimiter;
            IncludeChinese = includeChinese;
            JsonKeys = jsonKeys;
            Key = key;
            SqlFlag = sqlFlag;
            ValueType = valueType;
        }
    }
}
