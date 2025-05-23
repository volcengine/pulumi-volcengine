// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cen
{
    public static class GetBandwidthPackages
    {
        /// <summary>
        /// Use this data source to query detailed information of cen bandwidth packages
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var fooBandwidthPackage = new List&lt;Volcengine.Cen.BandwidthPackage&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; 2; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         fooBandwidthPackage.Add(new Volcengine.Cen.BandwidthPackage($"fooBandwidthPackage-{range.Value}", new()
        ///         {
        ///             LocalGeographicRegionSetId = "China",
        ///             PeerGeographicRegionSetId = "China",
        ///             Bandwidth = 2,
        ///             CenBandwidthPackageName = $"acc-test-cen-bp-{range.Value}",
        ///             Description = "acc-test",
        ///             BillingType = "PrePaid",
        ///             PeriodUnit = "Month",
        ///             Period = 1,
        ///             ProjectName = "default",
        ///             Tags = new[]
        ///             {
        ///                 new Volcengine.Cen.Inputs.BandwidthPackageTagArgs
        ///                 {
        ///                     Key = "k1",
        ///                     Value = "v1",
        ///                 },
        ///             },
        ///         }));
        ///     }
        ///     var fooBandwidthPackages = Volcengine.Cen.GetBandwidthPackages.Invoke(new()
        ///     {
        ///         Ids = fooBandwidthPackage.Select(__item =&gt; __item.Id).ToList(),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetBandwidthPackagesResult> InvokeAsync(GetBandwidthPackagesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBandwidthPackagesResult>("volcengine:cen/getBandwidthPackages:getBandwidthPackages", args ?? new GetBandwidthPackagesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cen bandwidth packages
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var fooBandwidthPackage = new List&lt;Volcengine.Cen.BandwidthPackage&gt;();
        ///     for (var rangeIndex = 0; rangeIndex &lt; 2; rangeIndex++)
        ///     {
        ///         var range = new { Value = rangeIndex };
        ///         fooBandwidthPackage.Add(new Volcengine.Cen.BandwidthPackage($"fooBandwidthPackage-{range.Value}", new()
        ///         {
        ///             LocalGeographicRegionSetId = "China",
        ///             PeerGeographicRegionSetId = "China",
        ///             Bandwidth = 2,
        ///             CenBandwidthPackageName = $"acc-test-cen-bp-{range.Value}",
        ///             Description = "acc-test",
        ///             BillingType = "PrePaid",
        ///             PeriodUnit = "Month",
        ///             Period = 1,
        ///             ProjectName = "default",
        ///             Tags = new[]
        ///             {
        ///                 new Volcengine.Cen.Inputs.BandwidthPackageTagArgs
        ///                 {
        ///                     Key = "k1",
        ///                     Value = "v1",
        ///                 },
        ///             },
        ///         }));
        ///     }
        ///     var fooBandwidthPackages = Volcengine.Cen.GetBandwidthPackages.Invoke(new()
        ///     {
        ///         Ids = fooBandwidthPackage.Select(__item =&gt; __item.Id).ToList(),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetBandwidthPackagesResult> Invoke(GetBandwidthPackagesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBandwidthPackagesResult>("volcengine:cen/getBandwidthPackages:getBandwidthPackages", args ?? new GetBandwidthPackagesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBandwidthPackagesArgs : global::Pulumi.InvokeArgs
    {
        [Input("cenBandwidthPackageNames")]
        private List<string>? _cenBandwidthPackageNames;

        /// <summary>
        /// A list of cen bandwidth package names.
        /// </summary>
        public List<string> CenBandwidthPackageNames
        {
            get => _cenBandwidthPackageNames ?? (_cenBandwidthPackageNames = new List<string>());
            set => _cenBandwidthPackageNames = value;
        }

        /// <summary>
        /// A cen id.
        /// </summary>
        [Input("cenId")]
        public string? CenId { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of cen bandwidth package IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A local geographic region set id.
        /// </summary>
        [Input("localGeographicRegionSetId")]
        public string? LocalGeographicRegionSetId { get; set; }

        /// <summary>
        /// A Name Regex of cen bandwidth package.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// A peer geographic region set id.
        /// </summary>
        [Input("peerGeographicRegionSetId")]
        public string? PeerGeographicRegionSetId { get; set; }

        [Input("tags")]
        private List<Inputs.GetBandwidthPackagesTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public List<Inputs.GetBandwidthPackagesTagArgs> Tags
        {
            get => _tags ?? (_tags = new List<Inputs.GetBandwidthPackagesTagArgs>());
            set => _tags = value;
        }

        public GetBandwidthPackagesArgs()
        {
        }
        public static new GetBandwidthPackagesArgs Empty => new GetBandwidthPackagesArgs();
    }

    public sealed class GetBandwidthPackagesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("cenBandwidthPackageNames")]
        private InputList<string>? _cenBandwidthPackageNames;

        /// <summary>
        /// A list of cen bandwidth package names.
        /// </summary>
        public InputList<string> CenBandwidthPackageNames
        {
            get => _cenBandwidthPackageNames ?? (_cenBandwidthPackageNames = new InputList<string>());
            set => _cenBandwidthPackageNames = value;
        }

        /// <summary>
        /// A cen id.
        /// </summary>
        [Input("cenId")]
        public Input<string>? CenId { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of cen bandwidth package IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A local geographic region set id.
        /// </summary>
        [Input("localGeographicRegionSetId")]
        public Input<string>? LocalGeographicRegionSetId { get; set; }

        /// <summary>
        /// A Name Regex of cen bandwidth package.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// A peer geographic region set id.
        /// </summary>
        [Input("peerGeographicRegionSetId")]
        public Input<string>? PeerGeographicRegionSetId { get; set; }

        [Input("tags")]
        private InputList<Inputs.GetBandwidthPackagesTagInputArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.GetBandwidthPackagesTagInputArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.GetBandwidthPackagesTagInputArgs>());
            set => _tags = value;
        }

        public GetBandwidthPackagesInvokeArgs()
        {
        }
        public static new GetBandwidthPackagesInvokeArgs Empty => new GetBandwidthPackagesInvokeArgs();
    }


    [OutputType]
    public sealed class GetBandwidthPackagesResult
    {
        /// <summary>
        /// The collection of cen bandwidth package query.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBandwidthPackagesBandwidthPackageResult> BandwidthPackages;
        public readonly ImmutableArray<string> CenBandwidthPackageNames;
        public readonly string? CenId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// The local geographic region set id of the cen bandwidth package.
        /// </summary>
        public readonly string? LocalGeographicRegionSetId;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The peer geographic region set id of the cen bandwidth package.
        /// </summary>
        public readonly string? PeerGeographicRegionSetId;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBandwidthPackagesTagResult> Tags;
        /// <summary>
        /// The total count of cen bandwidth package query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetBandwidthPackagesResult(
            ImmutableArray<Outputs.GetBandwidthPackagesBandwidthPackageResult> bandwidthPackages,

            ImmutableArray<string> cenBandwidthPackageNames,

            string? cenId,

            string id,

            ImmutableArray<string> ids,

            string? localGeographicRegionSetId,

            string? nameRegex,

            string? outputFile,

            string? peerGeographicRegionSetId,

            ImmutableArray<Outputs.GetBandwidthPackagesTagResult> tags,

            int totalCount)
        {
            BandwidthPackages = bandwidthPackages;
            CenBandwidthPackageNames = cenBandwidthPackageNames;
            CenId = cenId;
            Id = id;
            Ids = ids;
            LocalGeographicRegionSetId = localGeographicRegionSetId;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            PeerGeographicRegionSetId = peerGeographicRegionSetId;
            Tags = tags;
            TotalCount = totalCount;
        }
    }
}
