// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Kms.Inputs
{

    public sealed class GetKeysTagArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The key of the tag.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("values", required: true)]
        private List<string>? _values;

        /// <summary>
        /// The values of the tag.
        /// </summary>
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetKeysTagArgs()
        {
        }
        public static new GetKeysTagArgs Empty => new GetKeysTagArgs();
    }
}
