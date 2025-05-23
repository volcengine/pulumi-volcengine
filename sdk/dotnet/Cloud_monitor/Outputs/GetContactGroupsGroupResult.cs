// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cloud_monitor.Outputs
{

    [OutputType]
    public sealed class GetContactGroupsGroupResult
    {
        /// <summary>
        /// The id of the account.
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// Contact information in the contact group.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetContactGroupsGroupContactResult> Contacts;
        /// <summary>
        /// The create time.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The description of the contact group.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The id of the contact group.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Search for keywords in contact group names, supports fuzzy search.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The update time.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetContactGroupsGroupResult(
            string accountId,

            ImmutableArray<Outputs.GetContactGroupsGroupContactResult> contacts,

            string createdAt,

            string description,

            string id,

            string name,

            string updatedAt)
        {
            AccountId = accountId;
            Contacts = contacts;
            CreatedAt = createdAt;
            Description = description;
            Id = id;
            Name = name;
            UpdatedAt = updatedAt;
        }
    }
}
