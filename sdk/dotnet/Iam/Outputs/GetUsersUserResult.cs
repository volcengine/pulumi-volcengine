// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Iam.Outputs
{

    [OutputType]
    public sealed class GetUsersUserResult
    {
        /// <summary>
        /// Main account ID to which the sub-user belongs.
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// The create date of the user.
        /// </summary>
        public readonly string CreateDate;
        /// <summary>
        /// The description of the user.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The display name of the user.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The email of the user.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// Whether the email has been verified.
        /// </summary>
        public readonly bool EmailIsVerify;
        /// <summary>
        /// The mobile phone of the user.
        /// </summary>
        public readonly string MobilePhone;
        /// <summary>
        /// Whether the phone number has been verified.
        /// </summary>
        public readonly bool MobilePhoneIsVerify;
        /// <summary>
        /// The trn of the user.
        /// </summary>
        public readonly string Trn;
        /// <summary>
        /// The update date of the user.
        /// </summary>
        public readonly string UpdateDate;
        /// <summary>
        /// The id of the user.
        /// </summary>
        public readonly string UserId;
        /// <summary>
        /// The name of the user.
        /// </summary>
        public readonly string UserName;

        [OutputConstructor]
        private GetUsersUserResult(
            string accountId,

            string createDate,

            string description,

            string displayName,

            string email,

            bool emailIsVerify,

            string mobilePhone,

            bool mobilePhoneIsVerify,

            string trn,

            string updateDate,

            string userId,

            string userName)
        {
            AccountId = accountId;
            CreateDate = createDate;
            Description = description;
            DisplayName = displayName;
            Email = email;
            EmailIsVerify = emailIsVerify;
            MobilePhone = mobilePhone;
            MobilePhoneIsVerify = mobilePhoneIsVerify;
            Trn = trn;
            UpdateDate = updateDate;
            UserId = userId;
            UserName = userName;
        }
    }
}
