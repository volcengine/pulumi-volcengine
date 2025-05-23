// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds
{
    /// <summary>
    /// (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Provides a resource to manage rds parameter template
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
    ///     var foo = new Volcengine.Rds.ParameterTemplate("foo", new()
    ///     {
    ///         TemplateDesc = "created by terraform",
    ///         TemplateName = "tf-template",
    ///         TemplateParams = new[]
    ///         {
    ///             new Volcengine.Rds.Inputs.ParameterTemplateTemplateParamArgs
    ///             {
    ///                 Name = "auto_increment_increment",
    ///                 RunningValue = "2",
    ///             },
    ///             new Volcengine.Rds.Inputs.ParameterTemplateTemplateParamArgs
    ///             {
    ///                 Name = "slow_query_log",
    ///                 RunningValue = "ON",
    ///             },
    ///             new Volcengine.Rds.Inputs.ParameterTemplateTemplateParamArgs
    ///             {
    ///                 Name = "net_retry_count",
    ///                 RunningValue = "33",
    ///             },
    ///         },
    ///         TemplateType = "MySQL",
    ///         TemplateTypeVersion = "MySQL_Community_5_7",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RDS Instance can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:rds/parameterTemplate:ParameterTemplate default mysql-sys-80bb93aa14be22d0
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:rds/parameterTemplate:ParameterTemplate")]
    public partial class ParameterTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Parameter template description.
        /// </summary>
        [Output("templateDesc")]
        public Output<string?> TemplateDesc { get; private set; } = null!;

        /// <summary>
        /// Parameter template name.
        /// </summary>
        [Output("templateName")]
        public Output<string> TemplateName { get; private set; } = null!;

        /// <summary>
        /// Template parameters. InstanceParam only needs to pass Name and RunningValue.
        /// </summary>
        [Output("templateParams")]
        public Output<ImmutableArray<Outputs.ParameterTemplateTemplateParam>> TemplateParams { get; private set; } = null!;

        /// <summary>
        /// Parameter template database type, range of values:
        /// MySQL - MySQL database. (Defaults).
        /// </summary>
        [Output("templateType")]
        public Output<string?> TemplateType { get; private set; } = null!;

        /// <summary>
        /// Parameter template database version, value range:
        /// MySQL_Community_5_7 - MySQL 5.7 (default)
        /// MySQL_8_0 - MySQL 8.0.
        /// </summary>
        [Output("templateTypeVersion")]
        public Output<string?> TemplateTypeVersion { get; private set; } = null!;


        /// <summary>
        /// Create a ParameterTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ParameterTemplate(string name, ParameterTemplateArgs args, CustomResourceOptions? options = null)
            : base("volcengine:rds/parameterTemplate:ParameterTemplate", name, args ?? new ParameterTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ParameterTemplate(string name, Input<string> id, ParameterTemplateState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:rds/parameterTemplate:ParameterTemplate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/volcengine",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ParameterTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ParameterTemplate Get(string name, Input<string> id, ParameterTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new ParameterTemplate(name, id, state, options);
        }
    }

    public sealed class ParameterTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parameter template description.
        /// </summary>
        [Input("templateDesc")]
        public Input<string>? TemplateDesc { get; set; }

        /// <summary>
        /// Parameter template name.
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        [Input("templateParams", required: true)]
        private InputList<Inputs.ParameterTemplateTemplateParamArgs>? _templateParams;

        /// <summary>
        /// Template parameters. InstanceParam only needs to pass Name and RunningValue.
        /// </summary>
        public InputList<Inputs.ParameterTemplateTemplateParamArgs> TemplateParams
        {
            get => _templateParams ?? (_templateParams = new InputList<Inputs.ParameterTemplateTemplateParamArgs>());
            set => _templateParams = value;
        }

        /// <summary>
        /// Parameter template database type, range of values:
        /// MySQL - MySQL database. (Defaults).
        /// </summary>
        [Input("templateType")]
        public Input<string>? TemplateType { get; set; }

        /// <summary>
        /// Parameter template database version, value range:
        /// MySQL_Community_5_7 - MySQL 5.7 (default)
        /// MySQL_8_0 - MySQL 8.0.
        /// </summary>
        [Input("templateTypeVersion")]
        public Input<string>? TemplateTypeVersion { get; set; }

        public ParameterTemplateArgs()
        {
        }
        public static new ParameterTemplateArgs Empty => new ParameterTemplateArgs();
    }

    public sealed class ParameterTemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parameter template description.
        /// </summary>
        [Input("templateDesc")]
        public Input<string>? TemplateDesc { get; set; }

        /// <summary>
        /// Parameter template name.
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        [Input("templateParams")]
        private InputList<Inputs.ParameterTemplateTemplateParamGetArgs>? _templateParams;

        /// <summary>
        /// Template parameters. InstanceParam only needs to pass Name and RunningValue.
        /// </summary>
        public InputList<Inputs.ParameterTemplateTemplateParamGetArgs> TemplateParams
        {
            get => _templateParams ?? (_templateParams = new InputList<Inputs.ParameterTemplateTemplateParamGetArgs>());
            set => _templateParams = value;
        }

        /// <summary>
        /// Parameter template database type, range of values:
        /// MySQL - MySQL database. (Defaults).
        /// </summary>
        [Input("templateType")]
        public Input<string>? TemplateType { get; set; }

        /// <summary>
        /// Parameter template database version, value range:
        /// MySQL_Community_5_7 - MySQL 5.7 (default)
        /// MySQL_8_0 - MySQL 8.0.
        /// </summary>
        [Input("templateTypeVersion")]
        public Input<string>? TemplateTypeVersion { get; set; }

        public ParameterTemplateState()
        {
        }
        public static new ParameterTemplateState Empty => new ParameterTemplateState();
    }
}
