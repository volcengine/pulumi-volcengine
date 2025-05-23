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
    public sealed class GetContactsContactResult
    {
        /// <summary>
        /// The email of contact.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// The ID of contact.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of contact.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The phone of contact.
        /// </summary>
        public readonly string Phone;

        [OutputConstructor]
        private GetContactsContactResult(
            string email,

            string id,

            string name,

            string phone)
        {
            Email = email;
            Id = id;
            Name = name;
            Phone = phone;
        }
    }
}
