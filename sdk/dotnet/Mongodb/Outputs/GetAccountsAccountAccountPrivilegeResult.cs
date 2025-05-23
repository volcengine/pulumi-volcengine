// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Mongodb.Outputs
{

    [OutputType]
    public sealed class GetAccountsAccountAccountPrivilegeResult
    {
        /// <summary>
        /// The Name of DB.
        /// </summary>
        public readonly string DbName;
        /// <summary>
        /// The Name of role.
        /// </summary>
        public readonly string RoleName;

        [OutputConstructor]
        private GetAccountsAccountAccountPrivilegeResult(
            string dbName,

            string roleName)
        {
            DbName = dbName;
            RoleName = roleName;
        }
    }
}
