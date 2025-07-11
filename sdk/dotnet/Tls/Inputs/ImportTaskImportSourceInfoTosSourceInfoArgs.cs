// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls.Inputs
{

    public sealed class ImportTaskImportSourceInfoTosSourceInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The TOS bucket where the log file is located.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// The compression mode of data in the TOS bucket.
        /// </summary>
        [Input("compressType")]
        public Input<string>? CompressType { get; set; }

        /// <summary>
        /// The path of the file to be imported in the TOS bucket.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// The region where the TOS bucket is located. Support cross-regional data import.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ImportTaskImportSourceInfoTosSourceInfoArgs()
        {
        }
        public static new ImportTaskImportSourceInfoTosSourceInfoArgs Empty => new ImportTaskImportSourceInfoTosSourceInfoArgs();
    }
}
