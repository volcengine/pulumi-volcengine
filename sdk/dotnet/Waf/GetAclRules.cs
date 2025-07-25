// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Waf
{
    public static class GetAclRules
    {
        /// <summary>
        /// Use this data source to query detailed information of waf acl rules
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
        ///     var foo = Volcengine.Waf.GetAclRules.Invoke(new()
        ///     {
        ///         AclType = "Block",
        ///         Actions = new[]
        ///         {
        ///             "observe",
        ///         },
        ///         DefenceHosts = new[]
        ///         {
        ///             "www.tf-test.com",
        ///         },
        ///         Enables = new[]
        ///         {
        ///             1,
        ///         },
        ///         ProjectName = "default",
        ///         RuleName = "tf-test",
        ///         TimeOrderBy = "ASC",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAclRulesResult> InvokeAsync(GetAclRulesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAclRulesResult>("volcengine:waf/getAclRules:getAclRules", args ?? new GetAclRulesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of waf acl rules
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
        ///     var foo = Volcengine.Waf.GetAclRules.Invoke(new()
        ///     {
        ///         AclType = "Block",
        ///         Actions = new[]
        ///         {
        ///             "observe",
        ///         },
        ///         DefenceHosts = new[]
        ///         {
        ///             "www.tf-test.com",
        ///         },
        ///         Enables = new[]
        ///         {
        ///             1,
        ///         },
        ///         ProjectName = "default",
        ///         RuleName = "tf-test",
        ///         TimeOrderBy = "ASC",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAclRulesResult> Invoke(GetAclRulesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAclRulesResult>("volcengine:waf/getAclRules:getAclRules", args ?? new GetAclRulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAclRulesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The types of access control rules.
        /// </summary>
        [Input("aclType", required: true)]
        public string AclType { get; set; } = null!;

        [Input("actions")]
        private List<string>? _actions;

        /// <summary>
        /// Action to be taken on requests that match the rule.
        /// </summary>
        public List<string> Actions
        {
            get => _actions ?? (_actions = new List<string>());
            set => _actions = value;
        }

        [Input("defenceHosts")]
        private List<string>? _defenceHosts;

        /// <summary>
        /// The list of queried domain names.
        /// </summary>
        public List<string> DefenceHosts
        {
            get => _defenceHosts ?? (_defenceHosts = new List<string>());
            set => _defenceHosts = value;
        }

        [Input("enables")]
        private List<int>? _enables;

        /// <summary>
        /// The enabled status of the rule.
        /// </summary>
        public List<int> Enables
        {
            get => _enables ?? (_enables = new List<int>());
            set => _enables = value;
        }

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

        /// <summary>
        /// The name of the project to which your domain names belong.
        /// </summary>
        [Input("projectName")]
        public string? ProjectName { get; set; }

        /// <summary>
        /// Rule name, fuzzy search.
        /// </summary>
        [Input("ruleName")]
        public string? RuleName { get; set; }

        /// <summary>
        /// Rule unique identifier, precise search.
        /// </summary>
        [Input("ruleTag")]
        public string? RuleTag { get; set; }

        /// <summary>
        /// The list shows the timing sequence.
        /// </summary>
        [Input("timeOrderBy")]
        public string? TimeOrderBy { get; set; }

        public GetAclRulesArgs()
        {
        }
        public static new GetAclRulesArgs Empty => new GetAclRulesArgs();
    }

    public sealed class GetAclRulesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The types of access control rules.
        /// </summary>
        [Input("aclType", required: true)]
        public Input<string> AclType { get; set; } = null!;

        [Input("actions")]
        private InputList<string>? _actions;

        /// <summary>
        /// Action to be taken on requests that match the rule.
        /// </summary>
        public InputList<string> Actions
        {
            get => _actions ?? (_actions = new InputList<string>());
            set => _actions = value;
        }

        [Input("defenceHosts")]
        private InputList<string>? _defenceHosts;

        /// <summary>
        /// The list of queried domain names.
        /// </summary>
        public InputList<string> DefenceHosts
        {
            get => _defenceHosts ?? (_defenceHosts = new InputList<string>());
            set => _defenceHosts = value;
        }

        [Input("enables")]
        private InputList<int>? _enables;

        /// <summary>
        /// The enabled status of the rule.
        /// </summary>
        public InputList<int> Enables
        {
            get => _enables ?? (_enables = new InputList<int>());
            set => _enables = value;
        }

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

        /// <summary>
        /// The name of the project to which your domain names belong.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// Rule name, fuzzy search.
        /// </summary>
        [Input("ruleName")]
        public Input<string>? RuleName { get; set; }

        /// <summary>
        /// Rule unique identifier, precise search.
        /// </summary>
        [Input("ruleTag")]
        public Input<string>? RuleTag { get; set; }

        /// <summary>
        /// The list shows the timing sequence.
        /// </summary>
        [Input("timeOrderBy")]
        public Input<string>? TimeOrderBy { get; set; }

        public GetAclRulesInvokeArgs()
        {
        }
        public static new GetAclRulesInvokeArgs Empty => new GetAclRulesInvokeArgs();
    }


    [OutputType]
    public sealed class GetAclRulesResult
    {
        public readonly string AclType;
        /// <summary>
        /// Action to be taken on requests that match the rule.
        /// </summary>
        public readonly ImmutableArray<string> Actions;
        public readonly ImmutableArray<string> DefenceHosts;
        /// <summary>
        /// Whether to enable the rule.
        /// </summary>
        public readonly ImmutableArray<int> Enables;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        public readonly string? ProjectName;
        public readonly string? RuleName;
        /// <summary>
        /// Rule unique identifier.
        /// </summary>
        public readonly string? RuleTag;
        /// <summary>
        /// Details of the rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAclRulesRuleResult> Rules;
        public readonly string? TimeOrderBy;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetAclRulesResult(
            string aclType,

            ImmutableArray<string> actions,

            ImmutableArray<string> defenceHosts,

            ImmutableArray<int> enables,

            string id,

            string? nameRegex,

            string? outputFile,

            string? projectName,

            string? ruleName,

            string? ruleTag,

            ImmutableArray<Outputs.GetAclRulesRuleResult> rules,

            string? timeOrderBy,

            int totalCount)
        {
            AclType = aclType;
            Actions = actions;
            DefenceHosts = defenceHosts;
            Enables = enables;
            Id = id;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            ProjectName = projectName;
            RuleName = ruleName;
            RuleTag = ruleTag;
            Rules = rules;
            TimeOrderBy = timeOrderBy;
            TotalCount = totalCount;
        }
    }
}
