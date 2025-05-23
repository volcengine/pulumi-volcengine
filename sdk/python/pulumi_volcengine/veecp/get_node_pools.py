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
    'GetNodePoolsResult',
    'AwaitableGetNodePoolsResult',
    'get_node_pools',
    'get_node_pools_output',
]

@pulumi.output_type
class GetNodePoolsResult:
    """
    A collection of values returned by getNodePools.
    """
    def __init__(__self__, auto_scaling_enabled=None, cluster_id=None, cluster_ids=None, create_client_token=None, id=None, ids=None, name=None, name_regex=None, node_pools=None, output_file=None, statuses=None, total_count=None, update_client_token=None):
        if auto_scaling_enabled and not isinstance(auto_scaling_enabled, bool):
            raise TypeError("Expected argument 'auto_scaling_enabled' to be a bool")
        pulumi.set(__self__, "auto_scaling_enabled", auto_scaling_enabled)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
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
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if node_pools and not isinstance(node_pools, list):
            raise TypeError("Expected argument 'node_pools' to be a list")
        pulumi.set(__self__, "node_pools", node_pools)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if update_client_token and not isinstance(update_client_token, str):
            raise TypeError("Expected argument 'update_client_token' to be a str")
        pulumi.set(__self__, "update_client_token", update_client_token)

    @property
    @pulumi.getter(name="autoScalingEnabled")
    def auto_scaling_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "auto_scaling_enabled")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[str]:
        """
        The ClusterId of NodePool.
        """
        return pulumi.get(self, "cluster_id")

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
    def name(self) -> Optional[str]:
        """
        The Name of NodePool.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="nodePools")
    def node_pools(self) -> Sequence['outputs.GetNodePoolsNodePoolResult']:
        """
        The collection of NodePools query.
        """
        return pulumi.get(self, "node_pools")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def statuses(self) -> Optional[Sequence['outputs.GetNodePoolsStatusResult']]:
        return pulumi.get(self, "statuses")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="updateClientToken")
    def update_client_token(self) -> Optional[str]:
        """
        The ClientToken when last update was successful.
        """
        return pulumi.get(self, "update_client_token")


class AwaitableGetNodePoolsResult(GetNodePoolsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNodePoolsResult(
            auto_scaling_enabled=self.auto_scaling_enabled,
            cluster_id=self.cluster_id,
            cluster_ids=self.cluster_ids,
            create_client_token=self.create_client_token,
            id=self.id,
            ids=self.ids,
            name=self.name,
            name_regex=self.name_regex,
            node_pools=self.node_pools,
            output_file=self.output_file,
            statuses=self.statuses,
            total_count=self.total_count,
            update_client_token=self.update_client_token)


def get_node_pools(auto_scaling_enabled: Optional[bool] = None,
                   cluster_id: Optional[str] = None,
                   cluster_ids: Optional[Sequence[str]] = None,
                   create_client_token: Optional[str] = None,
                   ids: Optional[Sequence[str]] = None,
                   name: Optional[str] = None,
                   name_regex: Optional[str] = None,
                   output_file: Optional[str] = None,
                   statuses: Optional[Sequence[pulumi.InputType['GetNodePoolsStatusArgs']]] = None,
                   update_client_token: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNodePoolsResult:
    """
    Use this data source to query detailed information of veecp node pools
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_zones = volcengine.ecs.get_zones()
    foo_vpc = volcengine.vpc.Vpc("fooVpc",
        vpc_name="acc-test-project1",
        cidr_block="172.16.0.0/16")
    foo_subnet = volcengine.vpc.Subnet("fooSubnet",
        subnet_name="acc-subnet-test-2",
        cidr_block="172.16.0.0/24",
        zone_id=foo_zones.zones[0].id,
        vpc_id=foo_vpc.id)
    foo_security_group = volcengine.vpc.SecurityGroup("fooSecurityGroup",
        vpc_id=foo_vpc.id,
        security_group_name="acc-test-security-group2")
    foo_cluster = volcengine.veecp.Cluster("fooCluster",
        description="created by terraform",
        delete_protection_enabled=False,
        profile="Edge",
        cluster_config=volcengine.veecp.ClusterClusterConfigArgs(
            subnet_ids=[foo_subnet.id],
            api_server_public_access_enabled=True,
            api_server_public_access_config=volcengine.veecp.ClusterClusterConfigApiServerPublicAccessConfigArgs(
                public_access_network_config=volcengine.veecp.ClusterClusterConfigApiServerPublicAccessConfigPublicAccessNetworkConfigArgs(
                    billing_type="PostPaidByBandwidth",
                    bandwidth=1,
                ),
            ),
            resource_public_access_default_enabled=True,
        ),
        pods_config=volcengine.veecp.ClusterPodsConfigArgs(
            pod_network_mode="Flannel",
            flannel_config=volcengine.veecp.ClusterPodsConfigFlannelConfigArgs(
                pod_cidrs=["172.22.224.0/20"],
                max_pods_per_node=64,
            ),
        ),
        services_config=volcengine.veecp.ClusterServicesConfigArgs(
            service_cidrsv4s=["172.30.0.0/18"],
        ))
    foo_node_pool = volcengine.veecp.NodePool("fooNodePool",
        cluster_id=foo_cluster.id,
        client_token="FGAHIxa23412FGAIOHioj",
        auto_scaling=volcengine.veecp.NodePoolAutoScalingArgs(
            enabled=True,
            min_replicas=0,
            max_replicas=5,
            desired_replicas=0,
            priority=5,
            subnet_policy="ZoneBalance",
        ),
        node_config=volcengine.veecp.NodePoolNodeConfigArgs(
            instance_type_ids=["ecs.c1ie.xlarge"],
            subnet_ids=[foo_subnet.id],
            image_id="",
            system_volume=volcengine.veecp.NodePoolNodeConfigSystemVolumeArgs(
                type="ESSD_PL0",
                size=80,
            ),
            data_volumes=[
                volcengine.veecp.NodePoolNodeConfigDataVolumeArgs(
                    type="ESSD_PL0",
                    size=80,
                    mount_point="/tf1",
                ),
                volcengine.veecp.NodePoolNodeConfigDataVolumeArgs(
                    type="ESSD_PL0",
                    size=60,
                    mount_point="/tf2",
                ),
            ],
            initialize_script="ZWNobyBoZWxsbyB0ZXJyYWZvcm0h",
            security=volcengine.veecp.NodePoolNodeConfigSecurityArgs(
                login=volcengine.veecp.NodePoolNodeConfigSecurityLoginArgs(
                    password="UHdkMTIzNDU2",
                ),
                security_strategies=["Hids"],
                security_group_ids=[foo_security_group.id],
            ),
            additional_container_storage_enabled=False,
            instance_charge_type="PostPaid",
            name_prefix="acc-test",
            ecs_tags=[volcengine.veecp.NodePoolNodeConfigEcsTagArgs(
                key="ecs_k1",
                value="ecs_v1",
            )],
        ),
        kubernetes_config=volcengine.veecp.NodePoolKubernetesConfigArgs(
            labels=[volcengine.veecp.NodePoolKubernetesConfigLabelArgs(
                key="label1",
                value="value1",
            )],
            taints=[volcengine.veecp.NodePoolKubernetesConfigTaintArgs(
                key="taint-key/node-type",
                value="taint-value",
                effect="NoSchedule",
            )],
            cordon=True,
        ))
    foo_node_pools = volcengine.veecp.get_node_pools_output(ids=[foo_node_pool.id])
    ```


    :param bool auto_scaling_enabled: Is enabled of AutoScaling.
    :param str cluster_id: The ClusterId of NodePool.
    :param Sequence[str] cluster_ids: The ClusterIds of NodePool IDs.
    :param str create_client_token: The ClientToken when successfully created.
    :param Sequence[str] ids: The IDs of NodePool.
    :param str name: The Name of NodePool.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    :param Sequence[pulumi.InputType['GetNodePoolsStatusArgs']] statuses: The Status of NodePool.
    :param str update_client_token: The ClientToken when last update was successful.
    """
    __args__ = dict()
    __args__['autoScalingEnabled'] = auto_scaling_enabled
    __args__['clusterId'] = cluster_id
    __args__['clusterIds'] = cluster_ids
    __args__['createClientToken'] = create_client_token
    __args__['ids'] = ids
    __args__['name'] = name
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['statuses'] = statuses
    __args__['updateClientToken'] = update_client_token
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:veecp/getNodePools:getNodePools', __args__, opts=opts, typ=GetNodePoolsResult).value

    return AwaitableGetNodePoolsResult(
        auto_scaling_enabled=pulumi.get(__ret__, 'auto_scaling_enabled'),
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        cluster_ids=pulumi.get(__ret__, 'cluster_ids'),
        create_client_token=pulumi.get(__ret__, 'create_client_token'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name=pulumi.get(__ret__, 'name'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        node_pools=pulumi.get(__ret__, 'node_pools'),
        output_file=pulumi.get(__ret__, 'output_file'),
        statuses=pulumi.get(__ret__, 'statuses'),
        total_count=pulumi.get(__ret__, 'total_count'),
        update_client_token=pulumi.get(__ret__, 'update_client_token'))


@_utilities.lift_output_func(get_node_pools)
def get_node_pools_output(auto_scaling_enabled: Optional[pulumi.Input[Optional[bool]]] = None,
                          cluster_id: Optional[pulumi.Input[Optional[str]]] = None,
                          cluster_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                          create_client_token: Optional[pulumi.Input[Optional[str]]] = None,
                          ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                          name: Optional[pulumi.Input[Optional[str]]] = None,
                          name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                          output_file: Optional[pulumi.Input[Optional[str]]] = None,
                          statuses: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetNodePoolsStatusArgs']]]]] = None,
                          update_client_token: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNodePoolsResult]:
    """
    Use this data source to query detailed information of veecp node pools
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_zones = volcengine.ecs.get_zones()
    foo_vpc = volcengine.vpc.Vpc("fooVpc",
        vpc_name="acc-test-project1",
        cidr_block="172.16.0.0/16")
    foo_subnet = volcengine.vpc.Subnet("fooSubnet",
        subnet_name="acc-subnet-test-2",
        cidr_block="172.16.0.0/24",
        zone_id=foo_zones.zones[0].id,
        vpc_id=foo_vpc.id)
    foo_security_group = volcengine.vpc.SecurityGroup("fooSecurityGroup",
        vpc_id=foo_vpc.id,
        security_group_name="acc-test-security-group2")
    foo_cluster = volcengine.veecp.Cluster("fooCluster",
        description="created by terraform",
        delete_protection_enabled=False,
        profile="Edge",
        cluster_config=volcengine.veecp.ClusterClusterConfigArgs(
            subnet_ids=[foo_subnet.id],
            api_server_public_access_enabled=True,
            api_server_public_access_config=volcengine.veecp.ClusterClusterConfigApiServerPublicAccessConfigArgs(
                public_access_network_config=volcengine.veecp.ClusterClusterConfigApiServerPublicAccessConfigPublicAccessNetworkConfigArgs(
                    billing_type="PostPaidByBandwidth",
                    bandwidth=1,
                ),
            ),
            resource_public_access_default_enabled=True,
        ),
        pods_config=volcengine.veecp.ClusterPodsConfigArgs(
            pod_network_mode="Flannel",
            flannel_config=volcengine.veecp.ClusterPodsConfigFlannelConfigArgs(
                pod_cidrs=["172.22.224.0/20"],
                max_pods_per_node=64,
            ),
        ),
        services_config=volcengine.veecp.ClusterServicesConfigArgs(
            service_cidrsv4s=["172.30.0.0/18"],
        ))
    foo_node_pool = volcengine.veecp.NodePool("fooNodePool",
        cluster_id=foo_cluster.id,
        client_token="FGAHIxa23412FGAIOHioj",
        auto_scaling=volcengine.veecp.NodePoolAutoScalingArgs(
            enabled=True,
            min_replicas=0,
            max_replicas=5,
            desired_replicas=0,
            priority=5,
            subnet_policy="ZoneBalance",
        ),
        node_config=volcengine.veecp.NodePoolNodeConfigArgs(
            instance_type_ids=["ecs.c1ie.xlarge"],
            subnet_ids=[foo_subnet.id],
            image_id="",
            system_volume=volcengine.veecp.NodePoolNodeConfigSystemVolumeArgs(
                type="ESSD_PL0",
                size=80,
            ),
            data_volumes=[
                volcengine.veecp.NodePoolNodeConfigDataVolumeArgs(
                    type="ESSD_PL0",
                    size=80,
                    mount_point="/tf1",
                ),
                volcengine.veecp.NodePoolNodeConfigDataVolumeArgs(
                    type="ESSD_PL0",
                    size=60,
                    mount_point="/tf2",
                ),
            ],
            initialize_script="ZWNobyBoZWxsbyB0ZXJyYWZvcm0h",
            security=volcengine.veecp.NodePoolNodeConfigSecurityArgs(
                login=volcengine.veecp.NodePoolNodeConfigSecurityLoginArgs(
                    password="UHdkMTIzNDU2",
                ),
                security_strategies=["Hids"],
                security_group_ids=[foo_security_group.id],
            ),
            additional_container_storage_enabled=False,
            instance_charge_type="PostPaid",
            name_prefix="acc-test",
            ecs_tags=[volcengine.veecp.NodePoolNodeConfigEcsTagArgs(
                key="ecs_k1",
                value="ecs_v1",
            )],
        ),
        kubernetes_config=volcengine.veecp.NodePoolKubernetesConfigArgs(
            labels=[volcengine.veecp.NodePoolKubernetesConfigLabelArgs(
                key="label1",
                value="value1",
            )],
            taints=[volcengine.veecp.NodePoolKubernetesConfigTaintArgs(
                key="taint-key/node-type",
                value="taint-value",
                effect="NoSchedule",
            )],
            cordon=True,
        ))
    foo_node_pools = volcengine.veecp.get_node_pools_output(ids=[foo_node_pool.id])
    ```


    :param bool auto_scaling_enabled: Is enabled of AutoScaling.
    :param str cluster_id: The ClusterId of NodePool.
    :param Sequence[str] cluster_ids: The ClusterIds of NodePool IDs.
    :param str create_client_token: The ClientToken when successfully created.
    :param Sequence[str] ids: The IDs of NodePool.
    :param str name: The Name of NodePool.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    :param Sequence[pulumi.InputType['GetNodePoolsStatusArgs']] statuses: The Status of NodePool.
    :param str update_client_token: The ClientToken when last update was successful.
    """
    ...
