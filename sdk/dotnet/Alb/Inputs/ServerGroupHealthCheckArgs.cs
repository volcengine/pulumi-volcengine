// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Alb.Inputs
{

    public sealed class ServerGroupHealthCheckArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain of health check.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The enable status of health check function. Valid values: `on`, `off`. Default is `on`.
        /// </summary>
        [Input("enabled")]
        public Input<string>? Enabled { get; set; }

        /// <summary>
        /// The healthy threshold of health check. Valid value range in 2~10. Default is 3.
        /// </summary>
        [Input("healthyThreshold")]
        public Input<int>? HealthyThreshold { get; set; }

        /// <summary>
        /// The normal http status code of health check, the value can be `http_2xx` or `http_3xx` or `http_4xx` or `http_5xx`.
        /// </summary>
        [Input("httpCode")]
        public Input<string>? HttpCode { get; set; }

        /// <summary>
        /// The http version of health check. Valid values: `HTTP1.0`, `HTTP1.1`. Default is `HTTP1.0`.
        /// </summary>
        [Input("httpVersion")]
        public Input<string>? HttpVersion { get; set; }

        /// <summary>
        /// The interval executing health check. Unit: second. Valid value range in 1~300. Default is 2.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// The method of health check. Valid values: `GET` or `HEAD`. Default is `HEAD`.
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        /// <summary>
        /// The response timeout of health check. Unit: second. Valid value range in 1~60. Default is 2.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The unhealthy threshold of health check. Valid value range in 2~10. Default is 3.
        /// </summary>
        [Input("unhealthyThreshold")]
        public Input<int>? UnhealthyThreshold { get; set; }

        /// <summary>
        /// The uri of health check.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public ServerGroupHealthCheckArgs()
        {
        }
        public static new ServerGroupHealthCheckArgs Empty => new ServerGroupHealthCheckArgs();
    }
}