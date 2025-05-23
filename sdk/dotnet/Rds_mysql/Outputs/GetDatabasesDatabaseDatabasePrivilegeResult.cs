// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_mysql.Outputs
{

    [OutputType]
    public sealed class GetDatabasesDatabaseDatabasePrivilegeResult
    {
        /// <summary>
        /// The name of account.
        /// </summary>
        public readonly string AccountName;
        /// <summary>
        /// The privilege type of the account.
        /// </summary>
        public readonly string AccountPrivilege;
        /// <summary>
        /// The privilege detail of the account.
        /// </summary>
        public readonly string AccountPrivilegeDetail;

        [OutputConstructor]
        private GetDatabasesDatabaseDatabasePrivilegeResult(
            string accountName,

            string accountPrivilege,

            string accountPrivilegeDetail)
        {
            AccountName = accountName;
            AccountPrivilege = accountPrivilege;
            AccountPrivilegeDetail = accountPrivilegeDetail;
        }
    }
}
