// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Iam
{
    public static class GetAccessKeys
    {
        /// <summary>
        /// Use this data source to query detailed information of iam access keys
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
        ///     var foo = Volcengine.Iam.GetAccessKeys.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAccessKeysResult> InvokeAsync(GetAccessKeysArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccessKeysResult>("volcengine:iam/getAccessKeys:getAccessKeys", args ?? new GetAccessKeysArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of iam access keys
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
        ///     var foo = Volcengine.Iam.GetAccessKeys.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAccessKeysResult> Invoke(GetAccessKeysInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccessKeysResult>("volcengine:iam/getAccessKeys:getAccessKeys", args ?? new GetAccessKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccessKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A Name Regex of IAM.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The user names.
        /// </summary>
        [Input("userName")]
        public string? UserName { get; set; }

        public GetAccessKeysArgs()
        {
        }
        public static new GetAccessKeysArgs Empty => new GetAccessKeysArgs();
    }

    public sealed class GetAccessKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A Name Regex of IAM.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The user names.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public GetAccessKeysInvokeArgs()
        {
        }
        public static new GetAccessKeysInvokeArgs Empty => new GetAccessKeysInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccessKeysResult
    {
        /// <summary>
        /// The collection of access keys.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccessKeysAccessKeyMetadataResult> AccessKeyMetadatas;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of user query.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The user name.
        /// </summary>
        public readonly string? UserName;

        [OutputConstructor]
        private GetAccessKeysResult(
            ImmutableArray<Outputs.GetAccessKeysAccessKeyMetadataResult> accessKeyMetadatas,

            string id,

            string? nameRegex,

            string? outputFile,

            int totalCount,

            string? userName)
        {
            AccessKeyMetadatas = accessKeyMetadatas;
            Id = id;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            TotalCount = totalCount;
            UserName = userName;
        }
    }
}
