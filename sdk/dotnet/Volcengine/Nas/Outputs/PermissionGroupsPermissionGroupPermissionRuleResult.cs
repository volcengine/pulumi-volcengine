// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Nas.Outputs
{

    [OutputType]
    public sealed class PermissionGroupsPermissionGroupPermissionRuleResult
    {
        /// <summary>
        /// Client IP addresses that are allowed access.
        /// </summary>
        public readonly string CidrIp;
        /// <summary>
        /// The id of the permission rule.
        /// </summary>
        public readonly string PermissionRuleId;
        /// <summary>
        /// Permission group read and write rules.
        /// </summary>
        public readonly string RwMode;
        /// <summary>
        /// Permission group user permissions.
        /// </summary>
        public readonly string UserMode;

        [OutputConstructor]
        private PermissionGroupsPermissionGroupPermissionRuleResult(
            string cidrIp,

            string permissionRuleId,

            string rwMode,

            string userMode)
        {
            CidrIp = cidrIp;
            PermissionRuleId = permissionRuleId;
            RwMode = rwMode;
            UserMode = userMode;
        }
    }
}