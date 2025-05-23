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

__all__ = [
    'GetServerGroupServersResult',
    'AwaitableGetServerGroupServersResult',
    'get_server_group_servers',
    'get_server_group_servers_output',
]

@pulumi.output_type
class GetServerGroupServersResult:
    """
    A collection of values returned by getServerGroupServers.
    """
    def __init__(__self__, id=None, ids=None, name_regex=None, output_file=None, server_group_id=None, servers=None, total_count=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if server_group_id and not isinstance(server_group_id, str):
            raise TypeError("Expected argument 'server_group_id' to be a str")
        pulumi.set(__self__, "server_group_id", server_group_id)
        if servers and not isinstance(servers, list):
            raise TypeError("Expected argument 'servers' to be a list")
        pulumi.set(__self__, "servers", servers)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

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
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="serverGroupId")
    def server_group_id(self) -> str:
        return pulumi.get(self, "server_group_id")

    @property
    @pulumi.getter
    def servers(self) -> Sequence['outputs.GetServerGroupServersServerResult']:
        """
        The server list of ServerGroup.
        """
        return pulumi.get(self, "servers")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of ServerGroupServer query.
        """
        return pulumi.get(self, "total_count")


class AwaitableGetServerGroupServersResult(GetServerGroupServersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerGroupServersResult(
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            output_file=self.output_file,
            server_group_id=self.server_group_id,
            servers=self.servers,
            total_count=self.total_count)


def get_server_group_servers(ids: Optional[Sequence[str]] = None,
                             name_regex: Optional[str] = None,
                             output_file: Optional[str] = None,
                             server_group_id: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServerGroupServersResult:
    """
    Use this data source to query detailed information of server group servers
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_zones = volcengine.ecs.get_zones()
    foo_vpc = volcengine.vpc.Vpc("fooVpc",
        vpc_name="acc-test-vpc",
        cidr_block="172.16.0.0/16")
    foo_subnet = volcengine.vpc.Subnet("fooSubnet",
        subnet_name="acc-test-subnet",
        cidr_block="172.16.0.0/24",
        zone_id=foo_zones.zones[0].id,
        vpc_id=foo_vpc.id)
    foo_clb = volcengine.clb.Clb("fooClb",
        type="public",
        subnet_id=foo_subnet.id,
        load_balancer_spec="small_1",
        description="acc0Demo",
        load_balancer_name="acc-test-create",
        eip_billing_config=volcengine.clb.ClbEipBillingConfigArgs(
            isp="BGP",
            eip_billing_type="PostPaidByBandwidth",
            bandwidth=1,
        ))
    foo_server_group = volcengine.clb.ServerGroup("fooServerGroup",
        load_balancer_id=foo_clb.id,
        server_group_name="acc-test-create",
        description="hello demo11")
    foo_security_group = volcengine.vpc.SecurityGroup("fooSecurityGroup",
        vpc_id=foo_vpc.id,
        security_group_name="acc-test-security-group")
    foo_instance = volcengine.ecs.Instance("fooInstance",
        image_id="image-ycjwwciuzy5pkh54xx8f",
        instance_type="ecs.c3i.large",
        instance_name="acc-test-ecs-name",
        password="93f0cb0614Aab12",
        instance_charge_type="PostPaid",
        system_volume_type="ESSD_PL0",
        system_volume_size=40,
        subnet_id=foo_subnet.id,
        security_group_ids=[foo_security_group.id])
    foo_server_group_server = volcengine.clb.ServerGroupServer("fooServerGroupServer",
        server_group_id=foo_server_group.id,
        instance_id=foo_instance.id,
        type="ecs",
        weight=100,
        port=80,
        description="This is a acc test server")
    foo_server_group_servers = volcengine.clb.get_server_group_servers_output(ids=[pulumi.Output.all(foo_server_group_server.id.apply(lambda id: id.split(":")), len(foo_server_group_server.id.apply(lambda id: id.split(":")))).apply(lambda split, length: split[length - 1])],
        server_group_id=foo_server_group.id)
    ```


    :param Sequence[str] ids: The list of ServerGroupServer IDs.
    :param str name_regex: A Name Regex of ServerGroupServer.
    :param str output_file: File name where to save data source results.
    :param str server_group_id: The ID of the ServerGroup.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['serverGroupId'] = server_group_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:clb/getServerGroupServers:getServerGroupServers', __args__, opts=opts, typ=GetServerGroupServersResult).value

    return AwaitableGetServerGroupServersResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        server_group_id=pulumi.get(__ret__, 'server_group_id'),
        servers=pulumi.get(__ret__, 'servers'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(get_server_group_servers)
def get_server_group_servers_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                    name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                    output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                    server_group_id: Optional[pulumi.Input[str]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServerGroupServersResult]:
    """
    Use this data source to query detailed information of server group servers
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_zones = volcengine.ecs.get_zones()
    foo_vpc = volcengine.vpc.Vpc("fooVpc",
        vpc_name="acc-test-vpc",
        cidr_block="172.16.0.0/16")
    foo_subnet = volcengine.vpc.Subnet("fooSubnet",
        subnet_name="acc-test-subnet",
        cidr_block="172.16.0.0/24",
        zone_id=foo_zones.zones[0].id,
        vpc_id=foo_vpc.id)
    foo_clb = volcengine.clb.Clb("fooClb",
        type="public",
        subnet_id=foo_subnet.id,
        load_balancer_spec="small_1",
        description="acc0Demo",
        load_balancer_name="acc-test-create",
        eip_billing_config=volcengine.clb.ClbEipBillingConfigArgs(
            isp="BGP",
            eip_billing_type="PostPaidByBandwidth",
            bandwidth=1,
        ))
    foo_server_group = volcengine.clb.ServerGroup("fooServerGroup",
        load_balancer_id=foo_clb.id,
        server_group_name="acc-test-create",
        description="hello demo11")
    foo_security_group = volcengine.vpc.SecurityGroup("fooSecurityGroup",
        vpc_id=foo_vpc.id,
        security_group_name="acc-test-security-group")
    foo_instance = volcengine.ecs.Instance("fooInstance",
        image_id="image-ycjwwciuzy5pkh54xx8f",
        instance_type="ecs.c3i.large",
        instance_name="acc-test-ecs-name",
        password="93f0cb0614Aab12",
        instance_charge_type="PostPaid",
        system_volume_type="ESSD_PL0",
        system_volume_size=40,
        subnet_id=foo_subnet.id,
        security_group_ids=[foo_security_group.id])
    foo_server_group_server = volcengine.clb.ServerGroupServer("fooServerGroupServer",
        server_group_id=foo_server_group.id,
        instance_id=foo_instance.id,
        type="ecs",
        weight=100,
        port=80,
        description="This is a acc test server")
    foo_server_group_servers = volcengine.clb.get_server_group_servers_output(ids=[pulumi.Output.all(foo_server_group_server.id.apply(lambda id: id.split(":")), len(foo_server_group_server.id.apply(lambda id: id.split(":")))).apply(lambda split, length: split[length - 1])],
        server_group_id=foo_server_group.id)
    ```


    :param Sequence[str] ids: The list of ServerGroupServer IDs.
    :param str name_regex: A Name Regex of ServerGroupServer.
    :param str output_file: File name where to save data source results.
    :param str server_group_id: The ID of the ServerGroup.
    """
    ...
