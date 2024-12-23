// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds.Inputs
{

    public sealed class AccountPrivilegeDbPrivilegeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The privilege type of the account.
        /// </summary>
        [Input("accountPrivilege", required: true)]
        public Input<string> AccountPrivilege { get; set; } = null!;

        /// <summary>
        /// The privilege string of the account.
        /// </summary>
        [Input("accountPrivilegeStr")]
        public Input<string>? AccountPrivilegeStr { get; set; }

        /// <summary>
        /// The name of database.
        /// </summary>
        [Input("dbName", required: true)]
        public Input<string> DbName { get; set; } = null!;

        public AccountPrivilegeDbPrivilegeGetArgs()
        {
        }
        public static new AccountPrivilegeDbPrivilegeGetArgs Empty => new AccountPrivilegeDbPrivilegeGetArgs();
    }
}
