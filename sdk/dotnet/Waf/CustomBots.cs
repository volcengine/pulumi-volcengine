// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Waf
{
    [Obsolete(@"volcengine.waf.CustomBots has been deprecated in favor of volcengine.waf.getCustomBots")]
    public static class CustomBots
    {
        /// <summary>
        /// Use this data source to query detailed information of waf custom bots
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
        ///     var foo = Volcengine.Waf.GetCustomBots.Invoke(new()
        ///     {
        ///         Host = "www.tf-test.com",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<CustomBotsResult> InvokeAsync(CustomBotsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<CustomBotsResult>("volcengine:waf/customBots:CustomBots", args ?? new CustomBotsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of waf custom bots
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
        ///     var foo = Volcengine.Waf.GetCustomBots.Invoke(new()
        ///     {
        ///         Host = "www.tf-test.com",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<CustomBotsResult> Invoke(CustomBotsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<CustomBotsResult>("volcengine:waf/customBots:CustomBots", args ?? new CustomBotsInvokeArgs(), options.WithDefaults());
    }


    public sealed class CustomBotsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain names that need to be viewed.
        /// </summary>
        [Input("host")]
        public string? Host { get; set; }

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

        public CustomBotsArgs()
        {
        }
        public static new CustomBotsArgs Empty => new CustomBotsArgs();
    }

    public sealed class CustomBotsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain names that need to be viewed.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

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

        public CustomBotsInvokeArgs()
        {
        }
        public static new CustomBotsInvokeArgs Empty => new CustomBotsInvokeArgs();
    }


    [OutputType]
    public sealed class CustomBotsResult
    {
        /// <summary>
        /// The Details of Custom bot.
        /// </summary>
        public readonly ImmutableArray<Outputs.CustomBotsDataResult> Datas;
        public readonly string? Host;
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
        private CustomBotsResult(
            ImmutableArray<Outputs.CustomBotsDataResult> datas,

            string? host,

            string id,

            string? nameRegex,

            string? outputFile,

            int totalCount)
        {
            Datas = datas;
            Host = host;
            Id = id;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
