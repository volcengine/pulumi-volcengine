// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Veenedge.Outputs
{

    [OutputType]
    public sealed class GetInstanceTypesInstanceTypeConfigStorageResult
    {
        /// <summary>
        /// The amount of local storage.
        /// </summary>
        public readonly int LocalStorageAmount;
        /// <summary>
        /// The capacity of local storage.
        /// </summary>
        public readonly int LocalStorageCapacity;
        /// <summary>
        /// The local storage category.
        /// </summary>
        public readonly string LocalStorageCategory;
        /// <summary>
        /// The unit of local storage.
        /// </summary>
        public readonly string LocalStorageUnit;

        [OutputConstructor]
        private GetInstanceTypesInstanceTypeConfigStorageResult(
            int localStorageAmount,

            int localStorageCapacity,

            string localStorageCategory,

            string localStorageUnit)
        {
            LocalStorageAmount = localStorageAmount;
            LocalStorageCapacity = localStorageCapacity;
            LocalStorageCategory = localStorageCategory;
            LocalStorageUnit = localStorageUnit;
        }
    }
}
