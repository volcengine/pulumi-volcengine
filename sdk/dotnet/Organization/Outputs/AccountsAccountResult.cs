// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Organization.Outputs
{

    [OutputType]
    public sealed class AccountsAccountResult
    {
        /// <summary>
        /// The id of the account.
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// The name of the account.
        /// </summary>
        public readonly string AccountName;
        /// <summary>
        /// Whether to allow the account enable console. `0` means allowed, `1` means not allowed.
        /// </summary>
        public readonly int AllowConsole;
        /// <summary>
        /// Whether to allow exit the organization. `0` means allowed, `1` means not allowed.
        /// </summary>
        public readonly int AllowExit;
        /// <summary>
        /// The created time of the account.
        /// </summary>
        public readonly string CreatedTime;
        /// <summary>
        /// The delete uk of the account.
        /// </summary>
        public readonly string DeleteUk;
        /// <summary>
        /// The deleted time of the account.
        /// </summary>
        public readonly string DeletedTime;
        /// <summary>
        /// The description of the account.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The name of the iam role.
        /// </summary>
        public readonly string IamRole;
        /// <summary>
        /// The id of the account.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Whether the account is owner. `0` means not owner, `1` means owner.
        /// </summary>
        public readonly int IsOwner;
        /// <summary>
        /// The join type of the account. `0` means create, `1` means invitation.
        /// </summary>
        public readonly int JoinType;
        /// <summary>
        /// The id of the organization.
        /// </summary>
        public readonly string OrgId;
        /// <summary>
        /// The type of the organization. `1` means business organization.
        /// </summary>
        public readonly int OrgType;
        /// <summary>
        /// The id of the organization unit.
        /// </summary>
        public readonly string OrgUnitId;
        /// <summary>
        /// The name of the organization unit.
        /// </summary>
        public readonly string OrgUnitName;
        /// <summary>
        /// The id of the organization verification.
        /// </summary>
        public readonly string OrgVerificationId;
        /// <summary>
        /// The owner id of the account.
        /// </summary>
        public readonly string Owner;
        /// <summary>
        /// The show name of the account.
        /// </summary>
        public readonly string ShowName;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.AccountsAccountTagResult> Tags;
        /// <summary>
        /// The updated time of the account.
        /// </summary>
        public readonly string UpdatedTime;

        [OutputConstructor]
        private AccountsAccountResult(
            string accountId,

            string accountName,

            int allowConsole,

            int allowExit,

            string createdTime,

            string deleteUk,

            string deletedTime,

            string description,

            string iamRole,

            string id,

            int isOwner,

            int joinType,

            string orgId,

            int orgType,

            string orgUnitId,

            string orgUnitName,

            string orgVerificationId,

            string owner,

            string showName,

            ImmutableArray<Outputs.AccountsAccountTagResult> tags,

            string updatedTime)
        {
            AccountId = accountId;
            AccountName = accountName;
            AllowConsole = allowConsole;
            AllowExit = allowExit;
            CreatedTime = createdTime;
            DeleteUk = deleteUk;
            DeletedTime = deletedTime;
            Description = description;
            IamRole = iamRole;
            Id = id;
            IsOwner = isOwner;
            JoinType = joinType;
            OrgId = orgId;
            OrgType = orgType;
            OrgUnitId = orgUnitId;
            OrgUnitName = orgUnitName;
            OrgVerificationId = orgVerificationId;
            Owner = owner;
            ShowName = showName;
            Tags = tags;
            UpdatedTime = updatedTime;
        }
    }
}