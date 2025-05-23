# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'BatchEdgeMachinesResult',
    'AwaitableBatchEdgeMachinesResult',
    'batch_edge_machines',
    'batch_edge_machines_output',
]

warnings.warn("""volcengine.veecp.BatchEdgeMachines has been deprecated in favor of volcengine.veecp.getBatchEdgeMachines""", DeprecationWarning)

@pulumi.output_type
class BatchEdgeMachinesResult:
    """
    A collection of values returned by BatchEdgeMachines.
    """
    def __init__(__self__, cluster_ids=None, create_client_token=None, id=None, ids=None, ips=None, machines=None, name=None, need_bootstrap_script=None, output_file=None, statuses=None, total_count=None, zone_ids=None):
        if cluster_ids and not isinstance(cluster_ids, list):
            raise TypeError("Expected argument 'cluster_ids' to be a list")
        pulumi.set(__self__, "cluster_ids", cluster_ids)
        if create_client_token and not isinstance(create_client_token, str):
            raise TypeError("Expected argument 'create_client_token' to be a str")
        pulumi.set(__self__, "create_client_token", create_client_token)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if ips and not isinstance(ips, list):
            raise TypeError("Expected argument 'ips' to be a list")
        pulumi.set(__self__, "ips", ips)
        if machines and not isinstance(machines, list):
            raise TypeError("Expected argument 'machines' to be a list")
        pulumi.set(__self__, "machines", machines)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if need_bootstrap_script and not isinstance(need_bootstrap_script, str):
            raise TypeError("Expected argument 'need_bootstrap_script' to be a str")
        pulumi.set(__self__, "need_bootstrap_script", need_bootstrap_script)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if zone_ids and not isinstance(zone_ids, list):
            raise TypeError("Expected argument 'zone_ids' to be a list")
        pulumi.set(__self__, "zone_ids", zone_ids)

    @property
    @pulumi.getter(name="clusterIds")
    def cluster_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "cluster_ids")

    @property
    @pulumi.getter(name="createClientToken")
    def create_client_token(self) -> Optional[str]:
        """
        The ClientToken when successfully created.
        """
        return pulumi.get(self, "create_client_token")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter
    def ips(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "ips")

    @property
    @pulumi.getter
    def machines(self) -> Sequence['outputs.BatchEdgeMachinesMachineResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "machines")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The Name of NodePool.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="needBootstrapScript")
    def need_bootstrap_script(self) -> Optional[str]:
        return pulumi.get(self, "need_bootstrap_script")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def statuses(self) -> Optional[Sequence['outputs.BatchEdgeMachinesStatusResult']]:
        return pulumi.get(self, "statuses")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="zoneIds")
    def zone_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "zone_ids")


class AwaitableBatchEdgeMachinesResult(BatchEdgeMachinesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return BatchEdgeMachinesResult(
            cluster_ids=self.cluster_ids,
            create_client_token=self.create_client_token,
            id=self.id,
            ids=self.ids,
            ips=self.ips,
            machines=self.machines,
            name=self.name,
            need_bootstrap_script=self.need_bootstrap_script,
            output_file=self.output_file,
            statuses=self.statuses,
            total_count=self.total_count,
            zone_ids=self.zone_ids)


def batch_edge_machines(cluster_ids: Optional[Sequence[str]] = None,
                        create_client_token: Optional[str] = None,
                        ids: Optional[Sequence[str]] = None,
                        ips: Optional[Sequence[str]] = None,
                        name: Optional[str] = None,
                        need_bootstrap_script: Optional[str] = None,
                        output_file: Optional[str] = None,
                        statuses: Optional[Sequence[pulumi.InputType['BatchEdgeMachinesStatusArgs']]] = None,
                        zone_ids: Optional[Sequence[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableBatchEdgeMachinesResult:
    """
    Use this data source to query detailed information of veecp batch edge machines
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_batch_edge_machine = volcengine.veecp.BatchEdgeMachine("fooBatchEdgeMachine",
        cluster_id="ccvd7mte6t101fno98u60",
        node_pool_id="pcvd90uacnsr73g6bjic0",
        ttl_hours=1)
    foo_batch_edge_machines = volcengine.veecp.get_batch_edge_machines_output(cluster_ids=[foo_batch_edge_machine.cluster_id],
        ids=[foo_batch_edge_machine.id])
    ```


    :param Sequence[str] cluster_ids: The ClusterIds of NodePool IDs.
    :param str create_client_token: The ClientToken when successfully created.
    :param Sequence[str] ids: A list of IDs.
    :param Sequence[str] ips: The IPs.
    :param str name: The Name of NodePool.
    :param str need_bootstrap_script: Whether it is necessary to query the node management script.
    :param str output_file: File name where to save data source results.
    :param Sequence[pulumi.InputType['BatchEdgeMachinesStatusArgs']] statuses: The Status of NodePool.
    :param Sequence[str] zone_ids: The Zone Ids.
    """
    pulumi.log.warn("""batch_edge_machines is deprecated: volcengine.veecp.BatchEdgeMachines has been deprecated in favor of volcengine.veecp.getBatchEdgeMachines""")
    __args__ = dict()
    __args__['clusterIds'] = cluster_ids
    __args__['createClientToken'] = create_client_token
    __args__['ids'] = ids
    __args__['ips'] = ips
    __args__['name'] = name
    __args__['needBootstrapScript'] = need_bootstrap_script
    __args__['outputFile'] = output_file
    __args__['statuses'] = statuses
    __args__['zoneIds'] = zone_ids
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:veecp/batchEdgeMachines:BatchEdgeMachines', __args__, opts=opts, typ=BatchEdgeMachinesResult).value

    return AwaitableBatchEdgeMachinesResult(
        cluster_ids=pulumi.get(__ret__, 'cluster_ids'),
        create_client_token=pulumi.get(__ret__, 'create_client_token'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        ips=pulumi.get(__ret__, 'ips'),
        machines=pulumi.get(__ret__, 'machines'),
        name=pulumi.get(__ret__, 'name'),
        need_bootstrap_script=pulumi.get(__ret__, 'need_bootstrap_script'),
        output_file=pulumi.get(__ret__, 'output_file'),
        statuses=pulumi.get(__ret__, 'statuses'),
        total_count=pulumi.get(__ret__, 'total_count'),
        zone_ids=pulumi.get(__ret__, 'zone_ids'))


@_utilities.lift_output_func(batch_edge_machines)
def batch_edge_machines_output(cluster_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                               create_client_token: Optional[pulumi.Input[Optional[str]]] = None,
                               ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                               ips: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                               name: Optional[pulumi.Input[Optional[str]]] = None,
                               need_bootstrap_script: Optional[pulumi.Input[Optional[str]]] = None,
                               output_file: Optional[pulumi.Input[Optional[str]]] = None,
                               statuses: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['BatchEdgeMachinesStatusArgs']]]]] = None,
                               zone_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[BatchEdgeMachinesResult]:
    """
    Use this data source to query detailed information of veecp batch edge machines
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_batch_edge_machine = volcengine.veecp.BatchEdgeMachine("fooBatchEdgeMachine",
        cluster_id="ccvd7mte6t101fno98u60",
        node_pool_id="pcvd90uacnsr73g6bjic0",
        ttl_hours=1)
    foo_batch_edge_machines = volcengine.veecp.get_batch_edge_machines_output(cluster_ids=[foo_batch_edge_machine.cluster_id],
        ids=[foo_batch_edge_machine.id])
    ```


    :param Sequence[str] cluster_ids: The ClusterIds of NodePool IDs.
    :param str create_client_token: The ClientToken when successfully created.
    :param Sequence[str] ids: A list of IDs.
    :param Sequence[str] ips: The IPs.
    :param str name: The Name of NodePool.
    :param str need_bootstrap_script: Whether it is necessary to query the node management script.
    :param str output_file: File name where to save data source results.
    :param Sequence[pulumi.InputType['BatchEdgeMachinesStatusArgs']] statuses: The Status of NodePool.
    :param Sequence[str] zone_ids: The Zone Ids.
    """
    pulumi.log.warn("""batch_edge_machines is deprecated: volcengine.veecp.BatchEdgeMachines has been deprecated in favor of volcengine.veecp.getBatchEdgeMachines""")
    ...
