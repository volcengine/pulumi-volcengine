// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vefaas.Inputs
{

    public sealed class FunctionNasStorageGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable NAS storage mounting.
        /// </summary>
        [Input("enableNas", required: true)]
        public Input<bool> EnableNas { get; set; } = null!;

        [Input("nasConfigs")]
        private InputList<Inputs.FunctionNasStorageNasConfigGetArgs>? _nasConfigs;

        /// <summary>
        /// The configuration of NAS.
        /// </summary>
        public InputList<Inputs.FunctionNasStorageNasConfigGetArgs> NasConfigs
        {
            get => _nasConfigs ?? (_nasConfigs = new InputList<Inputs.FunctionNasStorageNasConfigGetArgs>());
            set => _nasConfigs = value;
        }

        public FunctionNasStorageGetArgs()
        {
        }
        public static new FunctionNasStorageGetArgs Empty => new FunctionNasStorageGetArgs();
    }
}
