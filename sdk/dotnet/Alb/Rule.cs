// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Alb
{
    /// <summary>
    /// Provides a resource to manage alb rule
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
    ///     var foo = new Volcengine.Alb.Rule("foo", new()
    ///     {
    ///         Description = "test",
    ///         Domain = "www.test.com",
    ///         ListenerId = "lsn-1iidd19u4oni874adhezjkyj3",
    ///         RedirectConfig = new Volcengine.Alb.Inputs.RuleRedirectConfigArgs
    ///         {
    ///             RedirectDomain = "www.testtest.com",
    ///             RedirectHttpCode = "302",
    ///             RedirectPort = "555",
    ///             RedirectUri = "/testtest",
    ///         },
    ///         RewriteConfig = new Volcengine.Alb.Inputs.RuleRewriteConfigArgs
    ///         {
    ///             RewritePath = "/test",
    ///         },
    ///         RewriteEnabled = "off",
    ///         RuleAction = "Redirect",
    ///         ServerGroupId = "rsp-1g72w74y4umf42zbhq4k4hnln",
    ///         TrafficLimitEnabled = "off",
    ///         TrafficLimitQps = 100,
    ///         Url = "/test",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// AlbRule can be imported using the listener id and rule id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:alb/rule:Rule default lsn-273yv0mhs5xj47fap8sehiiso:rule-****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:alb/rule:Rule")]
    public partial class Rule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the Rule.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The domain of Rule.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// The ID of listener.
        /// </summary>
        [Output("listenerId")]
        public Output<string> ListenerId { get; private set; } = null!;

        /// <summary>
        /// The redirect related configuration.
        /// </summary>
        [Output("redirectConfig")]
        public Output<Outputs.RuleRedirectConfig?> RedirectConfig { get; private set; } = null!;

        /// <summary>
        /// The list of rewrite configurations.
        /// </summary>
        [Output("rewriteConfig")]
        public Output<Outputs.RuleRewriteConfig?> RewriteConfig { get; private set; } = null!;

        /// <summary>
        /// Rewrite configuration switch for forwarding rules, only allows configuration and takes effect when RuleAction is empty (i.e., forwarding to server group). Only available for whitelist users, please submit an application to experience. Supported values are as follows:
        /// on: enable.
        /// off: disable.
        /// </summary>
        [Output("rewriteEnabled")]
        public Output<string?> RewriteEnabled { get; private set; } = null!;

        /// <summary>
        /// The forwarding rule action, if this parameter is empty(`""`), forward to server group, if value is `Redirect`, will redirect.
        /// </summary>
        [Output("ruleAction")]
        public Output<string> RuleAction { get; private set; } = null!;

        /// <summary>
        /// The ID of rule.
        /// </summary>
        [Output("ruleId")]
        public Output<string> RuleId { get; private set; } = null!;

        /// <summary>
        /// Server group ID, this parameter is required if `rule_action` is empty.
        /// </summary>
        [Output("serverGroupId")]
        public Output<string?> ServerGroupId { get; private set; } = null!;

        /// <summary>
        /// Forwarding rule QPS rate limiting switch:
        /// on: enable.
        /// off: disable (default).
        /// </summary>
        [Output("trafficLimitEnabled")]
        public Output<string?> TrafficLimitEnabled { get; private set; } = null!;

        /// <summary>
        /// When Rules.N.TrafficLimitEnabled is turned on, this field is required. Requests per second. Valid values are between 100 and 100000.
        /// </summary>
        [Output("trafficLimitQps")]
        public Output<int?> TrafficLimitQps { get; private set; } = null!;

        /// <summary>
        /// The Url of Rule.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Rule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Rule(string name, RuleArgs args, CustomResourceOptions? options = null)
            : base("volcengine:alb/rule:Rule", name, args ?? new RuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Rule(string name, Input<string> id, RuleState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:alb/rule:Rule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Rule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Rule Get(string name, Input<string> id, RuleState? state = null, CustomResourceOptions? options = null)
        {
            return new Rule(name, id, state, options);
        }
    }

    public sealed class RuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The domain of Rule.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The ID of listener.
        /// </summary>
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        /// <summary>
        /// The redirect related configuration.
        /// </summary>
        [Input("redirectConfig")]
        public Input<Inputs.RuleRedirectConfigArgs>? RedirectConfig { get; set; }

        /// <summary>
        /// The list of rewrite configurations.
        /// </summary>
        [Input("rewriteConfig")]
        public Input<Inputs.RuleRewriteConfigArgs>? RewriteConfig { get; set; }

        /// <summary>
        /// Rewrite configuration switch for forwarding rules, only allows configuration and takes effect when RuleAction is empty (i.e., forwarding to server group). Only available for whitelist users, please submit an application to experience. Supported values are as follows:
        /// on: enable.
        /// off: disable.
        /// </summary>
        [Input("rewriteEnabled")]
        public Input<string>? RewriteEnabled { get; set; }

        /// <summary>
        /// The forwarding rule action, if this parameter is empty(`""`), forward to server group, if value is `Redirect`, will redirect.
        /// </summary>
        [Input("ruleAction", required: true)]
        public Input<string> RuleAction { get; set; } = null!;

        /// <summary>
        /// Server group ID, this parameter is required if `rule_action` is empty.
        /// </summary>
        [Input("serverGroupId")]
        public Input<string>? ServerGroupId { get; set; }

        /// <summary>
        /// Forwarding rule QPS rate limiting switch:
        /// on: enable.
        /// off: disable (default).
        /// </summary>
        [Input("trafficLimitEnabled")]
        public Input<string>? TrafficLimitEnabled { get; set; }

        /// <summary>
        /// When Rules.N.TrafficLimitEnabled is turned on, this field is required. Requests per second. Valid values are between 100 and 100000.
        /// </summary>
        [Input("trafficLimitQps")]
        public Input<int>? TrafficLimitQps { get; set; }

        /// <summary>
        /// The Url of Rule.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public RuleArgs()
        {
        }
        public static new RuleArgs Empty => new RuleArgs();
    }

    public sealed class RuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The domain of Rule.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The ID of listener.
        /// </summary>
        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        /// <summary>
        /// The redirect related configuration.
        /// </summary>
        [Input("redirectConfig")]
        public Input<Inputs.RuleRedirectConfigGetArgs>? RedirectConfig { get; set; }

        /// <summary>
        /// The list of rewrite configurations.
        /// </summary>
        [Input("rewriteConfig")]
        public Input<Inputs.RuleRewriteConfigGetArgs>? RewriteConfig { get; set; }

        /// <summary>
        /// Rewrite configuration switch for forwarding rules, only allows configuration and takes effect when RuleAction is empty (i.e., forwarding to server group). Only available for whitelist users, please submit an application to experience. Supported values are as follows:
        /// on: enable.
        /// off: disable.
        /// </summary>
        [Input("rewriteEnabled")]
        public Input<string>? RewriteEnabled { get; set; }

        /// <summary>
        /// The forwarding rule action, if this parameter is empty(`""`), forward to server group, if value is `Redirect`, will redirect.
        /// </summary>
        [Input("ruleAction")]
        public Input<string>? RuleAction { get; set; }

        /// <summary>
        /// The ID of rule.
        /// </summary>
        [Input("ruleId")]
        public Input<string>? RuleId { get; set; }

        /// <summary>
        /// Server group ID, this parameter is required if `rule_action` is empty.
        /// </summary>
        [Input("serverGroupId")]
        public Input<string>? ServerGroupId { get; set; }

        /// <summary>
        /// Forwarding rule QPS rate limiting switch:
        /// on: enable.
        /// off: disable (default).
        /// </summary>
        [Input("trafficLimitEnabled")]
        public Input<string>? TrafficLimitEnabled { get; set; }

        /// <summary>
        /// When Rules.N.TrafficLimitEnabled is turned on, this field is required. Requests per second. Valid values are between 100 and 100000.
        /// </summary>
        [Input("trafficLimitQps")]
        public Input<int>? TrafficLimitQps { get; set; }

        /// <summary>
        /// The Url of Rule.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public RuleState()
        {
        }
        public static new RuleState Empty => new RuleState();
    }
}