// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Alb.Inputs
{

    public sealed class RuleRedirectConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The redirect domain, only support exact domain name.
        /// </summary>
        [Input("redirectDomain")]
        public Input<string>? RedirectDomain { get; set; }

        /// <summary>
        /// The redirect http code, support 301(default), 302, 307, 308.
        /// </summary>
        [Input("redirectHttpCode")]
        public Input<string>? RedirectHttpCode { get; set; }

        /// <summary>
        /// The redirect port.
        /// </summary>
        [Input("redirectPort")]
        public Input<string>? RedirectPort { get; set; }

        /// <summary>
        /// The redirect protocol, support HTTP, HTTPS(default).
        /// </summary>
        [Input("redirectProtocol")]
        public Input<string>? RedirectProtocol { get; set; }

        /// <summary>
        /// The redirect URI.
        /// </summary>
        [Input("redirectUri")]
        public Input<string>? RedirectUri { get; set; }

        public RuleRedirectConfigArgs()
        {
        }
        public static new RuleRedirectConfigArgs Empty => new RuleRedirectConfigArgs();
    }
}