// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Alb.Outputs
{

    [OutputType]
    public sealed class CustomizedCfgsCfgListenerResult
    {
        /// <summary>
        /// The id of the listener.
        /// </summary>
        public readonly string ListenerId;
        /// <summary>
        /// The Name of Listener.
        /// </summary>
        public readonly string ListenerName;
        /// <summary>
        /// The port info of listener.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The protocol info of listener.
        /// </summary>
        public readonly string Protocol;

        [OutputConstructor]
        private CustomizedCfgsCfgListenerResult(
            string listenerId,

            string listenerName,

            int port,

            string protocol)
        {
            ListenerId = listenerId;
            ListenerName = listenerName;
            Port = port;
            Protocol = protocol;
        }
    }
}