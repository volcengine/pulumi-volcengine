// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Nas.Outputs
{

    [OutputType]
    public sealed class GetPermissionGroupsFilterResult
    {
        /// <summary>
        /// Filters permission groups for specified characteristics based on attributes. The parameters that support filtering are as follows: `PermissionGroupName`, `PermissionGroupId`.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The value of the filter item.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetPermissionGroupsFilterResult(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
