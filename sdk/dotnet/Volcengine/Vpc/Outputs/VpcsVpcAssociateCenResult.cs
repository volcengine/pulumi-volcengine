// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Vpc.Outputs
{

    [OutputType]
    public sealed class VpcsVpcAssociateCenResult
    {
        /// <summary>
        /// The ID of CEN.
        /// </summary>
        public readonly string CenId;
        /// <summary>
        /// The owner ID of CEN.
        /// </summary>
        public readonly string CenOwnerId;
        /// <summary>
        /// The status of CEN.
        /// </summary>
        public readonly string CenStatus;

        [OutputConstructor]
        private VpcsVpcAssociateCenResult(
            string cenId,

            string cenOwnerId,

            string cenStatus)
        {
            CenId = cenId;
            CenOwnerId = cenOwnerId;
            CenStatus = cenStatus;
        }
    }
}