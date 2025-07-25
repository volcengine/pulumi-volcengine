// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vefaas.Inputs
{

    public sealed class FunctionTosMountConfigCredentialsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessKeyId", required: true)]
        private Input<string>? _accessKeyId;

        /// <summary>
        /// The AccessKey ID (AK) of the Volcano Engine account.
        /// </summary>
        public Input<string>? AccessKeyId
        {
            get => _accessKeyId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessKeyId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("secretAccessKey", required: true)]
        private Input<string>? _secretAccessKey;

        /// <summary>
        /// The Secret Access Key (SK) of the Volcano Engine account.
        /// </summary>
        public Input<string>? SecretAccessKey
        {
            get => _secretAccessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretAccessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public FunctionTosMountConfigCredentialsGetArgs()
        {
        }
        public static new FunctionTosMountConfigCredentialsGetArgs Empty => new FunctionTosMountConfigCredentialsGetArgs();
    }
}
