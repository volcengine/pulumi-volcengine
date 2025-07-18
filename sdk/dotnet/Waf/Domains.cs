// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Waf
{
    [Obsolete(@"volcengine.waf.Domains has been deprecated in favor of volcengine.waf.getDomains")]
    public static class Domains
    {
        /// <summary>
        /// Use this data source to query detailed information of waf domains
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
        ///     var foo = Volcengine.Waf.GetDomains.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<DomainsResult> InvokeAsync(DomainsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<DomainsResult>("volcengine:waf/domains:Domains", args ?? new DomainsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of waf domains
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
        ///     var foo = Volcengine.Waf.GetDomains.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<DomainsResult> Invoke(DomainsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<DomainsResult>("volcengine:waf/domains:Domains", args ?? new DomainsInvokeArgs(), options.WithDefaults());
    }


    public sealed class DomainsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Matching mode.
        /// </summary>
        [Input("accurateQuery")]
        public int? AccurateQuery { get; set; }

        /// <summary>
        /// The domain name of the protected website that needs to be queried.
        /// </summary>
        [Input("domain")]
        public string? Domain { get; set; }

        /// <summary>
        /// A Name Regex of Resource.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public DomainsArgs()
        {
        }
        public static new DomainsArgs Empty => new DomainsArgs();
    }

    public sealed class DomainsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Matching mode.
        /// </summary>
        [Input("accurateQuery")]
        public Input<int>? AccurateQuery { get; set; }

        /// <summary>
        /// The domain name of the protected website that needs to be queried.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// A Name Regex of Resource.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public DomainsInvokeArgs()
        {
        }
        public static new DomainsInvokeArgs Empty => new DomainsInvokeArgs();
    }


    [OutputType]
    public sealed class DomainsResult
    {
        public readonly int? AccurateQuery;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.DomainsDataResult> Datas;
        /// <summary>
        /// domain names that need to be protected by WAF.
        /// </summary>
        public readonly string? Domain;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private DomainsResult(
            int? accurateQuery,

            ImmutableArray<Outputs.DomainsDataResult> datas,

            string? domain,

            string id,

            string? nameRegex,

            string? outputFile,

            int totalCount)
        {
            AccurateQuery = accurateQuery;
            Datas = datas;
            Domain = domain;
            Id = id;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
