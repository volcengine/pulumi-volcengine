// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Ebs
{
    /// <summary>
    /// Provides a resource to manage ebs auto snapshot policy
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
    ///     var foo = new Volcengine.Ebs.AutoSnapshotPolicy("foo", new()
    ///     {
    ///         AutoSnapshotPolicyName = "acc-test-auto-snapshot-policy",
    ///         ProjectName = "default",
    ///         RepeatWeekdays = new[]
    ///         {
    ///             "2",
    ///             "6",
    ///         },
    ///         RetentionDays = -1,
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Ebs.Inputs.AutoSnapshotPolicyTagArgs
    ///             {
    ///                 Key = "k1",
    ///                 Value = "v1",
    ///             },
    ///         },
    ///         TimePoints = new[]
    ///         {
    ///             "1",
    ///             "5",
    ///             "9",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// EbsAutoSnapshotPolicy can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import volcengine:ebs/autoSnapshotPolicy:AutoSnapshotPolicy default resource_id
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:ebs/autoSnapshotPolicy:AutoSnapshotPolicy")]
    public partial class AutoSnapshotPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the auto snapshot policy.
        /// </summary>
        [Output("autoSnapshotPolicyName")]
        public Output<string> AutoSnapshotPolicyName { get; private set; } = null!;

        /// <summary>
        /// The creation time of the auto snapshot policy.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The project name of the auto snapshot policy.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot. The value range is `1-30`. Only one of `repeat_weekdays, repeat_days` can be specified.
        /// </summary>
        [Output("repeatDays")]
        public Output<int?> RepeatDays { get; private set; } = null!;

        /// <summary>
        /// The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday. Only one of `repeat_weekdays, repeat_days` can be specified.
        /// </summary>
        [Output("repeatWeekdays")]
        public Output<ImmutableArray<string>> RepeatWeekdays { get; private set; } = null!;

        /// <summary>
        /// The retention days of the auto snapshot. Valid values: -1 and 1~65536. `-1` means permanently preserving the snapshot.
        /// </summary>
        [Output("retentionDays")]
        public Output<int> RetentionDays { get; private set; } = null!;

        /// <summary>
        /// The status of the auto snapshot policy.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.AutoSnapshotPolicyTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.
        /// </summary>
        [Output("timePoints")]
        public Output<ImmutableArray<string>> TimePoints { get; private set; } = null!;

        /// <summary>
        /// The updated time of the auto snapshot policy.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The number of volumes associated with the auto snapshot policy.
        /// </summary>
        [Output("volumeNums")]
        public Output<int> VolumeNums { get; private set; } = null!;


        /// <summary>
        /// Create a AutoSnapshotPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AutoSnapshotPolicy(string name, AutoSnapshotPolicyArgs args, CustomResourceOptions? options = null)
            : base("volcengine:ebs/autoSnapshotPolicy:AutoSnapshotPolicy", name, args ?? new AutoSnapshotPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AutoSnapshotPolicy(string name, Input<string> id, AutoSnapshotPolicyState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:ebs/autoSnapshotPolicy:AutoSnapshotPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AutoSnapshotPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AutoSnapshotPolicy Get(string name, Input<string> id, AutoSnapshotPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new AutoSnapshotPolicy(name, id, state, options);
        }
    }

    public sealed class AutoSnapshotPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the auto snapshot policy.
        /// </summary>
        [Input("autoSnapshotPolicyName", required: true)]
        public Input<string> AutoSnapshotPolicyName { get; set; } = null!;

        /// <summary>
        /// The project name of the auto snapshot policy.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot. The value range is `1-30`. Only one of `repeat_weekdays, repeat_days` can be specified.
        /// </summary>
        [Input("repeatDays")]
        public Input<int>? RepeatDays { get; set; }

        [Input("repeatWeekdays")]
        private InputList<string>? _repeatWeekdays;

        /// <summary>
        /// The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday. Only one of `repeat_weekdays, repeat_days` can be specified.
        /// </summary>
        public InputList<string> RepeatWeekdays
        {
            get => _repeatWeekdays ?? (_repeatWeekdays = new InputList<string>());
            set => _repeatWeekdays = value;
        }

        /// <summary>
        /// The retention days of the auto snapshot. Valid values: -1 and 1~65536. `-1` means permanently preserving the snapshot.
        /// </summary>
        [Input("retentionDays", required: true)]
        public Input<int> RetentionDays { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.AutoSnapshotPolicyTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.AutoSnapshotPolicyTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.AutoSnapshotPolicyTagArgs>());
            set => _tags = value;
        }

        [Input("timePoints", required: true)]
        private InputList<string>? _timePoints;

        /// <summary>
        /// The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.
        /// </summary>
        public InputList<string> TimePoints
        {
            get => _timePoints ?? (_timePoints = new InputList<string>());
            set => _timePoints = value;
        }

        public AutoSnapshotPolicyArgs()
        {
        }
        public static new AutoSnapshotPolicyArgs Empty => new AutoSnapshotPolicyArgs();
    }

    public sealed class AutoSnapshotPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the auto snapshot policy.
        /// </summary>
        [Input("autoSnapshotPolicyName")]
        public Input<string>? AutoSnapshotPolicyName { get; set; }

        /// <summary>
        /// The creation time of the auto snapshot policy.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The project name of the auto snapshot policy.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot. The value range is `1-30`. Only one of `repeat_weekdays, repeat_days` can be specified.
        /// </summary>
        [Input("repeatDays")]
        public Input<int>? RepeatDays { get; set; }

        [Input("repeatWeekdays")]
        private InputList<string>? _repeatWeekdays;

        /// <summary>
        /// The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday. Only one of `repeat_weekdays, repeat_days` can be specified.
        /// </summary>
        public InputList<string> RepeatWeekdays
        {
            get => _repeatWeekdays ?? (_repeatWeekdays = new InputList<string>());
            set => _repeatWeekdays = value;
        }

        /// <summary>
        /// The retention days of the auto snapshot. Valid values: -1 and 1~65536. `-1` means permanently preserving the snapshot.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        /// <summary>
        /// The status of the auto snapshot policy.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<Inputs.AutoSnapshotPolicyTagGetArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.AutoSnapshotPolicyTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.AutoSnapshotPolicyTagGetArgs>());
            set => _tags = value;
        }

        [Input("timePoints")]
        private InputList<string>? _timePoints;

        /// <summary>
        /// The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.
        /// </summary>
        public InputList<string> TimePoints
        {
            get => _timePoints ?? (_timePoints = new InputList<string>());
            set => _timePoints = value;
        }

        /// <summary>
        /// The updated time of the auto snapshot policy.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The number of volumes associated with the auto snapshot policy.
        /// </summary>
        [Input("volumeNums")]
        public Input<int>? VolumeNums { get; set; }

        public AutoSnapshotPolicyState()
        {
        }
        public static new AutoSnapshotPolicyState Empty => new AutoSnapshotPolicyState();
    }
}