// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls.Outputs
{

    [OutputType]
    public sealed class ShippersShipperContentInfoResult
    {
        /// <summary>
        /// CSV format log content configuration.
        /// </summary>
        public readonly Outputs.ShippersShipperContentInfoCsvInfoResult CsvInfo;
        /// <summary>
        /// Log content parsing format.
        /// </summary>
        public readonly string Format;
        /// <summary>
        /// JSON format log content configuration.
        /// </summary>
        public readonly Outputs.ShippersShipperContentInfoJsonInfoResult JsonInfo;

        [OutputConstructor]
        private ShippersShipperContentInfoResult(
            Outputs.ShippersShipperContentInfoCsvInfoResult csvInfo,

            string format,

            Outputs.ShippersShipperContentInfoJsonInfoResult jsonInfo)
        {
            CsvInfo = csvInfo;
            Format = format;
            JsonInfo = jsonInfo;
        }
    }
}
