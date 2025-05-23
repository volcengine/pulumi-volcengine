// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpc
{
    [Obsolete(@"volcengine.vpc.SecurityGroups has been deprecated in favor of volcengine.vpc.getSecurityGroups")]
    public static class SecurityGroups
    {
        /// <summary>
        /// Use this data source to query detailed information of security groups
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
        ///     var @default = Volcengine.Vpc.GetSecurityGroups.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "sg-273ycgql3ig3k7fap8t3dyvqx",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<SecurityGroupsResult> InvokeAsync(SecurityGroupsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<SecurityGroupsResult>("volcengine:vpc/securityGroups:SecurityGroups", args ?? new SecurityGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of security groups
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
        ///     var @default = Volcengine.Vpc.GetSecurityGroups.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "sg-273ycgql3ig3k7fap8t3dyvqx",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<SecurityGroupsResult> Invoke(SecurityGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<SecurityGroupsResult>("volcengine:vpc/securityGroups:SecurityGroups", args ?? new SecurityGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class SecurityGroupsArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of SecurityGroup IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of SecurityGroup.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The ProjectName of SecurityGroup.
        /// </summary>
        [Input("projectName")]
        public string? ProjectName { get; set; }

        [Input("securityGroupNames")]
        private List<string>? _securityGroupNames;

        /// <summary>
        /// The list of security group name to query.
        /// </summary>
        public List<string> SecurityGroupNames
        {
            get => _securityGroupNames ?? (_securityGroupNames = new List<string>());
            set => _securityGroupNames = value;
        }

        [Input("tags")]
        private List<Inputs.SecurityGroupsTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public List<Inputs.SecurityGroupsTagArgs> Tags
        {
            get => _tags ?? (_tags = new List<Inputs.SecurityGroupsTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of vpc where security group is located.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public SecurityGroupsArgs()
        {
        }
        public static new SecurityGroupsArgs Empty => new SecurityGroupsArgs();
    }

    public sealed class SecurityGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of SecurityGroup IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of SecurityGroup.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The ProjectName of SecurityGroup.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        [Input("securityGroupNames")]
        private InputList<string>? _securityGroupNames;

        /// <summary>
        /// The list of security group name to query.
        /// </summary>
        public InputList<string> SecurityGroupNames
        {
            get => _securityGroupNames ?? (_securityGroupNames = new InputList<string>());
            set => _securityGroupNames = value;
        }

        [Input("tags")]
        private InputList<Inputs.SecurityGroupsTagInputArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.SecurityGroupsTagInputArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.SecurityGroupsTagInputArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of vpc where security group is located.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public SecurityGroupsInvokeArgs()
        {
        }
        public static new SecurityGroupsInvokeArgs Empty => new SecurityGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class SecurityGroupsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The ProjectName of SecurityGroup.
        /// </summary>
        public readonly string? ProjectName;
        public readonly ImmutableArray<string> SecurityGroupNames;
        /// <summary>
        /// The collection of SecurityGroup query.
        /// </summary>
        public readonly ImmutableArray<Outputs.SecurityGroupsSecurityGroupResult> SecurityGroups;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.SecurityGroupsTagResult> Tags;
        /// <summary>
        /// The total count of SecurityGroup query.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The ID of Vpc.
        /// </summary>
        public readonly string? VpcId;

        [OutputConstructor]
        private SecurityGroupsResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            string? outputFile,

            string? projectName,

            ImmutableArray<string> securityGroupNames,

            ImmutableArray<Outputs.SecurityGroupsSecurityGroupResult> securityGroups,

            ImmutableArray<Outputs.SecurityGroupsTagResult> tags,

            int totalCount,

            string? vpcId)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            ProjectName = projectName;
            SecurityGroupNames = securityGroupNames;
            SecurityGroups = securityGroups;
            Tags = tags;
            TotalCount = totalCount;
            VpcId = vpcId;
        }
    }
}
