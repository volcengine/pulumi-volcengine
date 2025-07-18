// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Apig.Outputs
{

    [OutputType]
    public sealed class GetGatewayServicesGatewayServiceResult
    {
        /// <summary>
        /// The auth spec of the api gateway service.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGatewayServicesGatewayServiceAuthSpecResult> AuthSpecs;
        /// <summary>
        /// The comments of the api gateway service.
        /// </summary>
        public readonly string Comments;
        /// <summary>
        /// The create time of the api gateway service.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The custom domains of the api gateway service.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGatewayServicesGatewayServiceCustomDomainResult> CustomDomains;
        /// <summary>
        /// The domains of the api gateway service.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGatewayServicesGatewayServiceDomainResult> Domains;
        /// <summary>
        /// The gateway id of api gateway service.
        /// </summary>
        public readonly string GatewayId;
        /// <summary>
        /// The gateway name of the api gateway service.
        /// </summary>
        public readonly string GatewayName;
        /// <summary>
        /// The Id of the api gateway service.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The error message of the api gateway service.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// The name of api gateway service. This field support fuzzy query.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The protocol of the api gateway service.
        /// </summary>
        public readonly ImmutableArray<string> Protocols;
        /// <summary>
        /// The status of api gateway service.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetGatewayServicesGatewayServiceResult(
            ImmutableArray<Outputs.GetGatewayServicesGatewayServiceAuthSpecResult> authSpecs,

            string comments,

            string createTime,

            ImmutableArray<Outputs.GetGatewayServicesGatewayServiceCustomDomainResult> customDomains,

            ImmutableArray<Outputs.GetGatewayServicesGatewayServiceDomainResult> domains,

            string gatewayId,

            string gatewayName,

            string id,

            string message,

            string name,

            ImmutableArray<string> protocols,

            string status)
        {
            AuthSpecs = authSpecs;
            Comments = comments;
            CreateTime = createTime;
            CustomDomains = customDomains;
            Domains = domains;
            GatewayId = gatewayId;
            GatewayName = gatewayName;
            Id = id;
            Message = message;
            Name = name;
            Protocols = protocols;
            Status = status;
        }
    }
}
