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
    'SslVpnServersResult',
    'AwaitableSslVpnServersResult',
    'ssl_vpn_servers',
    'ssl_vpn_servers_output',
]

@pulumi.output_type
class SslVpnServersResult:
    """
    A collection of values returned by SslVpnServers.
    """
    def __init__(__self__, id=None, ids=None, output_file=None, ssl_vpn_server_name=None, ssl_vpn_servers=None, total_count=None, vpn_gateway_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if ssl_vpn_server_name and not isinstance(ssl_vpn_server_name, str):
            raise TypeError("Expected argument 'ssl_vpn_server_name' to be a str")
        pulumi.set(__self__, "ssl_vpn_server_name", ssl_vpn_server_name)
        if ssl_vpn_servers and not isinstance(ssl_vpn_servers, list):
            raise TypeError("Expected argument 'ssl_vpn_servers' to be a list")
        pulumi.set(__self__, "ssl_vpn_servers", ssl_vpn_servers)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if vpn_gateway_id and not isinstance(vpn_gateway_id, str):
            raise TypeError("Expected argument 'vpn_gateway_id' to be a str")
        pulumi.set(__self__, "vpn_gateway_id", vpn_gateway_id)

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
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="sslVpnServerName")
    def ssl_vpn_server_name(self) -> Optional[str]:
        """
        The name of the SSL server.
        """
        return pulumi.get(self, "ssl_vpn_server_name")

    @property
    @pulumi.getter(name="sslVpnServers")
    def ssl_vpn_servers(self) -> Sequence['outputs.SslVpnServersSslVpnServerResult']:
        """
        List of SSL VPN servers.
        """
        return pulumi.get(self, "ssl_vpn_servers")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of SSL VPN server query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> Optional[str]:
        """
        The vpn gateway id.
        """
        return pulumi.get(self, "vpn_gateway_id")


class AwaitableSslVpnServersResult(SslVpnServersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return SslVpnServersResult(
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            ssl_vpn_server_name=self.ssl_vpn_server_name,
            ssl_vpn_servers=self.ssl_vpn_servers,
            total_count=self.total_count,
            vpn_gateway_id=self.vpn_gateway_id)


def ssl_vpn_servers(ids: Optional[Sequence[str]] = None,
                    output_file: Optional[str] = None,
                    ssl_vpn_server_name: Optional[str] = None,
                    vpn_gateway_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableSslVpnServersResult:
    """
    Use this data source to query detailed information of ssl vpn servers
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_zones = volcengine.ecs.zones()
    foo_vpc = volcengine.vpc.Vpc("fooVpc",
        vpc_name="acc-test-vpc",
        cidr_block="172.16.0.0/16")
    foo_subnet = volcengine.vpc.Subnet("fooSubnet",
        subnet_name="acc-test-subnet",
        cidr_block="172.16.0.0/24",
        zone_id=foo_zones.zones[0].id,
        vpc_id=foo_vpc.id)
    foo_gateway = volcengine.vpn.Gateway("fooGateway",
        vpc_id=foo_vpc.id,
        subnet_id=foo_subnet.id,
        bandwidth=5,
        vpn_gateway_name="acc-test1",
        description="acc-test1",
        period=7,
        project_name="default",
        ssl_enabled=True,
        ssl_max_connections=5)
    foo_ssl_vpn_server = volcengine.vpn.SslVpnServer("fooSslVpnServer",
        vpn_gateway_id=foo_gateway.id,
        local_subnets=[foo_subnet.cidr_block],
        client_ip_pool="172.16.2.0/24",
        ssl_vpn_server_name="acc-test-ssl",
        description="acc-test",
        protocol="UDP",
        cipher="AES-128-CBC",
        auth="SHA1",
        compress=True)
    foo_ssl_vpn_servers = volcengine.vpn.ssl_vpn_servers_output(ids=[foo_ssl_vpn_server.id])
    ```


    :param Sequence[str] ids: The ids list.
    :param str output_file: File name where to save data source results.
    :param str ssl_vpn_server_name: The name of the ssl vpn server.
    :param str vpn_gateway_id: The id of the vpn gateway.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    __args__['sslVpnServerName'] = ssl_vpn_server_name
    __args__['vpnGatewayId'] = vpn_gateway_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:vpn/sslVpnServers:SslVpnServers', __args__, opts=opts, typ=SslVpnServersResult).value

    return AwaitableSslVpnServersResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        output_file=pulumi.get(__ret__, 'output_file'),
        ssl_vpn_server_name=pulumi.get(__ret__, 'ssl_vpn_server_name'),
        ssl_vpn_servers=pulumi.get(__ret__, 'ssl_vpn_servers'),
        total_count=pulumi.get(__ret__, 'total_count'),
        vpn_gateway_id=pulumi.get(__ret__, 'vpn_gateway_id'))


@_utilities.lift_output_func(ssl_vpn_servers)
def ssl_vpn_servers_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           ssl_vpn_server_name: Optional[pulumi.Input[Optional[str]]] = None,
                           vpn_gateway_id: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[SslVpnServersResult]:
    """
    Use this data source to query detailed information of ssl vpn servers
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_zones = volcengine.ecs.zones()
    foo_vpc = volcengine.vpc.Vpc("fooVpc",
        vpc_name="acc-test-vpc",
        cidr_block="172.16.0.0/16")
    foo_subnet = volcengine.vpc.Subnet("fooSubnet",
        subnet_name="acc-test-subnet",
        cidr_block="172.16.0.0/24",
        zone_id=foo_zones.zones[0].id,
        vpc_id=foo_vpc.id)
    foo_gateway = volcengine.vpn.Gateway("fooGateway",
        vpc_id=foo_vpc.id,
        subnet_id=foo_subnet.id,
        bandwidth=5,
        vpn_gateway_name="acc-test1",
        description="acc-test1",
        period=7,
        project_name="default",
        ssl_enabled=True,
        ssl_max_connections=5)
    foo_ssl_vpn_server = volcengine.vpn.SslVpnServer("fooSslVpnServer",
        vpn_gateway_id=foo_gateway.id,
        local_subnets=[foo_subnet.cidr_block],
        client_ip_pool="172.16.2.0/24",
        ssl_vpn_server_name="acc-test-ssl",
        description="acc-test",
        protocol="UDP",
        cipher="AES-128-CBC",
        auth="SHA1",
        compress=True)
    foo_ssl_vpn_servers = volcengine.vpn.ssl_vpn_servers_output(ids=[foo_ssl_vpn_server.id])
    ```


    :param Sequence[str] ids: The ids list.
    :param str output_file: File name where to save data source results.
    :param str ssl_vpn_server_name: The name of the ssl vpn server.
    :param str vpn_gateway_id: The id of the vpn gateway.
    """
    ...