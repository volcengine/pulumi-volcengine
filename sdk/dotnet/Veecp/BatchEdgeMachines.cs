// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Veecp
{
    [Obsolete(@"volcengine.veecp.BatchEdgeMachines has been deprecated in favor of volcengine.veecp.getBatchEdgeMachines")]
    public static class BatchEdgeMachines
    {
        /// <summary>
        /// Use this data source to query detailed information of veecp batch edge machines
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
        ///     var fooBatchEdgeMachine = new Volcengine.Veecp.BatchEdgeMachine("fooBatchEdgeMachine", new()
        ///     {
        ///         ClusterId = "ccvd7mte6t101fno98u60",
        ///         NodePoolId = "pcvd90uacnsr73g6bjic0",
        ///         TtlHours = 1,
        ///     });
        /// 
        ///     var fooBatchEdgeMachines = Volcengine.Veecp.GetBatchEdgeMachines.Invoke(new()
        ///     {
        ///         ClusterIds = new[]
        ///         {
        ///             fooBatchEdgeMachine.ClusterId,
        ///         },
        ///         Ids = new[]
        ///         {
        ///             fooBatchEdgeMachine.Id,
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<BatchEdgeMachinesResult> InvokeAsync(BatchEdgeMachinesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<BatchEdgeMachinesResult>("volcengine:veecp/batchEdgeMachines:BatchEdgeMachines", args ?? new BatchEdgeMachinesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of veecp batch edge machines
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
        ///     var fooBatchEdgeMachine = new Volcengine.Veecp.BatchEdgeMachine("fooBatchEdgeMachine", new()
        ///     {
        ///         ClusterId = "ccvd7mte6t101fno98u60",
        ///         NodePoolId = "pcvd90uacnsr73g6bjic0",
        ///         TtlHours = 1,
        ///     });
        /// 
        ///     var fooBatchEdgeMachines = Volcengine.Veecp.GetBatchEdgeMachines.Invoke(new()
        ///     {
        ///         ClusterIds = new[]
        ///         {
        ///             fooBatchEdgeMachine.ClusterId,
        ///         },
        ///         Ids = new[]
        ///         {
        ///             fooBatchEdgeMachine.Id,
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<BatchEdgeMachinesResult> Invoke(BatchEdgeMachinesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<BatchEdgeMachinesResult>("volcengine:veecp/batchEdgeMachines:BatchEdgeMachines", args ?? new BatchEdgeMachinesInvokeArgs(), options.WithDefaults());
    }


    public sealed class BatchEdgeMachinesArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterIds")]
        private List<string>? _clusterIds;

        /// <summary>
        /// The ClusterIds of NodePool IDs.
        /// </summary>
        public List<string> ClusterIds
        {
            get => _clusterIds ?? (_clusterIds = new List<string>());
            set => _clusterIds = value;
        }

        /// <summary>
        /// The ClientToken when successfully created.
        /// </summary>
        [Input("createClientToken")]
        public string? CreateClientToken { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("ips")]
        private List<string>? _ips;

        /// <summary>
        /// The IPs.
        /// </summary>
        public List<string> Ips
        {
            get => _ips ?? (_ips = new List<string>());
            set => _ips = value;
        }

        /// <summary>
        /// The Name of NodePool.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Whether it is necessary to query the node management script.
        /// </summary>
        [Input("needBootstrapScript")]
        public string? NeedBootstrapScript { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("statuses")]
        private List<Inputs.BatchEdgeMachinesStatusArgs>? _statuses;

        /// <summary>
        /// The Status of NodePool.
        /// </summary>
        public List<Inputs.BatchEdgeMachinesStatusArgs> Statuses
        {
            get => _statuses ?? (_statuses = new List<Inputs.BatchEdgeMachinesStatusArgs>());
            set => _statuses = value;
        }

        [Input("zoneIds")]
        private List<string>? _zoneIds;

        /// <summary>
        /// The Zone Ids.
        /// </summary>
        public List<string> ZoneIds
        {
            get => _zoneIds ?? (_zoneIds = new List<string>());
            set => _zoneIds = value;
        }

        public BatchEdgeMachinesArgs()
        {
        }
        public static new BatchEdgeMachinesArgs Empty => new BatchEdgeMachinesArgs();
    }

    public sealed class BatchEdgeMachinesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterIds")]
        private InputList<string>? _clusterIds;

        /// <summary>
        /// The ClusterIds of NodePool IDs.
        /// </summary>
        public InputList<string> ClusterIds
        {
            get => _clusterIds ?? (_clusterIds = new InputList<string>());
            set => _clusterIds = value;
        }

        /// <summary>
        /// The ClientToken when successfully created.
        /// </summary>
        [Input("createClientToken")]
        public Input<string>? CreateClientToken { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("ips")]
        private InputList<string>? _ips;

        /// <summary>
        /// The IPs.
        /// </summary>
        public InputList<string> Ips
        {
            get => _ips ?? (_ips = new InputList<string>());
            set => _ips = value;
        }

        /// <summary>
        /// The Name of NodePool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether it is necessary to query the node management script.
        /// </summary>
        [Input("needBootstrapScript")]
        public Input<string>? NeedBootstrapScript { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("statuses")]
        private InputList<Inputs.BatchEdgeMachinesStatusInputArgs>? _statuses;

        /// <summary>
        /// The Status of NodePool.
        /// </summary>
        public InputList<Inputs.BatchEdgeMachinesStatusInputArgs> Statuses
        {
            get => _statuses ?? (_statuses = new InputList<Inputs.BatchEdgeMachinesStatusInputArgs>());
            set => _statuses = value;
        }

        [Input("zoneIds")]
        private InputList<string>? _zoneIds;

        /// <summary>
        /// The Zone Ids.
        /// </summary>
        public InputList<string> ZoneIds
        {
            get => _zoneIds ?? (_zoneIds = new InputList<string>());
            set => _zoneIds = value;
        }

        public BatchEdgeMachinesInvokeArgs()
        {
        }
        public static new BatchEdgeMachinesInvokeArgs Empty => new BatchEdgeMachinesInvokeArgs();
    }


    [OutputType]
    public sealed class BatchEdgeMachinesResult
    {
        public readonly ImmutableArray<string> ClusterIds;
        /// <summary>
        /// The ClientToken when successfully created.
        /// </summary>
        public readonly string? CreateClientToken;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<string> Ips;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.BatchEdgeMachinesMachineResult> Machines;
        /// <summary>
        /// The Name of NodePool.
        /// </summary>
        public readonly string? Name;
        public readonly string? NeedBootstrapScript;
        public readonly string? OutputFile;
        public readonly ImmutableArray<Outputs.BatchEdgeMachinesStatusResult> Statuses;
        /// <summary>
        /// The total count of query.
        /// </summary>
        public readonly int TotalCount;
        public readonly ImmutableArray<string> ZoneIds;

        [OutputConstructor]
        private BatchEdgeMachinesResult(
            ImmutableArray<string> clusterIds,

            string? createClientToken,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<string> ips,

            ImmutableArray<Outputs.BatchEdgeMachinesMachineResult> machines,

            string? name,

            string? needBootstrapScript,

            string? outputFile,

            ImmutableArray<Outputs.BatchEdgeMachinesStatusResult> statuses,

            int totalCount,

            ImmutableArray<string> zoneIds)
        {
            ClusterIds = clusterIds;
            CreateClientToken = createClientToken;
            Id = id;
            Ids = ids;
            Ips = ips;
            Machines = machines;
            Name = name;
            NeedBootstrapScript = needBootstrapScript;
            OutputFile = outputFile;
            Statuses = statuses;
            TotalCount = totalCount;
            ZoneIds = zoneIds;
        }
    }
}
