// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vefaas.Inputs
{

    public sealed class GetReleasesFilterInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Filter key enumeration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// The filtering value of the query.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public GetReleasesFilterInputArgs()
        {
        }
        public static new GetReleasesFilterInputArgs Empty => new GetReleasesFilterInputArgs();
    }
}
