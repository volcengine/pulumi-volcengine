// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Ecs
{
    public static class LaunchTemplates
    {
        /// <summary>
        /// Use this data source to query detailed information of ecs launch templates
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var @default = Output.Create(Volcengine.Ecs.LaunchTemplates.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<LaunchTemplatesResult> InvokeAsync(LaunchTemplatesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<LaunchTemplatesResult>("volcengine:ecs/launchTemplates:LaunchTemplates", args ?? new LaunchTemplatesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of ecs launch templates
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var @default = Output.Create(Volcengine.Ecs.LaunchTemplates.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<LaunchTemplatesResult> Invoke(LaunchTemplatesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<LaunchTemplatesResult>("volcengine:ecs/launchTemplates:LaunchTemplates", args ?? new LaunchTemplatesInvokeArgs(), options.WithDefaults());
    }


    public sealed class LaunchTemplatesArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of launch template ids.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("launchTemplateNames")]
        private List<string>? _launchTemplateNames;

        /// <summary>
        /// A list of launch template names.
        /// </summary>
        public List<string> LaunchTemplateNames
        {
            get => _launchTemplateNames ?? (_launchTemplateNames = new List<string>());
            set => _launchTemplateNames = value;
        }

        /// <summary>
        /// A Name Regex of scaling policy.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public LaunchTemplatesArgs()
        {
        }
    }

    public sealed class LaunchTemplatesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of launch template ids.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("launchTemplateNames")]
        private InputList<string>? _launchTemplateNames;

        /// <summary>
        /// A list of launch template names.
        /// </summary>
        public InputList<string> LaunchTemplateNames
        {
            get => _launchTemplateNames ?? (_launchTemplateNames = new InputList<string>());
            set => _launchTemplateNames = value;
        }

        /// <summary>
        /// A Name Regex of scaling policy.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public LaunchTemplatesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class LaunchTemplatesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<string> LaunchTemplateNames;
        /// <summary>
        /// The collection of launch templates.
        /// </summary>
        public readonly ImmutableArray<Outputs.LaunchTemplatesLaunchTemplateResult> LaunchTemplates;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of scaling policy query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private LaunchTemplatesResult(
            string id,

            ImmutableArray<string> ids,

            ImmutableArray<string> launchTemplateNames,

            ImmutableArray<Outputs.LaunchTemplatesLaunchTemplateResult> launchTemplates,

            string? nameRegex,

            string? outputFile,

            int totalCount)
        {
            Id = id;
            Ids = ids;
            LaunchTemplateNames = launchTemplateNames;
            LaunchTemplates = launchTemplates;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}